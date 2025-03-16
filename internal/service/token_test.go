package service

import (
	"testing"
)

func Test_ErrorWhenThereIsNoTokenInEnviromentVariables(t *testing.T) {

	_, err := GetToken()

	if err == nil {
		t.Fatal("error should have occureed, but it was nil")
	}

	want := "not found mnit notion token"
	if err.Error() != want {
		t.Errorf("got error %q, want %q", err.Error(), want)
	}

}

func Test_GetTokenFromEnviromentVariables(t *testing.T) {
	t.Setenv("MNIT_NOTION_TOKEN", "test_token")

	got, _ := GetToken()
	want := "test_token"

	if got != want {
		t.Errorf("got %q, expect %q", got, want)
	}
}
