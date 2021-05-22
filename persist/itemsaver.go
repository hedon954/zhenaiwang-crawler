package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"zhenaiwang-crawler/model"
)

/**
ItemSaver：负责数据的存储工作
*/

var savedItem map[model.Profile]bool

//执行信息保存操作
func ItemSaver(esIndex string, esType string) (chan interface{}, error) {

	//创建一个 Elastic Search 客户端
	client, err := elastic.NewClient(
		//运行在 Docker，处于一个内网环境的话，需要关闭 sniff
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan interface{})
	savedItem = make(map[model.Profile]bool)
	itemIndex := 1
	go func() {
		for {
			item := <-out
			profile := item.(model.Profile)
			//去重
			if itemIsDuplicate(profile) {
				continue
			}
			log.Printf("=========== Got item %d ===========\n", itemIndex)
			profile.PrintProfile()

			//保存数据到 ElasticSearch
			err := save(esIndex, esType, client, profile)
			if err != nil {
				log.Printf("Error occurred at saving profile: %v", err)
				continue
			}
			itemIndex++
		}
	}()
	return out, nil
}

//判断信息是否重复
func itemIsDuplicate(profile model.Profile) bool {
	if savedItem[profile] {
		return true
	} else {
		savedItem[profile] = true
		return false
	}
}

//保存信息
func save(esIndex string, esType string, client *elastic.Client, profile model.Profile) error {

	//存数据
	response, err := client.Index().
		Index(esIndex).
		Type(esType).
		Id(profile.Id).
		BodyJson(profile).
		Do(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", response)
	return nil
}
