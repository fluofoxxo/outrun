package netobj

type MileageFriend struct {
    ID              string          `json:"friendId"`
    Name            string          `json:"name"`
    URL             string          `json:"url"`
    MileageMapState MileageMapState `json:"mapState"`
}
