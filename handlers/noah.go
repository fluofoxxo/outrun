package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/helper"
)

func SetNoahIDHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("HSFDHSDFA HSDF HDSFDFS DFSDFSSDF")
    helper.Respond([]byte("WHOA WHAT"), w)
}
