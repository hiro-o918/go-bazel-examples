package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/hiro-o918/go-bazel-examples/mock"
)

func Test_echo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := mock.NewMockEcho(ctrl)
	type args struct {
		msg string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		mockFunc func()
	}{
		{
			name: "echo foo",
			args: args{
				msg: "foo",
			},
			want: "foo",
			mockFunc: func() {
				e.EXPECT().Echo("foo").Return("foo")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if got := echo(e, tt.args.msg); got != tt.want {
				t.Errorf("echo() = %v, want %v", got, tt.want)
			}
		})
	}
}
