package imgconv

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// @Title        http.go
// @Description
// @Create       XdpCs 2024-05-08 17:48
// @Update       XdpCs 2024-05-08 17:48

func HttpInputImageFile(url string) (*InputImage, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get error: %v", err)
	}
	defer resp.Body.Close()
	// 获取content-type
	contentType := resp.Header.Get("Content-Type")
	format, err := ContentTypeFormat(contentType)
	if err != nil {
		return nil, fmt.Errorf("content type format error: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	return &InputImage{
		Reader: bytes.NewReader(bodyBytes),
		Format: format,
	}, nil
}

func ContentTypeFormat(contentType string) (string, error) {
	switch strings.ToLower(contentType) {
	case "image/png":
		return "png", nil
	case "image/jpeg":
		return "jpeg", nil
	case "image/jpg":
		return "jpeg", nil
	case "image/gif":
		return "gif", nil
	case "image/webp":
		return "webp", nil
	}
	return "", fmt.Errorf("unsupported image format")
}
