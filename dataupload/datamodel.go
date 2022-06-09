package dataupload

type Auth struct {
	Token string
}

type Metadata struct {
	Filename    string `json:"filename"`
	Size        int64  `json:"size"`
	Contenttype string `json:"contenttype"`
	Agent       string `json:"agent"`
	ClientIP    string `json:"clientip"`
}
