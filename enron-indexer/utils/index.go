package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/Paoladevelopment/enron-indexer/models"
	zincutilities "github.com/Paoladevelopment/enron-indexer/zincUtilities"
)

// Represents the desired result after processing file
type Result struct {
	workerID int
	filePath string
	email    models.Email
	err      error
}

// A goroutine that processes files and send the results through a channel.
type Worker struct {
	id            int
	filePathQueue <-chan string
	resultChan    chan<- Result
}

func (w *Worker) Start() {
	go func() {
		for filePath := range w.filePathQueue {
			email, err := GenerateEmail(filePath)
			w.resultChan <- Result{workerID: w.id, filePath: filePath, email: email, err: err}
		}
	}()
}

// Worker pool manages the workers, distributes file processing and collects the emails.
type WorkerPool struct {
	filePathQueue chan string
	resultChan    chan Result
	workerCount   int
	emails        []models.Email
	mu            sync.Mutex
}

func NewWorkerPool(workerCount, bufferedSize int) *WorkerPool {
	return &WorkerPool{
		filePathQueue: make(chan string, bufferedSize),
		resultChan:    make(chan Result, bufferedSize),
		workerCount:   workerCount,
		emails:        make([]models.Email, 0),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{id: i, filePathQueue: wp.filePathQueue, resultChan: wp.resultChan}
		worker.Start()
	}

	go func() {
		for result := range wp.resultChan {
			if result.err != nil {
				fmt.Printf("Worker ID: %d, Path: %s, Error: %v\n", result.workerID, result.filePath, result.err)
			} else {
				wp.mu.Lock()
				wp.emails = append(wp.emails, result.email)
				if len(wp.emails) == 1000 {
					wp.sendBulk()
				}
				wp.mu.Unlock()
			}
		}
	}()
}

func (wp *WorkerPool) SubmitFile(filePath string) {
	wp.filePathQueue <- filePath
}

func (wp *WorkerPool) Close() {
	close(wp.filePathQueue)
	close(wp.resultChan)
	if len(wp.emails) > 0 {
		wp.sendBulk()
	}
}

func (wp *WorkerPool) sendBulk() {
	zincutilities.SaveBulk("emails", wp.emails)
	wp.emails = make([]models.Email, 0)
}

func IndexEmails(pathName string) {
	workerPool := NewWorkerPool(1000, 1000)
	workerPool.Start()

	root := filepath.Join(pathName, "maildir")
	err := filepath.Walk(root, func(path string, fs os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fs.IsDir() {
			workerPool.SubmitFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory tree: %v\n", err)
	}

	workerPool.Close()
}
