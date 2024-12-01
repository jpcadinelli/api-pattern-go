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
	ErrUsuarioNaoEncontrado              = fmt.Errorf("usuário não encontrado")
	ErrCredenciaisInvalidas              = fmt.Errorf("credenciais inválidas do usuário")
	ErrNaoPodeMudadarDadosDeOutroUsuario = fmt.Errorf("seu usuário não tem permissão de mudar os dados de outro usuário")
)

var (
	ErrPermissaoNaoEncontrada = fmt.Errorf("permissão não encontrado")
)

var (
	ErrUsuarioNaoTemPermissao = fmt.Errorf("usuário não tem permissão")
)
