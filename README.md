# Cadastro de Pessoas

Sistema de cadastro de pessoas desenvolvido seguindo princípios de Clean Code e Clean Architecture.

## 🏛️ Princípios de Desenvolvimento

Este projeto segue rigorosamente a **Constitution** definida em `.specify/memory/constitution.md`, que estabelece os seguintes princípios fundamentais:

### Princípios Fundamentais (NON-NEGOTIABLE)

1. **Clean Code** - Código limpo, legível e auto-documentado
2. **Clean Architecture** - Separação de camadas e inversão de dependências
3. **Code Reusability & DRY** - Máximo reaproveitamento através de abstração
4. **Design Patterns** - Uso apropriado de padrões de projeto
5. **Function & Method Cohesion** - Funções coesas com responsabilidade única

## 🚀 Começando

Este projeto utiliza **Spec-Driven Development** através do [GitHub Spec Kit](https://github.com/github/spec-kit).

### Comandos Disponíveis

Use os seguintes comandos slash com GitHub Copilot:

- `/speckit.constitution` - Ver/atualizar princípios do projeto
- `/speckit.specify` - Criar especificação de feature
- `/speckit.plan` - Criar plano de implementação técnica
- `/speckit.tasks` - Gerar lista de tarefas acionáveis
- `/speckit.implement` - Executar implementação

### Comandos Opcionais de Qualidade

- `/speckit.clarify` - Esclarecer áreas ambíguas
- `/speckit.analyze` - Análise de consistência entre artefatos
- `/speckit.checklist` - Gerar checklists de qualidade

## 📁 Estrutura do Projeto

```
cadastro-pessoas/
├── .github/              # Configurações GitHub (workflows, prompts, agents)
├── .specify/             # Templates e memória do Spec Kit
│   ├── memory/          # Constitution e documentos fundamentais
│   ├── templates/       # Templates para specs, plans, tasks
│   └── scripts/         # Scripts de automação
├── .vscode/             # Configurações do VS Code
├── specs/               # Especificações de features (gerado pelo Spec Kit)
├── src/                 # Código fonte (a ser criado)
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
- ✅ Code review
- ✅ Verificação contra Constitution

## 📚 Documentação

- **Constitution**: `.specify/memory/constitution.md` - Princípios fundamentais do projeto
- **Templates**: `.specify/templates/` - Templates para documentação
- **Specs**: `specs/` - Especificações de features (quando criadas)

## 🔧 Tecnologias

_(A ser definido durante o planejamento de features)_

## 🤝 Contribuindo

1. Leia a Constitution (`.specify/memory/constitution.md`)
2. Siga os princípios de Clean Code e Clean Architecture
3. Use o workflow Spec-Driven Development
4. Garanta que todos os quality gates passem
5. Solicite code review antes de mergear

## 📄 Licença

_(A ser definida)_
