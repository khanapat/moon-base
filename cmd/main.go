package main

import (
	"context"
	"fmt"
	"log"
	"moon-base/coin"
	"moon-base/docs"
	"moon-base/internal/database"
	"moon-base/logz"
	"moon-base/middleware"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	_ "moon-base/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	runtime.GOMAXPROCS(1)
	initTimezone()
	initViper()
}

// @title Moon Coin
// @version 1.0
// @description Moon Coin for exchange.
// @termsOfService http://swagger.io/terms/
// @contact.name Khanapat.A
// @contact.url http://www.swagger.io/support
// @contact.email k.apiwattanawong@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /moon
// @schemes http https
func main() {
	route := mux.NewRouter()

	cfgCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                    // All origins
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
		AllowedHeaders:   []string{"Content-Type", "Origin", "Authorization", "Accept"},
		AllowCredentials: true,
	})

	logger := logz.NewLogConfig()

	middle := middleware.NewMiddleware(logger)

	route.Use(middle.ContextLocaleMiddleware)

	swag := route.NewRoute().Subrouter()

	swag.Use(middle.BasicAuthenication)

	moonApi := route.PathPrefix(viper.GetString("app.context")).Subrouter()

	moonApi.Use(middle.JSONMiddleware)
	moonApi.Use(middle.ContextLogAndLoggingMiddleware)

	mssqlDB, err := database.NewMSSQLConn()
	if err != nil {
		logger.Error(err.Error())
	}
	defer mssqlDB.Close()

	moonApi.Handle("/buy", coin.NewBuyCoin(
		coin.NewGetSupplyByIDFn(mssqlDB),
		coin.NewUpdateSupplyByIDFn(mssqlDB),
		coin.NewCreateHistoryLogsFn(mssqlDB),
	)).Methods(http.MethodPost)

	moonApi.Handle("/history", coin.NewGetHistory(
		coin.NewGetHistoryLogsFn(mssqlDB),
	)).Methods(http.MethodGet)

	moonApi.Handle("/supply", coin.NewGetSupplyCoin(
		coin.NewGetSupplyFn(mssqlDB),
	)).Methods(http.MethodGet)

	moonApi.Handle("/reset", coin.NewResetMoonCoin(
		coin.NewResetSupplyFn(mssqlDB),
	)).Methods(http.MethodGet)

	registerSwaggerRoute(swag)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", viper.GetString("app.port")),
		Handler:      cfgCors.Handler(route),
		ReadTimeout:  viper.GetDuration("app.timeout"),
		WriteTimeout: viper.GetDuration("app.timeout"),
		IdleTimeout:  viper.GetDuration("app.timeout"),
	}

	logger.Info(fmt.Sprintf("⇨ http server started on [::]:%s", viper.GetString("APP.PORT")))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Info(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("app.timeout"))
	defer cancel()

	srv.Shutdown(ctx)

	logger.Info("shutting down")
	os.Exit(0)
}

func initViper() {
	viper.SetDefault("app.name", "moon-base")
	viper.SetDefault("app.port", "9090")
	viper.SetDefault("app.timeout", "60s")
	viper.SetDefault("app.context", "/moon")

	viper.SetDefault("log.env", "dev")
	viper.SetDefault("log.level", "debug")

	viper.SetDefault("swagger.host", "localhost:9090")
	viper.SetDefault("swagger.user", "admin")
	viper.SetDefault("swagger.password", "password")

	viper.SetDefault("mssql.type", "sqlserver")
	viper.SetDefault("mssql.host", "localhost")
	viper.SetDefault("mssql.username", "sa")
	viper.SetDefault("mssql.password", "Password1234")
	viper.SetDefault("mssql.database", "master")
	viper.SetDefault("mssql.timeout", "60")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("error loading location 'Asia/Bangkok': %v\n", err)
	}
	time.Local = ict
}

func registerSwaggerRoute(route *mux.Router) {
	route.PathPrefix("/moon-coin/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/moon-coin/swagger/doc.json", viper.GetString("swagger.host"))),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = viper.GetString("app.context")
}
