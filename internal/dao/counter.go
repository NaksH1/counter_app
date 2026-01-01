package dao

import (
	"counterapp/internal/config"
	"counterapp/internal/model"
	"counterapp/internal/util"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	_= godotenv.Load()

	cfg := config.Load()
	dsn := cfg.GetDSN()

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Profile{}, &model.Locker{}, &model.Visit{}, &model.Schedule{}, &model.SevaType{}, &model.StayArea{}, &model.Feedback{})
}

type GetProfilesDataResponse struct {
	ID            string
	Name          string
	PhoneNumber   string
	Email         string
	Gender        string
	Category      string
	IsBlocked     bool
	Remarks       string
	ActiveVisitID *string
	DepartureDate *string
	StayArea      *string
	Status        *string
}

func GetProfilesData(db *gorm.DB) ([]GetProfilesDataResponse, error) {
	sql := `
	SELECT    
	    p.id,
		p.NAME,
		p.phone_number,
		p.email,
		p.gender,
		p.category,
		p.is_blocked,
		p.remarks,
		v.id as active_visit_id,
		v.departure_date,
		sa.name as stay_area,
		v.status
FROM profiles p
LEFT JOIN LATERAL
		(
			select   *
			from     visits v
			WHERE    v.profile_id = p.id AND
			v.status = 'checked-in'
) v
ON true
LEFT JOIN stay_areas sa ON v.stay_area_id = sa.id
	`
	var result []sqlGetProfileData
	err := db.Raw(sql).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	var profilesData []GetProfilesDataResponse
	for _, r := range result {
		response := GetProfilesDataResponse{
			ID:            r.ID,
			Name:          r.Name,
			PhoneNumber:   r.PhoneNumber,
			Email:         r.Email,
			Category:      string(r.Category),
			IsBlocked:     r.IsBlocked,
			Remarks:       r.Remarks,
			StayArea:      r.StayArea,
			ActiveVisitID: r.ActiveVisitID,
		}

		if r.DepartureDate != nil {
			formatted := util.FormatDate(*r.DepartureDate)
			response.DepartureDate = &formatted
		}
		if r.Status != nil {
			statusPtr := string(*r.Status)
			response.Status = &statusPtr
		}

		profilesData = append(profilesData, response)
	}

	return profilesData, nil
}

func CreateProfile(db *gorm.DB, profile *model.Profile) (*model.Profile, error) {
	if profile == nil {
		return nil, errors.New("cannot create an empty profile")
	}
	result := db.Create(profile)
	return profile, result.Error
}

type ProfileUpdate struct {
	Name        *string
	PhoneNumber *string
	Gender      *model.Gender
	Category    *model.Category
	IsBlocked   *bool
	Remarks     *string
}

func UpdateProfile(db *gorm.DB, profileID string, updates *ProfileUpdate) (*model.Profile, error) {
	result := db.Model(&model.Profile{}).Where("id = ?", profileID).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	var updatedProfile model.Profile
	db.First(&updatedProfile, "id = ?", profileID)
	return &updatedProfile, nil
}

func GetVisitsByProfileID(db *gorm.DB, profileID string) ([]model.Visit, error) {
	var visits []model.Visit
	result := db.Preload("StayArea").Preload("Locker").Find(&visits, "profile_id = ?", profileID)
	if result.Error != nil {
		return nil, result.Error
	}
	return visits, nil
}

func GetAllLockers(db *gorm.DB) ([]model.Locker, error) {
	var lockers []model.Locker
	result := db.Find(&lockers)
	if result.Error != nil {
		fmt.Print("error while fetching lockers")
		return nil, result.Error
	}

	return lockers, nil
}

func GetScheduleForDateRange(db *gorm.DB, startDate time.Time, endDate time.Time) ([]model.Schedule, error) {
	var schedules []model.Schedule
	result := db.
		Preload("Profile").
		Preload("SevaType").
		Preload("Visit.StayArea").
		Preload("Visit.Locker").
		Where("date >= ? AND date <= ?", startDate, endDate).
		Find(&schedules)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedules, nil
}

func GetProfileByID(db *gorm.DB, profileID string) (*model.Profile, error) {
	var profile model.Profile
	result := db.Find(&profile, "id = ?", profileID)
	if result.Error != nil {
		fmt.Printf("error while fetching profile for id: %s", profileID)
		return nil, result.Error
	}
	return &profile, nil
}

type UpdateVisitRequest struct {
	DepartureDate *time.Time
	StayAreaID    *uuid.UUID
	Status        *model.ProfileStatus
	LockerID      *uuid.UUID
	Remarks       *string
}

func UpdateVisit(db *gorm.DB, visitID string, req UpdateVisitRequest) (*model.Visit, error) {
	result := db.Model(&model.Visit{}).Where("id = ?", visitID).Updates(req)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedVisit *model.Visit
	db.First(&updatedVisit, "id = ?", visitID)
	return updatedVisit, nil
}

type AddVisitRequest struct {
	ProfileID     uuid.UUID
	ArrivalDate   time.Time
	DepartureDate *time.Time
	StayAreaID    uuid.UUID
	Status        model.ProfileStatus
	LockerID      *uuid.UUID
	Remarks       *string
}

func AddVisit(db *gorm.DB, req AddVisitRequest) (*model.Visit, error) {
	var visit *model.Visit
	err := db.Transaction(func(tx *gorm.DB) error {
		visit = &model.Visit{
			ProfileID:     req.ProfileID,
			ArrivalDate:   req.ArrivalDate,
			DepartureDate: req.DepartureDate,
			StayAreaID:    req.StayAreaID,
			Status:        req.Status,
			LockerID:      req.LockerID,
			Remarks:       req.Remarks,
		}
		return tx.Create(&visit).Error
	})
	if err != nil {
		return nil, err
	}
	return visit, nil
}

func GetVisitByID(db *gorm.DB, visitID string) (*model.Visit, error) {
	var visit model.Visit
	result := db.Preload("StayArea").Preload("Locker").Find(&visit, "id = ?", visitID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &visit, nil
}

func GetScheduleByProfileAndDate(db *gorm.DB, profileID string, date time.Time) (*model.Schedule, error) {
	var sch model.Schedule
	result := db.Find(&sch, "profile_id = ? AND date = ?", profileID, date)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &sch, nil
}

type AddScheduleRequest struct {
	ProfileID  uuid.UUID
	VisitID    uuid.UUID
	SevaTypeID uuid.UUID
	Location   *string
	Date       time.Time
}

func AddSchedule(db *gorm.DB, req AddScheduleRequest) (*model.Schedule, error) {
	schedule := &model.Schedule{
		ProfileID:  req.ProfileID,
		VisitID:    req.VisitID,
		SevaTypeID: req.SevaTypeID,
		Location:   req.Location,
		Date:       req.Date,
	}
	result := db.Create(schedule)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedule, nil
}

func GetAllFeedbacks(db *gorm.DB) ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	result := db.Preload("Profile").Find(&feedbacks)
	if result.Error != nil {
		return nil, result.Error
	}
	return feedbacks, nil
}

func GetFeedbacksForProfile(db *gorm.DB, profileID string) ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	result := db.Where("profile_id = ?", profileID).Find(&feedbacks)
	if result.Error != nil {
		return nil, result.Error
	}
	return feedbacks, nil
}

type AddFeedbackRequest struct {
	ProfileID uuid.UUID
	VisitID   *uuid.UUID
	Content   string
	Type      model.FeedbackType
	CreatedBy *string
}

func AddFeedback(db *gorm.DB, req AddFeedbackRequest) (*model.Feedback, error) {
	feedback := &model.Feedback{
		ProfileID: req.ProfileID,
		VisitID:   req.VisitID,
		Content:   req.Content,
		Type:      req.Type,
		CreatedBy: req.CreatedBy,
	}
	result := db.Create(feedback)
	if result.Error != nil {
		return nil, result.Error
	}
	return feedback, nil
}

func GetAllSevaTypes(db *gorm.DB) ([]model.SevaType, error) {
	var sevaTypes []model.SevaType
	result := db.Where("is_active = ?", true).Find(&sevaTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return sevaTypes, nil
}

func GetSevaTypeByName(db *gorm.DB, name string) (*model.SevaType, error) {
	var sevaType model.SevaType
	result := db.Where("name = ? AND is_active = ?", name, true).First(&sevaType)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sevaType, nil
}

func GetSevaTypeByID(db *gorm.DB, id string) (*model.SevaType, error) {
	var sevaType model.SevaType
	result := db.Where("id = ? AND is_active = ?", id, true).First(&sevaType)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sevaType, nil
}

type AddSevaTypeRequest struct {
	Name        string
	Description *string
}

func AddSevaType(db *gorm.DB, req AddSevaTypeRequest) (*model.SevaType, error) {
	sevaType := &model.SevaType{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}
	result := db.Create(sevaType)
	if result.Error != nil {
		return nil, result.Error
	}
	return sevaType, nil
}

func GetAllStayAreas(db *gorm.DB) ([]model.StayArea, error) {
	var stayAreas []model.StayArea
	result := db.Find(&stayAreas)
	if result.Error != nil {
		return nil, result.Error
	}
	return stayAreas, nil
}

func GetStayAreaByID(db *gorm.DB, id string) (*model.StayArea, error) {
	var stayArea model.StayArea
	result := db.First(&stayArea, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stayArea, nil
}

func GetStayAreaByName(db *gorm.DB, name string) (*model.StayArea, error) {
	var stayArea model.StayArea
	result := db.Where("name = ?", name).First(&stayArea)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stayArea, nil
}

type AddStayAreaRequest struct {
	Name     string
	Capacity int
}

func AddStayArea(db *gorm.DB, req AddStayAreaRequest) (*model.StayArea, error) {
	stayArea := &model.StayArea{
		Name:     req.Name,
		Capacity: req.Capacity,
	}
	result := db.Create(stayArea)
	if result.Error != nil {
		return nil, result.Error
	}
	return stayArea, nil
}

type GetStayAreaOccupancyResponse struct {
	StayAreaID           string
	StayName             string
	StayCapacity         int
	CurrentOccupiedCount int
}

func GetAllStayAreasWithOccupancy(db *gorm.DB) ([]GetStayAreaOccupancyResponse, error) {
	sql := `
		SELECT
			sa.id,
			sa.name,
			sa.capacity,
			COUNT(v.id) as occupied_count
		FROM stay_areas sa
		LEFT JOIN visits v ON v.stay_area_id = sa.id AND v.status = 'checked-in'
		GROUP BY sa.id, sa.name, sa.capacity
		ORDER BY sa.name
	`

	var results []sqlGetStayAreaOccupancy
	err := db.Raw(sql).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	responses := make([]GetStayAreaOccupancyResponse, 0, len(results))
	for _, result := range results {
		responses = append(responses, GetStayAreaOccupancyResponse{
			StayAreaID:           result.ID,
			StayName:             result.Name,
			StayCapacity:         result.Capacity,
			CurrentOccupiedCount: result.OccupiedCount,
		})
	}

	return responses, nil
}

type sqlGetProfileData struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	PhoneNumber   string               `json:"phone_number"`
	Email         string               `json:"email"`
	Gender        model.Gender         `json:"gender"`
	Category      model.Category       `json:"category"`
	IsBlocked     bool                 `json:"is_blocked"`
	Remarks       string               `json:"remarks"`
	ActiveVisitID *string              `json:"active_visit_id"`
	DepartureDate *time.Time           `json:"departure_date"`
	StayArea      *string              `json:"stay_area"`
	Status        *model.ProfileStatus `json:"status"`
}

type sqlGetStayAreaOccupancy struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Capacity      int    `json:"capacity"`
	OccupiedCount int    `json:"occupied_count"`
}
