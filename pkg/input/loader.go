package input

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

//mySession holds a Session Cookie so I can download the input fast :D
var mySession = "53616c7465645f5f230686c3505f68db91af869d933816c14850e69d30dde877ce7cc2b5c861a7a892770cef116cc3c8"

func Load(year string, exercise string) (*os.File, error) {
	fileName := fileNameSetter(year, exercise)

	if !fileExists(fileName) {
		//https://adventofcode.com/2018/day/1/input
		err := downloadFile(fileName, fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, exercise))
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func fileNameSetter(year string, exercise string) string {
	exerciseFolder := exercise
	if i, err := strconv.Atoi(exercise); err == nil && i < 10 {
		exerciseFolder = "0" + exerciseFolder
	}
	return fmt.Sprintf("%s/%s_%s.txt", exerciseFolder, year, exercise)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Get the data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: mySession})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
