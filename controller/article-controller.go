package controller

import (
	"fmt"
	"net/http"
	"sampah/helper"
	"sampah/service"

	"github.com/gin-gonic/gin"
)

type ArticleContorller interface {
	GetAllArticle(ctx *gin.Context)
	GetLatestArticles(ctx *gin.Context)
	GetArticleByKey(ctx *gin.Context)
	ShowArticle(ctx *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
	jwtService     service.JWTService
}

func NewArticleController(art service.ArticleService , jwt service.JWTService) ArticleContorller {
	return &articleController{
		articleService: art,
		jwtService:     jwt,
	}
}

func (c *articleController) GetAllArticle(ctx *gin.Context) {
	articles := c.articleService.GetAllArticle()
	res := helper.BuildResponse(true, "Success", articles)
	ctx.JSON(http.StatusOK, res)
}
func (c *articleController) GetLatestArticles(ctx *gin.Context) {
	articles := c.articleService.GetLatestArticles()
	res := helper.BuildResponse(true, "Success", articles)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) GetArticleByKey(ctx *gin.Context){
	title := ctx.Param("title")
	articles := c.articleService.GetArticleByKey(title)
	res := helper.BuildResponse(true, "Success", articles)
	fmt.Println("Ini makanan ", title)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) ShowArticle(ctx *gin.Context){
	slug := ctx.Param("slug")
	articles := c.articleService.ShowArticle(slug)
	res := helper.BuildResponse(true, "Success", articles)
	fmt.Println("Ini makanan ", slug)
	ctx.JSON(http.StatusOK, res)
}