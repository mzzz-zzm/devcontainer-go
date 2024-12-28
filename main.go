package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	// wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutting down server...")

	// create a deadline to wait for the server to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}

func setupRouter() *gin.Engine {
	// simple
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// GET with param
	r.GET("/hi/:name", func(c *gin.Context) {
		n := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + n,
		})
	})
	// GET with query param
	r.GET("/greet", func(c *gin.Context) {
		n := c.DefaultQuery("name", "Guest")
		c.String(http.StatusOK, "Hello %s!", n) // non-json response
	})
	// POST
	r.POST("/user", func(c *gin.Context) {
		msg := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"received": msg,
		})
	})
	// PUT
	r.PUT("/user:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "update with user id " + id,
		})
	})
	// DELETE
	r.DELETE("/user:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "delete with user id " + id,
		})
	})
	return r
}
