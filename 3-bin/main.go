package main

import (
	"bin-cli/api"
	"bin-cli/bins"
	"bin-cli/config"
	"bin-cli/files"
	"bin-cli/interfaces"
	"bin-cli/storage"
	"flag"
	"fmt"
	"os"
)

const binsFile = "bins.json"

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

func (a *App) Create(filePath, name string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	id, err := a.apiClient.Create(data, name)
	if err != nil {
		fmt.Println("Ошибка создания bin:", err)
		return
	}

	list := a.loadOrNewList()
	list.Bins = append(list.Bins, *bins.NewBin(id, false, name))

	err = a.binStorage.SaveBinList(list, binsFile)
	if err != nil {
		fmt.Println("Ошибка сохранения списка:", err)
		return
	}

	fmt.Printf("Bin создан: id=%s  name=%s\n", id, name)
}

func (a *App) Get(id string) {
	result, err := a.apiClient.Get(id)
	if err != nil {
		fmt.Println("Ошибка получения bin:", err)
		return
	}
	fmt.Println(result)
}

func (a *App) Update(id, filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	err = a.apiClient.Update(id, data)
	if err != nil {
		fmt.Println("Ошибка обновления bin:", err)
		return
	}

	fmt.Printf("Bin %s обновлён\n", id)
}

func (a *App) Delete(id string) {
	err := a.apiClient.Delete(id)
	if err != nil {
		fmt.Println("Ошибка удаления bin из API:", err)
		return
	}

	list, err := a.binStorage.LoadBinList(binsFile)
	if err == nil {
		filtered := []bins.Bin{}
		for _, b := range list.Bins {
			if b.Id != id {
				filtered = append(filtered, b)
			}
		}
		list.Bins = filtered
		a.binStorage.SaveBinList(list, binsFile)
	}

	fmt.Printf("Bin %s удалён\n", id)
}

func (a *App) List() {
	list, err := a.binStorage.LoadBinList(binsFile)
	if err != nil {
		fmt.Println("Список пуст")
		return
	}

	if len(list.Bins) == 0 {
		fmt.Println("Список пуст")
		return
	}

	for _, b := range list.Bins {
		fmt.Printf("id=%-30s  name=%s\n", b.Id, b.Name)
	}
}

func (a *App) loadOrNewList() *bins.BinList {
	list, err := a.binStorage.LoadBinList(binsFile)
	if err != nil {
		return bins.NewBinList()
	}
	return list
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <команда> [флаги]")
		fmt.Println("Команды: create, update, delete, get, list")
		os.Exit(1)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileStorage := files.NewFileStorage()
	binStorage := storage.NewBinStorage(fileStorage)
	apiClient := api.NewClient(cfg.Key)
	app := NewApp(binStorage, apiClient)

	command := os.Args[1]
	flags := flag.NewFlagSet(command, flag.ExitOnError)

	filePath := flags.String("file", "", "Путь к JSON файлу")
	name := flags.String("name", "", "Имя bin")
	id := flags.String("id", "", "ID bin")

	flags.Parse(os.Args[2:])

	switch command {
	case "create":
		if *filePath == "" || *name == "" {
			fmt.Println("Нужно указать --file и --name")
			os.Exit(1)
		}
		app.Create(*filePath, *name)
	case "update":
		if *filePath == "" || *id == "" {
			fmt.Println("Нужно указать --file и --id")
			os.Exit(1)
		}
		app.Update(*id, *filePath)
	case "delete":
		if *id == "" {
			fmt.Println("Нужно указать --id")
			os.Exit(1)
		}
		app.Delete(*id)
	case "get":
		if *id == "" {
			fmt.Println("Нужно указать --id")
			os.Exit(1)
		}
		app.Get(*id)
	case "list":
		app.List()
	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		os.Exit(1)
	}
}
