package main

import (
	"fmt"
	"net/http"

	"github.com/arieljimenez/mybancaAPI/src/models/user"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", user.HandleGetUsers)
	appengine.Main()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "MyBanca - API")
}
