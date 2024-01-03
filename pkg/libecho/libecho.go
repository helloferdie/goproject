package libecho

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"spun/pkg/liblogger"
	"time"

	"github.com/labstack/echo/v4"
)

func StartHttp(e *echo.Echo) {
	// Start server
	go func() {
		liblogger.Infow("Start HTTP server")
		if err := e.Start(":" + os.Getenv("app_port")); err != nil && err != http.ErrServerClosed {
			liblogger.Errorf("Fail to start http server %v", err)
		}
		liblogger.Infow("Shutdown HTTP server")
		liblogger.Sync()
		os.Exit(0)
	}()

	// Grace full shut down when received interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		liblogger.Errorf("Fail to shutting down server %v", err)
	}
}
