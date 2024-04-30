package database

import (
	"database/sql"
	// "fmt"
	// "time"

    _ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(db_name, host, user, password string) (*Postgres, error) {
    db, err := sql.Open("postgres", "postgresql://" + user + ":" + password + "@" + host + "/" + db_name + "?sslmode=disable")
	if err != nil {
		return nil, err
	}

    return &Postgres{DB: db}, nil

	// Try connecting to the database
	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()
	//
	// timeoutExceeded := time.After(60 * time.Second)
	// for {
	// 	select {
	// 	case <-timeoutExceeded:
	// 		return nil, fmt.Errorf("connection timeout")
	//
	// 	case <-ticker.C:
	// 		if err := db.Ping(); err == nil {
	// 			return &Postgres{DB: db}, nil
	// 		}
	// 	}
	// }
}

