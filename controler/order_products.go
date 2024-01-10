package controler

import (
	"fmt"
	"log"
)

func (c Controler) InsertOrderProducts() {
	price := 0
	fmt.Println("enter price ")
	fmt.Scan(&price)

	quantity := 0
	fmt.Println("enter quantity")
	fmt.Scan(quantity)

	if err := c.Store.OrderProductStorage.CreateOrderProducts(price, quantity); err != nil {
		log.Fatal("this is error  from orderproductfunction", err.Error())
	}

	fmt.Println("success full order product insert")
}
