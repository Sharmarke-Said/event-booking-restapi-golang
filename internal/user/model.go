package user

import (
	"errors"

	"example.com/event-booking-restapi/internal/auth"
	"example.com/event-booking-restapi/internal/database"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
	Role			string 
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password, role) VALUES (?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.Role)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password, role FROM users WHERE email = ?`
	row := database.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword, &u.Role)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := auth.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
    query := "SELECT id, email, role FROM users"
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Email, &user.Role); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// GetUserByID fetches a single user by their ID.
func GetUserByID(id int64) (*User, error) {
    query := "SELECT id, email, role FROM users WHERE id = ?"
    row := database.DB.QueryRow(query, id)
    var user User
    if err := row.Scan(&user.ID, &user.Email, &user.Role); err != nil {
        return nil, err
    }
    return &user, nil
}

// Update updates a user's email and role.
func (u *User) Update() error {
    query := "UPDATE users SET email = ?, role = ? WHERE id = ?"
    stmt, err := database.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(u.Email, u.Role, u.ID)
    return err
}

// DeleteUserByID deletes a user from the database.
func DeleteUserByID(id int64) error {
    query := "DELETE FROM users WHERE id = ?"
    stmt, err := database.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(id)
    return err
}