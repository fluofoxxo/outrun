package objects

type ChaoCampaign struct { // TODO: Is this the same as a normal Campaign? Look in the game code.
}

type ChaoWheelOptions struct {
    Rarity               []int64        `json:"rarity"`
    ItemWeight           []int64        `json:"itemWeight"`
    SpinCost             int64          `json:"spinCost"`
    ChaoRouletteType     int64          `json:"chaoRouletteType"`
    NumSpecialEgg        int64          `json:"numSpecialEgg"`
    RouletteAvailable    int64          `json:"rouletteAvailable"`
    CampaignList         []ChaoCampaign `json:"campaignList"`
    NumChaoRoulette      int64          `json:"numChaoRoulette"`
    NumChaoRouletteToken int64          `json:"numChaoRouletteToken"`
    StartTime            int64          `json:"startTime"`
    EndTime              int64          `json:"endTime"`
}
