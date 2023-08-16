package models

import "github.com/web-application/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Units       int
}

func GetProducts() []Product {
	db := db.DataBaseConection()

	selectEachProduct, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectEachProduct.Next() {
		var id, units int
		var name, description string
		var price float64

		err = selectEachProduct.Scan(&id, &name, &description, &price, &units)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Units = units

		products = append(products, p)
	}
	defer db.Close()
	return products
}
