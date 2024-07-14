package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	charLimit := 100
	if membershipType == TypePremium {
		charLimit = 1000
	}
	user := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: charLimit,
		},
	}
	return user
}
