package main

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
    mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool {
    fmt.Println("Mocked charge notification function")
    fmt.Printf("Value passed in: %d\n", value)
  args := m.Called(value)
 
    return args.Bool(0)
}


func (m *smsServiceMock) DummyFunc() {
    fmt.Println("Dummy")
}


func TestChargeCustomer(t *testing.T) {
    smsService := new(smsServiceMock)
    smsService.On("SendChargeNotification", 100).Return(true)

  myService := MyService{smsService}
    myService.ChargeCustomer(100)
    smsService.AssertExpectations(t)
}