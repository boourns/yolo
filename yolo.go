package yolo

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
)

func Load(filename string, result interface{}) error {
	f, err := os.Open(filename)

	if err != nil {
		return err
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(result)

	return err
}

func Save(filename string, data interface{}) error {
	tmpFile := randomName()

	f, err := os.Create(tmpFile)
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(f)
	enc.Encode(data)
	f.Sync()
	f.Close()

	return os.Rename(tmpFile, filename)
}

func randomName() string {
	r := make([]byte, 16)

	_, err := rand.Read(r)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", os.TempDir(), fmt.Sprintf("%x", r))
}
