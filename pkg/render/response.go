// @description 响应封装
// @author zkp15
// @date 2023/8/9 23:15
// @version 1.0

package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v-tiktok/model/user"
)

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Error(c *gin.Context, statusCode int, statusMsg string) {
	c.JSON(http.StatusOK, user.Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	})
	c.Abort()
}
