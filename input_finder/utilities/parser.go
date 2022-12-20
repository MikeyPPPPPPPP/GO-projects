package utilities

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//use regular expressions to find input fields
func FindInputTags(body string) []string {
	re := regexp.MustCompile(`(<input .*?>)`) // (type="text" | type="email").*?>)`) //.*?<\/input>`)
	FoundFields := re.FindAllString(body, -1)

	var FinalFeild []string
	for _, item := range FoundFields {
		if strings.Contains(item, `type="hidden"`) == false || strings.Contains(item, `type="checkbox"`) || strings.Contains(item, `type='hidden'`) {
			FinalFeild = append(FinalFeild, item)
		}
	}
	return FinalFeild
}

//returns html as a string
func GetHtml(url string) string {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err := client.Get(url)
	Check(err)

	body, err := ioutil.ReadAll(req.Body)
	Check(err)

	sb := string(body)

	return sb
}
