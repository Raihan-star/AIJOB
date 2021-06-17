package models

import (
	"aijob/db"
	"net/http"
)

type DataJob struct {
	Id         int    `json:"id"`
	Namajob    string `json:"namajob"`
	Keterangan string `json:"keterangan"`
	Bukti      string `json:"bukti"`
	Status     string `json:"status"`
	Tanggal    string `json:"tanggal"`
}

// GET DATA JOB
func FetchAllDataJobMeraih() (Response, error) {

	var obj DataJob
	var arrobj []DataJob
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM data_job_meraih ORDER BY id DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Namajob, &obj.Keterangan, &obj.Bukti, &obj.Status, &obj.Tanggal)
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
func StoreDataJobMeraih(job DataJob) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO data_job_meraih (namajob, keterangan, bukti, status, tanggal) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(job.Namajob, job.Keterangan, job.Bukti, job.Status, job.Tanggal)
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
func UpdateDataJobMeraih(job DataJob) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE data_job_meraih SET namajob = ?, keterangan = ?, bukti = ?, status = ?, tanggal = ? Where id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(job.Namajob, job.Keterangan, job.Bukti, job.Status, job.Tanggal, job.Id)
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
func DeleteDataJobMeraih(job DataJob) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM data_job_meraih WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(job.Id)
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
