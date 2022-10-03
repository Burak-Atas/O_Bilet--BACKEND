package controllers

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMy_Profil(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := My_Profil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("My_Profil() = %v, want %v", got, tt.want)
			}
		})
	}
}
