package dataupload

import (
	"database/sql"
	"log"
)

type UploadSQLservice struct {
	DB *sql.DB
}

type InterfaceSQLite interface {
	Savemetadata(requestID int64, data Metadata) error
	Getdata(requestID int64) ([]Metadata, error)
}

func (us *UploadSQLservice) Savemetadata(requestID int64, data Metadata) error {

	res, err := us.DB.Exec("INSERT INTO imgmeta(originalname, newname, filesize, contenttype, agent, clientIP) values(?,?,?,?,?,?)",
		data.Originalname, data.Newname, data.Filesize, data.Contenttype, data.Agent, data.ClientIP)
	if err != nil {
		log.Println("[", requestID, "] | Savemetadata() - Unable to execute the query. | ", err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("[", requestID, "] | Savemetadata() - Unable to retrieve the last inserted ID. | ", err)
		return err
	}

	log.Println("[", requestID, "] | Savemetadata() - Data saved successfully with ID[", id, "]")
	// Return the last inserted ID
	return nil

}

// Getdata : This function fetches all the entries from the images table
func (us *UploadSQLservice) Getdata(requestID int64) (dataarray []Metadata, err error) {

	rows, err := us.DB.Query("SELECT originalname, newname, filesize, contenttype, agent, clientip FROM imgmeta")
	if err != nil {
		log.Println("[", requestID, "] | Getdata() - Unable to execute query. | ", err)
		return nil, err
	}
	defer rows.Close()

	var all []Metadata
	for rows.Next() {
		var mdata Metadata
		if err := rows.Scan(&mdata.Originalname, &mdata.Newname, &mdata.Filesize, &mdata.Contenttype, &mdata.Agent, &mdata.ClientIP); err != nil {
			return nil, err
		}
		all = append(all, mdata)
	}
	log.Println("[", requestID, "] | Getdata() - Successfully retrieved all the data. ")
	return all, nil

}
