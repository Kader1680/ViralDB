
package storage
import (
    "os"
    "tinydb/internal/record"
)

type Storage struct {
	file *os.File
}

func NewStorage(path string) (*Storage, error) {
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	if err != nil {
		return nil, err
	}

	return &Storage{
		file: f,
	}, nil
}



func (s *Storage) Append(r record.Record) error {
	data, err := record.Encode(r)
	if err != nil {
		return err
	}

	_, err = s.file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

