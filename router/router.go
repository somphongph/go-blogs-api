package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-blogs-api/blog"
)

func NewRouter(e *echo.Echo) {
	// Blog
	b := e.Group("/blogs")
	hBlog := blog.New(blog.NewMongoDBStore())
	{
		b.GET(":id", hBlog.GetById)
		b.POST("", hBlog.CreateHandler)
	}

	// Graceful Shutdown
	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
