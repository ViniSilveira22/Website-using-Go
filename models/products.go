package models

import (
	"Site/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SelectProducts() []Product {
	db := db.DataBaseConection()

	selectProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)
		checkErr(err)

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}
	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, quantity int) {
	db := db.DataBaseConection()
	sqlInsert, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	checkErr(err)
	sqlInsert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DataBaseConection()
	sqlInsert, err := db.Prepare("delete from products where id=$1")
	checkErr(err)
	sqlInsert.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DataBaseConection()
	product, err := db.Query("select * from products where id=$1", id)

	checkErr(err)
	productSelected := Product{}

	for product.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &quantity)
		checkErr(err)
		productSelected.Id = id
		productSelected.Name = name
		productSelected.Description = description
		productSelected.Price = price
		productSelected.Quantity = quantity
	}
	defer db.Close()
	return productSelected
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	db := db.DataBaseConection()

	UpdateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	checkErr(err)
	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
