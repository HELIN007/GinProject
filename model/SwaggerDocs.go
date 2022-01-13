package model

type Success struct {
	Status int         `json:"status" example:"200"` //返回成功码
	Msg    string      `json:"msg"`                  //返回信息
	Data   interface{} `json:"data"`                 //返回数据
}

type UserInfo struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Role     int    `json:"role"`     // 权限码1：管理员；2：普通用户
}

type UserLogin struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type Token struct {
	Token string `json:"token"` // token
}

type Error struct {
	Status int    `json:"status"` //返回错误码
	Msg    string `json:"msg"`    //返回信息
}
