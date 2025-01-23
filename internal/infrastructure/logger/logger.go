package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	file    *os.File
	mu      sync.Mutex
	logPath string
}

func NewLogger(logPath string) (*Logger, error) {
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Logger{
		file:    file,
		logPath: logPath,
	}, nil
}

func (l *Logger) Log(level, message string, fields map[string]interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
	
	if len(fields) > 0 {
		logEntry += fmt.Sprintf(" | fields: %v", fields)
	}
	logEntry += "\n"

	_, err := l.file.WriteString(logEntry)
	return err
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l *Logger) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.file.Close(); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102150405")
	backupPath := fmt.Sprintf("%s.%s", l.logPath, timestamp)
	
	if err := os.Rename(l.logPath, backupPath); err != nil {
		return err
	}

	file, err := os.OpenFile(l.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	l.file = file
	return nil
} 