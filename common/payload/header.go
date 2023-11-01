package payload

type Header struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

func (p *Header) GetMajor() string {
	return p.Major
}

func (p *Header) GetMinor() string {
	return p.Minor
}

