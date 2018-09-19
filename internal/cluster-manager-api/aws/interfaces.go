package aws

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaws"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/aws"
)

type Client struct {
	cmaAWSClient cmaaws.ClientInterface
	secretClient awsk8sutil.ClientInterface
}

type ClientInterface interface {
	CreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error)
	GetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error)
	GetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error)
	DeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error)
	UpdateCredentials(in *pb.UpdateAWSCredentialsMsg) (*pb.UpdateAWSCredentialsReply, error)
	Close() error
}
