package main

import (
	"context"
	"os/signal"
	"productAPI/cmd"
	"productAPI/infra"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	config := infra.NewConfig()
	containerDI := infra.NewContainerDI(config)
	defer containerDI.ShutDown()

	containerDI.Migration.MakeMigrations()
	cmd.StartHttp(ctx, containerDI)
}
