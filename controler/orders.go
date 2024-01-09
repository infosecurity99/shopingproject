package controler

import (
	"conectionmyprojectpath/structfortable"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controler) InsertOrders() {
	var amount int
	fmt.Println("Enter amount: ")
	fmt.Scan(&amount)
	var idStr string
	fmt.Println("Enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("This id is not uuid type ", err.Error())
	}
	if err := c.Store.OrderStorage.InsertOrder(amount, id); err != nil {
		log.Fatal("Error while insert orders ", err.Error())
	}
	fmt.Println("Order inserting succesfully...")
}

func (c Controler) GetOrderById() {
	var idStr string
	fmt.Println("Enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("This id is not uuid type ", err.Error())
	}
	order, err := c.Store.OrderStorage.GetByIdOrder(id)
	if err != nil {
		log.Fatal("Error while get order by id err: ", err.Error())
	}
	fmt.Println(order)
}

func (c Controler) GetOrdersList() {
	orders, err := c.Store.OrderStorage.SelectOrdresulter()
	if err != nil {
		log.Fatal("Error while get orders list: ", err.Error())
	}
	fmt.Println(orders)
}

func (c Controler) UpdateOrder() {
	orders := ForUpdateOrders()
	if err := c.Store.OrderStorage.UpdateOrder(orders); err != nil {
		log.Fatal("Error while updating order ", err.Error())
	}
	fmt.Println("order updated succesfully...")
}
func ForUpdateOrders() structfortable.Orders {
	var amount int
	fmt.Println("Enter amount: ")
	fmt.Scan(&amount)
	var idStr string
	fmt.Println("Enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("This id is not uuid type ", err.Error())
	}
	return structfortable.Orders{
		Amount:  amount,
		User_Id: id,
	}
}

func (c Controler) DeleteOrder() {
	var idStr string
	fmt.Println("Enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("This id is not uuid type ", err.Error())
	}
	if err := c.Store.OrderStorage.DeleteOrder(id); err != nil {
		log.Fatal("Error while delete order...")
	}
	fmt.Println("Order succesfully deleted...")
}
