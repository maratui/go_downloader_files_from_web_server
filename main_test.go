package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestWithoutStoppingWebServer(t *testing.T) {
	runCommand("sudo systemctl start nginx && mkdir ./bin")
	downloadAllFiles()
	runCommand("sudo nginx -s stop")

	want := []byte("maksim-drugaja-realnost.mp3\nmaksim-shtampy.mp3\nmaksim-znaesh-li-ty.mp3\n")
	text, err := os.ReadFile("./bin/successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	want = []byte("")
	text, err = os.ReadFile("./bin/not-successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/not-successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	runCommand("rm -rf ./bin/")
}

func TestWithStoppingWebServerBeforeHttpGet5sec(t *testing.T) {
	flag := true

	runCommand("sudo systemctl start nginx && mkdir ./bin")
	successfulFile := createFile(successfulFileName)
	defer successfulFile.Close()
	notSuccessfulFile := createFile(notSuccessfulFileName)
	defer notSuccessfulFile.Close()
	logFile = createFile(logFileName)
	defer logFile.Close()
	for _, fileUrl := range fileUrls {
		fileName := getFileName(fileUrl)
		file := createFile(binPath + fileName)
		defer file.Close()

		if flag {
			runCommand("sudo nginx -s stop")
		}
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
			} else if j < 5 && i < 12 {
				time.Sleep(5 * time.Second)
				writeFile(logFile, "5 sec")
				if flag {
					runCommand("sudo systemctl start nginx")
					flag = false
				}
			} else if j < 5 && i >= 12 {
				time.Sleep(1 * time.Minute)
				writeFile(logFile, "1 min")

				i = 0
				j++
			} else {
				writeFile(logFile, "10 min")
				break
			}

			i++
		}

		if err == nil {
			writeFile(successfulFile, fileName)
		} else {
			writeFile(notSuccessfulFile, fileName)
		}
	}
	runCommand("sudo nginx -s stop")

	want := []byte("maksim-drugaja-realnost.mp3\nmaksim-shtampy.mp3\nmaksim-znaesh-li-ty.mp3\n")
	text, err := os.ReadFile("./bin/successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	want = []byte("")
	text, err = os.ReadFile("./bin/not-successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/not-successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	runCommand("rm -rf ./bin")
}

func TestWithStoppingWebServerBeforeHttpGet1min(t *testing.T) {
	flag := true

	runCommand("sudo systemctl start nginx && mkdir ./bin")
	successfulFile := createFile(successfulFileName)
	defer successfulFile.Close()
	notSuccessfulFile := createFile(notSuccessfulFileName)
	defer notSuccessfulFile.Close()
	logFile = createFile(logFileName)
	defer logFile.Close()
	for _, fileUrl := range fileUrls {
		fileName := getFileName(fileUrl)
		file := createFile(binPath + fileName)
		defer file.Close()

		if flag {
			runCommand("sudo nginx -s stop")
		}
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
			} else if j < 5 && i < 12 {
				time.Sleep(5 * time.Second)
				writeFile(logFile, "5 sec")
			} else if j < 5 && i >= 12 {
				time.Sleep(1 * time.Minute)
				writeFile(logFile, "1 min")
				if flag {
					runCommand("sudo systemctl start nginx")
					flag = false
				}

				i = 0
				j++
			} else {
				writeFile(logFile, "10 min")
				break
			}

			i++
		}

		if err == nil {
			writeFile(successfulFile, fileName)
		} else {
			writeFile(notSuccessfulFile, fileName)
		}
	}
	runCommand("sudo nginx -s stop")

	want := []byte("maksim-drugaja-realnost.mp3\nmaksim-shtampy.mp3\nmaksim-znaesh-li-ty.mp3\n")
	text, err := os.ReadFile("./bin/successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	want = []byte("")
	text, err = os.ReadFile("./bin/not-successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/not-successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	runCommand("rm -rf ./bin")
}

func TestWithStoppingWebServerBeforeHttpGet10min(t *testing.T) {
	flag := true

	runCommand("sudo systemctl start nginx && mkdir ./bin")
	successfulFile := createFile(successfulFileName)
	defer successfulFile.Close()
	notSuccessfulFile := createFile(notSuccessfulFileName)
	defer notSuccessfulFile.Close()
	logFile = createFile(logFileName)
	defer logFile.Close()
	for _, fileUrl := range fileUrls {
		fileName := getFileName(fileUrl)
		file := createFile(binPath + fileName)
		defer file.Close()

		if flag {
			runCommand("sudo nginx -s stop")
		}
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
			} else if j < 5 && i < 12 {
				time.Sleep(5 * time.Second)
				writeFile(logFile, "5 sec")
			} else if j < 5 && i >= 12 {
				time.Sleep(1 * time.Minute)
				writeFile(logFile, "1 min")

				i = 0
				j++
			} else {
				writeFile(logFile, "10 min")
				if flag {
					runCommand("sudo systemctl start nginx")
					flag = false
				}
				break
			}

			i++
		}

		if err == nil {
			writeFile(successfulFile, fileName)
		} else {
			writeFile(notSuccessfulFile, fileName)
		}
	}
	runCommand("sudo nginx -s stop")

	want := []byte("maksim-shtampy.mp3\nmaksim-znaesh-li-ty.mp3\n")
	text, err := os.ReadFile("./bin/successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	want = []byte("maksim-drugaja-realnost.mp3\n")
	text, err = os.ReadFile("./bin/not-successful.txt")
	if err != nil || !bytes.Equal(want, text) {
		t.Fatalf("./bin/not-successful.txt = %q, %v; want match for %q, <nil>\n", text, err, want)
	}
	runCommand("rm -rf ./bin")
}

func runCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Run()
}
