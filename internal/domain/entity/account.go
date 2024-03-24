package domain

type Account struct {
	Id               uint64
	Username         string
	Email            string
	IsEmailConfirmed bool
	IsBanned         bool
	HashedPassword   string
	Claims           []Claim
}
