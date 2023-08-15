package handler

import (
	"github.com/gofiber/fiber/v2"
	"productAPI/user_case/request"
	"productAPI/user_case/service"
)

type TransactionHandler struct {
	InterfaceTransactionService service.InterfaceTransactionService
}

func NewTransactionHandler(InterfaceTransactionService service.InterfaceTransactionService) TransactionHandler {
	return TransactionHandler{InterfaceTransactionService}
}

func (t TransactionHandler) Create(c *fiber.Ctx) error {
	var transaction request.CreateTransaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err := t.InterfaceTransactionService.Create(transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("success")
}

func (t TransactionHandler) List(c *fiber.Ctx) error {
	result, err := t.InterfaceTransactionService.List()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
