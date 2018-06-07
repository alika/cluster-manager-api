package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/grpc"
	"github.com/samsung-cnct/cma-operator/pkg/layouts"
	"github.com/samsung-cnct/cma-operator/pkg/layouts/poc"
	"github.com/samsung-cnct/cma-operator/pkg/util/ccutil"
	"github.com/samsung-cnct/cma-operator/pkg/util/cma"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	var layout layouts.Layout
	layout = poc.NewLayout()

	options := cma.SDSClusterOptions{
		Name:     in.Name,
		Provider: in.Provider.Name,
		AWS: cma.AWSOptions{
			Region:          in.Provider.Aws.Region,
			SecretAccessKey: in.Provider.Aws.SecretAccessKey,
			SecretKeyId:     in.Provider.Aws.SecretKeyId,
		},
		MaaS: cma.MaaSOptions{
			Endpoint: in.Provider.Maas.Endpoint,
			Username: in.Provider.Maas.Username,
			OAuthKey: in.Provider.Maas.OauthKey,
		},
	}

	sdsCluster, err := cma.CreateSDSCluster(layout.GenerateSDSCluster(options), "default", nil)
	if err != nil {
		if k8sutil.IsResourceAlreadyExistsError(err) {
			return nil, grpc.GenerateError(codes.AlreadyExists, "SDSCluster already exists: %s", in.Name)
		}
		return nil, grpc.GenerateError(codes.Internal, "An internal error occured creating SDSCluster: %s", in.Name)
	}

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     string(sdsCluster.ObjectMeta.UID),
			Name:   sdsCluster.Name,
			Status: string(sdsCluster.Status.Phase),
		},
	}, nil
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	krakenCluster, err := ccutil.GetKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.NotFound, "KrakenCluster not found: %s", in.Name)
	}

	sdsCluster, err := cma.GetSDSCluster(in.Name, "default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.NotFound, "SDSCluster not found: %s", in.Name)
	}

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         string(sdsCluster.ObjectMeta.UID),
			Name:       sdsCluster.Name,
			Status:     string(sdsCluster.Status.Phase),
			Kubeconfig: krakenCluster.Status.Kubeconfig,
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	ok, err := cma.DeleteSDSCluster(in.Name, "default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "An error occured deleting a SDSCluster: %v", err)
	}

	// Shouldn't be needed, but just doing it for now
	ok, err = ccutil.DeleteKrakenCluster(in.Name, "default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "An error occured deleting a KrakenCluster: %v", err)
	}
	return &pb.DeleteClusterReply{Ok: ok, Status: "Deleting"}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	reply = &pb.GetClusterListReply{Ok: true}
	list, err := cma.ListSDSClusters("default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "An error occured getting a list of SDSClusters: %v", err)
	}
	for _, cluster := range list {
		reply.Clusters = append(reply.Clusters, &pb.ClusterItem{
			Id:     string(cluster.ObjectMeta.UID),
			Name:   cluster.Name,
			Status: string(cluster.Status.Phase)})
	}
	return reply, nil
}
