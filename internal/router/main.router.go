package router

import (
	_ "github.com/alfannurfaiz9/shortlink_backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	router.GET("documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitAuthRouter(router, db, rdb)
}
