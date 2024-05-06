package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book //slice not Array

func main() {
	app := fiber.New()
	//! Cors Config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi Folks!")
	})
	//todo : Router
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Post("/upload", uploadfile)

	app.Listen(":8080")
}
func uploadfile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	err = c.SaveFile(file, "./assets/"+file.Filename)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File uploaded successfully")
}

// todo : Controller
