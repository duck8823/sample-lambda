package main_test

import (
	"context"
	"github.com/duck8823/sample-lambda/generated/go/user"
	"os"
	"reflect"
	"testing"
)

func Test_APIClient(t *testing.T) {
	if _, isCI := os.LookupEnv("CI"); isCI {
		t.Skip("起動しとかんとあかんからCIでは動かないよ")
	}

	// given
	want := "\"Hello John!\""

	// when
	got, _, err := user.NewAPIClient(user.NewConfiguration()).UserApi.
		CreateUser(context.Background()).
		User(user.User{Name: "John"}).
		Execute()

	// then
	if err != nil {
		t.Error(err)
	}

	// and
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %#v, but got: %#v", want, got)
	}
}
