#meli User - userwrite

Este repo é responsável por um cadastro onde teremos somente um recurso que será /v1/user e teremos os métodos: POST,PUT e GET.
As operações serão:

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


#### Docker Build
Para que possa funcionar os serviços docker-compose no raiz com os dois serviços precisamos buildar nossa imagem seja local ou para algum registry.

```bash
$ docker build --no-cache -f Dockerfile -t jeffotoni/userwrite
```

#### Docker Compose

Compilando e removendo cache
```bash
$ docker-compose build --no-cache
```

Subindo nosso serviço
```bash
$ docker-compose up -d
```

#### user POST PING

Foi criado este endpoint somente para checar o serviço e nada mais seria nosso health para o serviço.

```bash
$ curl -i -XPOST \
-H "Content-Type:application/json" \
localhost:8081/v1/user/ping 

```

#### user POST

Json a ser enviado tem a seguinte estrutura:

```bash
{
   "first_name":"Jefferson",
   "last_name":"Otoni",
   "birthday":"1975-08-20",
   "cpf":"201.823.840-04",
   "email":"jeff.otoni1@gmail.com",
   "password":"123456",
   "confirm_password":"123456"
}
```

```bash
$ curl -i -XPOST \
-H "Content-Type:application/json" \
localhost:8081/v1/user \
-d @json/user.post.json

```

Ò que será gravado no banco será:

```bash
{
   "_id":"706291f28228e3e64617b50e91cf68c6c42b37a1",
   "first_name":"Jefferson",
   "last_name":"Otoni",
   "birthday":"1975-08-20",
   "cpf":"20182384004",
   "email":"jeff.otoni1@gmail.com",
   "password":"$2a$10$TjaUv52pSfYJepIcEVyvqOLPqoDeJ/6qpz8JHyHkyzWTIrVQe1ot2",
   "create":"2021-09-14T05:22:26.057Z",
   "update":"2021-09-14T05:22:26.057Z",
   "ip":"127.0.0.1",
   "agent":"curl/7.68.0"
}
```

#### user PUT

Será atualizado somente alguns campos são eles: FirstName, LastName, Birthday. Qualquer outro campo que tentar colocar ou inserir no json será ignorado.

A nossa chave será o CPF e ser enviar dígitos no CPF ou não irá funcionar de qualquer forma, API já está validando.
Nosso json ficaria:
```bash
{
   "first_name":"Jefferson",
   "last_name":"Otoni",
   "birthday":"1980-08-20"
}
```

Nosso cURL:
```bash
$ curl -i -XPUT \
-H "Content-Type:application/json" \
localhost:8081/v1/user/20182384004 \
-d @json/user.put.json

```

Os campos email, cpf, passwords são todos campos que não serão alterados por este endpoint.
Para que possamos altera-los seria necessário criar novos serviços para cada um deles tratando suas particularidades de validação um a um.
