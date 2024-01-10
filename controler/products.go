package controler

import (
	"conectionmyprojectpath/structfortable"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Controler) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateProducts(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdProduct(w, r)
		} else {
			c.SelectProducts(w, r)
		}
	case http.MethodDelete:
	     c.DeleteProduct(w, r)
	case http.MethodPut:
		c.UpdateProduct(w, r)
	default:
		fmt.Println("this is not case")
	}
}

func (c Controler) CreateProducts(w http.ResponseWriter, r *http.Request) {

	products := structfortable.Products{}

	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		fmt.Println([]byte(`this is while erroring newdecoder problame` + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`this is suscces ` + err.Error()))
		return
	}

	id, err := c.Store.ProductStorage.CreateProductInsert(products)
	if err != nil {
		log.Fatal("Error while insert product ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is suscces ` + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is suces` + id))
}

func (c Controler) GetByIdProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	product, err := c.Store.ProductStorage.GetByidProduct(id)

	if err != nil {
		log.Fatal("eror  this is cannot product", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is error ` + id))
	}

	js, err := json.Marshal(product)
	if err != nil {
		fmt.Println("THIS IS ERROR marshal")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`this is erro `))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func (c Controler) SelectProducts(w http.ResponseWriter, r *http.Request) {
	product, err := c.Store.ProductStorage.SelectProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is error getlistusers` + err.Error()))
		log.Fatal("this is erro product package ")
		return
	}

	js, err := json.Marshal(product)
	if err != nil {
		fmt.Println("erroring while marshal product")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is success error`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func (c Controler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	value := r.URL.Query()
	id := value["id"][0]

	if err := c.Store.ProductStorage.DeleteProduct(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is erroing delete product`))
		fmt.Println("error this is deleteproducts")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`this is erroing delete product` + id))
}


func (c Controler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := structfortable.Products{}
	c.Store.ProductStorage.UPdateProduct(product)


	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error decoding update product request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request product body: " + err.Error()))
		return
	}

	if err := c.Store.ProductStorage.UPdateProduct(product); err != nil {
		fmt.Println("error updating product:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating product: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successful user update"))
}

