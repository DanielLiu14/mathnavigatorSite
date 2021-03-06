package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllAchievements_Success(t *testing.T) {
	achieveRepo.mockSelectAll = func(publishedOnly bool) ([]domains.Achieve, error) {
		return []domains.Achieve{
			createMockAchievement(1, 2020, "message1"),
			createMockAchievement(2, 2021, "message2"),
		}, nil
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)
	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)
	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get Published
//
func TestGetPublishedAchievements_Success(t *testing.T) {
	achieveRepo.mockSelectAll = func(publishedOnly bool) ([]domains.Achieve, error) {
		return []domains.Achieve{
			createMockAchievement(1, 2020, "message1"),
			createMockAchievement(2, 2021, "message2"),
		}, nil
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/all?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)

	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)

	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get All Grouped By Year
//
func TestGetAllAchievementsGroupedByYear_Success(t *testing.T) {
	achieveRepo.mockSelectAllGroupedByYear = func() ([]domains.AchieveYearGroup, error) {
		return []domains.AchieveYearGroup{
			{
				Year: 2021,
				Achievements: []domains.Achieve{
					createMockAchievement(1, 2021, "message1"),
				},
			},
			{
				Year: 2020,
				Achievements: []domains.Achieve{
					createMockAchievement(2, 2020, "message2"),
				},
			},
		}, nil
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/years", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.AchieveYearGroup
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Achievements[0].Id)
	assert.EqualValues(t, 2021, achieves[0].Achievements[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Achievements[0].Message)
	assert.EqualValues(t, 2, achieves[1].Achievements[0].Id)
	assert.EqualValues(t, 2020, achieves[1].Achievements[0].Year)
	assert.EqualValues(t, "message2", achieves[1].Achievements[0].Message)
	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get Achievement
//
func TestGetAchievement_Success(t *testing.T) {
	achieveRepo.mockSelectById = func(id uint) (domains.Achieve, error) {
		achieve := createMockAchievement(1, 2020, "message1")
		return achieve, nil
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieve domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieve); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieve.Id)
	assert.EqualValues(t, 2020, achieve.Year)
	assert.EqualValues(t, "message1", achieve.Message)
}

func TestGetAchievement_Failure(t *testing.T) {
	achieveRepo.mockSelectById = func(id uint) (domains.Achieve, error) {
		return domains.Achieve{}, errors.New("not found")
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAchievement_Success(t *testing.T) {
	achieveRepo.mockInsert = func(achieve domains.Achieve) error {
		return nil
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAchievement_Failure(t *testing.T) {
	// no mock needed
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 0, "")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAchievement_Success(t *testing.T) {
	achieveRepo.mockUpdate = func(id uint, achieve domains.Achieve) error {
		return nil // Successful update
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAchievement_Invalid(t *testing.T) {
	// no mock needed
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 0, "")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAchievement_Failure(t *testing.T) {
	achieveRepo.mockUpdate = func(id uint, achieve domains.Achieve) error {
		return errors.New("not found")
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishAchievement_Success(t *testing.T) {
	achieveRepo.mockPublish = func(ids []uint) error {
		return nil // Return no error, successful publish!
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	ids := []uint{1}
	marshal, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestPublishAchievement_Failure(t *testing.T) {
	achieveRepo.mockPublish = func(ids []uint) error {
		return errors.New("not found")
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	ids := []uint{1}
	marshal, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAchievement_Success(t *testing.T) {
	achieveRepo.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAchievement_Failure(t *testing.T) {
	achieveRepo.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.AchieveRepo = &achieveRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockAchievement(id uint, year uint, message string) domains.Achieve {
	return domains.Achieve{
		Id:      id,
		Year:    year,
		Message: message,
	}
}

func createBodyFromAchieve(achieve domains.Achieve) io.Reader {
	marshal, err := json.Marshal(&achieve)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
