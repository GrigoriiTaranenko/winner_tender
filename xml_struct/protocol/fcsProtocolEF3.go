package protocol

import "gitlab.com/garanteka/parsexsd/xsd"

// FcsProtocolEF3 Протокол подведения итогов электронного аукциона
type FcsProtocolEF3 struct {
}

// ZfcsProtocolEF3Type
type ZfcsProtocolEF3Type struct {
	SchemeVersion            string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                               `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                               `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                               `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                               `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             xsd.Date                             `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 xsd.Date                             `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	DirectDate               xsd.Date                             `xml:"directDate,omitempty" bson:"directDate,omitempty"`
	PublishDate              xsd.Date                             `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType                   `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ZfcsProtocolEF3TypeProtocolPublisher `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType         `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ZfcsProtocolEF3TypeProtocolLot       `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	IsProcessed              bool                                 `xml:"is_processed,omitempty" bson:"is_processed,omitempty"`
}

// ZfcsProtocolEF3TypeProtocolPublisher Информация об организации, разместившей протокол
type ZfcsProtocolEF3TypeProtocolPublisher struct {
	PublisherOrg  ZfcsPurchaseOrganizationType `xml:"publisherOrg,omitempty" bson:"publisherOrg,omitempty"`
	PublisherRole string                       `xml:"publisherRole,omitempty" bson:"publisherRole,omitempty"`
}

// ZfcsProtocolEF3TypeProtocolLot Лот протокола
type ZfcsProtocolEF3TypeProtocolLot struct {
	Applications    ProtocolLotApplications `xml:"applications,omitempty" bson:"applications,omitempty"`
	AbandonedReason ZfcsAbandonedReasonType `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
}

// ProtocolLotApplications Заявки по лоту
type ProtocolLotApplications struct {
	Application []ApplicationsApplication `xml:"application,omitempty" bson:"application,omitempty"`
}

// ApplicationsApplication Заявка
type ApplicationsApplication struct {
	JournalNumber              string                      `xml:"journalNumber,omitempty" bson:"journalNumber,omitempty"`
	AppParticipant             ZfcsParticipantType         `xml:"appParticipant,omitempty" bson:"appParticipant,omitempty"`
	AdmissionResults           ZfcsAdmissionResults        `xml:"admissionResults,omitempty" bson:"admissionResults,omitempty"`
	AppRating                  int64                       `xml:"appRating,omitempty" bson:"appRating,omitempty"`
	WinnerPrice                string                      `xml:"winnerPrice,omitempty" bson:"winnerPrice,omitempty"`
	IncreaseWinnerInitialPrice bool                        `xml:"increaseWinnerInitialPrice,omitempty" bson:"increaseWinnerInitialPrice,omitempty"`
	AppRejectedReason          []ZfcsAppRejectedReasonType `xml:"appRejectedReason,omitempty" bson:"appRejectedReason,omitempty"`
	Admitted                   bool                        `xml:"admitted,omitempty" bson:"admitted,omitempty"`
	NotConsidered              bool                        `xml:"notConsidered,omitempty" bson:"notConsidered,omitempty"`
}

// ZfcsCommissionType
type ZfcsCommissionType struct {
	CommissionName    string                              `xml:"commissionName,omitempty" bson:"commissionName,omitempty"`
	CommissionMembers ZfcsCommissionTypeCommissionMembers `xml:"commissionMembers,omitempty" bson:"commissionMembers,omitempty"`
	Competent         bool                                `xml:"competent,omitempty" bson:"competent,omitempty"`
	AddInfo           string                              `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsCommissionTypeCommissionMembers Участники комиссии
type ZfcsCommissionTypeCommissionMembers struct {
	CommissionMember          []CommissionMembersCommissionMember `xml:"commissionMember,omitempty" bson:"commissionMember,omitempty"`
	SpelledMembersNoVoteCount string                              `xml:"spelledMembersNoVoteCount,omitempty" bson:"spelledMembersNoVoteCount,omitempty"`
	SpelledMembersCount       string                              `xml:"spelledMembersCount,omitempty" bson:"spelledMembersCount,omitempty"`
}

// CommissionMembersCommissionMember Участник комиссии
type CommissionMembersCommissionMember struct {
	MemberNumber uint64                 `xml:"memberNumber,omitempty" bson:"memberNumber,omitempty"`
	LastName     string                 `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName    string                 `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName   string                 `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	Role         ZfcsCommissionRoleType `xml:"role,omitempty" bson:"role,omitempty"`
}

// ZfcsCommissionRoleType
type ZfcsCommissionRoleType struct {
	ID        int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	RightVote bool   `xml:"rightVote,omitempty" bson:"rightVote,omitempty"`
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

// CryptoSignsSignature Электронная подпись
type CryptoSignsSignature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsProtocolModificationType
type ZfcsProtocolModificationType struct {
	ModificationNumber int64                              `xml:"modificationNumber,omitempty" bson:"modificationNumber,omitempty"`
	Info               string                             `xml:"info,omitempty" bson:"info,omitempty"`
	AddInfo            string                             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Reason             ZfcsProtocolModificationReasonType `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsProtocolModificationReasonType
type ZfcsProtocolModificationReasonType struct {
	ResponsibleDecision   ZfcsProtocolModificationReasonTypeResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	CourtDecision         ZfcsProtocolModificationReasonTypeCourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	AuthorityPrescription ZfcsProtocolModificationReasonTypeAuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
}

// ZfcsProtocolModificationReasonTypeResponsibleDecision По решению заказчика (организации, осуществляющей определение поставщика для заказчика)
type ZfcsProtocolModificationReasonTypeResponsibleDecision struct {
	DecisionDate xsd.Date `xml:"decisionDate,omitempty" bson:"decisionDate,omitempty"`
}

// ZfcsProtocolModificationReasonTypeCourtDecision Решение судебного органа
type ZfcsProtocolModificationReasonTypeCourtDecision struct {
	CourtName string   `xml:"courtName,omitempty" bson:"courtName,omitempty"`
	DocName   string   `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate   xsd.Date `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// ZfcsProtocolModificationReasonTypeAuthorityPrescription Предписание органа, уполномоченного на осуществление контроля
type ZfcsProtocolModificationReasonTypeAuthorityPrescription struct {
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

// ZfcsParticipantType
type ZfcsParticipantType struct {
	ParticipantType   string                `xml:"participantType,omitempty" bson:"participantType,omitempty"`
	Inn               string                `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp               string                `xml:"kpp,omitempty" bson:"kpp,omitempty"`
	Ogrn              string                `xml:"ogrn,omitempty" bson:"ogrn,omitempty"`
	LegalForm         ZfcsOkopfRef          `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	IDNumber          string                `xml:"idNumber,omitempty" bson:"idNumber,omitempty"`
	IDNumberExtension string                `xml:"idNumberExtension,omitempty" bson:"idNumberExtension,omitempty"`
	OrganizationName  string                `xml:"organizationName,omitempty" bson:"organizationName,omitempty"`
	FirmName          string                `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Country           ZfcsCountryRef        `xml:"country,omitempty" bson:"country,omitempty"`
	FactualAddress    string                `xml:"factualAddress,omitempty" bson:"factualAddress,omitempty"`
	PostAddress       string                `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	ContactInfo       ZfcsContactPersonType `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	ContactEMail      string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone      string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax        string                `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AdditionalInfo    string                `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	Status            string                `xml:"status,omitempty" bson:"status,omitempty"`
}

// ZfcsOkopfRef
type ZfcsOkopfRef struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ZfcsCountryRef
type ZfcsCountryRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsContactPersonType
type ZfcsContactPersonType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
}

// ZfcsAdmissionResults
type ZfcsAdmissionResults struct {
	AdmissionResult []ZfcsAdmissionResultsAdmissionResult `xml:"admissionResult,omitempty" bson:"admissionResult,omitempty"`
}

// ZfcsAdmissionResultsAdmissionResult
type ZfcsAdmissionResultsAdmissionResult struct {
	ProtocolCommissionMember ZfcsCommissionMemberInAppType    `xml:"protocolCommissionMember,omitempty" bson:"protocolCommissionMember,omitempty"`
	Admitted                 bool                             `xml:"admitted,omitempty" bson:"admitted,omitempty"`
	AppRejectedReason        []ZfcsAppRejectedReasonNotIDType `xml:"appRejectedReason,omitempty" bson:"appRejectedReason,omitempty"`
}

// ZfcsCommissionMemberInAppType
type ZfcsCommissionMemberInAppType struct {
	MemberNumber uint64 `xml:"memberNumber,omitempty" bson:"memberNumber,omitempty"`
}

// ZfcsAppRejectedReasonNotIDType
type ZfcsAppRejectedReasonNotIDType struct {
	NsiRejectReason ZfcsAppRejectedReasonNotIDTypeNsiRejectReason `xml:"nsiRejectReason,omitempty" bson:"nsiRejectReason,omitempty"`
	Explanation     string                                        `xml:"explanation,omitempty" bson:"explanation,omitempty"`
}

// ZfcsAppRejectedReasonNotIDTypeNsiRejectReason Причина отказа из справочника "Справочник причин отказа в допуске заявок" (nsiPurchaseRejectReason)
type ZfcsAppRejectedReasonNotIDTypeNsiRejectReason struct {
	Reason string `xml:"reason,omitempty" bson:"reason,omitempty"`
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
}

// ZfcsAppRejectedReasonType
type ZfcsAppRejectedReasonType struct {
	NsiRejectReason ZfcsAppRejectedReasonTypeNsiRejectReason `xml:"nsiRejectReason,omitempty" bson:"nsiRejectReason,omitempty"`
	Explanation     string                                   `xml:"explanation,omitempty" bson:"explanation,omitempty"`
}

// ZfcsAppRejectedReasonTypeNsiRejectReason Причина отказа
type ZfcsAppRejectedReasonTypeNsiRejectReason struct {
	Reason string `xml:"reason,omitempty" bson:"reason,omitempty"`
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
}

// ZfcsAbandonedReasonType
type ZfcsAbandonedReasonType struct {
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	ObjectName string `xml:"objectName,omitempty" bson:"objectName,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Type       string `xml:"type,omitempty" bson:"type,omitempty"`
}
