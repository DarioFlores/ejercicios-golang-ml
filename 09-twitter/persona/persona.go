package persona

import "github.com/DarioFlores/09-twitter/datos"

type Persona struct {
	Nombre string
	Tweets []datos.Tweet
}

func (p Persona) TodosLosTweets() []datos.Tweet {
	return p.Tweets;
}

func (p Persona) CantidadDeTweets() int {
	return len(p.Tweets);
}