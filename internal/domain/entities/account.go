package entities

type Account struct {
	Id             uint64
	Username       string
	Email          string
	HashedPassword string
	Claims         []Claim
}
