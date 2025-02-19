/*
# O que são os direcionadores de canais?
Em Go, os direcionadores de canais são usados para especificar a direção do fluxo de dados.
Ou seja, você pode definir se um canal pode ser usado para enviar dados, para receber dados, ou para ambos.
Essa direção pode ser explicitada na assinatura da função ou do tipo do canal.

chan<- T: canal somente para envio (somente escrever no canal).
<-chan T: canal somente para recebimento (somente ler do canal).
Esses direcionadores ajudam a garantir a segurança do programa e a evitar erros de concorrência,
pois asseguram que o canal será usado de forma apropriada, sem permitir que dados sejam enviados ou recebidos de maneira errada.

# chan<- string:
O direcionador chan<- na assinatura da função pushToChannel indica que o canal pode apenas ser usado para envio de dados.
Ou seja, essa função só pode enviar (escrever) dados para o canal, não pode ler do canal.

# <-chan string
O direcionador <-chan na assinatura da função pullFromChannel indica que o canal pode apenas ser usado para recebimento de dados.
Ou seja, essa função só pode ler (receber) dados do canal, não pode enviar dados para o canal.

Usar direcionadores de canais ajuda a organizar o fluxo de dados, evitar erros de concorrência e tornar o código mais seguro e legível.
*/

package main

func main() {
	channel := make(chan string)
	go pushToChannel("Hello World", channel)
	pullFromChannel(channel)
}

func pushToChannel(msg string, channel chan<- string) {
	channel <- msg
}

func pullFromChannel(channel <-chan string) {
	println(<-channel)
}
