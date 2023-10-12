package data

import (
	"fmt"
	"time"
)

type User struct {
	ID uint64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	EmailVerified bool `json:"verified"`
	Role string `json:"role"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func Create(user *User) error {
	statement := "insert into users(email, password, email_verified, access, created, updated) values($1, $2, $3, $4, $5, $6)"
	timestamp := time.Unix(time.Now().Unix(), 0)
	fmt.Println("Timestamp: ", timestamp)
	_, err := db.Exec(statement, user.Email, user.Password, user.EmailVerified, user.Role, timestamp, timestamp)
	fmt.Println("Result: ", err)
	return err;
}

func Get(id string) (User, error){
	var user User
	statement := "SELECT * FROM users WHERE id=$1"
	row := db.QueryRow(statement, id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.EmailVerified, &user.Role)
		if err != nil {
			return User{}, err
		}
	return user, err
}

func CheckEmail(email string, user *User) bool {

	statement := "SELECT * FROM users WHERE email=$1"
	rows, err := db.Query(statement, email)
	if err != nil {
		return false
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.EmailVerified, &user.Role)
		if err != nil {
			return false
		}
	}
	return true
}