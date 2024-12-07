package entity

import "errors"

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
	Stock      int     `json:"stock"`
	ImageURL   string  `json:"image_url"`
}

type GetProduct struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
	Stock    int      `json:"stock"`
	Sold     int      `json:"sold"`
	ImageURL string   `json:"image_url"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("Required fields must be filled")
	}
	if p.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}
	return nil
}
