package storage

import (
	"bin-cli/bins"
	"bin-cli/files"
	"encoding/json"
)

func SaveBinList(list *bins.BinList, path string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return files.WriteFile(path, data)
}

func LoadBinList(path string) (*bins.BinList, error) {
	data, err := files.ReadFile(path)
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
