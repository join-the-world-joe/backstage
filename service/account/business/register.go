package business

import (
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/db/mysql/server/user"
	"backstage/common/protocol/account"
	"backstage/global/log"
	"context"
	"github.com/google/uuid"
)

func Register(ctx context.Context, req *account.RegisterReq, rsp *account.RegisterRsp) error {
	// check if sms verification valid
	err := verification_code.Check(verification_code.Register, req.CountryCode, req.PhoneNumber, req.VerificationCode)
	if err != nil {
		log.Error("Register.verification_code.Check failure, err: ", err.Error())
		rsp.Code = code.AppDataExpired
		return nil
	}

	// check if mobile already exists
	usr, err := user.GetModelByMobile(req.CountryCode, req.PhoneNumber)
	if usr != nil {
		log.ErrorF("Register.user.Get failure, CountryCode[%v] PhoneNumber[%v] exists", req.CountryCode, req.PhoneNumber)
		rsp.Code = code.EntryAlreadyExists
		return nil
	}

	// create new user
	_, err = user.Insert(&user.Model{CountryCode: req.CountryCode, PhoneNumber: req.PhoneNumber, MemberId: uuid.New().String()})
	if err != nil {
		log.Error("Register.user.Insert failure, err: ", err.Error())
		rsp.Code = code.DatabaseFailure
		return err
	}

	rsp.Code = code.Success
	return nil
}
