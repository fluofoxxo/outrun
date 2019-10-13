package netobj

import (
    "fmt"
    "strconv"
)

type ChaoSpinPrize struct {
    ID     string `json:"chaoId"`
    Level  int64  `json:"level"`
    Rarity int64  `json:"rarity"`
}

func NetChaoToChaoSpinPrize(chao Chao) ChaoSpinPrize {
    id := chao.ID
    level := chao.Level
    rarity := chao.Rarity
    return ChaoSpinPrize{
        id,
        level,
        rarity,
    }
}

func ChaoIDToChaoSpinPrize(chid string) ChaoSpinPrize {
    id := chid
    level := int64(0)                               // TODO: check if the game is accepting of this value...
    rarityInt, err := strconv.Atoi(string(chid[2])) // rarity is third digit
    if err != nil {
        panic(err) // TODO: handle better
    }
    rarity := int64(rarityInt)
    return ChaoSpinPrize{
        id,
        level,
        rarity,
    }
}

func CharacterToChaoSpinPrize(char Character) ChaoSpinPrize {
    id := char.ID
    level := char.Level
    rarity := int64(100) // used so the game recognizes that it is a character
    return ChaoSpinPrize{
        id,
        level,
        rarity,
    }
}

func CharacterIDToChaoSpinPrize(cid string) ChaoSpinPrize {
    id := cid
    level := int64(0)    // TODO: check if the game is accepting of this value...
    rarity := int64(100) // used so the game recognizes that it is a character
    return ChaoSpinPrize{
        id,
        level,
        rarity,
    }
}

func GenericIDToChaoSpinPrize(id string) ChaoSpinPrize {
    idStarter := string(id[0])
    if idStarter == "3" { // Character ID
        return CharacterIDToChaoSpinPrize(id)
    } else if idStarter == "4" { // Chao ID
        return ChaoIDToChaoSpinPrize(id)
    } else {
        // TODO: Gracefully handle error
        err := fmt.Errorf("invalid ID '%v'", id)
        panic(err)
    }
}
