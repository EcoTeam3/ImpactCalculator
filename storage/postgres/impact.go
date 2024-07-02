package postgres

import (
	"database/sql"
	pb "impactcalculator/generated"
)


type NewImpact struct{
	Db *sql.DB
}

func NewImpactRepo(db *sql.DB)*NewImpact{
	return &NewImpact{Db: db}
}

func(I *NewImpact) GetLeaderBoardUsers(leaderBoard *pb.LeaderBoard)(*pb.LeaderBoardUsers, error){
	return nil, nil
}

func(I *NewImpact) GetLeaderBoardGroups(leaderBoard *pb.LeaderBoard)(*pb.LeaderBoardGroups, error){
	return nil, nil
}

func(I *NewImpact) CreateDonation(donation *pb.Donation)(*pb.Status, error){
	_, err := I.Db.Exec(`INSERT INTO donations(user_id, cause, amount, donated_at) VALUES`, 
							donation.UserId, donation.Cause, donation.Amount, donation.DonatedAt)
	if err != nil{
		return &pb.Status{Status: false}, err
	}	
	return &pb.Status{Status: true}, nil
}

func(I *NewImpact) GetDonations(dCause *pb.DonationCause)(*pb.Donations, error){
	donation := &pb.Donations{}
	err := I.Db.QueryRow(`SELECT amount, cause FROM donations`).
				Scan(&donation.Amount, &donation.Cause)
	return donation, err
}