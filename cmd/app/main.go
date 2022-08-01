package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godruoyi/go-snowflake"
	_cakeHttpDelivery "github.com/nmfzone/privy-cake-store/cake/delivery/http"
	_cakeRepo "github.com/nmfzone/privy-cake-store/cake/repository/mysql"
	_cakeUsecase "github.com/nmfzone/privy-cake-store/cake/usecase"
	"github.com/nmfzone/privy-cake-store/config"
	_ "github.com/nmfzone/privy-cake-store/internal"
	"github.com/nmfzone/privy-cake-store/internal/logger"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	snowflake.SetMachineID(1)
	snowflake.SetStartTime(
		time.Date(
			2022,
			1,
			1,
			0,
			0,
			0,
			0,
			time.FixedZone("Asia/Jakarta", 7),
		),
	)

	config.InitConfig()
	logger.InitLogger()

	cfg := config.Get()

	dbConn, err := sql.Open(`mysql`, config.Dbdsn())

	if err != nil {
		log.Fatal().Err(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal().Err(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal().Err(err)
		}
	}()

	engine := gin.Default()
	engine.Use(cors.Default())

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Cake Store",
		})
	})

	cakeRepo := _cakeRepo.NewMysqlCakeRepository(dbConn)
	cakeUsecase := _cakeUsecase.NewCakeUsecase(cakeRepo, time.Duration(2)*time.Second)
	_cakeHttpDelivery.NewCakeHandler(engine, cakeUsecase)

	log.Fatal().Err(engine.Run(cfg.App.Host + ":" + cfg.App.Port))
}
