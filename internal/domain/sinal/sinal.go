package sinal

type Sinal struct {
	ID          int
	Significado string
	Descricao   string
	Url         string
}

func NewSinal(significado string, desc string, url string) *Sinal {
	return &Sinal{
		Significado: significado,
		Descricao:   desc,
		Url:         url,
	}
}
