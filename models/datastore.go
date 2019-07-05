package models

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

type DataStoreInsert struct {
	Cpf           string
	Private       string
	Incompleto    string
	LastDate      string
	AvgTicket     string
	LastTicket    string
	StoreFrequent string
	StoreLast     string
}

type DataStore struct {
	Cpf           string
	Private       NullInt64   `json:"private"`
	Incompleto    NullInt64   `json:"incompleto"`
	LastDate      NullString  `json:"lastdate"`
	AvgTicket     NullFloat64 `json:"avgticket"`
	LastTicket    NullFloat64 `json:"lastticket"`
	StoreFrequent NullString  `json:"storefrequent"`
	StoreLast     NullString  `json:"storelast"`
}

func AllDataStores(db *sql.DB) ([]*DataStore, error) {
	rows, err := db.Query("SELECT cpf, private, incompleto, lastDate, avgTicket, lastTicket, StoreFrequent, StoreLast FROM datastore")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataStores := make([]*DataStore, 0)
	for rows.Next() {
		ds := new(DataStore)
		err := rows.Scan(&ds.Cpf, &ds.Private, &ds.Incompleto, &ds.LastDate, &ds.AvgTicket, &ds.LastTicket, &ds.StoreFrequent, &ds.StoreLast)
		if err != nil {
			return nil, err
		}
		dataStores = append(dataStores, ds)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return dataStores, nil
}

func CreateDataStore(db *sql.DB, ds DataStoreInsert) error {
	sql := "INSERT INTO datastore (cpf, private, incompleto, lastdate, avgticket, lastTicket,  storefrequent, storelast) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) "

	cpf := NewNullString(ds.Cpf)
	private := NewNullInt(ds.Private)
	incompleto := NewNullInt(ds.Incompleto)
	lastDate := NewNullString(ds.LastDate)
	avgTicket := NewNullFloat(ds.AvgTicket)
	lastTicket := NewNullFloat(ds.LastTicket)
	storeFrequent := NewNullString(ds.StoreFrequent)
	storeLast := NewNullString(ds.StoreLast)

	result, err := db.Exec(sql, cpf, private, incompleto, lastDate, avgTicket, lastTicket, storeFrequent, storeLast)

	if err != nil {
		log.Panic(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Panic(err)
		return err
	}

	strconv.FormatInt(rowsAffected, 10)
	return nil
}

func DeleteAllDataStore(db *sql.DB) error {
	sql := "DELETE FROM datastore"
	result, err := db.Exec(sql)

	if err != nil {
		log.Panic(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Panic(err)
		return err
	}

	strconv.FormatInt(rowsAffected, 10)
	return nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 || s == "NULL" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt(s string) sql.NullInt64 {
	if len(s) == 0 || s == "NULL" {
		return sql.NullInt64{}
	}

	i64, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return sql.NullInt64{
		Int64: i64,
		Valid: true,
	}
}

func NewNullFloat(s string) sql.NullFloat64 {
	if len(s) == 0 || s == "NULL" {
		return sql.NullFloat64{}
	}

	sr := strings.Replace(s, ",", ".", -1)
	f64, err := strconv.ParseFloat(sr, 64)

	if err != nil {
		log.Fatal(err)
	}

	return sql.NullFloat64{
		Float64: f64,
		Valid:   true,
	}
}
