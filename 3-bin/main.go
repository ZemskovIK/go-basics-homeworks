package main

import (
	"bin-cli/api"
	"bin-cli/config"
	"bin-cli/files"
	"bin-cli/interfaces"
	"bin-cli/storage"
	"fmt"
	"os"
)

type App struct {
	binStorage interfaces.BinStorage
	apiClient  *api.Client
}

func NewApp(bs interfaces.BinStorage, ac *api.Client) *App {
	return &App{
		binStorage: bs,
		apiClient:  ac,
	}
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileStorage := files.NewFileStorage()
	binStorage := storage.NewBinStorage(fileStorage)

	apiClient := api.NewClient(cfg.Key)

	app := NewApp(binStorage, apiClient)
	_ = app
}
