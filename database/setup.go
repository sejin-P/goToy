package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=ab180 password=ab180!*)"+
		" dbname=ab180 host=5000 sslmode=disable")

	if db != nil {
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(10)
	}
	if err != nil {
		return nil, err
	}
	return db, err
}

func Time(c *gin.Context) {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var time string
	err = db.QueryRow("SELECT now()").Scan(&time)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"Time": time})
}
