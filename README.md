# Ashram Connect API

Backend API for managing volunteer profiles, visits, schedules, lockers, and feedback at ashram facilities.

## 🌟 Features

- **Profile Management**: Create and manage volunteer profiles with blocking capability
- **Visit Tracking**: Track volunteer visits with stay area assignments and status management
- **Schedule Management**: Assign volunteers to seva types with date-based scheduling
- **Locker Allocation**: Manage locker assignments with section-based organization
- **Feedback System**: Collect and categorize feedback (Positive/Negative/Neutral)
- **Occupancy Tracking**: Real-time stay area capacity and occupancy monitoring
- **Business Rules**: 
  - One active visit per profile at any time
  - Automatic capacity validation for stay areas
  - Duplicate schedule prevention

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- **ORM**: [GORM](https://gorm.io/) - Database ORM
- **Database**: PostgreSQL 14+
- **API Documentation**: OpenAPI 3.0 specification
- **Deployment**: Fly.io with Docker

## 📁 Project Structure

```
counter_app/
├── api/
│   └── openapi.yml          # OpenAPI 3.0 specification
├── internal/
│   ├── config/              # Configuration management
│   ├── dao/                 # Data Access Objects
│   ├── handler/             # HTTP request handlers
│   ├── model/               # Database models
│   └── util/                # Utility functions
├── scripts/
│   ├── seed_db.go           # Database seeding script
│   └── clear_db.sql         # Database cleanup script
├── server/
│   ├── api/
│   │   └── router.go        # Route definitions
│   └── main.go              # Application entry point
├── .env.example             # Environment variables template
├── Dockerfile               # Multi-stage Docker build
├── fly.toml                 # Fly.io deployment config
└── go.mod                   # Go module dependencies
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Git

### Local Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/counter_app.git
   cd counter_app
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Create database**
   ```bash
   createdb counter_app
   ```

5. **Run the application**
   ```bash
   go run server/main.go
   ```
   
   The server will start on `http://localhost:8080`

6. **Seed the database (optional)**
   ```bash
   go run scripts/seed_db.go
   ```

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | `counter_app` |
| `PORT` | Server port | `8080` |

## 📚 API Documentation

The complete API documentation is available in the [OpenAPI specification](./api/openapi.yml).

### Key Endpoints

#### Profiles
- `GET /api/profiles` - Get all profiles with active visit info
- `POST /api/profiles` - Create a new profile
- `PUT /api/profiles/:id` - Update profile details
- `DELETE /api/profiles/:id` - Delete a profile

#### Visits
- `POST /api/visits` - Create a new visit (checks capacity)
- `PUT /api/visits/:id` - Update visit details
- `DELETE /api/visits/:id` - Delete a visit

#### Schedules
- `GET /api/schedules?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD` - Get schedules for date range
- `POST /api/schedules` - Create a new schedule
- `PUT /api/schedules/:id` - Update schedule
- `DELETE /api/schedules/:id` - Delete schedule

#### Stay Areas
- `GET /api/stay-areas` - Get all stay areas
- `GET /api/stay-areas/occupancy` - Get occupancy details for all stay areas
- `POST /api/stay-areas` - Create a new stay area
- `PUT /api/stay-areas/:id` - Update stay area
- `DELETE /api/stay-areas/:id` - Delete stay area

#### Seva Types
- `GET /api/seva-types` - Get all seva types
- `POST /api/seva-types` - Create a new seva type
- `PUT /api/seva-types/:id` - Update seva type
- `DELETE /api/seva-types/:id` - Delete seva type

#### Lockers
- `GET /api/lockers` - Get all lockers

#### Feedbacks
- `POST /api/feedbacks` - Submit feedback
- `GET /api/feedbacks` - Get all feedbacks

## 🗄️ Database Schema

### Core Models

- **Profile**: Volunteer information (name, email, phone, gender, category)
- **Visit**: Visit tracking with stay area and locker assignment
- **Schedule**: Seva assignments with date and location
- **StayArea**: Accommodation areas with capacity management
- **SevaType**: Types of seva activities
- **Locker**: Locker inventory with section organization
- **Feedback**: Feedback entries linked to profiles and visits

See [FRONTEND_INTEGRATION_GUIDE.md](./docs/FRONTEND_INTEGRATION_GUIDE.md) for detailed schema information.

## 🚢 Deployment

### Deploy to Fly.io

1. **Install Fly.io CLI**
   ```bash
   brew install flyctl
   ```

2. **Login to Fly.io**
   ```bash
   flyctl auth login
   ```

3. **Launch the app**
   ```bash
   flyctl launch
   # Follow prompts to create app and PostgreSQL database
   ```

4. **Set environment secrets**
   ```bash
   flyctl secrets set DB_HOST=your-db-host.fly.dev
   flyctl secrets set DB_PORT=5432
   flyctl secrets set DB_USER=postgres
   flyctl secrets set DB_PASSWORD=your-password
   flyctl secrets set DB_NAME=your-db-name
   ```

5. **Deploy**
   ```bash
   flyctl deploy
   ```

6. **View logs**
   ```bash
   flyctl logs
   ```

## 📝 Development Notes


## 📄 License

This project is licensed under the MIT License.

## 🙏 Acknowledgments

- Built with [Gin](https://github.com/gin-gonic/gin)
- Database ORM by [GORM](https://gorm.io/)
- Deployed on [Fly.io](https://fly.io/)