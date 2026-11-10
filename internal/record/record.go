
package record

import (
	"encoding/binary"
	"errors"
	"time"
)

type Record struct {
	Key      string
	Value    string
	Deleted  bool
	Createat time.Time
}

func Encode(r Record) ([]byte, error) {
	keyBytes := []byte(r.Key)
	valueBytes := []byte(r.Value)

	header := make([]byte, 9)

	if r.Deleted {
		header[0] = 1
	}

	binary.LittleEndian.PutUint32(header[1:5], uint32(len(keyBytes)))
	binary.LittleEndian.PutUint32(header[5:9], uint32(len(valueBytes)))

	data := append(header, keyBytes...)
	data = append(data, valueBytes...)

	return data, nil
}
 
func Decode(data []byte) (Record, error) {
	var r Record

 
	if len(data) < 9 {
		return r, errors.New("invalid record: header too short")
	}

 
	if data[0] > 1 {
		return r, errors.New("invalid record: invalid deleted flag")
	}
	r.Deleted = data[0] == 1

 	// Read Key and Value lengths.
	keyLength := binary.LittleEndian.Uint32(data[1:5])
	valueLength := binary.LittleEndian.Uint32(data[5:9])

	expectedSize := uint64(9) + uint64(keyLength) + uint64(valueLength)

	if expectedSize != uint64(len(data)) {
		return r, errors.New("invalid record: size mismatch")
	}

	keyStart := 9
	keyEnd := keyStart + int(keyLength)
	valueEnd := keyEnd + int(valueLength)

	r.Key = string(data[keyStart:keyEnd])
	r.Value = string(data[keyEnd:valueEnd])

	return r, nil
}
