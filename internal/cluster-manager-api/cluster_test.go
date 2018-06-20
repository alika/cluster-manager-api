package cluster_manager_api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/generated/mocks"
	"github.com/satori/go.uuid"
)

func TestCreateCluster(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()
	mockCluster := mocks.NewMockClusterClient(mocker)

	var testCases = []struct {
		name          string                                    // cluster name
		kind          string                                    // kind of cluster
		success       bool                                      // expected success or not
		genClusterMsg func(string, string) *pb.CreateClusterMsg // function to generate test message
		//		genClusterReply func(string, bool) *pb.CreateClusterReply // function to generate test reply
	}{
		{"test-create-maas-succes", "maas", true, generateValidClusterMsg},
		{"test-create-aws-succes", "aws", true, generateValidClusterMsg},
		// {"test-create-maas-fail-bad-request", "maas", false, generateInvalidMaaSClusterMsg},
		// {"test-create-aws-fail-bad-request", "aws", false, generateInvalidAwsClusterMsg},
	}

	for _, tc := range testCases {
		mockCluster.EXPECT().CreateCluster(gomock.Any(), tc.genClusterMsg(tc.name, tc.kind)).Return(generateClusterReply(tc.name, tc.success), nil)
		mockCluster.CreateCluster(context.Background(), tc.genClusterMsg(tc.name, tc.kind))
	}
}

func generateClusterReply(name string, success bool) *pb.CreateClusterReply {
	if !success {
		return &pb.CreateClusterReply{
			Ok: success,
			ClusterOrError: &pb.CreateClusterReply_Error{
				Error: &pb.Error{
					Code:    "500",
					Message: "Internal Server Error",
				},
			},
		}
	}

	return &pb.CreateClusterReply{
		Ok: success,
		ClusterOrError: &pb.CreateClusterReply_Cluster{
			Cluster: &pb.ClusterItem{
				Id:   uuid.NewV4().String(),
				Name: name,
			},
		},
	}
}

func generateValidClusterMsg(name string, clusterType string) *pb.CreateClusterMsg {
	return &pb.CreateClusterMsg{
		Name: name,
		Provider: &pb.CreateClusterProviderSpec{
			Name: clusterType,
			Aws: &pb.CreateClusterAWSSpec{
				Region:          "us-west-2",
				SecretKeyId:     "MYTESTKEY",
				SecretAccessKey: "SUPERSECRET",
			},
			Maas: &pb.CreateClusterMaaSSpec{
				Endpoint: "http://192.168.1.1/MAAS/api/2.0",
				Username: "MaasOperator",
				OauthKey: "MYTESTOAUTHKEY",
			},
		},
	}
}

func TestGetClusterList(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()
	mockCluster := mocks.NewMockClusterClient(mocker)

	mockCluster.EXPECT().GetClusterList(
		gomock.Any(),
		gomock.Any(),
	).Return(&pb.GetClusterListReply{Ok: true}, nil)

	mockCluster.GetClusterList(nil, &pb.GetClusterListMsg{})
}
