package engine

type ParserFun func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFun
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
