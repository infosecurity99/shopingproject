package controler

import (
	"conectionmyprojectpath/structfortable"
	"fmt"

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

	id , err:= uuid.Parse(idStr)
	if err!=nil{
		fmt.Println("this is getbyidcretea .go  error  id")
		return
	}

	user, err:=c.Store.UsersStorage.GetByIdUser(id)

	fmt.Println(user)

}

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
