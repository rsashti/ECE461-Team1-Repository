package main

import (
	"ECE461-Team1-Repository/metrics"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	NPM    = 0
	GITHUB = 1
)

type Link struct {
	site int
	name string
}

var GITHUB_TOKEN string
var LOG_LEVEL string
var LOG_FILE string

func writeLog() {
	logFileLocation := os.Getenv("LOG_FILE")
	logFileLocation += "/log.txt"
	fmt.Println("Log file created at: ", logFileLocation) //for debugging purpose. take it out later

	logFile, err := os.Create(logFileLocation)
	if err != nil {
		log.Fatalf("Failed to create log file")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// write to the log file
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "0" {
		log.Println("LOG_LEVEL is set to 0. don't print anything")
	} else if logLevel == "1" {
		log.Println("LOG_LEVEL is set to 1. do something meaningful here")
	} else if logLevel == "2" {
		log.Println("LOG_LEVEL is set to 2. give more details")
	}
}

func init() {
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	LOG_FILE = os.Getenv("LOG_FILE")
}

func main() {
	//logFile :=os.Getenv("LOG_FILE")
	args := os.Args[1:]
	str := strings.Join(args, "")
	file, err := os.Open(str)
	if err != nil {
		log.Fatalf("Failed to open file!")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called on the os.File object to close the file
	file.Close()

	var links []Link

	// A loop iterates through and prints each of the slice values.
	for _, each_ln := range text {
		var tmpSite int
		var tmpName string
		gitMatch := strings.Contains(each_ln, "github")
		if gitMatch {
			gitLinkMatch := regexp.MustCompile(".*github.com/(.*)")
			tmpName = gitLinkMatch.FindStringSubmatch(each_ln)[1]
			tmpSite = 1
		} else {
			npmLinkMatch := regexp.MustCompile(".*package/(.*)")
			tmpName = npmLinkMatch.FindStringSubmatch(each_ln)[1]
			tmpSite = 0
		}
		newLink := Link{site: tmpSite, name: tmpName}
		links = append(links, newLink)
	}

	for _, tst_print := range links {
		// fmt.Printf("%+v\n", tst_print)
		// metrics.GetMetrics(tst_print.name, GITHUB_TOKEN)
		// if tst_print.site == GITHUB {
		metrics.GetMetrics(tst_print.site, tst_print.name, GITHUB_TOKEN)
		// }
	}

	// GETS ALL THE METRICS IN THIS FUNCTION GIVEN THE URL (ONLY WORKS FOR GITHUB CURRENTLY)
	// metrics.GetMetrics("cloudinary/cloudinary_npm", GITHUB_TOKEN)
	// metrics.GetMetrics("lodash/lodash", GITHUB_TOKEN)
	// metrics.GetMetrics("nullivex/nodist", GITHUB_TOKEN)
}
