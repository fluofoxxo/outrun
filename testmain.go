package main

import (
	"fmt"

	_ "github.com/fluofoxxo/outrun/bgtasks"
	_ "github.com/fluofoxxo/outrun/consts"
	_ "github.com/fluofoxxo/outrun/cryption"
	_ "github.com/fluofoxxo/outrun/db"
	_ "github.com/fluofoxxo/outrun/db/dbaccess"
	_ "github.com/fluofoxxo/outrun/enums"
	_ "github.com/fluofoxxo/outrun/helper"
	_ "github.com/fluofoxxo/outrun/log"
	_ "github.com/fluofoxxo/outrun/muxhandlers"
	_ "github.com/fluofoxxo/outrun/muxhandlers/muxobj"
	_ "github.com/fluofoxxo/outrun/netobj"
	_ "github.com/fluofoxxo/outrun/obj"
	_ "github.com/fluofoxxo/outrun/obj/constobjs"
	_ "github.com/fluofoxxo/outrun/requests"
	_ "github.com/fluofoxxo/outrun/responses"
)

func main() {
	fmt.Println("We're good...")
}
