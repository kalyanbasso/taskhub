# TaskHub API

Uma API REST para gerenciamento de tarefas (TaskHub) construída e## Como Executar

### Pré-requisitos
- Docker e Docker Compose
- Go 1.23+ (para desenvolvimento local)
- jq (para testes com formatação JSON)
- [mise](https://mise.jdx.dev/) (opcional, para gerenciamento de dependências)

### Usando mise (Recomendado para Desenvolvimento)

Se você usa mise para gerenciar dependências:

1. Clone o repositório:
```bash
git clone <repository-url>
cd taskhub
```

2. Confie no projeto e instale dependências:
```bash
mise trust
mise install
```

3. Execute a aplicação:
```bash
mise run docker-up  # Inicia containers
mise run run         # Executa a aplicação localmente
```

4. Outros comandos úteis:
```bash
mise tasks           # Lista todos os tasks disponíveis
mise run test        # Executa testes
mise run fmt         # Formata código Go
mise run build       # Constrói a aplicação
```

> Para mais detalhes sobre o mise, veja [MISE.md](./MISE.md)

### Com Docker Compose

1. Clone o repositório:
```bash
git clone <repository-url>
cd taskhub
```

2. Execute com Docker Compose:
```bash
docker-compose up --build -d
```

3. A API estará disponível em `http://localhost:8080` PostgreSQL.

## Características

- ✅ CRUD completo para tarefas
- ✅ Filtragem por status (completo/incompleto)
- ✅ Filtragem por prioridade (baixa/média/alta)
- ✅ Listagem de tarefas vencidas
- ✅ Validação de dados
- ✅ Containerização com Docker
- ✅ Banco de dados PostgreSQL

## Estrutura do Projeto

```
.
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── config/                 # Configurações e variáveis de ambiente
│   ├── controller/             # Controllers HTTP (handlers)
│   ├── model/                  # Modelos de dados e structs
│   ├── repository/             # Interações com banco de dados
│   └── usecase/               # Regras de negócio
├── docker-compose.yml         # Configuração do Docker Compose
├── Dockerfile                 # Configuração do container da aplicação
└── README.md
```

## API Endpoints

### Health Check
- `GET /health` - Verificar status da API

### Tasks
- `POST /api/v1/tasks` - Criar nova tarefa
- `GET /api/v1/tasks` - Listar todas as tarefas
- `GET /api/v1/tasks/:id` - Obter tarefa por ID
- `PUT /api/v1/tasks/:id` - Atualizar tarefa
- `DELETE /api/v1/tasks/:id` - Deletar tarefa
- `PATCH /api/v1/tasks/:id/complete` - Marcar tarefa como completa
- `GET /api/v1/tasks/status/:completed` - Filtrar por status (true/false)
- `GET /api/v1/tasks/priority/:priority` - Filtrar por prioridade (low/medium/high)
- `GET /api/v1/tasks/overdue` - Listar tarefas vencidas

## Modelos de Dados

### Task
```json
{
  "id": 1,
  "title": "Título da tarefa",
  "description": "Descrição opcional",
  "completed": false,
  "priority": "medium",
  "due_date": "2023-12-31T23:59:59Z",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Criar Tarefa
```json
{
  "title": "Título da tarefa",
  "description": "Descrição opcional",
  "priority": "high",
  "due_date": "2023-12-31T23:59:59Z"
}
```

### Atualizar Tarefa
```json
{
  "title": "Novo título",
  "description": "Nova descrição",
  "completed": true,
  "priority": "low",
  "due_date": "2023-12-31T23:59:59Z"
}
```

## Como Executar

### Pré-requisitos
- Docker e Docker Compose
- Go 1.21+ (para desenvolvimento local)

### Com Docker Compose (Recomendado)

1. Clone o repositório:
```bash
git clone <repository-url>
cd taskhub
```

2. Execute com Docker Compose:
```bash
docker-compose up --build
```

3. A API estará disponível em `http://localhost:8080`

### Desenvolvimento Local

1. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env conforme necessário
```

2. Execute o PostgreSQL:
```bash
docker run --name postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=taskhub -p 5432:5432 -d postgres:15-alpine
```

3. Instale as dependências:
```bash
go mod download
```

4. Execute a aplicação:
```bash
go run cmd/main.go
```

## Exemplos de Uso

### Criar uma nova tarefa
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Estudar Go",
    "description": "Aprender sobre Gin e GORM",
    "priority": "high",
    "due_date": "2023-12-31T23:59:59Z"
  }'
```

### Listar todas as tarefas
```bash
curl http://localhost:8080/api/v1/tasks
```

### Marcar tarefa como completa
```bash
curl -X PATCH http://localhost:8080/api/v1/tasks/1/complete
```

### Filtrar tarefas por prioridade
```bash
curl http://localhost:8080/api/v1/tasks/priority/high
```

## Tecnologias Utilizadas

- **Go** - Linguagem de programação
- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Docker** - Containerização
- **Docker Compose** - Orquestração de containers

## Arquitetura

O projeto segue os princípios de Clean Architecture:

- **Controller**: Responsável por receber e processar requisições HTTP
- **UseCase**: Contém a lógica de negócio da aplicação
- **Repository**: Abstrai o acesso aos dados
- **Model**: Define as estruturas de dados
- **Config**: Gerencia configurações e variáveis de ambiente

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request
