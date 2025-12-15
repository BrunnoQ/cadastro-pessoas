# Contribuindo para Cadastro de Pessoas

Obrigado por considerar contribuir para este projeto! Este guia ajudará você a entender nosso processo de desenvolvimento e como submeter contribuições.

## 📋 Índice

- [Código de Conduta](#código-de-conduta)
- [Como Posso Contribuir?](#como-posso-contribuir)
- [Processo de Desenvolvimento](#processo-de-desenvolvimento)
- [Padrões de Código](#padrões-de-código)
- [Commits e Pull Requests](#commits-e-pull-requests)
- [Testes](#testes)

## 📜 Código de Conduta

Este projeto adere aos princípios de respeito mútuo e colaboração profissional. Esperamos que todos os contribuidores:

- Sejam respeitosos e construtivos nas discussões
- Aceitem feedback com profissionalismo
- Foquem no que é melhor para o projeto
- Demonstrem empatia com outros contribuidores

## 🤝 Como Posso Contribuir?

### Reportando Bugs

Antes de criar um bug report:

1. Verifique se o bug já não foi reportado
2. Colete informações relevantes:
   - Versão do Go
   - Sistema operacional
   - Logs de erro
   - Passos para reproduzir

Crie uma issue com:

- Título claro e descritivo
- Descrição detalhada do problema
- Passos para reproduzir
- Comportamento esperado vs. atual
- Screenshots (se aplicável)

### Sugerindo Melhorias

Para sugerir novas features:

1. Verifique se a feature já não foi sugerida
2. Explique **por que** a feature seria útil
3. Forneça exemplos de uso
4. Considere o impacto em performance e arquitetura

### Pull Requests

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Faça commit das mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 🔄 Processo de Desenvolvimento

### Spec-Driven Development

Este projeto utiliza **Spec-Driven Development** com [GitHub Spec Kit](https://github.com/github/spec-kit):

```bash
1. /speckit.specify      # Criar especificação
2. /speckit.plan         # Criar plano técnico
3. /speckit.tasks        # Gerar tarefas
4. /speckit.implement    # Implementar
```

### Constitution

**IMPORTANTE**: Todas as contribuições DEVEM seguir a Constitution (`.specify/memory/constitution.md`).

Os 7 princípios fundamentais são **NON-NEGOTIABLE**:

1. ✅ Clean Code
2. ✅ Clean Architecture
3. ✅ Code Reusability & DRY
4. ✅ Design Patterns
5. ✅ Function & Method Cohesion
6. ✅ Performance Standards
7. ✅ Security Requirements

### Workflow

1. **Planejamento**
   - Revise a Constitution
   - Crie spec para a feature
   - Gere plano de implementação
   - Quebre em tarefas acionáveis

2. **Implementação**
   - Siga Clean Architecture
   - Escreva testes primeiro (TDD)
   - Mantenha funções < 20 linhas
   - Use design patterns apropriados

3. **Validação**
   - Execute testes (min 80% cobertura)
   - Execute linter (sem erros)
   - Valide performance
   - Revise segurança

4. **Documentação**
   - Atualize README se necessário
   - Documente APIs públicas
   - Adicione comentários onde apropriado

## 📝 Padrões de Código

### Clean Code

```go
// ✅ BOM - Nome descritivo, função pequena
func ValidateEmail(email string) error {
    if !emailRegex.MatchString(email) {
        return errors.NewValidationError("invalid email format", nil)
    }
    return nil
}

// ❌ RUIM - Nome vago, lógica complexa em uma função
func validate(e string) bool {
    // múltiplas responsabilidades...
}
```

### Clean Architecture

```
Regra de Dependência: Camadas externas dependem de internas
Domain ← Application ← Infrastructure
Domain ← Application ← Presentation
```

```go
// ✅ BOM - Domain não depende de nada
package entities
type Person struct { ... }

// ✅ BOM - Use case depende de interface do domain
package usecases
func (uc *CreatePersonUseCase) Execute(ctx context.Context, req dto.CreatePersonRequest) {}

// ❌ RUIM - Domain importando infrastructure
package entities
import "go.mongodb.org/mongo-driver/bson"  // NUNCA!
```

### Naming Conventions

- **Packages**: lowercase, singular (`user`, não `users`)
- **Files**: snake_case (`person_repository.go`)
- **Types**: PascalCase (`PersonRepository`)
- **Functions**: PascalCase (exported), camelCase (private)
- **Constants**: PascalCase ou UPPER_SNAKE_CASE
- **Variables**: camelCase

### Error Handling

```go
// ✅ BOM - Erros contextualizados
if err != nil {
    return fmt.Errorf("failed to create person: %w", err)
}

// ✅ BOM - Erros customizados
return errors.NewValidationError("invalid email", map[string]interface{}{
    "email": req.Email,
})

// ❌ RUIM - Erro genérico
if err != nil {
    return err
}
```

### Function Size

```go
// ✅ BOM - Função com ~10 linhas, responsabilidade única
func (s *PersonService) ValidateAge(birthdate time.Time) error {
    age := calculateAge(birthdate)
    if age < 18 {
        return errors.NewValidationError("person must be 18+", nil)
    }
    return nil
}

// ❌ RUIM - Função > 20 linhas
func (s *PersonService) CreatePersonAndValidateAndSendEmail(...) {
    // 50+ linhas de código...
}
```

## 🔧 Commits e Pull Requests

### Commit Messages

Siga [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add delete person endpoint
fix: resolve nil pointer in person service
refactor: extract validation to separate package
test: add integration tests for update endpoint
docs: update README with new API endpoints
perf: optimize database queries with indexes
security: add input sanitization for logs
```

### Pull Request Template

```markdown
## Descrição
[Descreva o que este PR faz]

## Tipo de Mudança
- [ ] Bug fix
- [ ] Nova feature
- [ ] Breaking change
- [ ] Atualização de documentação

## Constitution Compliance
- [ ] Clean Code: Funções < 20 linhas, nomes significativos
- [ ] Clean Architecture: Separação de camadas apropriada
- [ ] Code Reusability: Sem duplicação
- [ ] Design Patterns: Padrões apropriados utilizados
- [ ] Function Cohesion: Responsabilidade única
- [ ] Performance: Sem regressões
- [ ] Security: Input validado, sem segredos, error handling seguro

## Testing
- [ ] Unit tests adicionados/atualizados
- [ ] Integration tests adicionados/atualizados
- [ ] Todos os testes passando
- [ ] Cobertura ≥ 80%

## Checklist
- [ ] Code review por pelo menos um peer
- [ ] Documentação atualizada
- [ ] golangci-lint passando
- [ ] CHANGELOG.md atualizado (se aplicável)
```

## 🧪 Testes

### Executando Testes

```bash
# Todos os testes
go test ./...

# Com cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Específicos
go test ./internal/domain/...
go test -v ./internal/application/usecases/

# Com race detector
go test -race ./...
```

### Escrevendo Testes

```go
// Formato de nome: Test<FunctionName>_<Scenario>_<ExpectedResult>
func TestCreatePerson_ValidInput_ReturnsSuccess(t *testing.T) {
    // Arrange
    repo := &MockPersonRepository{}
    uc := NewCreatePersonUseCase(repo, logger)
    
    // Act
    result, err := uc.Execute(ctx, validRequest)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
}

// Table-driven tests
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "user@example.com", false},
        {"invalid email", "invalid", true},
        {"empty email", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Cobertura Mínima

- Cobertura total: **80%+**
- Domain layer: **90%+**
- Application layer: **85%+**
- Infrastructure layer: **70%+** (mocks para externos)

## 🔍 Code Review

### O que Revisores Procuram

1. **Constitution Compliance**
   - Princípios seguidos?
   - Clean Architecture mantida?
   - Performance adequada?

2. **Qualidade de Código**
   - Nomes descritivos?
   - Funções pequenas (<20 linhas)?
   - Sem duplicação?

3. **Testes**
   - Cobertura adequada?
   - Casos edge cobertos?
   - Testes legíveis?

4. **Segurança**
   - Input validado?
   - Erros tratados adequadamente?
   - Sem secrets no código?

### Para Contribuidores

- Responda feedback construtivamente
- Faça mudanças solicitadas prontamente
- Mantenha discussões focadas no código
- Peça esclarecimentos quando necessário

## 📚 Recursos

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Effective Go](https://golang.org/doc/effective_go)
- [GitHub Spec Kit](https://github.com/github/spec-kit)
- [AGENTS.md](AGENTS.md) - Guia completo de desenvolvimento

## ❓ Perguntas?

Se tiver dúvidas:

1. Revise a Constitution
2. Consulte AGENTS.md
3. Procure em issues existentes
4. Abra uma nova issue com a tag `question`

## 🙏 Obrigado

Suas contribuições tornam este projeto melhor. Agradecemos seu tempo e esforço! 🚀
