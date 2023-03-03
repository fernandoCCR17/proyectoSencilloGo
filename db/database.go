package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Aqui colocar tu usuario y contraseña
const url = "<Usuario>:<Contraseña>@tcp(localhost:3306)/goweb_db" //goweb_db es el nombre de la bd

var db *sql.DB

func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = conection
}

// Cierra la conexion
func Close() {
	db.Close()
}

// Verifica la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s", tableName)
	rowes, err := Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rowes.Next()
}

// Crea una tabla
func CreateTable(schema, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("Truncate %s", tableName)
	Exec(sql)
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
