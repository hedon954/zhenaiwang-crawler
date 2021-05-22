package view

import (
	"os"
	"testing"
	"zhenaiwang-crawler/frontend/model"
	model2 "zhenaiwang-crawler/model"
)

func TestTemplate(t *testing.T) {

	view := CreateSearchResultView("template.html")

	page := model.SearchResult{}

	profile := model2.Profile{
		Id:           "1700471119",
		Name:         "风轻云淡",
		Gender:       "男士",
		Age:          45,
		Height:       172,
		Avatar:       "https://photo.zastatic.com/images/photo/425118/1700471119/3905198160261551.jpg",
		LivaPlace:    "福建龙岩",
		Income:       "12001-20000元",
		Marriage:     "离异",
		Education:    "",
		Introduction: "顺其自然，一切随缘，努力工作，努力生活，踏实过日子。",
		Url:          "http://album.zhenai.com/u/1700471119",
	}

	for i := 0; i < 10; i++ {
		page.Profiles = append(page.Profiles, profile)
	}

	out, err := os.Create("template-test.html")

	if err != nil {
		panic(err)
	}

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}

}
