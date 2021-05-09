package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/spf13/viper"
)

func NewMSSQLConn() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;connection timeout=%d;",
		viper.GetString("MSSQL.HOST"),
		viper.GetString("MSSQL.USERNAME"),
		viper.GetString("MSSQL.PASSWORD"),
		viper.GetString("MSSQL.DATABASE"),
		viper.GetInt("MSSQL.TIMEOUT"),
	)
	conn, err := sql.Open(viper.GetString("MSSQL.TYPE"), connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
