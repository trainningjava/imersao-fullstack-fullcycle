<h1 align="center">
    <img width="120" height="40" src="https://github.com/trainningjava/imersao-fullstack-fullcycle/blob/main/public/assets/images/FullCycle.png">
</h1>

### 🤔 Sobre o que se trata ? 
A Imersão Full Stack && Full Cycle é um evento online e 100% gratuito de uma semana (02/2021) que vai ajudar na prática programadores a desenvolverem as principais habilidades exigidas pelas empresas sendo capazes de trabalhar em projetos de grande porte, com total confiança e desenvolvendo do jeito certo.

Participe gratuitamente: https://imersao.fullcycle.com.br/

Sobre o projeto

É uma solução para simularmos transferências de valores entre banco fictícios através de chaves (email, cpf).
Simularemos diversos bancos e contas bancárias que possuem uma chave Pix atribuída.
Cada conta bancária poderá cadastrar suas chaves Pix.
Uma conta bancária poderá realizar uma transferência para outra conta em outro banco utilizando a chave Pix da conta de destino.
Uma transação não pode ser perdida mesmo que: o CodePix esteja fora do ar
Uma transação não pode ser perdida mesmo que: o Bando de destino esteja fora do ar

Sobre os bancos

O banco será um microserviço com funções limitadas a cadastro de contas e chaves Pix, bem como transferência de valores.
Utilizaremos a mesma aplicação para simularmos diversos bancos, mudando apenas cores, nome e código.
Nest.js no backend.
Nest.js como frontend.

Sobre o CodePix

O microserviço CodePix será responsável por intermediar as transferência bancárias
Receberá a transação de transferência
Encaminhará a transação para o banco de destino (Status: "pending")
Recebe a confirmação do bando de destino (Status: "confirmed")
Envia confirmação para o banco de origem informado quando o banco de destino processou
Recebe a confirmação do banco de origem de que ele processou (Status: "completed")
Marca a transação como completa (Status: "completed")

## :package: Desafios

- [desafios1](https://github.com/trainningjava/imersao-fullstack-fullcycle/tree/main/codepix)
Nesse desafio você deverá montar seu ambiente de desenvolvimento com Docker utilizando o Dockerfile e docker-compose de nosso projeto prático.

Crie uma struct de User com ID (uuid), Name e Email. Além disso, implemente uma função NewUser que seja capaz de validar que ID, Name e Email são obrigatórios, caso o contrário, ela deve retornar um erro.

- [desafios2](https://github.com/trainningjava/imersao-fullstack-fullcycle/tree/main/codepix)
Nesse desafio você deverá criar uma simples aplicação utilizando Golang que seja capaz de publicar uma mensagem em um tópico do Apache Kafka e também consumir mensagens desse tópico.

Para deixar mais claro a separação das responsabilidades, crie uma pasta chamada producer e outra consumer. Em cada uma delas crie um arquivo main.go que será responsável por produzir e consumir as mensagens respectivamente

Utilize o mesmo Dockerfile e docker-compose.yaml utilizados no projeto do CodePix. Fique na liberdade para escolher o nome do tópico que será utilizado



<p align="center">
<img width="600" src="./public/assets/images/Maratona.gif?raw=true">
</p>