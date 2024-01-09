package storage

import (
	"conectionmyprojectpath/structfortable"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) productRepo {
	return productRepo{
		DB: db,
	}
}

func (r productRepo) CreateProductInsert(price int, name string) error {
	id := uuid.New()
	if _, err := r.DB.Exec(`insert into products values ($1, $2, $3)`, id, price, name); err != nil {
		return err
	}
	return nil
}

func (c productRepo) GetByidProduct(id uuid.UUID) (structfortable.Products, error) {
	product := structfortable.Products{}
	row := c.DB.QueryRow(`select *from products where id=$1`, id)

	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {

		return structfortable.Products{}, nil
	}

	return product, nil

}

func (c productRepo) SelectProduct() ([]structfortable.Products, error) {
	products := []structfortable.Products{}

	rows, err := c.DB.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := structfortable.Products{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (c productRepo) DeleteProduct(id uuid.UUID) error {
	_, err := c.DB.Exec(`delete from product where id=$1`, id)
	if err != nil {
		log.Fatal("erorr this is productRpository")
	}
	return nil
}

func (c productRepo) UPdateProduct(product structfortable.Products) error {

	_, err := c.DB.Exec(`update products set price = $1, product_name = $2 where id = $3`, product.Price, product.Name, product.ID)
	if err != nil {
		return err
	}
	return nil
}
