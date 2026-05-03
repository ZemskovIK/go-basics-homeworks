package storage

import (
	"bin-cli/bins"
	"bin-cli/file"
	"encoding/json"
)

func SaveBinList(list *bins.BinList, path string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return file.WriteFile(path, data)
}

func LoadBinList(path string) (*bins.BinList, error) {
	data, err := file.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var list bins.BinList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
