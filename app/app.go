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
	ctx, _ := createContext()
	mh := NewMigrationHandler(service.NewDefaultMigrationService(domain.NewMigrationRepositoryDB(*ctx)))
	ch := NewCustomerHandler(service.NewCustomerServiceInterface(domain.NewCustomerRepositoryDB(ctx)))
	ah := NewAccountHandler(service.NewAccountServiceInterface(domain.NewAccountRepositoryDB(ctx)))
	ph := NewPaymitemHandler(service.NewPaymitemServiceInterface(domain.NewPaymItemRepositoryDB(ctx)))
	uh := NewUserHandler(service.NewAuthServiceInterface(domain.NewUserRepositoryDB(ctx)))
	auth := NewAuthHandler(service.NewAuthServiceInterface(domain.NewUserRepositoryDB(ctx)))

	router := mux.NewRouter()
	router.HandleFunc("/migrations", mh.Migrations).
		Methods(http.MethodPost).Name("")

	router.HandleFunc("/register", uh.Create).Methods(http.MethodPost)
	router.HandleFunc("/login", uh.Login).Methods(http.MethodPost)

	router.HandleFunc("/customer", ch.Create).Methods(http.MethodPost).Name("CreateCustomer")

	router.HandleFunc("/customer/{id}", ch.Get).Methods(http.MethodGet).Name("GetCustomer")

	router.HandleFunc("/account", ah.Create).Methods(http.MethodPost).Name("CreateAccount")
	router.HandleFunc("/account/{id}", ah.GetBalance).Methods(http.MethodGet).Name("GetAccount")
	router.HandleFunc("/account/lock/{id}", ah.Lock).Methods(http.MethodPost).Name("LockAccount")
	router.HandleFunc("/account/unlock/{id}", ah.Unlock).Methods(http.MethodPost).Name("UnlockAccount")

	router.HandleFunc("/paymitem", ph.Create).Methods(http.MethodPost).Name("CreatePaymitem")

	router.Use(auth.Authorization())

	logger.Info("listening on : 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		logger.Fatal(err.Error())
	}

}

// -------------------------------------
func createContext() (*context.Context, *envvars) {
	env := getEnvVars()
	clientdb := newClientDB(env.server, env.port, env.dbUser, env.dbPass, env.dbName)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "clientdb", clientdb)
	return &ctx, &env
}

// -------------------------------------
func newClientDB(server string, port string, user string, pass string, dbName string) *sqlx.DB {
	conn := user + ":" + pass + "@tcp(" + server + ":" + port + ")/"
	client, err := sqlx.Open("mysql", conn)
	err = client.Ping()

	if err != nil {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("try", i+1, ":", "mysql:", conn)
			client, err = sqlx.Open("mysql", conn)
			err = client.Ping()
			if err == nil {
				break
			}
		}
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
			port:   "3306",
			dbName: "banking",
			dbUser: "root",
			dbPass: "eder",
		}
	}
	return envvars{
		server: os.Getenv("SERVER"),
		port:   os.Getenv("PORT"),
		dbName: os.Getenv("DB_NAME"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASS"),
	}
}
