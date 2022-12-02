/*
This will read a file of urls and look for injectable prameters like
?, &, or = 

*/

package main

import (
	"fmt"
	"strings"
	"log"
	"io/ioutil"
)


func Check(err error){
	if err != nil{
		log.Fatal(err)
	}
}


//this will pase a file for links and return a slice of links
func ReadFile(filename string) string{

	//read the file
	file, err := ioutil.ReadFile(filename)
	Check(err)

	//this convets the byte array to a string
	contents := string(file)

	return contents
}
//sting of urls to a slice
func GetURLs(content string) []string{
	temp := strings.Split(content, "\n")
	return temp
}
//returns the last element in a path
func LastInThePath(url string) string{
	temp := strings.Split(url, "/")
	return temp[len(temp) - 1]
}
//last element containts &  or  ?  or =
func ContainsPrameter(url string) bool{
	if strings.Contains(LastInThePath(url), "?") || strings.Contains(LastInThePath(url), "&") || strings.Contains(LastInThePath(url), "="){
		return true
	} 
	return false
}

func main(){
	var contents []string
	contents = GetURLs(ReadFile("fortinet.txt"))

	for _, url := range contents{
		if ContainsPrameter(url){
			fmt.Println(url)
		}
	}
	

}
