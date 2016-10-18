// Log.go
package utils

import (
	"log"
	"os"
)

var strPathName string

func init() {
	var strFileName = "/logs/server.log"
	strPathName = getFileName(strFileName)

}

func getFileName(filename string) string {

	var strFilePath string
	strFilePath, _ = os.Getwd()
	return strFilePath + filename
}

func LogFatal(v ...interface{}) {

	log.Println(strPathName)
	logfile, err := os.Create(strPathName)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.Llongfile|log.Ldate|log.Ltime)
	logger.SetPrefix("[Fatal]")
	logger.Println(v...)
	defer logfile.Close()
}

func LogErr(v ...interface{}) {

	log.Println(strPathName)
	logfile, err := os.Create(strPathName)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.Llongfile|log.Ldate|log.Ltime)
	logger.SetPrefix("[Error]")
	logger.Println(v...)
	defer logfile.Close()
}

func Log(v ...interface{}) {

	log.Println(strPathName)
	logfile, err := os.Create(strPathName)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Info]")
	logger.Println(v...)
	defer logfile.Close()
}

func LogDebug(v ...interface{}) {
	log.Println(strPathName)
	logfile, err := os.Create(strPathName)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Debug]")
	logger.Println(v...)
	defer logfile.Close()
}

func LogNotice(v ...interface{}) {

	log.Println(strPathName)
	logfile, err := os.Create(strPathName)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[NOTICE]")
	logger.Println(v...)
	defer logfile.Close()
}
