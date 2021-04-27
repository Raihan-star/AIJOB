package models

import (
	"aijob/db"
	"net/http"
)

type DataJob struct {
	Id         int    `json:"id"`
	Code_Job   string `json:"code_Job"`
	Nama_Job   string `json:"nama_job"`
	Keterangan string `json:"keterangan"`
	Bukti      string `json:"bukti"`
	Status     string `json:"status"`
	Tanggal    string `json:"tanggal"`
}

// GET DATA JOB All
func FetchAllDataJobAll() (Response, error) {
	var obj DataJob
	var arrobj []DataJob
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM data_job"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Code_Job, &obj.Nama_Job, &obj.Keterangan, &obj.Bukti, &obj.Status, &obj.Tanggal)
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

// GET DATA JOB
func FetchAllDataJobMeraih() (Response, error) {

	var obj DataJob
	var arrobj []DataJob
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM data_job_meraih ORDER BY id_meraih DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Code_Job, &obj.Nama_Job, &obj.Keterangan, &obj.Bukti, &obj.Status, &obj.Tanggal)
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

// PUT DATA JOB
func StoreDataJobMeraih(code_Job string, nama_job string, keterangan string, bukti string, status string, tanggal string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT data_job_meraih (code_job, nama_job, keterangan, bukti, status, tanggal) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(code_Job, nama_job, keterangan, bukti, status, tanggal)
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

// UPDATE DATA JOB
func UpdateDataJobMeraih(id int, code_job string, nama_job string, keterangan string, bukti string, status string, tanggal string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE data_job_meraih SET code_job = ?, nama_job = ?, keterangan = ?, bukti = ?, status = ?, tanggal = ? Where id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(code_job, nama_job, keterangan, bukti, status, tanggal, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

// DELETE DATA JOB
func DeleteDataJobMeraih(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM data_job_meraih WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil
}
