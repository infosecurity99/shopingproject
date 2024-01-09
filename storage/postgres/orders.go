package storage

import (
	"conectionmyprojectpath/structfortable"
	"database/sql"
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
func (r orderRepo) InsertOrder(order structfortable.Orders) error {
	id := uuid.New()
	createdAt := time.Now()

	_, err := r.DB.Exec(`INSERT INTO orders VALUES ($1, $2, $3)`, id, order.Amount, createdAt)
	if err != nil {
		return err
	}

	return nil
}

//get by id
func (r  orderRepo) GetByIdOrder(id uuid.UUID)  {
	r.DB.QueryRow(`select from orders where id=$1`,id)
}

//select
func (r  orderRepo) SelectOrder()  {
	
}

//delete
func (r  orderRepo) DeleteOrder()  {
	
}

//update
func (r  orderRepo) UpdateOrder()  {
	
}
