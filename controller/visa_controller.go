package controller

import (
	"../models"
	"../DBConnection"
)

const GET_ALL_DATA = `SELECT * FROM tbl_th_visa;`

func GetAllData() ([]models.Visa) {

	var visas []models.Visa

	var visaID int
	var visaCountry,
	visaGroupName,
	visaGroupNameUpdateTime,
	visaType,
	visaTypeEntryUpdateTime,
	visaTypeEntry,
	visaAgent,
	visaSubmitByOther,
	visaQueAppoinnt,
	visaApprovePeriod,
	visaRealPrice,
	visaOfficialAgentCharge,
	visaSupplierCharge,
	visaSupplier,
	visaSellingPrice,
	visaRemark,
	visaUpdateDateChar,
	visaUpdateDate,
	visaUpdateBy string

	db := DBConnection.GetConnection()
	rows, err := db.Query(GET_ALL_DATA)
	if err != nil {
		println(err.Error())
	}
	for (rows.Next()) {
		err := rows.Scan(
			&visaID,
			&visaCountry,
			&visaGroupName,
			&visaGroupNameUpdateTime,
			&visaType,
			&visaTypeEntryUpdateTime,
			&visaTypeEntry,
			&visaAgent,
			&visaSubmitByOther,
			&visaQueAppoinnt,
			&visaApprovePeriod,
			&visaRealPrice,
			&visaOfficialAgentCharge,
			&visaSupplierCharge,
			&visaSupplier,
			&visaSellingPrice,
			&visaRemark,
			&visaUpdateDateChar,
			&visaUpdateDate,
			&visaUpdateBy,
		)

		visas = append(visas, models.Visa{
			VisaID:visaID,
			VisaCountry:visaCountry,
			VisaGroupName:visaGroupName,
			VisaGroupNameUpdateTime:visaGroupNameUpdateTime,
			VisaType:visaType,
			VisaTypeEntryUpdateTime:visaTypeEntryUpdateTime,
			VisaTypeEntry:visaTypeEntry,
			VisaAgent:visaAgent,
			VisaSubmitByOther:visaSubmitByOther,
			VisaQueAppoinnt:visaQueAppoinnt,
			VisaApprovePeriod:visaApprovePeriod,
			VisaRealPrice:visaRealPrice,
			VisaOfficialAgentCharge:visaOfficialAgentCharge,
			VisaSupplierCharge:visaSupplierCharge,
			VisaSupplier:visaSupplier,
			VisaSellingPrice:visaSellingPrice,
			VisaRemark:visaRemark,
			VisaUpdateDateChar:visaUpdateDateChar,
			VisaUpdateDate:visaUpdateDate,
			VisaUpdateBy:visaUpdateBy,




		})

		if err != nil {
			println(err.Error())
		}
	}
	return visas

}

const CREATE_DATA  = `INSERT INTO tbl_th_visa (th_visa_country, th_visa_groupname, th_visa_type, th_visa_type_entry, th_visa_agent, th_visa_submit_byother, th_visa_que_appoinnt, th_visa_approve_period, th_visa_real_price, th_visa_official_agent_charge, th_visa_supplier_charge, th_visa_supplier, th_visa_selling_price, th_visa_remark, th_visa_update_date_char, th_visa_update_date, th_visa_update_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

func CreateData(visa models.Visa){

	db := DBConnection.GetConnection()

	Stmt,err := db.Prepare(CREATE_DATA)

	if err != nil {
		println(err.Error())
	}

	defer Stmt.Close()

	_ , err = Stmt.Exec(
		visa.VisaCountry,
		visa.VisaGroupName,
		visa.VisaType,
		visa.VisaTypeEntry,
		visa.VisaAgent,
		visa.VisaSubmitByOther,
		visa.VisaQueAppoinnt,
		visa.VisaApprovePeriod,
		visa.VisaRealPrice,
		visa.VisaOfficialAgentCharge,
		visa.VisaSupplierCharge,
		visa.VisaSupplier,
		visa.VisaSellingPrice,
		visa.VisaRemark,
		visa.VisaUpdateDateChar,
		visa.VisaUpdateDate,
		visa.VisaUpdateBy,
	)

	if err != nil {
		println(err.Error())
	}

}

const DELETE_DATA  = `DELETE FROM tbl_th_visa WHERE th_visa_ID = ? ;`

func DeteleData(id string){

	db := DBConnection.GetConnection()

	Stmt,err := db.Prepare(DELETE_DATA)

	if err != nil{

		println(err.Error())

	}

	defer Stmt.Close()

	_,err = Stmt.Exec(id)

	if err != nil {

		println(err.Error())

	}

}

const UPDATE_DATA = `UPDATE tbl_th_visa SET th_visa_country = ?, th_visa_groupname = ?, th_visa_type = ?, th_visa_type_entry = ?, th_visa_agent = ? , th_visa_submit_byother = ? , th_visa_que_appoinnt = ? , th_visa_approve_period = ? , th_visa_real_price = ?, th_visa_official_agent_charge = ?, th_visa_supplier_charge = ?, th_visa_supplier = ?, th_visa_selling_price = ?, th_visa_remark = ?, th_visa_update_date_char = ?, th_visa_update_date = ?, th_visa_update_by = ? WHERE th_visa_ID = ? ;`

func UpdateData (visa models.Visa) {

	db := DBConnection.GetConnection()

	Stmt,err := db.Prepare(UPDATE_DATA)

	if err != nil {
		println(err.Error())
	}

	defer Stmt.Close()

	_ , err = Stmt.Exec(
		visa.VisaCountry,
		visa.VisaGroupName,
		visa.VisaType,
		visa.VisaTypeEntry,
		visa.VisaAgent,
		visa.VisaSubmitByOther,
		visa.VisaQueAppoinnt,
		visa.VisaApprovePeriod,
		visa.VisaRealPrice,
		visa.VisaOfficialAgentCharge,
		visa.VisaSupplierCharge,
		visa.VisaSupplier,
		visa.VisaSellingPrice,
		visa.VisaRemark,
		visa.VisaUpdateDateChar,
		visa.VisaUpdateDate,
		visa.VisaUpdateBy,
		visa.VisaID,
	)

	if err != nil {
		println(err.Error())
	}

}