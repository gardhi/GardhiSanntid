package main

import(
	."fmt"
	."net"
	"bufio"
	"os"
	"strings"
)


func ListenToNetwork(chanCon chan *UDPConn){
	
	Println("Start UDP server")
	udpAddr, err := ResolveUDPAddr("udp4", ":20666") //resolving
	if err != nil{
		println("ERROR: Resolving error")
	}
	conn, err := ListenUDP("udp", udpAddr) //initiating listening
	if err != nil{
		println("ERROR: Listening error")
	}

   	chanCon <- conn

	for{
		data := make([]byte,1024)
		_, _, err := conn.ReadFromUDP(data) //kan bruke addr til å sjekke hvor melding kommer fra f.eks if addr not [egen i.p]
		if err != nil{
			println("ERROR: while reading")
		}
		//if addr.IP.String() != "129.241.187.148"{
			Println("Homofil_mann_22: ",string(data))
		//}
	}	
}

func SendToNetwork(chanCon chan *UDPConn){
    
    sendAddr, err := ResolveUDPAddr("udp4","129.241.187.157:20666") //Spesifiserer adresse
	//connection,err := DialUDP("udp",nil, sendAddr) //setter opp "socket" for sending
	if err != nil {
		println("ERROR while resolving UDP addr")
	}

	connection := <- chanCon
	reader := bufio.NewReader(os.Stdin)

	for{
		text, _ := reader.ReadString('\n')
		testmsg := []byte(strings.TrimSpace(text))
		if connection ==  nil{
			println("ERROR, connection = nil")
		}	
		connection.WriteToUDP(testmsg, sendAddr)
	}	
}



func main(){

//host, _ := os.Hostname()
//addrs, _ := LookupIP()
//println(addrs)
chanCon := make(chan *UDPConn,1)
go ListenToNetwork(chanCon)
SendToNetwork(chanCon)


}