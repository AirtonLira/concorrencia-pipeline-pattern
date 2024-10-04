package images

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func LoadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, path := range paths {
			Job := Job{InputPath: path,
				OutPath: strings.Replace(path, "internal/images/", "internal/images/output/", 1)}
			Job.Image = readImage(path)
			out <- Job
		}
		close(out)
	}()
	return out
}

func readImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatalf("Erro ao realizar decode imagem: %s erro: %v", path, err)
		panic(err)
	}
	return img
}

func Resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			newWith := uint(500)
			newHeight := uint(500)
			resizedImg := resize.Resize(newWith, newHeight, job.Image, resize.Lanczos2)
			job.Image = resizedImg
			out <- job
		}
		close(out)
	}()
	return out
}

func GrayImage(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			bounds := job.Image.Bounds()
			grayImg := image.NewGray(bounds)

			//Convert each pixel to gray
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					originalPixel := job.Image.At(x, y)
					grayPixel := color.GrayModel.Convert(originalPixel)
					grayImg.Set(x, y, grayPixel)
				}
			}
			job.Image = grayImg
			out <- job
		}
		close(out)
	}()
	return out
}

func WriteImage(input <-chan Job) []string {
	var wg sync.WaitGroup
	resultWrite := []string{}
	mu := &sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for job := range input {
			wg.Add(1)
			go func(job Job) {
				defer wg.Done()
				outPutFile, err := os.Create(job.OutPath)
				if err != nil {
					log.Fatalf("Erro ao gravar saida da imagem: %s - error: %v", job.InputPath, err)
					panic(err)
				}
				defer outPutFile.Close()

				err = jpeg.Encode(outPutFile, job.Image, nil)
				if err != nil {
					log.Fatalf("Erro ao realizar encode da imagem: %s - error: %v", job.InputPath, err)
					panic(err)
				}

				mu.Lock()
				resultWrite = append(resultWrite, job.InputPath)
				mu.Unlock()
			}(job)
		}
	}()

	wg.Wait() // Espera todas as goroutines terminarem
	return resultWrite
}
