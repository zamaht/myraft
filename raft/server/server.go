package server

const SERVER_TYPE = "tcp"

type Server interface {
	Run(address string, port string)
	Stop()

	init()
	listenOn(address string, port string)
	applyRules()
	handleChannelsMsg()
}

type ServerRole int

const (
	Follower  ServerRole = iota
	Candidate ServerRole = iota
	Leader    ServerRole = iota
)
