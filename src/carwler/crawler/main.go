package main

import (
	"carwler/engine"
	"carwler/zhenai/parser"
)


func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}
