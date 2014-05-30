package masterdex

type DataStore interface {
	Find(id int) interface{}
	FindAll() interface{}
}

type Pokemon struct {
	DexNumber   int    `json:"dex_number"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
