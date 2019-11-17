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
	"go.opencensus.io/plugin/ochttp"
)

type ctxdb struct{}

func GetDbConnMiddleware(conn *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return ochttp.Handler{
			Handler: http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				ctx := context.WithValue(req.Context(), ctxdb{}, conn)

				next.ServeHTTP(res, req.WithContext(ctx))
			}),
		}.Handler
	}
}

func GetConnection(ctx context.Context) *sqlx.DB {
	l, ok := ctx.Value(ctxdb{}).(sqlx.DB)
	if !ok {
		logrus.Fatal("can't get db connection")
	}
	return &l
}

func EnsureDBConn(config *config.Config) (*sqlx.DB, error) {
	v := url.Values{}
	v.Add("ssl-mode", config.DB.SSLMode)

	p := url.URL{
		Scheme:     config.DB.Database,
		Opaque:     "",
		User:       url.UserPassword(config.DB.Username, config.DB.Password),
		Host:       config.DB.Host,
		Path:       config.DB.Name,
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   v.Encode(),
		Fragment:   "",
	}

	connectURL, err := pq.ParseURL(p.String())
	if err != nil {
		return nil, errors.Wrap(err, "can't create url for db connection")
	}

	instance, err := sqlx.Connect(config.DB.Database, connectURL)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect db")
	}

	return instance, nil
}
