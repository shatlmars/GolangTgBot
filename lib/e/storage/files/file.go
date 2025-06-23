package files

import (
	"main/lib/e"
	"main/lib/e/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	PathToFolder string
}

const (
	defaultPerm = 0774
)

func New(path string) Storage {
	return Storage{PathToFolder: path}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.Wrap("can't save", err) }()
	filePath := filepath.Join(s.PathToFolder, page.UserName) // затем можно будет заменить на sql бд
	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

}
