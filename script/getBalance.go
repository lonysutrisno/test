package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type VCN struct {
	Status  int `json:"status"`
	Message struct {
		En string `json:"en"`
		ID string `json:"id"`
	} `json:"message"`
	Data struct {
		PurchaseRequestID              int    `json:"PurchaseRequestID"`
		RequestStatus                  string `json:"RequestStatus"`
		PurchaseRequestTemplateDetails struct {
			TemplateID              int `json:"TemplateId"`
			FullTemplateRuleDetails []struct {
				RuleName        string `json:"RuleName"`
				RuleType        string `json:"RuleType"`
				TemplateControl []struct {
					MinAmount          int       `json:"MinAmount,omitempty"`
					MaxAmount          int       `json:"MaxAmount,omitempty"`
					StrictPreAuthCheck bool      `json:"StrictPreAuthCheck,omitempty"`
					Negate             bool      `json:"Negate,omitempty"`
					CurrencyType       string    `json:"CurrencyType,omitempty"`
					CurrencyCode       any       `json:"CurrencyCode,omitempty"`
					From               time.Time `json:"From,omitempty"`
					To                 time.Time `json:"To,omitempty"`
					TimeZone           string    `json:"TimeZone,omitempty"`
					MaxTrans           int       `json:"MaxTrans,omitempty"`
					CumulativeLimit    float64   `json:"CumulativeLimit,omitempty"`
					Period             string    `json:"Period,omitempty"`
					AvailableBalance   float64   `json:"AvailableBalance,omitempty"`
					EndDate            any       `json:"EndDate,omitempty"`
				} `json:"TemplateControl"`
			} `json:"FullTemplateRuleDetails"`
			TemplateCustomField struct {
				PurchaseType string `json:"Purchase Type"`
				KwUserID     string `json:"kw_user_id"`
				KwUserName   string `json:"kw_user_name"`
			} `json:"TemplateCustomField"`
		} `json:"PurchaseRequestTemplateDetails"`
		VcnInformation struct {
			ID            int    `json:"ID"`
			Pan           string `json:"Pan"`
			Expiry        string `json:"Expiry"`
			Avv           string `json:"Avv"`
			Status        string `json:"Status"`
			EVCNIndicator bool   `json:"EVCNIndicator"`
		} `json:"VcnInformation"`
		Addenda         any    `json:"Addenda"`
		CardImage       string `json:"CardImage"`
		SupplierDetails struct {
			SupplierID     int    `json:"SupplierId"`
			SupplierEmails string `json:"SupplierEmails"`
			NotifySupplier bool   `json:"NotifySupplier"`
		} `json:"SupplierDetails"`
		RawRequest     string `json:"RawRequest"`
		RawResponse    string `json:"RawResponse"`
		Pin            string `json:"Pin"`
		Name           string `json:"Name"`
		CardThemeKey   string `json:"CardThemeKey"`
		CardStatus     string `json:"CardStatus"`
		FlagCardStatus string `json:"FlagCardStatus"`
		IsWithdrawable bool   `json:"IsWithdrawable"`
	} `json:"data"`
}

func ExecuteGetVCN() {
	var pid []string
	pid = []string{
		"86312454",
		"123282465",
		"119459941",
		"101586302",
	}
	for i := 0; i < len(pid); i++ {

		client := &http.Client{}
		var jsonStr = []byte(`{"purchase_request_id":` + pid[i] + `}`)
		req, err := http.NewRequest("POST", "https://asgard.koinworks.com/v1/koinneo/bo/get-details-vcn/debug", bytes.NewBuffer(jsonStr))
		if err != nil {
			fmt.Print(err.Error())
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6MTg2LCJlbWFpbCI6Imxvbnkuc3V0cmlzbm9Aa29pbndvcmtzLmNvbSIsInJvbGUiOiJTdXBlciBBZG1pbiIsIm1lbnVzIjpbeyJpZCI6MTUsIm5hbWUiOiJBY3Rpdml0eSBMb2ciLCJyb2xlcyI6bnVsbCwicm91dGUiOiIvYm8vdmNuL2FjdGl2aXR5LyJ9LHsiaWQiOjQxLCJuYW1lIjoiVkNOIERhc2hib2FyZCAoVmlldyBEZXRhaWwpIiwicm9sZXMiOm51bGwsInJvdXRlIjoiL2JvL3Zjbi92Y24tZGFzaGJvYXJkL2RldGFpbCJ9LHsiaWQiOjM5LCJuYW1lIjoiVkNOIERhc2hib2FyZCIsInJvbGVzIjpudWxsLCJyb3V0ZSI6Ii9iby92Y24vdmNuLWRhc2hib2FyZCJ9LHsiaWQiOjQzLCJuYW1lIjoiVkNOIEJhbGFuY2UgQWRqdXN0bWVudCAoQ3JlYXRpb24pIiwicm9sZXMiOm51bGwsInJvdXRlIjoiL2JvL3Zjbi9iYWxhbmNlLWFkanVzdG1lbnQvY3JlYXRpb24ifSx7ImlkIjo0NiwibmFtZSI6IlZDTiBCYWxhbmNlIEFkanVzdG1lbnQgKFVwZGF0ZSkiLCJyb2xlcyI6bnVsbCwicm91dGUiOiIvYm8vdmNuL2JhbGFuY2UtYWRqdXN0bWVudC9wdXJjaGFzZS1yZXF1ZXN0In0seyJpZCI6NDgsIm5hbWUiOiJWQ04gbWVudSIsInJvbGVzIjpudWxsLCJyb3V0ZSI6Ii9iby92Y24ifSx7ImlkIjo1MywibmFtZSI6IlZDTiBCYWxhbmNlIEFkanVzdG1lbnQiLCJyb2xlcyI6bnVsbCwicm91dGUiOiIvYm8vdmNuL2JhbGFuY2UtYWRqdXN0bWVudCJ9LHsiaWQiOjc2LCJuYW1lIjoiVXNlciBDb3JlIERhdGEiLCJyb2xlcyI6bnVsbCwicm91dGUiOiIvYm8vdXNlcnMifSx7ImlkIjo5NywibmFtZSI6IldoaXRlbGlzdCBUcmFuc2ZlciBVc2VyIiwicm9sZXMiOm51bGwsInJvdXRlIjoiL2JvL2dldC1kb21lc3RpY3QtdHJhbnNmZXItdW5saW1pdGVkLXVzZXJzIn0seyJpZCI6MTAyLCJuYW1lIjoiV2hpdGVsaXN0IFRyYW5zZmVyIFVzZXIgKEFkZCkiLCJyb2xlcyI6bnVsbCwicm91dGUiOiIvYm8vY3JlYXRlLWRvbWVzdGljdC10cmFuc2Zlci11bmxpbWl0ZWQtdXNlciJ9LHsiaWQiOjE0MSwibmFtZSI6Ik1hbnVhbCBBcHByb3ZhbCBEb3dubG9hZCAoRG9tZXN0aWMgVHJhbnNmZXIpIiwicm9sZXMiOm51bGwsInJvdXRlIjoiL2JvL3RyYW5zZmVyL2xpc3QvZG93bmxvYWQifV0sImV4cCI6MTY4MDA2NzczNCwiaXNzIjoiYXNnYXJkLWtvaW5uZW8tYmFja29mZmljZSJ9.ZJSI66sKMpDC4Zt4dJAPHdZsGi0Eyc1_4uQ9txjN0a4")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var responseObject VCN
		json.Unmarshal(bodyBytes, &responseObject)
		fmt.Printf("pid %+v : %+v\n", pid[i], responseObject.Data.PurchaseRequestTemplateDetails.FullTemplateRuleDetails[0].TemplateControl)
	}

}
