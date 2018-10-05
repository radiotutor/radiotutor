package dbutils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/square/squalor"
	"os"
)

var (
	serverURL   = os.Getenv("SQL_URL")
	sqlPassword = os.Getenv("SQL_PASSWORD")
	sqlUsername = os.Getenv("SQL_USERNAME")
)

func DB(ex func(*squalor.DB) interface{}) interface{} {
	_db, err := sql.Open("mysql", sqlUsername+":"+sqlPassword+"@tcp("+serverURL+")/rt")
	if err != nil {
		panic(err)
	}
	defer _db.Close()

	db, _ := squalor.NewDB(_db)
	if err != nil {
		panic(err)
	}

	rtn := ex(db)
	return rtn
}
