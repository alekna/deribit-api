package models

type TradesNotification []Trade

func (tn TradesNotification) Count() int {
	return len(tn)
}

func (tn TradesNotification) VWAP() float64 {
	var size, sizeCurrency float64
	for _, t := range tn {
		size += t.Amount
		sizeCurrency += t.Amount / t.Price
	}
	return size / sizeCurrency
}

func (tn TradesNotification) FirstTrade() Trade {
	return tn[0]
}

func (tn TradesNotification) LastTrade() Trade {
	return tn[len(tn)-1]
}

func (tn TradesNotification) Direction() string {
	return tn.FirstTrade().Direction
}

func (tn TradesNotification) Gap() float64 {
	return tn.LastTrade().Price/tn.FirstTrade().Price - 1
}
