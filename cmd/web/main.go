package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(loggerHandler)
	//intialising our structured logger

	app := &application{
		logger: logger,
	}
	//Initialising our instance of the application struct whic currently
	//stores our structured log format

	//specifying the http method
	logger.Info("Starting Server", slog.Any("addr", *addr))

	//ListenAndServe starts our server and binds it to the address specified
	err := http.ListenAndServe(*addr, app.routes())
	//logs an error if err is non-nil
	logger.Error(err.Error())
	os.Exit(1)
}
