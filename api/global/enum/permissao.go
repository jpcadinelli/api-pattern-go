package enum

var (
	ListaPermissoes = []string{
		PermissaoSistemaAdmin,

		PermissaoPermissaoCriar,
		PermissaoPermissaoVisualizar,
		PermissaoPermissaoListar,
		PermissaoPermissaoAtualizar,
		PermissaoPermissaoDeletar,

		PermissaoUsuarioVisualizar,
		PermissaoUsuarioListar,
		PermissaoUsuarioDropdown,
		PermissaoUsuarioAtualizar,
		PermissaoUsuarioDeletar,

		PermissaoUsuarioAtribuirPermissao,
		PermissaoUsuarioRemoverPermissao,
	}
)

const (
	PermissaoSistemaAdmin = "SISTEMA_ADMIN"

	PermissaoPermissaoCriar      = "PERMISSAO_CRIAR"
	PermissaoPermissaoVisualizar = "PERMISSAO_VISUALIZAR"
	PermissaoPermissaoListar     = "PERMISSAO_LISTAR"
	PermissaoPermissaoAtualizar  = "PERMISSAO_ATUALIZAR"
	PermissaoPermissaoDeletar    = "PERMISSAO_DELETAR"

	PermissaoUsuarioVisualizar = "USUARIO_VISUALIZAR"
	PermissaoUsuarioListar     = "USUARIO_LISTAR"
	PermissaoUsuarioDropdown   = "USUARIO_DROPDOWN"
	PermissaoUsuarioAtualizar  = "USUARIO_ATUALIZAR"
	PermissaoUsuarioDeletar    = "USUARIO_DELETAR"

	PermissaoUsuarioAtribuirPermissao = "USUARIO_ATRIBUIR_PERMISSAO"
	PermissaoUsuarioRemoverPermissao  = "USUARIO_REMOVER_PERMISSAO"
)
