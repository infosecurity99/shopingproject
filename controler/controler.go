package controler

import storage "conectionmyprojectpath/storage/postgres"

type Controler struct {
	Store storage.Store
}

func New(store storage.Store) Controler {
	return Controler{
		Store: store,
	}
}
