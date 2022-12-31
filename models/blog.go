package models

import (
	"blog_go/dao"
	_ "blog_go/dao"
)

type Blog struct {
	Id         int    `json:"id,omitempty"`
	Category   string `json:"category,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Title      string `json:"title,omitempty"`
}

func GetBlogById(id int) (blog *Blog, err error) {
	err = dao.DB.Raw("SELECT * FROM blog WHERE 1=1 AND id = ?", id).First(&blog).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetAllBlogs() (blogs []*Blog, err error) {
	err = dao.DB.Raw("SELECT id,title,category,content,create_time FROM blog Where 1=1").Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return
}
