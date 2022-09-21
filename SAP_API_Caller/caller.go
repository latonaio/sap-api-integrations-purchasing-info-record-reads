package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchasing-info-record-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchasingInfoRecord(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, materialGroup, conditionType string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(purchasingInfoRecord)
				wg.Done()
			}()
		case "Material":
			func() {
				c.Material(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)
				wg.Done()
			}()
		case "MaterialGroup":
			func() {
				c.MaterialGroup(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(purchasingInfoRecord string) {
	generalData, err := c.callPurchasingInfoRecordSrvAPIRequirementGeneral("A_PurchasingInfoRecord", purchasingInfoRecord)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(generalData)
	}

	orgPlantData, err := c.callToPurgInfoRecdOrgPlantData(generalData[0].ToPurgInfoRecdOrgPlantData)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(orgPlantData)
	}

	validityData, err := c.callToPurInfoRecdPrcgCndnValidity(orgPlantData[0].ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(validityData)
	}

	cndnData, err := c.callToPurInfoRecdPrcgCndn(validityData[0].ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(cndnData)
	}
	return
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementGeneral(api, purchasingInfoRecord string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	param := c.getQueryWithGeneral(map[string]string{}, purchasingInfoRecord)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurgInfoRecdOrgPlantData(url string) ([]sap_api_output_formatter.ToPurgInfoRecdOrgPlant, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurgInfoRecdOrgPlant(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndnValidity(url string) ([]sap_api_output_formatter.ToPurInfoRecdPrcgCndnValidity, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndnValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurInfoRecdPrcgCndn(url string) (*sap_api_output_formatter.ToPurInfoRecdPrcgCndn, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurInfoRecdPrcgCndn(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Material(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) {
	materialData, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterial("A_PurgInfoRecdOrgPlantData", purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(materialData)
	}

	validityData, err := c.callToPurInfoRecdPrcgCndnValidity(materialData[0].ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(validityData)
	}

	cndnData, err := c.callToPurInfoRecdPrcgCndn(validityData[0].ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(cndnData)
	}
	return
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterial(api, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) ([]sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")

	param := c.getQueryWithMaterial(map[string]string{}, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialGroup(purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) {
	materialGroupData, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterialGroup("A_PurgInfoRecdOrgPlantData", purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(materialGroupData)
	}

	validityData, err := c.callToPurInfoRecdPrcgCndnValidity(materialGroupData[0].ToPurInfoRecdPrcgCndnValidity)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(validityData)
	}

	cndnData, err := c.callToPurInfoRecdPrcgCndn(validityData[0].ToPurInfoRecdPrcgCndn)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(cndnData)
	}
	return
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialGroup(api, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) ([]sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")

	param := c.getQueryWithMaterialGroup(map[string]string{}, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithGeneral(params map[string]string, purchasingInfoRecord string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchasingInfoRecord eq '%s'", purchasingInfoRecord)
	return params
}

func (c *SAPAPICaller) getQueryWithMaterial(params map[string]string, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and Material eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, material, purchasingOrganization, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithMaterialGroup(params map[string]string, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, materialGroup, purchasingOrganization, plant)
	return params
}
