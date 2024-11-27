CREATE TABLE teste (
    id uuid NOT NULL,
    nome text NOT NULL,
    CONSTRAINT pk_teste PRIMARY KEY (id)
);

COMMENT ON COLUMN teste.id IS 'Identificador único da tabela teste';
COMMENT ON COLUMN teste.nome IS 'Nome associado à tabela teste';

CREATE TABLE users (
    id UUID NOT NULL,
    primeiro_nome TEXT NOT NULL,
    ultimo_nome TEXT NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    data_nascimento DATE NOT NULL,
    created_at TIMESTAMP,
        CONSTRAINT pk_user PRIMARY KEY (id)
);

COMMENT ON COLUMN users.id IS 'Identificador único do usuário';
COMMENT ON COLUMN users.primeiro_nome IS 'Primeiro nome do usuário';
COMMENT ON COLUMN users.ultimo_nome IS 'Último nome do usuário';
COMMENT ON COLUMN users.cpf IS 'CPF do usuário (único)';
COMMENT ON COLUMN users.email IS 'Email do usuário (único)';
COMMENT ON COLUMN users.password IS 'Senha do usuário (armazenada de forma segura)';
COMMENT ON COLUMN users.data_nascimento IS 'Data de nascimento do usuário';
COMMENT ON COLUMN users.created_at IS 'Data de criação do registro';
