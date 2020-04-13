package repository

import (
	"database/sql"
	"log"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/assignment"
	"github.com/denisenkom/go-mssqldb"
)

// Assignment implementation of assignment repository
type Assignment struct {
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewAssignmentRepository create new instance of assignment repository
func NewAssignmentRepository(auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Assignment {
	return &Assignment{
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// GetEmployerAssignmentsByAuctionID get all employer assigments by auction id
func (a *Assignment) GetEmployerAssignmentsByAuctionID(auctionID string) ([]assignment.AuctionEmployer, error) {
	query := "select AuctionEmployerAssignmentId, AuctionId, ClientId, UserId from AuctionEmployerAssignments where AuctionId=? and IsDeleted=0"

	a.auctionDB.Ping()

	rows, err := a.auctionDB.Query(query, auctionID)

	if err != nil {
		return nil, err
	}

	var id, auction, clientID, userID *mssql.UniqueIdentifier
	var assignments []assignment.AuctionEmployer

	for rows.Next() {
		rows.Scan(&id, &auction, &clientID, &userID)

		employerAssignment := assignment.NewAuctionEmployer(id.String(), auction.String(), clientID.String(), userID.String())
		assignments = append(assignments, *employerAssignment)
	}

	return assignments, nil
}

// GetEmployerAssignmentsBySuccessFeeID get all employer assingments by success fee id
func (a *Assignment) GetEmployerAssignmentsBySuccessFeeID(successFeeID string) ([]assignment.SuccessFeeEmployer, error) {
	query := "select SuccessFeeEmployerAssignmentId, SuccessFeeId, ClientId, UserId from SuccessFeeEmployerAssignments where SuccessFeeId=? and IsDeleted=0"

	a.auctionDB.Ping()

	rows, err := a.auctionDB.Query(query, successFeeID)

	if err != nil {
		return nil, err
	}

	var id, successFee, clientID, userID mssql.UniqueIdentifier
	var assignments []assignment.SuccessFeeEmployer

	for rows.Next() {
		rows.Scan(&id, &successFee, &clientID, &userID)

		employerAssignment := assignment.NewSuccessFeeEmployer(id.String(), successFee.String(), clientID.String(), userID.String())
		assignments = append(assignments, *employerAssignment)
	}

	log.Print(assignments)
	return assignments, nil
}

// GetAgencyAssignmentsByAuctionID get all agency assignments by auction id
func (a *Assignment) GetAgencyAssignmentsByAuctionID(auctionID string, tenantID string) ([]assignment.AuctionAgency, error) {
	query := "select AuctionAgencyAssignmentId, AuctionId, SupplierId, UserId from AuctionAgencyAssignments where AuctionId=? and SupplierId=?  and IsDeleted=0"

	a.auctionDB.Ping()

	rows, err := a.auctionDB.Query(query, auctionID, tenantID)

	if err != nil {
		a.telemetryClient.TrackException(err)
		return nil, err
	}

	var id, auction, supplierID, userID *mssql.UniqueIdentifier
	var assignments []assignment.AuctionAgency

	for rows.Next() {
		rows.Scan(&id, &auction, &supplierID, &userID)

		agencyAssignment := assignment.NewAuctionAgency(id.String(), auction.String(), supplierID.String(), userID.String())
		assignments = append(assignments, *agencyAssignment)
	}

	return assignments, nil
}

// GetAgencyAssignmentsBySuccessFeeID get all agency assignments by success fee id
func (a *Assignment) GetAgencyAssignmentsBySuccessFeeID(successFeeID string, tenantID string) ([]assignment.SuccessFeeAgency, error) {
	query := "select SuccessFeeAgencyAssignmentId,SuccessFeeId, SupplierId, UserId from SuccessFeeAgencyAssignments where SuccessFeeId=? and SupplierId=? and IsDeleted=0"

	a.auctionDB.Ping()

	rows, err := a.auctionDB.Query(query, successFeeID, tenantID)

	if err != nil {
		a.telemetryClient.TrackException(err)
		return nil, err
	}

	var id, successFee, supplierID, userID *mssql.UniqueIdentifier
	var assignments []assignment.SuccessFeeAgency

	for rows.Next() {
		rows.Scan(&id, &successFee, &supplierID, &userID)

		agencyAssignment := assignment.NewSuccessFeeAgency(id.String(), successFee.String(), supplierID.String(), userID.String())
		assignments = append(assignments, *agencyAssignment)
	}

	return assignments, nil
}

// IsApprovedAuctionStatus check if the auction status is approved
func (a *Assignment) IsApprovedAuctionStatus(auctionID string) bool {
	auctionRepository := NewAuctionRepository(a.auctionDB)
	auction := auctionRepository.getAuctionByID(auctionID)

	return auction.IsApprovedAuction()
}

// IsApprovedSuccessFeeStatus check if the success fee status is approved
func (a *Assignment) IsApprovedSuccessFeeStatus(successFeeID string) bool {
	auctionRepository := NewAuctionRepository(a.auctionDB)
	successFee := auctionRepository.getSuccessFeeByID(successFeeID)

	return successFee.IsApprovedSuccessFee()
}
