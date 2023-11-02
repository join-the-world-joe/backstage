package business

import (
	"backstage/common/cache/string/verification_code"
	"backstage/common/code"
	"backstage/common/protocol/sms"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/sms/runtime"
	"backstage/utils/random_number"
	"backstage/validator"
	"context"
	"strings"
)

func SendVerificationCode(ctx context.Context, req *sms.SendVerificationCodeReq, rsp *sms.SendVerificationCodeRsp) error {
	log.DebugF("SendVerificationCode Behavior[%v], Code:[%v], Number:[%v] ", req.Behavior, req.CountryCode, req.PhoneNumber)
	// check if mobile valid
	ok := validator.IsMobileValid(req.CountryCode, req.PhoneNumber)
	if !ok {
		rsp.Code = code.InvalidDataType
		return nil
	}

	// generate verification code
	c, err := random_number.Generate(
		runtime.OTPBeginOfRegister(),
		runtime.OTPEndOfRegister(),
		runtime.OTPLenOfRegister(),
	)
	if err != nil {
		log.Error("SendVerificationCode.random_number.Generate failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	c = runtime.OTP(c) // for debug purposes

	// create short message
	if len(req.Behavior) <= 0 {
		log.Error("SendVerificationCode failure,  len(req.Behavior) <= 0 ")
		rsp.Code = code.InvalidData
		return nil
	}
	if _, exist := config.SMSConf().SMS.Behavior[req.Behavior]; !exist {
		log.Error("SendVerificationCode failure, behavior doesn't exist")
		rsp.Code = code.UnsupportedType
		return nil
	}

	// replace code
	message := strings.Replace(config.SMSConf().SMS.Behavior[req.Behavior].Template, "${code}", c, 1)
	log.Debug("message: ", message)

	// request to send verification code
	err = sendSMS(req.CountryCode, req.PhoneNumber, message)
	if err != nil {
		log.Error("SendVerificationCode.sendSMS failure, err: ", err.Error())
		rsp.Code = code.ServiceError
		return nil
	}

	// cache verification code
	err = verification_code.Create(req.Behavior, req.CountryCode, req.PhoneNumber, c)
	if err != nil {
		log.Error("SendVerificationCode.verification_code.Create failure, err: ", err.Error())
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rsp.Code = code.Success
	return nil
}
