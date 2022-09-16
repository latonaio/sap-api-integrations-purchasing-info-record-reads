package responses

type ToPurInfoRecdPrcgCndn struct {
	D *struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ConditionRecord               string `json:"ConditionRecord"`
		ConditionSequentialNumber     string `json:"ConditionSequentialNumber"`
		ConditionApplication          string `json:"ConditionApplication"`
		ConditionType                 string `json:"ConditionType"`
		ConditionValidityEndDate      string `json:"ConditionValidityEndDate"`
		ConditionValidityStartDate    string `json:"ConditionValidityStartDate"`
		CreatedByUser                 string `json:"CreatedByUser"`
		CreationDate                  string `json:"CreationDate"`
		ConditionTextID               string `json:"ConditionTextID"`
		PricingScaleType              string `json:"PricingScaleType"`
		PricingScaleBasis             string `json:"PricingScaleBasis"`
		ConditionScaleQuantity        string `json:"ConditionScaleQuantity"`
		ConditionScaleQuantityUnit    string `json:"ConditionScaleQuantityUnit"`
		ConditionScaleAmount          string `json:"ConditionScaleAmount"`
		ConditionScaleAmountCurrency  string `json:"ConditionScaleAmountCurrency"`
		ConditionCalculationType      string `json:"ConditionCalculationType"`
		ConditionRateValue            string `json:"ConditionRateValue"`
		ConditionRateValueUnit        string `json:"ConditionRateValueUnit"`
		ConditionQuantity             string `json:"ConditionQuantity"`
		ConditionQuantityUnit         string `json:"ConditionQuantityUnit"`
		ConditionToBaseQtyNmrtr       string `json:"ConditionToBaseQtyNmrtr"`
		ConditionToBaseQtyDnmntr      string `json:"ConditionToBaseQtyDnmntr"`
		BaseUnit                      string `json:"BaseUnit"`
		ConditionLowerLimit           string `json:"ConditionLowerLimit"`
		ConditionUpperLimit           string `json:"ConditionUpperLimit"`
		ConditionAlternativeCurrency  string `json:"ConditionAlternativeCurrency"`
		ConditionExclusion            string `json:"ConditionExclusion"`
		ConditionIsDeleted            bool   `json:"ConditionIsDeleted"`
		AdditionalValueDays           string `json:"AdditionalValueDays"`
		FixedValueDate                string `json:"FixedValueDate"`
		PaymentTerms                  string `json:"PaymentTerms"`
		CndnMaxNumberOfSalesOrders    string `json:"CndnMaxNumberOfSalesOrders"`
		MinimumConditionBasisValue    string `json:"MinimumConditionBasisValue"`
		MaximumConditionBasisValue    string `json:"MaximumConditionBasisValue"`
		MaximumConditionAmount        string `json:"MaximumConditionAmount"`
		IncrementalScale              string `json:"IncrementalScale"`
		PricingScaleLine              string `json:"PricingScaleLine"`
		ConditionReleaseStatus        string `json:"ConditionReleaseStatus"`
		ToPurInfoRecdPrcgCndnValidity struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"to_PurInfoRecdPrcgCndnValidity"`
		ToPurInfoRecdSuplmntPrcgCndn struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"to_PurInfoRecdSuplmntPrcgCndn"`
	} `json:"d"`
}
