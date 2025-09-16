package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRouterForEmployee(t *testing.T) {
	// Create a test router using the Gin framework
	router := gin.Default()
	routerGroup := router.Group("/api")
	CreateRouterForEmployee(routerGroup)

	// Health Check API
	healthCheckReq, _ := http.NewRequest("GET", "/api/employee/health", nil)
	healthCheckResp := performRequest(router, healthCheckReq)
	assert.Equal(t, http.StatusOK, healthCheckResp.Code)

	// Detailed Health Check API
	detailedHealthCheckReq, _ := http.NewRequest("GET", "/api/employee/health/detail", nil)
	detailedHealthCheckResp := performRequest(router, detailedHealthCheckReq)
	assert.Equal(t, http.StatusOK, detailedHealthCheckResp.Code)

	// Create Employee Data API (with dummy JSON)
	createBody := `{"name":"John Doe","designation":"Engineer"}`
	createEmployeeDataReq, _ := http.NewRequest("POST", "/api/employee/create", strings.NewReader(createBody))
	createEmployeeDataReq.Header.Set("Content-Type", "application/json")
	createEmployeeDataResp := performRequest(router, createEmployeeDataReq)
	assert.Equal(t, http.StatusCreated, createEmployeeDataResp.Code)

	// Read Employee Data API
	readEmployeeDataReq, _ := http.NewRequest("GET", "/api/employee/search", nil)
	readEmployeeDataResp := performRequest(router, readEmployeeDataReq)
	assert.Equal(t, http.StatusOK, readEmployeeDataResp.Code)

	// Read Complete Employees Data API
	readCompleteEmployeesDataReq, _ := http.NewRequest("GET", "/api/employee/search/all", nil)
	readCompleteEmployeesDataResp := performRequest(router, readCompleteEmployeesDataReq)
	assert.Equal(t, http.StatusOK, readCompleteEmployeesDataResp.Code)

	// Read Employees Location API
	readEmployeesLocationReq, _ := http.NewRequest("GET", "/api/employee/search/location", nil)
	readEmployeesLocationResp := performRequest(router, readEmployeesLocationReq)
	assert.Equal(t, http.StatusOK, readEmployeesLocationResp.Code)

	// Read Employees Designation API
	readEmployeesDesignationReq, _ := http.NewRequest("GET", "/api/employee/search/designation", nil)
	readEmployeesDesignationResp := performRequest(router, readEmployeesDesignationReq)
	assert.Equal(t, http.StatusOK, readEmployeesDesignationResp.Code)
}

// Helper function to perform the HTTP request and retrieve the response
func performRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
