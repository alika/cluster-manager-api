package cluster_manager_api

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/grpc"
	"github.com/samsung-cnct/cma-operator/pkg/util/cma"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func (s *Server) InstallHelmChart(ctx context.Context, in *pb.InstallHelmChartMsg) (*pb.InstallHelmChartReply, error) {
	res, err := cma.CreateSDSApplication(cma.GenerateSDSApplication(cma.SDSApplicationOptions{
		Name:           in.Chart.Name,
		Namespace:      in.Chart.Namespace,
		Values:         in.Chart.Values,
		PackageManager: in.PackageManagerName,
		Chart: cma.Chart{
			Name: in.Chart.Chart,
			Repository: cma.ChartRepository{
				Name: in.Repo.Name,
				URL:  in.Repo.Url,
			},
			Version: in.Chart.Version,
		},
	}), "default", nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "An error occured trying to create a SDSApplication: %v", err)
	}

	message := "Successfully installed " + in.Chart.Chart + " as " + in.Chart.Name + " in " + in.Chart.Namespace
	return &pb.InstallHelmChartReply{Ok: res, Message: message}, nil
}

func (s *Server) DeleteHelmChart(ctx context.Context, in *pb.DeleteHelmChartMsg) (*pb.DeleteHelmChartReply, error) {
	res, err := cma.DeleteSDSApplication(in.Name, in.Namespace, nil)
	if err != nil {
		return nil, grpc.GenerateError(codes.Internal, "An error occured trying to delete a SDSApplication: %v", err)
	}

	message := "Successfully deleted " + in.Name
	return &pb.DeleteHelmChartReply{Ok: res, Message: message}, nil
}
