package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	//"github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//command line flag to specifying the port to be used
	addr := flag.String("addr", ":4000", "HTTP network address")
	//data source name MySQL
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(loggerHandler)
	//intialising our structured logger

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
	}
	//Initialising our instance of the application struct whic currently
	//stores our structured log format

	//specifying the http method
	logger.Info("Starting Server", slog.Any("addr", *addr))

	//ListenAndServe starts our server and binds it to the address specified
	err = http.ListenAndServe(*addr, app.routes())
	//logs an error if err is non-nil
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
