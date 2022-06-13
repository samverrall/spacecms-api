package hasher

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hashed, pwd string) (bool, error)
}
