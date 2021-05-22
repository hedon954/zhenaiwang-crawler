package model

import (
	"fmt"
	"strconv"
)

//个人信息包装类
type Profile struct {
	Id           string
	Name         string
	Gender       string
	Age          int
	Height       int
	Avatar       string //头像
	LivaPlace    string //居住地
	Income       string //收入
	Marriage     string //婚姻状况
	Education    string //教育水平
	Introduction string //自我介绍
	Url          string //Url
}

func (p Profile) PrintProfile() {
	fmt.Println("Profile{\n" +
		"        Id: " + p.Id + "\n" +
		"        昵称: " + p.Name + "\n" +
		"        性别: " + p.Gender + "\n" +
		"        年龄: " + strconv.Itoa(p.Age) + "\n" +
		"        身高: " + strconv.Itoa(p.Height) + "\n" +
		"        头像: " + p.Avatar + "\n" +
		"        居住地: " + p.LivaPlace + "\n" +
		"        月薪 " + p.Income + "\n" +
		"        婚姻状况: " + p.Marriage + "\n" +
		"        教育水平: " + p.Education + "\n" +
		"        自我介绍: " + p.Introduction + "\n" +
		"        珍爱网链接: " + p.Url + "\n" +
		"}")
}

//func JsonToObj(o interface{}) (Profile, error){
//
//	marshal, err := json.Marshal(o)
//
//	if err != nil {
//		return Profile{},err
//	}
//
//
//
//}
