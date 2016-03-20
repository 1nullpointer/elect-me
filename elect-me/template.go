//The static pages are based on this tutorial
//http://www.alexedwards.net/blog/serving-static-sites-with-go

package electme

import (
	"appengine"
	//"appengine/log"
	//"appengine/cloudsql"
	//	"database/sql"
	"errors"

	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/ziutek/mymysql/godrv"
	//_ "github.com/ziutek/mymysql/mysql"
	//_ "github.com/ziutek/mymysql/native"
	//"golang.org/x/oauth2/google"
	"encoding/json"
	"html/template"
	//"log"
	"net/http"
	"os"
	"path"
)

/*
//sql testing still failing
func offices(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)

	datastoreName := os.Getenv("MYSQL_CONNECTION")
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return &appError{err, "Oauth failed", http.StatusInternalServerError}
	}

	ts, err := google.DefaultTokenSource(ctx, scope1, scope2, ...)
	  if err != nil {
	    // Handle error.
	  }
	  httpClient := oauth2.NewClient(ctx, ts)


	//maybe not...
	//appid := appengine.AppID(c)
	//db, err := sql.Open("mymysql", "cloudsql:elect-me-1255*election/grady/pass")
	db, err := sql.Open("mysql", datastoreName)

	rows, err := db.Query("SELECT * from offices limit 1")
	if err != nil {
		return &appError{err, "Query failed", http.StatusInternalServerError}
	}
	defer rows.Close()

	for rows.Next() {
		var firstName string
		var lastName string
		if err := rows.Scan(&firstName, &lastName); err != nil {
			return &appError{err, "Row scan failed", http.StatusInternalServerError}
			continue
		}
		c.Infof("First: %v - Last: %v", firstName, lastName)
	}
	if err := rows.Err(); err != nil {
		return &appError{err, "Row error", http.StatusInternalServerError}
	}

	fmt.Fprintf(w, "Got data!")
	return nil
}
*/

//it would be good for this function to pass a token to the page in case the page has a form (a lot of them will)
//it would also be good to default to a page if none is given in the r.URL.Path
func serveTemplate(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)

	c.Debugf("serving a non-static request")

	lp := path.Join(basePath+"/templates", "layout.html")
	fp := path.Join(basePath+"/templates", r.URL.Path)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return &appError{
				err,
				"Couldn't find the specified template",
				http.StatusNotFound,
			}
		}
		return &appError{err, "Unknown error finding template", http.StatusInternalServerError}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		return &appError{
			errors.New("Attempted to display directory " + r.URL.Path),
			"Can't display record",
			http.StatusNotFound,
		}
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

func getOffices() {
	tmp := []byte(data)

	if err := json.Unmarshal(tmp, &offices); err != nil {
		panic(err)
	}

	//for _,val:= range
}

type office struct {
	Position, Description, Requirments, Salary, Filing, ResidencyInDistrict, PetitionSignatures, MinAge, VoteCount string
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
