CREATE TABLE teste (
    id uuid NOT NULL,
    nome text NOT NULL,
    CONSTRAINT pk_teste PRIMARY KEY (id)
);

COMMENT ON COLUMN teste.id IS 'Identificador único da tabela teste';
COMMENT ON COLUMN teste.nome IS 'Nome associado à tabela teste';

CREATE TABLE usuario (
    id UUID NOT NULL,
    primeiro_nome TEXT NOT NULL,
    ultimo_nome TEXT NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    data_nascimento DATE NOT NULL,
    created_at TIMESTAMP,
        CONSTRAINT pk_usuario PRIMARY KEY (id)
);

COMMENT ON COLUMN usuario.id IS 'Identificador único do usuário';
COMMENT ON COLUMN usuario.primeiro_nome IS 'Primeiro nome do usuário';
COMMENT ON COLUMN usuario.ultimo_nome IS 'Último nome do usuário';
COMMENT ON COLUMN usuario.cpf IS 'CPF do usuário (único)';
COMMENT ON COLUMN usuario.email IS 'Email do usuário (único)';
COMMENT ON COLUMN usuario.password IS 'Senha do usuário (armazenada de forma segura)';
COMMENT ON COLUMN usuario.data_nascimento IS 'Data de nascimento do usuário';
COMMENT ON COLUMN usuario.created_at IS 'Data de criação do registro';

CREATE TABLE permissao (
    id UUID,
    nome VARCHAR(50) UNIQUE,
    descricao TEXT,
        CONSTRAINT pk_permissao PRIMARY KEY (id)
);

COMMENT ON COLUMN permissao.id IS 'Identificador único da tabela permissão';
COMMENT ON COLUMN permissao.nome IS 'Nome associado à tabela permissão';
COMMENT ON COLUMN permissao.descricao IS 'Descrição associada à tabela permissão';

CREATE TABLE permissao_usuario (
    id UUID,
    id_permissao UUID,
    id_usuario UUID,
        CONSTRAINT pk_permissao_usuario PRIMARY KEY (id),
        CONSTRAINT fk_permissao_usuario_permissao FOREIGN KEY (id_permissao) REFERENCES usuario(id),
        CONSTRAINT fk_permissao_usuario_usuario FOREIGN KEY (id_usuario) REFERENCES permissao(id)
);

COMMENT ON COLUMN permissao_usuario.id IS 'Identificador único da tabela permissao_usuario';
COMMENT ON COLUMN permissao_usuario.id_permissao IS 'Referência para a permissão na tabela permissao';
COMMENT ON COLUMN permissao_usuario.id_usuario IS 'Referência para o usuário na tabela usuario';
