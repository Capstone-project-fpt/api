package queue

type IBasePublisher[T any] interface {
	SendMessage(message T, delay int) error
}

