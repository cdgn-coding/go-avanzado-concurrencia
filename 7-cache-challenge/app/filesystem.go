package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	fsWrite  = "fsWrite"
	fsDelete = "fsDelete"
)

type writerAction struct {
	record *record
	action string
}

type filesystemWorker struct {
	fsActionQueue chan writerAction
	basepath      string
}

func newFilesystemWorker(basePath string) *filesystemWorker {
	return &filesystemWorker{
		fsActionQueue: make(chan writerAction),
		basepath:      basePath,
	}
}

func (worker *filesystemWorker) enqueuePersist(key, val string) {
	action := writerAction{
		record: &record{
			key:   key,
			value: val,
		},
		action: fsWrite,
	}
	worker.fsActionQueue <- action
}

func (worker *filesystemWorker) enqueueDelete(key string) {
	action := writerAction{
		record: &record{
			key: key,
		},
		action: fsDelete,
	}
	worker.fsActionQueue <- action
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

func (worker *filesystemWorker) deleteFile(path string) {
	errorDeleting := os.Remove(path)
	if errorDeleting != nil {
		log.Fatal("Error deleting file")
	}
}

func (worker *filesystemWorker) start() {
	for {
		enqueuedAction := <-worker.fsActionQueue
		filename, content, action := enqueuedAction.record.key, enqueuedAction.record.value, enqueuedAction.action
		path := filepath.Join(worker.basepath, filename)
		switch action {
		case fsWrite:
			worker.writeFile(path, content)
		case fsDelete:
			worker.deleteFile(path)
		}
	}
}
