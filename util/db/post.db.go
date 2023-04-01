package db

import (
	"gapi/models"
	nanoid "github.com/matoous/go-nanoid/v2"
	"time"
)

func CreatePost(title string, content string, authorID string) (bool, models.Post) {
	driver := Database()
	id, _ := nanoid.Generate("ABCDEFGHJKLIMNOPRSTUYVZ0123456789", 11)
	post := models.Post{
		CID:       id,
		Author:    authorID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	err := driver.Insert(post)
	if err != nil {
		return false, models.Post{}
	}
	AddLog("Post Created. - Title: " + post.Title + ", Content: " + post.Content + ", Author: " + post.Author + ", ID:" + post.CID)

	return true, post
}

func UpdatePost(title string, content string, post *models.Post) (bool, *models.Post) {
	driver := Database()

	post.Title = title
	post.Content = content

	err := driver.Update(post)

	if err != nil {
		return false, nil
	}
	AddLog("Post Updated. - Title: " + post.Title + ", Content: " + post.Content + ", Author: " + post.Author + ", ID:" + post.CID)

	return true, post
}

func DeletePost(post *models.Post) (bool, *models.Post) {
	driver := Database()

	err := driver.Delete(post)
	if err != nil {
		return false, nil
	}
	AddLog("Post Deleted. - Title: " + post.Title + ", Content: " + post.Content + ", Author: " + post.Author + ", ID:" + post.CID)

	return true, post
}

func FindPostByTitle(title string) (*models.Post, bool) {
	driver := Database()
	var post *models.Post
	err := driver.Open(models.Post{}).Where("title", "=", title).First().AsEntity(&post)
	if err != nil {
		return nil, false
	}
	return post, true
}

func FindPostByID(ID string) (*models.Post, bool) {
	driver := Database()
	var post *models.Post
	err := driver.Open(models.Post{}).Where("id", "=", ID).First().AsEntity(&post)
	if err != nil {
		return nil, false
	}
	return post, true
}

func GetAllPosts() ([]models.Post, bool) {
	driver := Database()
	var post []models.Post
	err := driver.Open(models.Post{}).AsEntity(&post)
	if err != nil {
		return nil, false
	}
	return post, true
}
