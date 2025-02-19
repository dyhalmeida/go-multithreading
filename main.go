/*
# O que é Deadlock?
Deadlock ocorre quando duas ou mais goroutines estão esperando por recursos ou condições que nunca vão ocorrer.
No contexto de canais em Go, isso geralmente acontece quando uma goroutine tenta ler de um canal que nunca será escrito
ou quando uma goroutine tenta escrever em um canal que nunca será lido.
Como resultado, ambas as goroutines ficam presas esperando algo que nunca acontece, e o programa não avança.

- Criação do canal
Um canal channel de tipo bool é criado.
Por padrão, canais em Go são bloqueantes: uma operação de leitura (<-channel) ou escrita (channel <- value)
vai bloquear a execução até que a operação seja possível.

- O código está tentando ler um valor do canal. A operação <-channel vai bloquear a execução do programa e
aguardar até que um valor seja enviado para o canal.

## Problema: Não há nenhuma goroutine escrevendo no canal.
O canal está vazio e a operação de leitura fica esperando para sempre por um valor.

# O Deadlock:
- A goroutine principal (no main()) está esperando indefinidamente por um valor do canal, mas nunca vai receber
esse valor porque não há nenhuma goroutine enviando um valor para o canal.

- Isso cria um deadlock, pois o programa fica parado, aguardando algo que nunca vai acontecer, já que não há um envio para o canal.
*/

package main

func main() {
	channel := make(chan bool) // canal criado
	<-channel                  // A operação de leitura espera um valor que nunca será enviado
}
