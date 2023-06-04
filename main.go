package main

import (
	"fmt"
	"vuelos/tickets"
	"log"
	"vuelos/reader"
	
	
)

func main() {

	 reader.ReaderTickets()
	
	
	totalTicketes, err := tickets.GetTotalTickets("Brazil")
	 if err != nil {
		log.Fatal("Vuelo de LAPA")
	}
	fmt.Println("el total de viajes al pais es", totalTicketes)	   
	
	totalTicketes1, err := tickets.GetTotalPorDia("Brazil")
	if err != nil {
	   log.Fatal("Vuelo de LAPA")
   }
   fmt.Println("el total de viajes al pais es", totalTicketes1)

    
   totalTime, err := tickets.GetCountbyPeriod("maniana")
   if err != nil {
	  log.Fatal("error en el horario")
   }
  fmt.Println("el total de viajes en la fraja horario es", totalTime)
 
  

  } 
   


