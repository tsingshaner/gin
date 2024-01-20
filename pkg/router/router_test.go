package router

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockRouter struct {
	mock.Mock
}

func (m *MockRouter) Handle(group *gin.RouterGroup) {
	m.Called(group)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	group := router.Group("/")

	mockRouter := new(MockRouter)
	mockRouter.On("Handle", group).Return()

	routes := &[]Router{mockRouter.Handle}

	Register(group, routes)

	mockRouter.AssertCalled(t, "Handle", group)
}
