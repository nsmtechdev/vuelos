package reader

import(
	"os"
	
	//"fmt"
	"log"
	//"vuelos/tickets"
	"encoding/csv"

)

type Ticket struct {
	ID      string
	Nombre  string
	Mail    string
	Pais	string
	Hora    string
	Precio  string}

var ListaDeTicketes []Ticket

func ReaderTickets(){
rawData, err := os.Open("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	 
	//data := strings.Split(string(rawData), "\n")
	
	data, err:= csv.NewReader(rawData).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	

	for _, linea := range data {
		ListaDeTicketes = append(ListaDeTicketes, Ticket{
		ID:      linea[0],
		Nombre:  linea[1],
		Mail:    linea[2],
		Pais:	 linea[3],
		Hora:    linea[4],
		Precio:  linea[5],})}
		}
		