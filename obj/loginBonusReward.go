package obj

type LoginBonusReward struct {
	SelectRewardList []SelectReward `json:"selectRewardList"`
}

func NewLoginBonusReward(selectRewardList []SelectReward) LoginBonusReward {
	return LoginBonusReward{
		selectRewardList,
	}
}
