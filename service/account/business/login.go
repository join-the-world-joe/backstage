package business

import (
	"backstage/common/cache/hash/token"
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/db/mysql/server/user"
	"backstage/common/service/account"
	"backstage/global/log"
	"context"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func Login(ctx context.Context, req *account.LoginReq, rsp *account.LoginRsp) error {
	if len(req.VerificationCode) > 0 {
		// check if verification code valid
		err := verification_code.Check(verification_code.Login, req.CountryCode, req.PhoneNumber, req.VerificationCode)
		if err != nil {
			log.Error("Login.verification_code.Check failure, err:", err.Error())
			rsp.Code = code.MissingData
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

		// create token
		unique := uuid.New().String()
		err = token.Create(req.CountryCode, req.PhoneNumber, cast.ToString(usr.Id), unique)
		if err != nil {
			log.Error("Login.token.Create failure, err:", err.Error())
			rsp.Code = code.DatabaseFailure
			return nil
		}

		rsp.Code = 0
		rsp.Token = unique
		rsp.UserId = usr.Id
		return nil
	} else if len(req.Token) > 0 {
		// check if token valid
		temp, err := token.Get(req.CountryCode, req.PhoneNumber, req.Token)
		if err != nil {
			log.Error("Login.token.Get failure, err:", err.Error())
			rsp.Code = code.MissingData
			return nil
		}

		rsp.Code = 0
		rsp.Token = req.Token
		rsp.UserId = cast.ToInt64(temp.UserId)
		return nil
	} else {
		rsp.Code = code.DataSourceProtocolError
		return nil
	}

	return nil
}
