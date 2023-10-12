package data

import "time"

type User struct {
	ID uint64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	EmailVerified bool `json:"verified"`
	Role string `json:"role"`
}

func Create(user *User) error {
	statement := "insert into users(email, password, email_verified, access, created, updated) values($1, $2, $3, $4, $5, $6)"
	timestamp := time.Unix(time.Now().Unix(), 0)
	_, err := DB.Exec(statement, user.Email, user.Password, user.EmailVerified, user.Role, timestamp, timestamp)
	return err;
}

func Get(id string) (User, error){
	var user User
	statement := "SELECT * FROM users WHERE id=$1"
	row := DB.QueryRow(statement, id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.EmailVerified, &user.Role)
		if err != nil {
			return User{}, err
		}
	return user, err
}

func CheckEmail(email string, user *User) bool {

	statement := "SELECT * FROM users WHERE email=$1"
	rows, err := DB.Query(statement, email)
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