package types

// UserSignData 用户签到数据
type UserSignData struct {
	RawJson string
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}
