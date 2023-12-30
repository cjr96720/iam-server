package main

import (
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"iam-service/api/router"
	_ "iam-service/docs"
	"iam-service/internal/database"
)

// @title			IAM Service
// @version		1.0
// @description	Provide API endpoints for IAM service.
// @host			localhost:8080
// @BasePath		/
func main() {
	// migrate database
	db := database.GetDatabase()
	database.AutoMigrate(db)

	// initialize an adapter
	// an adapter is a component that provides a bridge between Casbin's policy storage mechanism
	// and various data sources or databases
	casbinDBAdapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal(err)
	}

	// initialize an enforcer
	// an enforcer is an essential component responsible for enforcing access control policies
	enforcer, err := casbin.NewEnforcer(os.Getenv("CASBIN_MODEL"), casbinDBAdapter)
	if err != nil {
		log.Fatal(err)
	}
	// reload the policy from dratabase
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatal(err)
	}

	// create server
	gin := gin.Default()

	// add CORS
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))

	// add swagger
	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// setup router
	router.Setup(db, enforcer, gin)

	// run server
	gin.Run("0.0.0.0:8080")
}
