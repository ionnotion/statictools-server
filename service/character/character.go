package character_service

type CreateNewCharacterSpec struct {
	FirstName string
	LastName  string
	IsGuest   bool
}
