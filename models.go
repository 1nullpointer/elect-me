package electme

type Office struct {
	Id                  int    `json: "Id"`
	FirstName           string `json: "FirstName"`
	LastName            string `json: "LastName"`
	Position            string `json: "Position"`
	Description         string `json: "Description"`
	Requirments         string `json: "Requirments"`
	Salary              string `json: "Salary"`
	Filing              string `json: "Filing"`
	ResidencyInDistrict string `json: "ResidencyInDistrict"`
	PetitionSignatures  string `json: "PetitionSignatures"`
	MinAge              string `json: "MinAge"`
	VoteCount           string `json: "VoteCount"`
}
