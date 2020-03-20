package controllers_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/router"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Global test variables
var handler router.Handler

var programService mockProgramService

// Fake programService that implements ProgramService interface
type mockProgramService struct {
	mockGetAll         func() ([]domains.Program, error)
	mockGetByProgramId func(string) (domains.Program, error)
	mockCreate         func(domains.Program) error
	mockUpdate         func(string, domains.Program) error
	mockDelete         func(string) error
}

// Implement methods of ProgramService interface with mocked implementations
func (programService *mockProgramService) GetAll() ([]domains.Program, error) {
	return programService.mockGetAll()
}
func (programService *mockProgramService) GetByProgramId(programId string) (domains.Program, error) {
	return programService.mockGetByProgramId(programId)
}
func (programService *mockProgramService) Create(program domains.Program) error {
	return programService.mockCreate(program)
}
func (programService *mockProgramService) Update(programId string, program domains.Program) error {
	return programService.mockUpdate(programId, program)
}
func (programService *mockProgramService) Delete(programId string) error {
	return programService.mockDelete(programId)
}

var locationService mockLocationService

// Fake locationService that implements LocationService interface
type mockLocationService struct {
	mockGetAll          func() ([]domains.Location, error)
	mockGetByLocationId func(string) (domains.Location, error)
	mockCreate          func(domains.Location) error
	mockUpdate          func(string, domains.Location) error
	mockDelete          func(string) error
}

// Implement methods of LocationService interface with mocked implementations
func (locationService *mockLocationService) GetAll() ([]domains.Location, error) {
	return locationService.mockGetAll()
}
func (locationService *mockLocationService) GetByLocationId(locId string) (domains.Location, error) {
	return locationService.mockGetByLocationId(locId)
}
func (locationService *mockLocationService) Create(location domains.Location) error {
	return locationService.mockCreate(location)
}
func (locationService *mockLocationService) Update(locId string, location domains.Location) error {
	return locationService.mockUpdate(locId, location)
}
func (locationService *mockLocationService) Delete(locId string) error {
	return locationService.mockDelete(locId)
}

var announceService mockAnnounceService

// Fake announceService that implements AnnounceService interface
type mockAnnounceService struct {
	mockGetAll          func() ([]domains.Announce, error)
	mockGetByAnnounceId func(uint) (domains.Announce, error)
	mockCreate          func(domains.Announce) error
	mockUpdate          func(uint, domains.Announce) error
	mockDelete          func(uint) error
}

// Implement methods of AnnounceService interface with mocked implementations
func (announceService *mockAnnounceService) GetAll() ([]domains.Announce, error) {
	return announceService.mockGetAll()
}
func (announceService *mockAnnounceService) GetByAnnounceId(id uint) (domains.Announce, error) {
	return announceService.mockGetByAnnounceId(id)
}
func (announceService *mockAnnounceService) Create(announce domains.Announce) error {
	return announceService.mockCreate(announce)
}
func (announceService *mockAnnounceService) Update(id uint, announce domains.Announce) error {
	return announceService.mockUpdate(id, announce)
}
func (announceService *mockAnnounceService) Delete(id uint) error {
	return announceService.mockDelete(id)
}

var achieveService mockAchieveService

// Fake achieveService that implements AchieveService interface
type mockAchieveService struct {
	mockGetAll  func() ([]domains.Achieve, error)
	mockGetById func(uint) (domains.Achieve, error)
	mockCreate  func(domains.Achieve) error
	mockUpdate  func(uint, domains.Achieve) error
	mockDelete  func(uint) error
}

// Implement methods of AchieveService interface with mocked implementations
func (achieveService *mockAchieveService) GetAll() ([]domains.Achieve, error) {
	return achieveService.mockGetAll()
}
func (achieveService *mockAchieveService) GetById(id uint) (domains.Achieve, error) {
	return achieveService.mockGetById(id)
}
func (achieveService *mockAchieveService) Create(achieve domains.Achieve) error {
	return achieveService.mockCreate(achieve)
}
func (achieveService *mockAchieveService) Update(id uint, achieve domains.Achieve) error {
	return achieveService.mockUpdate(id, achieve)
}
func (achieveService *mockAchieveService) Delete(id uint) error {
	return achieveService.mockDelete(id)
}

var semesterService mockSemesterService

// Fake semesterService that implements SemesterService interface
type mockSemesterService struct {
	mockGetAll          func() ([]domains.Semester, error)
	mockGetBySemesterId func(string) (domains.Semester, error)
	mockCreate          func(domains.Semester) error
	mockUpdate          func(string, domains.Semester) error
	mockDelete          func(string) error
}

// Implement methods of SemesterService interface with mocked implementations
func (semesterService *mockSemesterService) GetAll() ([]domains.Semester, error) {
	return semesterService.mockGetAll()
}
func (semesterService *mockSemesterService) GetBySemesterId(semesterId string) (domains.Semester, error) {
	return semesterService.mockGetBySemesterId(semesterId)
}
func (semesterService *mockSemesterService) Create(semester domains.Semester) error {
	return semesterService.mockCreate(semester)
}
func (semesterService *mockSemesterService) Update(semesterId string, semester domains.Semester) error {
	return semesterService.mockUpdate(semesterId, semester)
}
func (semesterService *mockSemesterService) Delete(semesterId string) error {
	return semesterService.mockDelete(semesterId)
}

func init() {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	handler = router.Handler{Engine: engine}
	handler.SetupApiEndpoints()
}

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
	}
	w := httptest.NewRecorder()
	handler.Engine.ServeHTTP(w, req)
	return w
}
