package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  int
}

func main() {
	dsn := "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	newUser := User{Name: "Alice", Age: 25}
	db.Create(&newUser)
	fmt.Printf("Created User: %+v\n", newUser)

}
