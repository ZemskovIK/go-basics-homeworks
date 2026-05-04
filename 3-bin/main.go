package main

import (
	"bin-cli/files"
	"bin-cli/interfaces"
	"bin-cli/storage"
)

type App struct {
	binStorage interfaces.BinStorage
}

func NewApp(bs interfaces.BinStorage) *App {
	return &App{
		binStorage: bs,
	}
}

func main() {
	fileStorage := files.NewFileStorage()
	binStorage := storage.NewBinStorage(fileStorage)

	app := NewApp(binStorage)

	_ = app
}
