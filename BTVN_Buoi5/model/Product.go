package model

type Product struct {
	Id         int64    `json:"id"`
	CategoryId int64    `json:"categoryId"`
	Image      []string `json:"image"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Sale     uint8     `json:"isSale"`
	CreatedAt  int64    `json:"createdAt"`
	ModifiedAt int64    `json:"modifiedAt"`
	Rate       float32  `json:"rate"`
	
}

//  	"id": 1,
//      "categoryId": 2,
//      "image": "/uploads/images/item-02.jpg",
//      "name": "Herschel supply co 25l",
//      "price": 75,
//      "isSale": true,
//      "createdAt": 1614362898000,
//      "modifiedAt": 1615410795000
