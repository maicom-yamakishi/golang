package models

import (
	"errors"
	"strings"
)

type Offer struct {
	ID         uint64   `json:"id,omitempty"`
	Categoria  string   `json:"categoria,omitempty"`
	Titulo     string   `json:"titulo,omitempty"`
	Descricao  string   `json:"descricao,omitempty"`
	Anunciante string   `json:"anunciante,omitempty"`
	Valor      float64  `json:"valor,omitempty"`
	Destaque   bool     `json:"destaque,omitempty"`
	Imagens    []string `json:"imagens,omitempty"`
}

// Prepare call methods to validate and format fields
func (offer *Offer) Prepare() error {
	if err := offer.validate(); err != nil {
		return err
	}
	offer.format()
	return nil
}

// validate check empty fields
func (offer *Offer) validate() error {
	if offer.Categoria == "" {
		return errors.New("o campo categoria é obrigatório e não pode ser vazio")
	}
	if offer.Titulo == "" {
		return errors.New("o campo titulo é obrigatório e não pode ser vazio")
	}
	if offer.Descricao == "" {
		return errors.New("o campo descricao é obrigatório e não pode ser vazio")
	}
	if offer.Anunciante == "" {
		return errors.New("o campo anunciante é obrigatório e não pode ser vazio")
	}
	if offer.Valor <= 0 {
		return errors.New("o campo valor é obrigatório e deve ser positivo")
	}
	return nil
}

// format remove spaces
func (offer *Offer) format() {
	offer.Categoria = strings.TrimSpace((offer.Categoria))
	offer.Titulo = strings.TrimSpace((offer.Titulo))
	offer.Descricao = strings.TrimSpace((offer.Descricao))
	offer.Anunciante = strings.TrimSpace((offer.Anunciante))
}
