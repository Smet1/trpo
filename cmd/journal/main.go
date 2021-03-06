package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Smet1/trpo/internal/db"
	"github.com/Smet1/trpo/internal/handlers"

	"github.com/Smet1/trpo/internal/config"
	"github.com/Smet1/trpo/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

func main() {
	configPath := flag.String(
		"config",
		"./config.yaml",
		"path to config",
	)
	flag.Parse()

	filenameHook := filename.NewHook()
	filenameHook.Field = "sourcelog"

	log := logrus.New()
	log.AddHook(filenameHook)

	log.Formatter = &logrus.JSONFormatter{}

	cfg := &config.Config{}
	err := config.ReadConfig(*configPath, &cfg)
	if err != nil {
		log.WithError(err).Fatal("can't read config")
	}
	log.WithField("config", cfg).Info("started with config")

	conn, err := db.EnsureDBConn(cfg)
	if err != nil {
		log.WithError(err).Fatal("can't create db connection")
	}

	mux := chi.NewRouter()
	mux.Use(middleware.NoCache)
	mux.Use(logger.GetLoggerMiddleware(log))
	mux.Use(db.GetDbConnMiddleware(conn))
	mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodHead, http.MethodPatch, http.MethodPut},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "metadata"},
		ExposedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "metadata"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	ph := handlers.Posts{Conn: conn}
	uh := handlers.User{Conn: conn}

	mux.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", uh.CreateUser)
			r.Get("/{username}", uh.GetUser)
			r.Post("/login", uh.Auth)
		})

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", ph.CreatePost)
			r.Get("/{post_id}", ph.GetPost)
			r.Get("/", ph.GetPosts)
		})
	})

	server := http.Server{
		Handler: mux,
		Addr:    cfg.ServeAddr,
	}

	go func() {
		log.Infof("ejournal backend service started on port %s", cfg.ServeAddr)
		if err = server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("graceful shutdown")
			} else {
				log.WithError(err).Fatal("sync service")
			}
		}
	}()

	sgnl := make(chan os.Signal, 1)
	signal.Notify(sgnl,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	stop := <-sgnl

	if err = server.Shutdown(context.Background()); err != nil {
		log.WithError(err).Error("error on shutdown")
	}

	log.WithField("signal", stop).Info("stopping")
}
