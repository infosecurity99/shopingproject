package main

import (
	"conectionmyprojectpath/config"
	"conectionmyprojectpath/controler"
	storage "conectionmyprojectpath/storage/postgres"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.Load()

	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("Error while connecting to db err: ", err.Error())
		return
	}
	defer store.DB.Close()

	con := controler.New(store)

	http.HandleFunc("/users", con.Users)
	http.HandleFunc("/orders", con.Orders)
	http.HandleFunc("/products", con.Products)
	http.HandleFunc("/order_products", con.OrderProduct)

	http.ListenAndServe(":8080", nil)
	/*

		fmt.Println(`
		1-users
		2-orders
		3-products
		4-order products
		`)


		var choose int
		fmt.Scan(&choose)

		switch choose {

		case 1:

			fmt.Println(`
			1-create user
			2-get by id user
			3-select user
			4-delete user
			5-update user
			`)
			fmt.Printf("enter choose comand for user:=")
			fmt.Scan(&choose)

			switch choose {
			case 1:
				con.CreateInsert()
			case 2:
				con.GetByIdUserCreate()
			case 3:
				con.GetListUser()
			case 4:
				con.DeleteUser()
			case 5:
				con.UpdateUsers()
			default:
				fmt.Println("Command not found...")
			}

		case 2:

			var choosorder int
			fmt.Println("enter choosorder")
			fmt.Scan(&choosorder)

			switch choosorder {
			case 1:
				con.InsertOrders()
			case 2:
				con.GetOrderById()
			case 3:
				con.GetOrdersList()
			case 4:
				con.UpdateOrder()
			case 5:
				con.DeleteOrder()

			}

		case 3:
			var choosproduct int
			fmt.Println("enter choosproduct")
			fmt.Scan(&choosproduct)

			switch choosproduct {
			case 1:
				con.CreateProduct()
			case 2:
				con.GetByIdProduct()
			case 3:
				con.GetOrdersList()
			case 4:
				con.DeleteProduct()
			case 5:
				con.UpdateProduct()
			default:
				fmt.Println("this is not case")
			}
		case 4:
			var choosorderproduct int
			fmt.Println("enter chooseorderproduct")
			fmt.Scan(&choosorderproduct)

			switch choosorderproduct {
			case 1:
				con.InsertOrderProducts()
			case 2:
				con.GetByIdOrderProducts()
			case 3:
				con.GetListOrderProducts()
			case 4:
				con.DeleteOrderProduct()
			case 5:
				con.UpdateOrderProducts()
			default:
				fmt.Println("this is not case n in   chooseorderproduct")
			}


		default:
			fmt.Println("this is not case ")
		}*/
}
