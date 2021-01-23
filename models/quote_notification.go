package models

type QuoteNotification struct {
	Timestamp      int64   `json:"timestamp"`
	InstrumentName string  `json:"instrument_name"`
	BestBidPrice   float64 `json:"best_bid_price"`
	BestBidAmount  float64 `json:"best_bid_amount"`
	BestAskPrice   float64 `json:"best_ask_price"`
	BestAskAmount  float64 `json:"best_ask_amount"`
}

func (qn QuoteNotification) Spread() float64 {
	return qn.BestAskPrice/qn.BestBidPrice - 1
}

func (qn QuoteNotification) MidPrice() float64 {
	return (qn.BestBidPrice + qn.BestAskPrice) / 2
}
