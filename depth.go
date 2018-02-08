package hitbtc

import (
	"encoding/json"
	"time"
)
type Item struct  {
	Price          float64   `json:"price,string"`
	Amount         float64   `json:"size,string"`
}
type Depth struct {
	BidList         []Item `json:"bid,string"`
	AskList         []Item  `json:"ask,string"`
	Timestamp   time.Time `json:"timestamp"`
}

func (d *Depth) UnmarshalJSON(data []byte) error {
	var err error
	type Alias  Depth
	aux := &struct {
		//Timestamp string `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
