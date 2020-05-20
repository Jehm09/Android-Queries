package model

type Server struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Contry   string `json:"contry"`
	Owner    string `json:"owner"`
}
