package router

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/go-post-api/post"
)

func NewRouter(e *echo.Echo) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	// Expense
	expenseHandler := post.NewHandler(db)
	expenseHandler.InitDB()
	expense := e.Group("/expenses")
	{
		expense.POST("", expenseHandler.CreateExpenseHandler)
	}

	// Graceful Shutdown
	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
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
