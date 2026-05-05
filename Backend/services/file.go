package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func SaveImage(r *http.Request, field string) (string, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return "", err
	}

	file, handler, err := r.FormFile(field)
	if err != nil {
		return "/uploads/default.jpg", nil
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("el archivo debe ser una imagen válida")
	}

	if handler.Size > 1<<20 {
		return "", errors.New("imagen muy grande (max 1MB)")
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := "./uploads/" + filename

	dst, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("error creando archivo: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", fmt.Errorf("error guardando archivo: %w", err)
	}

	return "/uploads/" + filename, nil
}
