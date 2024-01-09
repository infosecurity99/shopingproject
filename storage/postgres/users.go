package storage

import (
	"conectionmyprojectpath/structfortable"
	"database/sql"

	"github.com/google/uuid"
)

type usersRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) usersRepo {
	return usersRepo{
		DB: db,
	}
}

// insert function users;
func (r usersRepo) Insert(user structfortable.Users) (string, error) {
	id := uuid.New()

	if _, err := r.DB.Exec(`insert into users values ($1, $2, $3, $4, $5)`, id, user.FirstName, user.LastName, user.Email, user.Phone); err != nil {
		return "", nil
	}
	return id.String(), nil
}

// getby id  function users
func (r usersRepo) GetByIdUser(id uuid.UUID) (structfortable.Users, error) {
	user := structfortable.Users{}
	row := r.DB.QueryRow(`select id, firstname, lastname, email, phone from users where id = $1`, id)

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone); err != nil {
		return structfortable.Users{}, err
	}
	return user, nil
}

//getlist user

func (r usersRepo) GetListUser() ([]structfortable.Users, error) {
	users := []structfortable.Users{}
	rows, err := r.DB.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := structfortable.Users{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

//delete funct   for delete  user
func (r usersRepo) Deleteusers(id uuid.UUID) error {

	_, err := r.DB.Exec(`delete from users where id=$1`, id)
	if err != nil {

		return err
	}
	return nil
}

//update  user for iuser

func (r usersRepo) UpdateUser(user structfortable.Users) error {
	_, err := r.DB.Exec(`update users set firstname=$1 ,lastname=$2 ,email=$3, phone=$4, id=$5 `, user.FirstName, user.LastName, user.Email, user.Phone, user.ID)
	if err != nil {
		return err
	}
	return nil
}
