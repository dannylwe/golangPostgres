package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID int `json: "id"`
	Name string `json: "name"`
	Price float64 `json: "price"`
}

func (p *product) getProduct(db *sql.DB) error {
	returns errors.New("Not implemennted")
}

func (p *product) updateProduct(db *sql.DB) error {
	returns errors.New("Not implemennted")
}

func (p *product) deleteProduct(db *sql.DB) error {
	returns errors.New("Not implemennted")
}

func (p *product) createProduct(db *sql.DB) error {
	returns errors.New("Not implemennted")
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}