package model

import (
	"database/sql"
	"fmt"
	logger "read-csv/services"

	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	id     int
	name   string
	sex    string
	age    int
	height int
	weight int
}

func Create(db *sql.DB, newPerson person) {
	stmt, _ := db.Prepare("INSERT INTO people (id, name, sex, age, height, weight) VALUES (?, ?, ?, ?, ?, ?)")
	stmt.Exec(nil, newPerson.name, newPerson.sex, newPerson.age, newPerson.height, newPerson.weight)
	defer stmt.Close()

	fmt.Printf("Added %v \n", newPerson.name)
}

func FindPerson(db *sql.DB, queryString string) []person {

	rows, err := db.Query("SELECT * FROM people WHERE name like '%" + queryString + "%'")
	logger.CheckErr(err)
	defer rows.Close()

	err = rows.Err()
	logger.CheckErr(err)

	people := make([]person, 0)

	for rows.Next() {
		ourPerson := person{}
		err = rows.Scan(&ourPerson.id, &ourPerson.name, &ourPerson.sex, &ourPerson.age, &ourPerson.height, &ourPerson.weight)
		logger.CheckErr(err)

		people = append(people, ourPerson)
	}

	err = rows.Err()
	logger.CheckErr(err)

	return people
}

func FindPersonById(db *sql.DB, personID string) person {

	rows, _ := db.Query("SELECT * FROM people WHERE id = '" + personID + "'")
	defer rows.Close()

	ourPerson := person{}

	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.name, &ourPerson.sex, &ourPerson.age, &ourPerson.height, &ourPerson.weight)
	}

	return ourPerson
}

func Update(db *sql.DB, ourPerson person) int64 {

	stmt, err := db.Prepare("UPDATE people set name = ?, sex = ?, age = ?, height = ? weight = ?, where id = ?")
	logger.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(ourPerson.name, ourPerson.sex, ourPerson.age, ourPerson.height, ourPerson.weight, ourPerson.id)
	logger.CheckErr(err)

	affected, err := res.RowsAffected()
	logger.CheckErr(err)

	return affected
}

func Delete(db *sql.DB, idToDelete string) int64 {

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	logger.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	logger.CheckErr(err)

	affected, err := res.RowsAffected()
	logger.CheckErr(err)

	return affected
}
