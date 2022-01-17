package v1

import (
	"GinProject/middleware"
	"GinProject/model"
	"GinProject/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	err := model.Rdb.Del(c, data.Username).Err()
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
