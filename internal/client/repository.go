package client

import (
	"database/sql"

	"github.com/dahenao/ReservSoccerField/internal/domain"
)

type repoCliente struct {
	db *sql.DB
}

func NewRepoCliente(dbConnection sql.DB) Repository {
	return &repoCliente{
		db: &dbConnection,
	}
}

type Repository interface {
	//Create(cl domain.Client) (int, error)
	GetById(id int) (domain.Client, error)
	//GetAll() ([]domain.Client, error)
	//Update(cl domain.Client) error
	//Delete(id int) error
}

func (repo *repoCliente) GetById(id int) (domain.Client, error) {
	query := `SELECT id, cid, first_name, last_name,address FROM tbl_cliente where id = ? ;`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return domain.Client{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	cl := domain.Client{}

	err = row.Scan(&cl.ID, &cl.CID, &cl.FirstName, &cl.Cel, &cl.Address)
	if err != nil {
		return domain.Client{}, err

	}

	return cl, nil

}
