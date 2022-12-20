package utilities

import (
	"io/ioutil"
	"log"
	"strings"
)

func Check(err error) error {
	if err != nil {
		//log.Printf("error: %q", err)
		log.Fatal(err)
		return nil
	}
	return nil
}

//this will pase a file for links and return a slice of links
func ReadFile(filename string) string {

	//read the file
	file, err := ioutil.ReadFile(filename)
	Check(err)

	//this convets the byte array to a string
	contents := string(file)

	return contents
}

//sting of urls to a slice
func GetURLs(content string) []string {
	temp := strings.Split(content, "\n")
	var urls []string

	for _, x := range temp {
		if strings.HasPrefix(x, "http") {
			urls = append(urls, x)
		}
	}
	return urls
}
