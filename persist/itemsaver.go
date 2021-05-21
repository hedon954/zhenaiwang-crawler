package persist

import (
	"log"
	"zhenaiwang-crawler/model"
)

/**
ItemSaver：负责数据的存储工作
*/

var savedItem map[model.Profile]bool

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	savedItem = make(map[model.Profile]bool)
	itemIndex := 1
	go func() {
		for {
			item := <-out
			profile := item.(model.Profile)
			if itemIsDuplicate(profile) {
				continue
			}
			log.Printf("=========== Got item %d ===========\n", itemIndex)
			profile.PrintProfile()
			itemIndex++
		}
	}()
	return out
}

func itemIsDuplicate(profile model.Profile) bool {
	if savedItem[profile] {
		return true
	} else {
		savedItem[profile] = true
		return false
	}
}
