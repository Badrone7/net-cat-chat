package NetCat

import (
	"log"
	"os"
	"strings"
)

func AppendToHistory(message string) {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	file, err := os.OpenFile(
		"history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		log.Println("error opening history file:", err)
		return
	}
	defer file.Close()
	cleanMessage := strings.TrimRight(message, "\r\n") + "\n"
	_, err = file.WriteString(cleanMessage)
	if err != nil {
		log.Println("error writing to history file:", err)
	}
}

func Historyclean() {
	err := os.Remove("history.txt")
	if err != nil && !os.IsNotExist(err) {
		log.Println("Warning: could not delete history.txt:", err)
	}
	_, err = os.Create("history.txt")
	if err != nil {
		log.Fatal("Failed to create history.txt:", err)
	}
}
