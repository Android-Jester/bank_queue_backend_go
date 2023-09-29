package client

import "time"

type ClientQueueData struct {
	Name             string
	NationalId       string
	Activity         string
	Position         int32
	SubQueuePosition int32
	AssignedServer   string
	ServerLocation   int8
	TimeDuration     uint32
	TimeJoined       time.Time
}

type ClientInputData struct {
	NationalId string
	Activity   string
}

type ClientDataQuery struct {
	Name          string
	NationalId    string
	AccountNumber string
}

// func (clientData *ClientQueueData) New() {
// 	clientData = &ClientQueueData{
// 	}
// }
