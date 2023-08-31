package script

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var query string = `
select
     sum(vcmt.amount),vcmt.user_id, vc.purchase_request_id
     from 
     koinneo.virtual_cards_mastercard_transactions vcmt
 inner join
     koinneo.virtual_cards vc
 on
     vcmt.source = vc.xid
 or
     vcmt.target = vc.xid
 where
     (
         vc.parent is not null
         and
         vcmt.description = 'Topup NEO Card'
         and
         vcmt.transaction_type = 'Debit/Money Out'
     ) != TRUE
 and
     (
         vc.parent is not null
         and
         vcmt.description = 'Withdrawal NEO Card'
         and
         vcmt.transaction_type = 'Credit/Money In'
     ) != TRUE
 and
     vc.purchase_request_id in(
     
     )
 and
     (vcmt.is_active = '1'::bit and vcmt.deleted_at is null)
 and
     vcmt.status = 'success'
 and
     vcmt.created_at < '2023-05-01'
    group by vcmt.user_id ,vc.purchase_request_id ;`

type Resp struct {
	Sum               float64
	UserID            int64
	PurchaseRequestID int64
}

func ExecuteSQL() {
	var pid []string
	var Resps []Resp
	pid = []string{
		"119153991",
	}
	for i := 0; i < len(pid); i++ {
		r := Resp{}
		host := "34.101.44.156"
		port := "5432"
		user := "ro_lony_sutrisno"
		password := "GLUymSFBgzcELk4V"
		dbname := "koinneo"
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		conn, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Tidak Konek DB Errornya : %s", err)
		}
		query :=
			`select
				sum(vcmt.amount),vcmt.user_id, vc.purchase_request_id
				from 
				koinneo.virtual_cards_mastercard_transactions vcmt
			inner join
				koinneo.virtual_cards vc
			on
				vcmt.source = vc.xid
			or
				vcmt.target = vc.xid
			where
				(
					vc.parent is not null
					and
					vcmt.description = 'Topup NEO Card'
					and
					vcmt.transaction_type = 'Debit/Money Out'
				) != TRUE
			and
				(
					vc.parent is not null
					and
					vcmt.description = 'Withdrawal NEO Card'
					and
					vcmt.transaction_type = 'Credit/Money In'
				) != TRUE
			and
				vc.purchase_request_id in(
				` + pid[i] + `
				)
			and
				(vcmt.is_active = '1'::bit and vcmt.deleted_at is null)
			and
				vcmt.status = 'success'
			and
				vcmt.created_at < '2023-05-01'
				group by vcmt.user_id ,vc.purchase_request_id ;`

		row, err := conn.Query(query)
		if err != nil {
			fmt.Println(err)
		}
		for row.Next() {
			err = row.Scan(&r.Sum, &r.UserID, &r.PurchaseRequestID)
			if err != nil {
				fmt.Println(err)
			}
			Resps = append(Resps, r)

		}
		conn.Close()
	}
	fmt.Println(Resps)
}

// /
