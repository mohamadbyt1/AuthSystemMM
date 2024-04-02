package user

import (
	"database/sql"

)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{db: db}

}

func (d *Repository) CreateUser(user *CreateUserReq) (error){
	_ , err := d.db.Exec("INSERT INTO users (username,password) VALUES ($1,$2)",user.Username,user.Password)
		if err != nil {
			return err
		}
		return nil
	

}
func (d *Repository)UserExists (uname string) (bool, error){
	var c int
	qeruy := "SELECT COUNT(*) FROM users WHERE username = $1"

	if err := d.db.QueryRow(qeruy,uname).Scan(&c); err != nil {
		return false ,err
	}
	return c > 0,nil
}
func (db *Repository) GetUser(uname string) (*User, error) {
	query := "SELECT id, username, password, role FROM users WHERE username = $1"
	var user User
	err := db.db.QueryRow(query, uname).Scan(&user.Id,&user.Username, &user.Password,&user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err	
	}
	return &user, nil
}