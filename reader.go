/*package reader

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(){
var TicketList []X.ticket 

rawData, er := os.ReadFile("./tickets.csv")
if err != nil{
	panic(err/Error())
}

err = json.Unmarshal(rawData, TicketList)
if err != nil{
	panic(err)
}

fmt.Println(TicketList)

for _, row := range rawData {
	ticketList = append(ticketList, X.{
			ID: row[0],
			Name: row[1], 
			Mail: row[2],
			Country: row[3],


	})
} 
/*type Ticket struct {
	ID      string
	Name    string
	Mail    string
	Country string
	Time    string
	Price  string
}
}*/