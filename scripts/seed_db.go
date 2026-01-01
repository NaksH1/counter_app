package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"counterapp/internal/model"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	log.Println("Clearing existing data...")
	if err := clearDatabase(db); err != nil {
		log.Fatalf("Failed to clear database: %v", err)
	}

	log.Println("Seeding SevaTypes...")
	if err := seedSevaTypes(db); err != nil {
		log.Fatalf("Failed to seed SevaTypes: %v", err)
	}

	log.Println("Seeding StayAreas...")
	if err := seedStayAreas(db); err != nil {
		log.Fatalf("Failed to seed StayAreas: %v", err)
	}

	log.Println("Seeding Lockers...")
	if err := seedLockers(db); err != nil {
		log.Fatalf("Failed to seed Lockers: %v", err)
	}

	log.Println("Seeding Sample Profiles...")
	profiles, err := seedProfiles(db)
	if err != nil {
		log.Fatalf("Failed to seed Profiles: %v", err)
	}

	log.Println("Seeding Sample Visits...")
	visits, err := seedVisits(db, profiles)
	if err != nil {
		log.Fatalf("Failed to seed Visits: %v", err)
	}

	log.Println("Seeding Sample Schedules...")
	if err := seedSchedules(db, profiles, visits); err != nil {
		log.Fatalf("Failed to seed Schedules: %v", err)
	}

	log.Println("Seeding Sample Feedbacks...")
	if err := seedFeedbacks(db, profiles, visits); err != nil {
		log.Fatalf("Failed to seed Feedbacks: %v", err)
	}

	log.Println("Database seeded successfully!")
}

func clearDatabase(db *gorm.DB) error {
	tables := []string{
		"schedules",
		"feedbacks",
		"visits",
		"profiles",
		"lockers",
		"stay_areas",
		"seva_types",
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)).Error; err != nil {
			return fmt.Errorf("failed to truncate %s: %w", table, err)
		}
		log.Printf("  Cleared table: %s", table)
	}

	return nil
}

func seedSevaTypes(db *gorm.DB) error {
	sevaTypes := []model.SevaType{
		{
			Name:        "Dining",
			Description: strPtr("Food service and dining hall seva"),
			IsActive:    true,
		},
		{
			Name:        "Kitchen Support",
			Description: strPtr("Kitchen preparation and cleaning seva"),
			IsActive:    true,
		},
		{
			Name:        "Counter",
			Description: strPtr("Reception and counter service seva"),
			IsActive:    true,
		},
		{
			Name:        "Cleaning",
			Description: strPtr("General cleaning and maintenance seva"),
			IsActive:    true,
		},
		{
			Name:        "Security",
			Description: strPtr("Security and monitoring seva"),
			IsActive:    true,
		},
		{
			Name:        "Garden",
			Description: strPtr("Gardening and landscaping seva"),
			IsActive:    true,
		},
		{
			Name:        "Transport",
			Description: strPtr("Transportation and logistics seva"),
			IsActive:    true,
		},
		{
			Name:        "Medical",
			Description: strPtr("Medical assistance and first aid seva"),
			IsActive:    true,
		},
	}

	for _, sevaType := range sevaTypes {
		if err := db.Create(&sevaType).Error; err != nil {
			return fmt.Errorf("failed to create SevaType %s: %w", sevaType.Name, err)
		}
		log.Printf("  Created SevaType: %s", sevaType.Name)
	}

	return nil
}

func seedStayAreas(db *gorm.DB) error {
	stayAreas := []model.StayArea{
		{Name: "Dormitory A - Men", Capacity: 50},
		{Name: "Dormitory B - Men", Capacity: 50},
		{Name: "Dormitory C - Women", Capacity: 40},
		{Name: "Dormitory D - Women", Capacity: 40},
		{Name: "Family Room Block 1", Capacity: 20},
		{Name: "Family Room Block 2", Capacity: 20},
		{Name: "Guest House - VIP", Capacity: 10},
		{Name: "Cottage Area", Capacity: 15},
	}

	for _, stayArea := range stayAreas {
		if err := db.Create(&stayArea).Error; err != nil {
			return fmt.Errorf("failed to create StayArea %s: %w", stayArea.Name, err)
		}
		log.Printf("  Created StayArea: %s (Capacity: %d)", stayArea.Name, stayArea.Capacity)
	}

	return nil
}

func seedLockers(db *gorm.DB) error {
	for i := 1; i <= 100; i++ {
		section := "A"
		if i > 50 {
			section = "B"
		}
		locker := model.Locker{
			LockerNumber: fmt.Sprintf("L%03d", i),
			Section:      section,
			IsOccupied:   false,
		}
		if err := db.Create(&locker).Error; err != nil {
			return fmt.Errorf("failed to create Locker %s: %w", locker.LockerNumber, err)
		}
	}
	log.Printf("  Created 100 lockers (L001 - L100)")

	return nil
}

func seedProfiles(db *gorm.DB) ([]model.Profile, error) {
	profiles := []model.Profile{
		{
			Name:        "Rajesh Kumar",
			Email:       "rajesh.kumar@example.com",
			PhoneNumber: "+91-9876543210",
			Gender:      model.GenderMale,
			Category:    model.CategorySTV,
			IsBlocked:   false,
			Remarks:     strPtr("Regular volunteer, very dedicated"),
		},
		{
			Name:        "Priya Sharma",
			Email:       "priya.sharma@example.com",
			PhoneNumber: "+91-9876543211",
			Gender:      model.GenderFemale,
			Category:    model.CategoryLTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Amit Patel",
			Email:       "amit.patel@example.com",
			PhoneNumber: "+91-9876543212",
			Gender:      model.GenderMale,
			Category:    model.CategorySTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Sneha Reddy",
			Email:       "sneha.reddy@example.com",
			PhoneNumber: "+91-9876543213",
			Gender:      model.GenderFemale,
			Category:    model.CategoryOverseas,
			IsBlocked:   false,
			Remarks:     strPtr("From USA, first time volunteer"),
		},
		{
			Name:        "Vikram Singh",
			Email:       "vikram.singh@example.com",
			PhoneNumber: "+91-9876543214",
			Gender:      model.GenderMale,
			Category:    model.CategoryLTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Ananya Iyer",
			Email:       "ananya.iyer@example.com",
			PhoneNumber: "+91-9876543215",
			Gender:      model.GenderFemale,
			Category:    model.CategorySTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Karthik Menon",
			Email:       "karthik.menon@example.com",
			PhoneNumber: "+91-9876543216",
			Gender:      model.GenderMale,
			Category:    model.CategorySTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Deepa Nair",
			Email:       "deepa.nair@example.com",
			PhoneNumber: "+91-9876543217",
			Gender:      model.GenderFemale,
			Category:    model.CategoryLTV,
			IsBlocked:   false,
			Remarks:     strPtr("Kitchen seva specialist"),
		},
		{
			Name:        "Arjun Desai",
			Email:       "arjun.desai@example.com",
			PhoneNumber: "+91-9876543218",
			Gender:      model.GenderMale,
			Category:    model.CategoryOverseas,
			IsBlocked:   false,
			Remarks:     nil,
		},
		{
			Name:        "Meera Joshi",
			Email:       "meera.joshi@example.com",
			PhoneNumber: "+91-9876543219",
			Gender:      model.GenderFemale,
			Category:    model.CategorySTV,
			IsBlocked:   false,
			Remarks:     nil,
		},
	}

	createdProfiles := make([]model.Profile, 0, len(profiles))
	for _, profile := range profiles {
		if err := db.Create(&profile).Error; err != nil {
			return nil, fmt.Errorf("failed to create Profile %s: %w", profile.Name, err)
		}
		createdProfiles = append(createdProfiles, profile)
		log.Printf("  Created Profile: %s (%s)", profile.Name, profile.Category)
	}

	return createdProfiles, nil
}

func seedVisits(db *gorm.DB, profiles []model.Profile) ([]model.Visit, error) {
	var stayAreas []model.StayArea
	if err := db.Find(&stayAreas).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch stay areas: %w", err)
	}

	var lockers []model.Locker
	if err := db.Limit(10).Find(&lockers).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch lockers: %w", err)
	}

	now := time.Now()
	visits := []model.Visit{
		{
			ProfileID:     profiles[0].ID,
			ArrivalDate:   now.AddDate(0, 0, -5),
			DepartureDate: timePtr(now.AddDate(0, 0, 5)),
			StayAreaID:    stayAreas[0].ID,
			Status:        model.StatusCheckedIn,
			LockerID:      &lockers[0].ID,
			Remarks:       strPtr("First visit, very enthusiastic"),
		},
		{
			ProfileID:     profiles[1].ID,
			ArrivalDate:   now.AddDate(0, 0, -3),
			DepartureDate: timePtr(now.AddDate(0, 0, 7)),
			StayAreaID:    stayAreas[2].ID,
			Status:        model.StatusCheckedIn,
			LockerID:      &lockers[1].ID,
			Remarks:       nil,
		},
		{
			ProfileID:     profiles[2].ID,
			ArrivalDate:   now.AddDate(0, 0, -2),
			DepartureDate: timePtr(now.AddDate(0, 0, 3)),
			StayAreaID:    stayAreas[0].ID,
			Status:        model.StatusCheckedIn,
			LockerID:      &lockers[2].ID,
			Remarks:       nil,
		},
		{
			ProfileID:     profiles[3].ID,
			ArrivalDate:   now.AddDate(0, 0, -10),
			DepartureDate: timePtr(now.AddDate(0, 0, -3)),
			StayAreaID:    stayAreas[4].ID,
			Status:        model.StatusCheckedOut,
			LockerID:      nil,
			Remarks:       strPtr("Completed stay, excellent feedback"),
		},
		{
			ProfileID:     profiles[4].ID,
			ArrivalDate:   now.AddDate(0, 0, -1),
			DepartureDate: timePtr(now.AddDate(0, 0, 14)),
			StayAreaID:    stayAreas[1].ID,
			Status:        model.StatusCheckedIn,
			LockerID:      &lockers[3].ID,
			Remarks:       nil,
		},
		{
			ProfileID:     profiles[5].ID,
			ArrivalDate:   now.AddDate(0, 0, -7),
			DepartureDate: timePtr(now.AddDate(0, 0, 2)),
			StayAreaID:    stayAreas[2].ID,
			Status:        model.StatusCheckedIn,
			LockerID:      &lockers[4].ID,
			Remarks:       nil,
		},
	}

	createdVisits := make([]model.Visit, 0, len(visits))
	for i, visit := range visits {
		if err := db.Create(&visit).Error; err != nil {
			return nil, fmt.Errorf("failed to create Visit for %s: %w", profiles[i].Name, err)
		}
		createdVisits = append(createdVisits, visit)
		log.Printf("  Created Visit: %s - %s (Status: %s)", profiles[i].Name, visit.StayAreaID, visit.Status)
	}

	if err := db.Model(&model.Locker{}).Where("id IN ?", []uuid.UUID{
		lockers[0].ID, lockers[1].ID, lockers[2].ID, lockers[3].ID, lockers[4].ID,
	}).Update("is_occupied", true).Error; err != nil {
		return nil, fmt.Errorf("failed to update locker status: %w", err)
	}

	return createdVisits, nil
}

func seedSchedules(db *gorm.DB, profiles []model.Profile, visits []model.Visit) error {
	var sevaTypes []model.SevaType
	if err := db.Find(&sevaTypes).Error; err != nil {
		return fmt.Errorf("failed to fetch seva types: %w", err)
	}

	now := time.Now()
	schedules := []model.Schedule{
		{
			ProfileID:  profiles[0].ID,
			VisitID:    visits[0].ID,
			SevaTypeID: sevaTypes[0].ID,
			Location:   strPtr("Main Dining Hall"),
			Date:       now.AddDate(0, 0, 1),
		},
		{
			ProfileID:  profiles[0].ID,
			VisitID:    visits[0].ID,
			SevaTypeID: sevaTypes[1].ID,
			Location:   strPtr("Kitchen Block A"),
			Date:       now.AddDate(0, 0, 2),
		},
		{
			ProfileID:  profiles[1].ID,
			VisitID:    visits[1].ID,
			SevaTypeID: sevaTypes[2].ID,
			Location:   strPtr("Reception Counter"),
			Date:       now.AddDate(0, 0, 1),
		},
		{
			ProfileID:  profiles[2].ID,
			VisitID:    visits[2].ID,
			SevaTypeID: sevaTypes[3].ID,
			Location:   strPtr("Dormitory Area"),
			Date:       now.AddDate(0, 0, 1),
		},
		{
			ProfileID:  profiles[4].ID,
			VisitID:    visits[4].ID,
			SevaTypeID: sevaTypes[4].ID,
			Location:   strPtr("Main Gate"),
			Date:       now.AddDate(0, 0, 2),
		},
		{
			ProfileID:  profiles[5].ID,
			VisitID:    visits[5].ID,
			SevaTypeID: sevaTypes[5].ID,
			Location:   strPtr("Garden Area"),
			Date:       now.AddDate(0, 0, 1),
		},
	}

	for _, schedule := range schedules {
		if err := db.Create(&schedule).Error; err != nil {
			return fmt.Errorf("failed to create Schedule: %w", err)
		}
		log.Printf("  Created Schedule: Profile %s - SevaType %s on %s",
			schedule.ProfileID, schedule.SevaTypeID, schedule.Date.Format("2006-01-02"))
	}

	return nil
}

func seedFeedbacks(db *gorm.DB, profiles []model.Profile, visits []model.Visit) error {
	feedbacks := []model.Feedback{
		{
			ProfileID: profiles[0].ID,
			VisitID:   &visits[0].ID,
			Content:   "Great experience! The facilities are excellent and the atmosphere is very peaceful.",
			Type:      model.TypePositive,
			CreatedBy: strPtr("Admin"),
		},
		{
			ProfileID: profiles[1].ID,
			VisitID:   &visits[1].ID,
			Content:   "The food quality could be improved. Otherwise everything is good.",
			Type:      model.TypeNeutral,
			CreatedBy: strPtr("Admin"),
		},
		{
			ProfileID: profiles[2].ID,
			VisitID:   &visits[2].ID,
			Content:   "Wonderful seva opportunity. Very well organized.",
			Type:      model.TypePositive,
			CreatedBy: strPtr("Admin"),
		},
		{
			ProfileID: profiles[3].ID,
			VisitID:   &visits[3].ID,
			Content:   "Had some issues with locker assignment initially, but resolved quickly.",
			Type:      model.TypeNeutral,
			CreatedBy: strPtr("Admin"),
		},
		{
			ProfileID: profiles[4].ID,
			VisitID:   nil,
			Content:   "General inquiry about upcoming programs.",
			Type:      model.TypeNeutral,
			CreatedBy: strPtr("System"),
		},
	}

	for _, feedback := range feedbacks {
		if err := db.Create(&feedback).Error; err != nil {
			return fmt.Errorf("failed to create Feedback: %w", err)
		}
		log.Printf("  Created Feedback: Profile %s - Type %s", feedback.ProfileID, feedback.Type)
	}

	return nil
}

func strPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}
