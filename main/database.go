package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"reflect"
	"regexp"
)

func dbConncetion() (*sql.DB, error) {
	var db *sql.DB
	config := mysql.Config{
		User:                     "",
		Passwd:                   "",
		Net:                      "tcp",
		Addr:                     "domain",
		DBName:                   "chema",
		Params:                   nil,
		Collation:                "",
		Loc:                      nil,
		MaxAllowedPacket:         0,
		ServerPubKey:             "",
		TLSConfig:                "",
		TLS:                      nil,
		Timeout:                  0,
		ReadTimeout:              0,
		WriteTimeout:             0,
		AllowAllFiles:            false,
		AllowCleartextPasswords:  false,
		AllowFallbackToPlaintext: false,
		AllowNativePasswords:     true,
		AllowOldPasswords:        false,
		CheckConnLiveness:        false,
		ClientFoundRows:          false,
		ColumnsWithAlias:         false,
		InterpolateParams:        false,
		MultiStatements:          false,
		ParseTime:                true,
		RejectReadOnly:           false,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	return db, nil
}

func editDatabaseFromRequest(request string, args ...any) (int, error) {
	db, err := dbConncetion()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	for i, arg := range args {
		if reflect.TypeOf(arg).String() == "string" && len(arg.(string)) < 1 {
			args[i] = nil
		}
	}
	response, err := db.Exec(request, args...)
	if err != nil {
		return 0, err
	}
	effectedRows, err := response.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(effectedRows), nil
}

func getInfoFromRequest(request string, args ...any) (*sql.Rows, error) {
	db, err := dbConncetion()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(request, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func getInfoWithNullParams(request string, args ...any) (*sql.Rows, error) {
	db, err := dbConncetion()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	for i, arg := range args {
		if reflect.TypeOf(arg).String() == "string" && len(arg.(string)) < 1 {
			args[i] = nil
		}
	}
	rows, err := db.Query(request, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func isValidUUID(input string) bool {
	uuidPattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	match, _ := regexp.MatchString(uuidPattern, input)
	return match
}
