package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pborman/uuid"
	"log"
	"time"
)

// go get github.com/lib/pq
// go get -u github.com/jinzhu/gorm
type Product struct {
	gorm.Model

	Code  string
	Price int `gorm:"default:200"`
}

type User struct {
	ID         string
	Name       string
	Department string `gorm:"default:'BIPRD'"`
	CreatedAt  time.Time
	Age        *int
}

// Take action on a reacord before inserting it in the database
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	// generate GUID
	user.ID = uuid.New()

	// Set Creation date to current timestamp
	user.CreatedAt = time.Now()

	return nil
}

// postgres / Go4thsas : 5432

// Create the database in the table to which you intend to connect using PGAdmin being one option (sinkar in this case)
func main() {
	// using 127.0.0.1 instead of dsinkarina4.apac.sas.com
	fmt.Println(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "127.0.0.1", "5432", "postgres", "sinkar", "Go4thsas", "disable"))
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "127.0.0.1", "5432", "postgres", "sinkar", "Go4thsas", "disable"))

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error while connecting to Postgres")
	}

	defer db.Close()

	// Create tables in public schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})

	// Create an entry
	db.Create(&Product{Code: "B001", Price: 20000})
	db.Create(&Product{Code: "A001"})
	var age int = 21
	db.Create(&User{Name: "Kaushal 2", Department: "Quest", Age: &age})
	age = 0
	db.Create(&User{Name: "Kaushal 3", Department: "BIPRD", Age: &age})
	// Read Entry
	var product Product
	var users []User

	//db.First(&product, "1") // Find product with 'id' equal to 1
	var user123 []User
	db.Find(&user123)
	//log.Fatal(users)
	log.Println(user123)

	db.Where("Department = ?", "BIPRD").Find(&users)
	//log.Fatal(users)
	log.Println(users)
	db.First(&product, "code = ?", "B001")

	log.Println(product)

	db.Where("Department in (?) AND Age = ?", []string{"Quest", "RQS"}, 0).Find(&users)
	log.Println("!!!: ", users)

	// Where clauses using Structs
	//age = 0
	db.Where(User{Department: "Quest", Age: &age}).Find(&users)
	log.Println(users)
	//log.Println("Age : ", *users[0].Age)

	// Update the record
	product.Price = 12345
	db.Save(&product)
	db.First(&product, "code = ?", "B001")
	log.Println("Final Product 2: ", product)

	product.Price = 1010101
	db.Model(&product).Update(&product)
	db.First(&product, "code = ?", "B001")
	log.Println("Final Product 2: ", product)

	db.Model(&product).Update("Price", 5000)
	db.First(&product, "code = ?", "B001")
	log.Println("Final Product: ", product)

}
