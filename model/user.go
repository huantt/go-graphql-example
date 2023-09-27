package model

type User struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	} `json:"geo"`
}
