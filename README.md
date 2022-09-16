# sap-api-integrations-purchasing-info-record-reads
sap-api-integrations-purchasing-info-record-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で購買情報レコードを取得するマイクロサービスです。    
sap-api-integrations-purchasing-info-record-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchasing-info-record-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_INFORECORD_PROCESS_SRV_0001/overview   

## 動作環境  
sap-api-integrations-purchasing-info-record-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-purchasing-info-record-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchasing-info-record-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_INFORECORD_PROCESS_SRV_0001/overview    
* APIサービス名(=baseURL): API_INFORECORD_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchasing-info-record-reads には、次の API をコールするためのリソースが含まれています。  

* A_PurchasingInfoRecord（購買情報 - 一般）※価格条件関連データを取得するために、ToPurgInfoRecdOrgPlantData、ToPurInfoRecdPrcgCndnValidity、ToPurInfoRecdPrcgCndn、と合わせて利用されます。
* A_PurgInfoRecdOrgPlantData（購買情報 - 購買組織プラント）※価格条件関連データを取得するために、ToPurInfoRecdPrcgCndnValidity、ToPurInfoRecdPrcgCndn、と合わせて利用されます。
* ToPurInfoRecdPrcgCndnValidity（購買情報 - 価格条件存在性）
* ToPurInfoRecdPrcgCndn（購買情報 - 価格条件）
* ToPurgInfoRecdOrgPlantData（購買情報 - 購買組織プラント）

## API への 値入力条件 の 初期値
sap-api-integrations-purchasing-info-record-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.PurchasingInfoRecord.PurchasingInfoRecord（購買情報）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingInfoRecordCategory（購買情報カテゴリ）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Supplier（仕入先）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Material（品目）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingOrganization（購買組織）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Plant（プラント）
* inputSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.MaterialGroup（品目グループ）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "SAPPurchasingInforecordReads",
	"accepter": ["General"],
	"purchasing_info_record": "5300000000",
	"deleted": null
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPPurchasingInforecordReads",
	"accepter": ["All"],
	"purchasing_info_record": "5300000000",
	"deleted": null
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 購買情報 の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"PurchasingInfoRecord" ～ "ToPurgInfoRecdOrgPlantData" は、/SAP_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller/caller.go#L65",
	"function": "sap-api-integrations-purchasing-info-record-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"PurchasingInfoRecord": "5300000000",
			"Supplier": "100000",
			"Material": "21",
			"MaterialGroup": "",
			"PurgDocOrderQuantityUnit": "PC",
			"SupplierMaterialNumber": "TEST SUPPLIER MATERIAL NO",
			"SupplierRespSalesPersonName": "",
			"SupplierPhoneNumber": "",
			"SupplierMaterialGroup": "",
			"IsRegularSupplier": false,
			"AvailabilityStartDate": "",
			"AvailabilityEndDate": "",
			"Manufacturer": "",
			"CreationDate": "2022-09-16",
			"PurchasingInfoRecordDesc": "",
			"LastChangeDateTime": "2022-09-16T17:28:36+09:00",
			"IsDeleted": false,
			"to_PurgInfoRecdOrgPlantData": "http://100.21.57.120:8080/sap/opu/odata/sap/API_INFORECORD_PROCESS_SRV/A_PurchasingInfoRecord('5300000000')/to_PurgInfoRecdOrgPlantData"
		}
	],
	"time": "2022-09-16T17:32:41+09:00"
}
```

