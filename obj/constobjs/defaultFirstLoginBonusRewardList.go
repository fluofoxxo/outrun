package constobjs

import "github.com/fluofoxxo/outrun/obj"

var DefaultFirstLoginBonusRewardList = func() []obj.LoginBonusReward {
	return []obj.LoginBonusReward{
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 10),
						obj.NewItem("910000", 10000),
						obj.NewItem("240000", 3),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 20),
						obj.NewItem("910000", 20000),
						obj.NewItem("240000", 3),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 20),
						obj.NewItem("910000", 30000),
						obj.NewItem("240000", 3),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("230000", 1),
						obj.NewItem("910000", 50000),
						obj.NewItem("240000", 3),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("300013", 1),
						obj.NewItem("910000", 50000),
						obj.NewItem("240000", 3),
					},
				),
			},
		),
	}
}()
