package track

// output
type Model struct {
	Operator   string `json:"operator"`
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	Permission string `json:"permission"`
	Request    string `json:"request"`
	Response   string `json:"response"`
	Timestamp  int64  `json:"timestamp"`
}
