package dataupload

import (
	"brankas_test/dependency/sqlservice"
	"log"
)

type dependencyholder struct {
	iUploadSQLservice InterfaceSQLite
}

var Dholder dependencyholder

func Initdependency(requestID int64) error {

	sqlserv := UploadSQLservice{DB: sqlservice.Db}
	Dholder = dependencyholder{
		iUploadSQLservice: &sqlserv,
	}
	log.Println("[", requestID, "] | Initdependency() - Initialised dependency-holder object for the data upload service.")
	return nil
}
