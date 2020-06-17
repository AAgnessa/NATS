package main
import(
	"fmt"
	"net"
)

func main(){

	conn,err:=net.Dial("tcp","demo.nats.io:4222")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()


for {

	//Создаем сообщение
	var ClientMessage string

	fmt.Print("Enter the message:")
	_ ,err:=fmt.Scanln(&ClientMessage)
	if err != nil {
		fmt.Println("Uncorrect imput", err)
		continue
	}

	//Отправляем сообщение серверу
	if n,err:=conn.Write([] byte(ClientMessage));
	n==0 || err != nil{
		fmt.Println(err)
		return
	}

	//Получаем ответ от сервера

}

}
