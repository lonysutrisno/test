package script

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type M map[string]interface{}

func ExecScript() {
	// 	export DB_ENGINE=postgres
	// export DB_HOST=34.101.139.27
	// export DB_PORT=‘5432’
	// export DB_USER=koinworkstech
	// export DB_PWD=“D3v3l0p3rw0rks!@#$”
	// export DB_NAME=“asgard”
	// export DB_SSL_MODE=“disable”
	f, err := os.Open("virtual_cards.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	host := "34.101.139.27"
	port := "5432"
	user := "koinworkstech"
	password := "“D3v3l0p3rw0rks!@#$"
	dbname := "asgard"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	_, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	str := ""

	for i := 0; i < len(data); i++ {
		if data[i][1] != "" {
			str = str + fmt.Sprintf("('%+v',%+v),", data[i][0], data[i][1])

		}
		// if counter == 10000 {

		// 	f, _ := os.Create("./docs/data.txt")
		// 	f.WriteString(script)

		// 	counter = 0
		// 	str = ""
		// }
	}
	str = str[:len(str)-1]
	script := `
			update temp_mambu_feb as t set
				mambu_id = c.mambu_id
			from (values
				` + str + `
			) as c(purchase_request_id, mambu_id)
			where c.purchase_request_id = t.purchase_request_id;
`
	fmt.Println(script)

}
