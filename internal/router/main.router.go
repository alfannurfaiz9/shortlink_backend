package router

import (
	_ "github.com/alfannurfaiz9/shortlink_backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool) {
	router.GET("documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitAuthRouter(router, db)
}
