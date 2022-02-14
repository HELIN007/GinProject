package v1

import (
	"GinProject/middleware"
	"GinProject/model"
	"GinProject/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Username struct {
	Username string `json:"username"`
}

func Logout(c *gin.Context) {
	var data Username
	_ = c.ShouldBindJSON(&data)
	fmt.Println("======", data)
	token, _ := model.Rdb.Get(c, data.Username).Result()
	fmt.Println(data.Username, "的token是", token)
	err := model.Rdb.Del(c, data.Username).Err()
	_ = model.Rdb.Del(c, token).Err()
	if err != nil {
		code = errmsg.ErrorTokenDel
		middleware.Infof("用户%v正在退出登录，清除token失败，原因是：%v", data.Username, err)
	} else {
		code = errmsg.SUCCESS
		middleware.Infof("用户%v退出成功，清除redis成功！")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
