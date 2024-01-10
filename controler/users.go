package controler

import (
	"conectionmyprojectpath/structfortable"
	"fmt"
	"log"

	"github.com/google/uuid"
)

//create insert  function
func (c Controler) CreateInsert() {
	user := GetInfoUsers()

	id, err := c.Store.UsersStorage.Insert(user)
	if err != nil {
		fmt.Println("errrothis insert create    users.go/controler")
		return
	}

	fmt.Println("users id", id)
}

//create getbyid users

func (c Controler) GetByIdUserCreate() {
	var idStr string
	fmt.Println("enter idStr")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("this is getbyidcretea .go  error  id")
		return
	}

	user, err := c.Store.UsersStorage.GetByIdUser(id)

	fmt.Println(user)
}

//get list users
func (c Controler) GetListUser() {

	users, err := c.Store.UsersStorage.GetListUser()
	if err != nil {
		log.Fatal("this is erro getlistuser package ")
	}

	fmt.Println(users)
}

//delete list users
func (c Controler) DeleteUser() {
	var idStr string
	fmt.Printf("enter str for delete by user")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("this is error id str", id)
	}
	c.Store.UsersStorage.Deleteusers(id)
	fmt.Println("delete this database ", id)
}
func (c Controler) UpdateUsers() {
	user := GetInfoUsersforUpdate()
	c.Store.UsersStorage.UpdateUser(user)

	fmt.Println("successful update this data ")
}
//selectall
func GetInfoUsers() structfortable.Users {

	var (
		firstname string
		lastname  string
		phone     string
		email     string
	)
	fmt.Printf("enter first name")
	fmt.Scan(&firstname)

	fmt.Printf("enter last name")
	fmt.Scan(&lastname)

	fmt.Printf("enter phone")
	fmt.Scan(&phone)

	fmt.Printf("enter email")
	fmt.Scan(&email)

	return structfortable.Users{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
	}
}
//getinfoforupdateusers
func GetInfoUsersforUpdate() structfortable.Users {
	var idStr string
	fmt.Printf("enter id  for update")
	fmt.Scan(&idStr)
	var (
		firstname string
		lastname  string
		phone     string
		email     string
	)
	fmt.Printf("enter first name")
	fmt.Scan(&firstname)

	fmt.Printf("enter last name")
	fmt.Scan(&lastname)

	fmt.Printf("enter phone")
	fmt.Scan(&phone)

	fmt.Printf("enter email")
	fmt.Scan(&email)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatal("this is eror", err.Error())
	}

	return structfortable.Users{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
	}
}
