package parser

import (
	"fmt"
	"regexp"
	"zhenaiwang-crawler/engine"
	"zhenaiwang-crawler/model"
)

/**
  城市解析器：提取出具体每个城市页面中所有用户的跳转链接
 */

//<a href="http://album.zhenai.com/u/1781087887" target="_blank">相遇的那天</a>
const urlAndNameRegex string = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]*>([^<]+)</a>`

//target="_blank"><img src="https://photo.zastatic.com/images/photo/370094/1480374558/24059686869976080.png?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="lucas"></a></div>
const avatarRegex string = `target="_blank"><img src="(https://photo.zastatic.com/images/photo/[^\.]+.[a-z]+g)?[^"]*" alt="[^"]+"></a></div>`

//<td width="180"><span class="grayL">性别：</span>男士</td>
const genderRegex string = `<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

//<td><span class="grayL">居住地：</span>四川成都</td>
const livePlaceRegex string = `<td><span class="grayL">居住地：</span>([^<]+)</td>`

//<td width="180"><span class="grayL">年龄：</span>37</td>
const ageRegex string = `<td width="180"><span class="grayL">年龄：</span>([^<]+)</td>`

//<td><span class="grayL">月   薪：</span>12001-20000元</td>
const salaryRegex string = `<td><span class="grayL">月[^薪]*薪：</span>([^<]+)</td>`

//<td><span class="grayL">学   历：</span>大专</td>
const educationRegex string = `<td><span class="grayL">学[^历]*历：</span>([^<]+)</td>`

//<tr><td width="180"><span class="grayL">婚况：</span>离异</td>
const marriageRegex string = `<tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`

//<td width="180"><span class="grayL">身   高：</span>162</td>
const heightRegex string = `<td width="180"><span class="grayL">身[^高]*高：</span>([^<]+)</td>`

//<div class="introduce">一直以来习惯了强大，才发现会哭的孩子才有糖吃</div>
const introductionRegex string = `</table>[^<]*<div class="introduce">([^<]+)</div>`


//下一页：<li class="paging-item"><a href="http://www.zhenai.com/zhenghun/shanghai/6">下一页</a> </li>
const nextPageRegex string = `<li class="paging-item"><a href="(http://www.zhenai.com/zhenghun/shanghai/[\d]+)">下一页</a> </li>`

//下方拓展链接：<a target="_blank" href="http://www.zhenai.com/zhenghun/shanghai/junren">上海军人征婚</a>
const extendLinkRegex string = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[^"]+)">[^<]+</a>`

//解析具体某一个城市页面
func ParseCity(contents []byte) engine.ParseResult {
	//封装到 ParseResult
	result := engine.ParseResult{}

	//提取出用户url和name
	re := regexp.MustCompile(urlAndNameRegex)
	urlAndNameMatches := re.FindAllSubmatch(contents, -1)

	//提取用户头像
	re = regexp.MustCompile(avatarRegex)
	avatarMatches := re.FindAllSubmatch(contents, -1)

	//提取用户性别
	re = regexp.MustCompile(genderRegex)
	genderMatches := re.FindAllSubmatch(contents, -1)

	//提取用户居住地
	re = regexp.MustCompile(livePlaceRegex)
	livePlaceMatches := re.FindAllSubmatch(contents, -1)

	//提取用户年龄
	re = regexp.MustCompile(ageRegex)
	ageMatches := re.FindAllSubmatch(contents, -1)

	//提取用户月薪（男） | 学历（女）
	re = regexp.MustCompile(salaryRegex)
	salaryMatches := re.FindAllSubmatch(contents, -1)

	re = regexp.MustCompile(educationRegex)
	educationMatches := re.FindAllSubmatch(contents, -1)

	//提取用户婚况
	re = regexp.MustCompile(marriageRegex)
	marriageMatches := re.FindAllSubmatch(contents, -1)

	//提取用户身高
	re = regexp.MustCompile(heightRegex)
	heightMatches := re.FindAllSubmatch(contents, -1)

	//提取用户自我介绍
	re = regexp.MustCompile(introductionRegex)
	introductionMatches := re.FindAllSubmatch(contents, -1)

	salaryI := 0
	educationI := 0
	for i,m := range urlAndNameMatches{
		//封装个人信息
		profile := model.Profile{}
		profile.Id = string(m[1])   								//ID
		profile.Name = string(m[2]) 								//Name
		profile.Gender = string(genderMatches[i][1])				//Gender
		profile.Age = string(ageMatches[i][1])						//Age
		profile.Height = string(heightMatches[i][1])				//Height
		profile.Avatar = string(avatarMatches[i][1])				//Avatar
		profile.LivaPlace = string(livePlaceMatches[i][1])			//LivePlace
		profile.Marriage = string(marriageMatches[i][1])			//Marriage
		profile.Introduction = string(introductionMatches[i][1])  	//Introduction
		profile.Url = "http://album.zhenai.com/u/" + string(m[1])
		if string(genderMatches[i][1]) == "男士" {
			//男的有月薪
			profile.Income = string(salaryMatches[salaryI][1])
			salaryI++
		}else{
			//女的有学历
			profile.Education = string(educationMatches[educationI][1])
			educationI ++
		}

		//用户信息
		result.Items = append(result.Items, profile)
		//用户 URL
		result.Requests = append(result.Requests,
			engine.Request{
				Url: "http://album.zhenai.com/u/" + string(m[1]),
				ParserFunc: engine.NilParser,
		})
		fmt.Printf("User: %s, Url: %s\n", string(m[2]), "http://album.zhenai.com/u/" + string(m[1]))
	}

	//放入下一页的链接
	re = regexp.MustCompile(nextPageRegex)
	nextPageMatches := re.FindAllSubmatch(contents, -1)
	for _, match := range nextPageMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(match[1]),
				ParserFunc: ParseCity,
			})
	}

	//放入网站下方拓展链接
	re = regexp.MustCompile(extendLinkRegex)
	extendLinkMatches := re.FindAllSubmatch(contents, -1)
	for _, match := range extendLinkMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(match[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}