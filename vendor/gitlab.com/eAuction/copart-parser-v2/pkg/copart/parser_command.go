package copart

type ParserCommand interface {
	Command() ParserCommandEnvelope
}

//ParserCommandEnvelope - represents parser command
type ParserCommandEnvelope struct {
	Command string      `json:"command"`
	Data    interface{} `json:"data"`
}

type ParserCommandStart struct {
	CopartUserName *string `json:"copart_username,omitempty"`
	CopartPassword *string `json:"copart_password,omitempty"`
}

func (p ParserCommandStart) Command() ParserCommandEnvelope {
	return ParserCommandEnvelope{
		Command: "start",
		Data:    p,
	}
}

type ParserCommandJoin struct {
	SaleID string `json:"sale_id"`
}

func (p ParserCommandJoin) Command() ParserCommandEnvelope {
	return ParserCommandEnvelope{
		Command: "join",
		Data:    p,
	}
}

type ParserCommandBid struct {
	SaleID string `json:"sale_id"`
	LotID  string `json:"lot_id"`
	MaxBid Amount `json:"max_bid"`
	ID     string `json:"e2e_id"`
}

func (p ParserCommandBid) Command() ParserCommandEnvelope {
	return ParserCommandEnvelope{
		Command: "bid",
		Data:    p,
	}
}
