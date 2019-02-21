package mock

type Retriever struct {
	Txt string
}

func (r Retriever) Get(url string) string {
	return r.Txt
}
