package service

import (
	"blog_go/models"
)

func GetBlogById(id int) (blog *models.Blog, err error) {
	blog, err = models.GetBlogById(id)
	if err != nil {
		return nil, err
	}
	//fmt.Println(blog)
	return
}

func GetAllBlogs() (blogs []*models.Blog, err error) {
	blogs, err = models.GetAllBlogs()
	if err != nil {
		return nil, err
	}
	return
}
