package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	Id    int
	Name  string
	Email string
	NoHp  string
}

func connectionDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "mountz:z8579973#&+@tcp(localhost:3306)/kontak")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return db, nil
}

func getData() ([]Contact, error) {
	db, err := connectionDatabase()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	query, err := db.Query("SELECT * FROM friendlist")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer query.Close()

	var contacts []Contact
	for query.Next() {
		var contact = Contact{}
		err := query.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.NoHp)
		if err != nil {
			fmt.Println(err.Error())
		}
		contacts = append(contacts, contact)
	}

	if err := query.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return contacts, nil
}

func insertData(name, email, nohp string) error {
	db, err := connectionDatabase()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO friendlist (name, email, nomor_hp) VALUES (?, ?, ?)", name, email, nohp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	return err
}
func retriveData(id string) (Contact, error) {
	db, err := connectionDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return Contact{}, err
	}
	defer db.Close()

	var contact Contact
	err = db.QueryRow("SELECT * FROM friendlist WHERE id = ?", id).Scan(&contact.Id, &contact.Name, &contact.NoHp, &contact.Email)
	if err != nil {
		fmt.Println(err.Error())
		return Contact{}, err
	}
	defer db.Close()

	return contact, err
}
func updateData(id, name, email, nohp string) error {
	db, err := connectionDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE friendlist SET name = ?, email = ?, nomor_hp = ? WHERE id = ?", name, email, nohp, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func deleteData(id string) error {
	db, err := connectionDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM FRIENDLIST WHERE ID = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
