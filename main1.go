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
	"github.com/fluofoxxo/outrun/muxhandlers"
	"github.com/fluofoxxo/outrun/muxhandlers/muxobj"
)

const UNKNOWN_REQUEST_DIRECTORY = "logging/unknown_requests/"

const (
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	h := muxobj.Handle
	mux := http.NewServeMux()
	// Login
	mux.HandleFunc("/Login/login/", h(muxhandlers.Login, LogExecutionTime))
	mux.HandleFunc("/Sgn/sendApollo/", h(muxhandlers.SendApollo, LogExecutionTime))
	mux.HandleFunc("/Login/getVariousParameter/", h(muxhandlers.GetVariousParameter, LogExecutionTime))
	mux.HandleFunc("/Player/getPlayerState/", h(muxhandlers.GetPlayerState, LogExecutionTime))
	mux.HandleFunc("/Player/getCharacterState/", h(muxhandlers.GetCharacterState, LogExecutionTime))
	mux.HandleFunc("/Player/getChaoState/", h(muxhandlers.GetChaoState, LogExecutionTime))
	mux.HandleFunc("/Spin/getWheelOptions/", h(muxhandlers.GetWheelOptions, LogExecutionTime))
	mux.HandleFunc("/Game/getDailyChalData/", h(muxhandlers.GetDailyChallengeData, LogExecutionTime))
	mux.HandleFunc("/Message/getMessageList/", h(muxhandlers.GetMessageList, LogExecutionTime))
	mux.HandleFunc("/Store/getRedstarExchangeList/", h(muxhandlers.GetRedStarExchangeList, LogExecutionTime))
	mux.HandleFunc("/Game/getCostList/", h(muxhandlers.GetCostList, LogExecutionTime))
	mux.HandleFunc("/Event/getEventList/", h(muxhandlers.GetEventList, LogExecutionTime))
	mux.HandleFunc("/Game/getMileageData/", h(muxhandlers.GetMileageData, LogExecutionTime))
	mux.HandleFunc("/Game/getCampaignList/", h(muxhandlers.GetCampaignList, LogExecutionTime))
	mux.HandleFunc("/Chao/getChaoWheelOptions/", h(muxhandlers.GetChaoWheelOptions, LogExecutionTime))
	mux.HandleFunc("/Chao/getPrizeChaoWheelSpin/", h(muxhandlers.GetPrizeChaoWheelSpin, LogExecutionTime))
	mux.HandleFunc("/login/getInformation/", h(muxhandlers.GetInformation, LogExecutionTime))
	mux.HandleFunc("/Leaderboard/getWeeklyLeaderboardOptions/", h(muxhandlers.GetWeeklyLeaderboardOptions, LogExecutionTime))
	mux.HandleFunc("/Leaderboard/getLeagueData/", h(muxhandlers.GetLeagueData, LogExecutionTime))
	mux.HandleFunc("/Leaderboard/getWeeklyLeaderboardEntries/", h(muxhandlers.GetWeeklyLeaderboardEntries, LogExecutionTime))
	mux.HandleFunc("/Player/setUserName/", h(muxhandlers.SetUsername, LogExecutionTime))
	mux.HandleFunc("/login/getTicker/", h(muxhandlers.GetTicker, LogExecutionTime))
	mux.HandleFunc("/Login/loginBonus/", h(muxhandlers.LoginBonus, LogExecutionTime))
	// Timed mode
	mux.HandleFunc("/Game/quickActStart/", h(muxhandlers.QuickActStart, LogExecutionTime))
	// Retry
	mux.HandleFunc("/Game/actRetry/", h(muxhandlers.ActRetry, LogExecutionTime))

	mux.HandleFunc("/", OutputUnknownRequest)
	log.Println("Starting server on port 9001")
	panic(http.ListenAndServe(":9001", mux))
}
