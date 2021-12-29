# Meli User - userwrite

Este repo é responsável por um cadastro onde teremos somente um recurso que será /v1/user e teremos os métodos: POST,PUT e GET.

**As operações serão:**

	Cadastro;
	- Atualização;
	- Busca de um único recurso;
	- Busca de vários recursos

Irá possuir algumas regras de validação:

	Algumas regras:
	 - Somente usuários acima de 18 anos serão cadastrados;
	 - Não será permitido usuários com e-mail e CPF duplicados;
	 - Quando buscar por vários usuários, deve permitir realizar um filtro pelo nome;
	 - Permitir a alteração parcial;
	 - Validar se o CPF é válido (dígitos verificadores);
	 - Validar email
	 - Checar a senha e confirm senha
	 - Enviar email para validação e confirmação do email pelo usuário
	 - Password deverá ser criptografado
	 - 
A API desenvolvida foi desfragmentada em 2 serviços:

	- Serviço userwrite
	- Serviço userget

Para que possamos escalar de forma eficiente e com maior controle dividimos o problema.
O Service Write será responsável por fazer todas operações de persistência e atualizações.
O Service Get será responsável por fazer as buscas e neste cenário poderemos usar CACHE para deixar tudo ainda mais flexível e rápido.

#### APICORE

Foi criado um core que chamamos de apicore, ele é um conjunto de pkgs e configurações globais que é usado em diversos serviços,
e como Go permite brincarmos com estes pkgs de forma flexível exploramos este recurso da lang ao nosso favor.
Tudo que será comum a todos serviços é criado um pkg no apicore, alguns exemplo: envio de email, cryptografia, concatenação, validadores etc.

#### Arquitetura Hexagonal

Foi utilizado a arquitetura hexagonal, onde iremos explorar todo seu poder quando trata-se de desacoplamento e facilidade de conectar dentre os pacotes etc.
O maior objetivo desta arquitetura é criar componentes de aplicativos fracamente acoplados que possam ser facilmente conectados ao ambiente de software.

#### Desenho da Arquitetura modelo One

Este é uma prévia do projeto e da sua implementação.

O promethes e grafana não subi nesta versão, para instrumentar o código precisamos de mais um tempinho.

Criamos uma saída para Stdout no padrão ElasticSearch para que possamos mapear todo Log para o Elastic ou qualquer outra tecnologia que venhamos utilizar.

![operator](arquitetura/UserOne.png?raw=true "Versão usando docker-compose")

#### Desenho da Arquitetura modelo Two

Este é uma prévia do projeto e da sua implementação se fossemos implementar em Kubernetes.
Neste cenário utilizamos um Gateway e autenticações, Cache, toda parte de instrumentações que teriamos em nossa API tudo seria feito pelo nosso Gateway que neste cenário coloquei o Kong.

![operator](arquitetura/UserTwo.png?raw=true "Versão usando k8s")

#### Swagger

Este é uma prévia do projeto em Swagger, preciso atualiza-lo mas já da para ter uma ideia.

![operator](swagger/swagger.png?raw=true "swagger do projeto")

#### Test Stress

Utilizei o k6 para fazer os test de stress em nossa API.
Preciso ainda implementar o POST para N.

![operator](k6/k6.png?raw=true "k6 result do Put")

**Test do Get**

![operator](k6/k6.get.png?raw=true "k6 result do Get")


#### A estrutura está organizada da seguinte forma:

```bash
├── gusermeli
│    ├── apicore
│	 │	 ├── config
│	 │	 ├── middleware
│	 │	 ├── pkg
├── userwrite
│    ├── controler
│    │   ├── handler
│    │       └── config.go
│    │       └── ping.go
│    │       └── ping_test.go
│    │       └── route.go
│    │       └── user.post.go
│    │       └── user.post_test.go
│    │       └── user.put.go
│    │       └── user.put_test.go
│    │       └── user.get.go
│    │       └── user.get_test.go
│    ├── domain
│    │ 	 ├── repo
│    │       └── user.create.go
│    │       └── user.update.go
│    ├── Makefile
│    ├── Dockerfile
│    ├── README.md
│    ├── docker-compose.yaml
│    ├── go.mod
│    ├── go.sum
├── userget
│    ├── controler
│    │   ├── handler
│    │       └── config.go
│    │       └── ping.go
│    │       └── ping_test.go
│    │       └── route.go
│    │       └── user.get.go
│    │       └── user.get_test.go
│    ├── domain
│    │ 	 ├── repo
│    │       └── user.get.go
│    ├── Makefile
│    ├── Dockerfile
│    ├── README.md
│    ├── docker-compose.yaml
│    ├── go.mod
│    ├── go.sum
├── swagger
├── arquitetura
├── docker-compose.yaml
├── LICENSE
├── README.md
└── 
```

#### Docker Compose

Este docker-compose.yaml está configurado para buscar as imagens locais que criamos de cada serviço.

#### Deploy

Criando as imagens de dando start em todo projeto
```bash
$ sh deploy.sh
```
Saída:

```bash
Creating gusermeli_redis_1       ... done
Creating userwrite               ... done
Creating gusermeli_mongo-users_1 ... done
Creating userget                 ... done
         Name                        Command               State                      Ports                    
---------------------------------------------------------------------------------------------------------------
gusermeli_mongo-users_1   docker-entrypoint.sh mongod      Up      0.0.0.0:27017->27017/tcp,:::27017->27017/tcp
gusermeli_redis_1         /opt/bitnami/scripts/redis ...   Up      6379/tcp                                    
userget                   /userget                         Up      0.0.0.0:8082->8082/tcp,:::8082->8082/tcp    
userwrite                 /userwrite                       Up      0.0.0.0:8081->8081/tcp,:::8081->8081/tcp    
 Prontinho .... 
```

Depois que subimos todos os serviços podemos agora alimentar para que nossos exemplos funcionem lindamente 😍.

**Buscar na Base usando Filtro para first_name**

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
"localhost:8082/v1/user?firstname=\[Paul,Jefferson\]"
```

**Buscando por Cpf**
```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
"localhost:8082/v1/user?cpf=\[29145037094,20182384004\]"
```


**Buscando por Last Name**
```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
"localhost:8082/v1/user?lastname=\[Otoni\]"
```

**Buscando por Email**
```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
"localhost:8082/v1/user?email=\[Otoni\]"
```