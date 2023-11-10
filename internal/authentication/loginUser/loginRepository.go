package loginuser

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) isUserNotExist(username string) (bool, error) {
	var res bool
	err := r.db.Get(&res,
		`
	SELECT  EXISTS
	    (SELECT 1 FROM auth.users WHERE username = $1)
	`, username,
	)
	if err != nil {
		return false, err
	}
	return res, err
}

func (r *Repository) returnHashPassword(username string) ([]byte, error) {
	var res []byte

	err := r.db.Get(&res,
		`SELECT hash_password FROM auth.users WHERE username = $1
	`, username,
	)
	if err != nil {
		return nil, err
	}
	return res, err
}
