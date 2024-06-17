package repository

import (
	"sampah/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetAllArticle() []entity.Article
	GetLatestArticles() []entity.Article
	GetArticleByKey(title string) []entity.Article
	ShowArticle(slug string) []entity.Article
}

type articleConnection struct {
	connection *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleConnection{
		connection: db,
	}
}

func (db *articleConnection) GetAllArticle() []entity.Article {
	var articles []entity.Article
	db.connection.Order("id desc").Find(&articles)
	return articles
}

func (db *articleConnection) GetLatestArticles() []entity.Article {
	var articles []entity.Article
	db.connection.Limit(5).Order("created_At desc").Find(&articles)
	return articles
}


func (db *articleConnection) GetArticleByKey(title string) []entity.Article {
	var articles []entity.Article
	db.connection.Where("title LIKE ?", "%"+title+"%").Find(&articles)
	return articles
}

func (db *articleConnection) ShowArticle(slug string) []entity.Article {
	var articles []entity.Article
	db.connection.Where("slug = ?", slug).Find(&articles)
	return articles
}