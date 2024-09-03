# Arquitetura Hexagonal
Aplicação de produto usando arquitetura hexagonal usando golang

## Coniguração
### Subindo container docker
```sh docker-compose up -d```

### Acessando o container
```sh docker exec -it appproduct bash```

### Gerando mocks
#### Instalando a biblioteca pra geração de mocks
```sh go install github.com/golang/mock/mockgen@v1.5.0```


```sh mockgen -destination=application/mocks/application.go -source=application/product.go application```

```sh mockgen -destination=caminho-onde-mock-sera-criado -source=caminho-da-interface-projeto nome-pacote-projeto```

#### Criando banco de dados
Instalando banco de dados sqlite3
```sh sudo apt-get install sqlite3```

Criando arquivo do banco de dados
```sh touch sqlite.db```

Abrindo o banco de dados
```sh sqlite3 sqlite.db```

Criando a tabela products
```sh create table products(id string, name string, price float, status string);```
Visualizando a tabela criada
```sh .tables ```


#### Cobra cli
[https://github.com/spf13/cobra?tab=readme-ov-file](https://github.com/spf13/cobra?tab=readme-ov-file)

```sh cobra-cli init --pkg-name=github.com/tayron/golang-arquitetura-hexagonal```
Comando para adicionar um comando, no caso ele se chama cli
```sh cobra-cli add cli ```

Rodando comando para ver as ações permitidas
```sh go run main.go cli --help```

Rodando comando para criar produto
```sh go run main.go cli -a=create -n=Product CLI -p=25.0 ```

Rodando comando para pegar produto
```sh go run main.go cli -a=get --id=752e6333-fbec-4af4-8741-3751e727c2ce ```


#### Testando HTTP service
Startando servidor HTTP:
```sh go run main.go http ```

Chamando a api:
```sh curl http://localhost:8080/product/f56b3370-a708-484b-b36d-a9af4cfabf1c ```