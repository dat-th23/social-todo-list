package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(db)

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "This is item 1",
		Description: "This is item 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	r := gin.Default()

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items/ (Create a new item)
	// GET /v1/items (list items) v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT || PATCH) /v1/items/:id (Update an item by id)
	// DELETE /v1/items/:id (Delete item by id)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("")
			items.GET("")
			items.GET("/:id")
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})

	r.Run()
}
