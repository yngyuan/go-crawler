package parser

import (
	"go-crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	maches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//limit := 10
	for _, m := range maches{
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}

