package main

import (
	"fmt"
	"vuelos/tickets"
	//"log"
	"vuelos/reader"
//	"math"
	
)

func main() {

/////////////////////////
ch := make(chan string)
//x := 1.5
//	error := math.Erf(x)
//	fmt.Println(error)


	 reader.ReaderTickets()
	///////////////////////////////////////
	
	/*go totalTicketes, err := tickets.GetTotalTickets(ch, "Brazil")
	 if err != nil {
		log.Fatal("Vuelo de LAPA")
	}*/

	//go tickets.GetTotalTickets(ch, "Brazil")
	
/*
	fmt.Println("el total de viajes al pais es", totalTicketes)	   
	//////////////////////////////
	
	 totalTicketes1, err := tickets.GetTotalPorDia("Brazil")
	if err != nil {
	   log.Fatal("Vuelo de LAPA")
   }
 
   fmt.Println("el total de viajes al pais es", totalTicketes1)
*/


// go tickets.GetTotalPorDia(ch, "Brazil")
   	
 
//////////////////////////////////////////////////
    /*
   totalTime, err := tickets.GetCountbyPeriod("maniana")
   if err != nil {
	  log.Fatal("error en el horario")
   }
  fmt.Println("el total de viajes en la fraja horario es", totalTime)
 */
 /*
  fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
*/
tickets.GetCountbyPeriod(ch, "noite")
 
fmt.Println(<-ch)
//fmt.Println(<-ch)
//fmt.Println(<-ch)
} 
   


