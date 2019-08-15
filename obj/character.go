package obj

type Character struct {
	ID            string `json:"characterId"`
	Cost          int64  `json:"numRings"`         // interestingly, is used for both buying the character and for levelling up...
	NumRedRings   int64  `json:"numRedRings"`      // ?
	Price         int64  `json:"priceNumRings"`    // should mirror Character.Cost, as it seems to be used in _some_ functions
	PriceRedRings int64  `json:"priceNumRedRings"` // ?
}
