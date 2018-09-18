package user

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/square/squalor"
	"os"
)

var (
	serverURL   = os.Getenv("SQL_URL")
	sqlPassword = os.Getenv("SQL_PASSWORD")
	sqlUsername = os.Getenv("SQL_USERNAME")
)

type User struct {
	UID      int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func init() {
	db, err := sql.Open("mysql", sqlUsername+":"+sqlPassword+"@tcp("+serverURL+")/rt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`
if not exists (select * from sysobjects where name='users' and xtype='U')
    create table users (
        id int primary key auto increment not null,
        username varchar(64) not null unique,
        password varchar(128) not null,
        email varchar(128) not null
    )
go
`)
}

func getUser(username string) (User, error) {
	_db, err := sql.Open("mysql", sqlUsername+":"+sqlPassword+"@tcp("+serverURL+")/rt")
	if err != nil {
		return User{}, err
	}
	defer _db.Close()

	db, _ := squalor.NewDB(_db)
	users, err := db.BindModel("users", User{})

	q := users.Select("*").Where(users.C("username").Eq(username))

	var u []User
	err = db.Select(&u, q)
	if err != nil {
		return User{}, err
	}
	if len(u) == 1 {
		return u[0], nil
	}
	return User{}, errors.New("User Not In DB")
}

// Ensure you hash passwords befor passing them to this function
// Also, only pass valid users
func insertUser(u User) error {
	_db, err := sql.Open("mysql", sqlUsername+":"+sqlPassword+"@tcp("+serverURL+")/rt")
	if err != nil {
		return err
	}
	defer _db.Close()

	db, _ := squalor.NewDB(_db)
	_, err = db.BindModel("users", User{})
	if err != nil {
		panic(err)
	}

	err = db.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}
