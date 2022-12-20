package utilities

type PageInputs struct {
	Url       string
	InputTags []string
}

func stringInStruct(in []string, check string) bool {
	good := false
	for _, item := range in {
		if item == check {
			good = true
		}
	}
	return good
}

func MakeUniqStruct(in []PageInputs) []PageInputs {
	var finalStruct []PageInputs

	var uniq []string

	//iderat over InputTags
	for _, item := range in {

		var inputtags []string

		for _, inputs := range item.InputTags {

			//if the input is not in the uniq slice, add it
			if stringInStruct(uniq, inputs) == false {

				uniq = append(uniq, inputs)
				inputtags = append(inputtags, inputs)
			}
		}
		//if there are iput tags in the url, add it to the final struct
		if len(inputtags) > 0 {
			temp := PageInputs{item.Url, inputtags}
			finalStruct = append(finalStruct, temp)
		}

	}

	return finalStruct
}
