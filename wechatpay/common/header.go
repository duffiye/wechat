package common

type Header struct {
	Timestamp string `json:"-"`
	NonceStr  string `json:"-"`
}
