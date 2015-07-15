package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strings"
)

var hostname string

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	var err error
	hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	if strings.Contains(hostname, "local") {
		orm.Debug = true
		orm.RegisterDataBase("default", "postgres", "user=rompei dbname=bgm sslmode=disable")
	} else {
		orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
	}
}

func main() {
	router := NewRouter()

	if strings.Contains(hostname, "local") {
		log.Fatal(http.ListenAndServe(":9000", router))
	} else {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
	}
}
