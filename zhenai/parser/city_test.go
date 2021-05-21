package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

const urlAndNameRegex2 string = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]*>([^<]+)</a>`
const avatarRegex2 string = `target="_blank"><img src="https://photo.zastatic.com/images/photo/[\d]+/[\d]+/[\d]+.png?[^"]+" alt="[^"]+"></a></div>`



func TestParseCity(t *testing.T) {

	//contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun/chengdu")
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	//url和name
	//<a href="http://album.zhenai.com/u/1848766259" target="_blank">红豆</a> 1848766259 红豆
	re := regexp.MustCompile(urlAndNameRegex2)
	urlAndNameMatches := re.FindAllSubmatch(contents, -1)
	for _,match := range urlAndNameMatches{
		for _,m := range match{
			fmt.Print(string(m) + " ")
		}
		fmt.Println()
	}

	//头像
	//target="_blank"><img src="https://photo.zastatic.com/images/photo/439092/1756365177/53152320750083910.png?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="木棉"></a></div>
	re = regexp.MustCompile(avatarRegex2)
	avatarMatches := re.FindAllSubmatch(contents, -1)
	for _,match := range avatarMatches{
		for _,m := range match{
			fmt.Print(string(m) + " ")
		}
		fmt.Println()
	}
	//
	//result := ParseCity(contents)
	//fmt.Println(result)

}
