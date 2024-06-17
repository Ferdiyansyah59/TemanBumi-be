package main

import (
	"sampah/config"
	"sampah/controller"
	"sampah/middleware"
	"sampah/repository"
	"sampah/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	// Repo
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)
	carbonsRepository repository.CarbonsRepository = repository.NewCarbonsRepository(db)

	// Service
	jwtService     service.JWTService     = service.NewJWTService()
	authService    service.AuthService    = service.NewAuthServie(userRepository)
	articleService service.ArticleService = service.NewArticleService(articleRepository)
	carbonsService service.CarbonsService = service.NewCarbonsService(carbonsRepository)

	// Controller
	authController         controller.AuthController         = controller.NewAuthController(authService, jwtService)
	articleController controller.ArticleContorller = controller.NewArticleController(articleService, jwtService)
	carbonsController controller.CarbonsController = controller.NewCarbonsController(carbonsService, jwtService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(CORSMiddleware())

	authRoutes := r.Group("api")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	
	articleRoutes := r.Group("api", middleware.AuthorizeJWT(jwtService))
	{
		// Untuk menampilkan semua data
		articleRoutes.GET("/getAllArticles", articleController.GetAllArticle)
		// Untuk menampilkan artikel terbaru
		articleRoutes.GET("/getLatestArticles", articleController.GetLatestArticles)
		// Untuk output klasifikasi
		articleRoutes.GET("/getArticleWithKey/:title", articleController.GetArticleByKey)
		// Show Article
		articleRoutes.GET("/showArticle/:slug", articleController.ShowArticle)
	}

	carbonsRoutes := r.Group("api", middleware.AuthorizeJWT(jwtService))
	{
		// Untuk menampilkan detail carbon user
		carbonsRoutes.GET("/getDetailCarbons/:user_id", carbonsController.GetDetailCarbons)
		// Untuk menampilkan foot print user
		carbonsRoutes.GET("/getFootPrint/:user_id", carbonsController.GetFootPrint)
		// Untuk insert carbon user
		carbonsRoutes.POST("/insertCarbons", carbonsController.InsertCarbons)
	}

	r.Run(":8081")
}