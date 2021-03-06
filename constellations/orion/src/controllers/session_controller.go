package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSessionsByClassId(c *gin.Context) {
	classId := c.Param("classId")
	publishedOnly := ParseParamPublishedOnly(c)

	sessionList, err := repos.SessionRepo.SelectAllByClassId(classId, publishedOnly)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, sessionList)
	}
	return
}

func GetSessionById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	session, err := repos.SessionRepo.SelectBySessionId(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &session)
	}
	return
}

func CreateSessions(c *gin.Context) {
	// Incoming JSON
	var sessionsJson []domains.Session
	c.BindJSON(&sessionsJson)

	err := repos.SessionRepo.Insert(sessionsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateSession(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var sessionJson domains.Session
	c.BindJSON(&sessionJson)

	if err := sessionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.SessionRepo.Update(id, sessionJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func PublishSessions(c *gin.Context) {
	// Incoming JSON
	var idsJson []uint
	c.BindJSON(&idsJson)

	err := repos.SessionRepo.Publish(idsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteSessions(c *gin.Context) {
	// Incoming Parameters
	var idsJson []uint
	c.BindJSON(&idsJson)

	err := repos.SessionRepo.Delete(idsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
