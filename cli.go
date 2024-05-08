//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"imgconv"
)

func main() {
	var url string
	var inputImagePath string
	var outputImageName string
	var outputDir string
	var convertFormat string

	flag.StringVar(&url, "url", "", "image url")
	flag.StringVar(&inputImagePath, "IPath", "", "input image path")
	flag.StringVar(&outputImageName, "OName", "output", "output image name, if you don't want to use the default name(output), you can specify it")
	flag.StringVar(&outputDir, "ODir", "./", "output dir, if you don't want to use the default dir, you can specify it")
	flag.StringVar(&convertFormat, "format", "", "output image format, support png, jpeg, jpg, gif, webp")
	flag.Usage = func() {
		fmt.Println("Usage: imgconv -url <url> -format <format>")
		flag.PrintDefaults()
	}
	flag.Parse()

	if url == "" && inputImagePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if imgconv.CheckImageFormat(convertFormat) != nil {
		fmt.Println("unsupported image format")
		os.Exit(1)
	}

	if url != "" {
		inputImage, err := imgconv.HttpInputImageFile(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputImage, err := imgconv.LocalOutputImageFile(outputImageName, outputDir, convertFormat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = imgconv.ConvertFormat(inputImage, outputImage)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		inputImage, err := imgconv.LocalInputImageFile(inputImagePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputImage, err := imgconv.LocalOutputImageFile(outputImageName, outputDir, convertFormat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = imgconv.ConvertFormat(inputImage, outputImage)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
