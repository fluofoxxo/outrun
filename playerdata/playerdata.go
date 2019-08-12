package pdata

import (
	"errors"
	"time"

	"github.com/fluofoxxo/outrun/consts"

	"encoding/json"
	"strconv"
)

// TODO: REFACTOR ME

var DEFAULT_ITEMS = fetchItems()

func fetchItems() []ItemInfo {
	items := []ItemInfo{}
	baseNum := 120000
	for i := range make([]byte, 11) { // 0...10
		n := baseNum + i
		s := strconv.Itoa(n)
		item := NewItemInfo(s)
		items = append(items, item)
	}
	return items
}

func PlayerFromJSON(jbytes []byte) (Player, error) {
	var player Player
	err := json.Unmarshal(jbytes, &player)
	return player, err
}

type Player struct {
	UserID          string           `json:"userID"`
	Username        string           `json:"username"`
	Password        string           `json:"password"`
	Key             string           `json:"key"`
	CharacterStates []CharacterState `json:"characterState"` // TODO: change the name of this to CharacterState
	PlayerState     `json:"playerState"`
	ChaoState       []Chao          `json:"chaoState"`
	MileageMapState MileageMapState `json:"mileageMapState"`
}

func (p Player) GetCharacter(cid string) (CharacterState, error) {
	for _, c := range p.CharacterStates {
		if c.CharacterID == cid {
			return c, nil
		}
	}
	return BLANK_CHARACTERSTATE, errors.New("no such character '" + cid + "'")
}
func (p Player) GetMainCharacter() (CharacterState, error) {
	cid := p.PlayerState.MainCharaID
	c, e := p.GetCharacter(cid)
	return c, e
}
func (p Player) GetSubCharacter() (CharacterState, error) {
	cid := p.PlayerState.SubCharaID
	c, e := p.GetCharacter(cid)
	return c, e
}

func MakePlayer(uid, username, password, key string, cs []CharacterState, ps PlayerState, chs []Chao, mms MileageMapState) Player {
	player := Player{
		uid,
		username,
		password,
		key,
		cs,
		ps,
		chs,
		mms,
	}
	return player
}

func NewPlayer(uid, username, password, key string) Player {
	cstates := []CharacterState{}
	for _, cid := range consts.CHAR_IDS {
		char := NewCharacterState(cid)
		cstates = append(cstates, char)
	}
	pstate := NewPlayerState()
	chaostate := []Chao{}
	for _, chid := range consts.CHAO_IDS {
		chao := NewChao(chid) // TODO: use proper values to create Chao (see notes in consts)
		chaostate = append(chaostate, chao)
	}
	mileageMapState := StartingMileageMapState()
	player := MakePlayer(
		uid,
		username,
		password,
		key,
		cstates,
		pstate,
		chaostate,
		mileageMapState,
	)
	return player
}

var BLANK_PLAYER = NewPlayer("", "", "", "")
var BLANK_CHARACTERSTATE = CharacterState{"", 0, 0, 0, []int64{}, []int64{}, 0, 0, 0, 0, 0, 0, 0, []string{}}

type CharacterState struct {
	CharacterID      string   `json:"characterId"`
	Level            int64    `json:"level,string"`
	NumRings         int64    `json:"numRings,string"`
	NumRedRings      int64    `json:"numRedRings,string"`
	AbilityLevel     []int64  `json:"abilityLevel"`
	AbilityNumRings  []int64  `json:"abilityNumRings"`
	Exp              int64    `json:"exp"`
	Star             int64    `json:"star,string"`
	StarMax          int64    `json:"starMax,string"`
	LockCondition    int64    `json:"lockCondition,string"`
	PriceNumRings    int64    `json:"priceNumRings,string"`
	PriceNumRedRings int64    `json:"priceNumRedRings,string"`
	Status           int64    `json:"status"`
	CampaignList     []string `json:"campaignList"` // unknown, do more research
}

func MakeCharacterState(cid string, lvl, nr, nrr int64, al, anr []int64, exp, st, stm, lc, pnr, pnrr, s int64, cl []string) CharacterState {
	cs := CharacterState{
		cid,
		lvl,
		nr,
		nrr,
		al,
		anr,
		exp,
		st,
		stm,
		lc,
		pnr,
		pnrr,
		s,
		cl,
	}
	return cs
}

func NewCharacterState(cid string) CharacterState {
	return MakeCharacterState(
		cid,
		//consts.USR_DEFAULT_CHARACTERSTATE_CHARACTERID,
		consts.USR_DEFAULT_CHARACTERSTATE_LEVEL,
		consts.USR_DEFAULT_CHARACTERSTATE_NUMRINGS,
		consts.USR_DEFAULT_CHARACTERSTATE_NUMREDRINGS,
		consts.USR_DEFAULT_CHARACTERSTATE_ABILITYLEVEL,
		consts.USR_DEFAULT_CHARACTERSTATE_ABILITYNUMRINGS,
		consts.USR_DEFAULT_CHARACTERSTATE_EXP,
		consts.USR_DEFAULT_CHARACTERSTATE_STAR,
		consts.USR_DEFAULT_CHARACTERSTATE_STARMAX,
		consts.USR_DEFAULT_CHARACTERSTATE_LOCKCONDITION,
		consts.USR_DEFAULT_CHARACTERSTATE_PRICENUMRINGS,
		consts.USR_DEFAULT_CHARACTERSTATE_PRICENUMREDRINGS,
		consts.USR_DEFAULT_CHARACTERSTATE_STATUS,
		consts.USR_DEFAULT_CHARACTERSTATE_CAMPAIGNLIST,
	)
}

type ItemInfo struct {
	ItemID  string `json:"itemId"`
	NumItem int64  `json:"numItem,string"`
}

func MakeItemInfo(itemID string, num int64) ItemInfo {
	ii := ItemInfo{
		itemID,
		num,
	}
	return ii
}

func NewItemInfo(itemID string) ItemInfo {
	return MakeItemInfo(itemID, 0) // TODO: make last param const
}

type PlayerState struct {
	NumRings               int64      `json:"numRings,string"`
	NumBuyRings            int64      `json:"numBuyRings,string"`
	NumRedRings            int64      `json:"numRedRings,string"`
	NumBuyRedRings         int64      `json:"numBuyRedRings,string"`
	Energy                 int64      `json:"energy,string"`
	EnergyBuy              int64      `json:"energyBuy,string"`
	EnergyRenewsAt         int64      `json:"energyRenewsAt"` // does 0 mean it is instant?
	Items                  []ItemInfo `json:"items"`
	MumMessages            int64      `json:"mumMessages"` // ?
	RankingLeague          int64      `json:"rankingLeague,string"`
	QuickRankingLeague     int64      `json:"quickRankingLeague,string"`
	NumRouletteTicket      int64      `json:"numRouletteTicket,string"`
	TotalHighScore         int64      `json:"totalHighScore,string"`
	TotalDistance          int64      `json:"totalDistance,string"`
	MaximumDistance        int64      `json:"maximumDistance,string"`
	DailyMissionId         int64      `json:"dailyMissionId,string"`
	DailyMissionEndTime    int64      `json:"dailyMissionEndTime"` // 11:59 pm of current day
	DailyChallengeValue    int64      `json:"dailyChallengeValue"`
	DailyChallengeComplete int64      `json:"dailyChallengeComplete"`
	NumDailyChalCont       int64      `json:"numDailyChalCont"`
	MainCharaID            string     `json:"mainCharaID"`
	SubCharaID             string     `json:"subCharaID"`
	MainChaoID             string     `json:"mainChaoID"`
	SubChaoID              string     `json:"subChaoID"`
	NumPlaying             int64      `json:"numPlaying,string"` // ?
	NumAnimals             int64      `json:"numAnimals,string"`
	NumRank                int64      `json:"numRank,string"`
	EquipItemList          []string   `json:"equipItemList"` // default is list of 3 "-1"s. look to be item ids
	QuickTotalHighScore    int64      `json:"quickTotalHighScore,string"`
}

func MakePlayerState(nrings, nbr, nrr, nbrr, e, eb, era int64, i []ItemInfo, mm, rl, qrl, nrt, ths, td, md, dmi, dmet, dcv, dcc, ndcc int64, mcid, scid, mchid, schid string, np, na, nr int64, eil []string, qths int64) PlayerState {
	ps := PlayerState{
		nrings,
		nbr,
		nrr,
		nbrr,
		e,
		eb,
		era,
		i,
		mm,
		rl,
		qrl,
		nrt,
		ths,
		td,
		md,
		dmi,
		dmet,
		dcv,
		dcc,
		ndcc,
		mcid,
		scid,
		mchid,
		schid,
		np,
		na,
		nr,
		eil,
		qths,
	}
	return ps
}
func NewPlayerState() PlayerState {
	ps := MakePlayerState(
		consts.USR_DEFAULT_PLAYERSTATE_NUMRINGS,
		consts.USR_DEFAULT_PLAYERSTATE_NUMBUYRINGS,
		consts.USR_DEFAULT_PLAYERSTATE_NUMREDRINGS,
		consts.USR_DEFAULT_PLAYERSTATE_NUMBUYREDRINGS,
		consts.USR_DEFAULT_PLAYERSTATE_ENERGY,
		consts.USR_DEFAULT_PLAYERSTATE_ENERGYBUY,
		consts.USR_DEFAULT_PLAYERSTATE_ENERGYRENEWSAT,
		DEFAULT_ITEMS,
		consts.USR_DEFAULT_PLAYERSTATE_MUMMESSAGES,
		consts.USR_DEFAULT_PLAYERSTATE_RANKINGLEAGUE,
		consts.USR_DEFAULT_PLAYERSTATE_QUICKRANKINGLEAGUE,
		consts.USR_DEFAULT_PLAYERSTATE_NUMROULETTETICKET,
		consts.USR_DEFAULT_PLAYERSTATE_TOTALHIGHSCORE,
		consts.USR_DEFAULT_PLAYERSTATE_TOTALDISTANCE,
		consts.USR_DEFAULT_PLAYERSTATE_MAXIMUMDISTANCE,
		consts.USR_DEFAULT_PLAYERSTATE_DAILYMISSIONID,
		consts.USR_DEFAULT_PLAYERSTATE_DAILYMISSIONENDTIME,
		consts.USR_DEFAULT_PLAYERSTATE_DAILYCHALLENGEVALUE,
		consts.USR_DEFAULT_PLAYERSTATE_DAILYCHALLENGECOMPLETE,
		consts.USR_DEFAULT_PLAYERSTATE_NUMDAILYCHALCONT,
		consts.USR_DEFAULT_PLAYERSTATE_MAINCHARAID,
		consts.USR_DEFAULT_PLAYERSTATE_SUBCHARAID,
		consts.USR_DEFAULT_PLAYERSTATE_MAINCHAOID,
		consts.USR_DEFAULT_PLAYERSTATE_SUBCHAOID,
		consts.USR_DEFAULT_PLAYERSTATE_NUMPLAYING,
		consts.USR_DEFAULT_PLAYERSTATE_NUMANIMALS,
		consts.USR_DEFAULT_PLAYERSTATE_NUMRANK,
		consts.USR_DEFAULT_PLAYERSTATE_EQUIPITEMLIST,
		consts.USR_DEFAULT_PLAYERSTATE_QUICKTOTALHIGHSCORE,
	)
	return ps
}

type PlayerVarious struct {
	CmSkipCount          int64 `json:"cmSkipCount,string"`
	EnergyRecoveryMax    int64 `json:"energyRecoveryMax,string"`
	EnergyRecoveryTime   int64 `json:"energyRecoveryTime,string"`
	OnePlayCmCount       int64 `json:"onePlayCmCount,string"`
	OnePlayContinueCount int64 `json:"onePlayContinueCount,string"`
	IsPurchased          int64 `json:"isPurchased"`
}

//func MakePlayerVarious(cmsc, erm, ert, opcc, opcct, ip string) PlayerVarious {
func MakePlayerVarious(cmsc, erm, ert, opcc, opcct, ip int64) PlayerVarious {
	pv := PlayerVarious{
		cmsc,
		erm,
		ert,
		opcc,
		opcct,
		ip,
	}
	return pv
}

type Chao struct {
	ChaoID   string `json:"chaoId"`
	Status   int64  `json:"status,string"` // consts.CHAO_STATUS_*
	Level    int64  `json:"level"`
	Dealing  int64  `json:"setStatus,string"` // consts.CHAO_DEALING_*
	Acquired int64  `json:"acquired"`
	Rarity   int64  `json:"rarity,string"`
	Hidden   int64  `json:"hidden,string"` // it is likely that if something is an int, they do not need to be stringed, as the game will automatically int any strings
}

func MakeChao(chid string, status, dealing, level, acquired, rarity, hidden int64) Chao {
	chao := Chao{
		chid,
		status,
		level,
		dealing,
		acquired,
		rarity,
		hidden,
	}
	return chao
}

func NewChao(chid string) Chao {
	chao := MakeChao(
		chid,
		consts.USR_DEFAULT_CHAO_STATUS,
		consts.USR_DEFAULT_CHAO_LEVEL,
		consts.USR_DEFAULT_CHAO_DEALING,
		consts.USR_DEFAULT_CHAO_ACQUIRED,
		consts.USR_DEFAULT_CHAO_RARITY,
		consts.USR_DEFAULT_CHAO_HIDDEN,
	)
	return chao
}

type MileageMapState struct {
	Episode          int64 `json:"episode"`
	Chapter          int64 `json:"chapter"`
	Point            int64 `json:"point"`
	MapDistance      int64 `json:"mapDistance"`      // this field is used very sparingly in the game...
	NumBossAttack    int64 `json:"numBossAttack"`    // number of boss fights per this level?
	StageDistance    int64 `json:"stageDistance"`    // how long the stage is?
	StageTotalScore  int64 `json:"stageTotalScore"`  // ?
	StageMaxScore    int64 `json:"stageMaxScore"`    // max score needed to pass?
	ChapterStartTime int64 `json:"chapterStartTime"` // when the chapter starts..?
}

func StartingMileageMapState() MileageMapState {
	// TODO: const the below!
	return MileageMapState{
		1,
		1,
		0,
		0,
		2,
		300,
		10000000,
		10000000,
		time.Now().UTC().Unix(),
	}
}
