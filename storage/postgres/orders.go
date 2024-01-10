package storage

import (
	"conectionmyprojectpath/structfortable"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) orderRepo {
	return orderRepo{
		DB: db,
	}
}

//inser
func (r orderRepo) InsertOrder(order structfortable.Orders) (string, error) {
	id := uuid.New()
	createdAt := time.Now()

	_, err := r.DB.Exec(`INSERT INTO orders (id, amount, userid, createat) VALUES ($1, $2, $3, $4)`, id, order.Amount, order.UserId, createdAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

//get by id
func (r orderRepo) GetByIdOrder(id uuid.UUID) (structfortable.Orders, error) {
	order := structfortable.Orders{}
	rows := r.DB.QueryRow(`select from orders where id=$1`, id)

	if err := rows.Scan(&order.ID, &order.Amount, &order.UserId, &order.CreateAt); err != nil {
		return structfortable.Orders{}, err
	}
	return order, nil
}

//select
func (r orderRepo) SelectOrdresulter() ([]structfortable.Orders, error) {
	orders := []structfortable.Orders{}

	rows, err := r.DB.Query(`SELECT * FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := structfortable.Orders{}
		if err := rows.Scan(&order.ID, &order.Amount, &order.UserId, &order.CreateAt); err != nil {
			log.Fatal(err.Error())
		}
		orders = append(orders, order)
	}

	return orders, nil
}

//delete

func (r orderRepo) DeleteOrder(id uuid.UUID) error {
	_, err := r.DB.Exec(`DELETE FROM orders WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

/*
//update
func (r orderRepo) UpdateOrder(order structfortable.Orders) error {
	_, err := r.DB.Exec(`update orders set amount = $1, user_id = $2 where id = $3`, order.Amount, order.UserId, order.ID)
	if err != nil {
		return err
	}
	return nil
}
*/
