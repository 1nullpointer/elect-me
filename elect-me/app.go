package electme

import (
	"net/http"
	//"time"

	"appengine"
)

const basePath = "elect-me"

func init() {
	//basePath = "rapdemo"
	fs := http.FileServer(http.Dir(basePath + "/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	//api
	//http.Handle("/positions", appHandler(resources))

	//handles the templated but otherwise mostly static html pages
	http.Handle("/", appHandler(serveTemplate))
}

//following the error pattern suggested in the Go Blog
//http://blog.golang.org/error-handling-and-go

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}
