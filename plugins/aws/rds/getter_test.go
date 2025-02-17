package rds

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

type mockGetRdsAPI func(ctx context.Context, input *rds.DescribeDBInstancesInput) (output *rds.DescribeDBInstancesOutput, err error)

func (m mockGetRdsAPI) DescribeDBInstances(ctx context.Context, input *rds.DescribeDBInstancesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error) {
	// Return an empty list of RDS instances
	return &rds.DescribeDBInstancesOutput{
		DBInstances: []types.DBInstance{},
	}, nil

}

func TestGetListRDS(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		want []types.DBInstance
	}{
		{
			name: "Empty list of RDS instances",
			want: []types.DBInstance{},
		},
	}
	mockGetRdsAPI := mockGetRdsAPI(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListRDS(mockGetRdsAPI); len(got) != 0 {
				t.Errorf("GetListRDS() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
