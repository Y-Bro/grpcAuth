package users

import (
	"database/sql"
	"fmt"
	"log"

	database "github.com/Y-bro/go-grpcAuth/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

func (u *User) Create() (string, error) {
	stmt, err := database.Db.Prepare("INSERT INTO Users(username,password) VALUES(?,?)")

	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := HashPassword(u.Password)

	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(u.Username, hashedPassword)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("User inserted", fmt.Sprint(id))

	return fmt.Sprint(id), nil

}

func (u *User) ValidateUser() bool {
	stmt, err := database.Db.Prepare("SELECT Password from Users where Username = ?")

	if err != nil {
		log.Fatal(err)
	}

	var hashPass string

	err = stmt.QueryRow(u.Username).Scan(&hashPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}

	return CompareHash(u.Password, hashPass)
}

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	hash := string(byte)

	return hash, nil
}

func CompareHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
