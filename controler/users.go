package controler

import (
	"conectionmyprojectpath/structfortable"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Controler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateInsert(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdUserCreate(w, r)
		} else {
			c.GetListUser(w, r)
		}
	case http.MethodDelete:
		c.DeleteUser(w, r)
	case http.MethodPut:
		//c.UpdateUsers(w, r)
	default:
		fmt.Println("this is not case ")
	}
}

//create insert  function
func (c Controler) CreateInsert(w http.ResponseWriter, r *http.Request) {

	users := structfortable.Users{}

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		fmt.Println([]byte(`this is while erroring newdecoder problame` + err.Error()))
		return
	}

	id, err := c.Store.UsersStorage.Insert(users)
	if err != nil {
		fmt.Println("errrothis insert create    users.go/controler")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this database add  new  data ` + id))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is creatStatus created` + id))
}

//create getbyid users

func (c Controler) GetByIdUserCreate(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	users, err := c.Store.UsersStorage.GetByIdUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is succus this id ` + id))
		return
	}

	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error while marshaling data ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro while marshaling " + err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(js)
}

//get list users
func (c Controler) GetListUser(w http.ResponseWriter, r *http.Request) {

	users, err := c.Store.UsersStorage.GetListUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is error getlistusers` + err.Error()))
		log.Fatal("this is erro getlistuser package ")
		return
	}

	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println("erroring while marshal users")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is success error`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(js)

}

//delete list users
func (c Controler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	if err := c.Store.UsersStorage.Deleteusers(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`this is id success delete`))
		fmt.Println("delete this database ", id)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`this is sucess  delete`))

}
func (c Controler) UpdateUsers(w http.ResponseWriter, r *http.Request) {

	userToUpdate := structfortable.Users{}

	if err := json.NewDecoder(r.Body).Decode(&userToUpdate); err != nil {
		fmt.Println("error decoding update user request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request body: " + err.Error()))
		return
	}

	if err := c.Store.UsersStorage.UpdateUser(userToUpdate); err != nil {
		fmt.Println("error updating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating user: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successful user update"))
}
