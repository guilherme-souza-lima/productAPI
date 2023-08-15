package handler

import (
	"github.com/gofiber/fiber/v2"
	"productAPI/user_case/request"
	"productAPI/user_case/service"
)

type ProductHandler struct {
	InterfaceProductService service.InterfaceProductService
}

func NewProductHandler(InterfaceProductService service.InterfaceProductService) ProductHandler {
	return ProductHandler{InterfaceProductService}
}

func (p ProductHandler) Create(c *fiber.Ctx) error {
	var product request.CreateProduct
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err := p.InterfaceProductService.Create(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("success")
}

func (p ProductHandler) Update(c *fiber.Ctx) error {
	var product request.UpdateProduct
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err := p.InterfaceProductService.Update(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("success")
}

func (p ProductHandler) Consult(c *fiber.Ctx) error {
	result, err := p.InterfaceProductService.ConsultAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (p ProductHandler) Delete(c *fiber.Ctx) error {
	var product request.DeleteProduct
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err := p.InterfaceProductService.Delete(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("success")
}
