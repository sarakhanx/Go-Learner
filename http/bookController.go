package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id")) //in 'err' only use 'err' can not use 'error' or 'errors' it, was implement like this
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
	// return c.Status(fiber.StatusNotFound).SendString("Book not found") //! can return like this too if you want to use custom message return
}
func createBook(c *fiber.Ctx) error {
	//? ประกาศตัวแปรใหม่อิงโครงสร้าง(struct) เพื่อมารับค่าจาก body
	book := new(Book)
	//? เอาค่าจาก body ใส่ในตัวแปร bookโดยให้ bodyParsher แปลงค่าจาก JSON เป็น struct
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	//? return ค่ากลับไป
	books = append(books, *book)
	return c.JSON(book)
}
func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookUpdated := new(Book)
	if err := c.BodyParser(bookUpdated); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdated.Title
			books[i].Author = bookUpdated.Author
			// books[i] = book
			return c.JSON(books[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
func deleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
