package constobjs

import (
    "strconv"

    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/obj"
)

var AreaReward = map[string][]obj.MileageReward{
    "1,1,1": invincibleItem(1),
    "1,1,2": invincibleItem(2),
    "1,1,3": invincibleItem(3),
    "1,1,4": invincibleItem(4),
    "1,1,5": invincibleItem(5),
}

func GetAreaReward(chapter, episode, point int64) []obj.MileageReward {
    chapS := strconv.Itoa(int(chapter))
    epS := strconv.Itoa(int(episode))
    pointS := strconv.Itoa(int(point))
    getS := chapS + "," + epS + "," + pointS
    return AreaReward[getS]
}

func invincibleItem(point int64) []obj.MileageReward {
    return []obj.MileageReward{obj.NewMileageReward(enums.ItemIDInvincible, point)}
}
