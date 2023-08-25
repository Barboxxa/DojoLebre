# DojoLebre
Repo para dojo da liber

## Executando a aplicação
Para executar a aplicação, basta rodar o comando:
```bash
make run
```

## Instalando ferramentas

### Linter
Como linter, usaremos o Staticcheck. Para instalar localmente, é necessário rodar o comando:
```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```
Para executar o linter, basta rodar o comando:
```bash
make lint
```
