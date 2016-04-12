//The static pages are based on this tutorial
//http://www.alexedwards.net/blog/serving-static-sites-with-go

// Serve the application form the `app.yaml` file.
// When running locally, use port 8080 so that the server is running behind the gateway.
// Disregard if there is no gateway running/configured locally on your system.
//
// dev_appserver.py --port=8080 elect-me-1255

package electme

import (
	"appengine"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
)

// Below is the proper syntax for placing objects into an array.
// However, this variable is not necessary because the mySQL driver
// does not support the use of a data structure, at least to my knowledge.
// All of the data coming in from the query must be place individually into
// a variable and then the driver points each piece of datum to the proper
// variable.

var offices []Office

func Offices(w http.ResponseWriter, r *http.Request) *appError {

	// At this time a there is no database present at this URI
	// cloudsql:elect-me-1255*election/grady/pass. Therefore, the authentication
	// below is not necessary. There were also some issues with this auth.
	// The previous snippet was attempting to create an authentication compatible
	// with Google's system (i.e. encryptions, etc.). The snippet below creates an
	// authentication token, which is stored in a variable below called `client`, to be
	// sent to the database server then a connection to mySQL can be established.
	//
	// "golang.org/x/net/context"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	//
	// ctx := context.TODO()
	//
	// ts, err := google.DefaultTokenSource(ctx, "token")
	//
	// if err != nil {
	//	 return &appError{err, "The Google Default Token Doesn't Exist", http.StatusInternalServerError}
	// }
	//
	// client := oauth2.NewClient(ctx, ts)
	//
	// "cloudsql:elect-me-1255*election/grady/pass"
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	db, err := sql.Open("mysql", "@tcp(127.0.0.1:3306)/test")

	if err != nil {
		return &appError{err, "Database Connection Error", http.StatusInternalServerError}
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return &appError{err, "Database Connection Error After Connection Established", http.StatusInternalServerError}
	}

	rows, err := db.Query("SELECT * from offices") // limit 1

	if err != nil {
		return &appError{err, "Database Query Failed", http.StatusInternalServerError}
	}

	c := appengine.NewContext(r)

	for rows.Next() {

		var (
			uid       int
			firstname string
			lastname  string
		)

		if err := rows.Scan(&uid, &firstname, &lastname); err != nil {
			return &appError{err, "Table Row Scan Failed", http.StatusInternalServerError}
			continue
		}

		c.Infof("First: %v - Last: %v", firstname, lastname)

	}

	if err := rows.Err(); err != nil {
		return &appError{err, "Table Row Error", http.StatusInternalServerError}
	}

	fmt.Fprintf(w, "Got Data")

	return nil
}

//it would be good for this function to pass a token to the page in case the page has a form (a lot of them will)
func serveTemplate(w http.ResponseWriter, r *http.Request) *appError {

	// Right now the server will redirect to `index.html` if the request URL
	// is `/`. However, this is not very good for the server. What should be
	// happening is this server should be handling routes, which right now it
	// is not. That looks like the following:
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//   w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	//   w.WriteHeader(http.StatusOK)
	//   w.Write([]byte("Hello World"))
	// })
	//
	// This is a callback function that points to the HTTP request and response.
	// Once the URL path request is made and there is a an adequate TCP connection,
	// the server will pass the request to the handler and the handle will use HTTP
	// to serve the request.

	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 301)
	}

	c := appengine.NewContext(r)

	c.Debugf("serving a non-static request")

	lp := path.Join(basePath+"/templates", "layout.html")
	fp := path.Join(basePath+"/templates", r.URL.Path)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return &appError{err, "Couldn't find the specified template", http.StatusNotFound}
		}
		return &appError{err, "Unknown error finding template", http.StatusInternalServerError}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		return &appError{errors.New("Attempted to display directory " + r.URL.Path), "Can't display record", http.StatusNotFound}
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		return &appError{err, "Error parsing template", http.StatusInternalServerError}
	}

	//log.Println(len(offices))
	if err := tmpl.ExecuteTemplate(w, "layout", offices); err != nil {
		return &appError{err, "Error executing template", http.StatusInternalServerError}
	}
	return nil
}

var data = `
[{"Salary": "95$", "VoteCount": "0", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-29-17", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "0", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-29-17", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-58-44", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-21-35", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-59-14", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-2-14", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-41-22", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-50-27", "Filing": "0$"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-22-29", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-17-16", "Filing": "$0"}]
`

func getOffices() {
	tmp := []byte(data)

	if err := json.Unmarshal(tmp, &offices); err != nil {
		panic(err)
	}

	//for _,val:= range
}
