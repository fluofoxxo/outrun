package netobj

type PlayerVarious struct {
	CmSkipCount          int64 `json:"cmSkipCount"` // no clear purpose
	EnergyRecoveryMax    int64 `json:"energyRecoveryMax"`
	EnergyRecoveryTime   int64 `json:"energyRecveryTime"` // time until...?
	OnePlayCmCount       int64 `json:"onePlayCmCount"`
	OnePlayContinueCount int64 `json:"onePlayContinueCount"`
	IsPurchased          int64 `json:"isPurchased"`
}

func DefaultPlayerVarious() PlayerVarious {
	cmSkipCount := int64(5)
	energyRecoveryMax := int64(5)  //max lives should be five
	energyRecoveryTime := int64(660) // eleven minutes
	onePlayCmCount := int64(0)
	onePlayContinueCount := int64(5)
	isPurchased := int64(0)
	return PlayerVarious{
		cmSkipCount,
		energyRecoveryMax,
		energyRecoveryTime,
		onePlayCmCount,
		onePlayContinueCount,
		isPurchased,
	}
}

