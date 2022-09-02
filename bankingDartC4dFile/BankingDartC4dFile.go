package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	banking "github.com/abar64/agentcommon/banking"
)

func BankingDartC4DVer() string {
	return "BankingDartC4DVer v0.0.0.1"
}

var fileIdentifier = "DartC4DFile.exe"
var routingGateway = "656565"

func main() {
	versionPtr := flag.Bool("version", false, "Display modue version numbers")
	flag.Parse()

	if *versionPtr {
		println(BankingDartC4DVer())
		println(banking.FileC4DVer())
		os.Exit(0)
	}

	// Load a TXT file.
	//	f, _ := os.Open("Retail_310522.dart")
	f, _ := os.Open("APBTestData.CSV")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		//		fmt.Println(record)
		//		fmt.Println(len(record))

		rec, _ := strconv.Atoi(record[0])
		//		fmt.Println(strconv.Atoi(record[0]))
		/*var tmp int = rec

		var currentState = banking.CheckState(tmp)
		if currentState == 9999 {
			panic(9)
		}
		*/
		var recordAsJSon string = banking.StringArraytoDartJson(record) //.StringArraytoDartTransactionJson(record)
		fmt.Println(recordAsJSon)

		switch rec {
		case banking.Header_Record:
			//			fmt.Println("Header Record")
			banking.C4D_Header(record)
		case banking.Accepted_Record:
			fmt.Println("Accepted Record")
			banking.C4D_Transaction_Accept(record)
		case banking.Pending_Record:
			banking.C4D_Transaction_Pending(record)
		case banking.Rejected_Record:
			banking.C4D_Transaction_Reject(record)
		case banking.Trailer_Record:
			//C4D_Transaction_Trailer(record)
		default:
			fmt.Printf("Don't know type %T\n", record[0])
		}
	}
}

/*
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
*/
