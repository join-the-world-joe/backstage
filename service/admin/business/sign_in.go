package business

import (
	"backstage/common/cache/string/token"
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/protocol/admin"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/utils/bcrypt"
	google_authentictor "backstage/utils/google_authenticator"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"time"
)

func SignIn(ctx context.Context, req *admin.SignInReq, rsp *admin.SignInRsp) error {
	if req.Behavior == 1 { // email
		if len(req.Email) <= 0 {
			log.ErrorF("SigIn failure(%v), err: %v", req.Email, "len(req.Email) <= 0")
			rsp.Code = code.InvalidData
			return nil
		}
		usr, err := user.GetModelByEmail(req.Email)
		if err != nil {
			log.ErrorF("SigIn failure(%v), err: %v", req.Email, err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}
		if usr.Status != 1 {
			log.ErrorF("SigIn failure(%v), err: %v", req.Email, "usr.Status != 1")
			rsp.Code = code.AccountDisable
			return nil
		}
		bytes, err := crypto.RSADecrypt(req.Password)
		if err != nil {
			log.ErrorF("SigIn.crypto.RSADecrypt failure(%v), err: %v", req.Email, err.Error())
			rsp.Code = code.UnsupportedType
			return nil
		}
		ok := bcrypt.PasswordVerify(string(bytes), usr.Password)
		if !ok {
			log.ErrorF("SigIn.bcrypt.PasswordVerify failure(%v), err: %v", req.Email, "!ok")
			rsp.Code = code.LogonFailure
			return nil
		}

		tokenValue, secret, err := encryptToken(usr.Id)
		if err != nil {
			log.ErrorF("SigIn.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("SigIn.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.Code = code.Success
		rsp.UserId = usr.Id
		rsp.Name = usr.Name
		return nil
	} else if req.Behavior == 2 { // mobile
		// check if verification code valid
		err := verification_code.Check(verification_code.SignIn, req.CountryCode, req.PhoneNumber, cast.ToString(req.VerificationCode))
		if err != nil {
			log.ErrorF("SigIn.verification_code.Check failure(+%v-%v), err: %v", req.CountryCode, req.PhoneNumber, err.Error())
			rsp.Code = code.InvalidData
			return nil
		}
		// verification code is ok now
		// get user
		usr, err := user.GetModelByMobile(req.CountryCode, req.PhoneNumber)
		if err != nil {
			log.ErrorF("SigIn.user.GetModelByMobile failure(+%v-%v), err: %v", req.CountryCode, req.PhoneNumber, err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}

		tokenValue, secret, err := encryptToken(usr.Id)
		if err != nil {
			log.ErrorF("SigIn.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("SigIn.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.Name = usr.Name
		rsp.UserId = usr.Id
		rsp.Code = code.Success
		return nil
	} else if req.Behavior == 3 { // token
		b, err := compareToken(req.UserId, req.MemberId, req.VerificationCode)
		if err != nil {
			log.ErrorF("SigIn.compareToken failure(%v), err: %v", req.UserId, err.Error())
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
			log.ErrorF("SigIn.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("SigIn.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.Name = usr.Name
		rsp.UserId = req.UserId
		rsp.Code = code.Success
		return nil
	} else if req.Behavior == 4 { // account
		if len(req.Account) <= 0 {
			log.ErrorF("SigIn failure(%v), err: %v", req.Account, "len(req.Account) <= 0")
			rsp.Code = code.InvalidData
			return nil
		}
		bytes, err := crypto.RSADecrypt(req.Password)
		if err != nil {
			log.ErrorF("SigIn.crypto.RSADecrypt failure(%v), err: %v", req.Account, err.Error())
			rsp.Code = code.UnsupportedType
			return nil
		}
		usr, err := user.GetModelByAccount(req.Account)
		if err != nil {
			log.ErrorF("SigIn.user.GetModelByAccount failure(%v), err: %v", req.Account, err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}

		ok := bcrypt.PasswordVerify(string(bytes), usr.Password)
		if !ok {
			log.ErrorF("SigIn.bcrypt.PasswordVerify failure(%v), err: %v", req.Account, "!ok")
			rsp.Code = code.LogonFailure
			return nil
		}

		tokenValue, secret, err := encryptToken(usr.Id)
		if err != nil {
			log.ErrorF("SigIn.encryptToken failure(%v), err: %v", rsp.UserId, err.Error())
		} else {
			err = token.Create(usr.MemberId, tokenValue)
			if err != nil {
				log.ErrorF("SigIn.token.Create failure(%v), err: %v", rsp.UserId, err.Error())
			}
			rsp.MemberId = usr.MemberId
			rsp.Secret = secret
		}

		rsp.Code = code.Success
		rsp.Name = usr.Name
		rsp.UserId = usr.Id
		return nil
	} else {
		rsp.Code = code.InvalidData
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
