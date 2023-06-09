package tickets

import (
	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
	"vuelos/reader"
)
///////////////////////////////////Ej1

func GetTotalTickets(ch chan string , id string) (int, error) {
	
	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)}
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
					total += 1}		
				}
	}
		if total == 0 {
			ch <- fmt.Sprintf("Sin pasajeros pa el vuelo: %s", id)}
		
			if total >= 1 {
		ch <- fmt.Sprintf("Actualmente hay %.0f pasajeros con destino a %s.", total, id)
	}
return int(total), nil
}

/////////////////////////////////////////////////////////////////////// Ej3
func GetTotalPorDia(ch chan string, id string) (int, error) {
		
	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)	}
	data := strings.Split(string(rawData), "\n")
//fmt.Println(data)

var total float64
	//var ListaTotal []reader.Ticket
	for i, v := range data {
		linea := strings.Split(v, ",")
		if i > 0  {
			if linea[3] != id{
				if err != nil {
					fmt.Println("Debe ser un vuelo de LAPA del dia")
			continue
			}}
				if linea[3] == id {
					total += 1}
						}
	}
var	porcentaje = (float64(total) * 100.0 )/ float64(len(data))
ch <- fmt.Sprintf("El porcentaje de vuelos con destino a %s es de %.2f%%.", id, porcentaje)

return int(porcentaje), nil
}

//////////////////////////////////Ej2
const (
	madrugada = "madrugada"
	maniana= "maniana"
	tarde = "tarde"
	noite = "noite"
)

func GetCountbyPeriod( periodo string ) (int, error) {
	var totalTime int
	switch periodo {
	
	case madrugada:
			
		for _, ticket := range reader.ListaDeTicketes {
			horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
			if err != nil {
			panic("No pude splitear o transformar a int madrugada ")}
				if horaM >= 0 && horaM <= 6  {
					totalTime += 1}		}
	return int(totalTime), nil
	
	case maniana:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("No pude splitear o transformar a int maniana ")}
		
			if horaM >= 7 && horaM <= 12  {
				totalTime += 1}		}
return int(totalTime), nil

	case tarde:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("No pude splitear o transformar a int tarde ")}
			if horaM >= 13 && horaM <= 19  {
				totalTime += 1}		}
				return int(totalTime), nil

	case noite:
			
	for _, ticket := range reader.ListaDeTicketes {
		horaM, err := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if err != nil {
		panic("No pude splitear o transformar a int noite ")}
			if horaM >= 20 && horaM <= 23  {
				totalTime += 1}		}
				return int(totalTime), nil}
				return int(totalTime), nil
}