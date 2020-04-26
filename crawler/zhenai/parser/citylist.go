package parser

import (
	"go-crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	maches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range maches{
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}
	return result
}

