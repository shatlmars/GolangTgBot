package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"main/lib/e"
	"main/lib/e/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	PathToFolder string
}

const (
	defaultPerm = 0774
)

var ErrNoSavedPages = errors.New("no saved pages")

func New(path string) Storage {
	return Storage{PathToFolder: path}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.Wrap("can't save", err) }()
	fPath := filepath.Join(s.PathToFolder, page.UserName) // затем можно будет заменить на sql бд
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}
	fName, err := fName(page)
	if err != nil {
		return err
	}
	fPath = filepath.Join(fPath, fName)
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	if err = gob.NewEncoder(file).Encode(nil); err != nil {
		return err
	}
	return nil
}
func (s Storage) PickRandom(username string) (page *storage.Page, err error) {
	defer func() {
		err = e.Wrap("can't pick random page", err)
	}()

	path := filepath.Join(s.PathToFolder, username)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(page *storage.Page) error {
	fileName, err := fName(page)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.PathToFolder, page.UserName, fileName)
	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file %s", path)
		e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExsist(page *storage.Page) (bool, error) {
	// defer func(){}()

	fileName, err := fName(page)
	if err != nil {
		return false, e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.PathToFolder, page.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, ErrNoSavedPages):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exsist", path)
		return false, e.Wrap(msg, err)
	}

	return true, nil
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode file", err)
	}

	defer func() { _ = f.Close() }()
	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode file", err)
	}

	return &p, nil
}

func fName(p *storage.Page) (string, error) {
	return p.Hash()
}
