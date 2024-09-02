package repository

import (
	"database/sql"
	"peixe-urbano/src/models"
)

// Offers represents users repository
type Offers struct {
	db *sql.DB
}

// NewOfferRepository create repository of offer
func NewOfferRepository(db *sql.DB) *Offers {
	return &Offers{db}
}

// Create inster offer in database
func (repository Offers) Create(offer models.Offer) (uint64, error) {
	statement, err := repository.db.Prepare("insert into offers (categoria, titulo, descricao, anunciante, valor, destaque) values(?,?,?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(offer.Categoria, offer.Titulo, offer.Descricao, offer.Anunciante, offer.Valor, offer.Destaque)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return uint64(lastId), nil
}
