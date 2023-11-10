package registeruser

import (
	"auth/internal/entities"

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

func (r *Repository) isUserExistsAndEmailExist(username, email string) (bool, error) {
	var res bool
	err := r.db.Get(&res,
		`
	SELECT  EXISTS
	    (SELECT 1 FROM auth.users WHERE username = $1 AND email = $2)
	`, username, email,
	)
	if err != nil {
		return false, err
	}
	return res, err
}

func (r *Repository) saveUser(user entities.User) error {
	_, err := r.db.Query(
		`
		INSERT INTO auth.users (username, hash_password, email, fullname)
		VALUES($1, $2, $3, $4)
		RETURNING user_id 
		`,
		user.Username,
		user.Password,
		user.Email,
		user.FullName,
	)
	if err != nil {
		return err
	}
	return nil
}

type Registration struct {
	Email          string
	HashedPassword string
	Username       string
	FullName       string
}
