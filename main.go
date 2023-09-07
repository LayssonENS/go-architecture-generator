package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LayssonENS/go-architecture-generator/config"
	_ "github.com/LayssonENS/go-architecture-generator/docs"
	generate "github.com/LayssonENS/go-architecture-generator/generate/delivery/http"
	"github.com/LayssonENS/go-architecture-generator/generate/usecase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const shutdownTimeout = 30 * time.Second

func main() {
	log := logrus.New()

	router := setupRouter()
	httpServer := setupHTTPServer(router)

	go startServer(httpServer, log)

	waitForShutdownSignal(log)

	shutdownServer(httpServer, log)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	}

	router.Static("/", "./static")

	useCaseGenerate := usecase.NewUserUseCase()
	generate.NewGenerateHandler(router, useCaseGenerate)

	return router
}

func setupHTTPServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetEnv().Port),
		Handler: router,
	}
}

func startServer(server *http.Server, log *logrus.Logger) {
	if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("listen: %s\n", err)
	}
}

func waitForShutdownSignal(log *logrus.Logger) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down API...")
}

func shutdownServer(server *http.Server, log *logrus.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("API Server forced to shutdown:", err)
	}

	log.Println("API Server exiting")
}
