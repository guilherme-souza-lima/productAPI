# productAPI

## :rocket: Docker Run App
### `Primeiro passo` vamos criar o arquivo .env
```sh
make env-docker
```
### `Segundo passo` vamos executar o docker-compose para subir os containers
```sh
make up
```
stop docker-compose comando:`make down`.

### `Terceiro passo`  Testar os endpoints da API.
## :books: Endpoints API
* ### `POST` `http://127.0.0.1:8080/create/product`
    * _Descrição_: Endpoint para cadastrar produto
        * `Body`
        ```JSON
        {
          "sku":"123456789",
          "name":"nome produto",
          "stock_total":10,
          "stock_cutting":1,
          "price_from":9.99,
          "price_per":8.88
        }
        ```
      
* ### `PUT` `http://127.0.0.1:8080/update/product`
    * _Descrição_: Endpoint para alterar produto
        * `Body`
        ```JSON
        {
          "uuid": "abc.....",
          "sku":"123456789",
          "name":"nome produto",
          "stock_total":10,
          "stock_cutting":1,
          "price_from":9.99,
          "price_per":8.88
        }
        ```
      
* ### `DELETE` `http://127.0.0.1:8080/delete/product`
    * _Descrição_: Endpoint para excluir um produto
        * `Body`
        ```JSON
        {
          "uuid": "abc.....",
          "sku":"123456789"
        }
        ```
      
* ### `GET` `http://127.0.0.1:8080/list-all/product`
    * _Descrição_: Endpoint para listar todos os produtos

* ### `POST` `http://127.0.0.1:8080/transaction`
    * _Descrição_: Endpoint para fazer uma transação
        * `Body`
        ```JSON
        {
          "uuid": "abc.....",
          "sku":"123456789",
          "amount": 5
        }
        ```
      
* ### `GET` `http://127.0.0.1:8080/list-all/transaction`
    * _Descrição_: Endpoint para listar todas as transações