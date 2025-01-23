package utils

import (
	"cursor-api/internal/domain/constants"
	"fmt"
	"os"
	"time"
)

func DebugLog(format string, args ...interface{}) {
	if !constants.DEBUG {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	logMessage := fmt.Sprintf("%s - %s\n", timestamp, message)

	// 异步写入日志
	go func() {
		f, err := os.OpenFile(constants.DEBUG_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening debug log file: %v\n", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(logMessage); err != nil {
			fmt.Printf("Error writing to debug log file: %v\n", err)
		}
	}()
} 