package model

type Cart struct {
	Id         int64  `json:"id"`
	CategoryId int64 `json:"categoryId"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	Quantity   int64  `json:"quantity"`
	CreatedAt  int64  `json:"createdAt"`
	ModifiedAt int64  `json:"modifiedAt"`
}

