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

	wonRewards := []obj.MileageReward{}

	if oldEpisode == newEpisode { // still in the same episode
		if config.CFile.DebugPrints {
			log.Println("oldEpisode == newEpisode")
		}
		if oldChapter == newChapter { // still in the same chapter
			if config.CFile.DebugPrints {
				log.Println("oldChapter == newChapter")
			}
			// player should only be moving across the map
			if oldPoint == newPoint { // no movement at all
				return wonRewards
			} else if oldPoint > newPoint { // something fishy is happening...
				log.Printf("[WARN] oldPoint (%v) > newPoint (%v), but still same area! Something is going wrong!\n", oldPoint, newPoint)
			} else if oldPoint < newPoint { // correct behavior, is moving along
				for i := oldPoint + 1; i <= newPoint; i++ { // range (oldPoint, newPoint]
					if i == 5 { // if boss point
						// Because we know that the chapter is the same,
						// we also know that the last point is a boss point,
						// which requires the showdown to get the rewards.
						// As such, we don't earn rewards when just arriving
						// at this point, only after.
						break
					}
					theseRewards := GetRewardsByPoint(oldEpisode, oldChapter, int64(i))
					wonRewards = append(wonRewards, theseRewards...)
				}
			}
		} else if oldChapter > newChapter { // something fishy is happening...
			if config.CFile.DebugPrints {
				log.Println("oldChapter > newChapter")
			}
			log.Printf("[WARN] Improper movement across chapters (%v to %v)!\n", oldChapter, newChapter)
		} else if oldChapter < newChapter { // same episode, new chapter
			if config.CFile.DebugPrints {
				log.Println("oldChapter < newChapter")
			}
			if newPoint != 0 { // how did they cross chapters this way?
				log.Printf("[WARN] Player entered a new chapter, but somehow reached point %v!\n", newPoint)
			} else { // normal
				for i := oldPoint + 1; i < 5; i++ { // player completed oldChapter, so get all rewards (oldPoint, 5]
					// don't ignore the point 5 of the oldChapter
					theseRewards := GetRewardsByPoint(oldEpisode, oldChapter, int64(i))
					wonRewards = append(wonRewards, theseRewards...)
				}
				lastRewards := GetRewardsByPoint(oldEpisode, oldChapter, 5)
				wonRewards = append(wonRewards, lastRewards...)
			}
		}
	} else if oldEpisode > newEpisode { // something fishy is happening...
		if config.CFile.DebugPrints {
			log.Println("oldEpisode > newEpisode")
		}
		log.Printf("[WARN] oldEpisode (%v) > newEpisode (%v)! Something is going wrong!\n", oldEpisode, newEpisode)
	} else if oldEpisode < newEpisode { // crossed episodes, get all rewards from previous chapter [oldPoint, 5]
		if config.CFile.DebugPrints {
			log.Println("oldEpisode < newEpisode")
		}
		if newPoint != 0 { // how did they cross episodes this way?
			log.Printf("[WARN] Player entered a new episode, but somehow reached point %v!\n", newPoint)
		} else {
			for i := oldPoint + 1; i < 5; i++ {
				// don't ignore the point 5 of the oldChapter
				theseRewards := GetRewardsByPoint(oldEpisode, oldChapter, int64(i))
				wonRewards = append(wonRewards, theseRewards...)
			}
			lastRewards := GetRewardsByPoint(oldEpisode, oldChapter, 5)
			wonRewards = append(wonRewards, lastRewards...)
		}
	}

	return wonRewards

	/*
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
	*/
}
