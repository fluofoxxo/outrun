package constobjs

import "github.com/fluofoxxo/outrun/obj"

var ItemIDs = []string{"110000", "110001", "110002", "120000", "120001", "120002", "120003", "120004", "120005", "120006", "120007", "120008", "120009", "120010"}
var AllItems = getAllItems()

func getAllItems() []obj.Item {
	items := []obj.Item{}
	for _, iid := range ItemIDs {
		item := obj.NewItem(iid, 1)
		items = append(items, item)
	}
	return items
}
