package main

import (
	ut "example.com/utilities"
)

func main() {

	var firstStruct []ut.PageInputs

	var filename string = "syxsense.txt"

	data := ut.ReadFile(filename)
	urls := ut.GetURLs(data) //slice

	for _, url := range urls {
		inputs := ut.FindInputTags(ut.GetHtml(url))

		firstStruct = append(firstStruct, ut.PageInputs{Url: url, InputTags: inputs})

	}

	//var FinalStruct []ut.PageInputs
	FinalStruct := ut.MakeUniqStruct(firstStruct)
	ut.StructToFile(FinalStruct)

}
