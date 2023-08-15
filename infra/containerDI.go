package infra

import (
	"gorm.io/gorm"
	"productAPI/infra/database"
	"productAPI/infra/database/postgresql"
	"productAPI/migration"
	"productAPI/user_case/handler"
	"productAPI/user_case/repository"
	"productAPI/user_case/service"
)

type ContainerDI struct {
	Config                Config
	PostgresqlDB          *gorm.DB
	Migration             migration.DatabaseMakeMigrations
	ProductHandler        handler.ProductHandler
	TransactionHandler    handler.TransactionHandler
	ProductService        service.ProductService
	TransactionService    service.TransactionService
	ProductRepository     repository.ProductRepository
	TransactionRepository repository.TransactionRepository
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}

	container.buildDB()
	container.buildRepository()
	container.buildService()
	container.buildHandler()

	return container
}

func (c *ContainerDI) buildDB() {
	gormConfig := database.Config{
		Hostname:    c.Config.PostgresqlHost,
		Port:        c.Config.PostgresqlPort,
		UserName:    c.Config.PostgresqlUser,
		Password:    c.Config.PostgresqlPassword,
		Database:    c.Config.PostgresqlDatabase,
		Environment: c.Config.Environment,
	}
	c.PostgresqlDB = postgresql.InitGorm(&gormConfig)
	//MIGRATION
	c.Migration = migration.NewDatabaseMakeMigrations(c.PostgresqlDB)
}

func (c *ContainerDI) buildRepository() {
	c.ProductRepository = repository.NewProductRepository(c.PostgresqlDB)
	c.TransactionRepository = repository.NewTransactionRepository(c.PostgresqlDB)
}

func (c *ContainerDI) buildService() {
	c.ProductService = service.NewProductService(c.ProductRepository)
	c.TransactionService = service.NewTransactionService(c.TransactionRepository, c.ProductRepository)
}

func (c *ContainerDI) buildHandler() {
	c.ProductHandler = handler.NewProductHandler(c.ProductService)
	c.TransactionHandler = handler.NewTransactionHandler(c.TransactionService)
}

func (c *ContainerDI) ShutDown() {}
