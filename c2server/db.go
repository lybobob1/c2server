package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	CreateImplant(implant *Implant) error
	GetImplants() ([]*Implant, error)
	doesImplantExist(identifier string) (bool, error)
}

type dbStore struct {
	db *sql.DB
}

var store Store

func createDB(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS c2server")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE c2server")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS implants (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
					 identifier CHAR(20),
					 ipaddress VARCHAR(11),
					 lastcheckin TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					)`)

	if err != nil {
		panic(err)
	}

}

func (store *dbStore) CreateImplant(implant *Implant) error {
	_, err := store.db.Query("INSERT INTO implants(identifier, ipaddress) VALUES (?,?)", implant.Identifier, implant.Ipaddress)
	return err
}

func (store *dbStore) GetImplants() ([]*Implant, error) {
	rows, err := store.db.Query("SELECT * FROM implants")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	implants := []*Implant{}
	for rows.Next() {
		implant := &Implant{}
		if err := rows.Scan(&implant.Dbid, &implant.Identifier, &implant.Ipaddress, &implant.Lastcheckin); err != nil {
			return nil, err
		}

		implants = append(implants, implant)
	}

	return implants, nil
}

func (store *dbStore) doesImplantExist(identifier string) (bool, error) {
	log.Println("Identifier: %s", identifier)
	rows, err := store.db.Query("SELECT * FROM implants WHERE identifier=?", identifier)

	if err != nil {
		return false, err
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		return true, err
	} else {
		return false, err
	}

}

func InitStore(s Store) {
	store = s
}
