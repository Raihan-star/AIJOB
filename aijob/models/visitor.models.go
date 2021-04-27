package models

import (
	"aijob/db"
	"net/http"
)

type Visitor struct {
	Id_visitor       int    `json:"id_visitor"`
	Username_visitor string `json:"username_visitor"`
	Institut_visitor string `json:"institut_visitor"`
	Email_visitor    string `json:"email_visitor"`
}

func FetchAllDataVisitor() (Response, error) {

	var obj Visitor
	var arrobj []Visitor
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM visitor"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id_visitor, &obj.Username_visitor, &obj.Institut_visitor, &obj.Email_visitor)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
func LoginDataVisitor(username_visitor string, institut_visitor string, email_visitor string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT visitor (username_visitor, institut_visitor, email_visitor) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username_visitor, institut_visitor, email_visitor)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertId,
	}

	return res, nil
}
