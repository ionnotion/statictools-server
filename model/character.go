package model

type Character struct {
	Id            uint
	FirstName     string
	LastName      string
	DataCenter    DataCenter
	CharacterJobs []CharacterJob
	IsGuest       bool
}
