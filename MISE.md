# Mise Configuration for TaskHub

Este projeto usa o [mise](https://mise.jdx.dev/) para gerenciar dependências e tarefas.

## Instalação

1. Instale o mise se ainda não tiver:
   ```bash
   curl https://mise.run | sh
   ```

2. Confie no arquivo de configuração do projeto:
   ```bash
   mise trust
   ```

3. Instale as ferramentas necessárias:
   ```bash
   mise install
   ```

## Configuração

O arquivo `.mise.toml` contém:

- **Ferramentas**: Go 1.23
- **Variáveis de ambiente**: Configurações do banco de dados e aplicação
- **Tasks**: Comandos úteis para desenvolvimento

## Tasks Disponíveis

Execute `mise tasks` para ver todos os tasks disponíveis:

```bash
# Construir a aplicação
mise run build

# Executar a aplicação
mise run run

# Executar testes
mise run test

# Formatar código Go
mise run fmt

# Limpar módulos Go
mise run mod-tidy

# Iniciar containers Docker
mise run docker-up

# Parar containers Docker
mise run docker-down
```

## Variáveis de Ambiente

O mise configura automaticamente as seguintes variáveis:

- `DB_HOST=localhost`
- `DB_PORT=5432`
- `DB_USER=postgres`
- `DB_PASSWORD=password`
- `DB_NAME=taskhub`
- `DB_SSLMODE=disable`
- `APP_PORT=8080`
- `GO111MODULE=on`
- `CGO_ENABLED=1`

## Comandos Úteis

```bash
# Ver versões atuais das ferramentas
mise current

# Ver variáveis de ambiente configuradas
mise env

# Ativar o ambiente no shell atual
eval "$(mise env)"

# Executar um comando com o ambiente mise
mise exec -- go version
```

## Integração com IDE

Para integrar com seu editor/IDE, você pode:

1. **VS Code**: Use a extensão "mise" ou configure o terminal integrado
2. **Outros IDEs**: Execute `mise env` para ver as variáveis que precisam ser configuradas

## Troubleshooting

Se encontrar problemas:

1. Verifique se o mise está instalado: `mise --version`
2. Confie no projeto: `mise trust`
3. Reinstale as ferramentas: `mise install --force`
4. Verifique o status: `mise doctor`
