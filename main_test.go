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

	// and
	cli := user.NewAPIClient(user.NewConfiguration())
	if host, exist := os.LookupEnv("SAMPLE_HOST"); exist {
		cli.GetConfig().Host = host
	}

	// when
	got, _, err := cli.UserApi.
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
