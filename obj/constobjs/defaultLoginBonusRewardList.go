package constobjs

import "github.com/fluofoxxo/outrun/obj"

var DefaultLoginBonusRewardList = func() []obj.LoginBonusReward {
	return []obj.LoginBonusReward{
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 3000),
						obj.NewItem("240000", 1),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 3000),
						obj.NewItem("240000", 1),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 5000),
						obj.NewItem("240000", 1),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 5000),
						obj.NewItem("240000", 1),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 10),
						obj.NewItem("910000", 5000),
						obj.NewItem("240000", 2),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 10000),
						obj.NewItem("240000", 2),
					},
				),
			},
		),
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 10),
						obj.NewItem("910000", 10000),
						obj.NewItem("240000", 2),
					},
				),
			},
		),
	}
}()
