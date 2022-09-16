package sap_api_output_formatter

type PurchasingInfoRecordReads struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	APISchema            string `json:"api_schema"`
	PurchasingInfoRecord string `json:"purchasing_info_record"`
	Deleted              bool   `json:"deleted"`
}

type General struct {
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
	ToPurgInfoRecdOrgPlantData  string `json:"to_PurgInfoRecdOrgPlantData"`
}

type PurchasingOrganizationPlant struct {
	PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
	PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	Plant                          string `json:"Plant"`
	Supplier                       string `json:"Supplier"`
	Material                       string `json:"Material"`
	MaterialGroup                  string `json:"MaterialGroup"`
	PurgDocOrderQuantityUnit       string `json:"PurgDocOrderQuantityUnit"`
	PurchasingGroup                string `json:"PurchasingGroup"`
	MinimumPurchaseOrderQuantity   string `json:"MinimumPurchaseOrderQuantity"`
	StandardPurchaseOrderQuantity  string `json:"StandardPurchaseOrderQuantity"`
	MaterialPlannedDeliveryDurn    string `json:"MaterialPlannedDeliveryDurn"`
	OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
	UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
	UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
	LastReferencingPurchaseOrder   string `json:"LastReferencingPurchaseOrder"`
	LastReferencingPurOrderItem    string `json:"LastReferencingPurOrderItem"`
	NetPriceAmount                 string `json:"NetPriceAmount"`
	MaterialPriceUnitQty           string `json:"MaterialPriceUnitQty"`
	PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
	PriceValidityEndDate           string `json:"PriceValidityEndDate"`
	InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
	TaxCode                        string `json:"TaxCode"`
	IncotermsClassification        string `json:"IncotermsClassification"`
	MaximumOrderQuantity           string `json:"MaximumOrderQuantity"`
	IsRelevantForAutomSrcg         string `json:"IsRelevantForAutomSrcg"`
	IsEvaluatedRcptSettlmtAllowed  bool   `json:"IsEvaluatedRcptSettlmtAllowed"`
	IsPurOrderAllwdForInbDeliv     bool   `json:"IsPurOrderAllwdForInbDeliv"`
	IsOrderAcknRqd                 bool   `json:"IsOrderAcknRqd"`
	IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
	ToPurInfoRecdPrcgCndnValidity  string `json:"to_PurInfoRecdPrcgCndnValidity"`
}

type ToPurgInfoRecdOrgPlant struct {
	PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
	PurchasingInfoRecordCategory   string `json:"PurchasingInfoRecordCategory"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	Plant                          string `json:"Plant"`
	Supplier                       string `json:"Supplier"`
	Material                       string `json:"Material"`
	MaterialGroup                  string `json:"MaterialGroup"`
	PurgDocOrderQuantityUnit       string `json:"PurgDocOrderQuantityUnit"`
	PurchasingGroup                string `json:"PurchasingGroup"`
	MinimumPurchaseOrderQuantity   string `json:"MinimumPurchaseOrderQuantity"`
	StandardPurchaseOrderQuantity  string `json:"StandardPurchaseOrderQuantity"`
	MaterialPlannedDeliveryDurn    string `json:"MaterialPlannedDeliveryDurn"`
	OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
	UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
	UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
	LastReferencingPurchaseOrder   string `json:"LastReferencingPurchaseOrder"`
	LastReferencingPurOrderItem    string `json:"LastReferencingPurOrderItem"`
	NetPriceAmount                 string `json:"NetPriceAmount"`
	MaterialPriceUnitQty           string `json:"MaterialPriceUnitQty"`
	PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
	PriceValidityEndDate           string `json:"PriceValidityEndDate"`
	InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
	TaxCode                        string `json:"TaxCode"`
	IncotermsClassification        string `json:"IncotermsClassification"`
	MaximumOrderQuantity           string `json:"MaximumOrderQuantity"`
	IsRelevantForAutomSrcg         string `json:"IsRelevantForAutomSrcg"`
	IsEvaluatedRcptSettlmtAllowed  bool   `json:"IsEvaluatedRcptSettlmtAllowed"`
	IsPurOrderAllwdForInbDeliv     bool   `json:"IsPurOrderAllwdForInbDeliv"`
	IsOrderAcknRqd                 bool   `json:"IsOrderAcknRqd"`
	ToPurInfoRecdPrcgCndnValidity  string `json:"to_PurInfoRecdPrcgCndnValidity"`
}

type ToPurInfoRecdPrcgCndnValidity struct {
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
	ToPurInfoRecdPrcgCndn        string `json:"to_PurInfoRecdPrcgCndn"`
}

type ToPurInfoRecdPrcgCndn struct {
	ConditionRecord              string `json:"ConditionRecord"`
	ConditionSequentialNumber    string `json:"ConditionSequentialNumber"`
	ConditionApplication         string `json:"ConditionApplication"`
	ConditionType                string `json:"ConditionType"`
	ConditionValidityEndDate     string `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate   string `json:"ConditionValidityStartDate"`
	CreatedByUser                string `json:"CreatedByUser"`
	CreationDate                 string `json:"CreationDate"`
	ConditionTextID              string `json:"ConditionTextID"`
	PricingScaleType             string `json:"PricingScaleType"`
	PricingScaleBasis            string `json:"PricingScaleBasis"`
	ConditionScaleQuantity       string `json:"ConditionScaleQuantity"`
	ConditionScaleQuantityUnit   string `json:"ConditionScaleQuantityUnit"`
	ConditionScaleAmount         string `json:"ConditionScaleAmount"`
	ConditionScaleAmountCurrency string `json:"ConditionScaleAmountCurrency"`
	ConditionCalculationType     string `json:"ConditionCalculationType"`
	ConditionRateValue           string `json:"ConditionRateValue"`
	ConditionRateValueUnit       string `json:"ConditionRateValueUnit"`
	ConditionQuantity            string `json:"ConditionQuantity"`
	ConditionQuantityUnit        string `json:"ConditionQuantityUnit"`
	ConditionToBaseQtyNmrtr      string `json:"ConditionToBaseQtyNmrtr"`
	ConditionToBaseQtyDnmntr     string `json:"ConditionToBaseQtyDnmntr"`
	BaseUnit                     string `json:"BaseUnit"`
	ConditionLowerLimit          string `json:"ConditionLowerLimit"`
	ConditionUpperLimit          string `json:"ConditionUpperLimit"`
	ConditionAlternativeCurrency string `json:"ConditionAlternativeCurrency"`
	ConditionExclusion           string `json:"ConditionExclusion"`
	ConditionIsDeleted           bool   `json:"ConditionIsDeleted"`
	AdditionalValueDays          string `json:"AdditionalValueDays"`
	FixedValueDate               string `json:"FixedValueDate"`
	PaymentTerms                 string `json:"PaymentTerms"`
	CndnMaxNumberOfSalesOrders   string `json:"CndnMaxNumberOfSalesOrders"`
	MinimumConditionBasisValue   string `json:"MinimumConditionBasisValue"`
	MaximumConditionBasisValue   string `json:"MaximumConditionBasisValue"`
	MaximumConditionAmount       string `json:"MaximumConditionAmount"`
	IncrementalScale             string `json:"IncrementalScale"`
	PricingScaleLine             string `json:"PricingScaleLine"`
	ConditionReleaseStatus       string `json:"ConditionReleaseStatus"`
}
