package parser

import (
	"regexp"
	"reptile/engine"
	"reptile/model"
	"strconv"
)

var ageRe = regexp.MustCompile(` <td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^>]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^>]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^>]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)
var guessRe = regexp.MustCompile(`a class="exp-user-name" [^>]*  href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)">([^>]+)</a>`)

var idUrlRe = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/([\d])+`)

func ParseProFile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	matchs := guessRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	for _, m := range matchs {

		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParserFun: ProfileParser(string(m[2])),
		})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) //获取一个
	if len(match) >= 2 {
		//fmt.Printf("句子： %,值：%s\n", match, match[1])
		return string(match[1])
	}
	return ""
}

func ProfileParser(name string) engine.ParserFun {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProFile(c, url, name)
	}
}
