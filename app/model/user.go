package model

import (
	"log"
	"sort"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsersFromDB() ([]User, error) {
	cmd := `SELECT * FROM users`
	rows, err := Db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	// IDで昇順にソートする
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	return users, nil
}

func GetUserByIDFromDB(id int) (User, error) {
	cmd := `SELECT * FROM users WHERE id = $1`

	var user User
	err := Db.QueryRow(cmd, id).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func InsertUserToDB(user User) error {
	cmd := `INSERT INTO users (email, password) VALUES ($1, $2)`

	_, err := Db.Exec(cmd, user.Email, user.Password)
	return err
}

func ValidateUserExists(id int) error {
	cmd := `SELECT id FROM users WHERE id = $1`
	var foundID int
	err := Db.QueryRow(cmd, id).Scan(&foundID)
	return err
}

func UpdateUserInDB(user User, id int) error {
	cmd := `UPDATE users SET email = $1, password = $2 WHERE id = $3`

	_, err := Db.Exec(cmd, user.Email, user.Password, id)
	return err
}

func DeleteUserFromDB(id int) error {
	cmd := `DELETE FROM users WHERE id = $1`

	_, err := Db.Exec(cmd, id)
	return err
}
