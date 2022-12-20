package utilities

//parse the struct of urls and uniq input fileds to a file
/*
url
dasfafdfasdf
sadfsdafsd


url
fasdfdsf
fdsafdsf

url
dfsffdsfsdfsdfsf


*/

import (
	"os"
)

func StructToFile(structure []PageInputs) {
	file, err := os.Create("All_Uniq_inputs.txt")
	Check(err)

	defer file.Close()

	for _, item := range structure {
		topString := "\n\n" + item.Url + "\n"
		_, err := file.WriteString(topString)
		Check(err)
		for _, inputs := range item.InputTags {
			tempString := inputs + "\n"
			_, err := file.WriteString(tempString)
			Check(err)
		}
	}
}
