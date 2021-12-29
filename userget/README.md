#meli User - userget

Este serviço é responsável por fazer as buscas por Nome na base de dados.
Foi implementado somente 1 filtro pelo campo first_name, mas tudo poderá extender.

A busca é um array de names, ex: "?names=[Jefferson,Paul]" lembrando que no bash precisa escapar os colchetes.

Este serviço utiliza do Redis para fazer cache, o time do cache poderá ser definido no config. Default do Cache é 60 segundos.

**Buscas**
	- Buscando vários nomes pelo names
	- Buscando vários cpfs pelo cpf
	- Buscando vários lastnames pelo lastname
	- Buscando vários emails pelo email

#### Index mongo

Para que as buscas funcionem para N elementos, vamos criar um index no mongo para isto precisa logar no mongo para rodar o script, não foi feito o migration.

```bash
$ db.stores.createIndex( { first_name: "text" } )
```
#### Deploy

Você poderá utilizar o sh deploy.sh para subir todo ambiente deste serviço, ele será composto por Mongo, Redis e o serviço userget.
Ele irá compilar a imagem e disponibiliza-la para que possa dar start.

Após rodar o deploy basta subir o serviço
```bash
$ docker-compose up -d
```

#### User GET PING

Foi criado este endpoint somente para checar o serviço e nada mais seria nosso health para o serviço.

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
localhost:8082/v1/user/ping 

```

#### User GET

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
"localhost:8082/v1/user?firstname=\[Jefferson,Paul\]"

```

A saida será:
```bash
[
   {
      "id":"6bd13c31d818238992002e6ad216455593b4f32d",
      "first_name":"Jefferson",
      "last_name":"Otoni",
      "birthday":"1975-08-20",
      "cpf":"20182384004",
      "email":"jeff.otoni1@gmail.com",
      "create":"2021-09-14T13:54:12.071Z",
      "update":"2021-09-14T13:54:12.071Z",
      "ip":"127.0.0.1",
      "agent":"curl/7.68.0"
   },
   {
      "id":"9bf8cf8ef01613305638b2be9aecd23418c50e93",
      "first_name":"Paul",
      "last_name":"Churchill",
      "birthday":"1920-08-20",
      "cpf":"29145037094",
      "email":"paul@gmail.com",
      "create":"2021-09-14T13:53:40.505Z",
      "update":"2021-09-14T13:53:40.505Z",
      "ip":"127.0.0.1",
      "agent":"no agent"
   }
]
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