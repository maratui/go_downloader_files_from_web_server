package main

import "os"

var fileUrls = []string{
	"http://localhost:9990/maksim-drugaja-realnost.mp3",
	"http://localhost:9990/maksim-shtampy.mp3",
	"http://localhost:9990/maksim-znaesh-li-ty.mp3"}

var binPath = "./bin/"

var successfulFileName = "./bin/successful.txt"
var notSuccessfulFileName = "./bin/not-successful.txt"
var logFileName = "./bin/log.txt"

var logFile *os.File
