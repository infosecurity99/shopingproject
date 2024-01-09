package storage

import (
	"conectionmyprojectpath/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB           *sql.DB
	UsersStorage usersRepo
}

func New(con config.Config) (Store, error) {
	urls := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", con.PostgresHost, con.PostgresPort, con.PostgresUser, con.PostgresPassword, con.PostgresDB)

	db, err := sql.Open("postgres", urls)
	if err != nil {
		return Store{}, err
	}

	usersRepo := NewUserRepo(db)

	return Store{
		DB:           db,
		UsersStorage: usersRepo,
	}, nil
}
