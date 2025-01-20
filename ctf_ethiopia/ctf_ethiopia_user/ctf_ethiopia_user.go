package ctf_ethiopia_user

import "github.com/jmoiron/sqlx"

type Ctf_Ethiopia_User struct {
	Userid   int
	Username string
	Email    string
	Password string
}

func InsertUser(username string, email string, password string) error {
	db, err := sqlx.Connect("sqlite3", "ctf_ethiopia")
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("Insert into Ctf_Ethiopia_User(username,email,password) values($1,$2,$3)", username, email, password)
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
