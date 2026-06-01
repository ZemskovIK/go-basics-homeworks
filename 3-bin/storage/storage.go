package storage

import (
	"bin-cli/bins"
	"bin-cli/interfaces"
	"encoding/json"
)

type BinStorage struct {
	fileStorage interfaces.FileStorage
}

func NewBinStorage(fs interfaces.FileStorage) *BinStorage {
	return &BinStorage{
		fileStorage: fs,
	}
}

func (s *BinStorage) SaveBinList(list *bins.BinList, path string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return s.fileStorage.WriteFile(path, data)
}

func (s *BinStorage) LoadBinList(path string) (*bins.BinList, error) {
	data, err := s.fileStorage.ReadFile(path)
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
