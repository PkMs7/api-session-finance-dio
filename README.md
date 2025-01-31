# API de Sistema de Planejanento Financeiro

API construída para fins de estudos de uma entrega completa, com testes, formatação de código e pipeline de Continuous Integration.

- Aplicação para CI: https://circleci.com/

- Biblioteca para métricas: https://pkg.go.dev/github.com/prometheus/client_golang/prometheus

## Passos para startar a aplicação

- Pré-requisito: Possuir docker na máquina

- Comandos para subir a aplicação:

    - Clonar para repositório local: ```git clone https://github.com/PkMs7/api-session-finance-dio```
    - Buildar a imagem: ```docker build -t finance .```
    - Subir o contâiner: ```docker-compose up```

- Comandos para teste:
    - Testes na pasta do arquivo
        ```go test -v```
    
    - Testes no projeto inteiro (procura por arquivos de teste "_test"):
        Comando na pasta raiz
        ```go test ./...```

- Comando para formatação do código:
    Comando na pasta raiz
        ```go fmt ./...```

<!-- Teste de Integração Jira com Pipeline CI/CD -->