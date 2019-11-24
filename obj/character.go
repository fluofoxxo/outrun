package obj

type Character struct {
	ID            string `json:"characterId"`
	Cost          int64  `json:"numRings"`         // interestingly, is used for both buying the character and for levelling up...
	NumRedRings   int64  `json:"numRedRings"`      // ?
	Price         int64  `json:"priceNumRings"`    // used to limit break, as far as I can tell?
	PriceRedRings int64  `json:"priceNumRedRings"` // ?
}
