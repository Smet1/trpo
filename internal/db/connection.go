package db

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Smet1/trpo/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ctxdb struct{}

func WithLogger(ctx context.Context, conn *sqlx.DB) context.Context {
	return context.WithValue(ctx, ctxdb{}, conn)
}

func GetConnection(ctx context.Context) *sqlx.DB {
	l, ok := ctx.Value(ctxdb{}).(sqlx.DB)
	if !ok {
		logrus.Fatal("can't get db connection")
	}
	return &l
}

func GetDbConnMiddleware(conn *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := WithLogger(req.Context(), conn)

			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}

func EnsureDBConn(cfg *config.Config) (*sqlx.DB, error) {
	v := url.Values{}
	v.Add("sslmode", cfg.DB.SSLMode)

	p := url.URL{
		Scheme:     cfg.DB.Database,
		Opaque:     "",
		User:       url.UserPassword(cfg.DB.Username, cfg.DB.Password),
		Host:       cfg.DB.Host,
		Path:       cfg.DB.Name,
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   v.Encode(),
		Fragment:   "",
	}

	connectURL, err := pq.ParseURL(p.String())
	if err != nil {
		return nil, errors.Wrap(err, "can't create url for db connection")
	}

	instance, err := sqlx.Connect(cfg.DB.Database, connectURL)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect db")
	}

	return instance, nil
}
