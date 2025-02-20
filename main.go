/*
# O que esse código faz?

Esse código cria 10.000 workers (goroutines) que escutam um canal de comunicação (channel), onde recebem valores e os processam.
O main() envia 100.000 números inteiros para o canal, e os workers os processam um por um.

# Multithreading: Como o código usa várias threads?

Go não trabalha diretamente com threads do sistema operacional (SO), mas sim com goroutines, que são muito mais leves do que threads tradicionais.
Aqui, qtdWorkers := 10000 cria 10.000 goroutines, cada uma executando a função worker().
Isso significa que 10.000 tarefas podem ser executadas simultaneamente, aproveitando múltiplos núcleos da CPU.
Cada worker entra em um loop for x := range data, onde aguarda valores do canal e os processa quando chegam.

# Comunicação entre threads: Como o canal (channel) funciona?

O canal (channel) é o meio pelo qual os valores são distribuídos para os workers.
make(chan int) cria um canal de inteiros sem buffer, o que significa que um valor enviado para o canal precisa ser recebido por alguma goroutine
antes que outro valor possa ser enviado (se não, ele bloqueia).

O main() envia números inteiros para o canal:
for i := 0; i <= 100000; i++ {
    channel <- i
}

Como os workers estão escutando esse canal, cada número é recebido por um único worker e processado.
*/

package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	channel := make(chan int)
	qtdWorkers := 10000

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, channel)
	}

	for i := 0; i < 100000; i++ {
		channel <- i
	}
	close(channel)

}
