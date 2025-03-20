package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        "time"

        "gorm.io/driver/mysql"
        "gorm.io/gorm"
)

type User struct {
        ID        uint      `gorm:"column:id;primaryKey"`
        Name      string    `gorm:"column:name"`
        Email     string    `gorm:"column:email"`
        Age       string    `gorm:"column:age"`
        CreatedAt time.Time `gorm:"column:createdAt"`
        UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func main() {
        dsn := "root:@tcp(localhost:3306)/openapi?charset=utf8mb4&parseTime=True&loc=Local"
        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
                fmt.Println("failed to connect database")
        }

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                var result User
                db.First(&result)
                response, _ := json.Marshal(result)
                fmt.Fprintf(w, string(response))
        })

        http.ListenAndServe(":3000", nil)
}