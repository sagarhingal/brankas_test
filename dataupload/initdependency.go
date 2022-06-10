package dataupload

import (
	"brankas_test/dependency/sqlservice"
)

type dependencyholder struct {
	iUploadSQLservice InterfaceSQLite
}

var Dholder dependencyholder

func Initdependency() error {

	sqlserv := UploadSQLservice{DB: sqlservice.Db}
	Dholder = dependencyholder{
		iUploadSQLservice: &sqlserv,
	}
	return nil
}
