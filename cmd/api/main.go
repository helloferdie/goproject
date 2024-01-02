package main

import (
	"os"
	"spun/internal/handler"
	"spun/internal/repository/mysql"
	"spun/internal/service"
	"spun/pkg/libdb"
	"spun/pkg/libecho/libmiddleware"
	"spun/pkg/liblogger"

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

	categoryRepo := mysql.NewMySQLCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoute := e.Group("/v1/category")
	categoryRoute.POST("/view", categoryHandler.ViewCategory)
	categoryRoute.POST("/restricted", categoryHandler.ViewCategory, libmiddleware.JWT)

	e.Logger.Fatal(e.Start(":1200"))
}

// func errorHandler(err error, c echo.Context) {
// 	fmt.Println(err)
// 	fmt.Println(c)

// 	return FormatResponse(c, resp)

// 	// code := http.StatusInternalServerError
// 	// if he, ok := err.(*echo.HTTPError); ok {
// 	// 	code = he.Code
// 	// }
// 	// c.Logger().Error(err)
// 	// errorPage := fmt.Sprintf("%d.html", code)
// 	// if err := c.File(errorPage); err != nil {
// 	// 	c.Logger().Error(err)
// 	// }
// }
