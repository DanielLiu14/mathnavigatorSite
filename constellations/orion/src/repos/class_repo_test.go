package repos_test

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

var now = time.Now().UTC()
var later = now.Add(time.Hour * 24 * 60)

func initClassTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.ClassRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestClassRepo(db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllClasses(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll(false)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Test Select Published
//
func TestSelectPublishedClasses(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes WHERE published_at IS NOT NULL").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll(true)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select Unpublished
//
func TestSelectAllUnpublishedClasses(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes WHERE published_at IS NULL").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAllUnpublished()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One
//
func TestSelectClass(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes WHERE class_id=?").
		ExpectQuery().
		WithArgs("program1_2020_spring_final_review").
		WillReturnRows(rows)
	got, err := repo.SelectByClassId("program1_2020_spring_final_review") // Correct classId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getClass()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By Program ID
//
func TestSelectClassesByProgramId(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes WHERE program_id=?").
		ExpectQuery().
		WithArgs("program1").
		WillReturnRows(rows)
	got, err := repo.SelectByProgramId("program1") // Correct programId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By Semester ID
//
func TestSelectClassesBySemesterId(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM classes WHERE semester_id=?").
		ExpectQuery().
		WithArgs("2020_spring").
		WillReturnRows(rows)
	got, err := repo.SelectBySemesterId("2020_spring") // Correct semesterId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By Program ID and Semester ID
//
func TestSelectClassesByProgramIdAndSemesterId(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getClassRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM classes WHERE program_id=\? AND semester_id=?`).
		ExpectQuery().
		WithArgs("program1", "2020_spring").
		WillReturnRows(rows)
	got, err := repo.SelectByProgramAndSemesterId("program1", "2020_spring") // Correct programId and semesterId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Class{getClass()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Create
//
func TestInsertClass(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO classes").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"program1",
			"2020_spring",
			domains.NewNullString("final_review"),
			"program1_2020_spring_final_review",
			"churchill",
			"3 pm - 5 pm",
			now,
			later,
		).WillReturnResult(result)
	class := getClass()
	err := repo.Insert(class)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Update
//
func TestUpdateClass(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE classes SET (.*) WHERE class_id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"program2",
			"2020_summer",
			domains.NewNullString(""),
			"program2_2020_summer",
			"churchill",
			"5 pm - 7 pm",
			now,
			later,
			"program1_2020_spring_final_review",
		).WillReturnResult(result)
	class := domains.Class{
		ProgramId:  "program2",
		SemesterId: "2020_summer",
		ClassKey:   domains.NewNullString(""),
		ClassId:    "program2_2020_summer",
		LocationId: "churchill",
		Times:      "5 pm - 7 pm",
		StartDate:  now,
		EndDate:    later,
	}
	err := repo.Update("program1_2020_spring_final_review", class)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Publish
//
func TestPublishClasses(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectPrepare(`^UPDATE classes SET published_at=\? WHERE class_id=\? AND published_at IS NULL`).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "program1_2020_spring_final_review").
		WillReturnResult(result)
	mock.ExpectCommit()
	err := repo.Publish([]string{"program1_2020_spring_final_review"})
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Delete
//
func TestDeleteClass(t *testing.T) {
	db, mock, repo := initClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM classes WHERE class_id=?").
		ExpectExec().
		WithArgs("program1_2020_spring_final_review").
		WillReturnResult(result)
	err := repo.Delete("program1_2020_spring_final_review")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Helper Methods
//
func getClassRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"PublishedAt",
		"ProgramId",
		"SemesterId",
		"ClassKey",
		"ClassId",
		"locationId",
		"Times",
		"StartDate",
		"EndDate",
	}).AddRow(
		1,
		now,
		now,
		domains.NullTime{},
		domains.NullTime{},
		"program1",
		"2020_spring",
		domains.NewNullString("final_review"),
		"program1_2020_spring_final_review",
		"churchill",
		"3 pm - 5 pm",
		now,
		later,
	)
}

func getClass() domains.Class {
	return domains.Class{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   domains.NullTime{},
		PublishedAt: domains.NullTime{},
		ProgramId:   "program1",
		SemesterId:  "2020_spring",
		ClassKey:    domains.NewNullString("final_review"),
		ClassId:     "program1_2020_spring_final_review",
		LocationId:  "churchill",
		Times:       "3 pm - 5 pm",
		StartDate:   now,
		EndDate:     later,
	}
}
