package notification

import "gitlab.com/garanteka/parsexsd/xsd"

// FcsNotificationEF Извещение о проведении ЭА (электронный аукцион);
//внесение изменений
type FcsNotificationEF struct {
}

// ZfcsNotificationEFType
type ZfcsNotificationEFType struct {
	SchemeVersion       string                                    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                                     `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                                    `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                                    `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DirectDate          xsd.Date                                  `xml:"directDate,omitempty" bson:"directDate,omitempty"`
	DocPublishDate      xsd.Date                                  `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                                    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                                    `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType                      `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                                    `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible ZfcsNotificationEFTypePurchaseResponsible `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType                        `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                                      `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ETP                 ZfcsETPType                               `xml:"ETP,omitempty" bson:"ETP,omitempty"`
	ProcedureInfo       ZfcsNotificationEFTypeProcedureInfo       `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                 ZfcsNotificationEFTypeLot                 `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType                    `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType          `xml:"modification,omitempty" bson:"modification,omitempty"`
	IsProcessed         bool                                      `xml:"is_processed,omitempty" bson:"is_processed,omitempty"`
}

// ZfcsNotificationEFTypePurchaseResponsible Информация об организации, осуществляющей закупку
type ZfcsNotificationEFTypePurchaseResponsible struct {
	ResponsibleOrg     ZfcsPurchaseOrganizationType `xml:"responsibleOrg,omitempty" bson:"responsibleOrg,omitempty"`
	ResponsibleRole    string                       `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
	ResponsibleInfo    ZfcsContactInfoType          `xml:"responsibleInfo,omitempty" bson:"responsibleInfo,omitempty"`
	SpecializedOrg     ZfcsPurchaseOrganizationType `xml:"specializedOrg,omitempty" bson:"specializedOrg,omitempty"`
	LastSpecializedOrg ZfcsPurchaseOrganizationType `xml:"lastSpecializedOrg,omitempty" bson:"lastSpecializedOrg,omitempty"`
}

// ZfcsNotificationEFTypeProcedureInfo Информация о процедуре закупки
type ZfcsNotificationEFTypeProcedureInfo struct {
	Collecting ZfcsPurchaseProcedureCollectingType `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Scoring    ProcedureInfoScoring                `xml:"scoring,omitempty" bson:"scoring,omitempty"`
	Bidding    ProcedureInfoBidding                `xml:"bidding,omitempty" bson:"bidding,omitempty"`
}

// ProcedureInfoScoring Информация о процедуре рассмотрения и оценки заявок на участие в аукционе
type ProcedureInfoScoring struct {
	Date xsd.Date `xml:"date,omitempty" bson:"date,omitempty"`
}

// ProcedureInfoBidding Информация о процедуре проведения аукциона в электронной форме
type ProcedureInfoBidding struct {
	Date    xsd.Date `xml:"date,omitempty" bson:"date,omitempty"`
	AddInfo string   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsNotificationEFTypeLot Лот извещения
type ZfcsNotificationEFTypeLot struct {
	MaxPrice                string                                  `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	PriceFormula            string                                  `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	StandardContractNumber  string                                  `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	Currency                ZfcsCurrencyRef                         `xml:"currency,omitempty" bson:"currency,omitempty"`
	FinanceSource           string                                  `xml:"financeSource,omitempty" bson:"financeSource,omitempty"`
	InterbudgetaryTransfer  bool                                    `xml:"interbudgetaryTransfer,omitempty" bson:"interbudgetaryTransfer,omitempty"`
	QuantityUndefined       bool                                    `xml:"quantityUndefined,omitempty" bson:"quantityUndefined,omitempty"`
	CustomerRequirements    LotCustomerRequirements                 `xml:"customerRequirements,omitempty" bson:"customerRequirements,omitempty"`
	Preferenses             LotPreferenses                          `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements            LotRequirements                         `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	Restrictions            LotRestrictions                         `xml:"restrictions,omitempty" bson:"restrictions,omitempty"`
	RestrictInfo            string                                  `xml:"restrictInfo,omitempty" bson:"restrictInfo,omitempty"`
	RestrictForeignsInfo    string                                  `xml:"restrictForeignsInfo,omitempty" bson:"restrictForeignsInfo,omitempty"`
	AddInfo                 string                                  `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	NoPublicDiscussion      bool                                    `xml:"noPublicDiscussion,omitempty" bson:"noPublicDiscussion,omitempty"`
	PublicDiscussion        LotPublicDiscussion                     `xml:"publicDiscussion,omitempty" bson:"publicDiscussion,omitempty"`
	MustPublicDiscussion    bool                                    `xml:"mustPublicDiscussion,omitempty" bson:"mustPublicDiscussion,omitempty"`
	PurchaseObjects         LotPurchaseObjects                      `xml:"purchaseObjects,omitempty" bson:"purchaseObjects,omitempty"`
	DrugPurchaseObjectsInfo ZfcsPurchaseDrugPurchaseObjectsInfoType `xml:"drugPurchaseObjectsInfo,omitempty" bson:"drugPurchaseObjectsInfo,omitempty"`
}

// LotCustomerRequirements Требования заказчиков
type LotCustomerRequirements struct {
	CustomerRequirement []CustomerRequirementsCustomerRequirement `xml:"customerRequirement,omitempty" bson:"customerRequirement,omitempty"`
}

// CustomerRequirementsCustomerRequirement Требование заказчика
type CustomerRequirementsCustomerRequirement struct {
	Customer                  ZfcsOrganizationRef         `xml:"customer,omitempty" bson:"customer,omitempty"`
	MaxPrice                  string                      `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	DeliveryTerm              string                      `xml:"deliveryTerm,omitempty" bson:"deliveryTerm,omitempty"`
	ApplicationGuarantee      ZfcsPaymentInfoType         `xml:"applicationGuarantee,omitempty" bson:"applicationGuarantee,omitempty"`
	ContractGuarantee         ZfcsPaymentInfoType         `xml:"contractGuarantee,omitempty" bson:"contractGuarantee,omitempty"`
	PurchaseCode              string                      `xml:"purchaseCode,omitempty" bson:"purchaseCode,omitempty"`
	TenderPlanInfo            ZfcsTendePlanInfoType       `xml:"tenderPlanInfo,omitempty" bson:"tenderPlanInfo,omitempty"`
	BudgetFinancings          ZfcsBudgetFinancingsType    `xml:"budgetFinancings,omitempty" bson:"budgetFinancings,omitempty"`
	NonbudgetFinancings       ZfcsNonbudgetFinancingsType `xml:"nonbudgetFinancings,omitempty" bson:"nonbudgetFinancings,omitempty"`
	BOInfo                    ZfcsPurchaseBOInfoType      `xml:"BOInfo,omitempty" bson:"BOInfo,omitempty"`
	PurchaseObjectDescription string                      `xml:"purchaseObjectDescription,omitempty" bson:"purchaseObjectDescription,omitempty"`
	DeliveryPlace             string                      `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	KladrPlaces               ZfcsKladrPlacesType         `xml:"kladrPlaces,omitempty" bson:"kladrPlaces,omitempty"`
}

// LotPreferenses Преимущества
type LotPreferenses struct {
	Preferense []ZfcsPreferenseType `xml:"preferense,omitempty" bson:"preferense,omitempty"`
}

// LotRequirements Требования
type LotRequirements struct {
	Requirement []ZfcsRequirementType `xml:"requirement,omitempty" bson:"requirement,omitempty"`
}

// LotRestrictions Ограничения
type LotRestrictions struct {
	Restriction []ZfcsRestrictionType `xml:"restriction,omitempty" bson:"restriction,omitempty"`
}

// LotPublicDiscussion Общественное обсуждение крупных закупок (для печатной формы). Игнорируется при приеме. Автоматически заполняется при передаче
type LotPublicDiscussion struct {
	Place                string                               `xml:"place,omitempty" bson:"place,omitempty"`
	Number               string                               `xml:"number,omitempty" bson:"number,omitempty"`
	OrganizationCh5St15  bool                                 `xml:"organizationCh5St15,omitempty" bson:"organizationCh5St15,omitempty"`
	Href                 string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PublicDiscussion2017 PublicDiscussionPublicDiscussion2017 `xml:"publicDiscussion2017,omitempty" bson:"publicDiscussion2017,omitempty"`
}

// PublicDiscussionPublicDiscussion2017 Информация об общественном обсуждении по лоту закупки с 01.01.2017
type PublicDiscussionPublicDiscussion2017 struct {
	PublicDiscussionLargePurchasePhase2 PublicDiscussion2017PublicDiscussionLargePurchasePhase2 `xml:"publicDiscussionLargePurchasePhase2,omitempty" bson:"publicDiscussionLargePurchasePhase2,omitempty"`
}

// PublicDiscussion2017PublicDiscussionLargePurchasePhase2 Информация о втором этапе ООКЗ
type PublicDiscussion2017PublicDiscussionLargePurchasePhase2 struct {
	ProtocolDate              xsd.Date                                       `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	ProtocolPublishDate       xsd.Date                                       `xml:"protocolPublishDate,omitempty" bson:"protocolPublishDate,omitempty"`
	PublicDiscussionPhase2Num string                                         `xml:"publicDiscussionPhase2Num,omitempty" bson:"publicDiscussionPhase2Num,omitempty"`
	HrefPhase2                string                                         `xml:"hrefPhase2,omitempty" bson:"hrefPhase2,omitempty"`
	Attachments               PublicDiscussionLargePurchasePhase2Attachments `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// PublicDiscussionLargePurchasePhase2Attachments Протокол этапа
type PublicDiscussionLargePurchasePhase2Attachments struct {
	Attachment []AttachmentsAttachment `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// AttachmentsAttachment Документ
type AttachmentsAttachment struct {
	PublishedContentID string                `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string                `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	FileSize           string                `xml:"fileSize,omitempty" bson:"fileSize,omitempty"`
	DocDescription     string                `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	DocDate            xsd.Date              `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	CryptoSigns        AttachmentCryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string                `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string                `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte                `xml:"content,omitempty" bson:"content,omitempty"`
	PlacingDate        xsd.Date              `xml:"placingDate,omitempty" bson:"placingDate,omitempty"`
}

// AttachmentCryptoSigns Электронная подпись документа
type AttachmentCryptoSigns struct {
	Signature []CryptoSignsSignature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// CryptoSignsSignature Электронная подпись
type CryptoSignsSignature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// LotPurchaseObjects Объекты закупки
type LotPurchaseObjects struct {
	PurchaseObject []PurchaseObjectsPurchaseObject `xml:"purchaseObject,omitempty" bson:"purchaseObject,omitempty"`
	TotalSum       string                          `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// PurchaseObjectsPurchaseObject
type PurchaseObjectsPurchaseObject struct {
	Name               string                           `xml:"name,omitempty" bson:"name,omitempty"`
	OKEI               ZfcsContractOKEIType             `xml:"OKEI,omitempty" bson:"OKEI,omitempty"`
	CustomerQuantities PurchaseObjectCustomerQuantities `xml:"customerQuantities,omitempty" bson:"customerQuantities,omitempty"`
	Price              string                           `xml:"price,omitempty" bson:"price,omitempty"`
	Quantity           PurchaseObjectQuantity           `xml:"quantity,omitempty" bson:"quantity,omitempty"`
	Sum                string                           `xml:"sum,omitempty" bson:"sum,omitempty"`
	OKPD               ZfcsOKPDRef                      `xml:"OKPD,omitempty" bson:"OKPD,omitempty"`
	OKPD2              PurchaseObjectOKPD2              `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
	KTRU               PurchaseObjectKTRU               `xml:"KTRU,omitempty" bson:"KTRU,omitempty"`
}

// PurchaseObjectCustomerQuantities Количество по заказчикам
type PurchaseObjectCustomerQuantities struct {
	CustomerQuantity []CustomerQuantitiesCustomerQuantity `xml:"customerQuantity,omitempty" bson:"customerQuantity,omitempty"`
}

// CustomerQuantitiesCustomerQuantity
type CustomerQuantitiesCustomerQuantity struct {
	Customer ZfcsOrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
	Quantity float64             `xml:"quantity,omitempty" bson:"quantity,omitempty"`
}

// PurchaseObjectQuantity Общее количество по объекту закупки
type PurchaseObjectQuantity struct {
	Value     float64 `xml:"value,omitempty" bson:"value,omitempty"`
	Undefined bool    `xml:"undefined,omitempty" bson:"undefined,omitempty"`
}

// PurchaseObjectOKPD2 Классификация по ОКПД2
type PurchaseObjectOKPD2 struct {
	Code            string               `xml:"code,omitempty" bson:"code,omitempty"`
	Name            string               `xml:"name,omitempty" bson:"name,omitempty"`
	Characteristics OKPD2Characteristics `xml:"characteristics,omitempty" bson:"characteristics,omitempty"`
}

// OKPD2Characteristics Характеристики товара, работы, услуги. Игнорируется при приеме. Добавлено на развитие
type OKPD2Characteristics struct {
	CharacteristicsUsingTextForm []ZfcsTenderPlan2017ManualKtruCharacteristicType `xml:"characteristicsUsingTextForm,omitempty" bson:"characteristicsUsingTextForm,omitempty"`
}

// PurchaseObjectKTRU Классификация по КТРУ
type PurchaseObjectKTRU struct {
	Code            string              `xml:"code,omitempty" bson:"code,omitempty"`
	Name            string              `xml:"name,omitempty" bson:"name,omitempty"`
	VersionID       int64               `xml:"versionId,omitempty" bson:"versionId,omitempty"`
	Characteristics KTRUCharacteristics `xml:"characteristics,omitempty" bson:"characteristics,omitempty"`
}

// KTRUCharacteristics Характеристики товара, работы, услуги позиции КТРУ.
//Контролируется обязательное заполнение хотя бы одного из дочерних блоков characteristicsUsingReferenceInfo и/или characteristicsUsingTextForm
type KTRUCharacteristics struct {
	CharacteristicsUsingReferenceInfo []ZfcsTenderPlan2017RefKtruCharacteristicType    `xml:"characteristicsUsingReferenceInfo,omitempty" bson:"characteristicsUsingReferenceInfo,omitempty"`
	CharacteristicsUsingTextForm      []ZfcsTenderPlan2017ManualKtruCharacteristicType `xml:"characteristicsUsingTextForm,omitempty" bson:"characteristicsUsingTextForm,omitempty"`
	AddCharacteristicInfoReason       string                                           `xml:"addCharacteristicInfoReason,omitempty" bson:"addCharacteristicInfoReason,omitempty"`
}

// ZfcsPrintFormType
type ZfcsPrintFormType struct {
	URL       string                     `xml:"url,omitempty" bson:"url,omitempty"`
	Signature ZfcsPrintFormTypeSignature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsPrintFormTypeSignature Электронная подпись печатной формы
type ZfcsPrintFormTypeSignature struct {
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

// ZfcsPurchaseOrganizationType
type ZfcsPurchaseOrganizationType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
}

// ZfcsContactInfoType
type ZfcsContactInfoType struct {
	OrgPostAddress string                `xml:"orgPostAddress,omitempty" bson:"orgPostAddress,omitempty"`
	OrgFactAddress string                `xml:"orgFactAddress,omitempty" bson:"orgFactAddress,omitempty"`
	ContactPerson  ZfcsContactPersonType `xml:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	ContactEMail   string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone   string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax     string                `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AddInfo        string                `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsContactPersonType
type ZfcsContactPersonType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
}

// ZfcsPlacingWayType
type ZfcsPlacingWayType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsETPType
type ZfcsETPType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
	URL  string `xml:"url,omitempty" bson:"url,omitempty"`
}

// ZfcsPurchaseProcedureCollectingType
type ZfcsPurchaseProcedureCollectingType struct {
	StartDate xsd.Date `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	Place     string   `xml:"place,omitempty" bson:"place,omitempty"`
	Order     string   `xml:"order,omitempty" bson:"order,omitempty"`
	EndDate   xsd.Date `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ZfcsCurrencyRef
type ZfcsCurrencyRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsOrganizationRef
type ZfcsOrganizationRef struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsPaymentInfoType
type ZfcsPaymentInfoType struct {
	Amount            string  `xml:"amount,omitempty" bson:"amount,omitempty"`
	Part              float64 `xml:"part,omitempty" bson:"part,omitempty"`
	ProcedureInfo     string  `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	SettlementAccount string  `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string  `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string  `xml:"bik,omitempty" bson:"bik,omitempty"`
}

// ZfcsTendePlanInfoType
type ZfcsTendePlanInfoType struct {
	PlanNumber            string `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	PositionNumber        string `xml:"positionNumber,omitempty" bson:"positionNumber,omitempty"`
	Purchase83st544       bool   `xml:"purchase83st544,omitempty" bson:"purchase83st544,omitempty"`
	Plan2017Number        string `xml:"plan2017Number,omitempty" bson:"plan2017Number,omitempty"`
	Position2017Number    string `xml:"position2017Number,omitempty" bson:"position2017Number,omitempty"`
	Position2017ExtNumber string `xml:"position2017ExtNumber,omitempty" bson:"position2017ExtNumber,omitempty"`
}

// ZfcsBudgetFinancingsType
type ZfcsBudgetFinancingsType struct {
	BudgetFinancing []ZfcsBudgetFinancingsTypeBudgetFinancing `xml:"budgetFinancing,omitempty" bson:"budgetFinancing,omitempty"`
	TotalSum        string                                    `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// ZfcsBudgetFinancingsTypeBudgetFinancing Зпись плана исполнения контракта за счет бюджетных средств
type ZfcsBudgetFinancingsTypeBudgetFinancing struct {
	Year        int64  `xml:"year,omitempty" bson:"year,omitempty"`
	Sum         string `xml:"sum,omitempty" bson:"sum,omitempty"`
	KbkCode     string `xml:"kbkCode,omitempty" bson:"kbkCode,omitempty"`
	KbkCode2016 string `xml:"kbkCode2016,omitempty" bson:"kbkCode2016,omitempty"`
}

// ZfcsNonbudgetFinancingsType
type ZfcsNonbudgetFinancingsType struct {
	NonbudgetFinancing []ZfcsNonbudgetFinancingsTypeNonbudgetFinancing `xml:"nonbudgetFinancing,omitempty" bson:"nonbudgetFinancing,omitempty"`
	TotalSum           string                                          `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// ZfcsNonbudgetFinancingsTypeNonbudgetFinancing Зпись плана исполнения контракта за счет внебюджетных средств
type ZfcsNonbudgetFinancingsTypeNonbudgetFinancing struct {
	Year      int64  `xml:"year,omitempty" bson:"year,omitempty"`
	Sum       string `xml:"sum,omitempty" bson:"sum,omitempty"`
	KosguCode string `xml:"kosguCode,omitempty" bson:"kosguCode,omitempty"`
	KvrCode   string `xml:"kvrCode,omitempty" bson:"kvrCode,omitempty"`
}

// ZfcsPurchaseBOInfoType
type ZfcsPurchaseBOInfoType struct {
	BONumber    string   `xml:"BONumber,omitempty" bson:"BONumber,omitempty"`
	BODate      xsd.Date `xml:"BODate,omitempty" bson:"BODate,omitempty"`
	InputBOFlag string   `xml:"inputBOFlag,omitempty" bson:"inputBOFlag,omitempty"`
}

// ZfcsKladrPlacesType
type ZfcsKladrPlacesType struct {
	KladrPlace []ZfcsKladrPlacesTypeKladrPlace `xml:"kladrPlace,omitempty" bson:"kladrPlace,omitempty"`
}

// ZfcsKladrPlacesTypeKladrPlace Место доставки товара, выполнения работы или оказания услуги по справочнику КЛАДР
type ZfcsKladrPlacesTypeKladrPlace struct {
	DeliveryPlace              string                               `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	NoKladrForRegionSettlement KladrPlaceNoKladrForRegionSettlement `xml:"noKladrForRegionSettlement,omitempty" bson:"noKladrForRegionSettlement,omitempty"`
	Kladr                      ZfcsKladrType                        `xml:"kladr,omitempty" bson:"kladr,omitempty"`
	Country                    ZfcsCountryRef                       `xml:"country,omitempty" bson:"country,omitempty"`
}

// KladrPlaceNoKladrForRegionSettlement КЛАДР не используется для задания района/города и населенного пункта
type KladrPlaceNoKladrForRegionSettlement struct {
	Region     string `xml:"region,omitempty" bson:"region,omitempty"`
	Settlement string `xml:"settlement,omitempty" bson:"settlement,omitempty"`
}

// ZfcsKladrType
type ZfcsKladrType struct {
	KladrType string `xml:"kladrType,omitempty" bson:"kladrType,omitempty"`
	KladrCode string `xml:"kladrCode,omitempty" bson:"kladrCode,omitempty"`
	FullName  string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsCountryRef
type ZfcsCountryRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsPreferenseType
type ZfcsPreferenseType struct {
	Name      string  `xml:"name,omitempty" bson:"name,omitempty"`
	PrefValue float64 `xml:"prefValue,omitempty" bson:"prefValue,omitempty"`
	Code      int64   `xml:"code,omitempty" bson:"code,omitempty"`
	ShortName string  `xml:"shortName,omitempty" bson:"shortName,omitempty"`
}

// ZfcsRequirementType
type ZfcsRequirementType struct {
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	Content   string `xml:"content,omitempty" bson:"content,omitempty"`
	Code      int64  `xml:"code,omitempty" bson:"code,omitempty"`
	ShortName string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
}

// ZfcsRestrictionType
type ZfcsRestrictionType struct {
	ShortName string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	Content   string `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsContractOKEIType
type ZfcsContractOKEIType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	NationalCode string `xml:"nationalCode,omitempty" bson:"nationalCode,omitempty"`
	FullName     string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsOKPDRef
type ZfcsOKPDRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsTenderPlan2017ManualKtruCharacteristicType
type ZfcsTenderPlan2017ManualKtruCharacteristicType struct {
	Name   string                                               `xml:"name,omitempty" bson:"name,omitempty"`
	Type   string                                               `xml:"type,omitempty" bson:"type,omitempty"`
	Values ZfcsTenderPlan2017ManualKtruCharacteristicTypeValues `xml:"values,omitempty" bson:"values,omitempty"`
}

// ZfcsTenderPlan2017ManualKtruCharacteristicTypeValues Допустимые значения характеристики
type ZfcsTenderPlan2017ManualKtruCharacteristicTypeValues struct {
	Value []ZfcsKTRUCharacteristicValueType `xml:"value,omitempty" bson:"value,omitempty"`
}

// ZfcsKTRUCharacteristicValueType
type ZfcsKTRUCharacteristicValueType struct {
	QualityDescription string                                  `xml:"qualityDescription,omitempty" bson:"qualityDescription,omitempty"`
	OKEI               ZfcsOKEIRef                             `xml:"OKEI,omitempty" bson:"OKEI,omitempty"`
	RangeSet           ZfcsKTRUCharacteristicValueTypeRangeSet `xml:"rangeSet,omitempty" bson:"rangeSet,omitempty"`
	ValueSet           ZfcsKTRUCharacteristicValueTypeValueSet `xml:"valueSet,omitempty" bson:"valueSet,omitempty"`
}

// ZfcsKTRUCharacteristicValueTypeRangeSet Набор диапазонов значений характеристик
type ZfcsKTRUCharacteristicValueTypeRangeSet struct {
	ValueRange RangeSetValueRange `xml:"valueRange,omitempty" bson:"valueRange,omitempty"`
}

// RangeSetValueRange Диапазон значений
type RangeSetValueRange struct {
	MinMathNotation string  `xml:"minMathNotation,omitempty" bson:"minMathNotation,omitempty"`
	Min             float64 `xml:"min,omitempty" bson:"min,omitempty"`
	MaxMathNotation string  `xml:"maxMathNotation,omitempty" bson:"maxMathNotation,omitempty"`
	Max             float64 `xml:"max,omitempty" bson:"max,omitempty"`
}

// ZfcsKTRUCharacteristicValueTypeValueSet Набор конкретных значений характеристики
type ZfcsKTRUCharacteristicValueTypeValueSet struct {
	ConcreteValue float64 `xml:"concreteValue,omitempty" bson:"concreteValue,omitempty"`
}

// ZfcsOKEIRef
type ZfcsOKEIRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsTenderPlan2017RefKtruCharacteristicType
type ZfcsTenderPlan2017RefKtruCharacteristicType struct {
	Code   string                                            `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string                                            `xml:"name,omitempty" bson:"name,omitempty"`
	Type   string                                            `xml:"type,omitempty" bson:"type,omitempty"`
	Values ZfcsTenderPlan2017RefKtruCharacteristicTypeValues `xml:"values,omitempty" bson:"values,omitempty"`
}

// ZfcsTenderPlan2017RefKtruCharacteristicTypeValues Допустимые значения характеристики позиции КТРУ
type ZfcsTenderPlan2017RefKtruCharacteristicTypeValues struct {
	Value []ZfcsKTRUCharacteristicValueType `xml:"value,omitempty" bson:"value,omitempty"`
}

// ZfcsPurchaseDrugPurchaseObjectsInfoType
type ZfcsPurchaseDrugPurchaseObjectsInfoType struct {
	DrugPurchaseObjectInfo []ZfcsPurchaseDrugPurchaseObjectsInfoTypeDrugPurchaseObjectInfo `xml:"drugPurchaseObjectInfo,omitempty" bson:"drugPurchaseObjectInfo,omitempty"`
	Total                  string                                                          `xml:"total,omitempty" bson:"total,omitempty"`
}

// ZfcsPurchaseDrugPurchaseObjectsInfoTypeDrugPurchaseObjectInfo Сведения об объекте закупки в том случае, когда объектом закупки является лекарственный препарат
type ZfcsPurchaseDrugPurchaseObjectsInfoTypeDrugPurchaseObjectInfo struct {
	IsZNVLP                      bool                                               `xml:"isZNVLP,omitempty" bson:"isZNVLP,omitempty"`
	DrugQuantityCustomersInfo    DrugPurchaseObjectInfoDrugQuantityCustomersInfo    `xml:"drugQuantityCustomersInfo,omitempty" bson:"drugQuantityCustomersInfo,omitempty"`
	PricePerUnit                 string                                             `xml:"pricePerUnit,omitempty" bson:"pricePerUnit,omitempty"`
	PositionPrice                string                                             `xml:"positionPrice,omitempty" bson:"positionPrice,omitempty"`
	ObjectInfoUsingReferenceInfo DrugPurchaseObjectInfoObjectInfoUsingReferenceInfo `xml:"objectInfoUsingReferenceInfo,omitempty" bson:"objectInfoUsingReferenceInfo,omitempty"`
	ObjectInfoUsingTextForm      DrugPurchaseObjectInfoObjectInfoUsingTextForm      `xml:"objectInfoUsingTextForm,omitempty" bson:"objectInfoUsingTextForm,omitempty"`
}

// DrugPurchaseObjectInfoDrugQuantityCustomersInfo Количество (объем) закупаемого лекарственного препарата в разбивке по заказчикам в основном варианте поставки. Поле total в составе блока рассчитывается автоматически
type DrugPurchaseObjectInfoDrugQuantityCustomersInfo struct {
	DrugQuantityCustomerInfo []DrugQuantityCustomersInfoDrugQuantityCustomerInfo `xml:"drugQuantityCustomerInfo,omitempty" bson:"drugQuantityCustomerInfo,omitempty"`
	Total                    float64                                             `xml:"total,omitempty" bson:"total,omitempty"`
}

// DrugQuantityCustomersInfoDrugQuantityCustomerInfo Количество (объем) закупаемого лекарственного препарата для заказчика
type DrugQuantityCustomersInfoDrugQuantityCustomerInfo struct {
	Customer ZfcsOrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
	Quantity float64             `xml:"quantity,omitempty" bson:"quantity,omitempty"`
}

// DrugPurchaseObjectInfoObjectInfoUsingReferenceInfo Информация о вариантах поставки лекарственных препаратов формируется с использованием справочной информации (не в текстовой форме)
type DrugPurchaseObjectInfoObjectInfoUsingReferenceInfo struct {
	DrugsInfo              ObjectInfoUsingReferenceInfoDrugsInfo              `xml:"drugsInfo,omitempty" bson:"drugsInfo,omitempty"`
	MustSpecifyDrugPackage ObjectInfoUsingReferenceInfoMustSpecifyDrugPackage `xml:"mustSpecifyDrugPackage,omitempty" bson:"mustSpecifyDrugPackage,omitempty"`
}

// ObjectInfoUsingReferenceInfoDrugsInfo Сведения о вариантах поставки лекарственных препаратов
type ObjectInfoUsingReferenceInfoDrugsInfo struct {
	DrugInfo []DrugsInfoDrugInfo `xml:"drugInfo,omitempty" bson:"drugInfo,omitempty"`
}

// DrugsInfoDrugInfo Сведения о варианте поставки лекарственного препарата
type DrugsInfoDrugInfo struct {
	MNNInfo              DrugInfoMNNInfo              `xml:"MNNInfo,omitempty" bson:"MNNInfo,omitempty"`
	TradeInfo            []ZfcsTradeInfoType          `xml:"tradeInfo,omitempty" bson:"tradeInfo,omitempty"`
	MedicamentalFormInfo DrugInfoMedicamentalFormInfo `xml:"medicamentalFormInfo,omitempty" bson:"medicamentalFormInfo,omitempty"`
	DosageInfo           DrugInfoDosageInfo           `xml:"dosageInfo,omitempty" bson:"dosageInfo,omitempty"`
	PackagingInfo        DrugInfoPackagingInfo        `xml:"packagingInfo,omitempty" bson:"packagingInfo,omitempty"`
	ManualUserOKEI       ZfcsOKEIRef                  `xml:"manualUserOKEI,omitempty" bson:"manualUserOKEI,omitempty"`
	BasicUnit            bool                         `xml:"basicUnit,omitempty" bson:"basicUnit,omitempty"`
	DrugQuantity         float64                      `xml:"drugQuantity,omitempty" bson:"drugQuantity,omitempty"`
}

// DrugInfoMNNInfo Международное, группировочное или химическое наименование лекарственного препарата (МНН). При приеме контролируется принадлежность каждого из набора МНН лекарственных препаратов к одному и тому же списку МНН (т.е. МНН в списке должны иметь одни и те же наименования)
type DrugInfoMNNInfo struct {
	MNNExternalCode string `xml:"MNNExternalCode,omitempty" bson:"MNNExternalCode,omitempty"`
	MNNName         string `xml:"MNNName,omitempty" bson:"MNNName,omitempty"`
}

// DrugInfoMedicamentalFormInfo Лекарственная форма. Игнорируется при приеме, автоматически заполняется при передаче по справочнику "Лекарственные препараты"
type DrugInfoMedicamentalFormInfo struct {
	MedicamentalFormName string `xml:"medicamentalFormName,omitempty" bson:"medicamentalFormName,omitempty"`
}

// DrugInfoDosageInfo Дозировка. Игнорируется при приеме, автоматически заполняется при передаче по справочнику "Лекарственные препараты"
type DrugInfoDosageInfo struct {
	DosageGRLSValue string      `xml:"dosageGRLSValue,omitempty" bson:"dosageGRLSValue,omitempty"`
	DosageUserOKEI  ZfcsOKEIRef `xml:"dosageUserOKEI,omitempty" bson:"dosageUserOKEI,omitempty"`
}

// DrugInfoPackagingInfo Сведения об упаковке. В  случае заполнения блока mustSpecifyDrugPackage при приеме контролируется заполненность блока packagingInfo во всех  вариантах поставки лекарственных препаратов
type DrugInfoPackagingInfo struct {
	Packaging1Quantity      float64 `xml:"packaging1Quantity,omitempty" bson:"packaging1Quantity,omitempty"`
	Packaging2Quantity      int64   `xml:"packaging2Quantity,omitempty" bson:"packaging2Quantity,omitempty"`
	SumaryPackagingQuantity float64 `xml:"sumaryPackagingQuantity,omitempty" bson:"sumaryPackagingQuantity,omitempty"`
}

// ObjectInfoUsingReferenceInfoMustSpecifyDrugPackage Необходимо указание сведений об упаковке закупаемого лекарственного препарата. В случае указания блока контролируется заполненность блока packagingInfo во всех  вариантах поставки лекарственных препаратов
type ObjectInfoUsingReferenceInfoMustSpecifyDrugPackage struct {
	SpecifyDrugPackageReason string `xml:"specifyDrugPackageReason,omitempty" bson:"specifyDrugPackageReason,omitempty"`
}

// DrugPurchaseObjectInfoObjectInfoUsingTextForm Информация о вариантах поставки лекарственных препаратов  формируется в текстовой форме
type DrugPurchaseObjectInfoObjectInfoUsingTextForm struct {
	DrugsInfo              ObjectInfoUsingTextFormDrugsInfo              `xml:"drugsInfo,omitempty" bson:"drugsInfo,omitempty"`
	MustSpecifyDrugPackage ObjectInfoUsingTextFormMustSpecifyDrugPackage `xml:"mustSpecifyDrugPackage,omitempty" bson:"mustSpecifyDrugPackage,omitempty"`
}

// ObjectInfoUsingTextFormDrugsInfo Сведения о вариантах поставки лекарственных препаратов
type ObjectInfoUsingTextFormDrugsInfo struct {
	DrugInfo []DrugsInfoDrugInfo `xml:"drugInfo,omitempty" bson:"drugInfo,omitempty"`
}

// ObjectInfoUsingTextFormMustSpecifyDrugPackage Необходимо указание сведений об упаковке закупаемого лекарственного препарата. В случае указания блока контролируется заполненность блока packagingInfo во всех  вариантах поставки лекарственных препаратов
type ObjectInfoUsingTextFormMustSpecifyDrugPackage struct {
	SpecifyDrugPackageReason string `xml:"specifyDrugPackageReason,omitempty" bson:"specifyDrugPackageReason,omitempty"`
}

// ZfcsTradeInfoType
type ZfcsTradeInfoType struct {
	PositionTradeNameExternalCode string `xml:"positionTradeNameExternalCode,omitempty" bson:"positionTradeNameExternalCode,omitempty"`
	TradeName                     string `xml:"tradeName,omitempty" bson:"tradeName,omitempty"`
}

// ZfcsTradeInfoUsingTextFormType
type ZfcsTradeInfoUsingTextFormType struct {
	TradeName string `xml:"tradeName,omitempty" bson:"tradeName,omitempty"`
}

// ZfcsAttachmentListType
type ZfcsAttachmentListType struct {
	Attachment []ZfcsAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsAttachmentType
type ZfcsAttachmentType struct {
	PublishedContentID string                        `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string                        `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	FileSize           string                        `xml:"fileSize,omitempty" bson:"fileSize,omitempty"`
	DocDescription     string                        `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	DocDate            xsd.Date                      `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	CryptoSigns        ZfcsAttachmentTypeCryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string                        `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string                        `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte                        `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsAttachmentTypeCryptoSigns Электронная подпись документа
type ZfcsAttachmentTypeCryptoSigns struct {
	Signature []CryptoSignsSignature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsNotificationModificationType
type ZfcsNotificationModificationType struct {
	ModificationNumber int64                  `xml:"modificationNumber,omitempty" bson:"modificationNumber,omitempty"`
	Info               string                 `xml:"info,omitempty" bson:"info,omitempty"`
	AddInfo            string                 `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Reason             ZfcsPurchaseChangeType `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsPurchaseChangeType
type ZfcsPurchaseChangeType struct {
	ResponsibleDecision   ZfcsPurchaseChangeTypeResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	AuthorityPrescription ZfcsPurchaseChangeTypeAuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
	CourtDecision         ZfcsPurchaseChangeTypeCourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	DiscussionResult      ZfcsPurchaseChangeTypeDiscussionResult      `xml:"discussionResult,omitempty" bson:"discussionResult,omitempty"`
}

// ZfcsPurchaseChangeTypeResponsibleDecision По решению заказчика (организации, осуществляющей определение поставщика для заказчика)
type ZfcsPurchaseChangeTypeResponsibleDecision struct {
	DecisionDate xsd.Date `xml:"decisionDate,omitempty" bson:"decisionDate,omitempty"`
}

// ZfcsPurchaseChangeTypeAuthorityPrescription Предписание органа, уполномоченного на осуществление контроля
type ZfcsPurchaseChangeTypeAuthorityPrescription struct {
	ReestrPrescription   AuthorityPrescriptionReestrPrescription   `xml:"reestrPrescription,omitempty" bson:"reestrPrescription,omitempty"`
	ExternalPrescription AuthorityPrescriptionExternalPrescription `xml:"externalPrescription,omitempty" bson:"externalPrescription,omitempty"`
}

// AuthorityPrescriptionReestrPrescription Данные о предписании, выданном КО
type AuthorityPrescriptionReestrPrescription struct {
	CheckResultNumber  string   `xml:"checkResultNumber,omitempty" bson:"checkResultNumber,omitempty"`
	PrescriptionNumber string   `xml:"prescriptionNumber,omitempty" bson:"prescriptionNumber,omitempty"`
	Foundation         string   `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	AuthorityName      string   `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	DocDate            xsd.Date `xml:"docDate,omitempty" bson:"docDate,omitempty"`
}

// AuthorityPrescriptionExternalPrescription Предписание отсутствует в реестре результатов контроля
type AuthorityPrescriptionExternalPrescription struct {
	AuthorityName string   `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	AuthorityType string   `xml:"authorityType,omitempty" bson:"authorityType,omitempty"`
	DocName       string   `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate       xsd.Date `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber     string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// ZfcsPurchaseChangeTypeCourtDecision Решение судебного органа
type ZfcsPurchaseChangeTypeCourtDecision struct {
	CourtName string   `xml:"courtName,omitempty" bson:"courtName,omitempty"`
	DocName   string   `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate   xsd.Date `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// ZfcsPurchaseChangeTypeDiscussionResult Общественное обсуждение
type ZfcsPurchaseChangeTypeDiscussionResult struct {
	DocName   string   `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate   xsd.Date `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}
