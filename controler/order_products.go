package controler

import (
	"conectionmyprojectpath/structfortable"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Controler) OrderProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.InsertOrderProducts(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdOrderProducts(w, r)
		} else {
			c.GetListOrderProducts(w, r)
		}
	case http.MethodDelete:
		c.DeleteUser(w, r)
	case http.MethodPut:
		c.UpdateUsers(w, r)
	default:
		fmt.Println("this is not case ")
	}
}
func (c Controler) InsertOrderProducts(w http.ResponseWriter, r *http.Request) {
	orderproduct := structfortable.OrderProducts{}

	if err := json.NewDecoder(r.Body).Decode(&orderproduct); err != nil {
		fmt.Println([]byte(`this is while erroring newdecoder problame` + err.Error()))
		return
	}
	id, err := c.Store.OrderProductStorage.CreateOrderProducts(orderproduct)
	if err != nil {
		log.Fatal("this is error  from orderproductfunction", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is error orderproducts` + id))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is  create new data on insertorder` + id))
}

func (c Controler) GetByIdOrderProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	products, err := c.Store.OrderProductStorage.GetByIdorderProduct(id)
	if err != nil {
		log.Fatal("this is eror  ")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is erro` + err.Error()))
	}

	js, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`this is erro` + err.Error()))
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(js)
}

func (c Controler) GetListOrderProducts(w http.ResponseWriter, r *http.Request) {
	orderProducts, err := c.Store.OrderProductStorage.GetListOrderProduct()
	if err != nil {
		log.Println("Error while getting list of order products:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(orderProducts)
	if err != nil {
		log.Println("Error while marshaling order products:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (c Controler) DeleteOrderProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	if err := c.Store.OrderProductStorage.DeleteOrderProducts(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is id success delete`))
		fmt.Println("delete this database ", id)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is sucess  delete`))
}

func (c Controler) UpdateOrderProducts(w http.ResponseWriter, r *http.Request) {

	orderproduct := structfortable.Users{}

	if err := json.NewDecoder(r.Body).Decode(&orderproduct); err != nil {
		fmt.Println("error decoding update order product request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request body: " + err.Error()))
		return
	}

	if err := c.Store.UsersStorage.UpdateUser(orderproduct); err != nil {
		fmt.Println("error updating order product:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating order product: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successful order product update"))
}
