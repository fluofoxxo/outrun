package constobjs

import (
    "github.com/fluofoxxo/outrun/obj"
)

func GetMileageIncentives(episode, chapter int64) []obj.MileageIncentive {
    rewards := GetAreaReward(chapter, episode)
    // convert rewards to mileageIncentives
    // TODO: probably not the best to do this at runtime...
    incentives := []obj.MileageIncentive{}
    for _, reward := range rewards {
        var incentive obj.MileageIncentive
        incentive.Type = reward.Type
        incentive.ItemID = reward.ItemID
        // TODO: incentive.FriendID should likely be satisfied at some point
        incentive.NumItem = reward.NumItem
        incentive.PointID = reward.Point
        incentives = append(incentives, incentive)
    }
    return incentives
}
