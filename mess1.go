package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/fluofoxxo/outrun/handlers"
)

/*
!!!!!!!!!!
PLEASE NOTE
WHEN THE USER CHOOSES A USERNAME, MAKE SURE TO SET usernameLookup/username = UID
AND MAKE SURE THAT THE USERNAME DOESN'T EXIST
!!!!!!!!!!
*/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	log.Println("Starting server on port 9001")
	mux := http.NewServeMux()
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
	panic(http.ListenAndServe(":9001", mux))
}
