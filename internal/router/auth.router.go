package router

import (
	"github.com/alfannurfaiz9/shortlink_backend/internal/controller"
	"github.com/alfannurfaiz9/shortlink_backend/internal/repo"
	"github.com/alfannurfaiz9/shortlink_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitAuthRouter(router *gin.Engine, db *pgxpool.Pool) {
	r := router.Group("/auth")

	ar := repo.NewAuthRepo(db)
	as := service.NewAuthService(ar)
	ac := controller.NewController(as)

	r.POST("register", ac.Register)
}
