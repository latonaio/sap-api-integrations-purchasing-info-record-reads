package responses

type General struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
		PurchasingInfoRecord        string `json:"PurchasingInfoRecord"`
		Supplier                    string `json:"Supplier"`
		Material                    string `json:"Material"`
		MaterialGroup               string `json:"MaterialGroup"`
		PurgDocOrderQuantityUnit    string `json:"PurgDocOrderQuantityUnit"`
		SupplierMaterialNumber      string `json:"SupplierMaterialNumber"`
		SupplierRespSalesPersonName string `json:"SupplierRespSalesPersonName"`
		SupplierPhoneNumber         string `json:"SupplierPhoneNumber"`
		SupplierMaterialGroup       string `json:"SupplierMaterialGroup"`
		IsRegularSupplier           bool   `json:"IsRegularSupplier"`
		AvailabilityStartDate       string `json:"AvailabilityStartDate"`
		AvailabilityEndDate         string `json:"AvailabilityEndDate"`
		Manufacturer                string `json:"Manufacturer"`
		CreationDate                string `json:"CreationDate"`
		PurchasingInfoRecordDesc    string `json:"PurchasingInfoRecordDesc"`
		LastChangeDateTime          string `json:"LastChangeDateTime"`
		IsDeleted                   bool   `json:"IsDeleted"`
		ToPurgInfoRecdOrgPlantData  struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurgInfoRecdOrgPlantData"`
		} `json:"results"`
	} `json:"d"`
}
