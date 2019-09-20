package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fluofoxxo/outrun/cryption"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/netobj/constnetobjs"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

const (
	// TODO: use proper naming conventions
	PREFIX_ERR             = "ERR"
	PREFIX_OUT             = "OUT"
	PREFIX_WARN            = "WARN"
	PREFIX_UNCATCHABLE_ERR = "UNCATCHABLE ERR"

	LOGOUT_BASE = "[%s] (%s) %s\n"
	LOGERR_BASE = "[%s] (%s) %s: %s\n"

	INTERNAL_SRV_ERR = "Internal server error"
	BAD_REQUEST      = "Bad request"

	DEFAULT_IV = "HotAndSunnyMiami"
)

type Helper struct {
	CallerName string
	RespW      http.ResponseWriter
	Request    *http.Request
}

func MakeHelper(callerName string, r http.ResponseWriter, request *http.Request) *Helper {
	return &Helper{
		callerName,
		r,
		request,
	}
}

func (r *Helper) GetGameRequest() []byte {
	recv := cryption.GetReceivedMessage(r.Request)
	return recv
}
func (r *Helper) SendResponse(i interface{}) error {
	out, err := json.Marshal(i)
	if err != nil {
		return err
	}
	r.Respond(out)
	return nil
}
func (r *Helper) SendInsecureResponse(i interface{}) error {
	out, err := json.Marshal(i)
	if err != nil {
		return err
	}
	r.RespondInsecure(out)
	return nil
}
func (r *Helper) RespondRaw(out []byte, secureFlag, iv string) {
	response := map[string]string{}
	if secureFlag != "0" && secureFlag != "1" {
		r.Warn("Improper secureFlag in call to RespondRaw!")
	}
	response["secure"] = secureFlag
	response["key"] = iv
	if secureFlag == "1" {
		encrypted := cryption.Encrypt(out, cryption.EncryptionKey, []byte(iv))
		encryptedBase64 := cryption.B64Encode(encrypted)
		response["param"] = encryptedBase64
	} else {
		response["param"] = string(out)
	}
	toClient, err := json.Marshal(response)
	if err != nil {
		r.InternalErr("Error marshalling in RespondRaw", err)
		return
	}
	r.RespW.Write(toClient)
}
func (r *Helper) Respond(out []byte) {
	r.RespondRaw(out, "1", DEFAULT_IV)
}
func (r *Helper) RespondInsecure(out []byte) {
	r.RespondRaw(out, "0", "")
}
func (r *Helper) Out(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_OUT, r.CallerName, msg)
}
func (r *Helper) Warn(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_WARN, r.CallerName, msg)
}
func (r *Helper) Uncatchable(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_OUT, r.CallerName, msg)
}
func (r *Helper) InternalErr(msg string, err error) {
	log.Printf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Helper) Err(msg string, err error) {
	log.Printf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Helper) ErrRespond(msg string, err error, response string) {
	// TODO: remove if never used in stable builds
	log.Printf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusInternalServerError) // ideally include an option for this, but for now it's inconsequential
	r.RespW.Write([]byte(response))
}
func (r *Helper) InternalFatal(msg string, err error) {
	log.Fatalf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Helper) Fatal(msg string, err error) {
	log.Fatalf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Helper) BaseInfo(em string, statusCode int64) responseobjs.BaseInfo {
	return responseobjs.NewBaseInfo(em, statusCode)
}
func (r *Helper) InvalidRequest() {
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Helper) GetCallingPlayer() (netobj.Player, error) {
	// Powerful function to get the player directly from the response
	recv := r.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(recv, &request)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	sid := request.SessionID
	player, err := db.GetPlayerBySessionID(sid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}
