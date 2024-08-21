package main

import (
	"cartoonydesu/database"
	"cartoonydesu/skill"
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db := database.NewPostgres()
	defer db.Close()
	if err := db.Ping(); err != nil {
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

func init() {
	// database.ResetDB()
}

// func SetRouter(r *gin.Engine, sh *skill.Handler) {

// }
