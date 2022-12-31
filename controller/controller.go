package controller

import (
	"blog_go/result"
	"blog_go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBlogById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	blog, err := service.GetBlogById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail(-1, "查询失败，请检查id"))
	} else {
		c.JSON(http.StatusOK, result.Success("查询成功", blog))
	}

}

func GetAllBlogs(c *gin.Context) {
	blogs, err := service.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail(-1, "查询失败"))
	} else {
		c.JSON(http.StatusOK, result.Success("查询成功", blogs))
	}

}
