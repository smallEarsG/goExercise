package parser

import (
	"regexp"
	"reptile/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/kelamayi/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User"+name) // 像切片元素中添加一个元素 并且返回切片
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParserFun: ProfileParser(string(m[2])),
			})

	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParserFun: ParseCity,
		})
	}
	return result
}
