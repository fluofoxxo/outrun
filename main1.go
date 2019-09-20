package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/cryption"
	"github.com/fluofoxxo/outrun/muxhandlers"
	"github.com/fluofoxxo/outrun/muxhandlers/muxobj"
	"github.com/fluofoxxo/outrun/orpc"
	"github.com/gorilla/mux"
)

const UNKNOWN_REQUEST_DIRECTORY = "logging/unknown_requests/"

var (
	LogExecutionTime = true
)

func OutputUnknownRequest(w http.ResponseWriter, r *http.Request) {
	recv := cryption.GetReceivedMessage(r)
	// make a new logging path
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	os.MkdirAll(UNKNOWN_REQUEST_DIRECTORY, 0644)
	normalizedReq := strings.ReplaceAll(r.URL.Path, "/", "-")
	path := UNKNOWN_REQUEST_DIRECTORY + normalizedReq + "_" + timeStr + ".txt"
	err := ioutil.WriteFile(path, recv, 0644)
	if err != nil {
		log.Println("[OUT] UNABLE TO WRITE UNKNOWN REQUEST: " + err.Error())
		w.Write([]byte(""))
		return
	}
	log.Println("[OUT] !!!!!!!!!!!! Unknown request, output to " + path)
	w.Write([]byte(""))
}

func removePrependingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for len(r.URL.Path) != 0 && string(r.URL.Path[0]) == "/" {
			r.URL.Path = r.URL.Path[1:]
		}
		r.URL.Path = "/" + r.URL.Path
		next.ServeHTTP(w, r)
	})
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	err := config.Parse("config.json")
	if err != nil {
		log.Printf("[INFO] No config file (config.json) found (%s), using defaults\n", err)
	} else {
		log.Println("[INFO] Config file (config.json) loaded")
	}

	orpc.Start()

	h := muxobj.Handle
	router := mux.NewRouter()
	router.StrictSlash(true)
	LogExecutionTime = config.CFile.DoTimeLogging
	// Login
	router.HandleFunc("/Login/login/", h(muxhandlers.Login, LogExecutionTime))
	router.HandleFunc("/Sgn/sendApollo/", h(muxhandlers.SendApollo, LogExecutionTime))
	router.HandleFunc("/Login/getVariousParameter/", h(muxhandlers.GetVariousParameter, LogExecutionTime))
	router.HandleFunc("/Player/getPlayerState/", h(muxhandlers.GetPlayerState, LogExecutionTime))
	router.HandleFunc("/Player/getCharacterState/", h(muxhandlers.GetCharacterState, LogExecutionTime))
	router.HandleFunc("/Player/getChaoState/", h(muxhandlers.GetChaoState, LogExecutionTime))
	router.HandleFunc("/Spin/getWheelOptions/", h(muxhandlers.GetWheelOptions, LogExecutionTime))
	router.HandleFunc("/Game/getDailyChalData/", h(muxhandlers.GetDailyChallengeData, LogExecutionTime))
	router.HandleFunc("/Message/getMessageList/", h(muxhandlers.GetMessageList, LogExecutionTime))
	router.HandleFunc("/Store/getRedstarExchangeList/", h(muxhandlers.GetRedStarExchangeList, LogExecutionTime))
	router.HandleFunc("/Game/getCostList/", h(muxhandlers.GetCostList, LogExecutionTime))
	router.HandleFunc("/Event/getEventList/", h(muxhandlers.GetEventList, LogExecutionTime))
	router.HandleFunc("/Game/getMileageData/", h(muxhandlers.GetMileageData, LogExecutionTime))
	router.HandleFunc("/Game/getCampaignList/", h(muxhandlers.GetCampaignList, LogExecutionTime))
	router.HandleFunc("/Chao/getChaoWheelOptions/", h(muxhandlers.GetChaoWheelOptions, LogExecutionTime))
	router.HandleFunc("/Chao/getPrizeChaoWheelSpin/", h(muxhandlers.GetPrizeChaoWheelSpin, LogExecutionTime))
	router.HandleFunc("/login/getInformation/", h(muxhandlers.GetInformation, LogExecutionTime))
	router.HandleFunc("/Leaderboard/getWeeklyLeaderboardOptions/", h(muxhandlers.GetWeeklyLeaderboardOptions, LogExecutionTime))
	router.HandleFunc("/Leaderboard/getLeagueData/", h(muxhandlers.GetLeagueData, LogExecutionTime))
	router.HandleFunc("/Leaderboard/getWeeklyLeaderboardEntries/", h(muxhandlers.GetWeeklyLeaderboardEntries, LogExecutionTime))
	router.HandleFunc("/Player/setUserName/", h(muxhandlers.SetUsername, LogExecutionTime))
	router.HandleFunc("/login/getTicker/", h(muxhandlers.GetTicker, LogExecutionTime))
	router.HandleFunc("/Login/loginBonus/", h(muxhandlers.LoginBonus, LogExecutionTime))
	// Timed mode
	router.HandleFunc("/Game/quickActStart/", h(muxhandlers.QuickActStart, LogExecutionTime))
	router.HandleFunc("/Game/quickPostGameResults/", h(muxhandlers.QuickPostGameResults, LogExecutionTime))
	// Story mode
	router.HandleFunc("/Game/actStart/", h(muxhandlers.ActStart, LogExecutionTime))
	router.HandleFunc("/Game/getMileageReward/", h(muxhandlers.GetMileageReward, LogExecutionTime))
	// Retry
	router.HandleFunc("/Game/actRetry/", h(muxhandlers.ActRetry, LogExecutionTime))
	// Gameplay
	router.HandleFunc("/Game/getFreeItemList/", h(muxhandlers.GetFreeItemList, LogExecutionTime))
	router.HandleFunc("/Game/postGameResults/", h(muxhandlers.PostGameResults, LogExecutionTime))
	// Misc.
	router.HandleFunc("/Character/changeCharacter/", h(muxhandlers.ChangeCharacter, LogExecutionTime))
	router.HandleFunc("/Character/upgradeCharacter/", h(muxhandlers.UpgradeCharacter, LogExecutionTime))
	router.HandleFunc("/Chao/equipChao/", h(muxhandlers.EquipChao, LogExecutionTime))
	// Shop
	router.HandleFunc("/Store/redstarExchange/", h(muxhandlers.RedStarExchange, LogExecutionTime))
	// Battle
	router.HandleFunc("/Battle/getDailyBattleData/", h(muxhandlers.GetDailyBattleData, LogExecutionTime))
	router.HandleFunc("/Battle/updateDailyBattleStatus/", h(muxhandlers.UpdateDailyBattleStatus, LogExecutionTime))

	if config.CFile.LogUnknownRequests {
		//router.HandleFunc("/", OutputUnknownRequest)
		router.PathPrefix("/").HandlerFunc(OutputUnknownRequest)
	}

	port := config.CFile.Port
	log.Printf("Starting server on port %s\n", port)
	panic(http.ListenAndServe(":"+port, removePrependingSlashes(router)))
}
