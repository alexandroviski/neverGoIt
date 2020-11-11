package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func main() {

	sliceOfString, err := lerTexto("texto.txt")
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Join(sliceOfString, "\n")

	sb := []byte(s)

	SendEmail("<email de destino>", sb)

}

func lerTexto(pathFile string) ([]string, error) {
	// abre o arquivo
	arquivo, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	// garante que o arquivo vai ser fechado apos o uso
	defer arquivo.Close()

	// cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

// funcao qe escreve um texto no arquivo e retorna erro caso tenha algum problema

// funcao pra enviar o email usando api do google
func SendEmail(userEmail string, msg []byte) error {
	auth := smtp.PlainAuth("", "<seu gmail>", "<senha do email ou senha pro app>", "smtp.gmail.com")
	// Here we do it all: connect to our server, set up a message and send it
	to := []string{userEmail}
	err := smtp.SendMail("smtp.gmail.com:587", auth, "<seu gmail de novo>", to, msg)
	if err != nil {
		return err
	}
	return nil
}
