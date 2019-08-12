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

	"github.com/fluofoxxo/outrun/cryption"
	"github.com/fluofoxxo/outrun/handlers"
)

const UNKNOWN_REQUEST_DIRECTORY = "logging/unknown_requests/"

/*
!!!!!!!!!!
PLEASE NOTE
WHEN THE USER CHOOSES A USERNAME, MAKE SURE TO SET usernameLookup/username = UID
AND MAKE SURE THAT THE USERNAME DOESN'T EXIST
!!!!!!!!!!
*/

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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	log.Println("Starting server on port 9001")
	mux := http.NewServeMux()
	// login and main menu (primarily)
	mux.HandleFunc("/Login/login/", handlers.LoginHandler)
	mux.HandleFunc("/Sgn/sendApollo/", handlers.SendApolloHandler)
	mux.HandleFunc("/Sgn/setNoahId/", handlers.SetNoahIDHandler)
	mux.HandleFunc("/Login/getVariousParameter/", handlers.GetVariousParameterHandler)
	mux.HandleFunc("/Player/getPlayerState/", handlers.GetPlayerStateHandler)
	mux.HandleFunc("/Player/getCharacterState/", handlers.GetCharacterStateHandler)
	mux.HandleFunc("/Player/getChaoState/", handlers.GetChaoStateHandler)
	mux.HandleFunc("/Spin/getWheelOptions/", handlers.GetWheelOptionsHandler)
	mux.HandleFunc("/Game/getDailyChalData/", handlers.GetDailyChalDataHandler)
	mux.HandleFunc("/Message/getMessageList/", handlers.GetMessageListHandler)
	mux.HandleFunc("/Store/getRedstarExchangeList/", handlers.GetRedStarExchangeListHandler)
	mux.HandleFunc("/Game/getCostList/", handlers.GetCostListHandler)
	mux.HandleFunc("/Event/getEventList/", handlers.GetEventListHandler)
	mux.HandleFunc("/Game/getMileageData/", handlers.GetMileageDataHandler)
	mux.HandleFunc("/Game/getCampaignList/", handlers.GetCampaignListHandler)
	mux.HandleFunc("/Chao/getChaoWheelOptions/", handlers.GetChaoWheelOptionsHandler)
	mux.HandleFunc("/Chao/getPrizeChaoWheelSpin/", handlers.GetPrizeChaoWheelSpinHandler)
	mux.HandleFunc("/login/getInformation/", handlers.GetInformationHandler)
	mux.HandleFunc("/Leaderboard/getWeeklyLeaderboardOptions/", handlers.GetWeeklyLeaderboardOptionsHandler)
	mux.HandleFunc("/Leaderboard/getLeagueData/", handlers.GetLeagueData)
	mux.HandleFunc("/Leaderboard/getWeeklyLeaderboardEntries/", handlers.GetWeeklyLeaderboardEntriesHandler)
	mux.HandleFunc("/Player/setUserName/", handlers.SetUserNameHandler)
	mux.HandleFunc("/login/getTicker/", handlers.GetTickerHandler)
	mux.HandleFunc("/Login/loginBonus/", handlers.LoginBonusHandler)
	// Play modes
	mux.HandleFunc("/Game/getFreeItemList/", handlers.GetFreeItemListHandler)
	// Timed mode
	mux.HandleFunc("/Game/quickActStart/", handlers.QuickActStartHandler)
	mux.HandleFunc("/Game/quickPostGameResults/", handlers.QuickPostGameResultsHandler)
	// Campaign mode
	mux.HandleFunc("/Game/actStart/", handlers.ActStartHandler)
	mux.HandleFunc("/Game/postGameResults/", handlers.PostGameResultsHandler)
	mux.HandleFunc("/Game/getMileageReward/", handlers.GetMileageRewardHandler)

	mux.HandleFunc("/", OutputUnknownRequest)
	panic(http.ListenAndServe(":9001", mux))
}
