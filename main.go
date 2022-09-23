package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/book"
	"pustaka-api/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=belajar-gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error")
	}

	db.AutoMigrate(&book.Book{})

	// Menambahkan data
	book := book.Book{}
	book.Title = "Main tiger episode 2"
	book.Price = 90000
	book.Description = "This book about tiger wong"
	book.Rating = 5

	err = db.Create(&book).Error

	if err != nil {
		log.Fatal("error ketika menambahkan data baru")
	}

	// Mengambil data
	err = db.Debug().Where("id = ?", 2).Find(&book).Error

	if err != nil {
		log.Fatal("error ketika mengambil data")
	}

	// fmt.Println(book)

	// Update data
	book.Title = "Main tiger is not tiger wong"
	err = db.Save(&book).Error
	if err != nil {
		log.Fatal("error untuk update data")
	}

	// Hapus data
	err = db.Debug().Delete(&book, 2).Error
	if err != nil {
		log.Fatal("gagal menghapus data")
	}

	// for _, b := range books {
	// 	fmt.Println("Title: ", b.Title)
	// 	fmt.Println("Book Object: %v", b)
	// }

	r := gin.Default()
	r.GET("/", handler.Home)
	r.GET("/books", handler.ShowBooks)
	r.POST("/book", handler.AddBook)
	r.Run()

}
