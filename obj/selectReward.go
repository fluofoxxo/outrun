package obj

type SelectReward struct {
	ItemList []Item `json:"itemList"`
}

func NewSelectReward(itemList []Item) SelectReward {
	return SelectReward{
		itemList,
	}
}
