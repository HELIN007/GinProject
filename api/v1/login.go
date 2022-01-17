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
		// 将token写入redis，时间是ns（1ms=10e6ns）
		err := model.Rdb.Set(c, data.Username, token, 0).Err()
		if err != nil {
			middleware.Infof("将用户%v的token存入redis失败，原因是：", err)
		}
		err = model.Rdb.Set(c, token, data.Username, 0).Err()
		if err != nil {
			middleware.Infof("将用户%v的token存入redis失败，原因是：", err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}
