package obj

type Chao struct {
    ID     string `json:"chaoId"`
    Rarity int64  `json:"rarity"`
    Hidden int64  `json:"hidden"` // this value is required, but is never really used in game... Best to keep at "1"
}

func NewChao(id string, rarity, hidden int64) Chao {
    return Chao{
        id,
        rarity,
        hidden,
    }
}
