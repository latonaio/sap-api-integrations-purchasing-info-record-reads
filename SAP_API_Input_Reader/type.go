package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	PurchasingInfoRecord struct {
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
		PurchasingOrganizationPlant struct {
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
			PricingCondition               struct {
				ConditionRecord            string `json:"ConditionRecord"`
				PurchasingOrganization     string `json:"PurchasingOrganization"`
				Supplier                   string `json:"Supplier"`
				Material                   string `json:"Material"`
				MaterialGroup              string `json:"MaterialGroup"`
				Plant                      string `json:"Plant"`
				ConditionValidityEndDate   string `json:"ConditionValidityEndDate"`
				ConditionValidityStartDate string `json:"ConditionValidityStartDate"`
				ConditionApplication       string `json:"ConditionApplication"`
				ConditionType              string `json:"ConditionType"`
				ConditionScaleQuantity     string `json:"ConditionScaleQuantity"`
				ConditionScaleQuantityUnit string `json:"ConditionScaleQuantityUnit"`
				ConditionScaleAmount       string `json:"ConditionScaleAmount"`
				ConditionRateValue         string `json:"ConditionRateValue"`
				ConditionQuantity          string `json:"ConditionQuantity"`
			} `json:"PricingCondition"`
		} `json:"PurchasingOrganizationPlant"`
	} `json:"PurchasingInfoRecord"`
	APISchema  string   `json:"api_schema"`
	Accepter   []string `json:"accepter"`
	InfoRecord string   `json:"purchasing_info_record"`
	Deleted    bool     `json:"deleted"`
}
