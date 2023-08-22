package message

type Store interface {
	SaveMessage(Message) error
}
