package user

/*
   SQL code
    create table users (
        id int primary key auto_increment not null,
        username varchar(64) not null unique,
        password varchar(128) not null,
        email varchar(128) not null unique
    );
*/

import (
	"errors"
	"github.com/pe5er/radiotutor/dbutils"
	"github.com/square/squalor"
)

type User struct {
	UID      int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func getUser(username string) (User, error) {

	type funcRtn struct {
		user User
		err  error
	}

	userData := dbutils.DB(func(db *squalor.DB) interface{} {

		users, err := db.BindModel("users", User{})
		if err != nil {
			panic(err)
		}
		q1 := users.Select("*").Where(users.C("username").Eq(username))
		q2 := users.Select("*").Where(users.C("email").Eq(username))

		var u []User
		err = db.Select(&u, q1)
		if err != nil {
			return funcRtn{User{}, err}
		}
		if len(u) == 1 {
			return funcRtn{u[0], nil}
		}

		err = db.Select(&u, q2)
		if err != nil {
			return funcRtn{User{}, err}
		}
		if len(u) == 1 {
			return funcRtn{u[0], nil}
		}
		return funcRtn{User{}, errors.New("User not found")}
	})
	return userData.(funcRtn).user, userData.(funcRtn).err
}

// Ensure you hash passwords befor passing them to this function
// Also, only pass valid users
func insertUser(u User) error {

	err := dbutils.DB(func(db *squalor.DB) interface{} {

		_, err := db.BindModel("users", User{})
		if err != nil {
			panic(err)
		}

		err = db.Insert(&u)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err.(error)
	}
	return nil
}

func deleteUser(u User) error {

	err := dbutils.DB(func(db *squalor.DB) interface{} {

		_, err := db.BindModel("users", User{})
		if err != nil {
			panic(err)
		}

		_, err = db.Delete(&u)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err.(error)
	}
	return nil
}
