package main

import (
	"os"
	"sync"
)

type WALogger struct {
	file  *os.File
	mutex *sync.Mutex
}
