package helpers

import (
	"fmt"
	"time"
)

func CustomLog(logType string, message string) {
	iconLog := ""

	switch logType {
	case "error":
		iconLog = "🚨"
	case "info":
		iconLog = "💡"
	case "success":
		iconLog = "✅"
	default:
		iconLog = "🚧"
	}

	messageLog := fmt.Sprintf("[%s] - %s - [message]: %s", iconLog, time.Now().Format("2006-01-02 15:04:05"), message)
	fmt.Println(messageLog)
}
