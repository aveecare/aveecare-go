// Code generated from the AveeCare OpenAPI specification. DO NOT EDIT.

package aveecare

import "context"

// PatientsService groups the patients operations. Patients are the clients your agency provides
// care to, including prospective clients you have not yet admitted.
type PatientsService struct {
	c *core
}

// List patients.
//
// Returns a page of patients. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *PatientsService) List(ctx context.Context, params *PatientListParams, opts ...RequestOption) (*Page[Patient], error) {
	return getPage[Patient](ctx, s.c, "/v1/patients", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *PatientsService) ListAutoPaging(ctx context.Context, params *PatientListParams, opts ...RequestOption) *Iterator[Patient] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Patient], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a patient.
//
// Creates a patient and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *PatientsService) Create(ctx context.Context, params *PatientCreate, opts ...RequestOption) (*Patient, error) {
	if params == nil {
		params = &PatientCreate{}
	}
	var body any = params
	var out Patient
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/patients", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a patient.
//
// Returns one patient by ID.
//
// Requires the read scope.
func (s *PatientsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Patient, error) {
	var out Patient
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/patients/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a patient.
//
// Changes the fields you send and returns the updated patient. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *PatientsService) Update(ctx context.Context, id string, params *PatientUpdate, opts ...RequestOption) (*Patient, error) {
	if params == nil {
		params = &PatientUpdate{}
	}
	var body any = params
	var out Patient
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/patients/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a patient.
//
// Deletes a patient. The patient's visits, notes and other related records are not deleted.
// Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *PatientsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*PatientDeleted, error) {
	var out PatientDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/patients/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ContactsService groups the contacts operations. Contacts are family members and other people
// connected to a patient, such as an emergency contact.
type ContactsService struct {
	c *core
}

// List contacts.
//
// Returns a page of contacts. Filters combine with AND. Use `next_cursor` to fetch the next page.
//
// Requires the read scope.
func (s *ContactsService) List(ctx context.Context, params *ContactListParams, opts ...RequestOption) (*Page[Contact], error) {
	return getPage[Contact](ctx, s.c, "/v1/contacts", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ContactsService) ListAutoPaging(ctx context.Context, params *ContactListParams, opts ...RequestOption) *Iterator[Contact] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Contact], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a contact.
//
// Creates a contact and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ContactsService) Create(ctx context.Context, params *ContactCreate, opts ...RequestOption) (*Contact, error) {
	if params == nil {
		params = &ContactCreate{}
	}
	var body any = params
	var out Contact
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/contacts", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a contact.
//
// Returns one contact by ID.
//
// Requires the read scope.
func (s *ContactsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Contact, error) {
	var out Contact
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/contacts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a contact.
//
// Changes the fields you send and returns the updated contact. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ContactsService) Update(ctx context.Context, id string, params *ContactUpdate, opts ...RequestOption) (*Contact, error) {
	if params == nil {
		params = &ContactUpdate{}
	}
	var body any = params
	var out Contact
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/contacts/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a contact.
//
// Deletes a contact. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *ContactsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ContactDeleted, error) {
	var out ContactDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/contacts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AllergiesService groups the allergies operations. Allergies recorded on a patient's profile.
type AllergiesService struct {
	c *core
}

// List allergies.
//
// Returns a page of allergies. Filters combine with AND. Use `next_cursor` to fetch the next page.
//
// Requires the read scope.
func (s *AllergiesService) List(ctx context.Context, params *AllergyListParams, opts ...RequestOption) (*Page[Allergy], error) {
	return getPage[Allergy](ctx, s.c, "/v1/allergies", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *AllergiesService) ListAutoPaging(ctx context.Context, params *AllergyListParams, opts ...RequestOption) *Iterator[Allergy] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Allergy], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an allergy.
//
// Creates an allergy and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AllergiesService) Create(ctx context.Context, params *AllergyCreate, opts ...RequestOption) (*Allergy, error) {
	if params == nil {
		params = &AllergyCreate{}
	}
	var body any = params
	var out Allergy
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/allergies", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an allergy.
//
// Returns one allergy by ID.
//
// Requires the read scope.
func (s *AllergiesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Allergy, error) {
	var out Allergy
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/allergies/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an allergy.
//
// Changes the fields you send and returns the updated allergy. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AllergiesService) Update(ctx context.Context, id string, params *AllergyUpdate, opts ...RequestOption) (*Allergy, error) {
	if params == nil {
		params = &AllergyUpdate{}
	}
	var body any = params
	var out Allergy
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/allergies/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an allergy.
//
// Deletes an allergy. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *AllergiesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*AllergyDeleted, error) {
	var out AllergyDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/allergies/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// MedicationsService groups the medications operations. Medications on a patient's medication
// list, with the schedule caregivers follow when they give or assist with them.
type MedicationsService struct {
	c *core
}

// List medications.
//
// Returns a page of medications. Filters combine with AND. Use `next_cursor` to fetch the next
// page.
//
// Requires the read scope.
func (s *MedicationsService) List(ctx context.Context, params *MedicationListParams, opts ...RequestOption) (*Page[Medication], error) {
	return getPage[Medication](ctx, s.c, "/v1/medications", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *MedicationsService) ListAutoPaging(ctx context.Context, params *MedicationListParams, opts ...RequestOption) *Iterator[Medication] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Medication], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a medication.
//
// Creates a medication and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *MedicationsService) Create(ctx context.Context, params *MedicationCreate, opts ...RequestOption) (*Medication, error) {
	if params == nil {
		params = &MedicationCreate{}
	}
	var body any = params
	var out Medication
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/medications", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a medication.
//
// Returns one medication by ID.
//
// Requires the read scope.
func (s *MedicationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Medication, error) {
	var out Medication
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/medications/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a medication.
//
// Changes the fields you send and returns the updated medication. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *MedicationsService) Update(ctx context.Context, id string, params *MedicationUpdate, opts ...RequestOption) (*Medication, error) {
	if params == nil {
		params = &MedicationUpdate{}
	}
	var body any = params
	var out Medication
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/medications/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a medication.
//
// Deletes a medication. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *MedicationsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*MedicationDeleted, error) {
	var out MedicationDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/medications/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CareGoalsService groups the careGoals operations. Goals of care set for a patient, such as
// walking to the mailbox unaided.
type CareGoalsService struct {
	c *core
}

// List care goals.
//
// Returns a page of care goals. Filters combine with AND. Use `next_cursor` to fetch the next
// page.
//
// Requires the read scope.
func (s *CareGoalsService) List(ctx context.Context, params *CareGoalListParams, opts ...RequestOption) (*Page[CareGoal], error) {
	return getPage[CareGoal](ctx, s.c, "/v1/care-goals", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *CareGoalsService) ListAutoPaging(ctx context.Context, params *CareGoalListParams, opts ...RequestOption) *Iterator[CareGoal] {
	return newIterator(ctx, func(ctx context.Context) (*Page[CareGoal], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a care goal.
//
// Creates a care goal and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CareGoalsService) Create(ctx context.Context, params *CareGoalCreate, opts ...RequestOption) (*CareGoal, error) {
	if params == nil {
		params = &CareGoalCreate{}
	}
	var body any = params
	var out CareGoal
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/care-goals", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a care goal.
//
// Returns one care goal by ID.
//
// Requires the read scope.
func (s *CareGoalsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*CareGoal, error) {
	var out CareGoal
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/care-goals/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a care goal.
//
// Changes the fields you send and returns the updated care goal. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CareGoalsService) Update(ctx context.Context, id string, params *CareGoalUpdate, opts ...RequestOption) (*CareGoal, error) {
	if params == nil {
		params = &CareGoalUpdate{}
	}
	var body any = params
	var out CareGoal
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/care-goals/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a care goal.
//
// Deletes a care goal. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *CareGoalsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*CareGoalDeleted, error) {
	var out CareGoalDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/care-goals/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatientNotesService groups the patientNotes operations. Notes your staff and caregivers keep on
// a patient's profile.
type PatientNotesService struct {
	c *core
}

// List patient notes.
//
// Returns a page of patient notes. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *PatientNotesService) List(ctx context.Context, params *PatientNoteListParams, opts ...RequestOption) (*Page[PatientNote], error) {
	return getPage[PatientNote](ctx, s.c, "/v1/patient-notes", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *PatientNotesService) ListAutoPaging(ctx context.Context, params *PatientNoteListParams, opts ...RequestOption) *Iterator[PatientNote] {
	return newIterator(ctx, func(ctx context.Context) (*Page[PatientNote], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a patient note.
//
// Creates a patient note and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *PatientNotesService) Create(ctx context.Context, params *PatientNoteCreate, opts ...RequestOption) (*PatientNote, error) {
	if params == nil {
		params = &PatientNoteCreate{}
	}
	var body any = params
	var out PatientNote
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/patient-notes", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a patient note.
//
// Returns one patient note by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *PatientNotesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*PatientNote, error) {
	var out PatientNote
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/patient-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a patient note.
//
// Changes the fields you send and returns the updated patient note. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *PatientNotesService) Update(ctx context.Context, id string, params *PatientNoteUpdate, opts ...RequestOption) (*PatientNote, error) {
	if params == nil {
		params = &PatientNoteUpdate{}
	}
	var body any = params
	var out PatientNote
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/patient-notes/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a patient note.
//
// Deletes a patient note. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *PatientNotesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*PatientNoteDeleted, error) {
	var out PatientNoteDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/patient-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// MedicationAdministrationsService groups the medicationAdministrations operations. The medication
// administration record (MAR): each dose a caregiver gave, or the reason it was not given.
// Read-only.
type MedicationAdministrationsService struct {
	c *core
}

// List medication administrations.
//
// Returns a page of medication administrations. Filters combine with AND. Use `next_cursor` to
// fetch the next page.
//
// Requires the read scope.
func (s *MedicationAdministrationsService) List(ctx context.Context, params *MedicationAdministrationListParams, opts ...RequestOption) (*Page[MedicationAdministration], error) {
	return getPage[MedicationAdministration](ctx, s.c, "/v1/medication-administrations", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *MedicationAdministrationsService) ListAutoPaging(ctx context.Context, params *MedicationAdministrationListParams, opts ...RequestOption) *Iterator[MedicationAdministration] {
	return newIterator(ctx, func(ctx context.Context) (*Page[MedicationAdministration], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve a medication administration.
//
// Returns one medication administration by ID.
//
// Requires the read scope.
func (s *MedicationAdministrationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*MedicationAdministration, error) {
	var out MedicationAdministration
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/medication-administrations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OasisAssessmentsService groups the oasisAssessments operations. OASIS assessments recorded for
// Medicare home health patients. This lists each assessment and its status; the item-level answers
// are not included. Read-only.
type OasisAssessmentsService struct {
	c *core
}

// List OASIS assessments.
//
// Returns a page of OASIS assessments. Filters combine with AND. Use `next_cursor` to fetch the
// next page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *OasisAssessmentsService) List(ctx context.Context, params *OasisAssessmentListParams, opts ...RequestOption) (*Page[OasisAssessment], error) {
	return getPage[OasisAssessment](ctx, s.c, "/v1/oasis-assessments", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *OasisAssessmentsService) ListAutoPaging(ctx context.Context, params *OasisAssessmentListParams, opts ...RequestOption) *Iterator[OasisAssessment] {
	return newIterator(ctx, func(ctx context.Context) (*Page[OasisAssessment], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve an OASIS assessment.
//
// Returns one OASIS assessment by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *OasisAssessmentsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*OasisAssessment, error) {
	var out OasisAssessment
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/oasis-assessments/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CaregiversService groups the caregivers operations. Caregivers are the aides, nurses and
// therapists your agency schedules to visit patients.
type CaregiversService struct {
	c *core
}

// List caregivers.
//
// Returns a page of caregivers. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CaregiversService) List(ctx context.Context, params *CaregiverListParams, opts ...RequestOption) (*Page[Caregiver], error) {
	return getPage[Caregiver](ctx, s.c, "/v1/caregivers", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *CaregiversService) ListAutoPaging(ctx context.Context, params *CaregiverListParams, opts ...RequestOption) *Iterator[Caregiver] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Caregiver], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a caregiver.
//
// Creates a caregiver and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CaregiversService) Create(ctx context.Context, params *CaregiverCreate, opts ...RequestOption) (*Caregiver, error) {
	if params == nil {
		params = &CaregiverCreate{}
	}
	var body any = params
	var out Caregiver
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/caregivers", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a caregiver.
//
// Returns one caregiver by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CaregiversService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Caregiver, error) {
	var out Caregiver
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/caregivers/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a caregiver.
//
// Changes the fields you send and returns the updated caregiver. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CaregiversService) Update(ctx context.Context, id string, params *CaregiverUpdate, opts ...RequestOption) (*Caregiver, error) {
	if params == nil {
		params = &CaregiverUpdate{}
	}
	var body any = params
	var out Caregiver
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/caregivers/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a caregiver.
//
// Deletes a caregiver. The caregiver's past visits and other related records are not deleted.
// Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CaregiversService) Delete(ctx context.Context, id string, opts ...RequestOption) (*CaregiverDeleted, error) {
	var out CaregiverDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/caregivers/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CaregiverNotesService groups the caregiverNotes operations. Notes your staff keep on a
// caregiver's profile.
type CaregiverNotesService struct {
	c *core
}

// List caregiver notes.
//
// Returns a page of caregiver notes. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CaregiverNotesService) List(ctx context.Context, params *CaregiverNoteListParams, opts ...RequestOption) (*Page[CaregiverNote], error) {
	return getPage[CaregiverNote](ctx, s.c, "/v1/caregiver-notes", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *CaregiverNotesService) ListAutoPaging(ctx context.Context, params *CaregiverNoteListParams, opts ...RequestOption) *Iterator[CaregiverNote] {
	return newIterator(ctx, func(ctx context.Context) (*Page[CaregiverNote], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a caregiver note.
//
// Creates a caregiver note and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CaregiverNotesService) Create(ctx context.Context, params *CaregiverNoteCreate, opts ...RequestOption) (*CaregiverNote, error) {
	if params == nil {
		params = &CaregiverNoteCreate{}
	}
	var body any = params
	var out CaregiverNote
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/caregiver-notes", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a caregiver note.
//
// Returns one caregiver note by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CaregiverNotesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*CaregiverNote, error) {
	var out CaregiverNote
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/caregiver-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a caregiver note.
//
// Changes the fields you send and returns the updated caregiver note. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CaregiverNotesService) Update(ctx context.Context, id string, params *CaregiverNoteUpdate, opts ...RequestOption) (*CaregiverNote, error) {
	if params == nil {
		params = &CaregiverNoteUpdate{}
	}
	var body any = params
	var out CaregiverNote
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/caregiver-notes/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a caregiver note.
//
// Deletes a caregiver note. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *CaregiverNotesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*CaregiverNoteDeleted, error) {
	var out CaregiverNoteDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/caregiver-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CertificationsService groups the certifications operations. Certifications, licenses and
// background checks on a caregiver's profile, with their expiry dates.
type CertificationsService struct {
	c *core
}

// List certifications.
//
// Returns a page of certifications. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CertificationsService) List(ctx context.Context, params *CertificationListParams, opts ...RequestOption) (*Page[Certification], error) {
	return getPage[Certification](ctx, s.c, "/v1/certifications", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *CertificationsService) ListAutoPaging(ctx context.Context, params *CertificationListParams, opts ...RequestOption) *Iterator[Certification] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Certification], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a certification.
//
// Creates a certification and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CertificationsService) Create(ctx context.Context, params *CertificationCreate, opts ...RequestOption) (*Certification, error) {
	if params == nil {
		params = &CertificationCreate{}
	}
	var body any = params
	var out Certification
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/certifications", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a certification.
//
// Returns one certification by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CertificationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Certification, error) {
	var out Certification
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/certifications/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a certification.
//
// Changes the fields you send and returns the updated certification. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CertificationsService) Update(ctx context.Context, id string, params *CertificationUpdate, opts ...RequestOption) (*Certification, error) {
	if params == nil {
		params = &CertificationUpdate{}
	}
	var body any = params
	var out Certification
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/certifications/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a certification.
//
// Deletes a certification. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *CertificationsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*CertificationDeleted, error) {
	var out CertificationDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/certifications/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AvailabilityService groups the availability operations. When a caregiver can or cannot work. A
// window is either a weekly recurring time on one weekday, or a one-off date range such as time
// off.
type AvailabilityService struct {
	c *core
}

// List availability windows.
//
// Returns a page of availability windows. Filters combine with AND. Use `next_cursor` to fetch the
// next page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *AvailabilityService) List(ctx context.Context, params *AvailabilityListParams, opts ...RequestOption) (*Page[Availability], error) {
	return getPage[Availability](ctx, s.c, "/v1/availability", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *AvailabilityService) ListAutoPaging(ctx context.Context, params *AvailabilityListParams, opts ...RequestOption) *Iterator[Availability] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Availability], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an availability window.
//
// Creates an availability window and returns it. Send an `Idempotency-Key` header to make the
// request safe to retry. Available to keys belonging to office staff and caregivers. With a
// caregiver's or patient's key, a saved record that key cannot read back is answered with only
// `object` and `id`. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AvailabilityService) Create(ctx context.Context, params *AvailabilityCreate, opts ...RequestOption) (*Availability, error) {
	if params == nil {
		params = &AvailabilityCreate{}
	}
	var body any = params
	var out Availability
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/availability", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an availability window.
//
// Returns one availability window by ID. Available to keys belonging to office staff and
// caregivers.
//
// Requires the read scope.
func (s *AvailabilityService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Availability, error) {
	var out Availability
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/availability/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an availability window.
//
// Changes the fields you send and returns the updated availability window. Send `null` to clear an
// optional field. Available to keys belonging to office staff and caregivers. With a caregiver's
// or patient's key, a saved record that key cannot read back is answered with only `object` and
// `id`. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AvailabilityService) Update(ctx context.Context, id string, params *AvailabilityUpdate, opts ...RequestOption) (*Availability, error) {
	if params == nil {
		params = &AvailabilityUpdate{}
	}
	var body any = params
	var out Availability
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/availability/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an availability window.
//
// Deletes an availability window. Available to keys belonging to office staff and caregivers.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AvailabilityService) Delete(ctx context.Context, id string, opts ...RequestOption) (*AvailabilityDeleted, error) {
	var out AvailabilityDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/availability/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SpecializationsService groups the specializations operations. Skills and specialties on a
// caregiver's profile, such as dementia care or Hoyer lift transfers.
type SpecializationsService struct {
	c *core
}

// List specializations.
//
// Returns a page of specializations. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *SpecializationsService) List(ctx context.Context, params *SpecializationListParams, opts ...RequestOption) (*Page[Specialization], error) {
	return getPage[Specialization](ctx, s.c, "/v1/specializations", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *SpecializationsService) ListAutoPaging(ctx context.Context, params *SpecializationListParams, opts ...RequestOption) *Iterator[Specialization] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Specialization], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a specialization.
//
// Creates a specialization and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *SpecializationsService) Create(ctx context.Context, params *SpecializationCreate, opts ...RequestOption) (*Specialization, error) {
	if params == nil {
		params = &SpecializationCreate{}
	}
	var body any = params
	var out Specialization
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/specializations", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a specialization.
//
// Returns one specialization by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *SpecializationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Specialization, error) {
	var out Specialization
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/specializations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a specialization.
//
// Changes the fields you send and returns the updated specialization. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *SpecializationsService) Update(ctx context.Context, id string, params *SpecializationUpdate, opts ...RequestOption) (*Specialization, error) {
	if params == nil {
		params = &SpecializationUpdate{}
	}
	var body any = params
	var out Specialization
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/specializations/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a specialization.
//
// Deletes a specialization. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *SpecializationsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*SpecializationDeleted, error) {
	var out SpecializationDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/specializations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ApplicantsService groups the applicants operations. People who applied to work for your agency,
// tracked through your hiring stages until they are hired as caregivers.
type ApplicantsService struct {
	c *core
}

// List applicants.
//
// Returns a page of applicants. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ApplicantsService) List(ctx context.Context, params *ApplicantListParams, opts ...RequestOption) (*Page[Applicant], error) {
	return getPage[Applicant](ctx, s.c, "/v1/applicants", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ApplicantsService) ListAutoPaging(ctx context.Context, params *ApplicantListParams, opts ...RequestOption) *Iterator[Applicant] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Applicant], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an applicant.
//
// Creates an applicant and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ApplicantsService) Create(ctx context.Context, params *ApplicantCreate, opts ...RequestOption) (*Applicant, error) {
	if params == nil {
		params = &ApplicantCreate{}
	}
	var body any = params
	var out Applicant
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/applicants", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an applicant.
//
// Returns one applicant by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ApplicantsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Applicant, error) {
	var out Applicant
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/applicants/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an applicant.
//
// Changes the fields you send and returns the updated applicant. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ApplicantsService) Update(ctx context.Context, id string, params *ApplicantUpdate, opts ...RequestOption) (*Applicant, error) {
	if params == nil {
		params = &ApplicantUpdate{}
	}
	var body any = params
	var out Applicant
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/applicants/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an applicant.
//
// Deletes an applicant. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *ApplicantsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ApplicantDeleted, error) {
	var out ApplicantDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/applicants/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// VisitsService groups the visits operations. Visits are scheduled or completed caregiver shifts
// at a patient's home. A visit's patient, caregiver, times and status are set when it is created;
// reschedule or reassign a visit in AveeCare.
type VisitsService struct {
	c *core
}

// List visits.
//
// Returns a page of visits, ordered by scheduled start, earliest first. Filters combine with AND.
// Use `next_cursor` to fetch the next page.
//
// Requires the read scope.
func (s *VisitsService) List(ctx context.Context, params *VisitListParams, opts ...RequestOption) (*Page[Visit], error) {
	return getPage[Visit](ctx, s.c, "/v1/visits", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *VisitsService) ListAutoPaging(ctx context.Context, params *VisitListParams, opts ...RequestOption) *Iterator[Visit] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Visit], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a visit.
//
// Creates a visit and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitsService) Create(ctx context.Context, params *VisitCreate, opts ...RequestOption) (*Visit, error) {
	if params == nil {
		params = &VisitCreate{}
	}
	var body any = params
	var out Visit
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/visits", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a visit.
//
// Returns one visit by ID.
//
// Requires the read scope.
func (s *VisitsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Visit, error) {
	var out Visit
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/visits/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a visit.
//
// Changes the fields you send and returns the updated visit. Send `null` to clear an optional
// field. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitsService) Update(ctx context.Context, id string, params *VisitUpdate, opts ...RequestOption) (*Visit, error) {
	if params == nil {
		params = &VisitUpdate{}
	}
	var body any = params
	var out Visit
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/visits/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a visit.
//
// Deletes a visit. The visit moves to the schedule's recycle bin in AveeCare, where staff can
// restore it. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*VisitDeleted, error) {
	var out VisitDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/visits/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// VisitNotesService groups the visitNotes operations. Notes written about a visit by the caregiver
// or office staff.
type VisitNotesService struct {
	c *core
}

// List visit notes.
//
// Returns a page of visit notes. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *VisitNotesService) List(ctx context.Context, params *VisitNoteListParams, opts ...RequestOption) (*Page[VisitNote], error) {
	return getPage[VisitNote](ctx, s.c, "/v1/visit-notes", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *VisitNotesService) ListAutoPaging(ctx context.Context, params *VisitNoteListParams, opts ...RequestOption) *Iterator[VisitNote] {
	return newIterator(ctx, func(ctx context.Context) (*Page[VisitNote], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a visit note.
//
// Creates a visit note and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitNotesService) Create(ctx context.Context, params *VisitNoteCreate, opts ...RequestOption) (*VisitNote, error) {
	if params == nil {
		params = &VisitNoteCreate{}
	}
	var body any = params
	var out VisitNote
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/visit-notes", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a visit note.
//
// Returns one visit note by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *VisitNotesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*VisitNote, error) {
	var out VisitNote
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/visit-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a visit note.
//
// Changes the fields you send and returns the updated visit note. Send `null` to clear an optional
// field. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitNotesService) Update(ctx context.Context, id string, params *VisitNoteUpdate, opts ...RequestOption) (*VisitNote, error) {
	if params == nil {
		params = &VisitNoteUpdate{}
	}
	var body any = params
	var out VisitNote
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/visit-notes/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a visit note.
//
// Deletes a visit note. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *VisitNotesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*VisitNoteDeleted, error) {
	var out VisitNoteDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/visit-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// VisitRequestsService groups the visitRequests operations. Requests from caregivers to pick up
// visits with a patient, which office staff approve or deny.
type VisitRequestsService struct {
	c *core
}

// List visit requests.
//
// Returns a page of visit requests. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *VisitRequestsService) List(ctx context.Context, params *VisitRequestListParams, opts ...RequestOption) (*Page[VisitRequest], error) {
	return getPage[VisitRequest](ctx, s.c, "/v1/visit-requests", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *VisitRequestsService) ListAutoPaging(ctx context.Context, params *VisitRequestListParams, opts ...RequestOption) *Iterator[VisitRequest] {
	return newIterator(ctx, func(ctx context.Context) (*Page[VisitRequest], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a visit request.
//
// Creates a visit request and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitRequestsService) Create(ctx context.Context, params *VisitRequestCreate, opts ...RequestOption) (*VisitRequest, error) {
	if params == nil {
		params = &VisitRequestCreate{}
	}
	var body any = params
	var out VisitRequest
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/visit-requests", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a visit request.
//
// Returns one visit request by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *VisitRequestsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*VisitRequest, error) {
	var out VisitRequest
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/visit-requests/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a visit request.
//
// Changes the fields you send and returns the updated visit request. Send `null` to clear an
// optional field. Available to keys belonging to office staff and caregivers. With a caregiver's
// or patient's key, a saved record that key cannot read back is answered with only `object` and
// `id`. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *VisitRequestsService) Update(ctx context.Context, id string, params *VisitRequestUpdate, opts ...RequestOption) (*VisitRequest, error) {
	if params == nil {
		params = &VisitRequestUpdate{}
	}
	var body any = params
	var out VisitRequest
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/visit-requests/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a visit request.
//
// Deletes a visit request. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *VisitRequestsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*VisitRequestDeleted, error) {
	var out VisitRequestDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/visit-requests/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OpenShiftsService groups the openShifts operations. Open shifts are unassigned visits offered to
// caregivers so one of them can claim it.
type OpenShiftsService struct {
	c *core
}

// List open shifts.
//
// Returns a page of open shifts. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *OpenShiftsService) List(ctx context.Context, params *OpenShiftListParams, opts ...RequestOption) (*Page[OpenShift], error) {
	return getPage[OpenShift](ctx, s.c, "/v1/open-shifts", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *OpenShiftsService) ListAutoPaging(ctx context.Context, params *OpenShiftListParams, opts ...RequestOption) *Iterator[OpenShift] {
	return newIterator(ctx, func(ctx context.Context) (*Page[OpenShift], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an open shift.
//
// Creates an open shift and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *OpenShiftsService) Create(ctx context.Context, params *OpenShiftCreate, opts ...RequestOption) (*OpenShift, error) {
	if params == nil {
		params = &OpenShiftCreate{}
	}
	var body any = params
	var out OpenShift
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/open-shifts", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an open shift.
//
// Returns one open shift by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *OpenShiftsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*OpenShift, error) {
	var out OpenShift
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/open-shifts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an open shift.
//
// Changes the fields you send and returns the updated open shift. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *OpenShiftsService) Update(ctx context.Context, id string, params *OpenShiftUpdate, opts ...RequestOption) (*OpenShift, error) {
	if params == nil {
		params = &OpenShiftUpdate{}
	}
	var body any = params
	var out OpenShift
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/open-shifts/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an open shift.
//
// Deletes an open shift. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *OpenShiftsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*OpenShiftDeleted, error) {
	var out OpenShiftDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/open-shifts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ShiftSwapsService groups the shiftSwaps operations. Requests from one caregiver to hand a visit
// to another, which the other caregiver and then the office approve.
type ShiftSwapsService struct {
	c *core
}

// List shift swaps.
//
// Returns a page of shift swaps. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *ShiftSwapsService) List(ctx context.Context, params *ShiftSwapListParams, opts ...RequestOption) (*Page[ShiftSwap], error) {
	return getPage[ShiftSwap](ctx, s.c, "/v1/shift-swaps", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ShiftSwapsService) ListAutoPaging(ctx context.Context, params *ShiftSwapListParams, opts ...RequestOption) *Iterator[ShiftSwap] {
	return newIterator(ctx, func(ctx context.Context) (*Page[ShiftSwap], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a shift swap.
//
// Creates a shift swap and returns it. Send an `Idempotency-Key` header to make the request safe
// to retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ShiftSwapsService) Create(ctx context.Context, params *ShiftSwapCreate, opts ...RequestOption) (*ShiftSwap, error) {
	if params == nil {
		params = &ShiftSwapCreate{}
	}
	var body any = params
	var out ShiftSwap
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/shift-swaps", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a shift swap.
//
// Returns one shift swap by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *ShiftSwapsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*ShiftSwap, error) {
	var out ShiftSwap
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/shift-swaps/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a shift swap.
//
// Changes the fields you send and returns the updated shift swap. Send `null` to clear an optional
// field. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ShiftSwapsService) Update(ctx context.Context, id string, params *ShiftSwapUpdate, opts ...RequestOption) (*ShiftSwap, error) {
	if params == nil {
		params = &ShiftSwapUpdate{}
	}
	var body any = params
	var out ShiftSwap
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/shift-swaps/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a shift swap.
//
// Deletes a shift swap. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *ShiftSwapsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ShiftSwapDeleted, error) {
	var out ShiftSwapDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/shift-swaps/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CarePlansService groups the carePlans operations. Care plans describe the services a patient
// receives, how often, and when the plan is next reviewed. The plan's goals and task lists are
// managed in AveeCare and are not included here.
type CarePlansService struct {
	c *core
}

// List care plans.
//
// Returns a page of care plans. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CarePlansService) List(ctx context.Context, params *CarePlanListParams, opts ...RequestOption) (*Page[CarePlan], error) {
	return getPage[CarePlan](ctx, s.c, "/v1/care-plans", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *CarePlansService) ListAutoPaging(ctx context.Context, params *CarePlanListParams, opts ...RequestOption) *Iterator[CarePlan] {
	return newIterator(ctx, func(ctx context.Context) (*Page[CarePlan], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a care plan.
//
// Creates a care plan and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CarePlansService) Create(ctx context.Context, params *CarePlanCreate, opts ...RequestOption) (*CarePlan, error) {
	if params == nil {
		params = &CarePlanCreate{}
	}
	var body any = params
	var out CarePlan
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/care-plans", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a care plan.
//
// Returns one care plan by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *CarePlansService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*CarePlan, error) {
	var out CarePlan
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/care-plans/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a care plan.
//
// Changes the fields you send and returns the updated care plan. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *CarePlansService) Update(ctx context.Context, id string, params *CarePlanUpdate, opts ...RequestOption) (*CarePlan, error) {
	if params == nil {
		params = &CarePlanUpdate{}
	}
	var body any = params
	var out CarePlan
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/care-plans/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a care plan.
//
// Deletes a care plan. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *CarePlansService) Delete(ctx context.Context, id string, opts ...RequestOption) (*CarePlanDeleted, error) {
	var out CarePlanDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/care-plans/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ProgressNotesService groups the progressNotes operations. Progress notes are the clinical
// summaries caregivers write after a visit.
type ProgressNotesService struct {
	c *core
}

// List progress notes.
//
// Returns a page of progress notes. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *ProgressNotesService) List(ctx context.Context, params *ProgressNoteListParams, opts ...RequestOption) (*Page[ProgressNote], error) {
	return getPage[ProgressNote](ctx, s.c, "/v1/progress-notes", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ProgressNotesService) ListAutoPaging(ctx context.Context, params *ProgressNoteListParams, opts ...RequestOption) *Iterator[ProgressNote] {
	return newIterator(ctx, func(ctx context.Context) (*Page[ProgressNote], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a progress note.
//
// Creates a progress note and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ProgressNotesService) Create(ctx context.Context, params *ProgressNoteCreate, opts ...RequestOption) (*ProgressNote, error) {
	if params == nil {
		params = &ProgressNoteCreate{}
	}
	var body any = params
	var out ProgressNote
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/progress-notes", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a progress note.
//
// Returns one progress note by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *ProgressNotesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*ProgressNote, error) {
	var out ProgressNote
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/progress-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a progress note.
//
// Changes the fields you send and returns the updated progress note. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ProgressNotesService) Update(ctx context.Context, id string, params *ProgressNoteUpdate, opts ...RequestOption) (*ProgressNote, error) {
	if params == nil {
		params = &ProgressNoteUpdate{}
	}
	var body any = params
	var out ProgressNote
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/progress-notes/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a progress note.
//
// Deletes a progress note. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *ProgressNotesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ProgressNoteDeleted, error) {
	var out ProgressNoteDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/progress-notes/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AuthorizationsService groups the authorizations operations. Payer authorizations: how many units
// of a service a payer has approved for a patient, and over what dates.
type AuthorizationsService struct {
	c *core
}

// List authorizations.
//
// Returns a page of authorizations. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *AuthorizationsService) List(ctx context.Context, params *AuthorizationListParams, opts ...RequestOption) (*Page[Authorization], error) {
	return getPage[Authorization](ctx, s.c, "/v1/authorizations", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *AuthorizationsService) ListAutoPaging(ctx context.Context, params *AuthorizationListParams, opts ...RequestOption) *Iterator[Authorization] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Authorization], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an authorization.
//
// Creates an authorization and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AuthorizationsService) Create(ctx context.Context, params *AuthorizationCreate, opts ...RequestOption) (*Authorization, error) {
	if params == nil {
		params = &AuthorizationCreate{}
	}
	var body any = params
	var out Authorization
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/authorizations", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an authorization.
//
// Returns one authorization by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *AuthorizationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Authorization, error) {
	var out Authorization
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/authorizations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an authorization.
//
// Changes the fields you send and returns the updated authorization. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AuthorizationsService) Update(ctx context.Context, id string, params *AuthorizationUpdate, opts ...RequestOption) (*Authorization, error) {
	if params == nil {
		params = &AuthorizationUpdate{}
	}
	var body any = params
	var out Authorization
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/authorizations/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an authorization.
//
// Deletes an authorization. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *AuthorizationsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*AuthorizationDeleted, error) {
	var out AuthorizationDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/authorizations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IncidentsService groups the incidents operations. Incident reports: falls, injuries, medication
// errors and other events that need follow-up, often for regulatory reporting.
type IncidentsService struct {
	c *core
}

// List incidents.
//
// Returns a page of incidents. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *IncidentsService) List(ctx context.Context, params *IncidentListParams, opts ...RequestOption) (*Page[Incident], error) {
	return getPage[Incident](ctx, s.c, "/v1/incidents", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *IncidentsService) ListAutoPaging(ctx context.Context, params *IncidentListParams, opts ...RequestOption) *Iterator[Incident] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Incident], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an incident.
//
// Creates an incident and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *IncidentsService) Create(ctx context.Context, params *IncidentCreate, opts ...RequestOption) (*Incident, error) {
	if params == nil {
		params = &IncidentCreate{}
	}
	var body any = params
	var out Incident
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/incidents", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an incident.
//
// Returns one incident by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *IncidentsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Incident, error) {
	var out Incident
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/incidents/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an incident.
//
// Changes the fields you send and returns the updated incident. Send `null` to clear an optional
// field. Available to keys belonging to office staff and caregivers. With a caregiver's or
// patient's key, a saved record that key cannot read back is answered with only `object` and `id`.
// Requires a key with `write` scope.
//
// Requires the write scope.
func (s *IncidentsService) Update(ctx context.Context, id string, params *IncidentUpdate, opts ...RequestOption) (*Incident, error) {
	if params == nil {
		params = &IncidentUpdate{}
	}
	var body any = params
	var out Incident
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/incidents/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an incident.
//
// Deletes an incident. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *IncidentsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*IncidentDeleted, error) {
	var out IncidentDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/incidents/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EventsService groups the events operations. Events your agency documents about a patient,
// caregiver or visit, such as a complaint, a compliment or an attendance issue.
type EventsService struct {
	c *core
}

// List events.
//
// Returns a page of events. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *EventsService) List(ctx context.Context, params *EventListParams, opts ...RequestOption) (*Page[Event], error) {
	return getPage[Event](ctx, s.c, "/v1/events", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *EventsService) ListAutoPaging(ctx context.Context, params *EventListParams, opts ...RequestOption) *Iterator[Event] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Event], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an event.
//
// Creates an event and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *EventsService) Create(ctx context.Context, params *EventCreate, opts ...RequestOption) (*Event, error) {
	if params == nil {
		params = &EventCreate{}
	}
	var body any = params
	var out Event
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/events", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an event.
//
// Returns one event by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *EventsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Event, error) {
	var out Event
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/events/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an event.
//
// Changes the fields you send and returns the updated event. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *EventsService) Update(ctx context.Context, id string, params *EventUpdate, opts ...RequestOption) (*Event, error) {
	if params == nil {
		params = &EventUpdate{}
	}
	var body any = params
	var out Event
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/events/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an event.
//
// Deletes an event. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *EventsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*EventDeleted, error) {
	var out EventDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/events/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ReferralSourcesService groups the referralSources operations. Hospitals, physicians, case
// managers and other organizations that refer clients to your agency.
type ReferralSourcesService struct {
	c *core
}

// List referral sources.
//
// Returns a page of referral sources. Filters combine with AND. Use `next_cursor` to fetch the
// next page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ReferralSourcesService) List(ctx context.Context, params *ReferralSourceListParams, opts ...RequestOption) (*Page[ReferralSource], error) {
	return getPage[ReferralSource](ctx, s.c, "/v1/referral-sources", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ReferralSourcesService) ListAutoPaging(ctx context.Context, params *ReferralSourceListParams, opts ...RequestOption) *Iterator[ReferralSource] {
	return newIterator(ctx, func(ctx context.Context) (*Page[ReferralSource], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a referral source.
//
// Creates a referral source and returns it. Send an `Idempotency-Key` header to make the request
// safe to retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ReferralSourcesService) Create(ctx context.Context, params *ReferralSourceCreate, opts ...RequestOption) (*ReferralSource, error) {
	if params == nil {
		params = &ReferralSourceCreate{}
	}
	var body any = params
	var out ReferralSource
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/referral-sources", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a referral source.
//
// Returns one referral source by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ReferralSourcesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*ReferralSource, error) {
	var out ReferralSource
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/referral-sources/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a referral source.
//
// Changes the fields you send and returns the updated referral source. Send `null` to clear an
// optional field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ReferralSourcesService) Update(ctx context.Context, id string, params *ReferralSourceUpdate, opts ...RequestOption) (*ReferralSource, error) {
	if params == nil {
		params = &ReferralSourceUpdate{}
	}
	var body any = params
	var out ReferralSource
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/referral-sources/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a referral source.
//
// Deletes a referral source. Available to keys belonging to office staff. Requires a key with
// `write` scope.
//
// Requires the write scope.
func (s *ReferralSourcesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ReferralSourceDeleted, error) {
	var out ReferralSourceDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/referral-sources/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ReferralsService groups the referrals operations. Individual referrals of a prospective client,
// tracked from first contact to admission.
type ReferralsService struct {
	c *core
}

// List referrals.
//
// Returns a page of referrals. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ReferralsService) List(ctx context.Context, params *ReferralListParams, opts ...RequestOption) (*Page[Referral], error) {
	return getPage[Referral](ctx, s.c, "/v1/referrals", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ReferralsService) ListAutoPaging(ctx context.Context, params *ReferralListParams, opts ...RequestOption) *Iterator[Referral] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Referral], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a referral.
//
// Creates a referral and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ReferralsService) Create(ctx context.Context, params *ReferralCreate, opts ...RequestOption) (*Referral, error) {
	if params == nil {
		params = &ReferralCreate{}
	}
	var body any = params
	var out Referral
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/referrals", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a referral.
//
// Returns one referral by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ReferralsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Referral, error) {
	var out Referral
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/referrals/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a referral.
//
// Changes the fields you send and returns the updated referral. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *ReferralsService) Update(ctx context.Context, id string, params *ReferralUpdate, opts ...RequestOption) (*Referral, error) {
	if params == nil {
		params = &ReferralUpdate{}
	}
	var body any = params
	var out Referral
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/referrals/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete a referral.
//
// Deletes a referral. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *ReferralsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*ReferralDeleted, error) {
	var out ReferralDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/referrals/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InvoicesService groups the invoices operations. Invoices billed to patients and their families
// for private-pay care. Read-only.
type InvoicesService struct {
	c *core
}

// List invoices.
//
// Returns a page of invoices. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *InvoicesService) List(ctx context.Context, params *InvoiceListParams, opts ...RequestOption) (*Page[Invoice], error) {
	return getPage[Invoice](ctx, s.c, "/v1/invoices", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *InvoicesService) ListAutoPaging(ctx context.Context, params *InvoiceListParams, opts ...RequestOption) *Iterator[Invoice] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Invoice], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve an invoice.
//
// Returns one invoice by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *InvoicesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Invoice, error) {
	var out Invoice
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/invoices/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InvoiceLineItemsService groups the invoiceLineItems operations. The individual charges on an
// invoice. Read-only.
type InvoiceLineItemsService struct {
	c *core
}

// List invoice line items.
//
// Returns a page of invoice line items. Filters combine with AND. Use `next_cursor` to fetch the
// next page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *InvoiceLineItemsService) List(ctx context.Context, params *InvoiceLineItemListParams, opts ...RequestOption) (*Page[InvoiceLineItem], error) {
	return getPage[InvoiceLineItem](ctx, s.c, "/v1/invoice-line-items", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *InvoiceLineItemsService) ListAutoPaging(ctx context.Context, params *InvoiceLineItemListParams, opts ...RequestOption) *Iterator[InvoiceLineItem] {
	return newIterator(ctx, func(ctx context.Context) (*Page[InvoiceLineItem], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve an invoice line item.
//
// Returns one invoice line item by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *InvoiceLineItemsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*InvoiceLineItem, error) {
	var out InvoiceLineItem
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/invoice-line-items/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ClaimsService groups the claims operations. Insurance claims billed to Medicaid, Medicare and
// other payers, with their adjudication status. Read-only.
type ClaimsService struct {
	c *core
}

// List claims.
//
// Returns a page of claims. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ClaimsService) List(ctx context.Context, params *ClaimListParams, opts ...RequestOption) (*Page[Claim], error) {
	return getPage[Claim](ctx, s.c, "/v1/claims", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ClaimsService) ListAutoPaging(ctx context.Context, params *ClaimListParams, opts ...RequestOption) *Iterator[Claim] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Claim], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve a claim.
//
// Returns one claim by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *ClaimsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Claim, error) {
	var out Claim
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/claims/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PayrollEntriesService groups the payrollEntries operations. One caregiver's hours and pay for
// one pay period. Read-only.
type PayrollEntriesService struct {
	c *core
}

// List payroll entries.
//
// Returns a page of payroll entries. Filters combine with AND. Use `next_cursor` to fetch the next
// page. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *PayrollEntriesService) List(ctx context.Context, params *PayrollEntryListParams, opts ...RequestOption) (*Page[PayrollEntry], error) {
	return getPage[PayrollEntry](ctx, s.c, "/v1/payroll-entries", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *PayrollEntriesService) ListAutoPaging(ctx context.Context, params *PayrollEntryListParams, opts ...RequestOption) *Iterator[PayrollEntry] {
	return newIterator(ctx, func(ctx context.Context) (*Page[PayrollEntry], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve a payroll entry.
//
// Returns one payroll entry by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *PayrollEntriesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*PayrollEntry, error) {
	var out PayrollEntry
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/payroll-entries/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ConversationsService groups the conversations operations. In-app message threads between people
// at your agency: one-to-one chats and named group chats. Read-only; start conversations in
// AveeCare.
type ConversationsService struct {
	c *core
}

// List conversations.
//
// Returns a page of conversations. Filters combine with AND. Use `next_cursor` to fetch the next
// page.
//
// Requires the read scope.
func (s *ConversationsService) List(ctx context.Context, params *ConversationListParams, opts ...RequestOption) (*Page[Conversation], error) {
	return getPage[Conversation](ctx, s.c, "/v1/conversations", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *ConversationsService) ListAutoPaging(ctx context.Context, params *ConversationListParams, opts ...RequestOption) *Iterator[Conversation] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Conversation], error) {
		return s.List(ctx, params, opts...)
	})
}

// Retrieve a conversation.
//
// Returns one conversation by ID.
//
// Requires the read scope.
func (s *ConversationsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Conversation, error) {
	var out Conversation
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/conversations/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// MessagesService groups the messages operations. Messages in a conversation. Messages you send
// through the API are sent as the API key's user.
type MessagesService struct {
	c *core
}

// List messages.
//
// Returns a page of messages. Filters combine with AND. Use `next_cursor` to fetch the next page.
//
// Requires the read scope.
func (s *MessagesService) List(ctx context.Context, params *MessageListParams, opts ...RequestOption) (*Page[Message], error) {
	return getPage[Message](ctx, s.c, "/v1/messages", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *MessagesService) ListAutoPaging(ctx context.Context, params *MessageListParams, opts ...RequestOption) *Iterator[Message] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Message], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create a message.
//
// Creates a message and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. With a caregiver's or patient's key, a saved record that key cannot read back is answered
// with only `object` and `id`. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *MessagesService) Create(ctx context.Context, params *MessageCreate, opts ...RequestOption) (*Message, error) {
	if params == nil {
		params = &MessageCreate{}
	}
	var body any = params
	var out Message
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/messages", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve a message.
//
// Returns one message by ID.
//
// Requires the read scope.
func (s *MessagesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Message, error) {
	var out Message
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/messages/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AlertsService groups the alerts operations. Alerts on your agency's alert board, such as a
// credential about to expire or a missed clock-in, and announcements staff post for caregivers.
type AlertsService struct {
	c *core
}

// List alerts.
//
// Returns a page of alerts. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *AlertsService) List(ctx context.Context, params *AlertListParams, opts ...RequestOption) (*Page[Alert], error) {
	return getPage[Alert](ctx, s.c, "/v1/alerts", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *AlertsService) ListAutoPaging(ctx context.Context, params *AlertListParams, opts ...RequestOption) *Iterator[Alert] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Alert], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an alert.
//
// Creates an alert and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AlertsService) Create(ctx context.Context, params *AlertCreate, opts ...RequestOption) (*Alert, error) {
	if params == nil {
		params = &AlertCreate{}
	}
	var body any = params
	var out Alert
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/alerts", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an alert.
//
// Returns one alert by ID. Available to keys belonging to office staff and caregivers.
//
// Requires the read scope.
func (s *AlertsService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Alert, error) {
	var out Alert
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/alerts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an alert.
//
// Changes the fields you send and returns the updated alert. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *AlertsService) Update(ctx context.Context, id string, params *AlertUpdate, opts ...RequestOption) (*Alert, error) {
	if params == nil {
		params = &AlertUpdate{}
	}
	var body any = params
	var out Alert
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/alerts/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an alert.
//
// Deletes an alert. Available to keys belonging to office staff. Requires a key with `write`
// scope.
//
// Requires the write scope.
func (s *AlertsService) Delete(ctx context.Context, id string, opts ...RequestOption) (*AlertDeleted, error) {
	var out AlertDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/alerts/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OfficesService groups the offices operations. Your agency's offices (branches). Patients,
// caregivers and applicants can each belong to one office. There is always exactly one default
// office.
type OfficesService struct {
	c *core
}

// List offices.
//
// Returns a page of offices. Filters combine with AND. Use `next_cursor` to fetch the next page.
// Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *OfficesService) List(ctx context.Context, params *OfficeListParams, opts ...RequestOption) (*Page[Office], error) {
	return getPage[Office](ctx, s.c, "/v1/offices", params.values(), opts)
}

// ListAutoPaging returns an Iterator over every item of List across all pages, fetching pages
// lazily.
func (s *OfficesService) ListAutoPaging(ctx context.Context, params *OfficeListParams, opts ...RequestOption) *Iterator[Office] {
	return newIterator(ctx, func(ctx context.Context) (*Page[Office], error) {
		return s.List(ctx, params, opts...)
	})
}

// Create an office.
//
// Creates an office and returns it. Send an `Idempotency-Key` header to make the request safe to
// retry. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *OfficesService) Create(ctx context.Context, params *OfficeCreate, opts ...RequestOption) (*Office, error) {
	if params == nil {
		params = &OfficeCreate{}
	}
	var body any = params
	var out Office
	if err := s.c.do(ctx, requestSpec{method: "POST", path: "/v1/offices", body: body, idempotent: true}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieve an office.
//
// Returns one office by ID. Available to keys belonging to office staff.
//
// Requires the read scope.
func (s *OfficesService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Office, error) {
	var out Office
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/offices/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update an office.
//
// Changes the fields you send and returns the updated office. Send `null` to clear an optional
// field. Available to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *OfficesService) Update(ctx context.Context, id string, params *OfficeUpdate, opts ...RequestOption) (*Office, error) {
	if params == nil {
		params = &OfficeUpdate{}
	}
	var body any = params
	var out Office
	if err := s.c.do(ctx, requestSpec{method: "PATCH", path: "/v1/offices/" + pathEscape(id), body: body}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete an office.
//
// Deletes an office. The default office, and an agency's only office, cannot be deleted. Available
// to keys belonging to office staff. Requires a key with `write` scope.
//
// Requires the write scope.
func (s *OfficesService) Delete(ctx context.Context, id string, opts ...RequestOption) (*OfficeDeleted, error) {
	var out OfficeDeleted
	if err := s.c.do(ctx, requestSpec{method: "DELETE", path: "/v1/offices/" + pathEscape(id)}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// MeService groups the me operations. The API key making the request: its scopes, its user's role
// and its agency.
type MeService struct {
	c *core
}

// Retrieve the current API key.
//
// Returns the API key making the request, its scopes, the role of the user it belongs to, and its
// agency. Useful to check a key works.
//
// Requires the read scope.
func (s *MeService) Retrieve(ctx context.Context, opts ...RequestOption) (*Me, error) {
	var out Me
	if err := s.c.do(ctx, requestSpec{method: "GET", path: "/v1/me"}, newRequestOptions(opts), &out); err != nil {
		return nil, err
	}
	return &out, nil
}
