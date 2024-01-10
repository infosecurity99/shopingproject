package storage

import (
	"conectionmyprojectpath/structfortable"
	"database/sql"

	"github.com/google/uuid"
)

type orderproductRepo struct {
	DB *sql.DB
}

func NewOrderProduct(db *sql.DB) orderproductRepo {
	return orderproductRepo{
		DB: db,
	}
}

//insert orderproduct
func (r orderproductRepo) CreateOrderProducts(Price int, Quantity int) error {
	id := uuid.New()
	_, err := r.DB.Exec(`insert into order_products  values ($1, $2, $3 )`, id, Quantity, Price)
	if err != nil {
		return err
	}

	return nil
}

//getbyid  orderproduct
func (r orderproductRepo) GetByIdorderProduct(id uuid.UUID) (structfortable.OrderProducts, error) {
	orderproduct := structfortable.OrderProducts{}

	row := r.DB.QueryRow(`select from order_products where id=$1`, id)

	err := row.Scan(&orderproduct.ID, &orderproduct.Quantity, &orderproduct.Price)
	if err != nil {
		return structfortable.OrderProducts{}, nil
	}
	return orderproduct, nil
}

//getlist  orderproduct
func (r orderproductRepo) GetListOrderProduct() ([]structfortable.OrderProducts, error) {
	orderproducts := []structfortable.OrderProducts{}

	rows, err := r.DB.Query(`select*from order_products`)
	if err != nil {
		return []structfortable.OrderProducts{}, nil
	}

	for rows.Next() {
		orderproduct := structfortable.OrderProducts{}

		if err := rows.Scan(&orderproduct.ID, &orderproduct.Price, &orderproduct.Quantity); err != nil {
			return []structfortable.OrderProducts{}, nil
		}
		orderproducts = append(orderproducts, orderproduct)
	}

	return orderproducts, nil
}

//delete orderproduct
func (r orderproductRepo) DeleteOrderProducts(id uuid.UUID) error {
	_, err := r.DB.Exec(`delete from order_products`)
	if err != nil {
		return err
	}
	return nil
}

//update orderproduct

func (r orderproductRepo) UpdateOrderProduct(orderproduct structfortable.OrderProducts) error {
	_, err := r.DB.Exec(`update order_products set id=$1, price=$2, quantity=$3`, orderproduct.ID, orderproduct.Quantity, orderproduct.Price)
	if err != nil {
		return err
	}
	return nil
}
