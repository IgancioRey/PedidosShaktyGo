package entities

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	CellNumber  string `json:"cell_number"`
	Observation string `json:"observation"`
}
