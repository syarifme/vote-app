package main

import (
	"fmt"
	"net/http"

	"github.com/syarifme/vote-app/db"
	"github.com/syarifme/vote-app/handlers"
)

func main() {
	http.HandleFunc("/", handlers.UserIndex)
	gormDB := db.Connect()
	db.Migrate(gormDB)
	defer db.Connect().Close()

	fmt.Println("listen localhost:9090")

	http.ListenAndServe("localhost:9090", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
