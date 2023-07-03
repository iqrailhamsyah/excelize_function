package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	excelgetstruct := NewExcelize()
	arrayid := excelgetstruct.AmbilStructTransactionID("C:/Users/User/Downloads/JUNIOSMART-2022-10-04-settlement.xlsx")
	for _, id := range arrayid {
		fmt.Println(id)
	}
	arraypaymentdate := excelgetstruct.AmbilStructPaymentDate("C:/Users/User/Downloads/JUNIOSMART-2022-10-04-settlement.xlsx")
	for _, date := range arraypaymentdate {
		fmt.Println(date)
	}
	arraypaymenttotal := excelgetstruct.AmbilStructPaymentTotal("C:/Users/User/Downloads/JUNIOSMART-2022-10-04-settlement.xlsx")
	for _, total := range arraypaymenttotal {
		fmt.Println(total)
	}
}

func NewExcelize() ExcelizeFunction {
	//membuat objek exelcize
	excelize := new(excelize.File)
	return excelizefunction{
		libs:         excelize,
		libsopenfile: excelizewrap{},
	}
}

type ExcelizeFunction interface {
	AmbilStructTransactionID(filedirectory string) []string
	AmbilStructPaymentDate(filedirectory string) []string
	AmbilStructPaymentTotal(filedirectory string) []string
}

type excelizefunction struct {
	libs         *excelize.File
	libsopenfile excelizewrap
}

type excelizewrap struct {
}

func (e excelizewrap) OpenFileFunction(fileDr string) (*excelize.File, error) {
	return excelize.OpenFile(fileDr)
}

func (e excelizefunction) AmbilStructTransactionID(filedirectory string) []string {
	//Melakukan Pembukaan File Excel
	file, err := e.libsopenfile.OpenFileFunction(filedirectory)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	//Melakukan Pengambilan Baris-baris pada suatu Sheet
	rows := file.GetRows("Worksheet")

	var kumpulan_transactionID []string
	for _, row := range rows {
		//membuat transactionID untuk ditambahkan ke kumpulan_transactionID
		transactionID := row[4]

		//mengecualikan cell kosong dan cell berisi string "Transaction ID"
		if row[4] != "Transaction ID" && row[4] != "" {
			kumpulan_transactionID = append(kumpulan_transactionID, transactionID)
		}
	}

	//melakukan pengembalian array struct id
	return kumpulan_transactionID

}

func (e excelizefunction) AmbilStructPaymentDate(filedirectory string) []string {
	//Melakukan Pembukaan File Excel
	file, err := e.libsopenfile.OpenFileFunction(filedirectory)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	//Melakukan Pengambilan Baris-baris pada suatu Sheet
	rows := file.GetRows("Worksheet")

	var kumpulan_transactionID []string
	for _, row := range rows {
		//membuat transactionID untuk ditambahkan ke kumpulan_transactionID
		transactionID := row[6]

		//mengecualikan cell kosong dan cell berisi string "Transaction ID"
		if row[6] != "Payment Date" && row[6] != "" {
			kumpulan_transactionID = append(kumpulan_transactionID, transactionID)
		}
	}

	//melakukan pengembalian array struct id
	return kumpulan_transactionID

}

func (e excelizefunction) AmbilStructPaymentTotal(filedirectory string) []string {
	//Melakukan Pembukaan File Excel
	file, err := e.libsopenfile.OpenFileFunction(filedirectory)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	//Melakukan Pengambilan Baris-baris pada suatu Sheet
	rows := file.GetRows("Worksheet")

	var kumpulan_transactionID []string
	for _, row := range rows {
		//membuat transactionID untuk ditambahkan ke kumpulan_transactionID
		transactionID := row[8]

		//mengecualikan cell kosong dan cell berisi string "Transaction ID"
		if row[8] != "Payment Total" && row[8] != "" {
			kumpulan_transactionID = append(kumpulan_transactionID, transactionID)
		}
	}

	//melakukan pengembalian array struct id
	return kumpulan_transactionID

}
