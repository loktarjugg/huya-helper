package jwt

type HuYaJWT interface {
	GetJWTToken (data interface{}) (string, error)
}
