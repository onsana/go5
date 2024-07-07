package manager

import (
	"encoding/json"
	"io"
	"os"
)

type Storage struct {
	filename  string
	passwords map[string]string
}

func NewStorage(filename string) *Storage {
	s := &Storage{
		filename:  filename,
		passwords: make(map[string]string),
	}
	s.load()
	return s
}

func (s *Storage) load() {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			s.passwords = make(map[string]string)
			return
		}
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &s.passwords)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) save() {
	bytes, err := json.MarshalIndent(s.passwords, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(s.filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) ListNames() []string {
	names := make([]string, 0, len(s.passwords))
	for name := range s.passwords {
		names = append(names, name)
	}
	return names
}

func (s *Storage) SavePassword(name, password string) {
	s.passwords[name] = password
	s.save()
}

func (s *Storage) GetPassword(name string) (string, bool) {
	password, found := s.passwords[name]
	return password, found
}
