package v1

import (
	"GinProject/middleware"
	"GinProject/model"
	"GinProject/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary      用户登录
// @Description  用户登录，需要账号、密码
// @Tags         用户接口
// @Accept       json
// @Produce      json
// @Param        userinfo  body      model.UserLogin  true  "用户登录"
// @Success      200       {object}  model.Success{data=model.Token}
// @Failure      404       {object}  model.Error
// @Router       /login [post]
func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var token string
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}
