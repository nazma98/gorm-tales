package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{ID: 123, Name: "Shimul", Age: 13})

	// Read
	var user User
	db.First(&user, 1)
	db.First(&user, "name = ?", "Shimul")
	fmt.Println("📝 User Found:", user)

	// Update
	db.Model(&user).Update("Age", 14)
	fmt.Println("✅ User Updated: ", user)
	// Update - multiple fields
	db.Model(&user).Updates(User{Age: 14, Name: "Fatema Shimul"})
	db.Model(&user).Updates(map[string]interface{}{"Age": 14, "Name": "Fatema Shimul"})

	// Delete - delete product
	db.Delete(&user, 1)
	fmt.Println("❌ User Deleted")
}
