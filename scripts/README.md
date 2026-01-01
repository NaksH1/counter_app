# Database Scripts

This directory contains scripts to manage database seeding and clearing.

## Option 1: Using Go Script (Recommended)

### Build and Run the Seed Script

```bash
# From the project root directory
cd scripts
go run seed_db.go
```

This will:
1. Connect to your database using `.env` configuration
2. Clear all existing data (TRUNCATE all tables)
3. Seed SevaTypes (8 types)
4. Seed StayAreas (8 areas with capacities)
5. Seed Lockers (100 lockers)

## Option 2: Using SQL Scripts

### Clear Database

```bash
psql -h localhost -U naksh -d counter_db -f clear_db.sql
```

### Seed Database

```bash
psql -h localhost -U naksh -d counter_db -f seed_db.sql
```

## Option 3: Using psql Interactive

```bash
# Connect to database
psql -h localhost -U naksh -d counter_db

# Run the SQL files
\i clear_db.sql
\i seed_db.sql
```

## Seeded Data

### SevaTypes (8 types)
- Dining
- Kitchen Support
- Counter
- Cleaning
- Security
- Garden
- Transport
- Medical

### StayAreas (8 areas)
- Dormitory A - Men (50 capacity)
- Dormitory B - Men (50 capacity)
- Dormitory C - Women (40 capacity)
- Dormitory D - Women (40 capacity)
- Family Room Block 1 (20 capacity)
- Family Room Block 2 (20 capacity)
- Guest House - VIP (10 capacity)
- Cottage Area (15 capacity)

### Lockers
- 100 lockers (L001 - L100)
- All initially unoccupied

## Database Schema Dependency Chart

```
┌─────────────────────────────────────────────────────────────┐
│                    SEEDING ORDER                             │
└─────────────────────────────────────────────────────────────┘

Level 0 (No Dependencies - Reference Tables)
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  SevaTypes   │  │  StayAreas   │  │   Lockers    │
│              │  │              │  │              │
│ - id         │  │ - id         │  │ - id         │
│ - name       │  │ - name       │  │ - locker_num │
│ - desc       │  │ - capacity   │  │ - is_occupied│
│ - is_active  │  │              │  │              │
└──────────────┘  └──────────────┘  └──────────────┘
       │                  │                  │
       │                  │                  │
       └──────────┬───────┴──────────────────┘
                  │
                  ▼
Level 1 (Depends on: Nothing)
┌──────────────────────────────────┐
│           Profiles               │
│                                  │
│ - id                             │
│ - name, email, phone             │
│ - gender, category               │
│ - is_blocked, remarks            │
└──────────────────────────────────┘
                  │
                  │
                  ▼
Level 2 (Depends on: Profiles, StayAreas, Lockers)
┌──────────────────────────────────┐
│            Visits                │
│                                  │
│ - id                             │
│ - profile_id      ───────────────┼──► Profiles
│ - stay_area_id    ───────────────┼──► StayAreas
│ - locker_id       ───────────────┼──► Lockers
│ - arrival_date                   │
│ - departure_date                 │
│ - status, remarks                │
└──────────────────────────────────┘
                  │
                  │
        ┌─────────┴─────────┐
        │                   │
        ▼                   ▼
Level 3 (Depends on: Profiles, Visits, SevaTypes)
┌──────────────────┐  ┌──────────────────┐
│   Schedules      │  │    Feedbacks     │
│                  │  │                  │
│ - id             │  │ - id             │
│ - profile_id  ───┼──┼──► Profiles     │
│ - visit_id    ───┼──┼──► Visits       │
│ - seva_type_id ──┼──┘                 │
│ - location       │  │ - profile_id  ───┼──► Profiles
│ - date           │  │ - visit_id    ───┼──► Visits
│                  │  │ - content        │
└──────────────────┘  │ - type           │
                      │ - created_by     │
                      └──────────────────┘
```

### Seeding Order

1. **SevaTypes** (8 types) - No dependencies
2. **StayAreas** (8 areas) - No dependencies  
3. **Lockers** (100) - No dependencies
4. **Profiles** (10) - No dependencies
5. **Visits** (6) - Needs: Profiles, StayAreas, Lockers
6. **Schedules** (6) - Needs: Profiles, Visits, SevaTypes
7. **Feedbacks** (5) - Needs: Profiles, Visits (optional)

### Foreign Key Relationships

**Visits:**
- `profile_id` → Profiles.id
- `stay_area_id` → StayAreas.id
- `locker_id` → Lockers.id (nullable)

**Schedules:**
- `profile_id` → Profiles.id
- `visit_id` → Visits.id
- `seva_type_id` → SevaTypes.id

**Feedbacks:**
- `profile_id` → Profiles.id
- `visit_id` → Visits.id (nullable)

## Notes

- The Go script requires the `.env` file to be present in the project root
- The SQL scripts use `gen_random_uuid()` which requires PostgreSQL 13+
- All data will be cleared before seeding (TRUNCATE CASCADE)
- Foreign key constraints are respected during clearing and seeding
- Seeding order ensures all foreign key dependencies are satisfied
