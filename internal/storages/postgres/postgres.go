package storages

import (
	"auth/internal/storages"
	"auth/pkg/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	DBCnf := config.GetConfig().DBConfig
	fmt.Println(DBCnf)

	var err error
	DB, err = newInstance(DBCnf)

	if err != nil {
		panic("Failed postgres connect: " + err.Error())
	}
}

func newInstance(cfg storages.PostgresConfig) (db *sqlx.DB, err error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool params
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
