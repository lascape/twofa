package generate

type Authenticator struct {
	Secret string
	Expire int
	Code   int
}
