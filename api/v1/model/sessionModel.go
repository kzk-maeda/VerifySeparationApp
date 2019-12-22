package model

func GetSession(userID int) (session Session, err error) {
	sql, err := readSQLFile("model/sql/select_session_by_userid.sql")
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = Db.QueryRow(sql, userID).Scan(&session.ID, &session.UUID, &session.UserID)
	return
}
