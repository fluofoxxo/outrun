package obj

type ChaoPrize struct {
    ID     string `json:"chao_id"`
    Rarity int64  `json:"rarity,string"`
}

func NewChaoPrize(id string, rarity int64) ChaoPrize {
    return ChaoPrize{
        id,
        rarity,
    }
}
