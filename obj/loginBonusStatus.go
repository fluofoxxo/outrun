package obj

type LoginBonusStatus struct {
	NumLogin      int64 `json:"numLogin"`      // how many logins
	NumBonus      int64 `json:"numBonus"`      // how many bonuses?
	LastBonusTime int64 `json:"lastBonusTime"` // last time gotten a bonus (last login, essentially)
}

func NewLoginBonusStatus(numLogin, numBonus, lastBonusTime int64) LoginBonusStatus {
	return LoginBonusStatus{
		numLogin,
		numBonus,
		lastBonusTime,
	}
}
