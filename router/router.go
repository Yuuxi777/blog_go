package router

import (
	"blog_go/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/getBlogById", controller.GetBlogById)

	r.GET("/getAllBlogs", controller.GetAllBlogs)

	r.GET("/test", sayHello)

	return r
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello!")
}
