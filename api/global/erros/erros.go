package erros

import "fmt"

var (
	ErrUsuarioNaoEncontrado = fmt.Errorf("user not found")
	ErrCredenciaisInvalidas = fmt.Errorf("invalid user credentials")
)
