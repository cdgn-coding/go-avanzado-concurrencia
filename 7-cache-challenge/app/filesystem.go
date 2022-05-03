package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type cacheFile struct {
	filename string
	content  string
}

type filesystemWorker struct {
	writer   chan cacheFile
	basepath string
}

func newFilesystemWorker(basePath string) *filesystemWorker {
	return &filesystemWorker{
		writer:   make(chan cacheFile),
		basepath: basePath,
	}
}

func (worker *filesystemWorker) persist(key, val string) {
	worker.writer <- cacheFile{filename: key, content: val}
}

func (worker *filesystemWorker) writeFile(path, content string) {
	createdFile, _ := os.Create(path)
	defer createdFile.Close()
	_, errorWriting := createdFile.WriteString(content)
	if errorWriting != nil {
		log.Fatal("Error writing file")
	}
}

func (worker *filesystemWorker) load(recordsToSave chan record) {
	defer close(recordsToSave)
	files, errorReadingDir := ioutil.ReadDir(worker.basepath)

	if errorReadingDir != nil {
		log.Fatal("Cannot scan the base path")
	}

	for _, f := range files {
		key := f.Name()
		path := filepath.Join(worker.basepath, key)
		content := worker.readContent(path)
		recordsToSave <- record{key, content}
	}
}

func (worker *filesystemWorker) readContent(path string) string {
	openedFile, errorOpening := os.Open(path)
	if errorOpening != nil {
		log.Fatal("Cannot open file")
	}
	defer openedFile.Close()

	buf := new(bytes.Buffer)
	_, errorReading := buf.ReadFrom(openedFile)
	if errorReading != nil {
		log.Fatal("Error reading file")
	}
	contents := buf.String()

	return contents
}

func (worker *filesystemWorker) start() {
	for {
		persistent := <-worker.writer
		filename, content := persistent.filename, persistent.content
		path := filepath.Join(worker.basepath, filename)
		worker.writeFile(path, content)
	}
}
