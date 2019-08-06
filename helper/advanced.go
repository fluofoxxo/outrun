package helper

import (
    "encoding/json"
    "net/http"

    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/requests"
)

func GetBasicRequest(r *http.Request) (requests.BasicRequest, error) {
    recv := cryption.GetReceivedMessage(r)
    var request requests.BasicRequest
    err := json.Unmarshal(recv, &request)
    return request, err
}
