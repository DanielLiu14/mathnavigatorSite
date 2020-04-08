package integration_tests

import (
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Semesters and GetAll(false)
func Test_CreateSemesters(t *testing.T) {
	resetTable(t, domains.TABLE_SEMESTERS)

	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_fall", "Fall 2020")
	semester3 := createSemester("2020_winter", "Winter 2020")
	body1 := createJsonBody(&semester1)
	body2 := createJsonBody(&semester2)
	body3 := createJsonBody(&semester3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semesters []domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semesters); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", semesters[0].SemesterId)
	assert.EqualValues(t, "Spring 2020", semesters[0].Title)
	assert.EqualValues(t, "2020_fall", semesters[1].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters[1].Title)
	assert.EqualValues(t, "2020_winter", semesters[2].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[2].Title)
	assert.EqualValues(t, 3, len(semesters))
}

// Test: Create 2 Semesters with same semesterId. Then GetBySemesterId()
func Test_UniqueSemesterId(t *testing.T) {
	resetTable(t, domains.TABLE_SEMESTERS)

	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_spring", "Fall 2020") // Same semesterId
	body1 := createJsonBody(&semester1)
	body2 := createJsonBody(&semester2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder3.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", semester.SemesterId)
	assert.EqualValues(t, "Spring 2020", semester.Title)
}

// Test: Create 1 Semester, Update it, GetBySemesterId()
func Test_UpdateSemester(t *testing.T) {
	resetTable(t, domains.TABLE_SEMESTERS)

	// Create 1 Semester
	semester1 := createSemester("2020_spring", "Spring 2020")
	body1 := createJsonBody(&semester1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedSemester := createSemester("2020_fall", "Fall 2020")
	updatedBody := createJsonBody(&updatedSemester)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/semester/2020_spring", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_fall", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semester.SemesterId)
	assert.EqualValues(t, "Fall 2020", semester.Title)
}

// Test: Create 1 Semester, Delete it, GetBySemesterId()
func Test_DeleteSemester(t *testing.T) {
	resetTable(t, domains.TABLE_SEMESTERS)

	// Create
	semester1 := createSemester("2020_spring", "Spring 2020")
	body1 := createJsonBody(&semester1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/semesters/v1/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

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

	// Get
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

	// Get
	recorder5 := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var semesters2 []domains.Semester
	if err := json.Unmarshal(recorder5.Body.Bytes(), &semesters2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semesters2[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters2[0].Title)
	assert.EqualValues(t, 1, len(semesters2))
}

// Helper methods
func createSemester(semesterId string, title string) domains.Semester {
	return domains.Semester{
		SemesterId: semesterId,
		Title:      title,
	}
}
