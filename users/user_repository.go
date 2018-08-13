package users

import (
	"encoding/json"
	"time"

	"github.com/go-squads/reuni-server/helper"
)

type (
	verifiedUser struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		IAT      int64  `json:"iat"`
	}
)

const (
	createUserQuery  = "INSERT INTO users (name, username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	verifyLoginQuery = "SELECT id, name, username, email FROM users WHERE username=$1 AND password=$2"
)

type userRepositoryInterface interface {
	createUser(userstore user) error
	loginUser(loginData userv) ([]byte, error)
}

type userRepository struct {
	execer helper.QueryExecuter
}

func initRepository(execer helper.QueryExecuter) *userRepository {
	return &userRepository{
		execer: execer,
	}
}

func (u *userRepository) createUser(userstore user) error {
	userstore.CreatedAt = time.Now()
	userstore.UpdatedAt = userstore.CreatedAt
	_, err := u.execer.DoQueryRow(createUserQuery, userstore.Name, userstore.Username, userstore.Password, userstore.Email, userstore.CreatedAt, userstore.UpdatedAt)
	// db := context.GetDB()
	// _, err := db.Exec(createUserQuery, userstore.Name, userstore.Username, userstore.Password, userstore.Email, userstore.CreatedAt, userstore.UpdatedAt)
	return err
}

func (u *userRepository) loginUser(loginData userv) ([]byte, error) {
	v := verifiedUser{}
	data, err := u.execer.DoQueryRow(verifyLoginQuery, loginData.Username, loginData.Password)
	// data.Scan(&v.ID, &v.Name, &v.Username, &v.Email)
	if err != nil {
		return nil, err
	}
	v.ID = int(data["id"].(int64))
	v.Name = data["name"].(string)
	v.Username = data["username"].(string)
	v.Email = data["email"].(string)
	v.IAT = makeTimestamp()

	return json.Marshal(v)
}
