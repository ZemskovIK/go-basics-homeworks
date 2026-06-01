package interfaces

import "bin-cli/bins"

type FileStorage interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
}

type BinStorage interface {
	SaveBinList(list *bins.BinList, path string) error
	LoadBinList(path string) (*bins.BinList, error)
}
