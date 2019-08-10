package objects

import (
	"github.com/fluofoxxo/outrun/playerdata"
)

type Friend struct { // TODO: discover fields
}

type OperatorMessage struct { // is this the same as a normal message?
	MessageID  string      `json:"messageId"`
	Contents   string      `json:"contents"`
	Item       MessageItem `json:"item"`
	ExpireTime int64       `json:"expireTime"`
}

func NewOperatorMessage(mid, contents string, item MessageItem, expireTime int64) OperatorMessage {
	return OperatorMessage{
		mid,
		contents,
		item,
		expireTime,
	}
}

type MessageItem struct {
	ItemID          string `json:"itemId"`
	NumItem         int64  `json:"numItem"`
	AdditionalInfo1 int64  `json:"additionalInfo1"`
	AdditionalInfo2 int64  `json:"additionalInfo2"`
}

func NewMessageItem(iid string, ni, ai1, ai2 int64) MessageItem {
	return MessageItem{
		iid,
		ni,
		ai1,
		ai2,
	}
}

type PlayCharacterState struct {
	pdata.CharacterState
	AbilityLevelUp    []int64 `json:"abilityLevelup"`
	AbilityLevelUpExp []int64 `json:"abilityLevelupExp`
}

func NewPlayCharacterState(cs pdata.CharacterState, alu, alue []int64) PlayCharacterState {
	return PlayCharacterState{
		cs,
		alu,
		alue,
	}
}
