package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.alexedwards.net/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

type ExampleModel struct {
	DB *sql.DB
	InsertStmt *sql.Stmt
}

func NewExampleModel(db *sql.DB) (*ExampleModel, error){
	insertStmt, err := db.Prepare(`INSERT INTO ...`)
	if err != nil {
		return nil , err
	}

	return &ExampleModel{DB: db, InsertStmt: insertStmt}, nil 
}

func (m *ExampleModel) Insert(args...any) error{
	_,err := m.InsertStmt.Exec(args...)
	
	return err 
}

func (m *ExampleModel) ExampleTransaction()error{
	tx, err := m.DB.Begin()
	if err != nil {
		return err 
	}

	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO ...`)
	if err != nil {
		return err
	}
	
	_,err = tx.Exec(`UPDATE ...`)
	if err != nil {
		return err 
	}


	err = tx.Commit()
	return err
}

func main() {
	//command line flag to specifying the port to be used
	addr := flag.String("addr", ":4000", "HTTP network address")
	//data source name MySQL
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
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
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}
	//Initialising our instance of the application struct whic currently
	//stores our structured log format and our models.SnippetMode which contains
	//our database connection pool "db"

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
