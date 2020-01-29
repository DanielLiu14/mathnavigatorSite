package controllers
import (
  "net/http"
  "github.com/gin-gonic/gin"
  "orion/models"
)

func GetPrograms(c *gin.Context) {
  // Query DB
  results := []models.Program{}
  models.GetDb().Find(&results)

  // JSON Response
  c.JSON(http.StatusOK, results)
  return;
}

func GetProgram(c *gin.Context) {
  // Incoming parameters
  programId := c.Param("programId")

  // Query DB
  var foundProgram models.Program
  query := models.GetDb().Where(&models.Program{ProgramId: programId}).First(&foundProgram)

  // JSON Response
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    c.JSON(http.StatusOK, foundProgram)
  }
  return;
}

func CreateProgram(c *gin.Context) {
  // Incoming JSON
  var newProgram models.Program
  c.BindJSON(&newProgram)

  // *todo* implement for real!
  // isValid := CheckValidProgram(newProgram);
  // if (isValid) {
  //
  // } else {
  //
  // }

  // Query DB
  db := models.GetDb()
  programName := newProgram.Name
  var foundProgram models.Program
  query := db.Where(&models.Program{Name: programName}).First(&foundProgram)

  // JSON Response
  if query.RecordNotFound() {
    db.Create(&newProgram)
    c.JSON(http.StatusOK, gin.H{
      "program": newProgram,
    })
  } else {
    c.String(http.StatusBadRequest, "Program already exists " + programName)
  }
  return;
}

func UpdateProgram(c *gin.Context) {
  // Incoming JSON & Parameters
  programId := c.Param("programId")
  var updatedProgram models.Program
  c.BindJSON(&updatedProgram)

  // Query DB
  db := models.GetDb()
  var foundProgram models.Program
  query := db.Where(&models.Program{ProgramId: programId}).First(&foundProgram)

  // JSON Response
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    foundProgram.ProgramId = updatedProgram.ProgramId
    foundProgram.Name = updatedProgram.Name
    foundProgram.Grade1 = updatedProgram.Grade1
    foundProgram.Grade2 = updatedProgram.Grade2
    foundProgram.Description = updatedProgram.Description
    db.Save(&foundProgram)
    c.String(http.StatusOK, "Updated Program " + programId)
  }
  return;
}

func DeleteProgram(c *gin.Context) {
  // Incoming Parameters
  programId := c.Param("programId")

  // Query DB
  db := models.GetDb()
  var foundProgram models.Program
  query := db.Where(&models.Program{ProgramId: programId}).First(&foundProgram)

  // JSON Response
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    db.Delete(&foundProgram)
    c.String(http.StatusOK, "Deleted Program " + programId)
  }
  return;
}

func CheckValidProgram() bool {
  return true;
}