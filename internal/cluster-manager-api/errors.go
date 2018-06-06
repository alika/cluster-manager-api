package cluster_manager_api

import (
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//WORKAROUND: Should be able to use pb.ErrorReply, but because protobuf
//generates the Go 'Ok' boolean field with `json:"ok,omitempty"`, it causes
//boolean values of false not to be emitted.
type errorBody struct {
	Ok    bool      `protobuf:"varint,1,opt,name=ok" json:"ok"`
	Error *pb.Error `protobuf:"bytes,2,opt,name=error" json:"error"`
}

func GenerateError(c codes.Code, format string, a ...interface{}) error {
	grpcErr := status.Newf(c, format, a)
	logger.Errorf(format, a)
	return grpcErr.Err()
}

func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"ok": false, "error": {"code": "500", "message": "failed to marshal error message"}}`
	w.Header().Set("Content-type", marshaler.ContentType())

	grpcStatus, ok := status.FromError(err)
	if !ok {
		logger.Errorf("Was unable to get the GRPC status from the error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fallback))
		return
	}

	//CANDIDATE TODO: check grpcStatus.WithDetails() for more detailed error info.

	w.WriteHeader(runtime.HTTPStatusFromCode(grpcStatus.Code()))
	errReply := errorBody{
		Ok: false,
		Error: &pb.Error{
			Code:    strconv.FormatInt(int64(grpcStatus.Code()), 10),
			Message: grpcStatus.Message(),
		},
	}

	encodeErr := marshaler.NewEncoder(w).Encode(errReply)
	if encodeErr != nil {
		logger.Errorf("Was unable encode JSON error reply")
		w.Write([]byte(fallback))
	}
}
