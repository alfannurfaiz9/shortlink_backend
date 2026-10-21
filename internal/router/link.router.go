package router

import (
	"github.com/alfannurfaiz9/shortlink_backend/internal/controller"
	"github.com/alfannurfaiz9/shortlink_backend/internal/middleware"
	"github.com/alfannurfaiz9/shortlink_backend/internal/repo"
	"github.com/alfannurfaiz9/shortlink_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitLinkRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	r := router.Group("/links")

	r.Use(middleware.AuthMiddleware(rdb))

	lr := repo.NewLinkRepo(db)
	ls := service.NewLinkService(lr)
	lc := controller.NewLinkController(ls)

	r.POST("/", lc.CreateLink)
	r.GET("/", lc.GetUserLink)
	r.DELETE("/:id", lc.DeleteUserLink)
}
