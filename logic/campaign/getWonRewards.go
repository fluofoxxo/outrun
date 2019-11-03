package campaign

import (
	"log"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
)

func GetRewardsByPoint(episode, chapter, point int64) []obj.MileageReward {
	rewards := constobjs.GetAreaReward(chapter, episode)
	pointRewards := []obj.MileageReward{}
	for _, reward := range rewards {
		if reward.Point == point {
			pointRewards = append(pointRewards, reward)
		}
	}
	return pointRewards
}

func GetWonRewards(oldEpisode, oldChapter, oldPoint, newEpisode, newChapter, newPoint int64) []obj.MileageReward {
	if config.CFile.DebugPrints {
		log.Printf("%v %v %v %v %v %v\n", oldEpisode, oldChapter, oldPoint, newEpisode, newChapter, newPoint)
	}
	oldRewards := constobjs.GetAreaReward(oldChapter, oldEpisode)
	wonRewards := []obj.MileageReward{}
	if oldPoint == newPoint { // didn't win any rewards
		return wonRewards
	}
	if oldChapter != newChapter { // player crossed chapter. _Normally,_ this means that they completed the chapter, but this could also be due to debugging
		// return rewards from previous point of previous chapter
		wonRewards = GetRewardsByPoint(oldEpisode, oldChapter, oldPoint)
		return wonRewards
	}
	// populate a slice of rewards that have been won
	populating := false
	for _, reward := range oldRewards {
		if reward.Point == oldPoint { // start filling out rewards
			populating = true
		} else if reward.Point == newPoint { // exclude the newPoint rewards, they are never earned when the player just arrives at that point
			break
		}
		if populating {
			wonRewards = append(wonRewards, reward)
		}
	}
	return wonRewards
}
