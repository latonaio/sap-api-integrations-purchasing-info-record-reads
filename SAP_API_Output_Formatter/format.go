package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) ([]General, error) {
	pm := &responses.General{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	general := make([]General, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		general = append(general, General{
			PurchasingInfoRecord:        data.PurchasingInfoRecord,
			Supplier:                    data.Supplier,
			Material:                    data.Material,
			MaterialGroup:               data.MaterialGroup,
			PurgDocOrderQuantityUnit:    data.PurgDocOrderQuantityUnit,
			SupplierMaterialNumber:      data.SupplierMaterialNumber,
			SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
			SupplierPhoneNumber:         data.SupplierPhoneNumber,
			SupplierMaterialGroup:       data.SupplierMaterialGroup,
			IsRegularSupplier:           data.IsRegularSupplier,
			AvailabilityStartDate:       data.AvailabilityStartDate,
			AvailabilityEndDate:         data.AvailabilityEndDate,
			Manufacturer:                data.Manufacturer,
			CreationDate:                data.CreationDate,
			PurchasingInfoRecordDesc:    data.PurchasingInfoRecordDesc,
			LastChangeDateTime:          data.LastChangeDateTime,
			IsDeleted:                   data.IsDeleted,
			ToPurgInfoRecdOrgPlantData:  data.ToPurgInfoRecdOrgPlantData.Deferred.URI,
		})
	}

	return general, nil
}

func ConvertToPurchasingOrganizationPlant(raw []byte, l *logger.Logger) ([]PurchasingOrganizationPlant, error) {
	pm := &responses.PurchasingOrganizationPlant{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchasingOrganizationPlant. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	purchasingOrganizationPlant := make([]PurchasingOrganizationPlant, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		purchasingOrganizationPlant = append(purchasingOrganizationPlant, PurchasingOrganizationPlant{
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingInfoRecordCategory:   data.PurchasingInfoRecordCategory,
			PurchasingOrganization:         data.PurchasingOrganization,
			Plant:                          data.Plant,
			Supplier:                       data.Supplier,
			Material:                       data.Material,
			MaterialGroup:                  data.MaterialGroup,
			PurgDocOrderQuantityUnit:       data.PurgDocOrderQuantityUnit,
			PurchasingGroup:                data.PurchasingGroup,
			MinimumPurchaseOrderQuantity:   data.MinimumPurchaseOrderQuantity,
			StandardPurchaseOrderQuantity:  data.StandardPurchaseOrderQuantity,
			MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			LastReferencingPurchaseOrder:   data.LastReferencingPurchaseOrder,
			LastReferencingPurOrderItem:    data.LastReferencingPurOrderItem,
			NetPriceAmount:                 data.NetPriceAmount,
			MaterialPriceUnitQty:           data.MaterialPriceUnitQty,
			PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
			PriceValidityEndDate:           data.PriceValidityEndDate,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			TaxCode:                        data.TaxCode,
			IncotermsClassification:        data.IncotermsClassification,
			MaximumOrderQuantity:           data.MaximumOrderQuantity,
			IsRelevantForAutomSrcg:         data.IsRelevantForAutomSrcg,
			IsEvaluatedRcptSettlmtAllowed:  data.IsEvaluatedRcptSettlmtAllowed,
			IsPurOrderAllwdForInbDeliv:     data.IsPurOrderAllwdForInbDeliv,
			IsOrderAcknRqd:                 data.IsOrderAcknRqd,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
			ToPurInfoRecdPrcgCndnValidity:  data.ToPurInfoRecdPrcgCndnValidity.Deferred.URI,
		})
	}

	return purchasingOrganizationPlant, nil
}

func ConvertToToPurgInfoRecdOrgPlant(raw []byte, l *logger.Logger) ([]ToPurgInfoRecdOrgPlant, error) {
	pm := &responses.ToPurgInfoRecdOrgPlantData{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPurgInfoRecdOrgPlant. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toPurgInfoRecdOrgPlant := make([]ToPurgInfoRecdOrgPlant, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPurgInfoRecdOrgPlant = append(toPurgInfoRecdOrgPlant, ToPurgInfoRecdOrgPlant{
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingInfoRecordCategory:   data.PurchasingInfoRecordCategory,
			PurchasingOrganization:         data.PurchasingOrganization,
			Plant:                          data.Plant,
			Supplier:                       data.Supplier,
			Material:                       data.Material,
			MaterialGroup:                  data.MaterialGroup,
			PurgDocOrderQuantityUnit:       data.PurgDocOrderQuantityUnit,
			PurchasingGroup:                data.PurchasingGroup,
			MinimumPurchaseOrderQuantity:   data.MinimumPurchaseOrderQuantity,
			StandardPurchaseOrderQuantity:  data.StandardPurchaseOrderQuantity,
			MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			LastReferencingPurchaseOrder:   data.LastReferencingPurchaseOrder,
			LastReferencingPurOrderItem:    data.LastReferencingPurOrderItem,
			NetPriceAmount:                 data.NetPriceAmount,
			MaterialPriceUnitQty:           data.MaterialPriceUnitQty,
			PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
			PriceValidityEndDate:           data.PriceValidityEndDate,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			TaxCode:                        data.TaxCode,
			IncotermsClassification:        data.IncotermsClassification,
			MaximumOrderQuantity:           data.MaximumOrderQuantity,
			IsRelevantForAutomSrcg:         data.IsRelevantForAutomSrcg,
			IsEvaluatedRcptSettlmtAllowed:  data.IsEvaluatedRcptSettlmtAllowed,
			IsPurOrderAllwdForInbDeliv:     data.IsPurOrderAllwdForInbDeliv,
			IsOrderAcknRqd:                 data.IsOrderAcknRqd,
			ToPurInfoRecdPrcgCndnValidity:  data.ToPurInfoRecdPrcgCndnValidity.Deferred.URI,
		})
	}

	return toPurgInfoRecdOrgPlant, nil
}

func ConvertToToPurInfoRecdPrcgCndnValidity(raw []byte, l *logger.Logger) ([]ToPurInfoRecdPrcgCndnValidity, error) {
	pm := &responses.ToPurInfoRecdPrcgCndnValidity{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPurInfoRecdPrcgCndnValidity. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toPurInfoRecdPrcgCndnValidity := make([]ToPurInfoRecdPrcgCndnValidity, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPurInfoRecdPrcgCndnValidity = append(toPurInfoRecdPrcgCndnValidity, ToPurInfoRecdPrcgCndnValidity{
			ConditionRecord:              data.ConditionRecord,
			ConditionValidityEndDate:     data.ConditionValidityEndDate,
			ConditionValidityStartDate:   data.ConditionValidityStartDate,
			ConditionApplication:         data.ConditionApplication,
			ConditionType:                data.ConditionType,
			PurgDocOrderQuantityUnit:     data.PurgDocOrderQuantityUnit,
			PurchasingOrganization:       data.PurchasingOrganization,
			PurchasingInfoRecordCategory: data.PurchasingInfoRecordCategory,
			PurchasingInfoRecord:         data.PurchasingInfoRecord,
			Supplier:                     data.Supplier,
			MaterialGroup:                data.MaterialGroup,
			Material:                     data.Material,
			Plant:                        data.Plant,
			ToPurInfoRecdPrcgCndn:        data.ToPurInfoRecdPrcgCndn.Deferred.URI,
		})
	}

	return toPurInfoRecdPrcgCndnValidity, nil
}

func ConvertToToPurInfoRecdPrcgCndn(raw []byte, l *logger.Logger) (*ToPurInfoRecdPrcgCndn, error) {
	pm := &responses.ToPurInfoRecdPrcgCndn{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPurInfoRecdPrcgCndnValidity. unmarshal error: %w", err)
	}
	if pm.D == nil {
		return nil, xerrors.New("Result data is not exist")
	}
	data := pm.D

	return &ToPurInfoRecdPrcgCndn{
		ConditionRecord:              data.ConditionRecord,
		ConditionSequentialNumber:    data.ConditionSequentialNumber,
		ConditionApplication:         data.ConditionApplication,
		ConditionType:                data.ConditionType,
		ConditionValidityEndDate:     data.ConditionValidityEndDate,
		ConditionValidityStartDate:   data.ConditionValidityStartDate,
		CreatedByUser:                data.CreatedByUser,
		CreationDate:                 data.CreationDate,
		ConditionTextID:              data.ConditionTextID,
		PricingScaleType:             data.PricingScaleType,
		PricingScaleBasis:            data.PricingScaleBasis,
		ConditionScaleQuantity:       data.ConditionScaleQuantity,
		ConditionScaleQuantityUnit:   data.ConditionScaleQuantityUnit,
		ConditionScaleAmount:         data.ConditionScaleAmount,
		ConditionScaleAmountCurrency: data.ConditionScaleAmountCurrency,
		ConditionCalculationType:     data.ConditionCalculationType,
		ConditionRateValue:           data.ConditionRateValue,
		ConditionRateValueUnit:       data.ConditionRateValueUnit,
		ConditionQuantity:            data.ConditionQuantity,
		ConditionQuantityUnit:        data.ConditionQuantityUnit,
		ConditionToBaseQtyNmrtr:      data.ConditionToBaseQtyNmrtr,
		ConditionToBaseQtyDnmntr:     data.ConditionToBaseQtyDnmntr,
		BaseUnit:                     data.BaseUnit,
		ConditionLowerLimit:          data.ConditionLowerLimit,
		ConditionUpperLimit:          data.ConditionUpperLimit,
		ConditionAlternativeCurrency: data.ConditionAlternativeCurrency,
		ConditionExclusion:           data.ConditionExclusion,
		ConditionIsDeleted:           data.ConditionIsDeleted,
		AdditionalValueDays:          data.AdditionalValueDays,
		FixedValueDate:               data.FixedValueDate,
		PaymentTerms:                 data.PaymentTerms,
		CndnMaxNumberOfSalesOrders:   data.CndnMaxNumberOfSalesOrders,
		MinimumConditionBasisValue:   data.MinimumConditionBasisValue,
		MaximumConditionBasisValue:   data.MaximumConditionBasisValue,
		MaximumConditionAmount:       data.MaximumConditionAmount,
		IncrementalScale:             data.IncrementalScale,
		PricingScaleLine:             data.PricingScaleLine,
		ConditionReleaseStatus:       data.ConditionReleaseStatus,
	}, nil
}
