package constobjs

import (
    "strconv"

    "github.com/fluofoxxo/outrun/obj"
)

var DefaultPlayerStateItems = func() []obj.Item {
    items := []obj.Item{}
    baseNum := 120000
    for i := range make([]byte, 11) { // 0...10
        n := baseNum + i
        s := strconv.Itoa(n)
        item := obj.NewItem(s, 0)
        items = append(items, item)
    }
    return items
}()
