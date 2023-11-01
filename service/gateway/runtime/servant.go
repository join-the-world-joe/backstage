package runtime

func Endpoint() string {
	return serviceConf.Servant.Endpoint
}

func IPLimit() bool {
	return serviceConf.Servant.IPLimit
}

func EncryptionEnable() bool {
	return serviceConf.Encryption.Enable
}

func QPS() int {
	return serviceConf.Servant.QPS
}

func FeedbackEnable() bool {
	return serviceConf.Feedback.Enable
}

func FeedbackMajor() string {
	return serviceConf.Feedback.Major
}

func FeedbackMinor() string {
	return serviceConf.Feedback.Minor
}

func FeedbackMessage() string {
	return serviceConf.Feedback.Message
}

func FeedbackWaitForClose() int {
	return serviceConf.Feedback.WaitForCloseInterval
}
