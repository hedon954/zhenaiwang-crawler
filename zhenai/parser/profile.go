package parser
//
//import (
//	"regexp"
//	"zhenaiwang-crawler/engine"
//	"zhenaiwang-crawler/model"
//)
//
///**
// 个人信息解析器
//*/
//
////<div class="id" data-v-499ba28c>ID：1592678438</div> 1592678438
//const idRegex string = `<div class="id" data-v-499ba28c>ID：([\d]+)</div>`
//
////<div class="m-btn purple" data-v-8b1eac0c>未婚</div> 未婚
////<div class="m-btn purple" data-v-8b1eac0c>22岁</div> 22岁
////<div class="m-btn purple" data-v-8b1eac0c>天蝎座(10.23-11.21)</div> 天蝎座(10.23-11.21)
////<div class="m-btn purple" data-v-8b1eac0c>162cm</div> 162cm
////<div class="m-btn purple" data-v-8b1eac0c>47kg</div> 47kg
////<div class="m-btn purple" data-v-8b1eac0c>工作地:北京朝阳区</div> 工作地:北京朝阳区
////<div class="m-btn purple" data-v-8b1eac0c>月收入:8千-1.2万</div> 月收入:8千-1.2万
////<div class="m-btn purple" data-v-8b1eac0c>专业顾问</div> 专业顾问
////<div class="m-btn purple" data-v-8b1eac0c>大学本科</div> 大学本科
//const basicRegex string = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
//
////<div class="m-btn pink" data-v-8b1eac0c>汉族</div> 汉族
////<div class="m-btn pink" data-v-8b1eac0c>籍贯:河南济源</div> 籍贯:河南济源
////<div class="m-btn pink" data-v-8b1eac0c>体型:苗条</div> 体型:苗条
////<div class="m-btn pink" data-v-8b1eac0c>不吸烟</div> 不吸烟
////<div class="m-btn pink" data-v-8b1eac0c>不喝酒</div> 不喝酒
////<div class="m-btn pink" data-v-8b1eac0c>租房</div> 租房
////<div class="m-btn pink" data-v-8b1eac0c>未买车</div> 未买车
////<div class="m-btn pink" data-v-8b1eac0c>没有小孩</div> 没有小孩
////<div class="m-btn pink" data-v-8b1eac0c>是否想要孩子:想要孩子</div> 是否想要孩子:想要孩子
////<div class="m-btn pink" data-v-8b1eac0c>何时结婚:时机成熟就结婚</div> 何时结婚:时机成熟就结婚
//const detailRegex string = `<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`
//
////"memberID": 1592678438
//const genderRegex string = `"genderString": "([^"]+)"`
//
////"nickname": "nideaikubao"
//const nameRegex string = `"nickname": "([^"]+)"`
//
////解析个人信息
//func ParseProfile(contents []byte) engine.ParseResult {
//	profile := model.Profile{}
//
//	//匹配ID
//	re := regexp.MustCompile(idRegex)
//	matches := re.FindAllSubmatch(contents, -1)
//	profile.Id = string(matches[0][1])
//
//	//匹配基本信息
//	re = regexp.MustCompile(basicRegex)
//	matches = re.FindAllSubmatch(contents, -1)
//	profile.Marriage = string(matches[0][1])
//	profile.Age = string(matches[1][1])
//	profile.Height = string(matches[3][1])
//	profile.Income = string(matches[6][1])
//	profile.Education = string(matches[8][1])
//
//	//匹配详细信息
//	re = regexp.MustCompile(detailRegex)
//	matches = re.FindAllSubmatch(contents, -1)
//
//	//匹配性别
//	re = regexp.MustCompile(genderRegex)
//	matches = re.FindAllSubmatch(contents, -1)
//	profile.Gender = string(matches[0][1])
//
//	//匹配昵称
//	re = regexp.MustCompile(nameRegex)
//	matches = re.FindAllSubmatch(contents, -1)
//	profile.Name = string(matches[0][1])
//
//	result := engine.ParseResult{
//		Items: []interface{}{profile},
//	}
//
//	profile.PrintProfile()
//
//	return result
//}
