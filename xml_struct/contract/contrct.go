package contract

import "gitlab.com/garanteka/parsexsd/xsd"

// Contract Информация (проект информации) о заключенном контракте
type Contract struct {
}

// ZfcsContract2015Type
type ZfcsContract2015Type struct {
	SchemeVersion                       string                                   `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                                  int64                                    `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                          string                                   `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DirectDate                          xsd.Date                                 `xml:"directDate,omitempty" bson:"directDate,omitempty"`
	PublishDate                         xsd.Date                                 `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	VersionNumber                       int64                                    `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Foundation                          ZfcsContract2015TypeFoundation           `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Customer                            ZfcsContract2015TypeCustomer             `xml:"customer,omitempty" bson:"customer,omitempty"`
	Placer                              ZfcsContract2015TypePlacer               `xml:"placer,omitempty" bson:"placer,omitempty"`
	Finances                            ZfcsContract2015TypeFinances             `xml:"finances,omitempty" bson:"finances,omitempty"`
	ProtocolDate                        xsd.Date                                 `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	DocumentBase                        string                                   `xml:"documentBase,omitempty" bson:"documentBase,omitempty"`
	DocumentCode                        string                                   `xml:"documentCode,omitempty" bson:"documentCode,omitempty"`
	SignDate                            xsd.Date                                 `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	RegNum                              string                                   `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	Number                              string                                   `xml:"number,omitempty" bson:"number,omitempty"`
	DefenseContractNumber               string                                   `xml:"defenseContractNumber,omitempty" bson:"defenseContractNumber,omitempty"`
	PriceInfo                           ZfcsContract2015TypePriceInfo            `xml:"priceInfo,omitempty" bson:"priceInfo,omitempty"`
	AdvancePaymentSum                   ZfcsContract2015TypeAdvancePaymentSum    `xml:"advancePaymentSum,omitempty" bson:"advancePaymentSum,omitempty"`
	SubContractorsSum                   ZfcsContract2015TypeSubContractorsSum    `xml:"subContractorsSum,omitempty" bson:"subContractorsSum,omitempty"`
	ExecutionPeriod                     ZfcsContract2015TypeExecutionPeriod      `xml:"executionPeriod,omitempty" bson:"executionPeriod,omitempty"`
	Enforcement                         ZfcsContract2015TypeEnforcement          `xml:"enforcement,omitempty" bson:"enforcement,omitempty"`
	QualityGuaranteeInfo                ZfcsContract2015TypeQualityGuaranteeInfo `xml:"qualityGuaranteeInfo,omitempty" bson:"qualityGuaranteeInfo,omitempty"`
	GuaranteeReturns                    ZfcsContract2015BankGuaranteeReturnType  `xml:"guaranteeReturns,omitempty" bson:"guaranteeReturns,omitempty"`
	EnergyServiceContractInfo           string                                   `xml:"energyServiceContractInfo,omitempty" bson:"energyServiceContractInfo,omitempty"`
	Products                            ZfcsContract2015TypeProducts             `xml:"products,omitempty" bson:"products,omitempty"`
	Suppliers                           ZfcsContract2015TypeSuppliers            `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	Href                                string                                   `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                           ZfcsContractPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm                        ZfcsExtPrintFormType                     `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ScanDocuments                       ZfcsContractAttachmentListType           `xml:"scanDocuments,omitempty" bson:"scanDocuments,omitempty"`
	MedicalDocuments                    ZfcsContractAttachmentListType           `xml:"medicalDocuments,omitempty" bson:"medicalDocuments,omitempty"`
	SingleSupplierP25Part1St93Documents ZfcsContractAttachmentListType           `xml:"singleSupplierP25Part1St93Documents,omitempty" bson:"singleSupplierP25Part1St93Documents,omitempty"`
	BudgetObligations                   ZfcsContractAttachmentType               `xml:"budgetObligations,omitempty" bson:"budgetObligations,omitempty"`
	Attachments                         ZfcsContractAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification                        ZfcsContract2015TypeModification         `xml:"modification,omitempty" bson:"modification,omitempty"`
	CurrentContractStage                string                                   `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
	Okpd2okved2                         bool                                     `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	IsProcessed                         bool                                     `xml:"is_processed,omitempty" bson:"is_processed,omitempty"`
}

// ZfcsContract2015TypeFoundation Основание заключения контракта
type ZfcsContract2015TypeFoundation struct {
	OosOrder FoundationOosOrder `xml:"oosOrder,omitempty" bson:"oosOrder,omitempty"`
	FcsOrder FoundationFcsOrder `xml:"fcsOrder,omitempty" bson:"fcsOrder,omitempty"`
}

// FoundationOosOrder Размещение заказа по закону №94-ФЗ
type FoundationOosOrder struct {
	Order          OosOrderOrder                      `xml:"order,omitempty" bson:"order,omitempty"`
	SingleCustomer ZfcsContract2015SingleCustomerType `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
	NotOosOrder    OosOrderNotOosOrder                `xml:"notOosOrder,omitempty" bson:"notOosOrder,omitempty"`
	Other          OosOrderOther                      `xml:"other,omitempty" bson:"other,omitempty"`
}

// OosOrderOrder Заказ
type OosOrderOrder struct {
	NotificationNumber string                             `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	LotNumber          int64                              `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	Placing            string                             `xml:"placing,omitempty" bson:"placing,omitempty"`
	SingleCustomer     ZfcsContract2015SingleCustomerType `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
}

// OosOrderNotOosOrder Заказ не размещался на ЕИС
type OosOrderNotOosOrder struct {
	Placing string `xml:"placing,omitempty" bson:"placing,omitempty"`
}

// OosOrderOther Контракт заключен по результатам процедур размещения заказов, начатых до 1 января 2011 года
type OosOrderOther struct {
	NotificationNumber string `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	Placing            string `xml:"placing,omitempty" bson:"placing,omitempty"`
}

// FoundationFcsOrder Размещение закупки по закону №44-ФЗ
type FoundationFcsOrder struct {
	Order          FcsOrderOrder          `xml:"order,omitempty" bson:"order,omitempty"`
	SingleCustomer FcsOrderSingleCustomer `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
	NotOosOrder    FcsOrderNotOosOrder    `xml:"notOosOrder,omitempty" bson:"notOosOrder,omitempty"`
}

// FcsOrderOrder Закупка
type FcsOrderOrder struct {
	NotificationNumber string                             `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	LotNumber          int64                              `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	Placing            string                             `xml:"placing,omitempty" bson:"placing,omitempty"`
	SingleCustomer     ZfcsContract2015SingleCustomerType `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
	PurchaseCode       string                             `xml:"purchaseCode,omitempty" bson:"purchaseCode,omitempty"`
	TenderPlanInfo     ZfcsContract2015TenderPlanInfoType `xml:"tenderPlanInfo,omitempty" bson:"tenderPlanInfo,omitempty"`
}

// FcsOrderSingleCustomer Закупка у единственного поставщика (извещение не размещалось на ЕИС)
type FcsOrderSingleCustomer struct {
	Reason         SingleCustomerReason               `xml:"reason,omitempty" bson:"reason,omitempty"`
	Document       SingleCustomerDocument             `xml:"document,omitempty" bson:"document,omitempty"`
	ReportBase     string                             `xml:"reportBase,omitempty" bson:"reportBase,omitempty"`
	ReportCode     string                             `xml:"reportCode,omitempty" bson:"reportCode,omitempty"`
	Attachments    ZfcsContractAttachmentListType     `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PurchaseCode   string                             `xml:"purchaseCode,omitempty" bson:"purchaseCode,omitempty"`
	TenderPlanInfo ZfcsContract2015TenderPlanInfoType `xml:"tenderPlanInfo,omitempty" bson:"tenderPlanInfo,omitempty"`
}

// SingleCustomerReason Основание заключения контракта с единственным поставщиком
type SingleCustomerReason struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// SingleCustomerDocument Наименование документа, подтверждающего заключение контракта
type SingleCustomerDocument struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// FcsOrderNotOosOrder Закупка не размещалась на ЕИС
type FcsOrderNotOosOrder struct {
	Placing        string                             `xml:"placing,omitempty" bson:"placing,omitempty"`
	Document       NotOosOrderDocument                `xml:"document,omitempty" bson:"document,omitempty"`
	PurchaseCode   string                             `xml:"purchaseCode,omitempty" bson:"purchaseCode,omitempty"`
	TenderPlanInfo ZfcsContract2015TenderPlanInfoType `xml:"tenderPlanInfo,omitempty" bson:"tenderPlanInfo,omitempty"`
}

// NotOosOrderDocument Наименование документа, подтверждающего заключение контракта с единственным поставщиком
type NotOosOrderDocument struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsContract2015TypeCustomer Заказчик
type ZfcsContract2015TypeCustomer struct {
	RegNum           string    `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum  string    `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName         string    `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName        string    `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	RegistrationDate xsd.Date  `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	Inn              string    `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp              string    `xml:"kpp,omitempty" bson:"kpp,omitempty"`
	LegalForm        OkopfType `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	OKPO             string    `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	CustomerCode     string    `xml:"customerCode,omitempty" bson:"customerCode,omitempty"`
}

// ZfcsContract2015TypePlacer Информация об организации, разместившей контракт. Обязательно указание блока в случае приема в ЕИС по альтернативной интеграции под логином пользователя организации, имеющей несколько ролей
type ZfcsContract2015TypePlacer struct {
	ResponsibleOrg  ZfcsOrganizationRef `xml:"responsibleOrg,omitempty" bson:"responsibleOrg,omitempty"`
	ResponsibleRole string              `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// ZfcsContract2015TypeFinances Финансирование
type ZfcsContract2015TypeFinances struct {
	BudgetFunds      FinancesBudgetFunds      `xml:"budgetFunds,omitempty" bson:"budgetFunds,omitempty"`
	ExtrabudgetFunds FinancesExtrabudgetFunds `xml:"extrabudgetFunds,omitempty" bson:"extrabudgetFunds,omitempty"`
}

// FinancesBudgetFunds Бюджетные средства
type FinancesBudgetFunds struct {
	Budget           ZfcsBudgetFundsContract2015 `xml:"budget,omitempty" bson:"budget,omitempty"`
	OKTMO            ZfcsOKTMORef                `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	BudgetLevel      string                      `xml:"budgetLevel,omitempty" bson:"budgetLevel,omitempty"`
	FundsBudgetLevel string                      `xml:"fundsBudgetLevel,omitempty" bson:"fundsBudgetLevel,omitempty"`
	Stages           []BudgetFundsStages         `xml:"stages,omitempty" bson:"stages,omitempty"`
}

// BudgetFundsStages Этапы исполнения контракта
type BudgetFundsStages struct {
	EndDate  xsd.Date         `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	Payments []StagesPayments `xml:"payments,omitempty" bson:"payments,omitempty"`
}

// StagesPayments Платежи по  этапу
type StagesPayments struct {
	Comment       string `xml:"comment,omitempty" bson:"comment,omitempty"`
	PaymentMonth  byte   `xml:"paymentMonth,omitempty" bson:"paymentMonth,omitempty"`
	PaymentYear   int64  `xml:"paymentYear,omitempty" bson:"paymentYear,omitempty"`
	PaymentSum    string `xml:"paymentSum,omitempty" bson:"paymentSum,omitempty"`
	PaymentSumRUR string `xml:"paymentSumRUR,omitempty" bson:"paymentSumRUR,omitempty"`
	KBK           string `xml:"KBK,omitempty" bson:"KBK,omitempty"`
	KBK2016       string `xml:"KBK2016,omitempty" bson:"KBK2016,omitempty"`
	KVR           string `xml:"KVR,omitempty" bson:"KVR,omitempty"`
}

// FinancesExtrabudgetFunds Внебюджетные средства
type FinancesExtrabudgetFunds struct {
	Extrabudget ZfcsExtraBudgetFundsContract2015 `xml:"extrabudget,omitempty" bson:"extrabudget,omitempty"`
	Stages      []ExtrabudgetFundsStages         `xml:"stages,omitempty" bson:"stages,omitempty"`
}

// ExtrabudgetFundsStages Этапы исполнения контракта
type ExtrabudgetFundsStages struct {
	EndDate  xsd.Date         `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	Payments []StagesPayments `xml:"payments,omitempty" bson:"payments,omitempty"`
}

// ZfcsContract2015TypePriceInfo Информация о цене контракта
type ZfcsContract2015TypePriceInfo struct {
	Price        string                       `xml:"price,omitempty" bson:"price,omitempty"`
	PriceType    string                       `xml:"priceType,omitempty" bson:"priceType,omitempty"`
	PriceFormula string                       `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	PriceRUR     string                       `xml:"priceRUR,omitempty" bson:"priceRUR,omitempty"`
	PriceVAT     string                       `xml:"priceVAT,omitempty" bson:"priceVAT,omitempty"`
	PriceVATRUR  string                       `xml:"priceVATRUR,omitempty" bson:"priceVATRUR,omitempty"`
}

// ZfcsContract2015TypeAdvancePaymentSum Предусмотрена выплата аванса. В блоке должно быть заполнено либо поле sumInPercents либо поле priceValue. Оба поля одновременно заполнены быть не могут
type ZfcsContract2015TypeAdvancePaymentSum struct {
	SumInPercents float64 `xml:"sumInPercents,omitempty" bson:"sumInPercents,omitempty"`
	PriceValue    string  `xml:"priceValue,omitempty" bson:"priceValue,omitempty"`
	PriceValueRUR string  `xml:"priceValueRUR,omitempty" bson:"priceValueRUR,omitempty"`
}

// ZfcsContract2015TypeSubContractorsSum Объем  привлечения к исполнению контракта субподрядчиков
type ZfcsContract2015TypeSubContractorsSum struct {
	SumInPercents  float64                         `xml:"sumInPercents,omitempty" bson:"sumInPercents,omitempty"`
	PriceValueRUR  string                          `xml:"priceValueRUR,omitempty" bson:"priceValueRUR,omitempty"`
	SubContractors SubContractorsSumSubContractors `xml:"subContractors,omitempty" bson:"subContractors,omitempty"`
}

// SubContractorsSumSubContractors Информация о субподрядчиках, соисполнителях (СМП, СОНО). Контролируется обязательность заполнения в случае если sumInPercents>0
type SubContractorsSumSubContractors struct {
	SubContractor []SubContractorsSubContractor `xml:"subContractor,omitempty" bson:"subContractor,omitempty"`
}

// SubContractorsSubContractor Субподрядчик, соисполнитель
type SubContractorsSubContractor struct {
	LegalEntityRF           SubContractorLegalEntityRF           `xml:"legalEntityRF,omitempty" bson:"legalEntityRF,omitempty"`
	IndividualBusinessmanRF SubContractorIndividualBusinessmanRF `xml:"individualBusinessmanRF,omitempty" bson:"individualBusinessmanRF,omitempty"`
}

// SubContractorLegalEntityRF Юридическое лицо РФ
type SubContractorLegalEntityRF struct {
	FullName        string                              `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	FirmName        string                              `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	INN             string                              `xml:"INN,omitempty" bson:"INN,omitempty"`
	Address         string                              `xml:"address,omitempty" bson:"address,omitempty"`
	SubContractInfo ZfcsContract2015SubContractInfoType `xml:"subContractInfo,omitempty" bson:"subContractInfo,omitempty"`
}

// SubContractorIndividualBusinessmanRF Индивидуальный предприниматель РФ
type SubContractorIndividualBusinessmanRF struct {
	LastName        string                              `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName       string                              `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName      string                              `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	INN             string                              `xml:"INN,omitempty" bson:"INN,omitempty"`
	Address         string                              `xml:"address,omitempty" bson:"address,omitempty"`
	SubContractInfo ZfcsContract2015SubContractInfoType `xml:"subContractInfo,omitempty" bson:"subContractInfo,omitempty"`
}

// ZfcsContract2015TypeExecutionPeriod Срок исполнения контракта
type ZfcsContract2015TypeExecutionPeriod struct {
	StartDate xsd.Date                `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	Stages    []ExecutionPeriodStages `xml:"stages,omitempty" bson:"stages,omitempty"`
	EndDate   xsd.Date                `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ExecutionPeriodStages Этапы исполнения контракта
type ExecutionPeriodStages struct {
	EndDate xsd.Date `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ZfcsContract2015TypeEnforcement Обеспечение исполнения контракта
type ZfcsContract2015TypeEnforcement struct {
	BankGuarantee EnforcementBankGuarantee `xml:"bankGuarantee,omitempty" bson:"bankGuarantee,omitempty"`
	CashAccount   EnforcementCashAccount   `xml:"cashAccount,omitempty" bson:"cashAccount,omitempty"`
}

// EnforcementBankGuarantee Банковская гарантия, выданная банком в соответствии со статьей 45
type EnforcementBankGuarantee struct {
	RegNumber          string                       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber          string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Currency           ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	GuaranteeAmount    string                       `xml:"guaranteeAmount,omitempty" bson:"guaranteeAmount,omitempty"`
	CurrencyRate       ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	GuaranteeAmountRUR string                       `xml:"guaranteeAmountRUR,omitempty" bson:"guaranteeAmountRUR,omitempty"`
}

// EnforcementCashAccount Внесение денежных средств на указанный заказчиком счет
type EnforcementCashAccount struct {
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// ZfcsContract2015TypeQualityGuaranteeInfo Информация о гарантии качества товара, работы услуги
type ZfcsContract2015TypeQualityGuaranteeInfo struct {
	ProvidedPeriod               QualityGuaranteeInfoProvidedPeriod               `xml:"providedPeriod,omitempty" bson:"providedPeriod,omitempty"`
	WarrantyReqsText             string                                           `xml:"warrantyReqsText,omitempty" bson:"warrantyReqsText,omitempty"`
	ManufacturerWarrantyReqsText string                                           `xml:"manufacturerWarrantyReqsText,omitempty" bson:"manufacturerWarrantyReqsText,omitempty"`
	ExecObligationsGuaranteeInfo QualityGuaranteeInfoExecObligationsGuaranteeInfo `xml:"execObligationsGuaranteeInfo,omitempty" bson:"execObligationsGuaranteeInfo,omitempty"`
}

// QualityGuaranteeInfoProvidedPeriod Срок, на который предоставляется гарантия
type QualityGuaranteeInfoProvidedPeriod struct {
	FromDate xsd.Date `xml:"fromDate,omitempty" bson:"fromDate,omitempty"`
	ToDate   xsd.Date `xml:"toDate,omitempty" bson:"toDate,omitempty"`
}

// QualityGuaranteeInfoExecObligationsGuaranteeInfo Обеспечение исполнения обязательств по предоставленной гарантии качества товаров, работ, услуг
type QualityGuaranteeInfoExecObligationsGuaranteeInfo struct {
	EnsuringWay      ExecObligationsGuaranteeInfoEnsuringWay `xml:"ensuringWay,omitempty" bson:"ensuringWay,omitempty"`
	GuaranteeReturns ZfcsContract2015BankGuaranteeReturnType `xml:"guaranteeReturns,omitempty" bson:"guaranteeReturns,omitempty"`
}

// ExecObligationsGuaranteeInfoEnsuringWay Способ обеспечения исполнения обязательств по предоставленной гарантии
type ExecObligationsGuaranteeInfoEnsuringWay struct {
	BankGuarantee EnsuringWayBankGuarantee `xml:"bankGuarantee,omitempty" bson:"bankGuarantee,omitempty"`
	CashAccount   EnsuringWayCashAccount   `xml:"cashAccount,omitempty" bson:"cashAccount,omitempty"`
}

// EnsuringWayBankGuarantee Банковская гарантия, выданная банком в соответствии со статьей 45
type EnsuringWayBankGuarantee struct {
	RegNumber    string                       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber    string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// EnsuringWayCashAccount Внесение денежных средств на указанный заказчиком счет
type EnsuringWayCashAccount struct {
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// ZfcsContract2015TypeProducts Объекты закупки
type ZfcsContract2015TypeProducts struct {
	Product        []ProductsProduct      `xml:"product,omitempty" bson:"product,omitempty"`
	ProductsChange ProductsProductsChange `xml:"productsChange,omitempty" bson:"productsChange,omitempty"`
}

// ProductsProduct Объект закупки
type ProductsProduct struct {
	Sid                    int64                                `xml:"sid,omitempty" bson:"sid,omitempty"`
	ExternalSid            string                               `xml:"externalSid,omitempty" bson:"externalSid,omitempty"`
	Name                   string                               `xml:"name,omitempty" bson:"name,omitempty"`
	OKEI                   ZfcsContractOKEIType                 `xml:"OKEI,omitempty" bson:"OKEI,omitempty"`
	Price                  string                               `xml:"price,omitempty" bson:"price,omitempty"`
	PriceRUR               string                               `xml:"priceRUR,omitempty" bson:"priceRUR,omitempty"`
	Quantity               float64                              `xml:"quantity,omitempty" bson:"quantity,omitempty"`
	Sum                    string                               `xml:"sum,omitempty" bson:"sum,omitempty"`
	SumRUR                 string                               `xml:"sumRUR,omitempty" bson:"sumRUR,omitempty"`
	DrugPurchaseObjectInfo ZfcsContract2015DrugPurchaseInfoType `xml:"drugPurchaseObjectInfo,omitempty" bson:"drugPurchaseObjectInfo,omitempty"`
	OKPD                   ZfcsOKPDRef                          `xml:"OKPD,omitempty" bson:"OKPD,omitempty"`
	OKPD2                  ZfcsOKPDRef                          `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
	KTRU                   ZfcsKTRURef                          `xml:"KTRU,omitempty" bson:"KTRU,omitempty"`
}

// ProductsProductsChange Сведения об изменении объектов закупки
type ProductsProductsChange struct {
	Documents  []ProductsChangeDocuments `xml:"documents,omitempty" bson:"documents,omitempty"`
	ChangeInfo string                    `xml:"changeInfo,omitempty" bson:"changeInfo,omitempty"`
}

// ProductsChangeDocuments Документы
type ProductsChangeDocuments struct {
	DocumentName string   `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
}

// ZfcsContract2015TypeSuppliers Поставщики
type ZfcsContract2015TypeSuppliers struct {
	Supplier []ZfcsContract2015SupplierType `xml:"supplier,omitempty" bson:"supplier,omitempty"`
}

// ZfcsContract2015TypeModification Описание внесения изменений
type ZfcsContract2015TypeModification struct {
	Attachments     ZfcsContractAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ContractChange  ModificationContractChange     `xml:"contractChange,omitempty" bson:"contractChange,omitempty"`
	ErrorCorrection ModificationErrorCorrection    `xml:"errorCorrection,omitempty" bson:"errorCorrection,omitempty"`
}

// ModificationContractChange Изменение контракта
type ModificationContractChange struct {
	Reason         ContractChangeReason           `xml:"reason,omitempty" bson:"reason,omitempty"`
	Document       ContractChangeDocument         `xml:"document,omitempty" bson:"document,omitempty"`
	DamagePayments []ContractChangeDamagePayments `xml:"damagePayments,omitempty" bson:"damagePayments,omitempty"`
}

// ContractChangeReason Причина изменений условий контракта
type ContractChangeReason struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ContractChangeDocument Документ, являющийся основанием изменений контракта
type ContractChangeDocument struct {
	Code         string   `xml:"code,omitempty" bson:"code,omitempty"`
	Name         string   `xml:"name,omitempty" bson:"name,omitempty"`
	Base         string   `xml:"base,omitempty" bson:"base,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
}

// ContractChangeDamagePayments Информация об оплате суммы возмещения фактически понесенного ущерба
type ContractChangeDamagePayments struct {
	Document     ZfcsContract2015DocumentInfo `xml:"document,omitempty" bson:"document,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// ModificationErrorCorrection Корректировка ошибок
type ModificationErrorCorrection struct {
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsContract2015SingleCustomerType
type ZfcsContract2015SingleCustomerType struct {
	Reason      ZfcsContract2015SingleCustomerTypeReason   `xml:"reason,omitempty" bson:"reason,omitempty"`
	Document    ZfcsContract2015SingleCustomerTypeDocument `xml:"document,omitempty" bson:"document,omitempty"`
	ReportBase  string                                     `xml:"reportBase,omitempty" bson:"reportBase,omitempty"`
	ReportCode  string                                     `xml:"reportCode,omitempty" bson:"reportCode,omitempty"`
	Attachments ZfcsContractAttachmentListType             `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsContract2015SingleCustomerTypeReason Основание заключения контракта с единственным поставщиком
type ZfcsContract2015SingleCustomerTypeReason struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsContract2015SingleCustomerTypeDocument Наименование документа, подтверждающего заключение контракта
type ZfcsContract2015SingleCustomerTypeDocument struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsContractAttachmentListType
type ZfcsContractAttachmentListType struct {
	Attachment []ZfcsContractAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsContractAttachmentType
type ZfcsContractAttachmentType struct {
	PublishedContentID string                                `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string                                `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	DocDescription     string                                `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	DocRegNumber       string                                `xml:"docRegNumber,omitempty" bson:"docRegNumber,omitempty"`
	CryptoSigns        ZfcsContractAttachmentTypeCryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string                                `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string                                `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte                                `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsContractAttachmentTypeCryptoSigns Электронная подпись документа
type ZfcsContractAttachmentTypeCryptoSigns struct {
	Signature []CryptoSignsSignature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// CryptoSignsSignature Электронная подпись
type CryptoSignsSignature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsContract2015TenderPlanInfoType
type ZfcsContract2015TenderPlanInfoType struct {
	Plan2017Number        string `xml:"plan2017Number,omitempty" bson:"plan2017Number,omitempty"`
	Position2017Number    string `xml:"position2017Number,omitempty" bson:"position2017Number,omitempty"`
	Position2017ExtNumber string `xml:"position2017ExtNumber,omitempty" bson:"position2017ExtNumber,omitempty"`
}

// OkopfType
type OkopfType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ZfcsOrganizationRef
type ZfcsOrganizationRef struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsBudgetFundsContract2015
type ZfcsBudgetFundsContract2015 struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsOKTMORef
type ZfcsOKTMORef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsExtraBudgetFundsContract2015
type ZfcsExtraBudgetFundsContract2015 struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsCurrencyRef
type ZfcsCurrencyRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsCurrencyRateContract2015
type ZfcsCurrencyRateContract2015 struct {
	Rate    float64 `xml:"rate,omitempty" bson:"rate,omitempty"`
	Raiting int64   `xml:"raiting,omitempty" bson:"raiting,omitempty"`
}

// ZfcsContract2015SubContractInfoType
type ZfcsContract2015SubContractInfoType struct {
	SubContractDate      xsd.Date                                                `xml:"subContractDate,omitempty" bson:"subContractDate,omitempty"`
	SubContractNumber    string                                                  `xml:"subContractNumber,omitempty" bson:"subContractNumber,omitempty"`
	SubContractSubject   string                                                  `xml:"subContractSubject,omitempty" bson:"subContractSubject,omitempty"`
	SubContractPriceInfo ZfcsContract2015SubContractInfoTypeSubContractPriceInfo `xml:"subContractPriceInfo,omitempty" bson:"subContractPriceInfo,omitempty"`
}

// ZfcsContract2015SubContractInfoTypeSubContractPriceInfo Информация о цене договора
type ZfcsContract2015SubContractInfoTypeSubContractPriceInfo struct {
	Price        string                       `xml:"price,omitempty" bson:"price,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	PriceRUR     string                       `xml:"priceRUR,omitempty" bson:"priceRUR,omitempty"`
}

// ZfcsContract2015BankGuaranteeReturnType
type ZfcsContract2015BankGuaranteeReturnType struct {
	GuaranteeReturn []ZfcsContract2015BankGuaranteeReturnTypeGuaranteeReturn `xml:"guaranteeReturn,omitempty" bson:"guaranteeReturn,omitempty"`
}

// ZfcsContract2015BankGuaranteeReturnTypeGuaranteeReturn Информация
type ZfcsContract2015BankGuaranteeReturnTypeGuaranteeReturn struct {
	BankGuaranteeReturn GuaranteeReturnBankGuaranteeReturn `xml:"bankGuaranteeReturn,omitempty" bson:"bankGuaranteeReturn,omitempty"`
	WaiverNotice        GuaranteeReturnWaiverNotice        `xml:"waiverNotice,omitempty" bson:"waiverNotice,omitempty"`
}

// GuaranteeReturnBankGuaranteeReturn Информация о возвращении заказчиком банковской гарантии гаранту
type GuaranteeReturnBankGuaranteeReturn struct {
	RegNumber         string   `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber         string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	ReturnDate        xsd.Date `xml:"returnDate,omitempty" bson:"returnDate,omitempty"`
	ReturnReason      string   `xml:"returnReason,omitempty" bson:"returnReason,omitempty"`
	ReturnPublishDate xsd.Date `xml:"returnPublishDate,omitempty" bson:"returnPublishDate,omitempty"`
}

// GuaranteeReturnWaiverNotice Информация об уведомлении, направленном заказчиком гаранту, об освобождении от обязательств по банковской гарантии
type GuaranteeReturnWaiverNotice struct {
	RegNumber         string   `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber         string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	NoticeDate        xsd.Date `xml:"noticeDate,omitempty" bson:"noticeDate,omitempty"`
	NoticeNumber      string   `xml:"noticeNumber,omitempty" bson:"noticeNumber,omitempty"`
	NoticeReason      string   `xml:"noticeReason,omitempty" bson:"noticeReason,omitempty"`
	NoticePublishDate xsd.Date `xml:"noticePublishDate,omitempty" bson:"noticePublishDate,omitempty"`
}

// ZfcsContractOKEIType
type ZfcsContractOKEIType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	NationalCode string `xml:"nationalCode,omitempty" bson:"nationalCode,omitempty"`
	FullName     string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsContract2015DrugPurchaseInfoType
type ZfcsContract2015DrugPurchaseInfoType struct {
	DrugInfoUsingReferenceInfo ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingReferenceInfo `xml:"drugInfoUsingReferenceInfo,omitempty" bson:"drugInfoUsingReferenceInfo,omitempty"`
	DrugInfoUsingTextForm      ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingTextForm      `xml:"drugInfoUsingTextForm,omitempty" bson:"drugInfoUsingTextForm,omitempty"`
}

// ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingReferenceInfo Информация о лекарственных препаратах формируется с использованием справочной информации (не в текстовой форме)
type ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingReferenceInfo struct {
	MNNsInfo                       DrugInfoUsingReferenceInfoMNNsInfo                       `xml:"MNNsInfo,omitempty" bson:"MNNsInfo,omitempty"`
	MNNInfo                        DrugInfoUsingReferenceInfoMNNInfo                        `xml:"MNNInfo,omitempty" bson:"MNNInfo,omitempty"`
	ExpirationDateCustomFormatInfo DrugInfoUsingReferenceInfoExpirationDateCustomFormatInfo `xml:"expirationDateCustomFormatInfo,omitempty" bson:"expirationDateCustomFormatInfo,omitempty"`
	IsZNVLP                        bool                                                     `xml:"isZNVLP,omitempty" bson:"isZNVLP,omitempty"`
	PositionsTradeName             DrugInfoUsingReferenceInfoPositionsTradeName             `xml:"positionsTradeName,omitempty" bson:"positionsTradeName,omitempty"`
}

// DrugInfoUsingReferenceInfoMNNsInfo Международные, группировочные или химические наименования лекарственных препаратов (МНН). При приеме контролируется обязательность заполнения блока и принадлежность всех МНН в блоке одному типу (т.е. все МНН в блоке должны иметь одинаковое наименование)
type DrugInfoUsingReferenceInfoMNNsInfo struct {
	MNNInfo []MNNsInfoMNNInfo `xml:"MNNInfo,omitempty" bson:"MNNInfo,omitempty"`
}

// MNNsInfoMNNInfo Международное, группировочное или химическое наименование лекарственного препарата (МНН)
type MNNsInfoMNNInfo struct {
	MNNExternalCode    string                    `xml:"MNNExternalCode,omitempty" bson:"MNNExternalCode,omitempty"`
	MNNDrugCode        string                    `xml:"MNNDrugCode,omitempty" bson:"MNNDrugCode,omitempty"`
	MNNName            string                    `xml:"MNNName,omitempty" bson:"MNNName,omitempty"`
	PositionsTradeName MNNInfoPositionsTradeName `xml:"positionsTradeName,omitempty" bson:"positionsTradeName,omitempty"`
}

// MNNInfoPositionsTradeName Позиции по торговому наименованию лекарственного средства
type MNNInfoPositionsTradeName struct {
	PositionTradeName []PositionsTradeNamePositionTradeName `xml:"positionTradeName,omitempty" bson:"positionTradeName,omitempty"`
}

// PositionsTradeNamePositionTradeName Позиция
type PositionsTradeNamePositionTradeName struct {
	PositionTradeNameExternalCode string                                `xml:"positionTradeNameExternalCode,omitempty" bson:"positionTradeNameExternalCode,omitempty"`
	DrugCode                      string                                `xml:"drugCode,omitempty" bson:"drugCode,omitempty"`
	TradeInfo                     PositionTradeNameTradeInfo            `xml:"tradeInfo,omitempty" bson:"tradeInfo,omitempty"`
	CertificateNumber             string                                `xml:"certificateNumber,omitempty" bson:"certificateNumber,omitempty"`
	MedicamentalFormInfo          PositionTradeNameMedicamentalFormInfo `xml:"medicamentalFormInfo,omitempty" bson:"medicamentalFormInfo,omitempty"`
	DosageInfo                    PositionTradeNameDosageInfo           `xml:"dosageInfo,omitempty" bson:"dosageInfo,omitempty"`
	CertificateKeeperName         string                                `xml:"certificateKeeperName,omitempty" bson:"certificateKeeperName,omitempty"`
	ManufacturerInfo              PositionTradeNameManufacturerInfo     `xml:"manufacturerInfo,omitempty" bson:"manufacturerInfo,omitempty"`
	PackagingsInfo                PositionTradeNamePackagingsInfo       `xml:"packagingsInfo,omitempty" bson:"packagingsInfo,omitempty"`
}

// PositionTradeNameTradeInfo Торговое наименование (ТН) лекарственного препарата. Игнорируется при приеме, автоматически заполняется при передаче из справочника "Лекарственные препараты" (блок MNNInfo\positionsTradeName\positionTradeName\tradeInfo документа nsiFarmDrugDictionary)
type PositionTradeNameTradeInfo struct {
	TradeName string `xml:"tradeName,omitempty" bson:"tradeName,omitempty"`
}

// PositionTradeNameMedicamentalFormInfo Лекарственная форма.   Игнорируется при приеме, автоматически заполняется при передаче из справочника "Лекарственные препараты" (блок MNNInfo\positionsTradeName\positionTradeName\medicamentalFormInfo документа nsiFarmDrugDictionary)
type PositionTradeNameMedicamentalFormInfo struct {
	MedicamentalFormName string `xml:"medicamentalFormName,omitempty" bson:"medicamentalFormName,omitempty"`
}

// PositionTradeNameDosageInfo Дозировка. Игнорируется при приеме, автоматически заполняется при передаче из справочника "Лекарственные препараты" (блок MNNInfo\positionsTradeName\positionTradeName\dosagesInfo документа nsiFarmDrugDictionary)
type PositionTradeNameDosageInfo struct {
	DosageName      string      `xml:"dosageName,omitempty" bson:"dosageName,omitempty"`
	DosageOKEI      ZfcsOKEIRef `xml:"dosageOKEI,omitempty" bson:"dosageOKEI,omitempty"`
	DosageValue     float64     `xml:"dosageValue,omitempty" bson:"dosageValue,omitempty"`
	DosageGRLSValue string      `xml:"dosageGRLSValue,omitempty" bson:"dosageGRLSValue,omitempty"`
}

// PositionTradeNameManufacturerInfo Производитель лекарственного препарата. Игнорируется при приеме, автоматически заполняется при передаче из справочника "Лекарственные препараты" (блок MNNInfo\positionsTradeName\positionTradeName\manufacturerInfo документа
type PositionTradeNameManufacturerInfo struct {
	ManufacturerOKSM ZfcsOKSMRef `xml:"manufacturerOKSM,omitempty" bson:"manufacturerOKSM,omitempty"`
	ManufacturerName string      `xml:"manufacturerName,omitempty" bson:"manufacturerName,omitempty"`
}

// PositionTradeNamePackagingsInfo Сведения об упаковках
type PositionTradeNamePackagingsInfo struct {
	PackagingInfo PackagingsInfoPackagingInfo `xml:"packagingInfo,omitempty" bson:"packagingInfo,omitempty"`
}

// PackagingsInfoPackagingInfo Сведения об упаковке
type PackagingsInfoPackagingInfo struct {
	PrimaryPackagingInfo    PackagingInfoPrimaryPackagingInfo `xml:"primaryPackagingInfo,omitempty" bson:"primaryPackagingInfo,omitempty"`
	Packaging1Quantity      float64                           `xml:"packaging1Quantity,omitempty" bson:"packaging1Quantity,omitempty"`
	Packaging2Quantity      int64                             `xml:"packaging2Quantity,omitempty" bson:"packaging2Quantity,omitempty"`
	SumaryPackagingQuantity float64                           `xml:"sumaryPackagingQuantity,omitempty" bson:"sumaryPackagingQuantity,omitempty"`
	Completeness            string                            `xml:"completeness,omitempty" bson:"completeness,omitempty"`
	DrugQuantity            float64                           `xml:"drugQuantity,omitempty" bson:"drugQuantity,omitempty"`
}

// PackagingInfoPrimaryPackagingInfo Сведения о первичной упаковке по справочнику "Лекарственные препараты" (блок MNNInfo\positionsTradeName\positionTradeName\packagingsInfo\packagingInfo\primaryPackagingInfo  документа nsiFarmDrugDictionary). При приеме контролируется принадлежность вида с указанным наименованием к ТН с кодом, указанным в поле positionTradeNameExternalCode
type PackagingInfoPrimaryPackagingInfo struct {
	PrimaryPackagingName string `xml:"primaryPackagingName,omitempty" bson:"primaryPackagingName,omitempty"`
}

// DrugInfoUsingReferenceInfoMNNInfo Международное, группировочное или химическое наименование лекарственного препарата (МНН). Устарело, не применяется
type DrugInfoUsingReferenceInfoMNNInfo struct {
	MNNExternalCode string `xml:"MNNExternalCode,omitempty" bson:"MNNExternalCode,omitempty"`
	MNNDrugCode     string `xml:"MNNDrugCode,omitempty" bson:"MNNDrugCode,omitempty"`
	MNNName         string `xml:"MNNName,omitempty" bson:"MNNName,omitempty"`
}

// DrugInfoUsingReferenceInfoExpirationDateCustomFormatInfo Срок годности (годен до) в пользовательском формате
type DrugInfoUsingReferenceInfoExpirationDateCustomFormatInfo struct {
	ExpirationDateMonthYear ExpirationDateCustomFormatInfoExpirationDateMonthYear `xml:"expirationDateMonthYear,omitempty" bson:"expirationDateMonthYear,omitempty"`
	ExpirationDate          xsd.Date                                              `xml:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
}

// ExpirationDateCustomFormatInfoExpirationDateMonthYear Срок годности в формате "месяц-год"
type ExpirationDateCustomFormatInfoExpirationDateMonthYear struct {
	Month byte  `xml:"month,omitempty" bson:"month,omitempty"`
	Year  int64 `xml:"year,omitempty" bson:"year,omitempty"`
}

// DrugInfoUsingReferenceInfoPositionsTradeName Позиции по торговому наименованию лекарственного средства. Устарело, не применяется
type DrugInfoUsingReferenceInfoPositionsTradeName struct {
	PositionTradeName []PositionsTradeNamePositionTradeName `xml:"positionTradeName,omitempty" bson:"positionTradeName,omitempty"`
}

// ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingTextForm Информация о лекарственных препаратах формируется в текстовой форме
type ZfcsContract2015DrugPurchaseInfoTypeDrugInfoUsingTextForm struct {
	MNNsInfo                       DrugInfoUsingTextFormMNNsInfo                       `xml:"MNNsInfo,omitempty" bson:"MNNsInfo,omitempty"`
	MNNInfo                        DrugInfoUsingTextFormMNNInfo                        `xml:"MNNInfo,omitempty" bson:"MNNInfo,omitempty"`
	ExpirationDateCustomFormatInfo DrugInfoUsingTextFormExpirationDateCustomFormatInfo `xml:"expirationDateCustomFormatInfo,omitempty" bson:"expirationDateCustomFormatInfo,omitempty"`
	IsZNVLP                        bool                                                `xml:"isZNVLP,omitempty" bson:"isZNVLP,omitempty"`
	PositionsTradeName             DrugInfoUsingTextFormPositionsTradeName             `xml:"positionsTradeName,omitempty" bson:"positionsTradeName,omitempty"`
}

// DrugInfoUsingTextFormMNNsInfo Международные, группировочные или химические наименования лекарственных препаратов (МНН). При приеме контролируется принадлежность всех МНН в блоке одному типу (т.е. все МНН в блоке должны иметь одинаковое наименование)
type DrugInfoUsingTextFormMNNsInfo struct {
	MNNInfo []MNNsInfoMNNInfo `xml:"MNNInfo,omitempty" bson:"MNNInfo,omitempty"`
}

// DrugInfoUsingTextFormMNNInfo Международное, группировочное или химическое наименование лекарственного препарата (МНН). Устарело, не применяется
type DrugInfoUsingTextFormMNNInfo struct {
	MNNName string `xml:"MNNName,omitempty" bson:"MNNName,omitempty"`
}

// DrugInfoUsingTextFormExpirationDateCustomFormatInfo Срок годности (годен до) в пользовательском формате
type DrugInfoUsingTextFormExpirationDateCustomFormatInfo struct {
	ExpirationDateMonthYear ExpirationDateCustomFormatInfoExpirationDateMonthYear `xml:"expirationDateMonthYear,omitempty" bson:"expirationDateMonthYear,omitempty"`
	ExpirationDate          xsd.Date                                              `xml:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
}

// DrugInfoUsingTextFormPositionsTradeName Позиции по торговому наименованию лекарственного средства. Устарело, не применяется
type DrugInfoUsingTextFormPositionsTradeName struct {
	PositionTradeName []PositionsTradeNamePositionTradeName `xml:"positionTradeName,omitempty" bson:"positionTradeName,omitempty"`
}

// ZfcsOKEIRef
type ZfcsOKEIRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsOKSMRef
type ZfcsOKSMRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsOKPDRef
type ZfcsOKPDRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsKTRURef
type ZfcsKTRURef struct {
	Code      string `xml:"code,omitempty" bson:"code,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	VersionID int64  `xml:"versionId,omitempty" bson:"versionId,omitempty"`
}

// ZfcsContract2015SupplierType
type ZfcsContract2015SupplierType struct {
	LegalEntityRF                         ZfcsContract2015SupplierTypeLegalEntityRF                         `xml:"legalEntityRF,omitempty" bson:"legalEntityRF,omitempty"`
	LegalEntityForeignState               ZfcsContract2015SupplierTypeLegalEntityForeignState               `xml:"legalEntityForeignState,omitempty" bson:"legalEntityForeignState,omitempty"`
	IndividualPersonRF                    ZfcsContract2015SupplierTypeIndividualPersonRF                    `xml:"individualPersonRF,omitempty" bson:"individualPersonRF,omitempty"`
	IndividualPersonForeignState          ZfcsContract2015SupplierTypeIndividualPersonForeignState          `xml:"individualPersonForeignState,omitempty" bson:"individualPersonForeignState,omitempty"`
	IndividualPersonRFisCulture           ZfcsContract2015SupplierTypeIndividualPersonRFisCulture           `xml:"individualPersonRFisCulture,omitempty" bson:"individualPersonRFisCulture,omitempty"`
	IndividualPersonForeignStateisCulture ZfcsContract2015SupplierTypeIndividualPersonForeignStateisCulture `xml:"individualPersonForeignStateisCulture,omitempty" bson:"individualPersonForeignStateisCulture,omitempty"`
}

// ZfcsContract2015SupplierTypeLegalEntityRF Юридическое лицо РФ
type ZfcsContract2015SupplierTypeLegalEntityRF struct {
	LegalForm        ZfcsOkopfRef          `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	FullName         string                `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName        string                `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	FirmName         string                `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Status           string                `xml:"status,omitempty" bson:"status,omitempty"`
	ContractPrice    string                `xml:"contractPrice,omitempty" bson:"contractPrice,omitempty"`
	OKPO             string                `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	INN              string                `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP              string                `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	RegistrationDate xsd.Date              `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	OKTMO            ZfcsOKTMORef          `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address          string                `xml:"address,omitempty" bson:"address,omitempty"`
	ContactInfo      ZfcsContactPersonType `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	ContactEMail     string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// ZfcsContract2015SupplierTypeLegalEntityForeignState Юридическое лицо иностранного государства
type ZfcsContract2015SupplierTypeLegalEntityForeignState struct {
	FullName                string                                         `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName               string                                         `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	FirmName                string                                         `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	FullNameLat             string                                         `xml:"fullNameLat,omitempty" bson:"fullNameLat,omitempty"`
	TaxPayerCode            string                                         `xml:"taxPayerCode,omitempty" bson:"taxPayerCode,omitempty"`
	RegisterInRFTaxBodies   LegalEntityForeignStateRegisterInRFTaxBodies   `xml:"registerInRFTaxBodies,omitempty" bson:"registerInRFTaxBodies,omitempty"`
	PlaceOfStayInRegCountry LegalEntityForeignStatePlaceOfStayInRegCountry `xml:"placeOfStayInRegCountry,omitempty" bson:"placeOfStayInRegCountry,omitempty"`
	PlaceOfStayInRF         LegalEntityForeignStatePlaceOfStayInRF         `xml:"placeOfStayInRF,omitempty" bson:"placeOfStayInRF,omitempty"`
}

// LegalEntityForeignStateRegisterInRFTaxBodies Поставщик состоит на учете в налоговых органах на территории РФ
type LegalEntityForeignStateRegisterInRFTaxBodies struct {
	INN              string   `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP              string   `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	RegistrationDate xsd.Date `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
}

// LegalEntityForeignStatePlaceOfStayInRegCountry Место нахождления в стране регистрации
type LegalEntityForeignStatePlaceOfStayInRegCountry struct {
	Country      ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
	Address      string         `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string         `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string         `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// LegalEntityForeignStatePlaceOfStayInRF Наличие у поставщика места пребывания на территории РФ
type LegalEntityForeignStatePlaceOfStayInRF struct {
	OKTMO        ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address      string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// ZfcsContract2015SupplierTypeIndividualPersonRF Физическое лицо РФ
type ZfcsContract2015SupplierTypeIndividualPersonRF struct {
	LastName         string       `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName        string       `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName       string       `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	INN              string       `xml:"INN,omitempty" bson:"INN,omitempty"`
	IsIP             bool         `xml:"isIP,omitempty" bson:"isIP,omitempty"`
	RegistrationDate xsd.Date     `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	OKTMO            ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address          string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail     string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	IsCulture        bool         `xml:"isCulture,omitempty" bson:"isCulture,omitempty"`
}

// ZfcsContract2015SupplierTypeIndividualPersonForeignState Физическое лицо иностранного государства
type ZfcsContract2015SupplierTypeIndividualPersonForeignState struct {
	LastName                string                                              `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName               string                                              `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName              string                                              `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	LastNameLat             string                                              `xml:"lastNameLat,omitempty" bson:"lastNameLat,omitempty"`
	FirstNameLat            string                                              `xml:"firstNameLat,omitempty" bson:"firstNameLat,omitempty"`
	MiddleNameLat           string                                              `xml:"middleNameLat,omitempty" bson:"middleNameLat,omitempty"`
	TaxPayerCode            string                                              `xml:"taxPayerCode,omitempty" bson:"taxPayerCode,omitempty"`
	RegisterInRFTaxBodies   IndividualPersonForeignStateRegisterInRFTaxBodies   `xml:"registerInRFTaxBodies,omitempty" bson:"registerInRFTaxBodies,omitempty"`
	PlaceOfStayInRegCountry IndividualPersonForeignStatePlaceOfStayInRegCountry `xml:"placeOfStayInRegCountry,omitempty" bson:"placeOfStayInRegCountry,omitempty"`
	PlaceOfStayInRF         IndividualPersonForeignStatePlaceOfStayInRF         `xml:"placeOfStayInRF,omitempty" bson:"placeOfStayInRF,omitempty"`
	IsCulture               bool                                                `xml:"isCulture,omitempty" bson:"isCulture,omitempty"`
}

// IndividualPersonForeignStateRegisterInRFTaxBodies Поставщик состоит на учете в налоговых органах на территории РФ
type IndividualPersonForeignStateRegisterInRFTaxBodies struct {
	INN              string   `xml:"INN,omitempty" bson:"INN,omitempty"`
	RegistrationDate xsd.Date `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
}

// IndividualPersonForeignStatePlaceOfStayInRegCountry Место нахождления в стране регистрации
type IndividualPersonForeignStatePlaceOfStayInRegCountry struct {
	Country      ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
	Address      string         `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string         `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string         `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// IndividualPersonForeignStatePlaceOfStayInRF Наличие у поставщика места пребывания на территории РФ
type IndividualPersonForeignStatePlaceOfStayInRF struct {
	OKTMO        ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address      string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// ZfcsContract2015SupplierTypeIndividualPersonRFisCulture Физическое лицо РФ. Поставщик культурных ценностей
type ZfcsContract2015SupplierTypeIndividualPersonRFisCulture struct {
	LastName         string       `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName        string       `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName       string       `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	INN              string       `xml:"INN,omitempty" bson:"INN,omitempty"`
	IsIP             bool         `xml:"isIP,omitempty" bson:"isIP,omitempty"`
	RegistrationDate xsd.Date     `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	OKTMO            ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address          string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail     string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// ZfcsContract2015SupplierTypeIndividualPersonForeignStateisCulture Физическое лицо иностранного государства. Поставщик культурных ценностей
type ZfcsContract2015SupplierTypeIndividualPersonForeignStateisCulture struct {
	LastName                string                                                       `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName               string                                                       `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName              string                                                       `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	LastNameLat             string                                                       `xml:"lastNameLat,omitempty" bson:"lastNameLat,omitempty"`
	FirstNameLat            string                                                       `xml:"firstNameLat,omitempty" bson:"firstNameLat,omitempty"`
	MiddleNameLat           string                                                       `xml:"middleNameLat,omitempty" bson:"middleNameLat,omitempty"`
	TaxPayerCode            string                                                       `xml:"taxPayerCode,omitempty" bson:"taxPayerCode,omitempty"`
	RegisterInRFTaxBodies   IndividualPersonForeignStateisCultureRegisterInRFTaxBodies   `xml:"registerInRFTaxBodies,omitempty" bson:"registerInRFTaxBodies,omitempty"`
	PlaceOfStayInRegCountry IndividualPersonForeignStateisCulturePlaceOfStayInRegCountry `xml:"placeOfStayInRegCountry,omitempty" bson:"placeOfStayInRegCountry,omitempty"`
	PlaceOfStayInRF         IndividualPersonForeignStateisCulturePlaceOfStayInRF         `xml:"placeOfStayInRF,omitempty" bson:"placeOfStayInRF,omitempty"`
}

// IndividualPersonForeignStateisCultureRegisterInRFTaxBodies Поставщик состоит на учете в налоговых органах на территории РФ
type IndividualPersonForeignStateisCultureRegisterInRFTaxBodies struct {
	INN              string   `xml:"INN,omitempty" bson:"INN,omitempty"`
	RegistrationDate xsd.Date `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
}

// IndividualPersonForeignStateisCulturePlaceOfStayInRegCountry Место нахождления в стране регистрации
type IndividualPersonForeignStateisCulturePlaceOfStayInRegCountry struct {
	Country      ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
	Address      string         `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string         `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string         `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// IndividualPersonForeignStateisCulturePlaceOfStayInRF Наличие у поставщика места пребывания на территории РФ
type IndividualPersonForeignStateisCulturePlaceOfStayInRF struct {
	OKTMO        ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address      string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// ZfcsOkopfRef
type ZfcsOkopfRef struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ZfcsContactPersonType
type ZfcsContactPersonType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
}

// ZfcsCountryRef
type ZfcsCountryRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsContractPrintFormType
type ZfcsContractPrintFormType struct {
	URL          string                               `xml:"url,omitempty" bson:"url,omitempty"`
	DocRegNumber string                               `xml:"docRegNumber,omitempty" bson:"docRegNumber,omitempty"`
	Signature    []ZfcsContractPrintFormTypeSignature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsContractPrintFormTypeSignature Электронная подпись печатной формы
type ZfcsContractPrintFormTypeSignature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsExtPrintFormType
type ZfcsExtPrintFormType struct {
	Signature                ZfcsExtPrintFormTypeSignature                `xml:"signature,omitempty" bson:"signature,omitempty"`
	FileType                 string                                       `xml:"fileType,omitempty" bson:"fileType,omitempty"`
	ControlPersonalSignature ZfcsExtPrintFormTypeControlPersonalSignature `xml:"controlPersonalSignature,omitempty" bson:"controlPersonalSignature,omitempty"`
	Content                  []byte                                       `xml:"content,omitempty" bson:"content,omitempty"`
	URL                      string                                       `xml:"url,omitempty" bson:"url,omitempty"`
}

// ZfcsExtPrintFormTypeSignature Электронная подпись электронного документа
type ZfcsExtPrintFormTypeSignature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsExtPrintFormTypeControlPersonalSignature Электронная подпись электронного документа лицом, уполномоченным на проведение контроля в соответствии с ч.5 ст.99 закона №44-ФЗ
type ZfcsExtPrintFormTypeControlPersonalSignature struct {
	Type                     string `xml:"type,attr" bson:"type,omitempty"`
	ControlPersonalSignature []byte `xml:"controlPersonalSignature,omitempty" bson:"controlPersonalSignature,omitempty"`
}

// ZfcsContract2015DocumentInfo
type ZfcsContract2015DocumentInfo struct {
	DocumentName string   `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
}
