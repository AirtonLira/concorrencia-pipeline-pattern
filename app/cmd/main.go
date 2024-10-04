package main

import (
	"image"
	"log"
	"os"
	"strings"

	"github.com/AirtonLira/concorrencia-pipeline-pattern/internal/images"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func listImages() []string {
	dir := "internal/images"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Iniciando coleta de imagens...")
	var jpegFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") {
			jpegFiles = append(jpegFiles, dir+"/"+file.Name())
			log.Printf("Imagem obtida com sucesso: %s", file.Name())
		}
	}
	return jpegFiles
}

func main() {
	imagePaths := listImages()
	canal1 := images.LoadImage(imagePaths)
	canal2 := images.Resize(canal1)
	canal3 := images.GrayImage(canal2)
	results := images.WriteImage(canal3)
	for _, image := range results {
		log.Printf("Imagem - %s convertida e gravada com sucesso", image)
	}
}
