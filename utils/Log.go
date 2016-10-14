// Log.go
package utils

import (
	"log"
	"os"
)

func LogFatal(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Llongfile|log.Ldate|log.Ltime)
	logger.SetPrefix("[Fatal]")
	logger.Println(v...)
	defer logfile.Close()
}

func LogErr(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Llongfile|log.Ldate|log.Ltime)
	logger.SetPrefix("[Error]")
	logger.Println(v...)
	defer logfile.Close()
}

func Log(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Info]")
	logger.Println(v...)
	defer logfile.Close()
}

func LogDebug(v ...interface{}) {
	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Debug]")
	logger.Println(v...)
	defer logfile.Close()
}
