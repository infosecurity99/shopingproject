package controler

import (
	"conectionmyprojectpath/structfortable"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controler) InsertOrderProducts() {
	price := 0
	fmt.Println("enter price ")
	fmt.Scan(&price)

	quantity := 0
	fmt.Println("enter quantity")
	fmt.Scan(&quantity)

	if err := c.Store.OrderProductStorage.CreateOrderProducts(price, quantity); err != nil {
		log.Fatal("this is error  from orderproductfunction", err.Error())
	}

	fmt.Println("success full order product insert")
}

func (c Controler) GetByIdOrderProducts() {
	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(idStr)

	id, err := uuid.Parse(idStr)

	products, err := c.Store.OrderProductStorage.GetByIdorderProduct(id)
	if err != nil {
		log.Fatal("this is eror  ")
	}

	fmt.Println(products)
}

func (c Controler) GetListOrderProducts() {
	orderproducts, err := c.Store.OrderProductStorage.GetListOrderProduct()
	if err != nil {
		log.Fatal("error getlistorderproducts", err.Error())
	}

	fmt.Println(orderproducts)
}

func (c Controler) DeleteOrderProduct() {
	var idStr string
	fmt.Println("enter id string")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal(err)
	}

	c.Store.OrderProductStorage.DeleteOrderProducts(id)

	fmt.Println("succes full delete")
}

func (c Controler) UpdateOrderProducts() {
	orderproduct := getInfoOrderProduct()
	if err := c.Store.OrderProductStorage.UpdateOrderProduct(orderproduct); err != nil {
		log.Fatal(err)
	}
	fmt.Println("success full   orderproduct")
}

func getInfoOrderProduct() structfortable.OrderProducts {
	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("error this is orderproduct")
	}

	price := 0
	fmt.Println("enter price")
	fmt.Scan(&price)

	quantity := 0
	fmt.Println("enter quantity")
	fmt.Scan(&quantity)

	return structfortable.OrderProducts{
		ID:       id,
		Price:    price,
		Quantity: quantity,
	}
}
