package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建路由
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()
	// 服务端给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			// maxAge int, 单位：秒
			// path, cookie所在目录
			// domain string, 域名
			// secure 是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie","value_cookie",60,"/",
				"localhost",false,true)
		}
		fmt.Printf("cookie的值是：%s\n",cookie)

	})
	r.Run(":8000")
}