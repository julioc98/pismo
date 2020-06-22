# Pismo

[![Build Status](https://travis-ci.com/julioc98/pismo.svg?token=4SjCRRz2dpNCgC3iccDx&branch=master)](https://travis-ci.com/julioc98/pismo)

[Acesse a API](https://pismo-api.herokuapp.com/)

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/julioc98/pismo)


### Pré-requisitos

* [Golang](https://github.com/golang/go)

OU (Recomendado)

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)*

### Como rodar localmente?**

Baixe o repositório, entre no diretório e rode o comando:

```
make run/docker
```
Depois acesse a url
```
http://localhost:5001/
```

### Como rodar os testes?**

```
make test/docker
```

### Como fazer o deploy?

- Temos um Dockerfile para quando queremos levar essa aplicação para produção(Dockerfile.production). Ela é diferente da desnvolvimento porque ela usa `multi-stage build` para deixar a imagem bem menor com um S.O. mais leve e apenas o binário da aplicação para rodar.
- No caso estou fazendo o deploy no [Heroku](https://www.heroku.com/). Então só "commitar na master" ou apertar o botão de deploy:

  - [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/julioc98/pismo)


##### *Para Facilitar | **Com Docker + Docker Compose

# Outras informações

- Usei libs(packages) padrões ao maximo, tentando extrair tudo da linguagem Go.

- Nos testes usei [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests).

- Para me organizar gosto de usar o Github Project pois ele é bem fácil de usar e tem um ótimo suporte a Markdown. Criei um board bem simples [aqui](https://github.com/julioc98/pismo/projects/1).

- Tambem separei em milestones por nivel, [aqui](https://github.com/julioc98/pismo/milestones) 


- Para organizar o código gosto do [Gitflow](https://github.com/nvie/gitflow) por ser uma forma “padrão” de trabalhar com git e se da muito bem com a recomendação [SEMVER](https://semver.org/).

- Um CI bem simples foi configurado no [TravisCI](https://travis-ci.com/) para pegar Pushs e PRs. A badge fica bem intuitiva no README ([![Build Status](https://travis-ci.com/julioc98/pismo.svg?token=4SjCRRz2dpNCgC3iccDx&branch=master)](https://travis-ci.com/julioc98/pismo)).

