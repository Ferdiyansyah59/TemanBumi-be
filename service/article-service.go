package service

import (
	"sampah/entity"
	"sampah/repository"
)

type ArticleService interface {
	GetAllArticle() []entity.Article
	GetLatestArticles() []entity.Article
	GetArticleByKey(title string) []entity.Article
	ShowArticle(slug string) []entity.Article
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(artRepo repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository: artRepo,
	}
}

func (serv *articleService) GetAllArticle() []entity.Article {
	return serv.articleRepository.GetAllArticle()
}
func (serv *articleService) GetLatestArticles() []entity.Article {
	return serv.articleRepository.GetLatestArticles()
}

func (serv *articleService) GetArticleByKey(title string) []entity.Article{
	return serv.articleRepository.GetArticleByKey(title)
}
func (serv *articleService) ShowArticle(slug string) []entity.Article{
	return serv.articleRepository.ShowArticle(slug)
}

