package handler

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "dasdasd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()
		})
	}
}
