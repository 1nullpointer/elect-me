package electme

import (
	"appengine"
	"net/http"
	"path"
)

const basePath = "./" // This server should serve from the root

func init() {

	getOffices()

	fs := http.FileServer(http.Dir(path.Join(basePath + "/static")))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// This is a waste of whitespace and the and will unnecessarily increase
	// the compile time. This function could be handled directy.
	//
	// http.HandleFunc("/offices", Offices)
	//

	http.Handle("/offices", appHandler(Offices))

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
