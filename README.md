# challengerklever
 
Uma API simples e rapida de UpVote System onde você pode votar na suas Criptomoedas favoritas!
Esta API foi criada utilizado a Linguagem Go juntamente com o banco de dados mongoDB.


Para ter acesso utilize http://localhost:PORT/coin

Métodos:

 >>> POST <<<
- /create:
Cria sua Criptomoeda e já atribui um voto a mesma.

{ name: "CARDANO", "code": "ADA" }

- /vote:
Vota na Criptomoeda que você escolher.


 >>> GET <<<
/coin/:
Retorna todas as suas Criptomoedas e seus votos. Necessário colocar o ID na requisição.




*recomendação do uso da extensão ThunderClient no Visual Studio Code, para uma interação mais rápida.
/// Será feito a implementação de RankingSystem
// Será implementado novas funções tais como: - Maior número de votos - Média de votos de todas as moedas listadas.
//Utilizei o canal do youtube (Dev Problems) como referencia e base de estudos.
