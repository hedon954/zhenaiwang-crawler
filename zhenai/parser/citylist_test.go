package parser

import (
	"io/ioutil"
	"testing"
)


const zhenAiUrl string = "https://www.zhenai.com/zhenghun"

func TestParseCityList(t *testing.T) {

	contents, err := ioutil.ReadFile("citylist_test_data.html")

	//contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	list := ParseCityList(contents)

	//准确数据
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string {
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}

	//测试url个数
	if len(list.Requests) != resultSize {
		t.Errorf("result should hava %d requests, but had %d", resultSize, len(list.Requests))
	}

	//测试前3个url
	for i, url := range expectedUrls {
		if list.Requests[i].Url != expectedUrls[i] {
			t.Errorf("expected url #%d: %s, but was %s", i, url, list.Requests[i].Url)
		}
	}

	//测试item个数
	if len(list.Items) != resultSize {
		t.Errorf("result should hava %d items, but had %d", resultSize, len(list.Items))
	}

	//测试前3个item
	for i, city := range expectedCities {
		if list.Items[i].(string) != expectedCities[i] {
			t.Errorf("expected city #%d: %s, but was %s", i, city, list.Items[i].(string))
		}
	}


}
