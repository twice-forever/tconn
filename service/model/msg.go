package model

type SendMsgJson struct {
	ID       int32    `json:"id"`
	Type     uint8    `json:"type"`
	Position Position `json:"positon"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
