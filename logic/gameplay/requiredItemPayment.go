package gameplay

import (
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/obj/constobjs"
)

func findItem(id string) obj.ConsumedItem {
    var result obj.ConsumedItem
    for _, citem := range constobjs.DefaultCostList {
        if citem.ID == id {
            result = citem
            break
        }
    }
    return result
}

func GetRequiredItemPayment(items []string) int64 {
    totalRingPayment := int64(0)
    for _, itemID := range items {
        citem := findItem(itemID)
        totalRingPayment += citem.Item.Amount
    }
    return totalRingPayment
}
