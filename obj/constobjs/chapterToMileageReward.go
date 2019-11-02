package constobjs

import (
	"log"
	"strconv"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/obj"
)

/*
Observances:
  - When at the final point expecting to go to the next episode-chapter, the rewards earned at the end of the episode-chapter should _actually_ be whatever the current
    episode-chapter's incentive's are. Weird.
*/

var AreaRewards = map[string][]obj.MileageReward{
	"1,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDCombo, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing1, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"2,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDTrampoline, 1, 5), // Spring
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"3,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDMagnet, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"4,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDBarrier, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"5,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDInvincible, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"6,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeChapter, enums.ItemIDRing, 10000, 5),
		obj.NewMileageReward(enums.IncentiveTypeChapter, enums.ItemIDRedRing, 10, 5),
	},
	"6,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDInvincible, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},

	"7,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDDrill, 2, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"8,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDLaser, 2, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"9,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 5000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDAsteroid, 2, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"10,1": []obj.MileageReward{ // NO DOCS. This is an assumption!
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 10, 5),
	},
	"11,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, enums.CharaTypeTails, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3), // TODO: Find animal ID, if needed

		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 4),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDCombo, 5, 5),
	},
	"12,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1), // TODO: Find story ID
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3), // TODO: Find story ID
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 3, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDDrill, 5, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
	},
	"13,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 5, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDAsteroid, 5, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
	},
	"14,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRedRing, 15, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDInvincible, 5, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
	},
	"15,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 3), // This is also story...?
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 20000, 5),
	},
	"16,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 2000, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDInvincible, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 2000, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 5, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDLaser, 5, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
	},
	"17,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDBarrier, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 4000, 4),
	},
	"22,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 10, 4),
	},
	"25,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 4),
	},
}

func GetAreaReward(chapter, episode int64) []obj.MileageReward { // TODO: change to (episode, chapter int64)
	chapS := strconv.Itoa(int(chapter))
	epS := strconv.Itoa(int(episode))
	getS := epS + "," + chapS
	if config.CFile.DebugPrints {
		log.Println("GetAreaReward(", epS, ", ", chapS, ")")
		log.Println(getS)
	}
	return AreaRewards[getS]
}

func invincibleItem(point int64) obj.MileageReward {
	return obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDInvincible, 5, point)
}
