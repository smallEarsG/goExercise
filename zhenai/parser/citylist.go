package parser

import (
	"regexp"
	"reptile/engine"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"+[^>]*>(.*[^<])</a>`

func ParserCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "City"+string(m[2])) // 像切片元素中添加一个元素 并且返回切片
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParserFun: ParseCity,
			})

	}
	return result
}
