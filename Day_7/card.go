package main

type Card int

const (
	Joker Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	//Jack
	Queen
	King
	Ass
)

func (c Card) Compare(s Card) int {
	return int(c) - int(s)
}

func (c Card) String() string {
	switch c {
	case Ass:
		return "A"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Joker:
		// return Jack
		return "J"
	case Ten:
		return "T"
	case Nine:
		return "9"
	case Eight:
		return "8"
	case Seven:
		return "7"
	case Six:
		return "6"
	case Five:
		return "5"
	case Four:
		return "4"
	case Three:
		return "3"
	case Two:
		return "2"
	}

	return "UNDEFINED"
}

func getCard(s rune) Card {
	switch s {
	case 'A':
		return Ass
	case 'K':
		return King
	case 'Q':
		return Queen
	case 'J':
		// return Jack
		return Joker
	case 'T':
		return Ten
	case '9':
		return Nine
	case '8':
		return Eight
	case '7':
		return Seven
	case '6':
		return Six
	case '5':
		return Five
	case '4':
		return Four
	case '3':
		return Three
	case '2':
		return Two
	}

	return 1 << 32
}
