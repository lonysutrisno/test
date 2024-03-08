package script

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/lib/pq"
)

const (
	host     = "34.101.139.27"
	port     = 5432
	user     = "koinworkstech"
	password = "D3v3l0p3rw0rks!@#$"
	dbname   = "asgard"
)

func ExecMambu() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected!")

	srcFile, err := excelize.OpenFile("db_balance.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	startOfRows := 62500
	numberOfRows := 70000

	for i := startOfRows; i <= numberOfRows; i++ {
		// Get the value of cell A1 from the source file
		purchaseRequestID := srcFile.GetCellValue("neo_card", "A"+strconv.Itoa(i))
		dbBalance := srcFile.GetCellValue("neo_card", "B"+strconv.Itoa(i))
		fmt.Printf("\n purchaseRequestID: %s \n", purchaseRequestID)
		fmt.Printf("dbBalance : %s \n", dbBalance)

		// save to db
		updateStmt := `UPDATE koinneo.temp_mambu_feb SET db_balance=$1, modified_at=now(), modified_by=$2 WHERE purchase_request_id=$3`
		_, err := db.Exec(updateStmt, dbBalance, "lony", purchaseRequestID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n Finish UPDATE ROW %d \n", i)
	}

	fmt.Println("Finish Exec!")
}
