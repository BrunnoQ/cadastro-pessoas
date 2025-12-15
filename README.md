# Cadastro de Pessoas - API REST

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![MongoDB](https://img.shields.io/badge/MongoDB-6.0+-47A248?style=flat&logo=mongodb)](https://www.mongodb.com)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Sistema completo de gerenciamento de pessoas desenvolvido em Go, seguindo princípios de **Clean Code** e **Clean Architecture**.

## ✨ Funcionalidades

- ✅ **CRUD Completo** - Criar, listar, buscar, atualizar e deletar pessoas
- ✅ **Gestão de Contatos** - Múltiplos endereços e meios de contato por pessoa
- ✅ **Paginação** - Lista com suporte a paginação (offset-based, limite max 100)
- ✅ **Validação Robusta** - Email, telefone, datas, nomes com validators customizados
- ✅ **Optimistic Locking** - Controle de concorrência com campo version
- ✅ **Logging Estruturado** - Logs detalhados com Zap
- ✅ **Error Handling** - Tratamento padronizado de erros com códigos específicos
- ✅ **Clean Architecture** - Separação em camadas (Domain, Application, Infrastructure, Presentation)

## 🚀 Quick Start

### Pré-requisitos

- Go 1.24 ou superior
- Docker e Docker Compose
- Make (opcional, mas recomendado)
- **VS Code** (recomendado - veja [VS Code Quick Start Guide](docs/VSCODE_GUIDE.md))

### Instalação

1. **Clone o repositório**

```bash
git clone https://github.com/BrunnoQ/cadastro-pessoas.git
cd cadastro-pessoas
```

2. **Configure as variáveis de ambiente**

```bash
cp configs/config.local.yaml.example configs/config.local.yaml
# Edite configs/config.local.yaml conforme necessário
```

3. **Abra no VS Code (recomendado)**

```bash
code .
# Pressione F5 para iniciar com debugger
# Ou veja docs/VSCODE_GUIDE.md para guia completo
```

4. **OU inicie manualmente:**

**Inicie o MongoDB com Docker**

```bash
docker-compose up -d
```

**Instale as dependências**

```bash
go mod download
```

5. **Execute a aplicação**

```bash
go run ./cmd/api
# ou usando Make
make run
```

A API estará disponível em `http://localhost:8080`

### Gerenciamento da Aplicação

**Parar a aplicação:**

```bash
# Método 1: Se iniciou com Ctrl+C, pressione Ctrl+C no terminal

# Método 2: Matar processo por porta
lsof -i :8080 | grep LISTEN | awk '{print $2}' | xargs kill -9

# Método 3: Matar processo por nome
pkill -9 cadastro-pessoas

# Método 4: No VS Code, pressione Shift+F5 para parar o debug
```

**Verificar se a aplicação está rodando:**

```bash
# Verificar processo
ps aux | grep cadastro-pessoas

# Testar endpoint
curl http://localhost:8080/api/v1/health
```

## 📚 API Endpoints

### Health Check

```http
GET /api/v1/health
```

### Pessoas

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/api/v1/persons` | Criar nova pessoa |
| `GET` | `/api/v1/persons` | Listar pessoas (com paginação) |
| `GET` | `/api/v1/persons/:id` | Buscar pessoa por ID |
| `PUT` | `/api/v1/persons/:id` | Atualizar pessoa |
| `DELETE` | `/api/v1/persons/:id` | Deletar pessoa |

### Exemplos de Uso

**Criar pessoa:**

```bash
curl -X POST http://localhost:8080/api/v1/persons \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João",
    "surname": "Silva",
    "birthdate": "1990-05-15",
    "sex": "M",
    "addresses": [{
      "street": "Rua das Flores",
      "number": "123",
      "neighborhood": "Centro",
      "city": "São Paulo",
      "state": "SP",
      "country": "Brasil",
      "zip_code": "01234-567"
    }],
    "contacts": [{
      "type": "email",
      "value": "joao.silva@example.com"
    }]
  }'
```

**Listar pessoas (com paginação):**

```bash
curl "http://localhost:8080/api/v1/persons?limit=10&offset=0"
```

**Buscar pessoa:**

```bash
curl http://localhost:8080/api/v1/persons/{id}
```

**Atualizar pessoa:**

```bash
curl -X PUT http://localhost:8080/api/v1/persons/{id} \
  -H "Content-Type: application/json" \
  -d '{"name": "João Pedro", "surname": "Silva Santos"}'
```

**Deletar pessoa:**

```bash
curl -X DELETE http://localhost:8080/api/v1/persons/{id}
```

## 🏛️ Arquitetura

Este projeto segue **Clean Architecture** com separação rigorosa de camadas:

### Camadas

```
┌─────────────────────────────────────────────────────┐
│            Presentation Layer (HTTP)                │
│  ┌─────────────────────────────────────────────┐   │
│  │   Handlers, Middleware, Routes, DTOs        │   │
│  └─────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────┐
│           Application Layer (Use Cases)             │
│  ┌─────────────────────────────────────────────┐   │
│  │   Business Logic, DTOs, Orchestration       │   │
│  └─────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────┐
│          Domain Layer (Entities & Rules)            │
│  ┌─────────────────────────────────────────────┐   │
│  │   Entities, Value Objects, Repositories     │   │
│  └─────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────┐
│        Infrastructure Layer (External)              │
│  ┌─────────────────────────────────────────────┐   │
│  │   MongoDB, Config, Logger, External APIs    │   │
│  └─────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
```

## 🏛️ Princípios de Desenvolvimento

Este projeto segue rigorosamente a **Constitution** definida em `.specify/memory/constitution.md`:

1. **Clean Code** - Código limpo, legível e auto-documentado
2. **Clean Architecture** - Separação de camadas e inversão de dependências
3. **Code Reusability & DRY** - Máximo reaproveitamento através de abstração
4. **Design Patterns** - Repository, Use Case, DTO, Factory patterns
5. **Function & Method Cohesion** - Funções coesas com responsabilidade única (< 20 linhas)
6. **Performance Standards** - Indexes MongoDB, paginação, connection pooling
7. **Security Requirements** - Validação de entrada, error handling seguro

## 📁 Estrutura do Projeto

```
cadastro-pessoas/
├── cmd/
│   └── api/                    # Application entry point
│       └── main.go
├── internal/
│   ├── domain/                 # Domain layer (entities, repositories)
│   │   ├── entities/          # Business entities
│   │   └── repositories/      # Repository interfaces
│   ├── application/            # Application layer (use cases, DTOs)
│   │   ├── usecases/          # Business workflows
│   │   └── dto/               # Data transfer objects
│   ├── infrastructure/         # Infrastructure layer
│   │   ├── persistence/       # Database implementations
│   │   │   └── mongodb/
│   │   ├── config/            # Configuration management
│   │   └── logger/            # Logging setup
│   └── presentation/           # Presentation layer (HTTP)
│       └── http/
│           ├── handlers/      # HTTP handlers
│           ├── middleware/    # HTTP middleware
│           └── router.go      # Route registration
├── pkg/                        # Public reusable packages
│   ├── errors/                # Custom error types
│   └── validator/             # Input validators
├── configs/                    # Configuration files
├── docker/                     # Docker files
├── tests/                      # Tests
├── specs/                      # Spec Kit specifications
└── .specify/                   # Spec Kit templates
└── tests/               # Testes (a ser criado)
```

## 🎯 Workflow de Desenvolvimento

1. **Especificação** - Definir requisitos com `/speckit.specify`
2. **Planejamento** - Criar arquitetura com `/speckit.plan`
3. **Tarefas** - Quebrar em tarefas com `/speckit.tasks`
4. **Implementação** - Desenvolver com `/speckit.implement`
5. **Refatoração** - Manter código limpo continuamente

### Quality Gates

Todo código deve passar por:

- ✅ Linting (zero erros)
- ✅ Formatação automática
- ✅ Testes unitários (mín. 80% cobertura)
- ✅ Testes de integração
- ✅ Validação de segurança (input validation, sem secrets)
- ✅ Testes de performance (sem regressões)
- ✅ Code review
- ✅ Verificação contra Constitution

## 📚 Documentação

- **Constitution**: `.specify/memory/constitution.md` - Princípios fundamentais do projeto
- **Templates**: `.specify/templates/` - Templates para documentação
- **Specs**: `specs/` - Especificações de features (quando criadas)

## 🔧 Tecnologias

```

## 🛠️ Desenvolvimento

### Comandos Make

```bash
make build          # Compilar aplicação
make run            # Executar aplicação
make test           # Executar testes
make test-coverage  # Executar testes com cobertura
make lint           # Executar linter
make docker-up      # Iniciar MongoDB com Docker
make docker-down    # Parar MongoDB
make clean          # Limpar binários
```

### Comandos Go

```bash
# Desenvolvimento com hot reload
air

# Executar testes
go test ./...
go test -v ./...
go test -cover ./...

# Linting
golangci-lint run

# Formatar código
gofmt -s -w .
goimports -w .

# Build para produção
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app ./cmd/api
```

### Spec-Driven Development

Este projeto utiliza [GitHub Spec Kit](https://github.com/github/spec-kit):

```bash
/speckit.constitution  # Ver/atualizar princípios
/speckit.specify      # Criar especificação de feature
/speckit.plan         # Criar plano de implementação
/speckit.tasks        # Gerar lista de tarefas
/speckit.implement    # Executar implementação
```

## 🔧 Configuração

### Variáveis de Ambiente

Crie `configs/config.local.yaml` baseado em `config.local.yaml.example`:

```yaml
server:
  port: 8080
  mode: debug  # debug, release

mongodb:
  uri: mongodb://localhost:27017
  database: cadastro_pessoas
  timeout: 10s

logging:
  level: info  # debug, info, warn, error
  format: json  # json, console
```

### MongoDB

O projeto requer MongoDB 6.0 ou superior. Use Docker Compose para desenvolvimento:

```bash
docker-compose up -d
```

Ou instale localmente seguindo a [documentação oficial](https://docs.mongodb.com/manual/installation/).

## 🧪 Testes

```bash
# Executar todos os testes
go test ./...

# Testes com cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Testes específicos
go test ./internal/domain/...
go test ./internal/application/...

# Testes com race detector
go test -race ./...
```

## 📊 Métricas de Qualidade

- ✅ Cobertura de testes: Alvo 80%+
- ✅ Linter: golangci-lint sem erros
- ✅ Performance: Respostas < 100ms (p95)
- ✅ Segurança: Validação completa de entrada
- ✅ Clean Architecture: Separação estrita de camadas

## 🔐 Segurança

- Validação de entrada em todas as camadas
- Sanitização de dados para logs (prevenção de log injection)
- Error handling sem exposição de detalhes internos
- Optimistic locking para prevenir conflitos de atualização
- Middleware de CORS configurável

## 🚀 Deployment

### Docker

```bash
# Build da imagem
docker build -t cadastro-pessoas:latest .

# Executar container
docker run -p 8080:8080 \
  -e MONGODB_URI=mongodb://mongo:27017 \
  cadastro-pessoas:latest
```

### Docker Compose

```bash
docker-compose up -d
```

## 📖 Documentação Adicional

- [AGENTS.md](AGENTS.md) - Guia completo para desenvolvimento em Go
- [Constitution](.specify/memory/constitution.md) - Princípios do projeto
- [Specs](specs/) - Especificações de features
- [API Contracts](specs/001-pessoa-crud-api/contracts/) - Contratos da API

## 🤝 Contribuindo

1. Leia a Constitution (`.specify/memory/constitution.md`) e [AGENTS.md](AGENTS.md)
2. Siga os princípios de Clean Code e Clean Architecture
3. Use o workflow Spec-Driven Development
4. Garanta 80%+ de cobertura de testes
5. Execute `golangci-lint run` antes de commitar
6. Solicite code review antes de mergear

### Checklist de PR

- [ ] Código segue Clean Architecture
- [ ] Funções < 20 linhas
- [ ] Testes escritos (cobertura 80%+)
- [ ] Linter passando
- [ ] Documentação atualizada
- [ ] Constitution respeitada

## 📄 Licença

MIT License - veja [LICENSE](LICENSE) para detalhes.

## 👥 Autores

- Brunno Q - [@BrunnoQ](https://github.com/BrunnoQ)

## 🙏 Agradecimentos

- [GitHub Spec Kit](https://github.com/github/spec-kit) - Spec-Driven Development workflow
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Princípios arquiteturais

---

**Status do Projeto**: ✅ Produção Ready - CRUD Completo Implementado
