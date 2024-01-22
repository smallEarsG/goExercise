package engine

type ParserFun func(contents []byte, url string) ParseResult
type Request struct {
	Url       string
	ParserFun ParserFun
}
type ParseResult struct {
	Requests []Request
	Items    []Item
}
type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
