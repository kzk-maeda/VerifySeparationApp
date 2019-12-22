package model

import "time"

func (user *User) CreateSession() (session Session, err error) {
	sql, err := readSQLFile("model/sql/insert_session.sql")
	if err != nil {
		return
	}
	satement := sql
	stmt, err := Db.Prepare(satement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

func (user *User) CreateUser() (err error) {
	sql, err := readSQLFile("model/sql/insert_user.sql")
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).
		Scan(&user.ID, &user.UUID, &user.CreatedAt)
	return
}

//
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	sql, err := readSQLFile("model/sql/select_user_by_email.sql")
	if err != nil {
		return
	}
	err = Db.QueryRow(sql, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
