package objects

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type ConsumedItem struct {
    ConsumedItemID string `json:"consumedItemId"`
    pdata.ItemInfo
}

/*
type ConsumedCostList struct {
    ConsumedItems []ConsumedItem `json:"consumedCostList"`
}
*/

var DEFAULT_CONSUMEDCOSTLIST = getDefaultCostList()

func getDefaultCostList() []ConsumedItem {
    result := []ConsumedItem{}
    cids1 := []string{"110000", "110001", "110002"}
    cids1IDs := []string{"910000", "910000", "910000"}
    cids1Num := []int64{6000, 1000, 4000}
    cids2 := []string{"120000", "120001", "120002", "120003", "120004", "120005", "120006", "120007", "120008", "120009", "120010"}
    cids2IDs := []string{"910000", "910000", "910000", "910000", "910000", "910000", "910000", "910000", "910000", "910000", "910000"}
    cids2Num := []int64{3000, 1000, 3000, 2000, 3000, 5000, 4000, 5000, 0, 0, 0}
    cids3 := []string{"950000", "980000", "980001", "980002"}
    cids3IDs := []string{"900000", "900000", "900000", "900000"}
    cids3Num := []int64{5, 2, 5, 10}
    cids := []string{}
    cids = append(cids, cids1...)
    cids = append(cids, cids2...)
    cids = append(cids, cids3...)
    cidsIDs := []string{}
    cidsIDs = append(cidsIDs, cids1IDs...)
    cidsIDs = append(cidsIDs, cids2IDs...)
    cidsIDs = append(cidsIDs, cids3IDs...)
    cidsNum := []int64{}
    cidsNum = append(cidsNum, cids1Num...)
    cidsNum = append(cidsNum, cids2Num...)
    cidsNum = append(cidsNum, cids3Num...)
    for i, cid := range cids {
        id := cidsIDs[i]
        num := cidsNum[i]
        iinfo := pdata.MakeItemInfo(
            id,
            num,
        )
        ci := ConsumedItem{
            cid,
            iinfo,
        }
        result = append(result, ci)
    }
    return result
}
