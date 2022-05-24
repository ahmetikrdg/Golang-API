package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Customer struct {
	Name     string `json:"full_name"`
	Username string `json:"username"`
	Phone    string `json:"phone_number"`
}

var customers = []Customer{ //fake data
	{Name: "Mert Sahin", Username: "mertshn", Phone: "05433434343"},
	{Name: "Tony Stark", Username: "ironMan", Phone: "3423423423"},
	{Name: "Blue", Username: "red", Phone: "443523423"},
}

func GetAllCustomer(c *fiber.Ctx) error {
	if len(customers) <= 0 {
		return c.Status(http.StatusNotFound).JSON("no data")
	}
	return c.Status(http.StatusOK).JSON(customers)
}

func GetByUsernameWithData(c *fiber.Ctx) error {
	username := c.Params("username")

	for _, element := range customers {
		if username == element.Username {
			return c.Status(http.StatusOK).JSON(element)
		}
	}

	return c.Status(http.StatusNotFound).JSON("Username you were looking for was not found :(")
}

func CreateCustomer(c *fiber.Ctx) error {
	var newCustomers Customer

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Message: Bad Request")
	}

	customers = append(customers, newCustomers)
	return c.Status(http.StatusOK).JSON("Success Created")
}

func UpdateCustomer(c *fiber.Ctx) error {
	var newCustomers Customer

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Message: Bad Request")
	}

	for i, element := range customers {
		if element.Username == newCustomers.Username {
			element.Name = newCustomers.Name
			element.Phone = newCustomers.Phone
			customers = append(customers[:i], customers[i+1:]...)
			customers = append(customers, element)
		}
	}

	return c.Status(http.StatusOK).JSON("Updated Data")
}

func DeleteCustomer(c *fiber.Ctx) error {
	username := c.Params("username")

	for i, element := range customers {
		if element.Username == username {
			customers = append(customers[:i], customers[i+1:]...)
		}
	}
	return c.Status(http.StatusOK).JSON("Deleted Data")
}

func main() {
	app := fiber.New()
	app.Get("/customers", GetAllCustomer)
	app.Get("/customer/:username", GetByUsernameWithData)
	app.Post("/customer", CreateCustomer)
	app.Post("/customer/update", UpdateCustomer)
	app.Delete("/customer/:username", DeleteCustomer)

	app.Listen(":8000")

}
