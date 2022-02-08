package generate

type Authenticator struct {
	Secret string
	Expire int
	Code   string
	Error  error
}
