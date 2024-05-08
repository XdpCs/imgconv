package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
)

// @Title        image.go
// @Description
// @Create       XdpCs 2024-05-08 17:05
// @Update       XdpCs 2024-05-08 17:05

func init() {
	imageFormat = map[string]bool{
		"png":  true,
		"jpeg": true,
		"jpg":  true,
		"gif":  true,
		"webp": true,
	}
}

var imageFormat map[string]bool

func CheckImageFormat(format string) error {
	if _, ok := imageFormat[format]; !ok {
		return fmt.Errorf("unsupported image format")
	}
	return nil
}

type InputImage struct {
	Reader io.Reader
	Format string
	Name   string
}

type OutputImage struct {
	Writer io.Writer
	Format string
	Name   string
}

func ConvertFormat(input *InputImage, output *OutputImage) error {
	if input == nil || output == nil {
		return fmt.Errorf("input or output is nil")
	}

	if input.Format == output.Format {
		return fmt.Errorf("input format and output format are the same")
	}

	decodeImageFunc, err := DecodeFactory(input.Format)
	if err != nil {
		return fmt.Errorf("get decodeImageFunc error: %v", err)
	}

	if input.Reader == nil {
		return fmt.Errorf("input reader is nil")
	}

	decodeImage, err := decodeImageFunc(input.Reader)
	if err != nil {
		return fmt.Errorf("image decode error: %v", err)
	}

	encodeImageFunc, err := EncodeFactory(output.Format)
	if err != nil {
		return fmt.Errorf("image format encode error: %v", err)
	}

	if output.Writer == nil {
		return fmt.Errorf("output writer is nil")
	}

	if err := encodeImageFunc(output.Writer, decodeImage); err != nil {
		return fmt.Errorf("image encode error: %v", err)
	}

	return nil
}

func DecodeFactory(format string) (func(io.Reader) (image.Image, error), error) {
	switch format {
	case "png":
		return func(r io.Reader) (image.Image, error) {
			decode, err := png.Decode(r)
			if err != nil {
				return nil, err
			}
			return decode, nil
		}, nil

	case "jpg", "jpeg":
		return func(r io.Reader) (image.Image, error) {
			decode, err := jpeg.Decode(r)
			if err != nil {
				return nil, err
			}
			return decode, nil
		}, nil
	case "gif":
		return func(r io.Reader) (image.Image, error) {
			decode, err := gif.Decode(r)
			if err != nil {
				return nil, err
			}
			return decode, nil
		}, nil
	case "webp":
		return func(r io.Reader) (image.Image, error) {
			decode, err := webp.Decode(r)
			if err != nil {
				return nil, err
			}
			return decode, nil
		}, nil
	}
	return nil, nil
}

func EncodeFactory(format string) (func(io.Writer, image.Image) error, error) {
	switch format {
	case "png":
		return func(w io.Writer, img image.Image) error {
			err := png.Encode(w, img)
			if err != nil {
				return err
			}
			return nil
		}, nil
	case "jpg", "jpeg":
		return func(w io.Writer, img image.Image) error {
			err := jpeg.Encode(w, img, nil)
			if err != nil {
				return err
			}
			return nil
		}, nil
	case "gif":
		return func(w io.Writer, img image.Image) error {
			err := gif.Encode(w, img, nil)
			if err != nil {
				return err
			}
			return nil
		}, nil
	case "webp":
		return func(w io.Writer, img image.Image) error {
			err := webp.Encode(w, img, nil)
			if err != nil {
				return err
			}
			return nil
		}, nil
	}
	return nil, nil
}
