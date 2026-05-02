package fileutil

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

func ReadYaml[T any](filename string) (*T, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	

	var v T
	if err := yaml.NewDecoder(bytes.NewReader(f)).Decode(&v); err != nil {
		return nil, fmt.Errorf("config load: %w", err)
	}
	return &v, nil
}
func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	record, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return record, nil
}
