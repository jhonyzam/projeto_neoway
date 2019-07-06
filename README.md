# Projeto Neoway
##Objetivo
Ler o arquivo `base/base_test.txt` em GOLANG, efetuar as conversões de campos necessarios e inseri-lo no banco de dados PostgreSQL.
One Paragraph of project description goes here

##Pré requisitos
Possuir o docker instalado, com docker-compose, ou um ambiente com Golang e PostgreSQL

##Passo 1
######Usando docker 
Apos baixar o repositorio no seu ambiente docker, rode o comando dentro do arquivo `run.sh`

##Passo 2
Conecte-se no banco de dados postgres:

| | |
| - | - |
| **Host:** | ip_adress_docker |
| **Port:** | 5432 |
| **Database:** | neoway |
| **User:** | postgres |
| **Password:** | @postSenha123 |

##Passo 2.1
Execute a query dentro de `postgres/create.sql`

```sql
CREATE TABLE datastore(
	   id serial PRIMARY KEY,
	   cpf VARCHAR (50) NOT NULL,
	   private INTEGER,
	   incompleto INTEGER,
	   lastDate DATE,
	   avgTicket NUMERIC (10, 2),
	   lastTicket NUMERIC (10, 2),
	   storeFrequent VARCHAR (50),
	   storeLast VARCHAR (50)
);
```
##Passo 3
Execute os seguintes comandos no navegador ou num endpoint tester (Postman por exemplo)

#####Carregando base de dados:
``` <ip_adress_docker>:3000/datastore/execute ```
<br>
#####Listando base de dados:
```<ip_adress_docker>:3000/datastore```

##Autor
Jhonatan Reis
