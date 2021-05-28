package main

import (
	"fmt"
	"log"
	"os"
	"errors"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")
	/*var city City
    if err := db.Get(&city, "SELECT * FROM city WHERE Name='Tokyo'"); errors.Is(err, sql.ErrNoRows) {
        log.Println("no such city Name=%s", "Tokyo")
    } else if err != nil {
        log.Fatalf("DB Error: %s", err)
    }

	fmt.Printf("Tokyoの人口は%d人です\n", city.Population)*/

	// take population
	var city City
	if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", os.Args[1]); errors.Is(err, sql.ErrNoRows) {
        log.Printf("no such city Name=%s", "Tokyo")
    } else if err != nil {
        log.Fatalf("DB Error: %s", err)
    }

	fmt.Printf("%sの人口は%d人です\n", os.Args[1], city.Population)

	// cities
	cities := []City{}
	db.Select(&cities, "SELECT * FROM city WHERE CountryCode='JPN'")

	fmt.Println("日本の都市一覧")
	for _, city := range cities {
		fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
	}
}