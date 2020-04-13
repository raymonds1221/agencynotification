package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/settings"
	"github.com/denisenkom/go-mssqldb"
)

// Settings implementation of settings repository
type Settings struct {
	employerDB      *sql.DB
	agencyDB        *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewSettingsRepository create new instance of settings repository
func NewSettingsRepository(employerDB *sql.DB, agencyDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Settings {
	return &Settings{
		employerDB:      employerDB,
		agencyDB:        agencyDB,
		telemetryClient: telemetryClient,
	}
}

// GetSettingsByClientID get the settings by specified client id
func (s *Settings) GetSettingsByClientID(clientID string) (settings.Settings, error) {
	const query = `SELECT 
								ActivityStreamSettingId,
								AuctionsScheduling,
								Prequalification,
								Applications,
								Clarifications,
								Bidding,
								Awarding,
								Fulfillment,
								Payments,
								Ubidy,
								Messages,
								Users
							FROM ActivityStreamSettings
							WHERE UserId=?`
	s.employerDB.Ping()

	rows, err := s.employerDB.Query(query, clientID)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return settings.Settings{}, err
	}

	if n := rows.Next(); !n {
		return settings.NewSettings("", true, true, true, true, true, true, true, true, true, true, true), errors.New("no record found")
	}

	var activityStreamSettingID *mssql.UniqueIdentifier
	var auctionsScheduling,
		prequalification,
		applications,
		clarifications,
		bidding,
		awarding,
		fulfillment,
		payments,
		ubidy,
		messages,
		users bool

	err = rows.Scan(&activityStreamSettingID,
		&auctionsScheduling,
		&prequalification,
		&applications,
		&clarifications,
		&bidding,
		&awarding,
		&fulfillment,
		&payments,
		&ubidy,
		&messages,
		&users)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return settings.Settings{}, err
	}

	newSettings := settings.NewSettings(activityStreamSettingID.String(), auctionsScheduling,
		prequalification, applications, clarifications, bidding, awarding, fulfillment, payments, ubidy, messages, users)

	return newSettings, nil
}

// GetSettingsBySupplierID get the settings by specified supplier id
func (s *Settings) GetSettingsBySupplierID(supplierID string) (settings.Settings, error) {
	const query = `SELECT 
								ActivityStreamSettingId,
								AuctionsScheduling,
								Prequalification,
								Applications,
								Clarifications,
								Bidding,
								Awarding,
								Fulfillment,
								Payments,
								Ubidy,
								Messages,
								Users
							FROM ActivityStreamSettings
							WHERE UserId=?`
	s.agencyDB.Ping()

	rows, err := s.agencyDB.Query(query, supplierID)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return settings.Settings{}, err
	}

	if n := rows.Next(); !n {
		return settings.NewSettings("", true, true, true, true, true, true, true, true, true, true, true), errors.New("no record found")
	}

	var activityStreamSettingID string
	var auctionsScheduling,
		prequalification,
		applications,
		clarifications,
		bidding,
		awarding,
		fulfillment,
		payments,
		ubidy,
		messages,
		users bool

	err = rows.Scan(&activityStreamSettingID,
		&auctionsScheduling,
		&prequalification,
		&applications,
		&clarifications,
		&bidding,
		&awarding,
		&fulfillment,
		&payments,
		&ubidy,
		&messages,
		&users)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return settings.Settings{}, err
	}

	newSettings := settings.NewSettings(activityStreamSettingID, auctionsScheduling,
		prequalification, applications, clarifications, bidding, awarding, fulfillment, payments, ubidy, messages, users)

	return newSettings, nil
}

// CreateOrUpdateSettings create or update a settings for the agency
func (s *Settings) CreateOrUpdateSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool, activeSettings string) bool {
	_, err := s.GetSettingsBySupplierID(userID)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return s.createSettings(userID, auctionsScheduling, prequalification, applications, clarifications,
			bidding, awarding, fulfillment, payments, ubidy, messages, users)
	}

	return s.updateSettings(userID, auctionsScheduling, prequalification, applications, clarifications,
		bidding, awarding, fulfillment, payments, ubidy, messages, users, activeSettings)
}

func (s *Settings) createSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool) bool {
	query := `INSERT INTO ActivityStreamSettings(
							AuctionsScheduling, 
							Prequalification, 
							Applications, 
							Clarifications, 
							Bidding, 
							Awarding, 
							Fulfillment, 
							Payments, 
							Ubidy, 
							Messages, 
							Users,
							UserId,
							CreatedBy
						) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	tx, err := s.agencyDB.Begin()

	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Print(err)
		s.telemetryClient.TrackException(err)
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(auctionsScheduling, prequalification, applications,
		clarifications, bidding, awarding, fulfillment, payments, ubidy, messages, users, userID, userID)

	if err != nil {
		s.telemetryClient.TrackException(err)
		return false
	}

	tx.Commit()

	return true
}

func (s *Settings) updateSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool, activeSettings string) bool {
	var query = ""

	switch activeSettings {
	case "Applications":
		query = "UPDATE ActivityStreamSettings SET Applications=? where UserId=?"
		break
	case "Clarifications":
		query = "UPDATE ActivityStreamSettings SET Clarifications=? where UserId=?"
		break
	case "Bidding":
		query = "UPDATE ActivityStreamSettings SET Bidding=? where UserId=?"
		break
	case "Fulfillment":
		query = "UPDATE ActivityStreamSettings SET Fulfillment=? where UserId=?"
		break
	}

	tx, err := s.agencyDB.Begin()

	if err != nil {
		s.telemetryClient.TrackException(err)
		panic(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(query)

	if err != nil {
		s.telemetryClient.TrackException(err)
		panic(err)
	}

	defer stmt.Close()

	switch activeSettings {
	case "Applications":
		_, err = stmt.Exec(applications, userID)
		break
	case "Clarifications":
		_, err = stmt.Exec(clarifications, userID)
		break
	case "Bidding":
		_, err = stmt.Exec(bidding, userID)
		break
	case "Fulfillment":
		_, err = stmt.Exec(fulfillment, userID)
		break
	}

	if err != nil {
		return false
	}

	tx.Commit()

	return true
}
