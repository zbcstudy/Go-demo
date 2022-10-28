package common

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

// 定义消息
type Message struct {
	Type string `json:"type"`
	Data string `json:"date"`
}

type LoginMes struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	Username string `json:"username"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码
	Error string `json:"error"` // 返回错误信息
}
