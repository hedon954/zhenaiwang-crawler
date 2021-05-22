package controller

import (
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"zhenaiwang-crawler/frontend/model"
	"zhenaiwang-crawler/frontend/view"
	model2 "zhenaiwang-crawler/model"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//创建一个搜索前端控制器
func CreateSearchResultHandler(templateFilename string) SearchResultHandler {

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(templateFilename),
		client: client,
	}
}

//localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("q"))
	from, err := strconv.Atoi(request.FormValue("from"))
	//填错就是0
	if err != nil {
		from = 0
	}
	page, err := h.GetSearchResult(q, from)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

}

const pageSize = 10

//执行具体的 ElasticSearch 查询操作
func (h SearchResultHandler) GetSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Profiles = resp.Each(reflect.TypeOf(model2.Profile{}))

	//分页
	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom =
			(result.Start - 1) /
				pageSize * pageSize
	}
	result.NextFrom =
		result.Start + len(result.Profiles)

	return result, nil
}
