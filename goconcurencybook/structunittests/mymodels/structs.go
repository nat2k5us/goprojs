package mymodels

import (
	"errors"
	"fmt"
)

type my struct {
	Branch bool
}

func (s *my) operations() error {

	fmt.Println("my struct operations")
	if s.Branch {
		fmt.Println("branching code")
		return errors.New("bad branch")
	}
	return nil
}
