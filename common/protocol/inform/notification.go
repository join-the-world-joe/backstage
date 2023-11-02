package inform

type Notification struct {
	Event   string `json:"event"`
	Message string `json:"message"`
}

const (
	EventForceOffline   = "Force Offline"
	MessageForceOffline = "Your account has been logon somewhere! Your have to logout immediately."
)
