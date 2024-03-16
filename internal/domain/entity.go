package domain

type Account struct {
	Id               uint64
	Username         string
	Email            string
	IsEmailConfirmed bool
	HashedPassword   string
	Claims           []Claim
}

type Claim struct {
	Id    uint64
	Title string
	Value string
}
