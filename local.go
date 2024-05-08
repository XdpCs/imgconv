package imgconv

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// @Title        local.go
// @Description
// @Create       XdpCs 2024-05-08 17:53
// @Update       XdpCs 2024-05-08 17:53

func LocalInputImageFile(path string) (*InputImage, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	format, err := FileFormat(path)
	if err != nil {
		return nil, fmt.Errorf("file format error: %v", err)
	}

	bytesFile, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file error: %v", err)
	}

	return &InputImage{
		Reader: bytes.NewReader(bytesFile),
		Format: format,
	}, nil
}

func LocalOutputImageFile(name, dir, format string) (*OutputImage, error) {
	file, err := os.Create(dir + name + "." + format)
	if err != nil {
		return nil, fmt.Errorf("create file error: %v", err)
	}
	return &OutputImage{
		Writer: file,
		Format: format,
		Name:   name,
	}, nil
}

func FileFormat(path string) (string, error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".png":
		return "png", nil
	case ".jpeg":
		return "jpeg", nil
	case ".jpg":
		return "jpg", nil
	case ".gif":
		return "gif", nil
	case ".webp":
		return "webp", nil
	}
	return "", fmt.Errorf("unsupported image format")
}
