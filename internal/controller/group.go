package controller

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	"github.com/fair-n-square-co/transactions/internal/db"
)

type GroupController interface {
	CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error)
	ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error)
}

// groupController is responsible for group requests.
type groupController struct {
	dbClient db.Client
}

// CreateGroup creates the group using the db package,
// and returns the proto response.
func (g *groupController) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	options := db.CreateGroupOptions{
		Name: req.Name,
	}
	_, err := g.dbClient.CreateGroup(ctx, options)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateGroupResponse{}
	return resp, nil
}

// ListGroups lists the groups using the db package,
// and returns the proto response.
func (g *groupController) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	groups, err := g.dbClient.ListGroups(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListGroupsResponse{}
	for _, group := range groups.Groups {
		resp.Groups = append(resp.Groups, &pb.Group{
			Id:   group.ID.String(),
			Name: group.Name,
		})
	}

	return resp, nil
}

// NewGroupController creates a new instance of GroupController.
func NewGroupController(dbClient db.Client) GroupController {
	return &groupController{
		dbClient: dbClient,
	}
}
