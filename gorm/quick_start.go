package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect databases")
	}
	Db = db
}

func main() {
	//create()
	//batchCreate()
	query()
	//page()
}

func page() {
	var product []Product
	Db.Scopes(paginate(1, 3)).Find(&product)
	fmt.Println(product)
}

func paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func query() {
	var product Product
	var products []Product

	fmt.Println("------------------获取第一条-------------------------")
	// 获取第一条
	Db.First(&product)
	fmt.Println(product)
	fmt.Println("------------------获取最后条-------------------------")
	// 获取最后条
	var product2 Product
	Db.Last(&product2)
	fmt.Println(product2)

	fmt.Println("------------------批量查询-------------------------")
	// 批量获取
	Db.Find(&products, []int{1, 3})
	fmt.Println(products)

	fmt.Println("------------------根据主键查询-------------------------")
	// 根据主键查询
	var product3 Product
	Db.First(&product3, "id = ?", 1)
	fmt.Println(product3)
	fmt.Println("------------------获取全部记录-------------------------")

	res := Db.Find(&products)
	fmt.Println(res.RowsAffected, res.Error, products)

	fmt.Println("------------------倒序-------------------------")
	Db.Order("price desc").Find(&products)
	fmt.Println(products)

	fmt.Println("-------------------Not--------------------------")
	var product4 []Product
	Db.Not("code = ?", "001").Find(&product4)
	fmt.Println(product4)

}

func batchCreate() {
	products := []Product{{Price: 10, Code: "001"}, {Price: 20, Code: "002"}, {Price: 30, Code: "003"}}
	Db.Create(&products)
	for _, product := range products {
		fmt.Println(product.ID)
	}
}

func create() {
	product := Product{Code: "F01", Price: 199}
	result := Db.Create(&product) // 通过数据的指针来创建

	id := product.ID                // 返回插入数据的主键
	err := result.Error             // 返回 error
	affected := result.RowsAffected // 返回插入记录的条数
	fmt.Println(id, err, affected)
}
