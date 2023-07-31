package main

//
//import (
//	"database/sql"
//	_ "embed"
//	"encoding/json"
//	"fmt"
//	"github.com/jackskj/carta"
//	"log"
//	"net/http"
//	"strconv"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type Person struct {
//	ID        int    `json:"id" db:"id"`
//	FirstName string `json:"first_name" db:"first_name" `
//	LastName  string `json:"last_name" db:"last_name"`
//}
//
////go:embed query_test.sql
//var queryEmbed string
//
//func main() {
//	db, err := sql.Open("mysql", "root:db_mysql@tcp(localhost:3319)/db_mysql")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	http.HandleFunc("/api/people", func(w http.ResponseWriter, r *http.Request) {
//		//pageSize := 10
//
//		// Obtener el cursor del parámetro de la URL
//		cursorStr := r.URL.Query().Get("cursor")
//		pageSize := r.URL.Query().Get("size")
//		var size int
//		size = 10
//		if pageSize != "" {
//			size, err = strconv.Atoi(pageSize)
//			if err != nil {
//				http.Error(w, "Invalid cursor", http.StatusBadRequest)
//				return
//			}
//		}
//		var cursor int
//		if cursorStr != "" {
//			cursor, err = strconv.Atoi(cursorStr)
//			if err != nil {
//				http.Error(w, "Invalid cursor", http.StatusBadRequest)
//				return
//			}
//		}
//		var people []Person
//
//		// Consultar los datos de la página actual utilizando el cursor
//
//		query := queryEmbed
//		rows, err := db.Query(query, cursor, size)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer rows.Close()
//
//		// Construir el cursor para la próxima página (el último ID de la página actual)
//		var nextCursor int
//		//for rows.Next() {
//		//	var person Person
//		//	err := rows.Scan(&person.ID, &person.FirstName, &person.LastName)
//		//	if err != nil {
//		//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		//		return
//		//	}
//		//	people = append(people, person)
//		//	nextCursor = person.ID
//		//}
//		carta.Map(rows, &people)
//
//		// Construir el cursor para la página anterior (el primer ID de la página actual)
//		var prevCursor int
//		if cursor > 0 {
//			prevCursor = cursor - size
//			fmt.Println(cursor)
//		}
//
//		// Declarar un slice para almacenar los datos de la página actual
//
//		// Obtener los datos de la página actual utilizando el cursor
//		//rows, err = db.Query(query)
//		//if err != nil {
//		//	http.Error(w, err.Error(), http.StatusInternalServerError)
//		//	return
//		//}
//		//defer rows.Close()
//
//		//for rows.Next() {
//		//	var person Person
//		//	err := rows.Scan(&person.ID, &person.FirstName, &person.LastName)
//		//	if err != nil {
//		//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		//		return
//		//	}
//		//	people = append(people, person)
//		//}
//
//		// Crear un mapa para enviar los resultados al frontend
//		response := map[string]interface{}{
//			"next_cursor": nextCursor,
//			"prev_cursor": prevCursor,
//			"people":      people,
//		}
//
//		// Convertir el mapa a JSON y enviarlo al frontend
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(response)
//	})
//
//	log.Println("Servidor iniciado en http://localhost:8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
