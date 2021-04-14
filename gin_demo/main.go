package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		cookie, err := c.Cookie("abc")
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"StatusUnauthorized"})
			// 若验证不通过，不再调用后续的函数处理
			c.Abort()
		}
		if cookie=="123"{
			c.Next()
			return
		}
	}
}

func main() {
	// 1. 创建路由
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()
    r.GET("/login", func(c *gin.Context) {
		// 设置cookie
    	c.SetCookie("abc","123",60,"/",
    		"localhost",false,true)

    	// 返回信息
    	c.String(200,"Login success!")
	})
    r.GET("/home",AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200,gin.H{"data":"home"})
	})
	r.Run(":8000")
}
