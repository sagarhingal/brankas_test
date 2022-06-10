package dataupload

type Auth struct {
	Token string
}

type Metadata struct {
	Originalname string `json:"originalname"`
	Newname      string `json:"newname"`
	Filesize     int64  `json:"filesize"`
	Contenttype  string `json:"contenttype"`
	Agent        string `json:"agent"`
	ClientIP     string `json:"clientip"`
}
