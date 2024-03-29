package main

import (
	"goproj/internal/handler"
	"goproj/internal/repository/mysql"
	"goproj/internal/service"
	"goproj/pkg/libdb"
	"goproj/pkg/libecho"
	"goproj/pkg/libecho/libmiddleware"
	"os"

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

	// Audit Trail

	auditTrailRepo := mysql.NewMySQLAuditTrailRepository(db)
	auditTrailService := service.NewAuditTrailService(auditTrailRepo)

	// Category

	categoryRepo := mysql.NewMySQLCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, auditTrailService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoute := e.Group("/v1/category")
	categoryRoute.POST("/create", categoryHandler.CreateCategory)
	categoryRoute.POST("/view", categoryHandler.ViewCategory)
	categoryRoute.POST("/update", categoryHandler.UpdateCategory)
	categoryRoute.POST("/delete", categoryHandler.DeleteCategory)
	categoryRoute.POST("/list", categoryHandler.ListCategory)
	categoryRoute.POST("/restricted", categoryHandler.ViewCategory, jwtMiddleware)
	categoryRoute.POST("/restricted2", categoryHandler.ViewCategory, jwtMiddleware)

	// Country

	countryRepo := mysql.NewMySQLCountryRepository(db)
	countryService := service.NewCountryService(countryRepo)
	countryHandler := handler.NewCountryHandler(countryService)

	countryRoute := e.Group("/v1/country")
	countryRoute.POST("/list", countryHandler.ListCountry)
	countryRoute.POST("/view", countryHandler.ViewCountry)

	// User

	userRepo := mysql.NewMySQLUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoute := e.Group("/v1/user")
	userRoute.POST("/list", userHandler.ListUser)
	userRoute.POST("/view", userHandler.ViewUser)

	libecho.StartHttp(e)
}
