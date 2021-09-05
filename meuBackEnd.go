package main
import "fmt"
import "net/http"

func main(){
	exibeMenu()
	logicaProg()	
}
func exibeMenu(){
	fmt.Println("1 - Esculta Servidor Local")
	fmt.Println("2 - Verificar logs")
	fmt.Println("3 - Servidor Internet")
	fmt.Println("0 - Sair")
}
func retornaDados()  int {
	var retornaComando int
	fmt.Scan(&retornaComando)
	fmt.Println("Você digitou o numero: ", retornaComando)
	return retornaComando
}
func logicaProg(){
	engine := retornaDados()

		switch engine{
		case 1:
			servidorLocal()
		case 2:
			fmt.Println("Rodando 2")
		case 3:
			fmt.Println("Rodando 3")
		}

}
func servidorLocal(){
	for{
		servidor := "https://edmy.com.br"
		engineServer, _ := http.Get(servidor)
		if engineServer.StatusCode == 200 {
			fmt.Println(servidor, "Seu Servidor esta rodando com sucesso!", engineServer.StatusCode)
			fmt.Println("Finalizar o monitoramento Ctrl + C")
		}else {
			fmt.Println("Seu Servidor", servidor, "Não esta funcionando. Status Code: ", engineServer.StatusCode)
		}
	}
}