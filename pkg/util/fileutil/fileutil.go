package fileutil

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"go.yaml.in/yaml/v4"
)

func ReadYAML[T any](fileName string) (T, error) {
	var v T

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

func ReadJSON[T any](r io.ReadCloser) (*T, error) {
	var v T

	err := json.NewDecoder(r).Decode(&v)
	return &v, errors.Join(err, r.Close())
}

func WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
