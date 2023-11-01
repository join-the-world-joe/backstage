package runtime

import "backstage/service/sms/conf"

func OTPLenOfRegister() int {
	if serviceConf.Servant.OTPLenOfRegister == 0 {
		return conf.DefaultOTPLenOfRegister
	}
	return serviceConf.Servant.OTPLenOfRegister
}

func OTPBeginOfRegister() int {
	if serviceConf.Servant.OTPBeginOfRegister == 0 {
		return conf.DefaultOTPBeginOfRegister
	}
	return serviceConf.Servant.OTPBeginOfRegister
}

func OTPEndOfRegister() int {
	if serviceConf.Servant.OTPEndOfRegister == 0 {
		return conf.DefaultOTPEndOfRegister
	}
	return serviceConf.Servant.OTPEndOfRegister
}
