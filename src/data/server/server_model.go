package server

type ServerQuery struct {
	ServerId     string
	Station      uint8
	Service_Time uint32
}

type ServerLogin struct {
	ServerId string
	Passowrd string
}
