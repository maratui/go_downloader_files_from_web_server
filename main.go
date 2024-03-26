package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	downloadAllFiles()
}

func downloadAllFiles() {
	successfulFile := createFile(successfulFileName)
	defer successfulFile.Close()
	notSuccessfulFile := createFile(notSuccessfulFileName)
	defer notSuccessfulFile.Close()
	logFile = createFile(logFileName)
	defer logFile.Close()

	for _, fileUrl := range fileUrls {
		downloadFile(fileUrl, successfulFile, notSuccessfulFile)
	}
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	checkError(err)

	return file
}

func downloadFile(fileUrl string, successfulFile *os.File, notSuccessfulFile *os.File) {
	fileName := getFileName(fileUrl)
	file := createFile(binPath + fileName)
	defer file.Close()

	if putFile(fileUrl, file) == nil {
		writeFile(successfulFile, fileName)
	} else {
		writeFile(notSuccessfulFile, fileName)
	}
}

func getFileName(fileUrl string) string {
	fileURL, err := url.Parse(fileUrl)
	checkError(err)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	return segments[len(segments)-1]
}

func putFile(fileUrl string, file *os.File) error {
	var err error
	i, j := 1, 0

	for {
		resp, er := http.Get(fileUrl)

		err = checkError(er)
		if err == nil {
			defer resp.Body.Close()
			err = second(io.Copy(file, resp.Body))
			checkError(err)
			break
		} else if j < 6 && i < 12 {
			time.Sleep(5 * time.Second)
		} else if j < 6 && i >= 12 {
			time.Sleep(1 * time.Minute)

			i = 0
			j++
		} else {
			break
		}

		i++
	}

	return err
}

func writeFile(file *os.File, text string) {
	w := bufio.NewWriter(file)
	err := second(w.WriteString(text + "\n"))

	checkError(err)
	w.Flush()
}

func checkError(err error) error {
	if err != nil {
		writeFile(logFile, fmt.Sprintf("error = %v", err))
	}

	return err
}

func second[F any, S any](_ F, val S) S {
	return val
}
