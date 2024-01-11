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
	w.Write([]byte("successful order insertion" + id))
}

func (c Controler) GetOrderById(w http.ResponseWriter, r http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	order, err := c.Store.OrderStorage.GetByIdOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is succus this id ` + id))
		return
	}

	js, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error while marshaling data ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro while marshaling " + err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(js)
}

func (c Controler) GetOrdersList(w http.ResponseWriter, r http.Request) {

	order, err := c.Store.OrderStorage.SelectOrdresulter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is error order` + err.Error()))
		log.Fatal("this is erro order package ")
		return
	}

	js, err := json.Marshal(order)
	if err != nil {
		fmt.Println("erroring while marshal order")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is success error`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func (c Controler) UpdateOrder(w http.ResponseWriter, r http.Request) {
	order := structfortable.Orders{}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error decoding update order request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request body: " + err.Error()))
		return
	}

	if err := c.Store.OrderStorage.UpdateOrder(order); err != nil {
		fmt.Println("error updating orders:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating orders: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successful ordrer update"))
}

func (c Controler) DeleteOrder(w http.ResponseWriter, r http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	if err := c.Store.OrderStorage.DeleteOrder(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is id success delete`))
		fmt.Println("delete this database ", id)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is sucess  delete`))
}
