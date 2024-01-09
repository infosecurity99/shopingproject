package controler

import (
	"conectionmyprojectpath/structfortable"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controler) CreateProduct() {

	price := 0
	fmt.Println("Enter amount: ")
	fmt.Scan(&price)

	var name string
	fmt.Println("enter name")
	fmt.Scan(&name)

	if err := c.Store.ProductStorage.CreateProductInsert(price, name); err != nil {
		log.Fatal("Error while insert product ", err.Error())
	}
	fmt.Println("product inserting succesfully...")
}

func (c Controler) GetByIdProduct() {
	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("error this is not reponse  get by id", err.Error())
		return
	}

	product, err := c.Store.ProductStorage.GetByidProduct(id)

	if err != nil {
		log.Fatal("eror  this is cannot product", err.Error())
	}

	fmt.Println(product)
}

func (c Controler) SelectProducts() {

}

func (c Controler) DeleteProduct() {

	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Fatal("error this is deleteproduct")
	}
	c.Store.OrderStorage.DeleteOrder(id)

	fmt.Println("successfukldekete product data")
}

func (c Controler) UpdateProduct() {
	product := GetProductInfo()
	c.Store.ProductStorage.UPdateProduct(product)
}

func GetProductInfo() structfortable.Products {
	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("error  getproductinfo idStr")
	}

	var (
		name  string
		price int
	)

	fmt.Println("enter name")
	fmt.Scan(&name)

	fmt.Println("enter price")
	fmt.Scan(&price)

	return structfortable.Products{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
