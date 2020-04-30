package parser

import (
	"go-crawler/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	maches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range maches{
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})

	}
	return result
}
