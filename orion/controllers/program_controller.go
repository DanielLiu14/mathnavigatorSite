package controllers

import (
    // "errors"
    "net/http"
    // "regexp"
    "github.com/gin-gonic/gin"
    "github.com/ahsu1230/mathnavigatorSite/orion/services"
)

func GetAllPrograms(c *gin.Context) {
	programList, err := services.ProgramService.GetAll()
    if err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
    } else {
        c.JSON(http.StatusOK, programList)
    }
	return
}

func GetProgramById(c *gin.Context) {
	return
}

func CreateProgram(c *gin.Context) {
	return
}

func UpdateProgram(c *gin.Context) {
	return
}

func DeleteProgram(c *gin.Context) {
	return
}

// const REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`
// const REGEX_NAME = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

// // Implements ProgramService in domains.ProgramService
// type ProgramService struct {
//     db *sql.DB       // golang native db connection
//     dbx *sqlx.DB     // sqlx wrapper over db connection
// }

// func (ps *ProgramService) GetAll(c *gin.Context) {
// 	programList, err := store.GetAllPrograms(ps.dbx)
//     if err != nil {
//         c.Status(http.StatusInternalServerError)
//     } else {
//         c.JSON(http.StatusOK, programList)
//     }
// 	return
// }

// func (ps *ProgramService) GetByProgramId(c *gin.Context) {
//     // Incoming parameters
// 	programId := c.Param("programId")

// 	// Query Repo
// 	program, err := store.GetProgramById(ps.dbx, programId)
// 	if err != nil {
// 		c.String(http.StatusNotFound, programId)
// 	} else {
// 		c.JSON(http.StatusOK, program)
// 	}
// 	return
// }

// func (ps *ProgramService) Create(c *gin.Context) {
//     // Incoming JSON
// 	var programJson domains.Program
// 	c.BindJSON(&programJson)

// 	if err := CheckValidProgram(programJson); err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	// Query Repo (INSERT & SELECT)
// 	err := store.InsertProgram(ps.dbx, programJson)
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 	} else {
// 		c.JSON(http.StatusOK, nil)
// 	}
// 	return
// }

// func (ps *ProgramService) Update(c *gin.Context) {
//     // Incoming JSON & Parameters
// 	programId := c.Param("programId")
// 	var programJson domains.Program
// 	c.BindJSON(&programJson)

// 	if err := CheckValidProgram(programJson); err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	// Query Repo (UPDATE & SELECT)
// 	err := store.UpdateProgram(ps.dbx, programId, programJson)
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 	} else {
// 		c.JSON(http.StatusOK, nil)
// 	}
// 	return
// }

// func (ps *ProgramService) Delete(c *gin.Context) {
//     // Incoming Parameters
//     programId := c.Param("programId")

//     // Query Repo (DELETE)
//     err := store.DeleteProgram(ps.dbx, programId)
//     if err != nil {
//         c.Status(http.StatusInternalServerError)
//     } else {
//         c.Status(http.StatusOK)
//     }
//     return
// }

// func CheckValidProgram(program domains.Program) error {
// 	// Retrieves the inputted values
// 	programId := program.ProgramId
// 	name := program.Name
// 	grade1 := program.Grade1
// 	grade2 := program.Grade2
// 	description := program.Description

// 	// Checks if the program ID is in the form of alphanumeric strings separated by underscores
// 	if matches, _ := regexp.MatchString(REGEX_PROGRAM_ID, programId); !matches {
// 		return errors.New("invalid program id")
// 	}

// 	// Name validation
// 	match, _ := regexp.MatchString(REGEX_NAME, name)
// 	if !match {
// 		return errors.New("invalid program name")
// 	}

// 	// Checks if the grades are valid
// 	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
// 		return errors.New("invalid grades")
// 	}

// 	// Description validation
// 	if description == "" {
// 		return errors.New("empty description")
// 	}

// 	return nil
// }
