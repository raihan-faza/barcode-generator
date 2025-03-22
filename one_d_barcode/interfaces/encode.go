package interfaces

type Encoder interface {
	Code128Encode(data string) string
	Code39Encode(data string) string
	UPCEncode(data string) string
}
