package tickets

import (
	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
	"vuelos/reader"
	//"errors"
)


func GetTotalTickets(id string) (int, error) {
	
	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	 
	data := strings.Split(string(rawData), "\n")


//fmt.Println(data)
	var total float64
	for i, v := range data {
		linea := strings.Split(v, ",")
		if i > 0  {
			if linea[3] != id{
				if err != nil {
					fmt.Println("Debe ser un vuelo de LAPA")
			continue
			}}
				if linea[3] == id {
					total += 1

		}
}
	
}

return int(total), nil

}

///////////////////////////////////////////////////////////////////////
func GetTotalPorDia(id string) (int, error) {
		
	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	 
	data := strings.Split(string(rawData), "\n")


//fmt.Println(data)
	var total float64
	//var ListaTotal []reader.Ticket
	for i, v := range data {
		linea := strings.Split(v, ",")
		if i > 0  {
			if linea[3] != id{
				if err != nil {
					fmt.Println("Debe ser un vuelo de LAPA")
			continue
			}}
				if linea[3] == id {
					total += 1}

				
			}
	
}


var	porcentaje = (float64(total) * 100.0 )/ float64(len(data))

return int(porcentaje), nil
}





//////////////////////////////////
const (
	madrugada = "madrugada"
	maniana= "maniana"
	tarde = "tarde"
	noite = "noite"
)

func GetCountbyPeriod(tipo string ) (int, error) {
	var totalTime int
	switch tipo {
	case madrugada:
			
		for _, ticket := range reader.ListaDeTicketes {
			horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
			if err != nil {
			panic("no pude splitear o transfomrar a int ")}
				if horaM >= 0 && horaM <= 6  {
					totalTime += 1}			
	}
		
	return int(totalTime), nil
	

	case maniana:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("no pude splitear o transfomrar a int ")}
			if horaM >= 7 && horaM <= 12  {
				totalTime += 1}			
}
return int(totalTime), nil

	case tarde:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("no pude splitear o transfomrar a int ")}
			if horaM >= 13 && horaM <= 19  {
				totalTime += 1}			
}
	
return int(totalTime), nil

	case noite:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("no pude splitear o transfomrar a int ")}
			if horaM >= 20 && horaM <= 23  {
				totalTime += 1}			
		}
	}
return int(totalTime), nil

}

///////////////////////////////
