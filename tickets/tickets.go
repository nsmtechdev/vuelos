package tickets

import (
	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
	"vuelos/reader"
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



func GetMornings(time string) (int, error) {
	/*rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	 
	data := strings.Split(string(rawData), "\n")
*/

//fmt.Println(data)
	var total float64
	
	//for i, v := range data {
		//linea := strings.Split(v, ",")
		//horaM := linea[4]
		//linea, err := strconv.Atoi(strings.Split(horaM, ":"))

		for _, ticket := range reader.ListaDeTicketes {
			horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
			panic("could not split")}

		if horaM > 0  {
				if err != nil {
					fmt.Println("error en el timeing")
			continue
			}
				if horaM >= 0 && horaM <= 6 {
					total += 1

				}
	}
	
}

return int(total), nil

}

///////////////////////////////
