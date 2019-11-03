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
    episode-chapter's incentives are. Weird.
*/

var AreaRewards = map[string][]obj.MileageReward{
	"1,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDCombo, 1, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
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
	"11,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDAsteroid, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 4),
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
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 2000, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 5, 4),

		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRing, 8000, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDLaser, 5, 5),
		obj.NewMileageReward(enums.IncentiveTypeEpisode, enums.ItemIDRedRing, 20, 5),
	},
	"16,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDTrampoline, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 2000, 4),
	},
	"17,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDBarrier, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 4000, 4),
	},
	"18,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 2000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 5, 4),
	},
	"19,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 5, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRedRing, 20, 4),
	},
	"20,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDAsteroid, 10, 4),
	},
	"22,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 10, 4),
	},
	"22,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 5000, 4),
	},
	"23,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDBarrier, 10, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDTrampoline, 10, 4),
	},
	"23,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 6000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 6000, 4),
	},
	"24,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 6000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 10, 4),
	},
	"24,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 10, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRedRing, 25, 4),
	},
	"25,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 6000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 6000, 4),
	},
	"26,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 7000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 10, 4),
	},
	"27,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 10, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDInvincible, 10, 4),
	},
	"28,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 10, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 10, 4),
	},
	"29,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 7000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 7000, 4),
	},
	"29,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDAsteroid, 10, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRedRing, 30, 4),
	},
	"30,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDInvincible, 15, 4),
	},
	"31,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 15, 4),
	},
	"31,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		//obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 10, 2),  // UNKNOWN!
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDTrampoline, 15, 4),
	},
	"32,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 15, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 4),
	},
	"33,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 15, 4),
	},
	"33,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDInvincible, 15, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 4),
	},
	"34,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRedRing, 35, 4),
	},
	"36,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 35, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDAsteroid, 35, 4),
	},
	"36,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 35, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 10000, 4),
	},
	"40,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 18000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 25, 4),
	},
	"44,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 18000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 18000, 4),
	},
	"50,1": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 35000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 35, 4),
	},
	"50,2": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDCombo, 35, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDTrampoline, 35, 4),
	},
	"50,3": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDMagnet, 40, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 42000, 4),
	},
	"50,4": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDRing, 47000, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDAsteroid, 45, 4),
	},
	"50,5": []obj.MileageReward{
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 1),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDDrill, 50, 2),
		obj.NewMileageReward(enums.IncentiveTypeFriend, -1, 1, 3),
		obj.NewMileageReward(enums.IncentiveTypePoint, enums.ItemIDLaser, 50, 4),
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
