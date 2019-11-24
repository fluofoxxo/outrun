package factors

const (
	AnalyticTypeLogins              = iota // fulfilled
	AnalyticTypeStoryStarts                // fulfilled
	AnalyticTypeStoryEnds                  // fulfilled
	AnalyticTypeTimedStarts                // fulfilled
	AnalyticTypeTimedEnds                  // fulfilled
	AnalyticTypePurchaseRings              // fulfilled
	AnalyticTypePurchaseEnergy             // fulfilled
	AnalyticTypePurchaseRedRings           // Placeholder!
	AnalyticTypeSpinItemRoulette           // fulfilled
	AnalyticTypeSpinChaoRoulette           // fulfilled
	AnalyticTypeChangeMainCharacter        // fulfilled
	AnalyticTypeChangeSubCharacter         // fulfilled
	AnalyticTypeChangeMainChao             // fulfilled
	AnalyticTypeChangeSubChao              // fulfilled
	AnalyticTypeAverageStoryScore
	AnalyticTypeAverageTimedScore
	AnalyticTypeSpendRings
	AnalyticTypeSpendRedRings
	AnalyticTypeRevives // fulfilled
)
