package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func listFiles(path string) ([]string, error) {
	var fileList []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			fileList = append(fileList, entry.Name())
		}
	}
	return fileList, nil
}
func printFiles(fileList []string) {
	fmt.Println("Znalezione pliki:")
	for _, file := range fileList {
		fmt.Println("-", file)
	}
}
func inputPath() string {
	var path string
	fmt.Println("Podaj sciezke do posortowania: ")
	fmt.Scan(&path)
	return path
}
func makeDecision() bool {
	var ans string
	fmt.Println("Czy te pliki chcesz posortować? (Y/N)")
	fmt.Scan(&ans)
	if ans == "y" || ans == "Y" {
		return true
	}
	return false
}
func sortFiles(fileList []string, basePath string) {
	folders := []string{"docs", "images", "compressed", "sql dumps", "videos", "sounds", "executables", "presentations"}
	fmt.Println("base path: " + basePath)
	for _, folder := range folders {
		targetFolder := filepath.Join(basePath, folder)
		err := os.MkdirAll(targetFolder, 0755)
		if err != nil {
			log.Fatal("Nie udało się przygotować folderów:", err)
		}
	}
	for _, file := range fileList {
		ext := filepath.Ext(file)
		oldPath := filepath.Join(basePath, file)
		var newPath string

		if ext == ".pdf" || ext == ".docx" {
			newPath = filepath.Join(basePath, "docs", file)
		} else if ext == ".png" || ext == ".jpg" || ext == ".JPG" || ext == ".PNG" {
			newPath = filepath.Join(basePath, "images", file)
		} else if ext == ".zip" || ext == ".rar" {
			newPath = filepath.Join(basePath, "compressed", file)
		} else if ext == ".sql" {
			newPath = filepath.Join(basePath, "sql dumps", file)
		} else if ext == ".mp4" {
			newPath = filepath.Join(basePath, "videos", file)
		} else if ext == ".mp3" {
			newPath = filepath.Join(basePath, "sounds", file)
		} else if ext == ".exe" && strings.TrimSuffix(file, filepath.Ext(file)) != "sorter" {
			newPath = filepath.Join(basePath, "executables", file)
		} else if ext == ".pptx" {
			newPath = filepath.Join(basePath, "presentations", file)
		} else {
			continue
		}
		e := os.Rename(oldPath, newPath)
		if e != nil {
			log.Fatal(e)
		}

	}
	fmt.Println("Posortowano pliki!")
}

func main() {
	var decision bool = false
	var files []string
	var err error
	var path string
	for decision == false {
		path = inputPath()
		files, err = listFiles(path)
		if err != nil {
			log.Fatal(err)
		}
		printFiles(files)
		decision = makeDecision()
	}

	sortFiles(files, path)
}
