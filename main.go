package main

import (
	"fmt"
	"vuelos/tickets"
	"vuelos/reader"
	"time"
	
	
)

func main() {

/////////////////////////
ch := make(chan string)



	 reader.ReaderTickets()
	///////////////////////////////////////
	
	
	go tickets.GetTotalTickets(ch, "Brazil")
	
/////////////////////////////////////////////

 go tickets.GetTotalPorDia(ch, "Brazil")
   	
 
//////////////////////////////////////////////////
   
periodo := "noite"
go func() {
	for {
		
		totalTime, err := tickets.GetCountbyPeriod(periodo)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Total de viaje en la %s es de: %d\n", periodo, totalTime)
		time.Sleep(time.Second)
		
	}
	
}()
fmt.Println(<-ch)
fmt.Println(<-ch)
}