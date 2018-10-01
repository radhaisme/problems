package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"io"
	"net/http"
)

var problems []domain.Problem

func HealthCheck(c *gin.Context) {
	if problems == nil || len(problems) == 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail", "reason": "problems setup failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "problemsCount": len(problems)})
}

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(http.StatusOK, problems)
}

// curl -i http://localhost:8080/api/v1/problems/fib
func GetProblem(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, filter(problems, func(problem domain.Problem) bool {
		return problem.Id == id
	}))
}

func LoadProblems(problemsJson io.Reader) error {
	jsonParser := json.NewDecoder(problemsJson)

	if err := jsonParser.Decode(&problems); err != nil {
		return err
	}

	return nil
}

func filter(problems []domain.Problem, f func(problem domain.Problem) bool) domain.Problem {
	for _, problem := range problems {
		if f(problem) {
			return problem
		}
	}

	return domain.Problem{}
}
