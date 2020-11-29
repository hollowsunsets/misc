package main

import (
	"fmt"
	"sort"
	"strings"
)

type Directory struct {
	subdirectories map[string]Directory
	files          map[string]string
}

type FileSystem struct {
	root Directory
}

func Constructor() FileSystem {
	fs := new(FileSystem)
	fs.root = Directory{
		subdirectories: make(map[string]Directory),
		files:          make(map[string]string),
	}
	return *fs
}

func (fs *FileSystem) Ls(path string) []string {
	dir := fs.root
	path = strings.TrimRight(path, "\r\n") // For platform independence
	var files []string
	if path == "/" {
		directoryNames := strings.Split("/", path)
		for i := 1; i < len(directoryNames)-1; i++ {
			name := directoryNames[i]
			dir = dir.subdirectories[name]
		}
		lastDirName := directoryNames[len(directoryNames)-1]
		if _, exists := dir.files[lastDirName]; exists {
			files = append(files, lastDirName)
			return files
		}
		dir = dir.subdirectories[lastDirName]
	}

	for key := range dir.subdirectories {
		files = append(files, key)
	}
	for key := range dir.files {
		files = append(files, key)
	}
	sort.Strings(files)
	return files
}

func (fs *FileSystem) Mkdir(path string) {
	dir := fs.root
	directoryNames := strings.Split("/", path)
	for i := 1; i < len(directoryNames); i++ {
		name := directoryNames[i]
		if dir, exists := dir.subdirectories[name]; !exists {
			dir.subdirectories[name] = Directory{}
		}
		dir = dir.subdirectories[name]
	}
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	dir := fs.root
	fileNames := strings.Split("/", filePath)
	for i := 1; i < len(fileNames)-1; i++ {
		name := fileNames[i]
		dir = dir.subdirectories[name]
	}
	lastFileName := fileNames[len(fileNames)-1]

	var fileContent strings.Builder
	if previousFileContent, exists := dir.files[lastFileName]; exists {
		fileContent.WriteString(previousFileContent)
	}
	fileContent.WriteString(content)

	dir.files[lastFileName] = fileContent.String()
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	dir := fs.root
	fileNames := strings.Split("/", filePath)
	for i := 1; i < len(fileNames)-1; i++ {
		name := fileNames[i]
		dir = dir.subdirectories[name]
	}
	lastFileName := fileNames[len(fileNames)-1]
	return dir.files[lastFileName]
}

func main() {
	fmt.Println("Hello")
	fmt.Println(FileSystem{})
}
