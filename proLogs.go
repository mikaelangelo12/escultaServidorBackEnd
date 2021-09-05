//copyright © 2021 | By Mikael Angelo
package main
import(
	"fmt"
	"net/http"
	"os"
)

const daley = 5

//Função Principal do Programa
func main(){
	fmt.Println(" ")
	fmt.Println("Bem Vindo ao Programa Pro Logs")
	fmt.Println(" ")
	for{
		exibeMenu()
		logicaProg()	
	}
}
//Função responsavel por imprimir o Menu
func exibeMenu(){
	fmt.Println("[1] - Esculta Servidor Local")
	fmt.Println("[2] - Esculta Varios Servidores")
	fmt.Println("[3] - Servidor Internet")
	fmt.Println("[0] - Sair")		
	fmt.Println(" ")
	fmt.Println("Copyright © 2021 | By Mikael Angelo")
	fmt.Println(" ")
	fmt.Println("Digite Sua Opção Aqui Em Baixo:")
}
//Função responsavel por execultar os dados do Menu
func retornaDados()  int {
	var retornaComando int
	fmt.Scan(&retornaComando)
	fmt.Println("Você digitou o numero: ", retornaComando)
	return retornaComando
}
//Função Responsavel por tomar decisão do Menu
func logicaProg(){
	engine := retornaDados()

		switch engine{
		case 1:
			servidorLocal()
		case 2:
			multServerLocal()
		case 3:
			fmt.Println("EmBreve")
		case 0:
			fmt.Println("Ate mais!")
			os.Exit(0)
		default:
			os.Exit(-1)
		}

}
//Função responsavel acessar o servidor e mostrar os Logs
func servidorLocal(){
	for{		
		//Coloca Aqui o link do seu servidor Local ou servidor em produção
		servidor := "https://edmy.com.br"
		engineServer, _ := http.Get(servidor)

		if engineServer.StatusCode == 200 {
			fmt.Println("Finalizar o monitoramento Ctrl + C")
			fmt.Println(servidor, "Seu Servidor esta rodando com sucesso!", engineServer.StatusCode)

		}else {
			fmt.Println("Seu Servidor", servidor, "Não esta funcionando. Status Code: ", engineServer.StatusCode)
		}
		time.Sleep(delay * time.Second)
	}
}
//Função responsavel por acessar em um array varios servidor e mostra logs de cada um
func multServerLocal(){
	servidores :=[]string{"https://edmy.com.br", "http://brotofm.com.br", "http://mikaelangelo.tech"}
	for{
		for i:=0; i<len(servidores); i++{
			servidoresLocal, _ := http.Get(servidores[i])

			if servidoresLocal.StatusCode == 200 {
				fmt.Println("Finalizar o monitoramento Ctrl + C")				
				fmt.Println(servidores[i], "Seu Servidor esta rodando com sucesso!", servidoresLocal.StatusCode)

			}else {
				fmt.Println("Seu Servidor", servidores[i], "Não esta funcionando. Status Code: ", servidoresLocal.StatusCode)
			}
		}
		time.Sleep(delay * time.Second)
	}
}
