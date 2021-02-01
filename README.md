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



<p align="center">
<img width="600" src="./public/assets/images/Maratona.gif?raw=true">
</p>