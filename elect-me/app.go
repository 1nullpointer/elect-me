package electme

import (
	"appengine"
	"net/http"
)

const basePath = "elect-me"

var offices []*office

func init() {
	getOffices()
	//basePath = "rapdemo"
	fs := http.FileServer(http.Dir(basePath + "/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	//api
	//http.Handle("/offices", appHandler(offices))

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
