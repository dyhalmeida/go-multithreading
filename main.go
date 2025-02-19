/*
# O que é o problema de concorrência neste caso?
O problema de concorrência no código é a interação entre duas goroutines:
a goroutine principal (main()) que está lendo do canal e a goroutine publish que está escrevendo no canal.
Em sistemas concorrentes, o problema de concorrência ocorre quando múltiplas goroutines tentam acessar
e modificar o mesmo recurso (neste caso, o canal) ao mesmo tempo de forma descoordenada,
podendo levar a resultados inconsistentes ou bloqueios (deadlock).

# Como o código resolve isso?
O código utiliza goroutines e canais para coordenar a comunicação entre elas de forma eficiente e segura

# Controle de fluxo com goroutines:
- A goroutine main aguarda os dados enviados pelo canal e imprime cada valor.
- A goroutine publish envia valores para o canal e, uma vez que todos os dados são enviados, fecha o canal.
Isso sinaliza para a goroutine main que não haverá mais dados.

# Uso do range no canal
No código, o loop for ch := range channel da goroutine main é uma maneira de consumir os dados do canal até que o canal seja fechado.
O range em Go cuida automaticamente da sincronização, interrompendo a iteração quando o canal é fechado e não há mais dados.

# Evitando de Deadlock
Deadlock ocorre quando duas ou mais goroutines ficam esperando uma pela outra indefinidamente, sem que nenhuma delas consiga prosseguir, causando um bloqueio no programa.
No caso deste código, o deadlock seria um problema se a goroutine principal ficasse esperando para ler do canal enquanto a goroutine publish nunca escrevesse nada ou não fechasse o canal.
A goroutine publish envia números para o canal de 0 a 10 e, em seguida, fecha o canal.
O fechamento do canal é crucial. Ele informa à goroutine principal que todos os dados foram enviados e que não há mais dados a serem lidos, permitindo que o loop range na goroutine principal termine de maneira segura.
*/

package main

func main() {
	channel := make(chan int)
	go publish(channel)

	for ch := range channel {
		println(ch)
	}
}

func publish(channel chan int) {
	for i := 0; i <= 10; i++ {
		channel <- i
	}
	close(channel)
}
