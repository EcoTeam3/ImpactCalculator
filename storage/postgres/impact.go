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

func (I *NewImpact) CreateFootprint(footPrint *pb.CarbonFootprint)(*pb.Status,error){
	_,err :=I.Db.Exec("INSERT INTO Carbon_footprint_logs(user_id,category,amount,unit,logget_at) VALUES($1,$2,$3,$4,$5)",
	footPrint.UserId,footPrint.Category,footPrint.Amount,footPrint.Unit,footPrint.LoggedAt)
	if err != nil{
		return &pb.Status{
			Status: false,
		},err
	}
	return &pb.Status{
		Status: true,
	},nil
}

func (I *NewImpact) GetUserImpact(userId *pb.UserId)(*pb.CarbonFootprints,error){
	rows,err := I.Db.Query("SELECT * FROM Carbon_footprint_logs WHERE user_id = $1",userId.UserId)
	if err != nil{
		return nil,err
	}
	footPrints := &pb.CarbonFootprints{}
	for rows.Next(){
		footPrint := &pb.CarbonFootprint{}
		err = rows.Scan(&footPrint.FootId,&footPrint.UserId,&footPrint.Category,&footPrint.Amount,&footPrint.Unit,&footPrint.LoggedAt)
		if err != nil{
			return nil,err
		}
		footPrints.CarbonFootprints = append(footPrints.CarbonFootprints, footPrint)
	}
	return footPrints,err
}

func (I *NewImpact) GetGroupImpact(groupId *pb.GroupId)(*pb.CarbonFootprint,error){
	return nil,nil

  
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