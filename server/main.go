package main

import (
	"blog_server/routes"
	"blog_server/utils"
	"blog_server/utils/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Start()
	defer utils.Close()
	gin.SetMode(config.GetString("server.mode"))
	router := routes.InitRouter()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	go func() {
		server.ListenAndServe()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("shutdown error -> %v", err.Error())
		os.Exit(0)
	}
	fmt.Printf("shutdown success")
}
