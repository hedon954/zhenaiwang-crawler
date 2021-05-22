package view

import (
	"html/template"
	"io"
	"zhenaiwang-crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

//创建查询结果
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
