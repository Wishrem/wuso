package storage

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Wishrem/wuso/client/types"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	_db, err := sql.Open("sqlite3", "./local/local.db")
	if err != nil {
		log.Fatal("Open local storage failed:", err)
	}
	_db.SetMaxOpenConns(5)
	_db.SetMaxIdleConns(3)
	_db.SetConnMaxLifetime(0)
	db = _db

	ping(context.Background())
	inter := make(chan os.Signal, 3)
	signal.Notify(inter, os.Interrupt)

	go func(db *sql.DB) {
		<-inter
		db.Close()
	}(db)

	prepareStmts()
}

func ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("unable to connect to database: ", err)
	}
}

func SaveMsg(msg *types.Message) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Begin tx failed:", err)
	}

	defer recoverTx(tx, err)

	_, err = CreateMsg(getTableName(msg.From, msg.IsGroup), msg.Id, msg.Content, msg.Index, msg.EOF, time.UnixMicro(msg.CreatedAt))
	return err
}

func getTableName(from int64, isGroup bool) string {
	f := strconv.FormatInt(from, 10)
	if isGroup {
		return "s" + f
	} else {
		return "g" + f
	}
}

func recoverTx(tx *sql.Tx, err error) {
	if r := recover(); r != nil {
		log.Panicln("rolling back tx:", r)
		err := tx.Rollback()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Println("rolling back tx:", err)
		err := tx.Rollback()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := tx.Commit()
		if err != nil {
			log.Fatal("committing tx:", err)
		}
	}
}
