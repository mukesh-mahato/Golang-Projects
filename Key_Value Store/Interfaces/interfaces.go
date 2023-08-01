package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

func Decode(b []byte) error {
	return nil
}

func DecodeReader(r io.Reader) error {
	return nil
}

func main() {

	data := []byte("fooo!")

	Decode(data)

	DecodeReader(bytes.NewReader(data))

	DoAttack(StrongAttacker{})
	DoAttack(SuperStrongAttacker{})
}

type Attacker interface {
	Attack() error
}

type StrongAttacker struct{}

func (sa StrongAttacker) Attack() error {
	fmt.Println("Woow, What a strong attack!")

	return nil
}

type SuperStrongAttacker struct{}

func (ssa SuperStrongAttacker) Attack() error {
	fmt.Println("Fcuk, What a super strong attacker")

	return nil
}

func DoAttack(a Attacker) error {
	return a.Attack()
}
