package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"context"
	"productAPI/infra"
	"time"
)

func StartHttp(ctx context.Context, containerDI *infra.ContainerDI) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	//PRODUCT
	app.Post("/create/product", containerDI.ProductHandler.Create)
	app.Put("/update/product", containerDI.ProductHandler.Update)
	app.Delete("/delete/product", containerDI.ProductHandler.Delete)
	app.Get("/list-all/product", containerDI.ProductHandler.Consult)

	//TRANSACTION
	app.Post("/transaction", containerDI.TransactionHandler.Create)
	app.Get("/list-all/transaction", containerDI.TransactionHandler.List)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
