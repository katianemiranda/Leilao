Objetivo

Adicionar uma funcionalidade crítica ao sistema de leilões existente: o fechamento automático. Atualmente, o projeto permite criar leilões e dar lances, mas o leilão nunca expira. Sua missão é utilizar Goroutines para garantir que o leilão seja encerrado automaticamente após um tempo pré-definido.

Requisitos de Testes

Implemente um teste automatizado que comprove que o fechamento está ocorrendo.

Cenário: Criar um leilão, aguardar o tempo configurado e verificar se o status mudou automaticamente para fechado sem intervenção manual.
 

Dicas de Implementação

Lembre-se que estamos lidando com concorrência. Certifique-se de que sua solução não bloqueie a thread principal.

Analise como o sistema atual verifica se um leilão é válido na rotina de criação de bid para entender a lógica de tempo e status.

 

Entregável

Código Fonte: O link para o seu repositório contendo a implementação completa.

#### Instruções para rodar o projeto
Na raiz, Leilao, executar o comando
1 - docker compose build
2 - docker compose up -d

Após subida da aplicação. na pasta leilao/test, tem um arquivo auction.http com alguns comandos para testes