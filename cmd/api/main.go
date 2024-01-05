package main

import (
	"os"
	"spun/internal/handler"
	"spun/internal/repository/mysql"
	"spun/internal/service"
	"spun/pkg/libdb"
	"spun/pkg/libecho"
	"spun/pkg/libecho/libmiddleware"

	"github.com/helloferdie/golib/v2/liblogger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := libdb.Open("")
	if err != nil {
		liblogger.Errorf("Error connect database %v", err)
		os.Exit(1)
	}
	defer db.Close()

	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(libmiddleware.Logger())
	e.Use(libmiddleware.Session)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	jwtByte := []byte(os.Getenv("jwt_secret"))
	jwtMiddleware := libmiddleware.JWT(jwtByte)

	categoryRepo := mysql.NewMySQLCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoute := e.Group("/v1/category")
	categoryRoute.POST("/create", categoryHandler.CreateCategory)
	categoryRoute.POST("/view", categoryHandler.ViewCategory)
	categoryRoute.POST("/update", categoryHandler.UpdateCategory)
	categoryRoute.POST("/delete", categoryHandler.DeleteCategory)
	categoryRoute.POST("/list", categoryHandler.ListCategory)
	categoryRoute.POST("/restricted", categoryHandler.ViewCategory, jwtMiddleware)
	categoryRoute.POST("/restricted2", categoryHandler.ViewCategory, jwtMiddleware)

	libecho.StartHttp(e)
}
