package main
import "fmt"
import "net/http"
import "os"

func main(){
	exibeMenu()
	logicaProg()	
}
func exibeMenu(){
	fmt.Println("1 - Esculta Servidor Local")
	fmt.Println("2 - Esculta Varios Servidor")
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
			multServerLocal()
		case 3:
			fmt.Println("EmBreve")
			os.Exit(0)
		default:
			os.Exit(-1)
		}

}
func servidorLocal(){
	for{
		servidor := "https://edmy.com.br"
		engineServer, _ := http.Get(servidor)
		if engineServer.StatusCode == 200 {
			fmt.Println("Finalizar o monitoramento Ctrl + C")
			fmt.Println(servidor, "Seu Servidor esta rodando com sucesso!", engineServer.StatusCode)

		}else {
			fmt.Println("Seu Servidor", servidor, "Não esta funcionando. Status Code: ", engineServer.StatusCode)
		}
	}
}
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
	}
}