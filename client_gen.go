// Code generated from the AveeCare OpenAPI specification. DO NOT EDIT.

package aveecare

// Client is the AveeCare API client. Create one with NewClient; it is safe for concurrent use.
// Each field is one API resource.
type Client struct {
	// Patients are the clients your agency provides care to, including prospective clients you have
	// not yet admitted.
	Patients *PatientsService
	// Contacts are family members and other people connected to a patient, such as an emergency
	// contact.
	Contacts *ContactsService
	// Allergies recorded on a patient's profile.
	Allergies *AllergiesService
	// Medications on a patient's medication list, with the schedule caregivers follow when they give
	// or assist with them.
	Medications *MedicationsService
	// Goals of care set for a patient, such as walking to the mailbox unaided.
	CareGoals *CareGoalsService
	// Notes your staff and caregivers keep on a patient's profile.
	PatientNotes *PatientNotesService
	// The medication administration record (MAR): each dose a caregiver gave, or the reason it was
	// not given. Read-only.
	MedicationAdministrations *MedicationAdministrationsService
	// OASIS assessments recorded for Medicare home health patients. This lists each assessment and
	// its status; the item-level answers are not included. Read-only.
	OasisAssessments *OasisAssessmentsService
	// Caregivers are the aides, nurses and therapists your agency schedules to visit patients.
	Caregivers *CaregiversService
	// Notes your staff keep on a caregiver's profile.
	CaregiverNotes *CaregiverNotesService
	// Certifications, licenses and background checks on a caregiver's profile, with their expiry
	// dates.
	Certifications *CertificationsService
	// When a caregiver can or cannot work. A window is either a weekly recurring time on one weekday,
	// or a one-off date range such as time off.
	Availability *AvailabilityService
	// Skills and specialties on a caregiver's profile, such as dementia care or Hoyer lift transfers.
	Specializations *SpecializationsService
	// People who applied to work for your agency, tracked through your hiring stages until they are
	// hired as caregivers.
	Applicants *ApplicantsService
	// Visits are scheduled or completed caregiver shifts at a patient's home. A visit's patient,
	// caregiver, times and status are set when it is created; reschedule or reassign a visit in
	// AveeCare.
	Visits *VisitsService
	// Notes written about a visit by the caregiver or office staff.
	VisitNotes *VisitNotesService
	// Requests from caregivers to pick up visits with a patient, which office staff approve or deny.
	VisitRequests *VisitRequestsService
	// Open shifts are unassigned visits offered to caregivers so one of them can claim it.
	OpenShifts *OpenShiftsService
	// Requests from one caregiver to hand a visit to another, which the other caregiver and then the
	// office approve.
	ShiftSwaps *ShiftSwapsService
	// Care plans describe the services a patient receives, how often, and when the plan is next
	// reviewed. The plan's goals and task lists are managed in AveeCare and are not included here.
	CarePlans *CarePlansService
	// Progress notes are the clinical summaries caregivers write after a visit.
	ProgressNotes *ProgressNotesService
	// Payer authorizations: how many units of a service a payer has approved for a patient, and over
	// what dates.
	Authorizations *AuthorizationsService
	// Incident reports: falls, injuries, medication errors and other events that need follow-up,
	// often for regulatory reporting.
	Incidents *IncidentsService
	// Events your agency documents about a patient, caregiver or visit, such as a complaint, a
	// compliment or an attendance issue.
	Events *EventsService
	// Hospitals, physicians, case managers and other organizations that refer clients to your agency.
	ReferralSources *ReferralSourcesService
	// Individual referrals of a prospective client, tracked from first contact to admission.
	Referrals *ReferralsService
	// Invoices billed to patients and their families for private-pay care. Read-only.
	Invoices *InvoicesService
	// The individual charges on an invoice. Read-only.
	InvoiceLineItems *InvoiceLineItemsService
	// Insurance claims billed to Medicaid, Medicare and other payers, with their adjudication status.
	// Read-only.
	Claims *ClaimsService
	// One caregiver's hours and pay for one pay period. Read-only.
	PayrollEntries *PayrollEntriesService
	// In-app message threads between people at your agency: one-to-one chats and named group chats.
	// Read-only; start conversations in AveeCare.
	Conversations *ConversationsService
	// Messages in a conversation. Messages you send through the API are sent as the API key's user.
	Messages *MessagesService
	// Alerts on your agency's alert board, such as a credential about to expire or a missed clock-in,
	// and announcements staff post for caregivers.
	Alerts *AlertsService
	// Your agency's offices (branches). Patients, caregivers and applicants can each belong to one
	// office. There is always exactly one default office.
	Offices *OfficesService
	// The API key making the request: its scopes, its user's role and its agency.
	Me *MeService

	core *core
}

func newClient(c *core) *Client {
	return &Client{
		Patients:                  &PatientsService{c: c},
		Contacts:                  &ContactsService{c: c},
		Allergies:                 &AllergiesService{c: c},
		Medications:               &MedicationsService{c: c},
		CareGoals:                 &CareGoalsService{c: c},
		PatientNotes:              &PatientNotesService{c: c},
		MedicationAdministrations: &MedicationAdministrationsService{c: c},
		OasisAssessments:          &OasisAssessmentsService{c: c},
		Caregivers:                &CaregiversService{c: c},
		CaregiverNotes:            &CaregiverNotesService{c: c},
		Certifications:            &CertificationsService{c: c},
		Availability:              &AvailabilityService{c: c},
		Specializations:           &SpecializationsService{c: c},
		Applicants:                &ApplicantsService{c: c},
		Visits:                    &VisitsService{c: c},
		VisitNotes:                &VisitNotesService{c: c},
		VisitRequests:             &VisitRequestsService{c: c},
		OpenShifts:                &OpenShiftsService{c: c},
		ShiftSwaps:                &ShiftSwapsService{c: c},
		CarePlans:                 &CarePlansService{c: c},
		ProgressNotes:             &ProgressNotesService{c: c},
		Authorizations:            &AuthorizationsService{c: c},
		Incidents:                 &IncidentsService{c: c},
		Events:                    &EventsService{c: c},
		ReferralSources:           &ReferralSourcesService{c: c},
		Referrals:                 &ReferralsService{c: c},
		Invoices:                  &InvoicesService{c: c},
		InvoiceLineItems:          &InvoiceLineItemsService{c: c},
		Claims:                    &ClaimsService{c: c},
		PayrollEntries:            &PayrollEntriesService{c: c},
		Conversations:             &ConversationsService{c: c},
		Messages:                  &MessagesService{c: c},
		Alerts:                    &AlertsService{c: c},
		Offices:                   &OfficesService{c: c},
		Me:                        &MeService{c: c},
		core:                      c,
	}
}
