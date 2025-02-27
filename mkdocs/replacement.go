package main

type ReplacementType int

const (
	Simple ReplacementType = iota
	Regexp
)

type Replacement struct {
	Type ReplacementType
	From string
	To   string
}
