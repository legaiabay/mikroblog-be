package routers

import (
	"net/http"
	"os"

	article "github.com/mikroblog-be/controllers/api/v1/article"
	"github.com/mikroblog-be/helpers"
)

//SetupRouter -> Setup route with Go Standard Library
func InitRouter() {

	r := http.NewServeMux()

	//Redirect HTTP request
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	//Article
	r.HandleFunc("/article/", article.Get)
	r.HandleFunc("/article/add/", article.Add)
	r.HandleFunc("/article/delete/", article.Delete)

	helpers.Logger.Info("Server " + os.Getenv("SERVER_ENV") + " started at " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusPermanentRedirect)
}
