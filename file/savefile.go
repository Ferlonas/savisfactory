package file

import (
	"fmt"
	"time"
)

type SessionInfo struct {
	LevelType string
	// UnknownID1 is bound to the game
	UnknownID1 uint32
	// StartLocation is terminated by 0x3F
	StartLocation string
	// Name is terminated by 0x3F
	Name string
	// Visibility is NULL terminated
	Visibility string
	// UnknownID2 is bound to the game
	UnknownID2 uint32
	// SecondName is NULL terminated and the same as Name
	SecondName string
	// Age is 4 bytes of left-aligned uint in little endian and needs to be parsed
	Age time.Duration
}

type FileInfo struct {
	Session SessionInfo
}

type SaveFile struct {
	Info FileInfo
}

func (f *SaveFile) Dump() {
	fmt.Printf("Satisfactory save file: \n")
}
