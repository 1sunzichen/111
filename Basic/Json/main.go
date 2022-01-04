package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	Id       string       `json:"id"`
	Name     string       `json:"name,omitempty"`
	Quantity string       `json:"quantity"`
	Total    string       `json:"total_price"`
	Item     *[]OrderItem `json:"item"`
}

func main() {
	o := Order{
		Id:    "1",
		Total: "30",
		Item: &[]OrderItem{
			{
				Id:    "111",
				Name:  "apple",
				Price: 16,
			},
			{
				Id:    "1112",
				Name:  "banana",
				Price: 26,
			},
		},
	}
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", b)

}
