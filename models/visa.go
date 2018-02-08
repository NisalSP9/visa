package models

type Visa struct {
	VisaID                  int    `db:"th_visa_ID" json:"visaID"`
	VisaCountry             string `db:"th_visa_country" json:"visaCountry"`
	VisaGroupName           string `db:"th_visa_groupname" json:"visaGroupName"`
	VisaGroupNameUpdateTime string `db:"th_visa_groupname_update_time" json:"visaGroupNameUpdateTime"`
	VisaType                string `db:"th_visa_type" json:"visaType"`
	VisaTypeEntryUpdateTime string `db:"th_visa_type_entry_update_time" json:"visaTypeEntryUpdateTime"`
	VisaTypeEntry           string `db:"th_visa_type_entry" json:"visaTypeEntry"`
	VisaAgent               string `db:"th_visa_agent" json:"visaAgent"`
	VisaSubmitByOther       string `db:"th_visa_submit_byother" json:"visaSubmitByOther"`
	VisaQueAppoinnt         string `db:"th_visa_que_appoinnt" json:"visaQueAppoinnt"`
	VisaApprovePeriod       string `db:"th_visa_approve_period" json:"visaApprovePeriod"`
	VisaRealPrice           string `db:"th_visa_real_price" json:"visaRealPrice"`
	VisaOfficialAgentCharge string `db:"th_visa_official_agent_charge" json:"visaOfficialAgentCharge"`
	VisaSupplierCharge      string `db:"th_visa_supplier_charge" json:"visaSupplierCharge"`
	VisaSupplier            string `db:"th_visa_supplier" json:"visaSupplier"`
	VisaSellingPrice        string `db:"th_visa_selling_price" json:"visaSellingPrice"`
	VisaRemark              string `db:"th_visa_remark" json:"visaRemark"`
	VisaUpdateDateChar      string `db:"th_visa_update_date_char" json:"visaUpdateDateChar"`
	VisaUpdateDate          string `db:"th_visa_update_date" json:"visaUpdateDate"`
	VisaUpdateBy            string `db:"th_visa_update_by" json:"visaUpdateBy"`
}
