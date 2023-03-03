package errors

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want error
	}{
		{
			"Test error context",
			context.DeadlineExceeded,
			status.Error(codes.Canceled, context.DeadlineExceeded.Error()),
		}, {
			"Test error empty playlist",
			ErrorEmptyPlaylist,
			status.Error(codes.FailedPrecondition, ErrorEmptyPlaylist.Error()),
		}, {
			"Test error not found",
			ErrorNotFound,
			status.Error(codes.NotFound, ErrorNotFound.Error()),
		}, {
			"Test error not valid data",
			ErrorNotValid,
			status.Error(codes.InvalidArgument, ErrorNotValid.Error()),
		}, {
			"Test unknown error",
			errors.New("unknown"),
			status.Error(codes.Unknown, "unknown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.arg); err.Error() != tt.want.Error() {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
