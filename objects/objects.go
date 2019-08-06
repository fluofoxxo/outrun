package objects

type Message struct { // no idea what this contains?
}

var DEFAULT_REDSTARITEMS_ITEMTYPE_1 = rsiDefaultsIT1()
var DEFAULT_REDSTARITEMS_ITEMTYPE_2 = rsiDefaultsIT2()
var DEFAULT_REDSTARITEMS_ITEMTYPE_4 = rsiDefaultsIT4()

func rsiDefaultsIT1() []RedStarItemSansCampaign {
    // TODO: Make variable names better
    rsiscs := []RedStarItemSansCampaign{}
    storeItemIds := []string{"915001", "915002", "915003", "915004", "915005"}
    itemIds := []string{"910000", "910000", "910000", "910000", "910000"}
    numItems := []int64{20000, 42000, 108000, 224000, 600000}
    prices := []int64{10, 20, 50, 100, 250}
    priceDisps := []string{"", "", "", "", ""}
    productIds := []string{"", "", "", "", ""}
    for i, storeItemId := range storeItemIds {
        itemId := itemIds[i]
        numItem := numItems[i]
        price := prices[i]
        priceDisp := priceDisps[i]
        productId := productIds[i]
        rsisc := MakeRedStarItemSansCampaign(
            storeItemId,
            itemId,
            numItem,
            price,
            priceDisp,
            productId,
        )
        rsiscs = append(rsiscs, rsisc)
    }
    return rsiscs
}

func rsiDefaultsIT2() []RedStarItem {
    rsis := []RedStarItem{}
    storeItemIds := []string{"920005", "920020", "920030", "920050", "920100"}
    itemIds := []string{"920000", "920000", "920000", "920000", "920000"}
    numItems := []int64{5, 20, 30, 50, 100}
    prices := []int64{10, 35, 50, 80, 150}
    priceDisps := []string{"", "", "", "", ""}
    productIds := []string{"", "", "", "", ""}
    for i, storeItemId := range storeItemIds {
        itemId := itemIds[i]
        numItem := numItems[i]
        price := prices[i]
        priceDisp := priceDisps[i]
        productId := productIds[i]
        campaign := NewCampaign(8, 2000, 0) // WARN: is this a possibility for a memory leak?
        rsi := MakeRedStarItem(
            storeItemId,
            itemId,
            numItem,
            price,
            priceDisp,
            productId,
            campaign,
        )
        rsis = append(rsis, rsi)
    }
    return rsis
}

func rsiDefaultsIT4() []RedStarItemSansCampaign {
    rsis := []RedStarItemSansCampaign{}
    storeItemIds := []string{"940005", "940020", "940030", "940050", "940100"}
    itemIds := []string{"940000", "940000", "940000", "940000", "940000"}
    numItems := []int64{5, 20, 30, 50, 100}
    prices := []int64{10, 35, 50, 80, 150}
    priceDisps := []string{"", "", "", "", ""}
    productIds := []string{"", "", "", "", ""}
    for i, storeItemId := range storeItemIds {
        itemId := itemIds[i]
        numItem := numItems[i]
        price := prices[i]
        priceDisp := priceDisps[i]
        productId := productIds[i]
        rsi := MakeRedStarItemSansCampaign(
            storeItemId,
            itemId,
            numItem,
            price,
            priceDisp,
            productId,
        )
        rsis = append(rsis, rsi)
    }
    return rsis
}
