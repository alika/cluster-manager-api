package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/grpc"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
	"golang.org/x/net/context"
	"k8s.io/client-go/kubernetes"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cma-operator/pkg/util"
	"google.golang.org/grpc/codes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	logger loggo.Logger
)

type Server struct{}

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldMsg) (*pb.HelloWorldReply, error) {
	return &pb.HelloWorldReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) GetPodCount(ctx context.Context, in *pb.GetPodCountMsg) (*pb.GetPodCountReply, error) {
	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(k8sutil.DefaultConfig)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "Cannot establish a client connection to kubernetes: %v", err)
	}

	pods, err := clientSet.CoreV1().Pods(in.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "Cannot establish a client connection to kubernetes: %v", err)
	}

	logger.Infof("Was asked to get pods on -->%s<-- namespace, answer was -->%d<--", in.Namespace, int32(len(pods.Items)))
	return &pb.GetPodCountReply{Pods: int32(len(pods.Items))}, nil
}

func init() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
