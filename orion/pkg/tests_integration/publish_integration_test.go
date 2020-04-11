package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 2 Semesters and Publish 1
func Test_PublishSemesters(t *testing.T) {
	resetTable(t, domains.TABLE_SEMESTERS)

	// Create
	semester1 := createSemester("2020_fall", "Fall 2020")
	semester2 := createSemester("2020_winter", "Winter 2020")
	body1 := createJsonBody(&semester1)
	body2 := createJsonBody(&semester2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var semesters1 []domains.Semester
	if err := json.Unmarshal(recorder3.Body.Bytes(), &semesters1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(semesters1))

	// Publish
	semesterIds := []string{"2020_fall"}
	body3 := createJsonBody(&semesterIds)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var semesters2 []domains.Semester
	if err := json.Unmarshal(recorder5.Body.Bytes(), &semesters2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, semesters2[0].Id)
	assert.EqualValues(t, "2020_fall", semesters2[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters2[0].Title)
	assert.EqualValues(t, 1, len(semesters2))

	// Get All Unpublished
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, unpublishedDomains.Semesters[0].Id)
	assert.EqualValues(t, "2020_winter", unpublishedDomains.Semesters[0].SemesterId)
	assert.EqualValues(t, "Winter 2020", unpublishedDomains.Semesters[0].Title)
	assert.EqualValues(t, 1, len(unpublishedDomains.Semesters))
}
