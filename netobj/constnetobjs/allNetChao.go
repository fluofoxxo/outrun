package constnetobjs

import (
	"strconv"

	"github.com/fluofoxxo/outrun/config/gameconf"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
)

var ChaoIDs = []string{"400000", "400001", "400002", "400003", "400004", "400005", "400006", "400007", "400008", "400009", "400010", "400011", "400012", "400013", "400014", "400015", "400016", "400017", "400018", "400019", "400020", "400021", "400022", "400023", "400024", "400025", "401000", "401001", "401002", "401003", "401004", "401005", "401006", "401007", "401008", "401009", "401010", "401011", "401012", "401013", "401014", "401015", "401016", "401017", "401018", "401019", "401020", "401021", "401022", "401023", "401024", "401025", "401026", "401027", "401028", "401029", "401030", "401031", "401032", "401033", "401034", "401035", "401036", "401037", "401038", "401039", "401040", "401041", "401042", "401043", "401044", "401045", "401046", "401047", "402000", "402001", "402002", "402003", "402004", "402005", "402006", "402007", "402008", "402009", "402010", "402011", "402012", "402013", "402014", "402015", "402016", "402017", "402018", "402019", "402020", "402021", "402022", "402023", "402024", "402025", "402026", "402027", "402028", "402029", "402030", "402031", "402032", "402033", "402034"}

var NetChao = GetAllNetChao()
var NetChaoList = GetAllNetChaoList()

func DefaultChaoState() []netobj.Chao {
	chaos := []netobj.Chao{}
	chaoStatus := int64(enums.ChaoStatusNotOwned)
	chaoLevel := int64(0)
	chaoDealing := int64(enums.ChaoDealingNone) // TODO: discover use
	acquired := int64(0)
	if gameconf.CFile.AllChaoUnlocked {
		chaoStatus = enums.ChaoStatusOwned
		acquired = int64(1)
	}
	for _, chaoID := range ChaoIDs {
		id := chaoID
		rarity, _ := strconv.Atoi(string(id[2])) // numerical rarity (third digit)
		hidden := int64(0)                       // TODO: discover what this is used for (see obj/chao.go)
		chao := obj.NewChao(
			id,
			int64(rarity),
			hidden,
		)
		netchao := netobj.NewNetChao(
			chao,
			chaoStatus,
			chaoLevel,
			chaoDealing,
			acquired,
		)
		chaos = append(chaos, netchao)
	}
	return chaos
}

func GetAllNetChao() map[string]netobj.Chao {
	// TODO: remove. Should not be used anymore.
	chaos := make(map[string]netobj.Chao)
	for _, chaoID := range ChaoIDs {
		id := chaoID
		rarity, _ := strconv.Atoi(string(id[2])) // numerical rarity (third digit)
		hidden := int64(0)                       // TODO: discover what this is used for (see obj/chao.go)
		chao := obj.NewChao(
			id,
			int64(rarity),
			hidden,
		)
		netchao := netobj.NewNetChao(
			chao,
			enums.ChaoStatusOwned,
			1,
			enums.ChaoDealingNone,
			1,
		)
		chaos[chaoID] = netchao
	}
	return chaos
}

func GetAllNetChaoList() []netobj.Chao {
	chaolist := []netobj.Chao{}
	for _, value := range GetAllNetChao() {
		chaolist = append(chaolist, value)
	}
	return chaolist
}
