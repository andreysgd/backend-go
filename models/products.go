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

	selectEachProduct, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectEachProduct.Next() {
		var id int
		var name, description string
		var price float64
		var units int

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

func CreateNewProduct(name, description string, price float64, units int) {
	db := db.DataBaseConection()

	insertIntoDB, err := db.Prepare("insert into products(name, description, price, units) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertIntoDB.Exec(name, description, price, units)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DataBaseConection()

	deleteFromDB, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteFromDB.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DataBaseConection()

	dbProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for dbProduct.Next() {
		var id int
		var name, description string
		var price float64
		var units int

		err = dbProduct.Scan(&id, &name, &description, &price, &units)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Units = units
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, units int) {
	db := db.DataBaseConection()

	updateDB, err := db.Prepare("update products set name=$1, description=$2, price=$3, units=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(name, description, price, units, id)

	defer db.Close()
}
