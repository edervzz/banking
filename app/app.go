package app

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Run() {
	ctx := createContext()
	env := getEnvVars()
	mh := NewMigrationHandler(service.NewDefaultMigrationService(domain.NewMigrationRepositoryDB(*ctx)))
	ch := NewCustomerHandler(service.NewCustomerServiceInterface(domain.NewCustomerRepositoryDB(*&ctx)))

	router := mux.NewRouter()
	router.HandleFunc("/migrations", mh.Migrations).Methods(http.MethodPost)
	router.HandleFunc("/customer", ch.Create).Methods(http.MethodPost)

	logger.Info("listening on " + env.server + ":" + env.port)
	if err := http.ListenAndServe(env.server+":"+env.port, router); err != nil {
		logger.Fatal(err.Error())
	}

}

// -------------------------------------
type envvars struct {
	server string
	port   string
	dbName string
	dbUser string
	dbPass string
}

func getEnvVars() envvars {
	if os.Getenv("SERVER") == "" ||
		os.Getenv("PORT") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASS") == "" {
		return envvars{
			server: "localhost",
			port:   "8000",
			dbName: "banking",
			dbUser: "root",
			dbPass: "eder",
		}
		//utils.Fatal("Please check environment variables")
	}
	return envvars{
		server: os.Getenv("SERVER"),
		port:   os.Getenv("PORT"),
		dbName: os.Getenv("DB_NAME"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASS"),
	}
}

// -------------------------------------
func newClientDB(user string, pass string, name string) *sqlx.DB {
	client, err := sqlx.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = client.Ping()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	return client
}

// -------------------------------------
func createContext() *context.Context {
	env := getEnvVars()
	clientdb := newClientDB(env.dbUser, env.dbPass, env.dbName)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "clientdb", clientdb)
	return &ctx
}
