package service

import (
	"context"
	pb "impactcalculator/generated/impact"
	"impactcalculator/storage/postgres"
)

type Server struct {
	pb.UnimplementedImpactServer
	C postgres.NewImpact
}

func NewImpactRepo(c postgres.NewImpact) *Server {
	return &Server{C: c}
}

func (S *Server) CreateFootprint(ctx context.Context, footprint *pb.CarbonFootprint) (*pb.Status, error) {
	status, err := S.C.CreateFootprint(footprint)
	if err != nil {
		return nil, err
	}
	return status, err
}

func (S *Server) GetUserImpact(ctx context.Context, userId *pb.UserId) (*pb.Amount, error) {
	amount, err := S.C.GetUserImpact(userId)
	if err != nil {
		return nil, err
	}
	return amount, nil
}

func (S *Server) GetGroupImpact(ctx context.Context, groupId *pb.GroupId) (*pb.Amount, error) {
	amount, err := S.C.GetGroupImpact(groupId)
	if err != nil {
		return nil, err
	}
	return amount, nil
}

func (S *Server) GetLeaderBoardUsers(ctx context.Context, board *pb.LeaderBoard) (*pb.LeaderBoardUsers, error) {
	leaderboard, err := S.C.GetLeaderBoardUsers(board)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (S *Server) GetLeaderBoardGroups(ctx context.Context, board *pb.LeaderBoard) (*pb.LeaderBoardGroups, error) {
	leaderboard, err := S.C.GetLeaderBoardGroups(board)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (S *Server) CreateDonation(ctx context.Context, donation *pb.Donation) (*pb.Status, error) {
	status, err := S.C.CreateDonation(donation)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (S *Server) GetDonations(ctx context.Context, donationCause *pb.DonationCause) (*pb.Donations, error) {
	donation, err := S.C.GetDonations(donationCause)
	if err != nil {
		return nil, err
	}
	return donation, nil
}