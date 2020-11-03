package fs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ReadFromTestData(f string) []byte {
	filename := filepath.Join(GetTestDataPath(), f)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("error reading file:", filename)
	}
	return content
}

func TestDataJoin(f string) string {
	return filepath.Join(GetTestDataPath(), f)
}

func GetTestDataPath() string {
	current, _ := os.Getwd()
	parent := filepath.Dir(current)
	root := filepath.Dir(parent)
	return filepath.Join(root, "test", "testdata")
}
