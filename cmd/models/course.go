package main

import "github.com/google/uuid"

type Course struct {
	Id       uuid.UUID
	Name     string
	Location string
	Holes    int
}
