package api

import "github.com/front-go/gateway/internal/model"

type SrvI interface {
	Summator(in *model.In)
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Passport struct {
	Number string `json:"number"`
	Series string `json:"series"`
}

type FullName struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// {
// 	"passport": {
// 	  "number": "1234567890",
// 	  "series": "AB"
// 	},
// 	"address": {
// 	  "street": "Улица Ленина",
// 	  "city": "Москва",
// 	  "state": "Москва"
// 	},
// 	"fullname": {
// 	  "name": "Иван",
// 	  "surname": "Иванов"
// 	}
//   }
