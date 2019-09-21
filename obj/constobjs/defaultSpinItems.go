package constobjs

import (
    "strconv"

    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/obj"
)

var DefaultSpinItems = []obj.Item{
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 1),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 2),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 3),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 4),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 5),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 6),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 7),
    obj.NewItem(strconv.Itoa(int(enums.ItemIDAsteroid)), 8),
}
