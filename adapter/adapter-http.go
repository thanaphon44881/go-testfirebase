package adapter

import (
	"TestUser/repository"
	"TestUser/service"

	"github.com/gofiber/fiber/v2"
)

type httpUser struct {
	repo service.ServiceUser
}

func Newhttpuser(s service.ServiceUser) httpUser {
	return httpUser{repo: s}
}

func (h httpUser) CreatUser(c *fiber.Ctx) error {
	var user repository.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid reauest"})
	}
	err := h.repo.Creat(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *httpUser) GetUsers(c *fiber.Ctx) error {
	users, err := h.repo.GetUsers()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(users)
}

func (h *httpUser) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.repo.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}
