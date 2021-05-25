package database_api

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var users []Users

	for rows.Next() {
		var user Users
		rows.Scan(&user.Id, &user.Email, &user.Password, &user.First_Name, &user.Last_Name)
		users = append(users, user)
	}

	usersBytes, _ := json.MarshalIndent(users, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersBytes)

	defer rows.Close()
	defer db.Close()
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, user.Email, user.Password, user.First_Name, user.Last_Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
