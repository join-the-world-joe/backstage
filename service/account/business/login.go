package business

import (
	"backstage/common/cache/string/token"
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/db/mysql/server/user"
	"backstage/common/protocol/account"
	"backstage/global/crypto"
	"backstage/global/log"
	google_authentictor "backstage/utils/google_authenticator"
	"backstage/validator"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"time"
)

func Login(ctx context.Context, req *account.LoginReq, rsp *account.LoginRsp) error {
	if req.Behavior == 2 { // mobile
		// check if mobile valid
		ok := validator.IsMobileValid(req.CountryCode, req.PhoneNumber)
		if !ok {
			rsp.Code = code.InvalidDataType
			return nil
		}
		// check if verification code valid
		err := verification_code.Check(verification_code.Login, req.CountryCode, req.PhoneNumber, cast.ToString(req.VerificationCode))
		if err != nil {
			log.ErrorF("Login.verification_code.Check failure(+%v-%v), err: %v", req.CountryCode, req.PhoneNumber, err.Error())
			rsp.Code = code.InvalidData
			return nil
		}
		// verification code is ok now
		// get user
		usr, err := user.GetModelByMobile(req.CountryCode, req.PhoneNumber)
		if err != nil {
			log.ErrorF("Login.user.GetModelByMobile failure(+%v-%v), err: %v", req.CountryCode, req.PhoneNumber, err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}

		tokenValue, secret, err := encryptToken(usr.Id)
		if err != nil {
			log.ErrorF("Login.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("Login.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.UserId = usr.Id
		rsp.Code = code.Success
		return nil
	} else if req.Behavior == 3 {
		b, err := compareToken(req.UserId, req.MemberId, req.VerificationCode)
		if err != nil {
			log.ErrorF("Login.compareToken failure(%v), err: %v", req.UserId, err.Error())
			rsp.Code = code.NoUserSessionKey
			return nil
		}
		if !b {
			rsp.Code = code.NotAuthenticated
			return nil
		}
		usr, err := user.GetModelById(req.UserId)
		if err != nil || usr == nil {
			rsp.Code = code.DatabaseFailure
			return nil
		}

		tokenValue, secret, err := encryptToken(usr.Id)
		if err != nil {
			log.ErrorF("Login.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("Login.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.UserId = req.UserId
		rsp.Code = code.Success
		return nil
	} else {
		rsp.Code = code.NotSupported
		return nil
	}
}

func encryptToken(userId int64) (string, string, error) {
	secret := google_authentictor.CreateGoogleSecret(cast.ToString(time.Now().Unix()))
	tkn := &token.Model{
		UserId: userId,
		Secret: secret,
	}
	plainText, err := json.Marshal(tkn)
	if err != nil {
		return "", "", err
	}
	cipherText, err := crypto.RSAEncrypt(plainText)
	if err != nil {
		log.ErrorF("encryptToken.crypto.RSAEncrypt failure(%v), err: %v", userId, err.Error())
		return "", "", err
	}
	return string(cipherText), secret, nil
}

func decryptToken(memberId string) (*token.Model, error) {
	cipherText, err := token.Get(memberId)
	if err != nil {
		return nil, err
	}
	plainText, err := crypto.RSADecrypt([]byte(cipherText))
	if err != nil {
		log.ErrorF("decryptToken.crypto.RSADecrypt failure(%v), err: %v", memberId, err.Error())
		return nil, err
	}
	model := &token.Model{}
	err = json.Unmarshal(plainText, model)
	if err != nil {
		log.ErrorF("decryptToken.json.Unmarshal failure(%v), err: %v", memberId, err.Error())
		return nil, err
	}
	return model, nil
}

func compareToken(userId int64, memberId string, code int32) (bool, error) {
	tkn, err := decryptToken(memberId)
	if err != nil {
		return false, err
	}
	calcCode, _ := google_authentictor.GetGoogleCode(tkn.Secret)
	if calcCode != code {
		return false, nil
	}
	if tkn.UserId != userId {
		return false, nil
	}
	return true, nil
}
