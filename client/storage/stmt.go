package storage

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

var (
	// Single Chat
	createSingleChatTable *sql.Stmt

	// Common
	createMsg *sql.Stmt
)

func prepareStmts() {
	if db == nil {
		log.Fatal("unknown error when preparing stmts")
	}
	var err error
	createSingleChatTable, err = db.Prepare(`CREATE TABLE IF NOT EXISTS ? (
		id VARCHAR(255),
		content VARCHAR(255) NOT NULL,
		index VARCHAR(255) NOT NULL,
		eof CHAR(1) NOT NULL CHECK (eof = '1' OR eof = '0'),
		created_at TEXT NOT NULL,
		updated_at TEXT,
		del_flag CHAR(1) CHECK (del_flag = '1' OR del_flag = '0'),
		PRIMARY KEY (id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	if err != nil {
		log.Fatal("Prepare createSingleChatTable stmt error:", err)
	}

	createMsg, err = db.Prepare(`INSERT INTO ?(id, content, index, eof, created_at) values(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal("Prepare createMsg stmt error:", err)
	}
}

func CreateSingleChatTable(tableName string) error {
	_, err := createSingleChatTable.Exec(tableName)
	return err
}

func CreateMsg(tableName string, id int64, content string, index int64, eof bool, createAt time.Time) (sql.Result, error) {
	idStr := strconv.FormatInt(id, 10)
	idxStr := strconv.FormatInt(index, 10)
	var eofStr string
	if eof {
		eofStr = "1"
	} else {
		eofStr = "0"
	}
	createAtStr := createAt.Format("2023-05-01 00:00:00")
	return createMsg.Exec(tableName, idStr, content, idxStr, eofStr, createAtStr)
}
