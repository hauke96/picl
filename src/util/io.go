package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hauke96/picl/src/log"
)

func DownloadFile(url string, fileName string) error {
	var err error = nil

	// Create output file
	var output *os.File = nil

	// Remove existing file
	if _, err := os.Stat(fileName); err == nil {
		log.Debug("Remove existing file")
		err = os.Remove(fileName)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while removing file %s: %s", fileName, err.Error()))
		}
	}

	// Create output file
	log.Debug("Create new file")
	output, err = os.Create(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while creating file %s: %s", fileName, err.Error()))
	}
	defer output.Close()

	// Donwload data
	log.Debug(fmt.Sprintf("Download data from %s", url))
	response, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while donwloading %s: %s", url, err.Error()))
	}
	defer response.Body.Close()

	// Write data into output File
	log.Debug("Copy response to file")
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while copying response stream from %s to %s: %s", url, fileName, err.Error()))
	}

	return nil
}

func ReadAllLines(file *os.File) ([]string, error) {
	lines := make([]string, 0)

	// defer closing
	defer file.Close()

	// read lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// return lines or error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
