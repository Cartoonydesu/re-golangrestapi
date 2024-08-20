package main

import (
	"cartoonydesu/skill"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Print(os.Getenv("POSTGRES_URI"))
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	connStr := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	r := gin.Default()
	skill.SkillRouter(r, &skill.Handler{Db: db})
	srv := http.Server{
		Addr:        ":" + os.Getenv("PORT"),
		Handler:     r,
		ReadTimeout: 3 * time.Second,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}
	slog.Info("Server Shuting down...")
}

// func SetRouter(r *gin.Engine, sh *skill.Handler) {

// }
