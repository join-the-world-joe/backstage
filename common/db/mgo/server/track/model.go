package track

// output
type Model struct {
	UserId    string `json:"user_id"`
	Major     string `json:"major"`
	Minor     string `json:"minor"`
	Request   string `json:"request"`
	Response  string `json:"response"`
	Timestamp int64  `json:"timestamp"`
}
