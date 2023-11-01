package business

import (
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/service/backend"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/utils/bcrypt"
	"context"
	"github.com/spf13/cast"
)

func SignIn(ctx context.Context, req *backend.SignInReq, rsp *backend.SignInRsp) error {
	if len(req.VerificationCode) > 0 {
		// check if verification code valid
		err := verification_code.Check(verification_code.SignIn, req.CountryCode, req.PhoneNumber, req.VerificationCode)
		if err != nil {
			log.Error("SigIn.verification_code.Check failure, err:", err.Error())
			rsp.Code = code.LogonFailure
			return nil
		}
		// verification code is ok now
		// get user
		usr, err := user.Get(req.CountryCode, req.PhoneNumber)
		if err != nil {
			log.Error("Login.user.GetUser failure, err:", err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}

		rsp.Code = code.Success
		rsp.Role = usr.Role
		rsp.UserId = usr.Id
		return nil
	} else if len(req.Password) > 0 {
		bytes, err := crypto.RSADecrypt(req.Password)
		if err != nil {
			rsp.Code = code.UnsupportedType
			return nil
		}
		usr, err := user.Get(req.CountryCode, req.PhoneNumber)
		if err != nil {
			log.Error("Login.user.GetUser failure, err:", err.Error())
			rsp.Code = code.EntryNotFound
			return nil
		}

		ok := bcrypt.PasswordVerify(string(bytes), usr.Password)
		if !ok {
			rsp.Code = code.LogonFailure
			return nil
		}
		rsp.Code = 0
		rsp.UserId = cast.ToInt64(usr.Id)
		rsp.Role = usr.Role
		return nil
	} else {
		rsp.Code = code.InvalidData
		return nil
	}
}
