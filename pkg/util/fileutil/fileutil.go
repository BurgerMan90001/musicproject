package fileutil

import (
	"context"
	"database/sql"
	"os"

	"go.yaml.in/yaml/v4"
)

func ReadYAML[T any](fileName string) (T, error) {
	var v T
	//log.Println(os.Getwd())
	f, err := os.Open(fileName)

	if err != nil {
		return v, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

func ExecSql(ctx context.Context, db *sql.DB, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, string(data)); err != nil {
		return err
	}
	return nil
}
