package main

type ResiValidator interface {
	Validate(resi string) bool
}

type RealResiValidator struct{}

func (v RealResiValidator) Validate(resi string) bool {

	if len(resi) < 5 {
		return false
	}

	return true
}