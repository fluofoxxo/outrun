package requests

type Base struct {
    SessionID string `json:"sessionId"`
    Version string `json:"version"`
    Seq int64 `json:"seq,string"`
}
