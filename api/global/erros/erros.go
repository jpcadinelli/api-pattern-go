package erros

import (
	"errors"
	"fmt"
)

var (
	ErrTokenInexistente = errors.New("token inexistente, acesso não autorizado")
	ErrTokenInvalido    = errors.New("token inválido, acesso não autorizado")
)

var (
	ErrUsuarioNaoEncontrado = fmt.Errorf("user not found")
	ErrCredenciaisInvalidas = fmt.Errorf("invalid user credentials")
)
