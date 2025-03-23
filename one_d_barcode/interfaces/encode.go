package interfaces

type Encoder interface {
	Code128Encode(data string) (string, error)
	Code39Encode(data string) (string, error)
	UPCEncode(data string) (string, error)
}
