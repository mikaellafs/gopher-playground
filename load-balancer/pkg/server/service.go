package server

type Service interface {
	Shutdown(*Server) error
}
