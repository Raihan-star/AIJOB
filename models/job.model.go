package models

import (
	"aijob/db"
	"net/http"
)

type DataJob struct {
	Id         int    `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Namajob    string `param:"namajob" query:"namajob" form:"namajob" json:"namajob" xml:"namajob"`
	Keterangan string `param:"keterangan" query:"keterangan" form:"keterangan" json:"keterangan" xml:"keterangan"`
	Bukti      string `param:"bukti" query:"bukti" form:"bukti" json:"bukti" xml:"bukti"`
	Status     string `param:"status" query:"status" form:"status" json:"status" xml:"status"`
	Tanggal    string `param:"tanggal" query:"tanggal" form:"tanggal" json:"tanggal" xml:"tanggal"`
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

func StoreDataJobMeraihQ(namajob string, keterangan string, bukti string, status string, tanggal string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO data_job_meraih (namajob, keterangan, bukti, status, tanggal) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namajob, keterangan, bukti, status, tanggal)
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
func UpdateDataJobMeraih(id int, namajob string, keterangan string, bukti string, status string, tanggal string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE data_job_meraih SET code_job = ?, namajob = ?, keterangan = ?, bukti = ?, status = ?, tanggal = ? Where id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namajob, keterangan, bukti, status, tanggal, id)
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
