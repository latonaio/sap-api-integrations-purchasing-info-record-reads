package responses

type ToPurInfoRecdPrcgCndnValidity struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ConditionRecord              string `json:"ConditionRecord"`
			ConditionValidityEndDate     string `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate   string `json:"ConditionValidityStartDate"`
			ConditionApplication         string `json:"ConditionApplication"`
			ConditionType                string `json:"ConditionType"`
			PurgDocOrderQuantityUnit     string `json:"PurgDocOrderQuantityUnit"`
			PurchasingOrganization       string `json:"PurchasingOrganization"`
			PurchasingInfoRecordCategory string `json:"PurchasingInfoRecordCategory"`
			PurchasingInfoRecord         string `json:"PurchasingInfoRecord"`
			Supplier                     string `json:"Supplier"`
			MaterialGroup                string `json:"MaterialGroup"`
			Material                     string `json:"Material"`
			Plant                        string `json:"Plant"`
			ToPurInfoRecdPrcgCndn        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurInfoRecdPrcgCndn"`
		} `json:"results"`
	} `json:"d"`
}
