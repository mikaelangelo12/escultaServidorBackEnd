package main
import "fmt"

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
			fmt.Println("Rodando 1")
		case 2:
			fmt.Println("Rodando 2")
		case 3:
			fmt.Println("Rodando 3")
		}

}