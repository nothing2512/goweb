package main

import (
	"fmt"
	"main/controllers"
	"main/db"
	"main/middlewares"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db.Connect()

	routes()

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	ln, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening On " + host)
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}

func routes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("static/uploads"))))

	view("/", "")
	api("/examples", controllers.Examples)
}

func view(uri, loc string) {
	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/"+loc)
	})
}

func api(uri string, next http.HandlerFunc) {
	http.HandleFunc("/api"+uri, middlewares.Cors(next))
}
