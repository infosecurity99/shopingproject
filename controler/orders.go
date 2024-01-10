package controler

import (
	"conectionmyprojectpath/structfortable"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (c Controler) Orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Createorders(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			//	c.GetOrderById(w, r)
		} else {
			//	c.GetOrdersList(w, r)
		}
	case http.MethodDelete:
		//c.DeleteOrder(w, r)
	case http.MethodPut:
		//c.UpdateOrder(w, r)
	default:
		fmt.Println("this is not case")
	}
}

func (c Controler) Createorders(w http.ResponseWriter, r *http.Request) {
	var order structfortable.Orders

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error decoding order insert request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request body: " + err.Error()))
		return
	}

	id, err := c.Store.OrderStorage.InsertOrder(order)
	if err != nil {
		fmt.Println("error inserting order into the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error inserting order into the database: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successful order insertion"+id))
}

func (c Controler) GetOrderById(w http.ResponseWriter, r http.Request) {
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

func (c Controler) GetOrdersList(w http.ResponseWriter, r http.Request) {
	orders, err := c.Store.OrderStorage.SelectOrdresulter()
	if err != nil {
		log.Fatal("Error while get orders list: ", err.Error())
	}
	fmt.Println(orders)
}

/*
func (c Controler) UpdateOrder(w http.ResponseWriter, r http.Request) {
	orders := ForUpdateOrders()
	if err := c.Store.OrderStorage.UpdateOrder(orders); err != nil {
		log.Fatal("Error while updating order ", err.Error())
	}
	fmt.Println("order updated succesfully...")
}*/
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
		Amount: amount,
		UserId: id,
	}
}

func (c Controler) DeleteOrder(w http.ResponseWriter, r http.Request) {
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
