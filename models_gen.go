// Code generated from the AveeCare OpenAPI specification. DO NOT EDIT.

package aveecare

import (
	"net/url"
	"strconv"
)

// PatientLifecycleStage. Whether this person is a prospective client or an admitted client. `null`
// means client.
type PatientLifecycleStage string

const (
	PatientLifecycleStageProspect PatientLifecycleStage = "Prospect"
	PatientLifecycleStageClient   PatientLifecycleStage = "Client"
)

// PatientServiceType. The main type of service the patient receives.
type PatientServiceType string

const (
	PatientServiceTypeSkilledNursing  PatientServiceType = "SkilledNursing"
	PatientServiceTypePhysicalTherapy PatientServiceType = "PhysicalTherapy"
	PatientServiceTypeHospice         PatientServiceType = "Hospice"
	PatientServiceTypeStandardCare    PatientServiceType = "StandardCare"
)

// PatientStatementDelivery. How statements are delivered.
type PatientStatementDelivery string

const (
	PatientStatementDeliveryEmail        PatientStatementDelivery = "Email"
	PatientStatementDeliveryMail         PatientStatementDelivery = "Mail"
	PatientStatementDeliveryEmailAndMail PatientStatementDelivery = "EmailAndMail"
	PatientStatementDeliveryNoStatement  PatientStatementDelivery = "NoStatement"
)

// PatientPaymentMethodPreference. How the patient usually pays.
type PatientPaymentMethodPreference string

const (
	PatientPaymentMethodPreferenceACH       PatientPaymentMethodPreference = "ACH"
	PatientPaymentMethodPreferenceCard      PatientPaymentMethodPreference = "Card"
	PatientPaymentMethodPreferenceCheck     PatientPaymentMethodPreference = "Check"
	PatientPaymentMethodPreferenceCash      PatientPaymentMethodPreference = "Cash"
	PatientPaymentMethodPreferencePayerOnly PatientPaymentMethodPreference = "PayerOnly"
	PatientPaymentMethodPreferenceOther     PatientPaymentMethodPreference = "Other"
)

// PatientCreateStatus. The client's status as shown in AveeCare. `Inactive` and `Discharged` mean
// the person is no longer receiving care. A prospective client (`lifecycleStage: Prospect`) reads
// `Prospect` even if you send `Active`.
type PatientCreateStatus string

const (
	PatientCreateStatusActive     PatientCreateStatus = "Active"
	PatientCreateStatusInactive   PatientCreateStatus = "Inactive"
	PatientCreateStatusProspect   PatientCreateStatus = "Prospect"
	PatientCreateStatusDischarged PatientCreateStatus = "Discharged"
)

// ContactNotificationPolicy. Which updates the contact is notified about.
type ContactNotificationPolicy string

const (
	ContactNotificationPolicyAll         ContactNotificationPolicy = "All"
	ContactNotificationPolicyEmergencies ContactNotificationPolicy = "Emergencies"
)

// MedicationScheduleType. `Scheduled` for regular doses, `PRN` for as-needed.
type MedicationScheduleType string

const (
	MedicationScheduleTypeScheduled MedicationScheduleType = "Scheduled"
	MedicationScheduleTypePRN       MedicationScheduleType = "PRN"
)

// MedicationStatus. Whether the medication is currently given.
type MedicationStatus string

const (
	MedicationStatusActive       MedicationStatus = "Active"
	MedicationStatusOnHold       MedicationStatus = "OnHold"
	MedicationStatusDiscontinued MedicationStatus = "Discontinued"
)

// MedicationAdministrationStatus. What happened to the dose.
type MedicationAdministrationStatus string

const (
	MedicationAdministrationStatusGiven            MedicationAdministrationStatus = "Given"
	MedicationAdministrationStatusRefused          MedicationAdministrationStatus = "Refused"
	MedicationAdministrationStatusHeld             MedicationAdministrationStatus = "Held"
	MedicationAdministrationStatusMissed           MedicationAdministrationStatus = "Missed"
	MedicationAdministrationStatusSelfAdministered MedicationAdministrationStatus = "SelfAdministered"
)

// OasisAssessmentTimepoint. Which OASIS time point this assessment covers.
type OasisAssessmentTimepoint string

const (
	OasisAssessmentTimepointSOC             OasisAssessmentTimepoint = "SOC"
	OasisAssessmentTimepointROC             OasisAssessmentTimepoint = "ROC"
	OasisAssessmentTimepointRecertification OasisAssessmentTimepoint = "Recertification"
	OasisAssessmentTimepointTransfer        OasisAssessmentTimepoint = "Transfer"
	OasisAssessmentTimepointDischarge       OasisAssessmentTimepoint = "Discharge"
	OasisAssessmentTimepointDeathAtHome     OasisAssessmentTimepoint = "DeathAtHome"
)

// OasisAssessmentStatus. Where the assessment is in its lifecycle.
type OasisAssessmentStatus string

const (
	OasisAssessmentStatusDraft       OasisAssessmentStatus = "Draft"
	OasisAssessmentStatusCompleted   OasisAssessmentStatus = "Completed"
	OasisAssessmentStatusExported    OasisAssessmentStatus = "Exported"
	OasisAssessmentStatusAccepted    OasisAssessmentStatus = "Accepted"
	OasisAssessmentStatusRejected    OasisAssessmentStatus = "Rejected"
	OasisAssessmentStatusInactivated OasisAssessmentStatus = "Inactivated"
)

// CaregiverStatus. Employment status.
type CaregiverStatus string

const (
	CaregiverStatusActive   CaregiverStatus = "Active"
	CaregiverStatusOnLeave  CaregiverStatus = "OnLeave"
	CaregiverStatusInactive CaregiverStatus = "Inactive"
)

// VisitStatus. Where the visit is in its day. New visits are `Scheduled`.
type VisitStatus string

const (
	VisitStatusScheduled  VisitStatus = "Scheduled"
	VisitStatusInProgress VisitStatus = "InProgress"
	VisitStatusCancelled  VisitStatus = "Cancelled"
	VisitStatusCompleted  VisitStatus = "Completed"
	VisitStatusLate       VisitStatus = "Late"
	VisitStatusOverTime   VisitStatus = "OverTime"
	VisitStatusIncomplete VisitStatus = "Incomplete"
)

// VisitEvvStatus. Electronic visit verification status.
type VisitEvvStatus string

const (
	VisitEvvStatusPending    VisitEvvStatus = "Pending"
	VisitEvvStatusValidated  VisitEvvStatus = "Validated"
	VisitEvvStatusFlagged    VisitEvvStatus = "Flagged"
	VisitEvvStatusOverridden VisitEvvStatus = "Overridden"
	VisitEvvStatusExported   VisitEvvStatus = "Exported"
	VisitEvvStatusSubmitted  VisitEvvStatus = "Submitted"
)

// VisitBillingStatus. Whether the visit has been billed.
type VisitBillingStatus string

const (
	VisitBillingStatusNotBilled VisitBillingStatus = "NotBilled"
	VisitBillingStatusPending   VisitBillingStatus = "Pending"
	VisitBillingStatusInvoiced  VisitBillingStatus = "Invoiced"
	VisitBillingStatusPaid      VisitBillingStatus = "Paid"
)

// VisitCaregiverRsvp. The caregiver's response to the assignment.
type VisitCaregiverRsvp string

const (
	VisitCaregiverRsvpConfirmed VisitCaregiverRsvp = "confirmed"
	VisitCaregiverRsvpDeclined  VisitCaregiverRsvp = "declined"
)

// VisitNoteAuthorType. Whether office staff (`User`) or a caregiver wrote the note. Set from the
// API key's user.
type VisitNoteAuthorType string

const (
	VisitNoteAuthorTypeUser      VisitNoteAuthorType = "User"
	VisitNoteAuthorTypeCaregiver VisitNoteAuthorType = "Caregiver"
)

// VisitRequestStatus. Where the request stands. A caregiver's key creates `Pending` requests and
// can only change its own pending request to `Cancelled`; approving and denying is for office
// staff.
type VisitRequestStatus string

const (
	VisitRequestStatusPending   VisitRequestStatus = "Pending"
	VisitRequestStatusApproved  VisitRequestStatus = "Approved"
	VisitRequestStatusDenied    VisitRequestStatus = "Denied"
	VisitRequestStatusCancelled VisitRequestStatus = "Cancelled"
)

// OpenShiftStatus. Whether the offer is still open.
type OpenShiftStatus string

const (
	OpenShiftStatusActive    OpenShiftStatus = "Active"
	OpenShiftStatusFilled    OpenShiftStatus = "Filled"
	OpenShiftStatusCancelled OpenShiftStatus = "Cancelled"
)

// ShiftSwapStatus. Where the swap stands. With a caregiver's key: new swaps are `Pending`; the
// asked caregiver changes a `Pending` swap to `Accepted` or `Declined`; the requesting caregiver
// can change it to `Cancelled`. `AdminApproved` and `AdminDenied` are for office staff.
type ShiftSwapStatus string

const (
	ShiftSwapStatusPending       ShiftSwapStatus = "Pending"
	ShiftSwapStatusAccepted      ShiftSwapStatus = "Accepted"
	ShiftSwapStatusDeclined      ShiftSwapStatus = "Declined"
	ShiftSwapStatusCancelled     ShiftSwapStatus = "Cancelled"
	ShiftSwapStatusAdminApproved ShiftSwapStatus = "AdminApproved"
	ShiftSwapStatusAdminDenied   ShiftSwapStatus = "AdminDenied"
)

// CarePlanStatus. Where the plan is in its lifecycle.
type CarePlanStatus string

const (
	CarePlanStatusActive       CarePlanStatus = "Active"
	CarePlanStatusDraft        CarePlanStatus = "Draft"
	CarePlanStatusArchived     CarePlanStatus = "Archived"
	CarePlanStatusUnderReview  CarePlanStatus = "UnderReview"
	CarePlanStatusCompleted    CarePlanStatus = "Completed"
	CarePlanStatusDiscontinued CarePlanStatus = "Discontinued"
)

// AuthorizationUnitBasis. Whether units count time, visits or days.
type AuthorizationUnitBasis string

const (
	AuthorizationUnitBasisTime   AuthorizationUnitBasis = "Time"
	AuthorizationUnitBasisVisits AuthorizationUnitBasis = "Visits"
	AuthorizationUnitBasisDays   AuthorizationUnitBasis = "Days"
)

// AuthorizationStatus. Whether the authorization is in effect.
type AuthorizationStatus string

const (
	AuthorizationStatusActive         AuthorizationStatus = "Active"
	AuthorizationStatusExpired        AuthorizationStatus = "Expired"
	AuthorizationStatusPendingRenewal AuthorizationStatus = "PendingRenewal"
)

// IncidentCategory. Kind of incident.
type IncidentCategory string

const (
	IncidentCategoryFall       IncidentCategory = "Fall"
	IncidentCategoryMedication IncidentCategory = "Medication"
	IncidentCategoryBehavioral IncidentCategory = "Behavioral"
	IncidentCategoryInjury     IncidentCategory = "Injury"
	IncidentCategoryEquipment  IncidentCategory = "Equipment"
	IncidentCategoryAttendance IncidentCategory = "Attendance"
	IncidentCategoryConduct    IncidentCategory = "Conduct"
	IncidentCategoryOther      IncidentCategory = "Other"
)

// IncidentSeverity. How serious it was.
type IncidentSeverity string

const (
	IncidentSeverityLow      IncidentSeverity = "Low"
	IncidentSeverityMedium   IncidentSeverity = "Medium"
	IncidentSeverityHigh     IncidentSeverity = "High"
	IncidentSeverityCritical IncidentSeverity = "Critical"
)

// IncidentStatus. Where follow-up stands. New incidents start `Open`.
type IncidentStatus string

const (
	IncidentStatusOpen             IncidentStatus = "Open"
	IncidentStatusMonitoring       IncidentStatus = "Monitoring"
	IncidentStatusAwaitingFollowUp IncidentStatus = "AwaitingFollowUp"
	IncidentStatusResolved         IncidentStatus = "Resolved"
	IncidentStatusClosed           IncidentStatus = "Closed"
)

// ReferralStatus. Where the referral stands.
type ReferralStatus string

const (
	ReferralStatusNew                  ReferralStatus = "New"
	ReferralStatusContacted            ReferralStatus = "Contacted"
	ReferralStatusAssessmentScheduled  ReferralStatus = "AssessmentScheduled"
	ReferralStatusAssessmentCompleted  ReferralStatus = "AssessmentCompleted"
	ReferralStatusServiceAgreementSent ReferralStatus = "ServiceAgreementSent"
	ReferralStatusAdmitted             ReferralStatus = "Admitted"
	ReferralStatusDeclined             ReferralStatus = "Declined"
	ReferralStatusLost                 ReferralStatus = "Lost"
)

// InvoiceStatus. Payment status.
type InvoiceStatus string

const (
	InvoiceStatusDraft         InvoiceStatus = "Draft"
	InvoiceStatusPending       InvoiceStatus = "Pending"
	InvoiceStatusProcessing    InvoiceStatus = "Processing"
	InvoiceStatusPaid          InvoiceStatus = "Paid"
	InvoiceStatusPartiallyPaid InvoiceStatus = "PartiallyPaid"
	InvoiceStatusOverdue       InvoiceStatus = "Overdue"
	InvoiceStatusCancelled     InvoiceStatus = "Cancelled"
	InvoiceStatusRefunded      InvoiceStatus = "Refunded"
	InvoiceStatusFailed        InvoiceStatus = "Failed"
)

// ClaimClaimType. Claim format: professional (837P) or institutional (837I).
type ClaimClaimType string

const (
	ClaimClaimTypeProfessional837P  ClaimClaimType = "Professional837P"
	ClaimClaimTypeInstitutional837I ClaimClaimType = "Institutional837I"
)

// ClaimStatus. Where the claim is in adjudication.
type ClaimStatus string

const (
	ClaimStatusDraft            ClaimStatus = "Draft"
	ClaimStatusValidated        ClaimStatus = "Validated"
	ClaimStatusValidationFailed ClaimStatus = "ValidationFailed"
	ClaimStatusSubmitted        ClaimStatus = "Submitted"
	ClaimStatusAcknowledged     ClaimStatus = "Acknowledged"
	ClaimStatusAccepted         ClaimStatus = "Accepted"
	ClaimStatusRejected         ClaimStatus = "Rejected"
	ClaimStatusInAdjudication   ClaimStatus = "InAdjudication"
	ClaimStatusPaid             ClaimStatus = "Paid"
	ClaimStatusPartiallyPaid    ClaimStatus = "PartiallyPaid"
	ClaimStatusDenied           ClaimStatus = "Denied"
	ClaimStatusAppealed         ClaimStatus = "Appealed"
	ClaimStatusVoided           ClaimStatus = "Voided"
)

// PayrollEntryStatus. Where the entry is in your payroll process.
type PayrollEntryStatus string

const (
	PayrollEntryStatusPending        PayrollEntryStatus = "Pending"
	PayrollEntryStatusApproved       PayrollEntryStatus = "Approved"
	PayrollEntryStatusProcessed      PayrollEntryStatus = "Processed"
	PayrollEntryStatusPaid           PayrollEntryStatus = "Paid"
	PayrollEntryStatusCancelled      PayrollEntryStatus = "Cancelled"
	PayrollEntryStatusFailed         PayrollEntryStatus = "Failed"
	PayrollEntryStatusHoursSubmitted PayrollEntryStatus = "HoursSubmitted"
	PayrollEntryStatusCheckCreated   PayrollEntryStatus = "CheckCreated"
	PayrollEntryStatusPaidConfirmed  PayrollEntryStatus = "PaidConfirmed"
	PayrollEntryStatusPaidExternally PayrollEntryStatus = "PaidExternally"
	PayrollEntryStatusVoided         PayrollEntryStatus = "Voided"
)

// PayrollEntryPaymentMethod. How the caregiver is paid.
type PayrollEntryPaymentMethod string

const (
	PayrollEntryPaymentMethodDirectDeposit PayrollEntryPaymentMethod = "DirectDeposit"
	PayrollEntryPaymentMethodCheck         PayrollEntryPaymentMethod = "Check"
	PayrollEntryPaymentMethodManual        PayrollEntryPaymentMethod = "Manual"
	PayrollEntryPaymentMethodACH           PayrollEntryPaymentMethod = "ACH"
)

// AlertType. How prominently the alert is shown.
type AlertType string

const (
	AlertTypeCritical  AlertType = "Critical"
	AlertTypeImportant AlertType = "Important"
	AlertTypeNotice    AlertType = "Notice"
)

// AlertAudience. Who sees the alert. `Caregiver`, `Patient`, `Visit` and `PatientCareTeam` need
// `audienceTargetId`.
type AlertAudience string

const (
	AlertAudienceCompanyWide     AlertAudience = "CompanyWide"
	AlertAudienceStaffOnly       AlertAudience = "StaffOnly"
	AlertAudienceCaregiversAll   AlertAudience = "CaregiversAll"
	AlertAudiencePatientsAll     AlertAudience = "PatientsAll"
	AlertAudienceCaregiver       AlertAudience = "Caregiver"
	AlertAudiencePatient         AlertAudience = "Patient"
	AlertAudienceVisit           AlertAudience = "Visit"
	AlertAudiencePatientCareTeam AlertAudience = "PatientCareTeam"
)

// MeScope.
type MeScope string

const (
	MeScopeRead  MeScope = "read"
	MeScopeWrite MeScope = "write"
)

// MeRole. Role of the person the key belongs to: `RootUser` (agency administrator), `User` (office
// staff), `Caregiver` or `Patient`. It decides which resources the key can reach.
type MeRole string

const (
	MeRoleRootUser  MeRole = "RootUser"
	MeRoleUser      MeRole = "User"
	MeRoleCaregiver MeRole = "Caregiver"
	MeRolePatient   MeRole = "Patient"
)

// Patients are the clients your agency provides care to, including prospective clients you have
// not yet admitted. Every field is always present; a field with no value is `null`.
type Patient struct {
	// String representing the object's type. Always `patient`.
	Object string `json:"object"`
	// Unique identifier of the patient.
	ID string `json:"id"`
	// Given name.
	FirstName *string `json:"firstName"`
	// Middle name.
	MiddleName *string `json:"middleName"`
	// Family name.
	LastName *string `json:"lastName"`
	// The name the patient prefers to be called.
	GoesByName *string `json:"goesByName"`
	// Date of birth.
	DateOfBirth *string `json:"dateOfBirth"`
	// Gender as recorded by the agency.
	Gender *string `json:"gender"`
	// Language the patient is most comfortable speaking.
	PrimaryLanguage *string `json:"primaryLanguage"`
	// Contact email address.
	Email *string `json:"email"`
	// Whether the patient is an active client. Only active clients are scheduled and billed. A
	// prospective client is always inactive.
	Active *bool `json:"active"`
	// The client's status as shown in AveeCare. `Inactive` and `Discharged` mean the person is no
	// longer receiving care. A prospective client (`lifecycleStage: Prospect`) reads `Prospect` even
	// if you send `Active`. Usually one of `Active`, `Inactive`, `Prospect`, `Discharged`; records
	// imported from another system can hold other values.
	Status *string `json:"status"`
	// Whether this person is a prospective client or an admitted client. `null` means client.
	LifecycleStage *PatientLifecycleStage `json:"lifecycleStage"`
	// The main type of service the patient receives.
	ServiceType *PatientServiceType `json:"serviceType"`
	// Street address where care is delivered.
	Address *string `json:"address"`
	// City of the service address.
	City *string `json:"city"`
	// State of the service address, as a two-letter code.
	State *string `json:"state"`
	// ZIP code of the service address.
	ZipCode *string `json:"zipCode"`
	// The office (branch) that serves this patient.
	OfficeID *string `json:"officeId"`
	// The referral source that sent this patient to you.
	ReferralSourceID *string `json:"referralSourceId"`
	// Primary diagnosis, in plain words.
	PrimaryDiagnosisDescription *string `json:"primaryDiagnosisDescription"`
	// Whether the patient is a fall risk.
	FallRisk *bool `json:"fallRisk"`
	// Caregiver gender the patient prefers, if any.
	PreferredCaregiverGender *string `json:"preferredCaregiverGender"`
	// Notes schedulers see when booking visits for this patient.
	SchedulingNotes *string `json:"schedulingNotes"`
	// Person legally responsible for the patient, if not the patient.
	ResponsiblePartyName *string `json:"responsiblePartyName"`
	// Medicaid member ID.
	MedicaidID *string `json:"medicaidId"`
	// Medicare beneficiary identifier.
	MedicareID *string `json:"medicareId"`
	// Primary insurance company name.
	InsuranceProvider *string `json:"insuranceProvider"`
	// Member ID with the primary insurance.
	InsuranceID *string `json:"insuranceId"`
	// Group number with the primary insurance.
	InsuranceGroupNumber *string `json:"insuranceGroupNumber"`
	// Who receives the bills.
	BillingContactName *string `json:"billingContactName"`
	// Email address bills are sent to.
	BillingContactEmail *string `json:"billingContactEmail"`
	// Phone number of the billing contact.
	BillingContactPhone *string `json:"billingContactPhone"`
	// How statements are delivered.
	StatementDelivery *PatientStatementDelivery `json:"statementDelivery"`
	// How the patient usually pays.
	PaymentMethodPreference *PatientPaymentMethodPreference `json:"paymentMethodPreference"`
	// Balance the patient currently owes. In cents.
	OutstandingBalanceCents *int64 `json:"outstandingBalanceCents"`
	// When the patient was created.
	CreatedAt *string `json:"createdAt"`
	// When the patient was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new patient. Any other field is refused with `unknown_field`.
type PatientCreate struct {
	// Given name.
	FirstName string `json:"firstName"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName string `json:"lastName"`
	// The name the patient prefers to be called.
	GoesByName Nullable[string] `json:"goesByName,omitempty"`
	// Date of birth.
	DateOfBirth Nullable[string] `json:"dateOfBirth,omitempty"`
	// Gender as recorded by the agency.
	Gender Nullable[string] `json:"gender,omitempty"`
	// Language the patient is most comfortable speaking.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Contact email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Whether the patient is an active client. Only active clients are scheduled and billed. A
	// prospective client is always inactive.
	Active Nullable[bool] `json:"active,omitempty"`
	// The client's status as shown in AveeCare. `Inactive` and `Discharged` mean the person is no
	// longer receiving care. A prospective client (`lifecycleStage: Prospect`) reads `Prospect` even
	// if you send `Active`.
	Status Nullable[PatientCreateStatus] `json:"status,omitempty"`
	// Whether this person is a prospective client or an admitted client. `null` means client.
	LifecycleStage Nullable[PatientLifecycleStage] `json:"lifecycleStage,omitempty"`
	// The main type of service the patient receives.
	ServiceType Nullable[PatientServiceType] `json:"serviceType,omitempty"`
	// Street address where care is delivered.
	Address Nullable[string] `json:"address,omitempty"`
	// City of the service address.
	City Nullable[string] `json:"city,omitempty"`
	// State of the service address, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code of the service address.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// The office (branch) that serves this patient.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// The referral source that sent this patient to you.
	ReferralSourceID Nullable[string] `json:"referralSourceId,omitempty"`
	// Primary diagnosis, in plain words.
	PrimaryDiagnosisDescription Nullable[string] `json:"primaryDiagnosisDescription,omitempty"`
	// Whether the patient is a fall risk.
	FallRisk Nullable[bool] `json:"fallRisk,omitempty"`
	// Caregiver gender the patient prefers, if any.
	PreferredCaregiverGender Nullable[string] `json:"preferredCaregiverGender,omitempty"`
	// Notes schedulers see when booking visits for this patient.
	SchedulingNotes Nullable[string] `json:"schedulingNotes,omitempty"`
	// Person legally responsible for the patient, if not the patient.
	ResponsiblePartyName Nullable[string] `json:"responsiblePartyName,omitempty"`
	// Medicaid member ID.
	MedicaidID Nullable[string] `json:"medicaidId,omitempty"`
	// Medicare beneficiary identifier.
	MedicareID Nullable[string] `json:"medicareId,omitempty"`
	// Primary insurance company name.
	InsuranceProvider Nullable[string] `json:"insuranceProvider,omitempty"`
	// Member ID with the primary insurance.
	InsuranceID Nullable[string] `json:"insuranceId,omitempty"`
	// Group number with the primary insurance.
	InsuranceGroupNumber Nullable[string] `json:"insuranceGroupNumber,omitempty"`
	// Who receives the bills.
	BillingContactName Nullable[string] `json:"billingContactName,omitempty"`
	// Email address bills are sent to.
	BillingContactEmail Nullable[string] `json:"billingContactEmail,omitempty"`
	// Phone number of the billing contact.
	BillingContactPhone Nullable[string] `json:"billingContactPhone,omitempty"`
	// How statements are delivered.
	StatementDelivery Nullable[PatientStatementDelivery] `json:"statementDelivery,omitempty"`
	// How the patient usually pays.
	PaymentMethodPreference Nullable[PatientPaymentMethodPreference] `json:"paymentMethodPreference,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m PatientCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("firstName", m.FirstName)
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	w.add("lastName", m.LastName)
	if !m.GoesByName.IsOmitted() {
		w.add("goesByName", m.GoesByName)
	}
	if !m.DateOfBirth.IsOmitted() {
		w.add("dateOfBirth", m.DateOfBirth)
	}
	if !m.Gender.IsOmitted() {
		w.add("gender", m.Gender)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Active.IsOmitted() {
		w.add("active", m.Active)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.LifecycleStage.IsOmitted() {
		w.add("lifecycleStage", m.LifecycleStage)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.ReferralSourceID.IsOmitted() {
		w.add("referralSourceId", m.ReferralSourceID)
	}
	if !m.PrimaryDiagnosisDescription.IsOmitted() {
		w.add("primaryDiagnosisDescription", m.PrimaryDiagnosisDescription)
	}
	if !m.FallRisk.IsOmitted() {
		w.add("fallRisk", m.FallRisk)
	}
	if !m.PreferredCaregiverGender.IsOmitted() {
		w.add("preferredCaregiverGender", m.PreferredCaregiverGender)
	}
	if !m.SchedulingNotes.IsOmitted() {
		w.add("schedulingNotes", m.SchedulingNotes)
	}
	if !m.ResponsiblePartyName.IsOmitted() {
		w.add("responsiblePartyName", m.ResponsiblePartyName)
	}
	if !m.MedicaidID.IsOmitted() {
		w.add("medicaidId", m.MedicaidID)
	}
	if !m.MedicareID.IsOmitted() {
		w.add("medicareId", m.MedicareID)
	}
	if !m.InsuranceProvider.IsOmitted() {
		w.add("insuranceProvider", m.InsuranceProvider)
	}
	if !m.InsuranceID.IsOmitted() {
		w.add("insuranceId", m.InsuranceID)
	}
	if !m.InsuranceGroupNumber.IsOmitted() {
		w.add("insuranceGroupNumber", m.InsuranceGroupNumber)
	}
	if !m.BillingContactName.IsOmitted() {
		w.add("billingContactName", m.BillingContactName)
	}
	if !m.BillingContactEmail.IsOmitted() {
		w.add("billingContactEmail", m.BillingContactEmail)
	}
	if !m.BillingContactPhone.IsOmitted() {
		w.add("billingContactPhone", m.BillingContactPhone)
	}
	if !m.StatementDelivery.IsOmitted() {
		w.add("statementDelivery", m.StatementDelivery)
	}
	if !m.PaymentMethodPreference.IsOmitted() {
		w.add("paymentMethodPreference", m.PaymentMethodPreference)
	}
	return w.bytes()
}

// Fields to change on a patient. Send only the fields to change; any other field is refused with
// `unknown_field`.
type PatientUpdate struct {
	// Given name.
	FirstName *string `json:"firstName,omitempty"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName *string `json:"lastName,omitempty"`
	// The name the patient prefers to be called.
	GoesByName Nullable[string] `json:"goesByName,omitempty"`
	// Date of birth.
	DateOfBirth Nullable[string] `json:"dateOfBirth,omitempty"`
	// Gender as recorded by the agency.
	Gender Nullable[string] `json:"gender,omitempty"`
	// Language the patient is most comfortable speaking.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Contact email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Whether the patient is an active client. Only active clients are scheduled and billed. A
	// prospective client is always inactive.
	Active Nullable[bool] `json:"active,omitempty"`
	// The client's status as shown in AveeCare. `Inactive` and `Discharged` mean the person is no
	// longer receiving care. A prospective client (`lifecycleStage: Prospect`) reads `Prospect` even
	// if you send `Active`.
	Status Nullable[PatientCreateStatus] `json:"status,omitempty"`
	// Whether this person is a prospective client or an admitted client. `null` means client.
	LifecycleStage Nullable[PatientLifecycleStage] `json:"lifecycleStage,omitempty"`
	// The main type of service the patient receives.
	ServiceType Nullable[PatientServiceType] `json:"serviceType,omitempty"`
	// Street address where care is delivered.
	Address Nullable[string] `json:"address,omitempty"`
	// City of the service address.
	City Nullable[string] `json:"city,omitempty"`
	// State of the service address, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code of the service address.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// The office (branch) that serves this patient.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// The referral source that sent this patient to you.
	ReferralSourceID Nullable[string] `json:"referralSourceId,omitempty"`
	// Primary diagnosis, in plain words.
	PrimaryDiagnosisDescription Nullable[string] `json:"primaryDiagnosisDescription,omitempty"`
	// Whether the patient is a fall risk.
	FallRisk Nullable[bool] `json:"fallRisk,omitempty"`
	// Caregiver gender the patient prefers, if any.
	PreferredCaregiverGender Nullable[string] `json:"preferredCaregiverGender,omitempty"`
	// Notes schedulers see when booking visits for this patient.
	SchedulingNotes Nullable[string] `json:"schedulingNotes,omitempty"`
	// Person legally responsible for the patient, if not the patient.
	ResponsiblePartyName Nullable[string] `json:"responsiblePartyName,omitempty"`
	// Medicaid member ID.
	MedicaidID Nullable[string] `json:"medicaidId,omitempty"`
	// Medicare beneficiary identifier.
	MedicareID Nullable[string] `json:"medicareId,omitempty"`
	// Primary insurance company name.
	InsuranceProvider Nullable[string] `json:"insuranceProvider,omitempty"`
	// Member ID with the primary insurance.
	InsuranceID Nullable[string] `json:"insuranceId,omitempty"`
	// Group number with the primary insurance.
	InsuranceGroupNumber Nullable[string] `json:"insuranceGroupNumber,omitempty"`
	// Who receives the bills.
	BillingContactName Nullable[string] `json:"billingContactName,omitempty"`
	// Email address bills are sent to.
	BillingContactEmail Nullable[string] `json:"billingContactEmail,omitempty"`
	// Phone number of the billing contact.
	BillingContactPhone Nullable[string] `json:"billingContactPhone,omitempty"`
	// How statements are delivered.
	StatementDelivery Nullable[PatientStatementDelivery] `json:"statementDelivery,omitempty"`
	// How the patient usually pays.
	PaymentMethodPreference Nullable[PatientPaymentMethodPreference] `json:"paymentMethodPreference,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m PatientUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.FirstName != nil {
		w.add("firstName", m.FirstName)
	}
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	if m.LastName != nil {
		w.add("lastName", m.LastName)
	}
	if !m.GoesByName.IsOmitted() {
		w.add("goesByName", m.GoesByName)
	}
	if !m.DateOfBirth.IsOmitted() {
		w.add("dateOfBirth", m.DateOfBirth)
	}
	if !m.Gender.IsOmitted() {
		w.add("gender", m.Gender)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Active.IsOmitted() {
		w.add("active", m.Active)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.LifecycleStage.IsOmitted() {
		w.add("lifecycleStage", m.LifecycleStage)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.ReferralSourceID.IsOmitted() {
		w.add("referralSourceId", m.ReferralSourceID)
	}
	if !m.PrimaryDiagnosisDescription.IsOmitted() {
		w.add("primaryDiagnosisDescription", m.PrimaryDiagnosisDescription)
	}
	if !m.FallRisk.IsOmitted() {
		w.add("fallRisk", m.FallRisk)
	}
	if !m.PreferredCaregiverGender.IsOmitted() {
		w.add("preferredCaregiverGender", m.PreferredCaregiverGender)
	}
	if !m.SchedulingNotes.IsOmitted() {
		w.add("schedulingNotes", m.SchedulingNotes)
	}
	if !m.ResponsiblePartyName.IsOmitted() {
		w.add("responsiblePartyName", m.ResponsiblePartyName)
	}
	if !m.MedicaidID.IsOmitted() {
		w.add("medicaidId", m.MedicaidID)
	}
	if !m.MedicareID.IsOmitted() {
		w.add("medicareId", m.MedicareID)
	}
	if !m.InsuranceProvider.IsOmitted() {
		w.add("insuranceProvider", m.InsuranceProvider)
	}
	if !m.InsuranceID.IsOmitted() {
		w.add("insuranceId", m.InsuranceID)
	}
	if !m.InsuranceGroupNumber.IsOmitted() {
		w.add("insuranceGroupNumber", m.InsuranceGroupNumber)
	}
	if !m.BillingContactName.IsOmitted() {
		w.add("billingContactName", m.BillingContactName)
	}
	if !m.BillingContactEmail.IsOmitted() {
		w.add("billingContactEmail", m.BillingContactEmail)
	}
	if !m.BillingContactPhone.IsOmitted() {
		w.add("billingContactPhone", m.BillingContactPhone)
	}
	if !m.StatementDelivery.IsOmitted() {
		w.add("statementDelivery", m.StatementDelivery)
	}
	if !m.PaymentMethodPreference.IsOmitted() {
		w.add("paymentMethodPreference", m.PaymentMethodPreference)
	}
	return w.bytes()
}

// Confirmation that a patient was deleted.
type PatientDeleted struct {
	// Always `patient`.
	Object string `json:"object"`
	// ID of the deleted patient.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Contacts are family members and other people connected to a patient, such as an emergency
// contact. Every field is always present; a field with no value is `null`.
type Contact struct {
	// String representing the object's type. Always `contact`.
	Object string `json:"object"`
	// Unique identifier of the contact.
	ID string `json:"id"`
	// The patient this contact belongs to.
	PatientID *string `json:"patientId"`
	// Full name.
	Name *string `json:"name"`
	// Relationship to the patient.
	Relationship *string `json:"relationship"`
	// Phone number.
	PhoneNumber *string `json:"phoneNumber"`
	// Email address.
	Email *string `json:"email"`
	// Whether the contact prefers email over phone.
	PreferEmail *bool `json:"preferEmail"`
	// Whether this person is an emergency contact.
	EmergencyContact *bool `json:"emergencyContact"`
	// Which updates the contact is notified about.
	NotificationPolicy *ContactNotificationPolicy `json:"notificationPolicy"`
	// When the contact was created.
	CreatedAt *string `json:"createdAt"`
	// When the contact was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new contact. Any other field is refused with `unknown_field`.
type ContactCreate struct {
	// The patient this contact belongs to.
	PatientID string `json:"patientId"`
	// Full name.
	Name string `json:"name"`
	// Relationship to the patient.
	Relationship Nullable[string] `json:"relationship,omitempty"`
	// Phone number.
	PhoneNumber Nullable[string] `json:"phoneNumber,omitempty"`
	// Email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Whether the contact prefers email over phone.
	PreferEmail Nullable[bool] `json:"preferEmail,omitempty"`
	// Whether this person is an emergency contact.
	EmergencyContact Nullable[bool] `json:"emergencyContact,omitempty"`
	// Which updates the contact is notified about.
	NotificationPolicy Nullable[ContactNotificationPolicy] `json:"notificationPolicy,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ContactCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("name", m.Name)
	if !m.Relationship.IsOmitted() {
		w.add("relationship", m.Relationship)
	}
	if !m.PhoneNumber.IsOmitted() {
		w.add("phoneNumber", m.PhoneNumber)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.PreferEmail.IsOmitted() {
		w.add("preferEmail", m.PreferEmail)
	}
	if !m.EmergencyContact.IsOmitted() {
		w.add("emergencyContact", m.EmergencyContact)
	}
	if !m.NotificationPolicy.IsOmitted() {
		w.add("notificationPolicy", m.NotificationPolicy)
	}
	return w.bytes()
}

// Fields to change on a contact. Send only the fields to change; any other field is refused with
// `unknown_field`. `patientId` is set when the contact is created and cannot be changed
// afterwards.
type ContactUpdate struct {
	// Full name.
	Name *string `json:"name,omitempty"`
	// Relationship to the patient.
	Relationship Nullable[string] `json:"relationship,omitempty"`
	// Phone number.
	PhoneNumber Nullable[string] `json:"phoneNumber,omitempty"`
	// Email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Whether the contact prefers email over phone.
	PreferEmail Nullable[bool] `json:"preferEmail,omitempty"`
	// Whether this person is an emergency contact.
	EmergencyContact Nullable[bool] `json:"emergencyContact,omitempty"`
	// Which updates the contact is notified about.
	NotificationPolicy Nullable[ContactNotificationPolicy] `json:"notificationPolicy,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ContactUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Relationship.IsOmitted() {
		w.add("relationship", m.Relationship)
	}
	if !m.PhoneNumber.IsOmitted() {
		w.add("phoneNumber", m.PhoneNumber)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.PreferEmail.IsOmitted() {
		w.add("preferEmail", m.PreferEmail)
	}
	if !m.EmergencyContact.IsOmitted() {
		w.add("emergencyContact", m.EmergencyContact)
	}
	if !m.NotificationPolicy.IsOmitted() {
		w.add("notificationPolicy", m.NotificationPolicy)
	}
	return w.bytes()
}

// Confirmation that a contact was deleted.
type ContactDeleted struct {
	// Always `contact`.
	Object string `json:"object"`
	// ID of the deleted contact.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Allergies recorded on a patient's profile. Every field is always present; a field with no value
// is `null`.
type Allergy struct {
	// String representing the object's type. Always `allergy`.
	Object string `json:"object"`
	// Unique identifier of the allergy.
	ID string `json:"id"`
	// The patient who has this allergy.
	PatientID *string `json:"patientId"`
	// What the patient is allergic to.
	Name *string `json:"name"`
	// Reaction and other details.
	Description *string `json:"description"`
	// When the allergy was created.
	CreatedAt *string `json:"createdAt"`
	// When the allergy was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new allergy. Any other field is refused with `unknown_field`.
type AllergyCreate struct {
	// The patient who has this allergy.
	PatientID string `json:"patientId"`
	// What the patient is allergic to.
	Name string `json:"name"`
	// Reaction and other details.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AllergyCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("name", m.Name)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Fields to change on an allergy. Send only the fields to change; any other field is refused with
// `unknown_field`. `patientId` is set when the allergy is created and cannot be changed
// afterwards.
type AllergyUpdate struct {
	// What the patient is allergic to.
	Name *string `json:"name,omitempty"`
	// Reaction and other details.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AllergyUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Confirmation that an allergy was deleted.
type AllergyDeleted struct {
	// Always `allergy`.
	Object string `json:"object"`
	// ID of the deleted allergy.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Medications on a patient's medication list, with the schedule caregivers follow when they give
// or assist with them. Every field is always present; a field with no value is `null`.
type Medication struct {
	// String representing the object's type. Always `medication`.
	Object string `json:"object"`
	// Unique identifier of the medication.
	ID string `json:"id"`
	// The patient who takes this medication.
	PatientID *string `json:"patientId"`
	// Drug name, without the strength.
	Name *string `json:"name"`
	// Strength per dose unit.
	Strength *string `json:"strength"`
	// Dose form.
	Form *string `json:"form"`
	// How it is taken.
	Route *string `json:"route"`
	// Directions as written on the prescription.
	Directions *string `json:"directions"`
	// What the medication is for.
	Indication *string `json:"indication"`
	// `Scheduled` for regular doses, `PRN` for as-needed.
	ScheduleType *MedicationScheduleType `json:"scheduleType"`
	// For as-needed medications, when to give it.
	PrnReason *string `json:"prnReason"`
	// For as-needed medications, the most doses allowed in a day.
	PrnMaxDailyDoses *int64 `json:"prnMaxDailyDoses"`
	// For as-needed medications, the minimum hours between doses.
	PrnMinIntervalHours *float64 `json:"prnMinIntervalHours"`
	// Date the patient started the medication.
	StartDate *string `json:"startDate"`
	// Whether the medication is currently given.
	Status *MedicationStatus `json:"status"`
	// Date the medication was stopped.
	DiscontinuedDate *string `json:"discontinuedDate"`
	// Why the medication was stopped.
	DiscontinuedReason *string `json:"discontinuedReason"`
	// Prescribing clinician.
	Prescriber *string `json:"prescriber"`
	// Prescriber's phone number.
	PrescriberPhone *string `json:"prescriberPhone"`
	// Pharmacy that fills the prescription.
	Pharmacy *string `json:"pharmacy"`
	// Pharmacy phone number.
	PharmacyPhone *string `json:"pharmacyPhone"`
	// Whether a second person must witness each dose.
	RequiresWitness *bool `json:"requiresWitness"`
	// Doses currently on hand.
	QuantityOnHand *float64 `json:"quantityOnHand"`
	// Refills left on the prescription.
	RefillsRemaining *int64 `json:"refillsRemaining"`
	// Other notes.
	Notes *string `json:"notes"`
	// When the medication was created.
	CreatedAt *string `json:"createdAt"`
	// When the medication was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new medication. Any other field is refused with `unknown_field`.
type MedicationCreate struct {
	// The patient who takes this medication.
	PatientID string `json:"patientId"`
	// Drug name, without the strength.
	Name string `json:"name"`
	// Strength per dose unit.
	Strength Nullable[string] `json:"strength,omitempty"`
	// Dose form.
	Form Nullable[string] `json:"form,omitempty"`
	// How it is taken.
	Route Nullable[string] `json:"route,omitempty"`
	// Directions as written on the prescription.
	Directions Nullable[string] `json:"directions,omitempty"`
	// What the medication is for.
	Indication Nullable[string] `json:"indication,omitempty"`
	// `Scheduled` for regular doses, `PRN` for as-needed.
	ScheduleType Nullable[MedicationScheduleType] `json:"scheduleType,omitempty"`
	// For as-needed medications, when to give it.
	PrnReason Nullable[string] `json:"prnReason,omitempty"`
	// For as-needed medications, the most doses allowed in a day.
	PrnMaxDailyDoses Nullable[int64] `json:"prnMaxDailyDoses,omitempty"`
	// For as-needed medications, the minimum hours between doses.
	PrnMinIntervalHours Nullable[float64] `json:"prnMinIntervalHours,omitempty"`
	// Date the patient started the medication.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Whether the medication is currently given.
	Status Nullable[MedicationStatus] `json:"status,omitempty"`
	// Date the medication was stopped.
	DiscontinuedDate Nullable[string] `json:"discontinuedDate,omitempty"`
	// Why the medication was stopped.
	DiscontinuedReason Nullable[string] `json:"discontinuedReason,omitempty"`
	// Prescribing clinician.
	Prescriber Nullable[string] `json:"prescriber,omitempty"`
	// Prescriber's phone number.
	PrescriberPhone Nullable[string] `json:"prescriberPhone,omitempty"`
	// Pharmacy that fills the prescription.
	Pharmacy Nullable[string] `json:"pharmacy,omitempty"`
	// Pharmacy phone number.
	PharmacyPhone Nullable[string] `json:"pharmacyPhone,omitempty"`
	// Whether a second person must witness each dose.
	RequiresWitness Nullable[bool] `json:"requiresWitness,omitempty"`
	// Doses currently on hand.
	QuantityOnHand Nullable[float64] `json:"quantityOnHand,omitempty"`
	// Refills left on the prescription.
	RefillsRemaining Nullable[int64] `json:"refillsRemaining,omitempty"`
	// Other notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m MedicationCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("name", m.Name)
	if !m.Strength.IsOmitted() {
		w.add("strength", m.Strength)
	}
	if !m.Form.IsOmitted() {
		w.add("form", m.Form)
	}
	if !m.Route.IsOmitted() {
		w.add("route", m.Route)
	}
	if !m.Directions.IsOmitted() {
		w.add("directions", m.Directions)
	}
	if !m.Indication.IsOmitted() {
		w.add("indication", m.Indication)
	}
	if !m.ScheduleType.IsOmitted() {
		w.add("scheduleType", m.ScheduleType)
	}
	if !m.PrnReason.IsOmitted() {
		w.add("prnReason", m.PrnReason)
	}
	if !m.PrnMaxDailyDoses.IsOmitted() {
		w.add("prnMaxDailyDoses", m.PrnMaxDailyDoses)
	}
	if !m.PrnMinIntervalHours.IsOmitted() {
		w.add("prnMinIntervalHours", m.PrnMinIntervalHours)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.DiscontinuedDate.IsOmitted() {
		w.add("discontinuedDate", m.DiscontinuedDate)
	}
	if !m.DiscontinuedReason.IsOmitted() {
		w.add("discontinuedReason", m.DiscontinuedReason)
	}
	if !m.Prescriber.IsOmitted() {
		w.add("prescriber", m.Prescriber)
	}
	if !m.PrescriberPhone.IsOmitted() {
		w.add("prescriberPhone", m.PrescriberPhone)
	}
	if !m.Pharmacy.IsOmitted() {
		w.add("pharmacy", m.Pharmacy)
	}
	if !m.PharmacyPhone.IsOmitted() {
		w.add("pharmacyPhone", m.PharmacyPhone)
	}
	if !m.RequiresWitness.IsOmitted() {
		w.add("requiresWitness", m.RequiresWitness)
	}
	if !m.QuantityOnHand.IsOmitted() {
		w.add("quantityOnHand", m.QuantityOnHand)
	}
	if !m.RefillsRemaining.IsOmitted() {
		w.add("refillsRemaining", m.RefillsRemaining)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Fields to change on a medication. Send only the fields to change; any other field is refused
// with `unknown_field`. `patientId` is set when the medication is created and cannot be changed
// afterwards.
type MedicationUpdate struct {
	// Drug name, without the strength.
	Name *string `json:"name,omitempty"`
	// Strength per dose unit.
	Strength Nullable[string] `json:"strength,omitempty"`
	// Dose form.
	Form Nullable[string] `json:"form,omitempty"`
	// How it is taken.
	Route Nullable[string] `json:"route,omitempty"`
	// Directions as written on the prescription.
	Directions Nullable[string] `json:"directions,omitempty"`
	// What the medication is for.
	Indication Nullable[string] `json:"indication,omitempty"`
	// `Scheduled` for regular doses, `PRN` for as-needed.
	ScheduleType Nullable[MedicationScheduleType] `json:"scheduleType,omitempty"`
	// For as-needed medications, when to give it.
	PrnReason Nullable[string] `json:"prnReason,omitempty"`
	// For as-needed medications, the most doses allowed in a day.
	PrnMaxDailyDoses Nullable[int64] `json:"prnMaxDailyDoses,omitempty"`
	// For as-needed medications, the minimum hours between doses.
	PrnMinIntervalHours Nullable[float64] `json:"prnMinIntervalHours,omitempty"`
	// Date the patient started the medication.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Whether the medication is currently given.
	Status Nullable[MedicationStatus] `json:"status,omitempty"`
	// Date the medication was stopped.
	DiscontinuedDate Nullable[string] `json:"discontinuedDate,omitempty"`
	// Why the medication was stopped.
	DiscontinuedReason Nullable[string] `json:"discontinuedReason,omitempty"`
	// Prescribing clinician.
	Prescriber Nullable[string] `json:"prescriber,omitempty"`
	// Prescriber's phone number.
	PrescriberPhone Nullable[string] `json:"prescriberPhone,omitempty"`
	// Pharmacy that fills the prescription.
	Pharmacy Nullable[string] `json:"pharmacy,omitempty"`
	// Pharmacy phone number.
	PharmacyPhone Nullable[string] `json:"pharmacyPhone,omitempty"`
	// Whether a second person must witness each dose.
	RequiresWitness Nullable[bool] `json:"requiresWitness,omitempty"`
	// Doses currently on hand.
	QuantityOnHand Nullable[float64] `json:"quantityOnHand,omitempty"`
	// Refills left on the prescription.
	RefillsRemaining Nullable[int64] `json:"refillsRemaining,omitempty"`
	// Other notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m MedicationUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Strength.IsOmitted() {
		w.add("strength", m.Strength)
	}
	if !m.Form.IsOmitted() {
		w.add("form", m.Form)
	}
	if !m.Route.IsOmitted() {
		w.add("route", m.Route)
	}
	if !m.Directions.IsOmitted() {
		w.add("directions", m.Directions)
	}
	if !m.Indication.IsOmitted() {
		w.add("indication", m.Indication)
	}
	if !m.ScheduleType.IsOmitted() {
		w.add("scheduleType", m.ScheduleType)
	}
	if !m.PrnReason.IsOmitted() {
		w.add("prnReason", m.PrnReason)
	}
	if !m.PrnMaxDailyDoses.IsOmitted() {
		w.add("prnMaxDailyDoses", m.PrnMaxDailyDoses)
	}
	if !m.PrnMinIntervalHours.IsOmitted() {
		w.add("prnMinIntervalHours", m.PrnMinIntervalHours)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.DiscontinuedDate.IsOmitted() {
		w.add("discontinuedDate", m.DiscontinuedDate)
	}
	if !m.DiscontinuedReason.IsOmitted() {
		w.add("discontinuedReason", m.DiscontinuedReason)
	}
	if !m.Prescriber.IsOmitted() {
		w.add("prescriber", m.Prescriber)
	}
	if !m.PrescriberPhone.IsOmitted() {
		w.add("prescriberPhone", m.PrescriberPhone)
	}
	if !m.Pharmacy.IsOmitted() {
		w.add("pharmacy", m.Pharmacy)
	}
	if !m.PharmacyPhone.IsOmitted() {
		w.add("pharmacyPhone", m.PharmacyPhone)
	}
	if !m.RequiresWitness.IsOmitted() {
		w.add("requiresWitness", m.RequiresWitness)
	}
	if !m.QuantityOnHand.IsOmitted() {
		w.add("quantityOnHand", m.QuantityOnHand)
	}
	if !m.RefillsRemaining.IsOmitted() {
		w.add("refillsRemaining", m.RefillsRemaining)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Confirmation that a medication was deleted.
type MedicationDeleted struct {
	// Always `medication`.
	Object string `json:"object"`
	// ID of the deleted medication.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Goals of care set for a patient, such as walking to the mailbox unaided. Every field is always
// present; a field with no value is `null`.
type CareGoal struct {
	// String representing the object's type. Always `care_goal`.
	Object string `json:"object"`
	// Unique identifier of the care goal.
	ID string `json:"id"`
	// The patient this goal is for.
	PatientID *string `json:"patientId"`
	// Short name of the goal.
	Name *string `json:"name"`
	// What success looks like.
	Description *string `json:"description"`
	// When the care goal was created.
	CreatedAt *string `json:"createdAt"`
	// When the care goal was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new care goal. Any other field is refused with `unknown_field`.
type CareGoalCreate struct {
	// The patient this goal is for.
	PatientID string `json:"patientId"`
	// Short name of the goal.
	Name string `json:"name"`
	// What success looks like.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CareGoalCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("name", m.Name)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Fields to change on a care goal. Send only the fields to change; any other field is refused with
// `unknown_field`. `patientId` is set when the care goal is created and cannot be changed
// afterwards.
type CareGoalUpdate struct {
	// Short name of the goal.
	Name *string `json:"name,omitempty"`
	// What success looks like.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CareGoalUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Confirmation that a care goal was deleted.
type CareGoalDeleted struct {
	// Always `care_goal`.
	Object string `json:"object"`
	// ID of the deleted care goal.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Notes your staff and caregivers keep on a patient's profile. Every field is always present; a
// field with no value is `null`.
type PatientNote struct {
	// String representing the object's type. Always `patient_note`.
	Object string `json:"object"`
	// Unique identifier of the patient note.
	ID string `json:"id"`
	// The patient the note is about.
	PatientID *string `json:"patientId"`
	// The note.
	Content *string `json:"content"`
	// The person who wrote the note. Set to the API key's user.
	AuthorID *string `json:"authorId"`
	// Display name of the author.
	AuthorName *string `json:"authorName"`
	// When the patient note was created.
	CreatedAt *string `json:"createdAt"`
	// When the patient note was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new patient note. Any other field is refused with `unknown_field`.
type PatientNoteCreate struct {
	// The patient the note is about.
	PatientID string `json:"patientId"`
	// The note.
	Content string `json:"content"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m PatientNoteCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("content", m.Content)
	return w.bytes()
}

// Fields to change on a patient note. Send only the fields to change; any other field is refused
// with `unknown_field`. `patientId` is set when the patient note is created and cannot be changed
// afterwards.
type PatientNoteUpdate struct {
	// The note.
	Content *string `json:"content,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m PatientNoteUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Content != nil {
		w.add("content", m.Content)
	}
	return w.bytes()
}

// Confirmation that a patient note was deleted.
type PatientNoteDeleted struct {
	// Always `patient_note`.
	Object string `json:"object"`
	// ID of the deleted patient note.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// The medication administration record (MAR): each dose a caregiver gave, or the reason it was not
// given. Read-only. Every field is always present; a field with no value is `null`.
type MedicationAdministration struct {
	// String representing the object's type. Always `medication_administration`.
	Object string `json:"object"`
	// Unique identifier of the medication administration.
	ID string `json:"id"`
	// The medication this dose was for.
	MedicationID *string `json:"medicationId"`
	// The patient.
	PatientID *string `json:"patientId"`
	// The visit during which the dose was recorded.
	VisitID *string `json:"visitId"`
	// Date the dose was due.
	ScheduledDate *string `json:"scheduledDate"`
	// Time slot the dose was due, as `HH:MM`.
	ScheduledSlot *string `json:"scheduledSlot"`
	// What happened to the dose.
	Status *MedicationAdministrationStatus `json:"status"`
	// When the dose was given.
	AdministeredAt *string `json:"administeredAt"`
	// Who gave or recorded the dose.
	AdministeredByName *string `json:"administeredByName"`
	// Why the dose was refused, held or missed.
	ReasonText *string `json:"reasonText"`
	// Note recorded with the dose.
	Note *string `json:"note"`
	// Witness, for medications that require one.
	WitnessName *string `json:"witnessName"`
	// For as-needed doses, why it was given.
	PrnReason *string `json:"prnReason"`
	// Whether this entry was later corrected.
	Amended *bool `json:"amended"`
	// When the medication administration was created.
	CreatedAt *string `json:"createdAt"`
	// When the medication administration was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// OASIS assessments recorded for Medicare home health patients. This lists each assessment and its
// status; the item-level answers are not included. Read-only. Every field is always present; a
// field with no value is `null`.
type OasisAssessment struct {
	// String representing the object's type. Always `oasis_assessment`.
	Object string `json:"object"`
	// Unique identifier of the OASIS assessment.
	ID string `json:"id"`
	// The patient assessed.
	PatientID *string `json:"patientId"`
	// Which OASIS time point this assessment covers.
	Timepoint *OasisAssessmentTimepoint `json:"timepoint"`
	// OASIS data set version.
	SpecVersion *string `json:"specVersion"`
	// Where the assessment is in its lifecycle.
	Status *OasisAssessmentStatus `json:"status"`
	// Date the assessment was completed (OASIS item M0090).
	M0090CompletedDate *string `json:"m0090CompletedDate"`
	// Start of the care episode.
	EpisodeStartDate *string `json:"episodeStartDate"`
	// Date the assessment is due.
	TargetDate *string `json:"targetDate"`
	// Clinician who performed the assessment.
	AssessorName *string `json:"assessorName"`
	// When the assessment was marked complete.
	CompletedAt *string `json:"completedAt"`
	// When the assessment was exported for submission.
	ExportedAt *string `json:"exportedAt"`
	// Correction number; 0 for the original.
	CorrectionNumber *int64 `json:"correctionNumber"`
	// Notes on the assessment.
	Notes *string `json:"notes"`
	// When the OASIS assessment was created.
	CreatedAt *string `json:"createdAt"`
	// When the OASIS assessment was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Caregivers are the aides, nurses and therapists your agency schedules to visit patients. Every
// field is always present; a field with no value is `null`.
type Caregiver struct {
	// String representing the object's type. Always `caregiver`.
	Object string `json:"object"`
	// Unique identifier of the caregiver.
	ID string `json:"id"`
	// Given name.
	FirstName *string `json:"firstName"`
	// Middle name.
	MiddleName *string `json:"middleName"`
	// Family name.
	LastName *string `json:"lastName"`
	// The name the caregiver prefers to be called.
	GoesByName *string `json:"goesByName"`
	// Work email address.
	Email *string `json:"email"`
	// Employment status.
	Status *CaregiverStatus `json:"status"`
	// Whether the caregiver can be scheduled.
	Active *bool `json:"active"`
	// Date of birth.
	DateOfBirth *string `json:"dateOfBirth"`
	// Gender as recorded by the agency.
	Gender *string `json:"gender"`
	// Language the caregiver is most comfortable speaking.
	PrimaryLanguage *string `json:"primaryLanguage"`
	// Home street address.
	Address *string `json:"address"`
	// City.
	City *string `json:"city"`
	// State, as a two-letter code.
	State *string `json:"state"`
	// ZIP code.
	ZipCode *string `json:"zipCode"`
	// The office (branch) the caregiver works from.
	OfficeID *string `json:"officeId"`
	// License or credential type.
	LicenseType *string `json:"licenseType"`
	// National Provider Identifier, for licensed clinicians.
	NpiNumber *string `json:"npiNumber"`
	// State license number.
	StateLicenseNumber *string `json:"stateLicenseNumber"`
	// State that issued the license.
	StateLicenseState *string `json:"stateLicenseState"`
	// License expiry date, as `YYYY-MM-DD`.
	StateLicenseExpiry *string `json:"stateLicenseExpiry"`
	// Whether the caregiver works live-in shifts.
	IsLiveInCaregiver *bool `json:"isLiveInCaregiver"`
	// Notes schedulers see when booking this caregiver.
	SchedulingNotes *string `json:"schedulingNotes"`
	// Date of the caregiver's most recent visit.
	LastVisit *string `json:"lastVisit"`
	// When the caregiver was created.
	CreatedAt *string `json:"createdAt"`
	// When the caregiver was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new caregiver. Any other field is refused with `unknown_field`.
type CaregiverCreate struct {
	// Given name.
	FirstName string `json:"firstName"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName string `json:"lastName"`
	// The name the caregiver prefers to be called.
	GoesByName Nullable[string] `json:"goesByName,omitempty"`
	// Work email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Employment status.
	Status Nullable[CaregiverStatus] `json:"status,omitempty"`
	// Whether the caregiver can be scheduled.
	Active Nullable[bool] `json:"active,omitempty"`
	// Date of birth.
	DateOfBirth Nullable[string] `json:"dateOfBirth,omitempty"`
	// Gender as recorded by the agency.
	Gender Nullable[string] `json:"gender,omitempty"`
	// Language the caregiver is most comfortable speaking.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Home street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// The office (branch) the caregiver works from.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// License or credential type.
	LicenseType Nullable[string] `json:"licenseType,omitempty"`
	// National Provider Identifier, for licensed clinicians.
	NpiNumber Nullable[string] `json:"npiNumber,omitempty"`
	// State license number.
	StateLicenseNumber Nullable[string] `json:"stateLicenseNumber,omitempty"`
	// State that issued the license.
	StateLicenseState Nullable[string] `json:"stateLicenseState,omitempty"`
	// License expiry date, as `YYYY-MM-DD`.
	StateLicenseExpiry Nullable[string] `json:"stateLicenseExpiry,omitempty"`
	// Whether the caregiver works live-in shifts.
	IsLiveInCaregiver Nullable[bool] `json:"isLiveInCaregiver,omitempty"`
	// Notes schedulers see when booking this caregiver.
	SchedulingNotes Nullable[string] `json:"schedulingNotes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CaregiverCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("firstName", m.FirstName)
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	w.add("lastName", m.LastName)
	if !m.GoesByName.IsOmitted() {
		w.add("goesByName", m.GoesByName)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.Active.IsOmitted() {
		w.add("active", m.Active)
	}
	if !m.DateOfBirth.IsOmitted() {
		w.add("dateOfBirth", m.DateOfBirth)
	}
	if !m.Gender.IsOmitted() {
		w.add("gender", m.Gender)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.LicenseType.IsOmitted() {
		w.add("licenseType", m.LicenseType)
	}
	if !m.NpiNumber.IsOmitted() {
		w.add("npiNumber", m.NpiNumber)
	}
	if !m.StateLicenseNumber.IsOmitted() {
		w.add("stateLicenseNumber", m.StateLicenseNumber)
	}
	if !m.StateLicenseState.IsOmitted() {
		w.add("stateLicenseState", m.StateLicenseState)
	}
	if !m.StateLicenseExpiry.IsOmitted() {
		w.add("stateLicenseExpiry", m.StateLicenseExpiry)
	}
	if !m.IsLiveInCaregiver.IsOmitted() {
		w.add("isLiveInCaregiver", m.IsLiveInCaregiver)
	}
	if !m.SchedulingNotes.IsOmitted() {
		w.add("schedulingNotes", m.SchedulingNotes)
	}
	return w.bytes()
}

// Fields to change on a caregiver. Send only the fields to change; any other field is refused with
// `unknown_field`.
type CaregiverUpdate struct {
	// Given name.
	FirstName *string `json:"firstName,omitempty"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName *string `json:"lastName,omitempty"`
	// The name the caregiver prefers to be called.
	GoesByName Nullable[string] `json:"goesByName,omitempty"`
	// Work email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Employment status.
	Status Nullable[CaregiverStatus] `json:"status,omitempty"`
	// Whether the caregiver can be scheduled.
	Active Nullable[bool] `json:"active,omitempty"`
	// Date of birth.
	DateOfBirth Nullable[string] `json:"dateOfBirth,omitempty"`
	// Gender as recorded by the agency.
	Gender Nullable[string] `json:"gender,omitempty"`
	// Language the caregiver is most comfortable speaking.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Home street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// The office (branch) the caregiver works from.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// License or credential type.
	LicenseType Nullable[string] `json:"licenseType,omitempty"`
	// National Provider Identifier, for licensed clinicians.
	NpiNumber Nullable[string] `json:"npiNumber,omitempty"`
	// State license number.
	StateLicenseNumber Nullable[string] `json:"stateLicenseNumber,omitempty"`
	// State that issued the license.
	StateLicenseState Nullable[string] `json:"stateLicenseState,omitempty"`
	// License expiry date, as `YYYY-MM-DD`.
	StateLicenseExpiry Nullable[string] `json:"stateLicenseExpiry,omitempty"`
	// Whether the caregiver works live-in shifts.
	IsLiveInCaregiver Nullable[bool] `json:"isLiveInCaregiver,omitempty"`
	// Notes schedulers see when booking this caregiver.
	SchedulingNotes Nullable[string] `json:"schedulingNotes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CaregiverUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.FirstName != nil {
		w.add("firstName", m.FirstName)
	}
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	if m.LastName != nil {
		w.add("lastName", m.LastName)
	}
	if !m.GoesByName.IsOmitted() {
		w.add("goesByName", m.GoesByName)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.Active.IsOmitted() {
		w.add("active", m.Active)
	}
	if !m.DateOfBirth.IsOmitted() {
		w.add("dateOfBirth", m.DateOfBirth)
	}
	if !m.Gender.IsOmitted() {
		w.add("gender", m.Gender)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.LicenseType.IsOmitted() {
		w.add("licenseType", m.LicenseType)
	}
	if !m.NpiNumber.IsOmitted() {
		w.add("npiNumber", m.NpiNumber)
	}
	if !m.StateLicenseNumber.IsOmitted() {
		w.add("stateLicenseNumber", m.StateLicenseNumber)
	}
	if !m.StateLicenseState.IsOmitted() {
		w.add("stateLicenseState", m.StateLicenseState)
	}
	if !m.StateLicenseExpiry.IsOmitted() {
		w.add("stateLicenseExpiry", m.StateLicenseExpiry)
	}
	if !m.IsLiveInCaregiver.IsOmitted() {
		w.add("isLiveInCaregiver", m.IsLiveInCaregiver)
	}
	if !m.SchedulingNotes.IsOmitted() {
		w.add("schedulingNotes", m.SchedulingNotes)
	}
	return w.bytes()
}

// Confirmation that a caregiver was deleted.
type CaregiverDeleted struct {
	// Always `caregiver`.
	Object string `json:"object"`
	// ID of the deleted caregiver.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Notes your staff keep on a caregiver's profile. Every field is always present; a field with no
// value is `null`.
type CaregiverNote struct {
	// String representing the object's type. Always `caregiver_note`.
	Object string `json:"object"`
	// Unique identifier of the caregiver note.
	ID string `json:"id"`
	// The caregiver the note is about.
	CaregiverID *string `json:"caregiverId"`
	// The note.
	Content *string `json:"content"`
	// The person who wrote the note. Set to the API key's user.
	AuthorID *string `json:"authorId"`
	// Display name of the author.
	AuthorName *string `json:"authorName"`
	// When the caregiver note was created.
	CreatedAt *string `json:"createdAt"`
	// When the caregiver note was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new caregiver note. Any other field is refused with `unknown_field`.
type CaregiverNoteCreate struct {
	// The caregiver the note is about.
	CaregiverID string `json:"caregiverId"`
	// The note.
	Content string `json:"content"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CaregiverNoteCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("caregiverId", m.CaregiverID)
	w.add("content", m.Content)
	return w.bytes()
}

// Fields to change on a caregiver note. Send only the fields to change; any other field is refused
// with `unknown_field`. `caregiverId` is set when the caregiver note is created and cannot be
// changed afterwards.
type CaregiverNoteUpdate struct {
	// The note.
	Content *string `json:"content,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CaregiverNoteUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Content != nil {
		w.add("content", m.Content)
	}
	return w.bytes()
}

// Confirmation that a caregiver note was deleted.
type CaregiverNoteDeleted struct {
	// Always `caregiver_note`.
	Object string `json:"object"`
	// ID of the deleted caregiver note.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Certifications, licenses and background checks on a caregiver's profile, with their expiry
// dates. Every field is always present; a field with no value is `null`.
type Certification struct {
	// String representing the object's type. Always `certification`.
	Object string `json:"object"`
	// Unique identifier of the certification.
	ID string `json:"id"`
	// The caregiver who holds it.
	CaregiverID *string `json:"caregiverId"`
	// Name of the certification.
	Name *string `json:"name"`
	// Kind of credential, such as a certification or a background check.
	CertKind *string `json:"certKind"`
	// Date it was obtained.
	Obtained *string `json:"obtained"`
	// Date it expires.
	Expires *string `json:"expires"`
	// For background checks, the agency that ran the check.
	CheckAgency *string `json:"checkAgency"`
	// For background checks, the state the check covers.
	CheckState *string `json:"checkState"`
	// For background checks, the result.
	CheckResult *string `json:"checkResult"`
	// Who verified the credential.
	VerifiedByName *string `json:"verifiedByName"`
	// When the credential was verified.
	VerifiedAt *string `json:"verifiedAt"`
	// When the certification was created.
	CreatedAt *string `json:"createdAt"`
	// When the certification was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new certification. Any other field is refused with `unknown_field`.
type CertificationCreate struct {
	// The caregiver who holds it.
	CaregiverID string `json:"caregiverId"`
	// Name of the certification.
	Name string `json:"name"`
	// Kind of credential, such as a certification or a background check.
	CertKind Nullable[string] `json:"certKind,omitempty"`
	// Date it was obtained.
	Obtained Nullable[string] `json:"obtained,omitempty"`
	// Date it expires.
	Expires Nullable[string] `json:"expires,omitempty"`
	// For background checks, the agency that ran the check.
	CheckAgency Nullable[string] `json:"checkAgency,omitempty"`
	// For background checks, the state the check covers.
	CheckState Nullable[string] `json:"checkState,omitempty"`
	// For background checks, the result.
	CheckResult Nullable[string] `json:"checkResult,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CertificationCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("caregiverId", m.CaregiverID)
	w.add("name", m.Name)
	if !m.CertKind.IsOmitted() {
		w.add("certKind", m.CertKind)
	}
	if !m.Obtained.IsOmitted() {
		w.add("obtained", m.Obtained)
	}
	if !m.Expires.IsOmitted() {
		w.add("expires", m.Expires)
	}
	if !m.CheckAgency.IsOmitted() {
		w.add("checkAgency", m.CheckAgency)
	}
	if !m.CheckState.IsOmitted() {
		w.add("checkState", m.CheckState)
	}
	if !m.CheckResult.IsOmitted() {
		w.add("checkResult", m.CheckResult)
	}
	return w.bytes()
}

// Fields to change on a certification. Send only the fields to change; any other field is refused
// with `unknown_field`. `caregiverId` is set when the certification is created and cannot be
// changed afterwards.
type CertificationUpdate struct {
	// Name of the certification.
	Name *string `json:"name,omitempty"`
	// Kind of credential, such as a certification or a background check.
	CertKind Nullable[string] `json:"certKind,omitempty"`
	// Date it was obtained.
	Obtained Nullable[string] `json:"obtained,omitempty"`
	// Date it expires.
	Expires Nullable[string] `json:"expires,omitempty"`
	// For background checks, the agency that ran the check.
	CheckAgency Nullable[string] `json:"checkAgency,omitempty"`
	// For background checks, the state the check covers.
	CheckState Nullable[string] `json:"checkState,omitempty"`
	// For background checks, the result.
	CheckResult Nullable[string] `json:"checkResult,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CertificationUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.CertKind.IsOmitted() {
		w.add("certKind", m.CertKind)
	}
	if !m.Obtained.IsOmitted() {
		w.add("obtained", m.Obtained)
	}
	if !m.Expires.IsOmitted() {
		w.add("expires", m.Expires)
	}
	if !m.CheckAgency.IsOmitted() {
		w.add("checkAgency", m.CheckAgency)
	}
	if !m.CheckState.IsOmitted() {
		w.add("checkState", m.CheckState)
	}
	if !m.CheckResult.IsOmitted() {
		w.add("checkResult", m.CheckResult)
	}
	return w.bytes()
}

// Confirmation that a certification was deleted.
type CertificationDeleted struct {
	// Always `certification`.
	Object string `json:"object"`
	// ID of the deleted certification.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// When a caregiver can or cannot work. A window is either a weekly recurring time on one weekday,
// or a one-off date range such as time off. Every field is always present; a field with no value
// is `null`.
type Availability struct {
	// String representing the object's type. Always `availability`.
	Object string `json:"object"`
	// Unique identifier of the availability window.
	ID string `json:"id"`
	// The caregiver.
	CaregiverID *string `json:"caregiverId"`
	// `true` when the caregiver can work in this window, `false` for time off.
	IsAvailable *bool `json:"isAvailable"`
	// Whether the window repeats weekly.
	Recurring *bool `json:"recurring"`
	// For weekly windows, the weekday: 0 is Sunday, 6 is Saturday.
	DayOfWeek *int64 `json:"dayOfWeek"`
	// For weekly windows, the start time in minutes after midnight.
	StartMinute *int64 `json:"startMinute"`
	// For weekly windows, the end time in minutes after midnight.
	EndMinute *int64 `json:"endMinute"`
	// For weekly windows, repeat every this many weeks. Defaults to every week.
	IntervalWeeks *int64 `json:"intervalWeeks"`
	// First date a weekly window applies.
	EffectiveFrom *string `json:"effectiveFrom"`
	// Last date a weekly window applies; `null` for no end.
	EffectiveTo *string `json:"effectiveTo"`
	// For one-off windows, when it starts.
	StartDateTime *string `json:"startDateTime"`
	// For one-off windows, when it ends.
	EndDateTime *string `json:"endDateTime"`
	// Whether these are the caregiver's preferred hours.
	IsPreferred *bool `json:"isPreferred"`
	// Why the caregiver is unavailable, for time off.
	Reason *string `json:"reason"`
	// Whether the window is in effect.
	IsActive *bool `json:"isActive"`
	// When the availability window was created.
	CreatedAt *string `json:"createdAt"`
	// When the availability window was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new availability window. Any other field is refused with `unknown_field`.
type AvailabilityCreate struct {
	// The caregiver.
	CaregiverID string `json:"caregiverId"`
	// `true` when the caregiver can work in this window, `false` for time off.
	IsAvailable Nullable[bool] `json:"isAvailable,omitempty"`
	// Whether the window repeats weekly.
	Recurring Nullable[bool] `json:"recurring,omitempty"`
	// For weekly windows, the weekday: 0 is Sunday, 6 is Saturday.
	DayOfWeek Nullable[int64] `json:"dayOfWeek,omitempty"`
	// For weekly windows, the start time in minutes after midnight.
	StartMinute Nullable[int64] `json:"startMinute,omitempty"`
	// For weekly windows, the end time in minutes after midnight.
	EndMinute Nullable[int64] `json:"endMinute,omitempty"`
	// For weekly windows, repeat every this many weeks. Defaults to every week.
	IntervalWeeks Nullable[int64] `json:"intervalWeeks,omitempty"`
	// First date a weekly window applies.
	EffectiveFrom Nullable[string] `json:"effectiveFrom,omitempty"`
	// Last date a weekly window applies; `null` for no end.
	EffectiveTo Nullable[string] `json:"effectiveTo,omitempty"`
	// For one-off windows, when it starts.
	StartDateTime Nullable[string] `json:"startDateTime,omitempty"`
	// For one-off windows, when it ends.
	EndDateTime Nullable[string] `json:"endDateTime,omitempty"`
	// Whether these are the caregiver's preferred hours.
	IsPreferred Nullable[bool] `json:"isPreferred,omitempty"`
	// Why the caregiver is unavailable, for time off.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Whether the window is in effect.
	IsActive Nullable[bool] `json:"isActive,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AvailabilityCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("caregiverId", m.CaregiverID)
	if !m.IsAvailable.IsOmitted() {
		w.add("isAvailable", m.IsAvailable)
	}
	if !m.Recurring.IsOmitted() {
		w.add("recurring", m.Recurring)
	}
	if !m.DayOfWeek.IsOmitted() {
		w.add("dayOfWeek", m.DayOfWeek)
	}
	if !m.StartMinute.IsOmitted() {
		w.add("startMinute", m.StartMinute)
	}
	if !m.EndMinute.IsOmitted() {
		w.add("endMinute", m.EndMinute)
	}
	if !m.IntervalWeeks.IsOmitted() {
		w.add("intervalWeeks", m.IntervalWeeks)
	}
	if !m.EffectiveFrom.IsOmitted() {
		w.add("effectiveFrom", m.EffectiveFrom)
	}
	if !m.EffectiveTo.IsOmitted() {
		w.add("effectiveTo", m.EffectiveTo)
	}
	if !m.StartDateTime.IsOmitted() {
		w.add("startDateTime", m.StartDateTime)
	}
	if !m.EndDateTime.IsOmitted() {
		w.add("endDateTime", m.EndDateTime)
	}
	if !m.IsPreferred.IsOmitted() {
		w.add("isPreferred", m.IsPreferred)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.IsActive.IsOmitted() {
		w.add("isActive", m.IsActive)
	}
	return w.bytes()
}

// Fields to change on an availability window. Send only the fields to change; any other field is
// refused with `unknown_field`. `caregiverId` is set when the availability window is created and
// cannot be changed afterwards.
type AvailabilityUpdate struct {
	// `true` when the caregiver can work in this window, `false` for time off.
	IsAvailable Nullable[bool] `json:"isAvailable,omitempty"`
	// Whether the window repeats weekly.
	Recurring Nullable[bool] `json:"recurring,omitempty"`
	// For weekly windows, the weekday: 0 is Sunday, 6 is Saturday.
	DayOfWeek Nullable[int64] `json:"dayOfWeek,omitempty"`
	// For weekly windows, the start time in minutes after midnight.
	StartMinute Nullable[int64] `json:"startMinute,omitempty"`
	// For weekly windows, the end time in minutes after midnight.
	EndMinute Nullable[int64] `json:"endMinute,omitempty"`
	// For weekly windows, repeat every this many weeks. Defaults to every week.
	IntervalWeeks Nullable[int64] `json:"intervalWeeks,omitempty"`
	// First date a weekly window applies.
	EffectiveFrom Nullable[string] `json:"effectiveFrom,omitempty"`
	// Last date a weekly window applies; `null` for no end.
	EffectiveTo Nullable[string] `json:"effectiveTo,omitempty"`
	// For one-off windows, when it starts.
	StartDateTime Nullable[string] `json:"startDateTime,omitempty"`
	// For one-off windows, when it ends.
	EndDateTime Nullable[string] `json:"endDateTime,omitempty"`
	// Whether these are the caregiver's preferred hours.
	IsPreferred Nullable[bool] `json:"isPreferred,omitempty"`
	// Why the caregiver is unavailable, for time off.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Whether the window is in effect.
	IsActive Nullable[bool] `json:"isActive,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AvailabilityUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.IsAvailable.IsOmitted() {
		w.add("isAvailable", m.IsAvailable)
	}
	if !m.Recurring.IsOmitted() {
		w.add("recurring", m.Recurring)
	}
	if !m.DayOfWeek.IsOmitted() {
		w.add("dayOfWeek", m.DayOfWeek)
	}
	if !m.StartMinute.IsOmitted() {
		w.add("startMinute", m.StartMinute)
	}
	if !m.EndMinute.IsOmitted() {
		w.add("endMinute", m.EndMinute)
	}
	if !m.IntervalWeeks.IsOmitted() {
		w.add("intervalWeeks", m.IntervalWeeks)
	}
	if !m.EffectiveFrom.IsOmitted() {
		w.add("effectiveFrom", m.EffectiveFrom)
	}
	if !m.EffectiveTo.IsOmitted() {
		w.add("effectiveTo", m.EffectiveTo)
	}
	if !m.StartDateTime.IsOmitted() {
		w.add("startDateTime", m.StartDateTime)
	}
	if !m.EndDateTime.IsOmitted() {
		w.add("endDateTime", m.EndDateTime)
	}
	if !m.IsPreferred.IsOmitted() {
		w.add("isPreferred", m.IsPreferred)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.IsActive.IsOmitted() {
		w.add("isActive", m.IsActive)
	}
	return w.bytes()
}

// Confirmation that an availability window was deleted.
type AvailabilityDeleted struct {
	// Always `availability`.
	Object string `json:"object"`
	// ID of the deleted availability window.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Skills and specialties on a caregiver's profile, such as dementia care or Hoyer lift transfers.
// Every field is always present; a field with no value is `null`.
type Specialization struct {
	// String representing the object's type. Always `specialization`.
	Object string `json:"object"`
	// Unique identifier of the specialization.
	ID string `json:"id"`
	// The caregiver.
	CaregiverID *string `json:"caregiverId"`
	// Name of the skill.
	Name *string `json:"name"`
	// Details.
	Description *string `json:"description"`
	// When the specialization was created.
	CreatedAt *string `json:"createdAt"`
	// When the specialization was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new specialization. Any other field is refused with `unknown_field`.
type SpecializationCreate struct {
	// The caregiver.
	CaregiverID string `json:"caregiverId"`
	// Name of the skill.
	Name string `json:"name"`
	// Details.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m SpecializationCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("caregiverId", m.CaregiverID)
	w.add("name", m.Name)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Fields to change on a specialization. Send only the fields to change; any other field is refused
// with `unknown_field`. `caregiverId` is set when the specialization is created and cannot be
// changed afterwards.
type SpecializationUpdate struct {
	// Name of the skill.
	Name *string `json:"name,omitempty"`
	// Details.
	Description Nullable[string] `json:"description,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m SpecializationUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	return w.bytes()
}

// Confirmation that a specialization was deleted.
type SpecializationDeleted struct {
	// Always `specialization`.
	Object string `json:"object"`
	// ID of the deleted specialization.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// People who applied to work for your agency, tracked through your hiring stages until they are
// hired as caregivers. Every field is always present; a field with no value is `null`.
type Applicant struct {
	// String representing the object's type. Always `applicant`.
	Object string `json:"object"`
	// Unique identifier of the applicant.
	ID string `json:"id"`
	// Given name.
	FirstName *string `json:"firstName"`
	// Middle name.
	MiddleName *string `json:"middleName"`
	// Family name.
	LastName *string `json:"lastName"`
	// Email address.
	Email *string `json:"email"`
	// Phone number.
	Phone *string `json:"phone"`
	// Hiring stage, as named in your pipeline.
	Stage *string `json:"stage"`
	// When the applicant last moved stage.
	StageChangedAt *string `json:"stageChangedAt"`
	// When the application was received.
	ApplicationDate *string `json:"applicationDate"`
	// Where the applicant found you.
	ApplicationSource *string `json:"applicationSource"`
	// Who referred the applicant.
	ReferredBy *string `json:"referredBy"`
	// City.
	City *string `json:"city"`
	// State, as a two-letter code.
	State *string `json:"state"`
	// ZIP code.
	ZipCode *string `json:"zipCode"`
	// Earliest date the applicant can start.
	DesiredStartDate *string `json:"desiredStartDate"`
	// Hourly pay the applicant is asking for, in dollars.
	DesiredPayRate *float64 `json:"desiredPayRate"`
	// Hours per week the applicant can work.
	AvailableHoursPerWeek *int64 `json:"availableHoursPerWeek"`
	// Years of caregiving experience.
	YearsExperience *int64 `json:"yearsExperience"`
	// Primary language.
	PrimaryLanguage *string `json:"primaryLanguage"`
	// Whether the applicant has a driver's license.
	HasDriversLicense *bool `json:"hasDriversLicense"`
	// Whether the applicant has reliable transportation.
	HasReliableTransportation *bool `json:"hasReliableTransportation"`
	// Whether the applicant holds a current CPR certification.
	HasCprCertification *bool `json:"hasCPRCertification"`
	// Scheduled or completed interview time.
	InterviewDate *string `json:"interviewDate"`
	// Status of any job offer.
	OfferStatus *string `json:"offerStatus"`
	// The office the applicant applied to.
	OfficeID *string `json:"officeId"`
	// Recruiter notes.
	Notes *string `json:"notes"`
	// The caregiver record created when the applicant was hired.
	ConvertedToCaregiverID *string `json:"convertedToCaregiverId"`
	// When the applicant was created.
	CreatedAt *string `json:"createdAt"`
	// When the applicant was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new applicant. Any other field is refused with `unknown_field`.
type ApplicantCreate struct {
	// Given name.
	FirstName string `json:"firstName"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName string `json:"lastName"`
	// Email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Phone number.
	Phone Nullable[string] `json:"phone,omitempty"`
	// Hiring stage, as named in your pipeline.
	Stage Nullable[string] `json:"stage,omitempty"`
	// When the application was received.
	ApplicationDate Nullable[string] `json:"applicationDate,omitempty"`
	// Where the applicant found you.
	ApplicationSource Nullable[string] `json:"applicationSource,omitempty"`
	// Who referred the applicant.
	ReferredBy Nullable[string] `json:"referredBy,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Earliest date the applicant can start.
	DesiredStartDate Nullable[string] `json:"desiredStartDate,omitempty"`
	// Hourly pay the applicant is asking for, in dollars.
	DesiredPayRate Nullable[float64] `json:"desiredPayRate,omitempty"`
	// Hours per week the applicant can work.
	AvailableHoursPerWeek Nullable[int64] `json:"availableHoursPerWeek,omitempty"`
	// Years of caregiving experience.
	YearsExperience Nullable[int64] `json:"yearsExperience,omitempty"`
	// Primary language.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Whether the applicant has a driver's license.
	HasDriversLicense Nullable[bool] `json:"hasDriversLicense,omitempty"`
	// Whether the applicant has reliable transportation.
	HasReliableTransportation Nullable[bool] `json:"hasReliableTransportation,omitempty"`
	// Whether the applicant holds a current CPR certification.
	HasCprCertification Nullable[bool] `json:"hasCPRCertification,omitempty"`
	// Scheduled or completed interview time.
	InterviewDate Nullable[string] `json:"interviewDate,omitempty"`
	// Status of any job offer.
	OfferStatus Nullable[string] `json:"offerStatus,omitempty"`
	// The office the applicant applied to.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// Recruiter notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ApplicantCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("firstName", m.FirstName)
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	w.add("lastName", m.LastName)
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Phone.IsOmitted() {
		w.add("phone", m.Phone)
	}
	if !m.Stage.IsOmitted() {
		w.add("stage", m.Stage)
	}
	if !m.ApplicationDate.IsOmitted() {
		w.add("applicationDate", m.ApplicationDate)
	}
	if !m.ApplicationSource.IsOmitted() {
		w.add("applicationSource", m.ApplicationSource)
	}
	if !m.ReferredBy.IsOmitted() {
		w.add("referredBy", m.ReferredBy)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.DesiredStartDate.IsOmitted() {
		w.add("desiredStartDate", m.DesiredStartDate)
	}
	if !m.DesiredPayRate.IsOmitted() {
		w.add("desiredPayRate", m.DesiredPayRate)
	}
	if !m.AvailableHoursPerWeek.IsOmitted() {
		w.add("availableHoursPerWeek", m.AvailableHoursPerWeek)
	}
	if !m.YearsExperience.IsOmitted() {
		w.add("yearsExperience", m.YearsExperience)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.HasDriversLicense.IsOmitted() {
		w.add("hasDriversLicense", m.HasDriversLicense)
	}
	if !m.HasReliableTransportation.IsOmitted() {
		w.add("hasReliableTransportation", m.HasReliableTransportation)
	}
	if !m.HasCprCertification.IsOmitted() {
		w.add("hasCPRCertification", m.HasCprCertification)
	}
	if !m.InterviewDate.IsOmitted() {
		w.add("interviewDate", m.InterviewDate)
	}
	if !m.OfferStatus.IsOmitted() {
		w.add("offerStatus", m.OfferStatus)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Fields to change on an applicant. Send only the fields to change; any other field is refused
// with `unknown_field`.
type ApplicantUpdate struct {
	// Given name.
	FirstName *string `json:"firstName,omitempty"`
	// Middle name.
	MiddleName Nullable[string] `json:"middleName,omitempty"`
	// Family name.
	LastName *string `json:"lastName,omitempty"`
	// Email address.
	Email Nullable[string] `json:"email,omitempty"`
	// Phone number.
	Phone Nullable[string] `json:"phone,omitempty"`
	// Hiring stage, as named in your pipeline.
	Stage Nullable[string] `json:"stage,omitempty"`
	// When the application was received.
	ApplicationDate Nullable[string] `json:"applicationDate,omitempty"`
	// Where the applicant found you.
	ApplicationSource Nullable[string] `json:"applicationSource,omitempty"`
	// Who referred the applicant.
	ReferredBy Nullable[string] `json:"referredBy,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Earliest date the applicant can start.
	DesiredStartDate Nullable[string] `json:"desiredStartDate,omitempty"`
	// Hourly pay the applicant is asking for, in dollars.
	DesiredPayRate Nullable[float64] `json:"desiredPayRate,omitempty"`
	// Hours per week the applicant can work.
	AvailableHoursPerWeek Nullable[int64] `json:"availableHoursPerWeek,omitempty"`
	// Years of caregiving experience.
	YearsExperience Nullable[int64] `json:"yearsExperience,omitempty"`
	// Primary language.
	PrimaryLanguage Nullable[string] `json:"primaryLanguage,omitempty"`
	// Whether the applicant has a driver's license.
	HasDriversLicense Nullable[bool] `json:"hasDriversLicense,omitempty"`
	// Whether the applicant has reliable transportation.
	HasReliableTransportation Nullable[bool] `json:"hasReliableTransportation,omitempty"`
	// Whether the applicant holds a current CPR certification.
	HasCprCertification Nullable[bool] `json:"hasCPRCertification,omitempty"`
	// Scheduled or completed interview time.
	InterviewDate Nullable[string] `json:"interviewDate,omitempty"`
	// Status of any job offer.
	OfferStatus Nullable[string] `json:"offerStatus,omitempty"`
	// The office the applicant applied to.
	OfficeID Nullable[string] `json:"officeId,omitempty"`
	// Recruiter notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ApplicantUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.FirstName != nil {
		w.add("firstName", m.FirstName)
	}
	if !m.MiddleName.IsOmitted() {
		w.add("middleName", m.MiddleName)
	}
	if m.LastName != nil {
		w.add("lastName", m.LastName)
	}
	if !m.Email.IsOmitted() {
		w.add("email", m.Email)
	}
	if !m.Phone.IsOmitted() {
		w.add("phone", m.Phone)
	}
	if !m.Stage.IsOmitted() {
		w.add("stage", m.Stage)
	}
	if !m.ApplicationDate.IsOmitted() {
		w.add("applicationDate", m.ApplicationDate)
	}
	if !m.ApplicationSource.IsOmitted() {
		w.add("applicationSource", m.ApplicationSource)
	}
	if !m.ReferredBy.IsOmitted() {
		w.add("referredBy", m.ReferredBy)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.DesiredStartDate.IsOmitted() {
		w.add("desiredStartDate", m.DesiredStartDate)
	}
	if !m.DesiredPayRate.IsOmitted() {
		w.add("desiredPayRate", m.DesiredPayRate)
	}
	if !m.AvailableHoursPerWeek.IsOmitted() {
		w.add("availableHoursPerWeek", m.AvailableHoursPerWeek)
	}
	if !m.YearsExperience.IsOmitted() {
		w.add("yearsExperience", m.YearsExperience)
	}
	if !m.PrimaryLanguage.IsOmitted() {
		w.add("primaryLanguage", m.PrimaryLanguage)
	}
	if !m.HasDriversLicense.IsOmitted() {
		w.add("hasDriversLicense", m.HasDriversLicense)
	}
	if !m.HasReliableTransportation.IsOmitted() {
		w.add("hasReliableTransportation", m.HasReliableTransportation)
	}
	if !m.HasCprCertification.IsOmitted() {
		w.add("hasCPRCertification", m.HasCprCertification)
	}
	if !m.InterviewDate.IsOmitted() {
		w.add("interviewDate", m.InterviewDate)
	}
	if !m.OfferStatus.IsOmitted() {
		w.add("offerStatus", m.OfferStatus)
	}
	if !m.OfficeID.IsOmitted() {
		w.add("officeId", m.OfficeID)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Confirmation that an applicant was deleted.
type ApplicantDeleted struct {
	// Always `applicant`.
	Object string `json:"object"`
	// ID of the deleted applicant.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Visits are scheduled or completed caregiver shifts at a patient's home. A visit's patient,
// caregiver, times and status are set when it is created; reschedule or reassign a visit in
// AveeCare. Every field is always present; a field with no value is `null`.
type Visit struct {
	// String representing the object's type. Always `visit`.
	Object string `json:"object"`
	// Unique identifier of the visit.
	ID string `json:"id"`
	// The patient being visited.
	PatientID *string `json:"patientId"`
	// The caregiver assigned to the visit; `null` for an open shift.
	CaregiverID *string `json:"caregiverId"`
	// Scheduled start.
	StartTime *string `json:"startTime"`
	// Scheduled end.
	EndTime *string `json:"endTime"`
	// Where the visit is in its day. New visits are `Scheduled`.
	Status *VisitStatus `json:"status"`
	// Short title shown on the schedule.
	Title *string `json:"title"`
	// Whether the visit spans the whole day, such as a live-in shift.
	IsAllDay *bool `json:"isAllDay"`
	// Tasks the caregiver should complete during the visit.
	Tasks []string `json:"tasks"`
	// Why the visit is scheduled, if it is not routine.
	Reason *string `json:"reason"`
	// Notes for the caregiver about this visit.
	Notes *string `json:"notes"`
	// Billing service code for the visit.
	ServiceCode *string `json:"serviceCode"`
	// Where the visit takes place, when it is not the patient's home address.
	LocationAddress *string `json:"locationAddress"`
	// When the caregiver clocked in.
	ActualStartTime *string `json:"actualStartTime"`
	// When the caregiver clocked out.
	ActualEndTime *string `json:"actualEndTime"`
	// Tasks the caregiver marked done.
	CompletedTasks []string `json:"completedTasks"`
	// Whether the visit passed electronic visit verification.
	EvvVerified *bool `json:"evvVerified"`
	// Electronic visit verification status.
	EvvStatus *VisitEvvStatus `json:"evvStatus"`
	// Whether the visit has been billed.
	BillingStatus *VisitBillingStatus `json:"billingStatus"`
	// The invoice the visit was billed on.
	InvoiceID *string `json:"invoiceId"`
	// Hours billed for the visit.
	BillableHours *float64 `json:"billableHours"`
	// Whether the visit belongs to a repeating series.
	IsRecurring *bool `json:"isRecurring"`
	// The repeating series the visit belongs to.
	RecurrenceSeriesID *string `json:"recurrenceSeriesId"`
	// The caregiver's response to the assignment.
	CaregiverRsvp *VisitCaregiverRsvp `json:"caregiverRSVP"`
	// When the visit was created.
	CreatedAt *string `json:"createdAt"`
	// When the visit was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new visit. Any other field is refused with `unknown_field`.
type VisitCreate struct {
	// The patient being visited.
	PatientID string `json:"patientId"`
	// The caregiver assigned to the visit; `null` for an open shift.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// Scheduled start.
	StartTime string `json:"startTime"`
	// Scheduled end.
	EndTime string `json:"endTime"`
	// Where the visit is in its day. New visits are `Scheduled`.
	Status Nullable[VisitStatus] `json:"status,omitempty"`
	// Short title shown on the schedule.
	Title Nullable[string] `json:"title,omitempty"`
	// Whether the visit spans the whole day, such as a live-in shift.
	IsAllDay Nullable[bool] `json:"isAllDay,omitempty"`
	// Tasks the caregiver should complete during the visit.
	Tasks Nullable[[]string] `json:"tasks,omitempty"`
	// Why the visit is scheduled, if it is not routine.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Notes for the caregiver about this visit.
	Notes Nullable[string] `json:"notes,omitempty"`
	// Billing service code for the visit.
	ServiceCode Nullable[string] `json:"serviceCode,omitempty"`
	// Where the visit takes place, when it is not the patient's home address.
	LocationAddress Nullable[string] `json:"locationAddress,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	w.add("startTime", m.StartTime)
	w.add("endTime", m.EndTime)
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.Title.IsOmitted() {
		w.add("title", m.Title)
	}
	if !m.IsAllDay.IsOmitted() {
		w.add("isAllDay", m.IsAllDay)
	}
	if !m.Tasks.IsOmitted() {
		w.add("tasks", m.Tasks)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	if !m.ServiceCode.IsOmitted() {
		w.add("serviceCode", m.ServiceCode)
	}
	if !m.LocationAddress.IsOmitted() {
		w.add("locationAddress", m.LocationAddress)
	}
	return w.bytes()
}

// Fields to change on a visit. Send only the fields to change; any other field is refused with
// `unknown_field`. `patientId`, `caregiverId`, `startTime`, `endTime`, `status` are set when the
// visit is created and cannot be changed afterwards.
type VisitUpdate struct {
	// Short title shown on the schedule.
	Title Nullable[string] `json:"title,omitempty"`
	// Whether the visit spans the whole day, such as a live-in shift.
	IsAllDay Nullable[bool] `json:"isAllDay,omitempty"`
	// Tasks the caregiver should complete during the visit.
	Tasks Nullable[[]string] `json:"tasks,omitempty"`
	// Why the visit is scheduled, if it is not routine.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Notes for the caregiver about this visit.
	Notes Nullable[string] `json:"notes,omitempty"`
	// Billing service code for the visit.
	ServiceCode Nullable[string] `json:"serviceCode,omitempty"`
	// Where the visit takes place, when it is not the patient's home address.
	LocationAddress Nullable[string] `json:"locationAddress,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.Title.IsOmitted() {
		w.add("title", m.Title)
	}
	if !m.IsAllDay.IsOmitted() {
		w.add("isAllDay", m.IsAllDay)
	}
	if !m.Tasks.IsOmitted() {
		w.add("tasks", m.Tasks)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	if !m.ServiceCode.IsOmitted() {
		w.add("serviceCode", m.ServiceCode)
	}
	if !m.LocationAddress.IsOmitted() {
		w.add("locationAddress", m.LocationAddress)
	}
	return w.bytes()
}

// Confirmation that a visit was deleted.
type VisitDeleted struct {
	// Always `visit`.
	Object string `json:"object"`
	// ID of the deleted visit.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Notes written about a visit by the caregiver or office staff. Every field is always present; a
// field with no value is `null`.
type VisitNote struct {
	// String representing the object's type. Always `visit_note`.
	Object string `json:"object"`
	// Unique identifier of the visit note.
	ID string `json:"id"`
	// The visit the note is about.
	VisitID *string `json:"visitId"`
	// The note.
	Content *string `json:"content"`
	// Whether office staff (`User`) or a caregiver wrote the note. Set from the API key's user.
	AuthorType *VisitNoteAuthorType `json:"authorType"`
	// ID of the staff user or caregiver who wrote the note. Set from the API key's user.
	AuthorID *string `json:"authorId"`
	// When the visit note was created.
	CreatedAt *string `json:"createdAt"`
	// When the visit note was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new visit note. Any other field is refused with `unknown_field`.
type VisitNoteCreate struct {
	// The visit the note is about.
	VisitID string `json:"visitId"`
	// The note.
	Content string `json:"content"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitNoteCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("visitId", m.VisitID)
	w.add("content", m.Content)
	return w.bytes()
}

// Fields to change on a visit note. Send only the fields to change; any other field is refused
// with `unknown_field`. `visitId` is set when the visit note is created and cannot be changed
// afterwards.
type VisitNoteUpdate struct {
	// The note.
	Content *string `json:"content,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitNoteUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Content != nil {
		w.add("content", m.Content)
	}
	return w.bytes()
}

// Confirmation that a visit note was deleted.
type VisitNoteDeleted struct {
	// Always `visit_note`.
	Object string `json:"object"`
	// ID of the deleted visit note.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Requests from caregivers to pick up visits with a patient, which office staff approve or deny.
// Every field is always present; a field with no value is `null`.
type VisitRequest struct {
	// String representing the object's type. Always `visit_request`.
	Object string `json:"object"`
	// Unique identifier of the visit request.
	ID string `json:"id"`
	// The caregiver asking for the visit. With a caregiver's API key this is always that caregiver
	// (leave it out); office staff must send it.
	CaregiverID *string `json:"caregiverId"`
	// Caregiver's display name, taken from the caregiver's record.
	CaregiverName *string `json:"caregiverName"`
	// The patient the caregiver would visit.
	PatientID *string `json:"patientId"`
	// Patient's display name.
	PatientName *string `json:"patientName"`
	// First date the caregiver could start.
	PreferredStartDate *string `json:"preferredStartDate"`
	// Last date of the request.
	PreferredEndDate *string `json:"preferredEndDate"`
	// Time of day the caregiver prefers.
	PreferredTimeOfDay *string `json:"preferredTimeOfDay"`
	// Why the caregiver is asking.
	Reason *string `json:"reason"`
	// Where the request stands. A caregiver's key creates `Pending` requests and can only change its
	// own pending request to `Cancelled`; approving and denying is for office staff.
	Status *VisitRequestStatus `json:"status"`
	// The office's reply. Set by office staff.
	AdminResponse *string `json:"adminResponse"`
	// Who answered the request.
	RespondedByName *string `json:"respondedByName"`
	// When the request was answered.
	RespondedAt *string `json:"respondedAt"`
	// When the visit request was created.
	CreatedAt *string `json:"createdAt"`
	// When the visit request was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new visit request. Any other field is refused with `unknown_field`.
type VisitRequestCreate struct {
	// The caregiver asking for the visit. With a caregiver's API key this is always that caregiver
	// (leave it out); office staff must send it.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The patient the caregiver would visit.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// Patient's display name.
	PatientName Nullable[string] `json:"patientName,omitempty"`
	// First date the caregiver could start.
	PreferredStartDate Nullable[string] `json:"preferredStartDate,omitempty"`
	// Last date of the request.
	PreferredEndDate Nullable[string] `json:"preferredEndDate,omitempty"`
	// Time of day the caregiver prefers.
	PreferredTimeOfDay Nullable[string] `json:"preferredTimeOfDay,omitempty"`
	// Why the caregiver is asking.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Where the request stands. A caregiver's key creates `Pending` requests and can only change its
	// own pending request to `Cancelled`; approving and denying is for office staff.
	Status Nullable[VisitRequestStatus] `json:"status,omitempty"`
	// The office's reply. Set by office staff.
	AdminResponse Nullable[string] `json:"adminResponse,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitRequestCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.PatientName.IsOmitted() {
		w.add("patientName", m.PatientName)
	}
	if !m.PreferredStartDate.IsOmitted() {
		w.add("preferredStartDate", m.PreferredStartDate)
	}
	if !m.PreferredEndDate.IsOmitted() {
		w.add("preferredEndDate", m.PreferredEndDate)
	}
	if !m.PreferredTimeOfDay.IsOmitted() {
		w.add("preferredTimeOfDay", m.PreferredTimeOfDay)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.AdminResponse.IsOmitted() {
		w.add("adminResponse", m.AdminResponse)
	}
	return w.bytes()
}

// Fields to change on a visit request. Send only the fields to change; any other field is refused
// with `unknown_field`. `caregiverId` is set when the visit request is created and cannot be
// changed afterwards.
type VisitRequestUpdate struct {
	// The patient the caregiver would visit.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// Patient's display name.
	PatientName Nullable[string] `json:"patientName,omitempty"`
	// First date the caregiver could start.
	PreferredStartDate Nullable[string] `json:"preferredStartDate,omitempty"`
	// Last date of the request.
	PreferredEndDate Nullable[string] `json:"preferredEndDate,omitempty"`
	// Time of day the caregiver prefers.
	PreferredTimeOfDay Nullable[string] `json:"preferredTimeOfDay,omitempty"`
	// Why the caregiver is asking.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Where the request stands. A caregiver's key creates `Pending` requests and can only change its
	// own pending request to `Cancelled`; approving and denying is for office staff.
	Status Nullable[VisitRequestStatus] `json:"status,omitempty"`
	// The office's reply. Set by office staff.
	AdminResponse Nullable[string] `json:"adminResponse,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m VisitRequestUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.PatientName.IsOmitted() {
		w.add("patientName", m.PatientName)
	}
	if !m.PreferredStartDate.IsOmitted() {
		w.add("preferredStartDate", m.PreferredStartDate)
	}
	if !m.PreferredEndDate.IsOmitted() {
		w.add("preferredEndDate", m.PreferredEndDate)
	}
	if !m.PreferredTimeOfDay.IsOmitted() {
		w.add("preferredTimeOfDay", m.PreferredTimeOfDay)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.AdminResponse.IsOmitted() {
		w.add("adminResponse", m.AdminResponse)
	}
	return w.bytes()
}

// Confirmation that a visit request was deleted.
type VisitRequestDeleted struct {
	// Always `visit_request`.
	Object string `json:"object"`
	// ID of the deleted visit request.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Open shifts are unassigned visits offered to caregivers so one of them can claim it. Every field
// is always present; a field with no value is `null`.
type OpenShift struct {
	// String representing the object's type. Always `open_shift`.
	Object string `json:"object"`
	// Unique identifier of the open shift.
	ID string `json:"id"`
	// The unassigned visit being offered.
	VisitID *string `json:"visitId"`
	// When the shift was offered.
	BroadcastDate *string `json:"broadcastDate"`
	// Whether the offer is still open.
	Status *OpenShiftStatus `json:"status"`
	// The caregiver who claimed the shift.
	AcceptedBy *string `json:"acceptedBy"`
	// When the open shift was created.
	CreatedAt *string `json:"createdAt"`
	// When the open shift was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new open shift. Any other field is refused with `unknown_field`.
type OpenShiftCreate struct {
	// The unassigned visit being offered.
	VisitID string `json:"visitId"`
	// When the shift was offered.
	BroadcastDate Nullable[string] `json:"broadcastDate,omitempty"`
	// Whether the offer is still open.
	Status Nullable[OpenShiftStatus] `json:"status,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m OpenShiftCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("visitId", m.VisitID)
	if !m.BroadcastDate.IsOmitted() {
		w.add("broadcastDate", m.BroadcastDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	return w.bytes()
}

// Fields to change on an open shift. Send only the fields to change; any other field is refused
// with `unknown_field`. `visitId` is set when the open shift is created and cannot be changed
// afterwards.
type OpenShiftUpdate struct {
	// When the shift was offered.
	BroadcastDate Nullable[string] `json:"broadcastDate,omitempty"`
	// Whether the offer is still open.
	Status Nullable[OpenShiftStatus] `json:"status,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m OpenShiftUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.BroadcastDate.IsOmitted() {
		w.add("broadcastDate", m.BroadcastDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	return w.bytes()
}

// Confirmation that an open shift was deleted.
type OpenShiftDeleted struct {
	// Always `open_shift`.
	Object string `json:"object"`
	// ID of the deleted open shift.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Requests from one caregiver to hand a visit to another, which the other caregiver and then the
// office approve. Every field is always present; a field with no value is `null`.
type ShiftSwap struct {
	// String representing the object's type. Always `shift_swap`.
	Object string `json:"object"`
	// Unique identifier of the shift swap.
	ID string `json:"id"`
	// The visit to hand over.
	VisitID *string `json:"visitId"`
	// The caregiver giving up the visit. With a caregiver's API key this is always that caregiver
	// (leave it out), and the visit must be assigned to them; office staff must send it.
	RequestingCaregiverID *string `json:"requestingCaregiverId"`
	// Display name of the requesting caregiver, taken from the caregiver's record.
	RequestingCaregiverName *string `json:"requestingCaregiverName"`
	// The caregiver asked to take the visit.
	TargetCaregiverID *string `json:"targetCaregiverId"`
	// Display name of the target caregiver.
	TargetCaregiverName *string `json:"targetCaregiverName"`
	// Why the swap is requested.
	Reason *string `json:"reason"`
	// Where the swap stands. With a caregiver's key: new swaps are `Pending`; the asked caregiver
	// changes a `Pending` swap to `Accepted` or `Declined`; the requesting caregiver can change it to
	// `Cancelled`. `AdminApproved` and `AdminDenied` are for office staff.
	Status *ShiftSwapStatus `json:"status"`
	// When the swap was requested.
	RequestDate *string `json:"requestDate"`
	// When the target caregiver answered.
	ResponseDate *string `json:"responseDate"`
	// When the office answered.
	AdminResponseDate *string `json:"adminResponseDate"`
	// When the shift swap was created.
	CreatedAt *string `json:"createdAt"`
	// When the shift swap was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new shift swap. Any other field is refused with `unknown_field`.
type ShiftSwapCreate struct {
	// The visit to hand over.
	VisitID string `json:"visitId"`
	// The caregiver giving up the visit. With a caregiver's API key this is always that caregiver
	// (leave it out), and the visit must be assigned to them; office staff must send it.
	RequestingCaregiverID Nullable[string] `json:"requestingCaregiverId,omitempty"`
	// The caregiver asked to take the visit.
	TargetCaregiverID Nullable[string] `json:"targetCaregiverId,omitempty"`
	// Display name of the target caregiver.
	TargetCaregiverName Nullable[string] `json:"targetCaregiverName,omitempty"`
	// Why the swap is requested.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Where the swap stands. With a caregiver's key: new swaps are `Pending`; the asked caregiver
	// changes a `Pending` swap to `Accepted` or `Declined`; the requesting caregiver can change it to
	// `Cancelled`. `AdminApproved` and `AdminDenied` are for office staff.
	Status Nullable[ShiftSwapStatus] `json:"status,omitempty"`
	// When the swap was requested.
	RequestDate Nullable[string] `json:"requestDate,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ShiftSwapCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("visitId", m.VisitID)
	if !m.RequestingCaregiverID.IsOmitted() {
		w.add("requestingCaregiverId", m.RequestingCaregiverID)
	}
	if !m.TargetCaregiverID.IsOmitted() {
		w.add("targetCaregiverId", m.TargetCaregiverID)
	}
	if !m.TargetCaregiverName.IsOmitted() {
		w.add("targetCaregiverName", m.TargetCaregiverName)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.RequestDate.IsOmitted() {
		w.add("requestDate", m.RequestDate)
	}
	return w.bytes()
}

// Fields to change on a shift swap. Send only the fields to change; any other field is refused
// with `unknown_field`. `visitId`, `requestingCaregiverId` are set when the shift swap is created
// and cannot be changed afterwards.
type ShiftSwapUpdate struct {
	// The caregiver asked to take the visit.
	TargetCaregiverID Nullable[string] `json:"targetCaregiverId,omitempty"`
	// Display name of the target caregiver.
	TargetCaregiverName Nullable[string] `json:"targetCaregiverName,omitempty"`
	// Why the swap is requested.
	Reason Nullable[string] `json:"reason,omitempty"`
	// Where the swap stands. With a caregiver's key: new swaps are `Pending`; the asked caregiver
	// changes a `Pending` swap to `Accepted` or `Declined`; the requesting caregiver can change it to
	// `Cancelled`. `AdminApproved` and `AdminDenied` are for office staff.
	Status Nullable[ShiftSwapStatus] `json:"status,omitempty"`
	// When the swap was requested.
	RequestDate Nullable[string] `json:"requestDate,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ShiftSwapUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.TargetCaregiverID.IsOmitted() {
		w.add("targetCaregiverId", m.TargetCaregiverID)
	}
	if !m.TargetCaregiverName.IsOmitted() {
		w.add("targetCaregiverName", m.TargetCaregiverName)
	}
	if !m.Reason.IsOmitted() {
		w.add("reason", m.Reason)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.RequestDate.IsOmitted() {
		w.add("requestDate", m.RequestDate)
	}
	return w.bytes()
}

// Confirmation that a shift swap was deleted.
type ShiftSwapDeleted struct {
	// Always `shift_swap`.
	Object string `json:"object"`
	// ID of the deleted shift swap.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Care plans describe the services a patient receives, how often, and when the plan is next
// reviewed. The plan's goals and task lists are managed in AveeCare and are not included here.
// Every field is always present; a field with no value is `null`.
type CarePlan struct {
	// String representing the object's type. Always `care_plan`.
	Object string `json:"object"`
	// Unique identifier of the care plan.
	ID string `json:"id"`
	// The patient the plan is for.
	PatientID *string `json:"patientId"`
	// Plan title.
	Title *string `json:"title"`
	// Where the plan is in its lifecycle.
	Status *CarePlanStatus `json:"status"`
	// Type of service the plan covers.
	ServiceType *string `json:"serviceType"`
	// How often visits happen under the plan.
	Frequency *string `json:"frequency"`
	// Date the plan takes effect.
	StartDate *string `json:"startDate"`
	// Date the plan ends; `null` if open-ended.
	EndDate *string `json:"endDate"`
	// Date the plan is next due for review.
	ReviewDate *string `json:"reviewDate"`
	// Days between reviews.
	ReviewIntervalDays *int64 `json:"reviewIntervalDays"`
	// When the plan was last reviewed.
	LastReviewedAt *string `json:"lastReviewedAt"`
	// Who last reviewed the plan.
	LastReviewedByName *string `json:"lastReviewedByName"`
	// Who created the plan.
	CreatedByName *string `json:"createdByName"`
	// When the care plan was created.
	CreatedAt *string `json:"createdAt"`
	// When the care plan was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new care plan. Any other field is refused with `unknown_field`.
type CarePlanCreate struct {
	// The patient the plan is for.
	PatientID string `json:"patientId"`
	// Plan title.
	Title string `json:"title"`
	// Where the plan is in its lifecycle.
	Status Nullable[CarePlanStatus] `json:"status,omitempty"`
	// Type of service the plan covers.
	ServiceType Nullable[string] `json:"serviceType,omitempty"`
	// How often visits happen under the plan.
	Frequency Nullable[string] `json:"frequency,omitempty"`
	// Date the plan takes effect.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Date the plan ends; `null` if open-ended.
	EndDate Nullable[string] `json:"endDate,omitempty"`
	// Date the plan is next due for review.
	ReviewDate Nullable[string] `json:"reviewDate,omitempty"`
	// Days between reviews.
	ReviewIntervalDays Nullable[int64] `json:"reviewIntervalDays,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CarePlanCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	w.add("title", m.Title)
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.Frequency.IsOmitted() {
		w.add("frequency", m.Frequency)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.EndDate.IsOmitted() {
		w.add("endDate", m.EndDate)
	}
	if !m.ReviewDate.IsOmitted() {
		w.add("reviewDate", m.ReviewDate)
	}
	if !m.ReviewIntervalDays.IsOmitted() {
		w.add("reviewIntervalDays", m.ReviewIntervalDays)
	}
	return w.bytes()
}

// Fields to change on a care plan. Send only the fields to change; any other field is refused with
// `unknown_field`. `patientId` is set when the care plan is created and cannot be changed
// afterwards.
type CarePlanUpdate struct {
	// Plan title.
	Title *string `json:"title,omitempty"`
	// Where the plan is in its lifecycle.
	Status Nullable[CarePlanStatus] `json:"status,omitempty"`
	// Type of service the plan covers.
	ServiceType Nullable[string] `json:"serviceType,omitempty"`
	// How often visits happen under the plan.
	Frequency Nullable[string] `json:"frequency,omitempty"`
	// Date the plan takes effect.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Date the plan ends; `null` if open-ended.
	EndDate Nullable[string] `json:"endDate,omitempty"`
	// Date the plan is next due for review.
	ReviewDate Nullable[string] `json:"reviewDate,omitempty"`
	// Days between reviews.
	ReviewIntervalDays Nullable[int64] `json:"reviewIntervalDays,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m CarePlanUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Title != nil {
		w.add("title", m.Title)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.Frequency.IsOmitted() {
		w.add("frequency", m.Frequency)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.EndDate.IsOmitted() {
		w.add("endDate", m.EndDate)
	}
	if !m.ReviewDate.IsOmitted() {
		w.add("reviewDate", m.ReviewDate)
	}
	if !m.ReviewIntervalDays.IsOmitted() {
		w.add("reviewIntervalDays", m.ReviewIntervalDays)
	}
	return w.bytes()
}

// Confirmation that a care plan was deleted.
type CarePlanDeleted struct {
	// Always `care_plan`.
	Object string `json:"object"`
	// ID of the deleted care plan.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Progress notes are the clinical summaries caregivers write after a visit. Every field is always
// present; a field with no value is `null`.
type ProgressNote struct {
	// String representing the object's type. Always `progress_note`.
	Object string `json:"object"`
	// Unique identifier of the progress note.
	ID string `json:"id"`
	// The visit the note documents.
	VisitID *string `json:"visitId"`
	// The patient.
	PatientID *string `json:"patientId"`
	// The caregiver who wrote the note. With a caregiver's API key this is always that caregiver
	// (leave it out), and the visit must be assigned to them.
	CaregiverID *string `json:"caregiverId"`
	// Display name of the caregiver, taken from the caregiver's record.
	CaregiverName *string `json:"caregiverName"`
	// The written summary of the visit.
	Summary *string `json:"summary"`
	// Whether the summary was drafted with AI assistance.
	AiGenerated *bool `json:"aiGenerated"`
	// When the progress note was created.
	CreatedAt *string `json:"createdAt"`
	// When the progress note was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new progress note. Any other field is refused with `unknown_field`.
type ProgressNoteCreate struct {
	// The visit the note documents.
	VisitID string `json:"visitId"`
	// The patient.
	PatientID string `json:"patientId"`
	// The caregiver who wrote the note. With a caregiver's API key this is always that caregiver
	// (leave it out), and the visit must be assigned to them.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The written summary of the visit.
	Summary string `json:"summary"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ProgressNoteCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("visitId", m.VisitID)
	w.add("patientId", m.PatientID)
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	w.add("summary", m.Summary)
	return w.bytes()
}

// Fields to change on a progress note. Send only the fields to change; any other field is refused
// with `unknown_field`. `visitId`, `patientId`, `caregiverId` are set when the progress note is
// created and cannot be changed afterwards.
type ProgressNoteUpdate struct {
	// The written summary of the visit.
	Summary *string `json:"summary,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ProgressNoteUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Summary != nil {
		w.add("summary", m.Summary)
	}
	return w.bytes()
}

// Confirmation that a progress note was deleted.
type ProgressNoteDeleted struct {
	// Always `progress_note`.
	Object string `json:"object"`
	// ID of the deleted progress note.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Payer authorizations: how many units of a service a payer has approved for a patient, and over
// what dates. Every field is always present; a field with no value is `null`.
type Authorization struct {
	// String representing the object's type. Always `authorization`.
	Object string `json:"object"`
	// Unique identifier of the authorization.
	ID string `json:"id"`
	// The patient the authorization covers.
	PatientID *string `json:"patientId"`
	// The payer that issued it.
	PayerName *string `json:"payerName"`
	// The payer's authorization number.
	AuthorizationNumber *string `json:"authorizationNumber"`
	// Service authorized.
	ServiceType *string `json:"serviceType"`
	// Procedure (HCPCS) code billed against it.
	ProcedureCode *string `json:"procedureCode"`
	// Procedure code modifiers, comma separated.
	Modifiers *string `json:"modifiers"`
	// First date the authorization covers.
	StartDate *string `json:"startDate"`
	// Last date the authorization covers.
	EndDate *string `json:"endDate"`
	// Units approved.
	AuthorizedUnits *int64 `json:"authorizedUnits"`
	// What a unit is, in the payer's words.
	UnitType *string `json:"unitType"`
	// Whether units count time, visits or days.
	UnitBasis *AuthorizationUnitBasis `json:"unitBasis"`
	// For time-based units, minutes in one unit.
	MinutesPerUnit *int64 `json:"minutesPerUnit"`
	// Units used so far.
	UsedUnits *int64 `json:"usedUnits"`
	// Units left.
	RemainingUnits *int64 `json:"remainingUnits"`
	// Frequency the payer approved.
	Frequency *string `json:"frequency"`
	// Place-of-service code.
	PlaceOfService *string `json:"placeOfService"`
	// Whether the authorization is in effect.
	Status *AuthorizationStatus `json:"status"`
	// Notes.
	Notes *string `json:"notes"`
	// When the authorization was created.
	CreatedAt *string `json:"createdAt"`
	// When the authorization was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new authorization. Any other field is refused with `unknown_field`.
type AuthorizationCreate struct {
	// The patient the authorization covers.
	PatientID string `json:"patientId"`
	// The payer that issued it.
	PayerName Nullable[string] `json:"payerName,omitempty"`
	// The payer's authorization number.
	AuthorizationNumber Nullable[string] `json:"authorizationNumber,omitempty"`
	// Service authorized.
	ServiceType Nullable[string] `json:"serviceType,omitempty"`
	// Procedure (HCPCS) code billed against it.
	ProcedureCode Nullable[string] `json:"procedureCode,omitempty"`
	// Procedure code modifiers, comma separated.
	Modifiers Nullable[string] `json:"modifiers,omitempty"`
	// First date the authorization covers.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Last date the authorization covers.
	EndDate Nullable[string] `json:"endDate,omitempty"`
	// Units approved.
	AuthorizedUnits Nullable[int64] `json:"authorizedUnits,omitempty"`
	// What a unit is, in the payer's words.
	UnitType Nullable[string] `json:"unitType,omitempty"`
	// Whether units count time, visits or days.
	UnitBasis Nullable[AuthorizationUnitBasis] `json:"unitBasis,omitempty"`
	// For time-based units, minutes in one unit.
	MinutesPerUnit Nullable[int64] `json:"minutesPerUnit,omitempty"`
	// Frequency the payer approved.
	Frequency Nullable[string] `json:"frequency,omitempty"`
	// Place-of-service code.
	PlaceOfService Nullable[string] `json:"placeOfService,omitempty"`
	// Whether the authorization is in effect.
	Status Nullable[AuthorizationStatus] `json:"status,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AuthorizationCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("patientId", m.PatientID)
	if !m.PayerName.IsOmitted() {
		w.add("payerName", m.PayerName)
	}
	if !m.AuthorizationNumber.IsOmitted() {
		w.add("authorizationNumber", m.AuthorizationNumber)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.ProcedureCode.IsOmitted() {
		w.add("procedureCode", m.ProcedureCode)
	}
	if !m.Modifiers.IsOmitted() {
		w.add("modifiers", m.Modifiers)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.EndDate.IsOmitted() {
		w.add("endDate", m.EndDate)
	}
	if !m.AuthorizedUnits.IsOmitted() {
		w.add("authorizedUnits", m.AuthorizedUnits)
	}
	if !m.UnitType.IsOmitted() {
		w.add("unitType", m.UnitType)
	}
	if !m.UnitBasis.IsOmitted() {
		w.add("unitBasis", m.UnitBasis)
	}
	if !m.MinutesPerUnit.IsOmitted() {
		w.add("minutesPerUnit", m.MinutesPerUnit)
	}
	if !m.Frequency.IsOmitted() {
		w.add("frequency", m.Frequency)
	}
	if !m.PlaceOfService.IsOmitted() {
		w.add("placeOfService", m.PlaceOfService)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Fields to change on an authorization. Send only the fields to change; any other field is refused
// with `unknown_field`. `patientId` is set when the authorization is created and cannot be changed
// afterwards.
type AuthorizationUpdate struct {
	// The payer that issued it.
	PayerName Nullable[string] `json:"payerName,omitempty"`
	// The payer's authorization number.
	AuthorizationNumber Nullable[string] `json:"authorizationNumber,omitempty"`
	// Service authorized.
	ServiceType Nullable[string] `json:"serviceType,omitempty"`
	// Procedure (HCPCS) code billed against it.
	ProcedureCode Nullable[string] `json:"procedureCode,omitempty"`
	// Procedure code modifiers, comma separated.
	Modifiers Nullable[string] `json:"modifiers,omitempty"`
	// First date the authorization covers.
	StartDate Nullable[string] `json:"startDate,omitempty"`
	// Last date the authorization covers.
	EndDate Nullable[string] `json:"endDate,omitempty"`
	// Units approved.
	AuthorizedUnits Nullable[int64] `json:"authorizedUnits,omitempty"`
	// What a unit is, in the payer's words.
	UnitType Nullable[string] `json:"unitType,omitempty"`
	// Whether units count time, visits or days.
	UnitBasis Nullable[AuthorizationUnitBasis] `json:"unitBasis,omitempty"`
	// For time-based units, minutes in one unit.
	MinutesPerUnit Nullable[int64] `json:"minutesPerUnit,omitempty"`
	// Frequency the payer approved.
	Frequency Nullable[string] `json:"frequency,omitempty"`
	// Place-of-service code.
	PlaceOfService Nullable[string] `json:"placeOfService,omitempty"`
	// Whether the authorization is in effect.
	Status Nullable[AuthorizationStatus] `json:"status,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AuthorizationUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.PayerName.IsOmitted() {
		w.add("payerName", m.PayerName)
	}
	if !m.AuthorizationNumber.IsOmitted() {
		w.add("authorizationNumber", m.AuthorizationNumber)
	}
	if !m.ServiceType.IsOmitted() {
		w.add("serviceType", m.ServiceType)
	}
	if !m.ProcedureCode.IsOmitted() {
		w.add("procedureCode", m.ProcedureCode)
	}
	if !m.Modifiers.IsOmitted() {
		w.add("modifiers", m.Modifiers)
	}
	if !m.StartDate.IsOmitted() {
		w.add("startDate", m.StartDate)
	}
	if !m.EndDate.IsOmitted() {
		w.add("endDate", m.EndDate)
	}
	if !m.AuthorizedUnits.IsOmitted() {
		w.add("authorizedUnits", m.AuthorizedUnits)
	}
	if !m.UnitType.IsOmitted() {
		w.add("unitType", m.UnitType)
	}
	if !m.UnitBasis.IsOmitted() {
		w.add("unitBasis", m.UnitBasis)
	}
	if !m.MinutesPerUnit.IsOmitted() {
		w.add("minutesPerUnit", m.MinutesPerUnit)
	}
	if !m.Frequency.IsOmitted() {
		w.add("frequency", m.Frequency)
	}
	if !m.PlaceOfService.IsOmitted() {
		w.add("placeOfService", m.PlaceOfService)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Confirmation that an authorization was deleted.
type AuthorizationDeleted struct {
	// Always `authorization`.
	Object string `json:"object"`
	// ID of the deleted authorization.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Incident reports: falls, injuries, medication errors and other events that need follow-up, often
// for regulatory reporting. Every field is always present; a field with no value is `null`.
type Incident struct {
	// String representing the object's type. Always `incident`.
	Object string `json:"object"`
	// Unique identifier of the incident.
	ID string `json:"id"`
	// Short summary.
	Title *string `json:"title"`
	// What happened.
	Description *string `json:"description"`
	// When the incident happened.
	IncidentDateTime *string `json:"incidentDateTime"`
	// Kind of incident.
	Category *IncidentCategory `json:"category"`
	// How serious it was.
	Severity *IncidentSeverity `json:"severity"`
	// Where follow-up stands. New incidents start `Open`.
	Status *IncidentStatus `json:"status"`
	// The patient involved.
	PatientID *string `json:"patientId"`
	// The caregiver involved.
	CaregiverID *string `json:"caregiverId"`
	// The visit during which it happened.
	VisitID *string `json:"visitId"`
	// What was done at the time.
	ActionsTaken *string `json:"actionsTaken"`
	// Body part injured, if any.
	InjuryBodyPart *string `json:"injuryBodyPart"`
	// Witness, if any.
	WitnessName *string `json:"witnessName"`
	// Whether 911 or other emergency services were called.
	EmergencyServicesContacted *bool `json:"emergencyServicesContacted"`
	// Whether the patient went to hospital.
	HospitalTransport *bool `json:"hospitalTransport"`
	// Hospital, if transported.
	HospitalName *string `json:"hospitalName"`
	// Whether the patient's physician was told.
	PhysicianNotified *bool `json:"physicianNotified"`
	// Whether family was told.
	FamilyNotified *bool `json:"familyNotified"`
	// Whether the incident must be reported to a regulator.
	RegulatoryReportRequired *bool `json:"regulatoryReportRequired"`
	// Whether follow-up is needed.
	FollowUpRequired *bool `json:"followUpRequired"`
	// When follow-up is due.
	FollowUpDueAt *string `json:"followUpDueAt"`
	// Follow-up notes.
	FollowUpNotes *string `json:"followUpNotes"`
	// Whether the incident is resolved.
	Resolved *bool `json:"resolved"`
	// How it was resolved.
	ResolutionNotes *string `json:"resolutionNotes"`
	// Whether only administrators can see it.
	IsConfidential *bool `json:"isConfidential"`
	// Who reported it.
	ReportedByName *string `json:"reportedByName"`
	// When it was reported.
	ReportedAt *string `json:"reportedAt"`
	// When it was resolved.
	ResolvedAt *string `json:"resolvedAt"`
	// When the incident was created.
	CreatedAt *string `json:"createdAt"`
	// When the incident was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new incident. Any other field is refused with `unknown_field`.
type IncidentCreate struct {
	// Short summary.
	Title string `json:"title"`
	// What happened.
	Description Nullable[string] `json:"description,omitempty"`
	// When the incident happened.
	IncidentDateTime Nullable[string] `json:"incidentDateTime,omitempty"`
	// Kind of incident.
	Category Nullable[IncidentCategory] `json:"category,omitempty"`
	// How serious it was.
	Severity Nullable[IncidentSeverity] `json:"severity,omitempty"`
	// Where follow-up stands. New incidents start `Open`. Only keys belonging to office staff can set
	// this.
	Status Nullable[IncidentStatus] `json:"status,omitempty"`
	// The patient involved.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// The caregiver involved.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The visit during which it happened.
	VisitID Nullable[string] `json:"visitId,omitempty"`
	// What was done at the time.
	ActionsTaken Nullable[string] `json:"actionsTaken,omitempty"`
	// Body part injured, if any.
	InjuryBodyPart Nullable[string] `json:"injuryBodyPart,omitempty"`
	// Witness, if any.
	WitnessName Nullable[string] `json:"witnessName,omitempty"`
	// Whether 911 or other emergency services were called.
	EmergencyServicesContacted Nullable[bool] `json:"emergencyServicesContacted,omitempty"`
	// Whether the patient went to hospital.
	HospitalTransport Nullable[bool] `json:"hospitalTransport,omitempty"`
	// Hospital, if transported.
	HospitalName Nullable[string] `json:"hospitalName,omitempty"`
	// Whether the patient's physician was told.
	PhysicianNotified Nullable[bool] `json:"physicianNotified,omitempty"`
	// Whether family was told.
	FamilyNotified Nullable[bool] `json:"familyNotified,omitempty"`
	// Whether the incident must be reported to a regulator.
	RegulatoryReportRequired Nullable[bool] `json:"regulatoryReportRequired,omitempty"`
	// Whether follow-up is needed.
	FollowUpRequired Nullable[bool] `json:"followUpRequired,omitempty"`
	// When follow-up is due. Only keys belonging to office staff can set this.
	FollowUpDueAt Nullable[string] `json:"followUpDueAt,omitempty"`
	// Follow-up notes.
	FollowUpNotes Nullable[string] `json:"followUpNotes,omitempty"`
	// Whether the incident is resolved. Only keys belonging to office staff can set this.
	Resolved Nullable[bool] `json:"resolved,omitempty"`
	// How it was resolved. Only keys belonging to office staff can set this.
	ResolutionNotes Nullable[string] `json:"resolutionNotes,omitempty"`
	// Whether only administrators can see it. Only keys belonging to office staff can set this.
	IsConfidential Nullable[bool] `json:"isConfidential,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m IncidentCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("title", m.Title)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.IncidentDateTime.IsOmitted() {
		w.add("incidentDateTime", m.IncidentDateTime)
	}
	if !m.Category.IsOmitted() {
		w.add("category", m.Category)
	}
	if !m.Severity.IsOmitted() {
		w.add("severity", m.Severity)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	if !m.VisitID.IsOmitted() {
		w.add("visitId", m.VisitID)
	}
	if !m.ActionsTaken.IsOmitted() {
		w.add("actionsTaken", m.ActionsTaken)
	}
	if !m.InjuryBodyPart.IsOmitted() {
		w.add("injuryBodyPart", m.InjuryBodyPart)
	}
	if !m.WitnessName.IsOmitted() {
		w.add("witnessName", m.WitnessName)
	}
	if !m.EmergencyServicesContacted.IsOmitted() {
		w.add("emergencyServicesContacted", m.EmergencyServicesContacted)
	}
	if !m.HospitalTransport.IsOmitted() {
		w.add("hospitalTransport", m.HospitalTransport)
	}
	if !m.HospitalName.IsOmitted() {
		w.add("hospitalName", m.HospitalName)
	}
	if !m.PhysicianNotified.IsOmitted() {
		w.add("physicianNotified", m.PhysicianNotified)
	}
	if !m.FamilyNotified.IsOmitted() {
		w.add("familyNotified", m.FamilyNotified)
	}
	if !m.RegulatoryReportRequired.IsOmitted() {
		w.add("regulatoryReportRequired", m.RegulatoryReportRequired)
	}
	if !m.FollowUpRequired.IsOmitted() {
		w.add("followUpRequired", m.FollowUpRequired)
	}
	if !m.FollowUpDueAt.IsOmitted() {
		w.add("followUpDueAt", m.FollowUpDueAt)
	}
	if !m.FollowUpNotes.IsOmitted() {
		w.add("followUpNotes", m.FollowUpNotes)
	}
	if !m.Resolved.IsOmitted() {
		w.add("resolved", m.Resolved)
	}
	if !m.ResolutionNotes.IsOmitted() {
		w.add("resolutionNotes", m.ResolutionNotes)
	}
	if !m.IsConfidential.IsOmitted() {
		w.add("isConfidential", m.IsConfidential)
	}
	return w.bytes()
}

// Fields to change on an incident. Send only the fields to change; any other field is refused with
// `unknown_field`.
type IncidentUpdate struct {
	// Short summary.
	Title *string `json:"title,omitempty"`
	// What happened.
	Description Nullable[string] `json:"description,omitempty"`
	// When the incident happened.
	IncidentDateTime Nullable[string] `json:"incidentDateTime,omitempty"`
	// Kind of incident.
	Category Nullable[IncidentCategory] `json:"category,omitempty"`
	// How serious it was.
	Severity Nullable[IncidentSeverity] `json:"severity,omitempty"`
	// Where follow-up stands. New incidents start `Open`. Only keys belonging to office staff can set
	// this.
	Status Nullable[IncidentStatus] `json:"status,omitempty"`
	// The patient involved.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// The caregiver involved.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The visit during which it happened.
	VisitID Nullable[string] `json:"visitId,omitempty"`
	// What was done at the time.
	ActionsTaken Nullable[string] `json:"actionsTaken,omitempty"`
	// Body part injured, if any.
	InjuryBodyPart Nullable[string] `json:"injuryBodyPart,omitempty"`
	// Witness, if any.
	WitnessName Nullable[string] `json:"witnessName,omitempty"`
	// Whether 911 or other emergency services were called.
	EmergencyServicesContacted Nullable[bool] `json:"emergencyServicesContacted,omitempty"`
	// Whether the patient went to hospital.
	HospitalTransport Nullable[bool] `json:"hospitalTransport,omitempty"`
	// Hospital, if transported.
	HospitalName Nullable[string] `json:"hospitalName,omitempty"`
	// Whether the patient's physician was told.
	PhysicianNotified Nullable[bool] `json:"physicianNotified,omitempty"`
	// Whether family was told.
	FamilyNotified Nullable[bool] `json:"familyNotified,omitempty"`
	// Whether the incident must be reported to a regulator.
	RegulatoryReportRequired Nullable[bool] `json:"regulatoryReportRequired,omitempty"`
	// Whether follow-up is needed.
	FollowUpRequired Nullable[bool] `json:"followUpRequired,omitempty"`
	// When follow-up is due. Only keys belonging to office staff can set this.
	FollowUpDueAt Nullable[string] `json:"followUpDueAt,omitempty"`
	// Follow-up notes.
	FollowUpNotes Nullable[string] `json:"followUpNotes,omitempty"`
	// Whether the incident is resolved. Only keys belonging to office staff can set this.
	Resolved Nullable[bool] `json:"resolved,omitempty"`
	// How it was resolved. Only keys belonging to office staff can set this.
	ResolutionNotes Nullable[string] `json:"resolutionNotes,omitempty"`
	// Whether only administrators can see it. Only keys belonging to office staff can set this.
	IsConfidential Nullable[bool] `json:"isConfidential,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m IncidentUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Title != nil {
		w.add("title", m.Title)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.IncidentDateTime.IsOmitted() {
		w.add("incidentDateTime", m.IncidentDateTime)
	}
	if !m.Category.IsOmitted() {
		w.add("category", m.Category)
	}
	if !m.Severity.IsOmitted() {
		w.add("severity", m.Severity)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	if !m.VisitID.IsOmitted() {
		w.add("visitId", m.VisitID)
	}
	if !m.ActionsTaken.IsOmitted() {
		w.add("actionsTaken", m.ActionsTaken)
	}
	if !m.InjuryBodyPart.IsOmitted() {
		w.add("injuryBodyPart", m.InjuryBodyPart)
	}
	if !m.WitnessName.IsOmitted() {
		w.add("witnessName", m.WitnessName)
	}
	if !m.EmergencyServicesContacted.IsOmitted() {
		w.add("emergencyServicesContacted", m.EmergencyServicesContacted)
	}
	if !m.HospitalTransport.IsOmitted() {
		w.add("hospitalTransport", m.HospitalTransport)
	}
	if !m.HospitalName.IsOmitted() {
		w.add("hospitalName", m.HospitalName)
	}
	if !m.PhysicianNotified.IsOmitted() {
		w.add("physicianNotified", m.PhysicianNotified)
	}
	if !m.FamilyNotified.IsOmitted() {
		w.add("familyNotified", m.FamilyNotified)
	}
	if !m.RegulatoryReportRequired.IsOmitted() {
		w.add("regulatoryReportRequired", m.RegulatoryReportRequired)
	}
	if !m.FollowUpRequired.IsOmitted() {
		w.add("followUpRequired", m.FollowUpRequired)
	}
	if !m.FollowUpDueAt.IsOmitted() {
		w.add("followUpDueAt", m.FollowUpDueAt)
	}
	if !m.FollowUpNotes.IsOmitted() {
		w.add("followUpNotes", m.FollowUpNotes)
	}
	if !m.Resolved.IsOmitted() {
		w.add("resolved", m.Resolved)
	}
	if !m.ResolutionNotes.IsOmitted() {
		w.add("resolutionNotes", m.ResolutionNotes)
	}
	if !m.IsConfidential.IsOmitted() {
		w.add("isConfidential", m.IsConfidential)
	}
	return w.bytes()
}

// Confirmation that an incident was deleted.
type IncidentDeleted struct {
	// Always `incident`.
	Object string `json:"object"`
	// ID of the deleted incident.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Events your agency documents about a patient, caregiver or visit, such as a complaint, a
// compliment or an attendance issue. Every field is always present; a field with no value is
// `null`.
type Event struct {
	// String representing the object's type. Always `event`.
	Object string `json:"object"`
	// Unique identifier of the event.
	ID string `json:"id"`
	// Short summary.
	Title *string `json:"title"`
	// What happened.
	Description *string `json:"description"`
	// Category, as named in your agency's event types.
	Category *string `json:"category"`
	// How serious it was.
	Severity *IncidentSeverity `json:"severity"`
	// Where follow-up stands.
	Status *IncidentStatus `json:"status"`
	// When it happened.
	OccurredAt *string `json:"occurredAt"`
	// The patient involved.
	PatientID *string `json:"patientId"`
	// The caregiver involved.
	CaregiverID *string `json:"caregiverId"`
	// The visit involved.
	VisitID *string `json:"visitId"`
	// Attendance points this event counts for, if your agency uses a points policy.
	PointValue *float64 `json:"pointValue"`
	// What was done.
	ActionsTaken *string `json:"actionsTaken"`
	// Outcome.
	Outcome *string `json:"outcome"`
	// Whether follow-up is needed.
	FollowUpRequired *bool `json:"followUpRequired"`
	// When follow-up is due.
	FollowUpDueAt *string `json:"followUpDueAt"`
	// Follow-up notes.
	FollowUpNotes *string `json:"followUpNotes"`
	// How it was resolved.
	ResolutionNotes *string `json:"resolutionNotes"`
	// Whether only administrators can see it.
	IsConfidential *bool `json:"isConfidential"`
	// Who documented it.
	ReportedByName *string `json:"reportedByName"`
	// When it was documented.
	ReportedAt *string `json:"reportedAt"`
	// When it was resolved.
	ResolvedAt *string `json:"resolvedAt"`
	// When the event was created.
	CreatedAt *string `json:"createdAt"`
	// When the event was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new event. Any other field is refused with `unknown_field`.
type EventCreate struct {
	// Short summary.
	Title string `json:"title"`
	// What happened.
	Description Nullable[string] `json:"description,omitempty"`
	// Category, as named in your agency's event types.
	Category Nullable[string] `json:"category,omitempty"`
	// How serious it was.
	Severity Nullable[IncidentSeverity] `json:"severity,omitempty"`
	// Where follow-up stands.
	Status Nullable[IncidentStatus] `json:"status,omitempty"`
	// When it happened.
	OccurredAt Nullable[string] `json:"occurredAt,omitempty"`
	// The patient involved.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// The caregiver involved.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The visit involved.
	VisitID Nullable[string] `json:"visitId,omitempty"`
	// Attendance points this event counts for, if your agency uses a points policy.
	PointValue Nullable[float64] `json:"pointValue,omitempty"`
	// What was done.
	ActionsTaken Nullable[string] `json:"actionsTaken,omitempty"`
	// Outcome.
	Outcome Nullable[string] `json:"outcome,omitempty"`
	// Whether follow-up is needed.
	FollowUpRequired Nullable[bool] `json:"followUpRequired,omitempty"`
	// When follow-up is due.
	FollowUpDueAt Nullable[string] `json:"followUpDueAt,omitempty"`
	// Follow-up notes.
	FollowUpNotes Nullable[string] `json:"followUpNotes,omitempty"`
	// How it was resolved.
	ResolutionNotes Nullable[string] `json:"resolutionNotes,omitempty"`
	// Whether only administrators can see it.
	IsConfidential Nullable[bool] `json:"isConfidential,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m EventCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("title", m.Title)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.Category.IsOmitted() {
		w.add("category", m.Category)
	}
	if !m.Severity.IsOmitted() {
		w.add("severity", m.Severity)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.OccurredAt.IsOmitted() {
		w.add("occurredAt", m.OccurredAt)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	if !m.VisitID.IsOmitted() {
		w.add("visitId", m.VisitID)
	}
	if !m.PointValue.IsOmitted() {
		w.add("pointValue", m.PointValue)
	}
	if !m.ActionsTaken.IsOmitted() {
		w.add("actionsTaken", m.ActionsTaken)
	}
	if !m.Outcome.IsOmitted() {
		w.add("outcome", m.Outcome)
	}
	if !m.FollowUpRequired.IsOmitted() {
		w.add("followUpRequired", m.FollowUpRequired)
	}
	if !m.FollowUpDueAt.IsOmitted() {
		w.add("followUpDueAt", m.FollowUpDueAt)
	}
	if !m.FollowUpNotes.IsOmitted() {
		w.add("followUpNotes", m.FollowUpNotes)
	}
	if !m.ResolutionNotes.IsOmitted() {
		w.add("resolutionNotes", m.ResolutionNotes)
	}
	if !m.IsConfidential.IsOmitted() {
		w.add("isConfidential", m.IsConfidential)
	}
	return w.bytes()
}

// Fields to change on an event. Send only the fields to change; any other field is refused with
// `unknown_field`.
type EventUpdate struct {
	// Short summary.
	Title *string `json:"title,omitempty"`
	// What happened.
	Description Nullable[string] `json:"description,omitempty"`
	// Category, as named in your agency's event types.
	Category Nullable[string] `json:"category,omitempty"`
	// How serious it was.
	Severity Nullable[IncidentSeverity] `json:"severity,omitempty"`
	// Where follow-up stands.
	Status Nullable[IncidentStatus] `json:"status,omitempty"`
	// When it happened.
	OccurredAt Nullable[string] `json:"occurredAt,omitempty"`
	// The patient involved.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// The caregiver involved.
	CaregiverID Nullable[string] `json:"caregiverId,omitempty"`
	// The visit involved.
	VisitID Nullable[string] `json:"visitId,omitempty"`
	// Attendance points this event counts for, if your agency uses a points policy.
	PointValue Nullable[float64] `json:"pointValue,omitempty"`
	// What was done.
	ActionsTaken Nullable[string] `json:"actionsTaken,omitempty"`
	// Outcome.
	Outcome Nullable[string] `json:"outcome,omitempty"`
	// Whether follow-up is needed.
	FollowUpRequired Nullable[bool] `json:"followUpRequired,omitempty"`
	// When follow-up is due.
	FollowUpDueAt Nullable[string] `json:"followUpDueAt,omitempty"`
	// Follow-up notes.
	FollowUpNotes Nullable[string] `json:"followUpNotes,omitempty"`
	// How it was resolved.
	ResolutionNotes Nullable[string] `json:"resolutionNotes,omitempty"`
	// Whether only administrators can see it.
	IsConfidential Nullable[bool] `json:"isConfidential,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m EventUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Title != nil {
		w.add("title", m.Title)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.Category.IsOmitted() {
		w.add("category", m.Category)
	}
	if !m.Severity.IsOmitted() {
		w.add("severity", m.Severity)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.OccurredAt.IsOmitted() {
		w.add("occurredAt", m.OccurredAt)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if !m.CaregiverID.IsOmitted() {
		w.add("caregiverId", m.CaregiverID)
	}
	if !m.VisitID.IsOmitted() {
		w.add("visitId", m.VisitID)
	}
	if !m.PointValue.IsOmitted() {
		w.add("pointValue", m.PointValue)
	}
	if !m.ActionsTaken.IsOmitted() {
		w.add("actionsTaken", m.ActionsTaken)
	}
	if !m.Outcome.IsOmitted() {
		w.add("outcome", m.Outcome)
	}
	if !m.FollowUpRequired.IsOmitted() {
		w.add("followUpRequired", m.FollowUpRequired)
	}
	if !m.FollowUpDueAt.IsOmitted() {
		w.add("followUpDueAt", m.FollowUpDueAt)
	}
	if !m.FollowUpNotes.IsOmitted() {
		w.add("followUpNotes", m.FollowUpNotes)
	}
	if !m.ResolutionNotes.IsOmitted() {
		w.add("resolutionNotes", m.ResolutionNotes)
	}
	if !m.IsConfidential.IsOmitted() {
		w.add("isConfidential", m.IsConfidential)
	}
	return w.bytes()
}

// Confirmation that an event was deleted.
type EventDeleted struct {
	// Always `event`.
	Object string `json:"object"`
	// ID of the deleted event.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Hospitals, physicians, case managers and other organizations that refer clients to your agency.
// Every field is always present; a field with no value is `null`.
type ReferralSource struct {
	// String representing the object's type. Always `referral_source`.
	Object string `json:"object"`
	// Unique identifier of the referral source.
	ID string `json:"id"`
	// Name of the source.
	Name *string `json:"name"`
	// Kind of source, such as Hospital or Physician.
	Type *string `json:"type"`
	// Organization the contact works for.
	Organization *string `json:"organization"`
	// Main contact.
	ContactName *string `json:"contactName"`
	// Contact's job title.
	ContactTitle *string `json:"contactTitle"`
	// Contact's phone number.
	ContactPhone *string `json:"contactPhone"`
	// Contact's email address.
	ContactEmail *string `json:"contactEmail"`
	// Street address.
	Address *string `json:"address"`
	// City.
	City *string `json:"city"`
	// State, as a two-letter code.
	State *string `json:"state"`
	// ZIP code.
	ZipCode *string `json:"zipCode"`
	// Notes.
	Notes *string `json:"notes"`
	// Whether you still receive referrals from this source.
	IsActive *bool `json:"isActive"`
	// Referrals received from this source.
	TotalReferrals *int64 `json:"totalReferrals"`
	// Referrals from this source that became clients.
	TotalConversions *int64 `json:"totalConversions"`
	// When the most recent referral arrived.
	LastReferralDate *string `json:"lastReferralDate"`
	// When the referral source was created.
	CreatedAt *string `json:"createdAt"`
	// When the referral source was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new referral source. Any other field is refused with `unknown_field`.
type ReferralSourceCreate struct {
	// Name of the source.
	Name string `json:"name"`
	// Kind of source, such as Hospital or Physician.
	Type Nullable[string] `json:"type,omitempty"`
	// Organization the contact works for.
	Organization Nullable[string] `json:"organization,omitempty"`
	// Main contact.
	ContactName Nullable[string] `json:"contactName,omitempty"`
	// Contact's job title.
	ContactTitle Nullable[string] `json:"contactTitle,omitempty"`
	// Contact's phone number.
	ContactPhone Nullable[string] `json:"contactPhone,omitempty"`
	// Contact's email address.
	ContactEmail Nullable[string] `json:"contactEmail,omitempty"`
	// Street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
	// Whether you still receive referrals from this source.
	IsActive Nullable[bool] `json:"isActive,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ReferralSourceCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("name", m.Name)
	if !m.Type.IsOmitted() {
		w.add("type", m.Type)
	}
	if !m.Organization.IsOmitted() {
		w.add("organization", m.Organization)
	}
	if !m.ContactName.IsOmitted() {
		w.add("contactName", m.ContactName)
	}
	if !m.ContactTitle.IsOmitted() {
		w.add("contactTitle", m.ContactTitle)
	}
	if !m.ContactPhone.IsOmitted() {
		w.add("contactPhone", m.ContactPhone)
	}
	if !m.ContactEmail.IsOmitted() {
		w.add("contactEmail", m.ContactEmail)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	if !m.IsActive.IsOmitted() {
		w.add("isActive", m.IsActive)
	}
	return w.bytes()
}

// Fields to change on a referral source. Send only the fields to change; any other field is
// refused with `unknown_field`.
type ReferralSourceUpdate struct {
	// Name of the source.
	Name *string `json:"name,omitempty"`
	// Kind of source, such as Hospital or Physician.
	Type Nullable[string] `json:"type,omitempty"`
	// Organization the contact works for.
	Organization Nullable[string] `json:"organization,omitempty"`
	// Main contact.
	ContactName Nullable[string] `json:"contactName,omitempty"`
	// Contact's job title.
	ContactTitle Nullable[string] `json:"contactTitle,omitempty"`
	// Contact's phone number.
	ContactPhone Nullable[string] `json:"contactPhone,omitempty"`
	// Contact's email address.
	ContactEmail Nullable[string] `json:"contactEmail,omitempty"`
	// Street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
	// Whether you still receive referrals from this source.
	IsActive Nullable[bool] `json:"isActive,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ReferralSourceUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Type.IsOmitted() {
		w.add("type", m.Type)
	}
	if !m.Organization.IsOmitted() {
		w.add("organization", m.Organization)
	}
	if !m.ContactName.IsOmitted() {
		w.add("contactName", m.ContactName)
	}
	if !m.ContactTitle.IsOmitted() {
		w.add("contactTitle", m.ContactTitle)
	}
	if !m.ContactPhone.IsOmitted() {
		w.add("contactPhone", m.ContactPhone)
	}
	if !m.ContactEmail.IsOmitted() {
		w.add("contactEmail", m.ContactEmail)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	if !m.IsActive.IsOmitted() {
		w.add("isActive", m.IsActive)
	}
	return w.bytes()
}

// Confirmation that a referral source was deleted.
type ReferralSourceDeleted struct {
	// Always `referral_source`.
	Object string `json:"object"`
	// ID of the deleted referral source.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Individual referrals of a prospective client, tracked from first contact to admission. Every
// field is always present; a field with no value is `null`.
type Referral struct {
	// String representing the object's type. Always `referral`.
	Object string `json:"object"`
	// Unique identifier of the referral.
	ID string `json:"id"`
	// Who made the referral.
	ReferralSourceID *string `json:"referralSourceId"`
	// The patient record, once one exists.
	PatientID *string `json:"patientId"`
	// Name of the person referred.
	PatientName *string `json:"patientName"`
	// Phone number to reach the person or their family.
	ContactPhone *string `json:"contactPhone"`
	// Email address to reach the person or their family.
	ContactEmail *string `json:"contactEmail"`
	// When the referral arrived.
	ReferralDate *string `json:"referralDate"`
	// Where the referral stands.
	Status *ReferralStatus `json:"status"`
	// Care the person is expected to need.
	EstimatedCareNeeds *string `json:"estimatedCareNeeds"`
	// Expected hours of care per week.
	EstimatedWeeklyHours *float64 `json:"estimatedWeeklyHours"`
	// How soon care is needed.
	UrgencyLevel *string `json:"urgencyLevel"`
	// Why the referral was declined or lost.
	DeclineReason *string `json:"declineReason"`
	// Notes.
	Notes *string `json:"notes"`
	// When the person was admitted as a client.
	ConvertedDate *string `json:"convertedDate"`
	// When the referral was created.
	CreatedAt *string `json:"createdAt"`
	// When the referral was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new referral. Any other field is refused with `unknown_field`.
type ReferralCreate struct {
	// Who made the referral.
	ReferralSourceID Nullable[string] `json:"referralSourceId,omitempty"`
	// The patient record, once one exists.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// Name of the person referred.
	PatientName string `json:"patientName"`
	// Phone number to reach the person or their family.
	ContactPhone Nullable[string] `json:"contactPhone,omitempty"`
	// Email address to reach the person or their family.
	ContactEmail Nullable[string] `json:"contactEmail,omitempty"`
	// When the referral arrived.
	ReferralDate Nullable[string] `json:"referralDate,omitempty"`
	// Where the referral stands.
	Status Nullable[ReferralStatus] `json:"status,omitempty"`
	// Care the person is expected to need.
	EstimatedCareNeeds Nullable[string] `json:"estimatedCareNeeds,omitempty"`
	// Expected hours of care per week.
	EstimatedWeeklyHours Nullable[float64] `json:"estimatedWeeklyHours,omitempty"`
	// How soon care is needed.
	UrgencyLevel Nullable[string] `json:"urgencyLevel,omitempty"`
	// Why the referral was declined or lost.
	DeclineReason Nullable[string] `json:"declineReason,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ReferralCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.ReferralSourceID.IsOmitted() {
		w.add("referralSourceId", m.ReferralSourceID)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	w.add("patientName", m.PatientName)
	if !m.ContactPhone.IsOmitted() {
		w.add("contactPhone", m.ContactPhone)
	}
	if !m.ContactEmail.IsOmitted() {
		w.add("contactEmail", m.ContactEmail)
	}
	if !m.ReferralDate.IsOmitted() {
		w.add("referralDate", m.ReferralDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.EstimatedCareNeeds.IsOmitted() {
		w.add("estimatedCareNeeds", m.EstimatedCareNeeds)
	}
	if !m.EstimatedWeeklyHours.IsOmitted() {
		w.add("estimatedWeeklyHours", m.EstimatedWeeklyHours)
	}
	if !m.UrgencyLevel.IsOmitted() {
		w.add("urgencyLevel", m.UrgencyLevel)
	}
	if !m.DeclineReason.IsOmitted() {
		w.add("declineReason", m.DeclineReason)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Fields to change on a referral. Send only the fields to change; any other field is refused with
// `unknown_field`.
type ReferralUpdate struct {
	// Who made the referral.
	ReferralSourceID Nullable[string] `json:"referralSourceId,omitempty"`
	// The patient record, once one exists.
	PatientID Nullable[string] `json:"patientId,omitempty"`
	// Name of the person referred.
	PatientName *string `json:"patientName,omitempty"`
	// Phone number to reach the person or their family.
	ContactPhone Nullable[string] `json:"contactPhone,omitempty"`
	// Email address to reach the person or their family.
	ContactEmail Nullable[string] `json:"contactEmail,omitempty"`
	// When the referral arrived.
	ReferralDate Nullable[string] `json:"referralDate,omitempty"`
	// Where the referral stands.
	Status Nullable[ReferralStatus] `json:"status,omitempty"`
	// Care the person is expected to need.
	EstimatedCareNeeds Nullable[string] `json:"estimatedCareNeeds,omitempty"`
	// Expected hours of care per week.
	EstimatedWeeklyHours Nullable[float64] `json:"estimatedWeeklyHours,omitempty"`
	// How soon care is needed.
	UrgencyLevel Nullable[string] `json:"urgencyLevel,omitempty"`
	// Why the referral was declined or lost.
	DeclineReason Nullable[string] `json:"declineReason,omitempty"`
	// Notes.
	Notes Nullable[string] `json:"notes,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m ReferralUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if !m.ReferralSourceID.IsOmitted() {
		w.add("referralSourceId", m.ReferralSourceID)
	}
	if !m.PatientID.IsOmitted() {
		w.add("patientId", m.PatientID)
	}
	if m.PatientName != nil {
		w.add("patientName", m.PatientName)
	}
	if !m.ContactPhone.IsOmitted() {
		w.add("contactPhone", m.ContactPhone)
	}
	if !m.ContactEmail.IsOmitted() {
		w.add("contactEmail", m.ContactEmail)
	}
	if !m.ReferralDate.IsOmitted() {
		w.add("referralDate", m.ReferralDate)
	}
	if !m.Status.IsOmitted() {
		w.add("status", m.Status)
	}
	if !m.EstimatedCareNeeds.IsOmitted() {
		w.add("estimatedCareNeeds", m.EstimatedCareNeeds)
	}
	if !m.EstimatedWeeklyHours.IsOmitted() {
		w.add("estimatedWeeklyHours", m.EstimatedWeeklyHours)
	}
	if !m.UrgencyLevel.IsOmitted() {
		w.add("urgencyLevel", m.UrgencyLevel)
	}
	if !m.DeclineReason.IsOmitted() {
		w.add("declineReason", m.DeclineReason)
	}
	if !m.Notes.IsOmitted() {
		w.add("notes", m.Notes)
	}
	return w.bytes()
}

// Confirmation that a referral was deleted.
type ReferralDeleted struct {
	// Always `referral`.
	Object string `json:"object"`
	// ID of the deleted referral.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Invoices billed to patients and their families for private-pay care. Read-only. Every field is
// always present; a field with no value is `null`.
type Invoice struct {
	// String representing the object's type. Always `invoice`.
	Object string `json:"object"`
	// Unique identifier of the invoice.
	ID string `json:"id"`
	// Invoice number shown to the patient.
	InvoiceNumber *string `json:"invoiceNumber"`
	// The patient billed.
	PatientID *string `json:"patientId"`
	// Payment status.
	Status *InvoiceStatus `json:"status"`
	// Date of the invoice.
	InvoiceDate *string `json:"invoiceDate"`
	// Date payment is due.
	DueDate *string `json:"dueDate"`
	// First day of service billed.
	PeriodStart *string `json:"periodStart"`
	// Last day of service billed.
	PeriodEnd *string `json:"periodEnd"`
	// Description shown on the invoice.
	Description *string `json:"description"`
	// Three-letter currency code.
	Currency *string `json:"currency"`
	// Total before tax, discounts and credits. In cents.
	SubtotalCents *int64 `json:"subtotalCents"`
	// Tax. In cents.
	TaxCents *int64 `json:"taxCents"`
	// Discounts applied. In cents.
	DiscountCents *int64 `json:"discountCents"`
	// Account credit applied. In cents.
	CreditAppliedCents *int64 `json:"creditAppliedCents"`
	// Invoice total. In cents.
	TotalCents *int64 `json:"totalCents"`
	// Amount paid so far. In cents.
	AmountPaidCents *int64 `json:"amountPaidCents"`
	// Amount still owed. In cents.
	AmountDueCents *int64 `json:"amountDueCents"`
	// Visits billed on this invoice.
	VisitIDs []string `json:"visitIds"`
	// Whether AveeCare created the invoice automatically when a visit was completed.
	IsAutoBilled *bool `json:"isAutoBilled"`
	// When the invoice was sent.
	SentAt *string `json:"sentAt"`
	// When the invoice was paid in full.
	PaidDate *string `json:"paidDate"`
	// When the invoice was created.
	CreatedAt *string `json:"createdAt"`
	// When the invoice was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// The individual charges on an invoice. Read-only. Every field is always present; a field with no
// value is `null`.
type InvoiceLineItem struct {
	// String representing the object's type. Always `invoice_line_item`.
	Object string `json:"object"`
	// Unique identifier of the invoice line item.
	ID string `json:"id"`
	// The invoice this line belongs to.
	InvoiceID *string `json:"invoiceId"`
	// Line description.
	Description *string `json:"description"`
	// Quantity, such as hours.
	Quantity *float64 `json:"quantity"`
	// Price per unit. In cents.
	UnitPriceCents *int64 `json:"unitPriceCents"`
	// Line total. In cents.
	TotalCents *int64 `json:"totalCents"`
	// Service code.
	ServiceCode *string `json:"serviceCode"`
	// Service category.
	ServiceCategory *string `json:"serviceCategory"`
	// The visit this line bills.
	VisitID *string `json:"visitId"`
	// Date of the visit.
	VisitDate *string `json:"visitDate"`
	// Position of the line on the invoice.
	SortOrder *int64 `json:"sortOrder"`
	// When the invoice line item was created.
	CreatedAt *string `json:"createdAt"`
	// When the invoice line item was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Insurance claims billed to Medicaid, Medicare and other payers, with their adjudication status.
// Read-only. Every field is always present; a field with no value is `null`.
type Claim struct {
	// String representing the object's type. Always `claim`.
	Object string `json:"object"`
	// Unique identifier of the claim.
	ID string `json:"id"`
	// The patient the claim is for.
	PatientID *string `json:"patientId"`
	// The caregiver who provided the service.
	CaregiverID *string `json:"caregiverId"`
	// The visit billed, for single-visit claims.
	VisitID *string `json:"visitId"`
	// The related invoice, if any.
	InvoiceID *string `json:"invoiceId"`
	// Your claim number.
	ClaimNumber *string `json:"claimNumber"`
	// Claim format: professional (837P) or institutional (837I).
	ClaimType *ClaimClaimType `json:"claimType"`
	// The payer billed.
	PayerName *string `json:"payerName"`
	// The payer's electronic payer ID.
	PayerID *string `json:"payerId"`
	// First date of service, as `YYYY-MM-DD`.
	ServiceDateStart *string `json:"serviceDateStart"`
	// Last date of service, as `YYYY-MM-DD`.
	ServiceDateEnd *string `json:"serviceDateEnd"`
	// Total billed. In cents.
	TotalChargeAmountCents *int64 `json:"totalChargeAmountCents"`
	// Where the claim is in adjudication.
	Status *ClaimStatus `json:"status"`
	// Latest status detail from the payer or clearinghouse.
	StatusMessage *string `json:"statusMessage"`
	// When the status last changed, ISO 8601.
	StatusUpdatedAt *string `json:"statusUpdatedAt"`
	// Amount the payer paid. In cents.
	PaidAmountCents *int64 `json:"paidAmountCents"`
	// Amount the patient owes after the payer. In cents.
	PatientResponsibilityCents *int64 `json:"patientResponsibilityCents"`
	// Authorization number sent on the claim.
	PriorAuthNumber *string `json:"priorAuthNumber"`
	// Payment check or EFT trace number.
	CheckNumber *string `json:"checkNumber"`
	// When the claim was submitted, ISO 8601.
	SubmittedAt *string `json:"submittedAt"`
	// When the claim was created.
	CreatedAt *string `json:"createdAt"`
	// When the claim was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// One caregiver's hours and pay for one pay period. Read-only. Every field is always present; a
// field with no value is `null`.
type PayrollEntry struct {
	// String representing the object's type. Always `payroll_entry`.
	Object string `json:"object"`
	// Unique identifier of the payroll entry.
	ID string `json:"id"`
	// The caregiver paid.
	CaregiverID *string `json:"caregiverId"`
	// First day of the pay period.
	PayPeriodStart *string `json:"payPeriodStart"`
	// Last day of the pay period.
	PayPeriodEnd *string `json:"payPeriodEnd"`
	// Where the entry is in your payroll process.
	Status *PayrollEntryStatus `json:"status"`
	// Regular hours.
	RegularHours *float64 `json:"regularHours"`
	// Overtime hours.
	OvertimeHours *float64 `json:"overtimeHours"`
	// Holiday hours.
	HolidayHours *float64 `json:"holidayHours"`
	// Paid time off hours.
	PtoHours *float64 `json:"ptoHours"`
	// Sick leave hours.
	SickLeaveHours *float64 `json:"sickLeaveHours"`
	// Gross pay. In cents.
	GrossPayCents *int64 `json:"grossPayCents"`
	// Total deductions. In cents.
	DeductionsCents *int64 `json:"deductionsCents"`
	// Net pay. In cents.
	NetPayCents *int64 `json:"netPayCents"`
	// Mileage reimbursement. In cents.
	MileageCents *int64 `json:"mileageCents"`
	// Visits in the pay period.
	VisitCount *int64 `json:"visitCount"`
	// Visits paid on this entry.
	VisitIDs []string `json:"visitIds"`
	// How the caregiver is paid.
	PaymentMethod *PayrollEntryPaymentMethod `json:"paymentMethod"`
	// Who approved the entry.
	ApprovedByName *string `json:"approvedByName"`
	// When it was approved.
	ApprovedAt *string `json:"approvedAt"`
	// When it was paid.
	PaidAt *string `json:"paidAt"`
	// Notes.
	Notes *string `json:"notes"`
	// When the payroll entry was created.
	CreatedAt *string `json:"createdAt"`
	// When the payroll entry was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// In-app message threads between people at your agency: one-to-one chats and named group chats.
// Read-only; start conversations in AveeCare. Every field is always present; a field with no value
// is `null`.
type Conversation struct {
	// String representing the object's type. Always `conversation`.
	Object string `json:"object"`
	// Unique identifier of the conversation.
	ID string `json:"id"`
	// User IDs of everyone in the conversation.
	ParticipantIDs []string `json:"participantIds"`
	// Whether this is a group chat.
	IsGroup *bool `json:"isGroup"`
	// Group chat name; `null` for one-to-one chats.
	Title *string `json:"title"`
	// When the most recent message was sent.
	LastMessageTime *string `json:"lastMessageTime"`
	// Preview of the most recent message.
	LastMessageContent *string `json:"lastMessageContent"`
	// User ID of whoever sent the most recent message.
	LastMessageSenderID *string `json:"lastMessageSenderId"`
	// When the conversation was created.
	CreatedAt *string `json:"createdAt"`
	// When the conversation was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Messages in a conversation. Messages you send through the API are sent as the API key's user.
// Every field is always present; a field with no value is `null`.
type Message struct {
	// String representing the object's type. Always `message`.
	Object string `json:"object"`
	// Unique identifier of the message.
	ID string `json:"id"`
	// The conversation the message belongs to.
	ConversationID *string `json:"conversationId"`
	// Message text.
	Content *string `json:"content"`
	// User ID of the sender. Set to the API key's user.
	SenderID *string `json:"senderId"`
	// When the message was sent.
	SentAt *string `json:"sentAt"`
	// When the message was last edited.
	EditedAt *string `json:"editedAt"`
	// Whether the sender deleted the message. Deleted messages keep their place in the thread.
	IsDeleted *bool `json:"isDeleted"`
	// When the message was created.
	CreatedAt *string `json:"createdAt"`
	// When the message was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new message. Any other field is refused with `unknown_field`.
type MessageCreate struct {
	// The conversation the message belongs to.
	ConversationID string `json:"conversationId"`
	// Message text.
	Content string `json:"content"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m MessageCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("conversationId", m.ConversationID)
	w.add("content", m.Content)
	return w.bytes()
}

// Alerts on your agency's alert board, such as a credential about to expire or a missed clock-in,
// and announcements staff post for caregivers. Every field is always present; a field with no
// value is `null`.
type Alert struct {
	// String representing the object's type. Always `alert`.
	Object string `json:"object"`
	// Unique identifier of the alert.
	ID string `json:"id"`
	// Headline.
	Name *string `json:"name"`
	// Details.
	Description *string `json:"description"`
	// How prominently the alert is shown.
	Type *AlertType `json:"type"`
	// Who sees the alert. `Caregiver`, `Patient`, `Visit` and `PatientCareTeam` need
	// `audienceTargetId`.
	Audience *AlertAudience `json:"audience"`
	// The caregiver, patient or visit the alert is for, when the audience names one.
	AudienceTargetID *string `json:"audienceTargetId"`
	// When the alert takes effect.
	Time *string `json:"time"`
	// When the alert should be reviewed or cleared.
	ReviewDate *string `json:"reviewDate"`
	// When the alert was created.
	CreatedAt *string `json:"createdAt"`
	// When the alert was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new alert. Any other field is refused with `unknown_field`.
type AlertCreate struct {
	// Headline.
	Name string `json:"name"`
	// Details.
	Description Nullable[string] `json:"description,omitempty"`
	// How prominently the alert is shown.
	Type Nullable[AlertType] `json:"type,omitempty"`
	// Who sees the alert. `Caregiver`, `Patient`, `Visit` and `PatientCareTeam` need
	// `audienceTargetId`.
	Audience Nullable[AlertAudience] `json:"audience,omitempty"`
	// The caregiver, patient or visit the alert is for, when the audience names one.
	AudienceTargetID Nullable[string] `json:"audienceTargetId,omitempty"`
	// When the alert takes effect.
	Time Nullable[string] `json:"time,omitempty"`
	// When the alert should be reviewed or cleared.
	ReviewDate Nullable[string] `json:"reviewDate,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AlertCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("name", m.Name)
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.Type.IsOmitted() {
		w.add("type", m.Type)
	}
	if !m.Audience.IsOmitted() {
		w.add("audience", m.Audience)
	}
	if !m.AudienceTargetID.IsOmitted() {
		w.add("audienceTargetId", m.AudienceTargetID)
	}
	if !m.Time.IsOmitted() {
		w.add("time", m.Time)
	}
	if !m.ReviewDate.IsOmitted() {
		w.add("reviewDate", m.ReviewDate)
	}
	return w.bytes()
}

// Fields to change on an alert. Send only the fields to change; any other field is refused with
// `unknown_field`. `time` is set when the alert is created and cannot be changed afterwards.
type AlertUpdate struct {
	// Headline.
	Name *string `json:"name,omitempty"`
	// Details.
	Description Nullable[string] `json:"description,omitempty"`
	// How prominently the alert is shown.
	Type Nullable[AlertType] `json:"type,omitempty"`
	// Who sees the alert. `Caregiver`, `Patient`, `Visit` and `PatientCareTeam` need
	// `audienceTargetId`.
	Audience Nullable[AlertAudience] `json:"audience,omitempty"`
	// The caregiver, patient or visit the alert is for, when the audience names one.
	AudienceTargetID Nullable[string] `json:"audienceTargetId,omitempty"`
	// When the alert should be reviewed or cleared.
	ReviewDate Nullable[string] `json:"reviewDate,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m AlertUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Description.IsOmitted() {
		w.add("description", m.Description)
	}
	if !m.Type.IsOmitted() {
		w.add("type", m.Type)
	}
	if !m.Audience.IsOmitted() {
		w.add("audience", m.Audience)
	}
	if !m.AudienceTargetID.IsOmitted() {
		w.add("audienceTargetId", m.AudienceTargetID)
	}
	if !m.ReviewDate.IsOmitted() {
		w.add("reviewDate", m.ReviewDate)
	}
	return w.bytes()
}

// Confirmation that an alert was deleted.
type AlertDeleted struct {
	// Always `alert`.
	Object string `json:"object"`
	// ID of the deleted alert.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// Your agency's offices (branches). Patients, caregivers and applicants can each belong to one
// office. There is always exactly one default office. Every field is always present; a field with
// no value is `null`.
type Office struct {
	// String representing the object's type. Always `office`.
	Object string `json:"object"`
	// Unique identifier of the office.
	ID string `json:"id"`
	// Office name.
	Name *string `json:"name"`
	// Street address.
	Address *string `json:"address"`
	// City.
	City *string `json:"city"`
	// State, as a two-letter code.
	State *string `json:"state"`
	// ZIP code.
	ZipCode *string `json:"zipCode"`
	// Main phone number.
	Phone *string `json:"phone"`
	// Whether this is the default office. Making another office the default clears this one.
	IsDefault *bool `json:"isDefault"`
	// When the office was created.
	CreatedAt *string `json:"createdAt"`
	// When the office was last changed.
	UpdatedAt *string `json:"updatedAt"`
}

// Fields for a new office. Any other field is refused with `unknown_field`.
type OfficeCreate struct {
	// Office name.
	Name string `json:"name"`
	// Street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Main phone number.
	Phone Nullable[string] `json:"phone,omitempty"`
	// Whether this is the default office. Making another office the default clears this one.
	IsDefault Nullable[bool] `json:"isDefault,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m OfficeCreate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	w.add("name", m.Name)
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.Phone.IsOmitted() {
		w.add("phone", m.Phone)
	}
	if !m.IsDefault.IsOmitted() {
		w.add("isDefault", m.IsDefault)
	}
	return w.bytes()
}

// Fields to change on an office. Send only the fields to change; any other field is refused with
// `unknown_field`.
type OfficeUpdate struct {
	// Office name.
	Name *string `json:"name,omitempty"`
	// Street address.
	Address Nullable[string] `json:"address,omitempty"`
	// City.
	City Nullable[string] `json:"city,omitempty"`
	// State, as a two-letter code.
	State Nullable[string] `json:"state,omitempty"`
	// ZIP code.
	ZipCode Nullable[string] `json:"zipCode,omitempty"`
	// Main phone number.
	Phone Nullable[string] `json:"phone,omitempty"`
	// Whether this is the default office. Making another office the default clears this one.
	IsDefault Nullable[bool] `json:"isDefault,omitempty"`
}

// MarshalJSON implements json.Marshaler. Optional fields that were not set are left out of the request.
func (m OfficeUpdate) MarshalJSON() ([]byte, error) {
	var w objectWriter
	if m.Name != nil {
		w.add("name", m.Name)
	}
	if !m.Address.IsOmitted() {
		w.add("address", m.Address)
	}
	if !m.City.IsOmitted() {
		w.add("city", m.City)
	}
	if !m.State.IsOmitted() {
		w.add("state", m.State)
	}
	if !m.ZipCode.IsOmitted() {
		w.add("zipCode", m.ZipCode)
	}
	if !m.Phone.IsOmitted() {
		w.add("phone", m.Phone)
	}
	if !m.IsDefault.IsOmitted() {
		w.add("isDefault", m.IsDefault)
	}
	return w.bytes()
}

// Confirmation that an office was deleted.
type OfficeDeleted struct {
	// Always `office`.
	Object string `json:"object"`
	// ID of the deleted office.
	ID string `json:"id"`
	// Always `true`.
	Deleted bool `json:"deleted"`
}

// The API key making the request.
type Me struct {
	// Always `me`.
	Object string `json:"object"`
	// The key's public identifier: the part of the key between `ak_` and the second underscore.
	ID string `json:"id"`
	// The name given to the key when it was created.
	Name *string `json:"name"`
	// What the key may do. A key without `write` can only read.
	Scopes []MeScope `json:"scopes"`
	// Role of the person the key belongs to: `RootUser` (agency administrator), `User` (office
	// staff), `Caregiver` or `Patient`. It decides which resources the key can reach.
	Role MeRole `json:"role"`
	// Email address of the person the key belongs to.
	Email string `json:"email"`
	// When the key stops working.
	ExpiresAt string `json:"expiresAt"`
	// The agency the key belongs to.
	Organization MeOrganization `json:"organization"`
}

// The agency the key belongs to.
type MeOrganization struct {
	// ID of the agency.
	ID string `json:"id"`
	// Agency name.
	Name *string `json:"name"`
}

// PatientListParams holds the query parameters for patients.list.
type PatientListParams struct {
	// Only return patients with this status.
	Status *string
	// Only return active (`true`) or inactive (`false`) patients.
	Active *bool
	// Only return prospects or clients.
	LifecycleStage *PatientLifecycleStage
	// Only return patients served by this office.
	OfficeID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *PatientListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	if p.Active != nil {
		v.Set("active", strconv.FormatBool(*p.Active))
	}
	if p.LifecycleStage != nil {
		v.Set("lifecycle_stage", string(*p.LifecycleStage))
	}
	if p.OfficeID != nil {
		v.Set("office_id", *p.OfficeID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ContactListParams holds the query parameters for contacts.list.
type ContactListParams struct {
	// Only return contacts of this patient.
	PatientID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ContactListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// AllergyListParams holds the query parameters for allergies.list.
type AllergyListParams struct {
	// Only return allergies of this patient.
	PatientID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *AllergyListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// MedicationListParams holds the query parameters for medications.list.
type MedicationListParams struct {
	// Only return medications of this patient.
	PatientID *string
	// Only return medications with this status.
	Status *MedicationStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *MedicationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// CareGoalListParams holds the query parameters for careGoals.list.
type CareGoalListParams struct {
	// Only return goals of this patient.
	PatientID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *CareGoalListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// PatientNoteListParams holds the query parameters for patientNotes.list.
type PatientNoteListParams struct {
	// Only return notes about this patient.
	PatientID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *PatientNoteListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// MedicationAdministrationListParams holds the query parameters for
// medicationAdministrations.list.
type MedicationAdministrationListParams struct {
	// Only return doses for this patient.
	PatientID *string
	// Only return doses of this medication.
	MedicationID *string
	// Only return doses due on this date.
	ScheduledDate *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *MedicationAdministrationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.MedicationID != nil {
		v.Set("medication_id", *p.MedicationID)
	}
	if p.ScheduledDate != nil {
		v.Set("scheduled_date", *p.ScheduledDate)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// OasisAssessmentListParams holds the query parameters for oasisAssessments.list.
type OasisAssessmentListParams struct {
	// Only return assessments of this patient.
	PatientID *string
	// Only return assessments with this status.
	Status *OasisAssessmentStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *OasisAssessmentListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// CaregiverListParams holds the query parameters for caregivers.list.
type CaregiverListParams struct {
	// Only return caregivers with this employment status.
	Status *CaregiverStatus
	// Only return caregivers who can (`true`) or cannot (`false`) be scheduled.
	Active *bool
	// Only return caregivers from this office.
	OfficeID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *CaregiverListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.Active != nil {
		v.Set("active", strconv.FormatBool(*p.Active))
	}
	if p.OfficeID != nil {
		v.Set("office_id", *p.OfficeID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// CaregiverNoteListParams holds the query parameters for caregiverNotes.list.
type CaregiverNoteListParams struct {
	// Only return notes about this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *CaregiverNoteListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// CertificationListParams holds the query parameters for certifications.list.
type CertificationListParams struct {
	// Only return certifications of this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *CertificationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// AvailabilityListParams holds the query parameters for availability.list.
type AvailabilityListParams struct {
	// Only return windows of this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *AvailabilityListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// SpecializationListParams holds the query parameters for specializations.list.
type SpecializationListParams struct {
	// Only return specializations of this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *SpecializationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ApplicantListParams holds the query parameters for applicants.list.
type ApplicantListParams struct {
	// Only return applicants in this hiring stage.
	Stage *string
	// Only return applicants to this office.
	OfficeID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ApplicantListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Stage != nil {
		v.Set("stage", *p.Stage)
	}
	if p.OfficeID != nil {
		v.Set("office_id", *p.OfficeID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// VisitListParams holds the query parameters for visits.list. Fields that are not pointers are
// required.
type VisitListParams struct {
	// Start of the window (inclusive), ISO 8601 with a time zone. Visits are listed by scheduled
	// start. Both `start_from` and `start_to` are required, and the window can be at most 93 days.
	StartFrom string
	// End of the window (exclusive), ISO 8601 with a time zone. At most 93 days after `start_from`.
	StartTo string
	// Only return visits for this patient.
	PatientID *string
	// Only return visits assigned to this caregiver.
	CaregiverID *string
	// Only return visits with this status.
	Status *VisitStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *VisitListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	v.Set("start_from", p.StartFrom)
	v.Set("start_to", p.StartTo)
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// VisitNoteListParams holds the query parameters for visitNotes.list.
type VisitNoteListParams struct {
	// Only return notes about this visit.
	VisitID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *VisitNoteListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.VisitID != nil {
		v.Set("visit_id", *p.VisitID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// VisitRequestListParams holds the query parameters for visitRequests.list.
type VisitRequestListParams struct {
	// Only return requests from this caregiver.
	CaregiverID *string
	// Only return requests with this status.
	Status *VisitRequestStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *VisitRequestListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// OpenShiftListParams holds the query parameters for openShifts.list.
type OpenShiftListParams struct {
	// Only return offers with this status.
	Status *OpenShiftStatus
	// Only return offers for this visit.
	VisitID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *OpenShiftListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.VisitID != nil {
		v.Set("visit_id", *p.VisitID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ShiftSwapListParams holds the query parameters for shiftSwaps.list.
type ShiftSwapListParams struct {
	// Only return swaps with this status.
	Status *ShiftSwapStatus
	// Only return swaps for this visit.
	VisitID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ShiftSwapListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.VisitID != nil {
		v.Set("visit_id", *p.VisitID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// CarePlanListParams holds the query parameters for carePlans.list.
type CarePlanListParams struct {
	// Only return plans for this patient.
	PatientID *string
	// Only return plans with this status.
	Status *CarePlanStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *CarePlanListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ProgressNoteListParams holds the query parameters for progressNotes.list.
type ProgressNoteListParams struct {
	// Only return notes for this patient.
	PatientID *string
	// Only return notes for this visit.
	VisitID *string
	// Only return notes by this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ProgressNoteListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.VisitID != nil {
		v.Set("visit_id", *p.VisitID)
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// AuthorizationListParams holds the query parameters for authorizations.list.
type AuthorizationListParams struct {
	// Only return authorizations for this patient.
	PatientID *string
	// Only return authorizations with this status.
	Status *AuthorizationStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *AuthorizationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// IncidentListParams holds the query parameters for incidents.list.
type IncidentListParams struct {
	// Only return incidents with this follow-up status.
	Status *IncidentStatus
	// Only return incidents involving this patient.
	PatientID *string
	// Only return incidents involving this caregiver.
	CaregiverID *string
	// Only return resolved (`true`) or unresolved (`false`) incidents.
	Resolved *bool
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *IncidentListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.Resolved != nil {
		v.Set("resolved", strconv.FormatBool(*p.Resolved))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// EventListParams holds the query parameters for events.list.
type EventListParams struct {
	// Only return events with this follow-up status.
	Status *IncidentStatus
	// Only return events involving this patient.
	PatientID *string
	// Only return events involving this caregiver.
	CaregiverID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *EventListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ReferralSourceListParams holds the query parameters for referralSources.list.
type ReferralSourceListParams struct {
	// Only return active (`true`) or inactive (`false`) sources.
	IsActive *bool
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ReferralSourceListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.IsActive != nil {
		v.Set("is_active", strconv.FormatBool(*p.IsActive))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ReferralListParams holds the query parameters for referrals.list.
type ReferralListParams struct {
	// Only return referrals with this status.
	Status *ReferralStatus
	// Only return referrals from this source.
	ReferralSourceID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ReferralListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.ReferralSourceID != nil {
		v.Set("referral_source_id", *p.ReferralSourceID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// InvoiceListParams holds the query parameters for invoices.list.
type InvoiceListParams struct {
	// Only return invoices for this patient.
	PatientID *string
	// Only return invoices with this status.
	Status *InvoiceStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *InvoiceListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// InvoiceLineItemListParams holds the query parameters for invoiceLineItems.list.
type InvoiceLineItemListParams struct {
	// Only return lines of this invoice.
	InvoiceID *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *InvoiceLineItemListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.InvoiceID != nil {
		v.Set("invoice_id", *p.InvoiceID)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ClaimListParams holds the query parameters for claims.list.
type ClaimListParams struct {
	// Only return claims for this patient.
	PatientID *string
	// Only return claims with this status.
	Status *ClaimStatus
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ClaimListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.PatientID != nil {
		v.Set("patient_id", *p.PatientID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// PayrollEntryListParams holds the query parameters for payrollEntries.list.
type PayrollEntryListParams struct {
	// Only return entries for this caregiver.
	CaregiverID *string
	// Only return entries with this status.
	Status *PayrollEntryStatus
	// Only return entries for the pay period starting on this date.
	PayPeriodStart *string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *PayrollEntryListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CaregiverID != nil {
		v.Set("caregiver_id", *p.CaregiverID)
	}
	if p.Status != nil {
		v.Set("status", string(*p.Status))
	}
	if p.PayPeriodStart != nil {
		v.Set("pay_period_start", *p.PayPeriodStart)
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// ConversationListParams holds the query parameters for conversations.list.
type ConversationListParams struct {
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *ConversationListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// MessageListParams holds the query parameters for messages.list. Fields that are not pointers are
// required.
type MessageListParams struct {
	// The conversation to list messages from.
	ConversationID string
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *MessageListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	v.Set("conversation_id", p.ConversationID)
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// AlertListParams holds the query parameters for alerts.list.
type AlertListParams struct {
	// Only return alerts of this type.
	Type *AlertType
	// Only return alerts for this audience.
	Audience *AlertAudience
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *AlertListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Type != nil {
		v.Set("type", string(*p.Type))
	}
	if p.Audience != nil {
		v.Set("audience", string(*p.Audience))
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// OfficeListParams holds the query parameters for offices.list.
type OfficeListParams struct {
	// Only return objects created after this time (exclusive). ISO 8601 with a time zone.
	CreatedAfter *string
	// Only return objects created before this time (exclusive). ISO 8601 with a time zone.
	CreatedBefore *string
	// Only return objects last changed after this time (exclusive). ISO 8601 with a time zone. Use it
	// for incremental sync: pass the newest `updatedAt` you already have.
	UpdatedAfter *string
	// Only return objects last changed before this time (exclusive). ISO 8601 with a time zone.
	UpdatedBefore *string
	// Number of objects to return, from 1 to 100.
	Limit *int64
	// The `next_cursor` from the previous page. Send the same filters with it.
	Cursor *string
}

func (p *OfficeListParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.CreatedAfter != nil {
		v.Set("created_after", *p.CreatedAfter)
	}
	if p.CreatedBefore != nil {
		v.Set("created_before", *p.CreatedBefore)
	}
	if p.UpdatedAfter != nil {
		v.Set("updated_after", *p.UpdatedAfter)
	}
	if p.UpdatedBefore != nil {
		v.Set("updated_before", *p.UpdatedBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}
