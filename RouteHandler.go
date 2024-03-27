package main

import (
	"encoding/json"
	"net/http"
	"path"
	"text/template"
)

type ContactUser struct {
	Name  string `json:"name"`
	Nohp  string `json:"nohp"`
	Email string `json:"email"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	temp, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts, err := getData()
	data := struct {
		Contacts []Contact
	}{
		Contacts: contacts,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func insert(w http.ResponseWriter, r *http.Request) {
	var contact ContactUser
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = insertData(contact.Name, contact.Email, contact.Nohp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]
	contact, err := retriveData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// data := struct {
	// 	Contact Contact
	// }{
	// 	Contact: contact,
	// }

	// tmpl, err := template.New("edit").ParseFiles("edit.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func update(w http.ResponseWriter, r *http.Request) {
	var contact ContactUser
	id := r.URL.Path[len("/update/"):]
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := contact.Name
	email := contact.Email
	nohp := contact.Nohp
	err = updateData(id, name, email, nohp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/delete/"):]
	err := deleteData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
