package connection

import (
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"loan.com/config"
)

type ReplicationDB struct {
	Primary *sqlx.DB
	Standby *sqlx.DB
}

func NewPostgresSQL(
	psqlConf *config.DBConfig,
	driverName string,
) *sqlx.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		psqlConf.Username,
		url.QueryEscape(psqlConf.Password),
		psqlConf.Host,
		psqlConf.Port,
		psqlConf.Name,
	)

	conn, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("postgresql: cannot open connection: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("initializer: cannot connect to postgresql: %s", err)
	}

	conn.DB.SetMaxOpenConns(int(psqlConf.MaxOpen))
	conn.DB.SetMaxIdleConns(int(psqlConf.MaxIdle))
	conn.DB.SetConnMaxIdleTime(psqlConf.MaxIdleTime)

	return conn
}
