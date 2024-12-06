package entity

import "errors"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("Required fields must be filled")
	}
	return nil
}
