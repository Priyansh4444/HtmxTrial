package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		if session.Values["counter"] == nil {
			session.Values["counter"] = 0
		}
		session.Values["counter"] = session.Values["counter"].(int) + 1
		session.Save(r, w)
		fmt.Fprintf(w, "%d", session.Values["counter"])

		// Print count in CLI
		fmt.Println("Count:", session.Values["counter"])
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../public/index.html")
	})

	http.ListenAndServe(":8080", nil)
}