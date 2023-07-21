package main

import (
	"log"
	"os"
	"time"
)

func main() {
	var hardCodedString = "/var/log/myapp"
	logFile, err := os.OpenFile(hardCodedString+"/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	for {
		logMessage := "Logging at " + time.Now().String() + "generated by SURAJ KUMAR"
		log.Println(logMessage)

		time.Sleep(30 * time.Second)
	}
}
