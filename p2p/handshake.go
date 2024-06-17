package p2p

// type Handshaker interface {
// 	Handshake() error
// }

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error {
	return nil
}

// type DefaultHandshaker struct { }
