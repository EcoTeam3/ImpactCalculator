package postgres

import (
	"database/sql"
	"impactcalculator/generated/impact"
	pb "impactcalculator/generated/impact"
)

type NewImpact struct {
	Db *sql.DB
}

func NewImpactRepo(db *sql.DB) *NewImpact {
	return &NewImpact{Db: db}
}

func (I *NewImpact) CreateFootprint(footPrint *pb.CarbonFootprint) (*pb.Status, error) {
	_, err := I.Db.Exec("INSERT INTO Carbon_footprint_logs(user_id,category,amount,unit,logget_at) VALUES($1,$2,$3,$4,$5)",
		footPrint.UserId, footPrint.Category, footPrint.Amount, footPrint.Unit, footPrint.LoggedAt)
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}

func (I *NewImpact) GetUserImpact(userId *pb.UserId) (*pb.Amount, error) {
	amount := &pb.Amount{}
	err := I.Db.QueryRow("SELECT Sum(amount) FROM Carbon_footprint_logs WHERE user_id = $1", userId.UserId).Scan(amount.Amount)
	if err != nil {
		return nil, err
	}
	return amount, err
}

func (I *NewImpact) GetGroupImpact(groupId *pb.GroupId) (*pb.Amount, error) {
	amount := &pb.Amount{}
	err := I.Db.QueryRow("SELECT Sum(amount) FROM Carbon_footprint_logs WHERE user_id = $1", groupId.GroupId).Scan(amount.Amount)
	if err != nil {
		return nil, err
	}
	return amount, err
}

func (I *NewImpact) GetLeaderBoardUsers(leaderBoard *pb.LeaderBoard) (*pb.LeaderBoardUsers, error) {
	users := []string{}
	rows, err := I.Db.Query(`SELECT user_id FROM  Carbon_footprint_logs order by amount`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user string
		err := rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &impact.LeaderBoardUsers{Users: users}, nil
}

func (I *NewImpact) GetLeaderBoardGroups(leaderBoard *pb.LeaderBoard) (*pb.LeaderBoardGroups, error) {
	groups := []string{}
	rows, err := I.Db.Query(`SELECT user_id FROM  Carbon_footprint_logs order by amount`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var group string
		err := rows.Scan(&group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return &impact.LeaderBoardGroups{Groups: groups}, nil
}

func (I *NewImpact) CreateDonation(donation *pb.Donation) (*pb.Status, error) {
	_, err := I.Db.Exec(`INSERT INTO donations(user_id, cause, amount, donated_at) VALUES`,
		donation.UserId, donation.Cause, donation.Amount, donation.DonatedAt)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (I *NewImpact) GetDonations(dCause *pb.DonationCause) (*pb.Donations, error) {
	donation := &pb.Donations{}
	err := I.Db.QueryRow(`SELECT amount, cause FROM donations`).
		Scan(&donation.Amount, &donation.Cause)
	return donation, err
}
