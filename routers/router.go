package routers

import (
	"net/http"
	"os"

	"github.com/mikroblog/helpers"
)

//SetupRouter -> Setup route with Go Standard Library
func SetupRouter() {

	r := http.NewServeMux()

	//Redirect HTTP request
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	// r.HandleFunc("", )
	// r.HandleFunc("", )

	helpers.Logger.Info("Server " + os.Getenv("SERVER_ENV") + " started at " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusPermanentRedirect)
}
