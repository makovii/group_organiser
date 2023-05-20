// integration test for user

package repository

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
)

var dbURI string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file, using system env variables")
	}

	dbURI = os.Getenv("DATABASE_URL")
}

func initDB() func() {
	m, err := migrate.New("file://../../migration/", dbURI+"?sslmode=disable")
	if err != nil {
		log.Fatal(err, "1")
	}
	if err := m.Up(); err != nil {
		log.Fatal(err, "2")
	}
	log.Println("db scheme is up to date")

	return func() {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}

func TestUserRepository(t *testing.T) {
	ctx := context.Background()
	resetDB := initDB()
	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() {
		resetDB()
		pool.Close()
	}()


	t.Run("get user", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		user, err := repo.GetUserById(0)

		if err != nil {
			t.Errorf("failed to find user with error: %v", err)
		}

		admin := database.User{
			Id:uint(0),
			Name: "admin",
			Email: "admin@gmail.com",
			Password: "1234",
			Ban: false,
			Role: 1,
		}

		if !reflect.DeepEqual(user, &admin) {
			t.Errorf("user data is corrupted; actual: %v, expected: %v", user, admin)
		}
	})

	t.Run("create request, save request and get reques by 'from'", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		var request database.Request
		request.From = uint(1)
		request.To = uint(0)
		request.StatusId = uint(cfg.Status.WaitId)
		request.TypeId = uint(cfg.Type.JoinTeamId)

		_, errCreate := repo.CreateRequest(request)
		if errCreate != nil {
			t.Errorf("failed to create request with error: %v", err)
		}

		dbRequests, errGet := repo.GetRequestsByFrom(int(request.From))
		if errGet != nil {
			t.Errorf("failed to get request by 'from' with error: %v", err)
		}

		var canceledRequest database.Request
		for _, n := range *dbRequests {
				n.StatusId = uint(cfg.Status.CancelId)
				canceledRequest = n
		}

		requestSaveId, errSave := repo.SaveRequest(canceledRequest)
		if errSave != nil {
			t.Errorf("failed to save request with error: %v", err)
		}

		dbRequests2, errGet := repo.GetRequestsByFrom(int(canceledRequest.From))
		if errGet != nil {
			t.Errorf("failed to get request by 'from' with error: %v", err)
		}

		var savedRequest database.Request
		for _, n := range *dbRequests2 {
				n.StatusId = uint(cfg.Status.CancelId)
				savedRequest = n
		}


		if !reflect.DeepEqual(savedRequest.Id, requestSaveId) {
			t.Errorf("request data is corrupted; actual: %v, expected: %v", savedRequest.Id, requestSaveId)
		}
	})


	t.Run("get all managers", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		managers, err := repo.GetAllManagers()
		if err != nil {
			t.Errorf("failed to find managers with error: %v", err)
		}

		expected :=  database.User{
			Id: 1,
		}

		var manager database.User
		for _, n := range *managers {
			if 1 == n.Id{
				manager = n
			}
		}

		if !reflect.DeepEqual(manager.Id, expected.Id) {
			t.Errorf("manager Id is corrupted; actual: %v, expected: %v", manager.Id, expected.Id)
		}
	})

	t.Run("get all teams", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		teams, err := repo.GetAllTeams()
		if err != nil {
			t.Errorf("failed to find teams with error: %v", err)
		}

		expected :=  database.Team{
			Id: 1,
		}

		var team database.Team
		for _, n := range *teams {
			if 1 == n.Id{
				team = n
			}
		}

		if !reflect.DeepEqual(team.Id, expected.Id) {
			t.Errorf("manager Id is corrupted; actual: %v, expected: %v", team.Id, expected.Id)
		}
	})


	t.Run("get team by id", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		team := repo.GetTeamById(1)

		expected :=  database.Team{
			Id: 1,
		}

		if !reflect.DeepEqual(team.Id, expected.Id) {
			t.Errorf("manager Id is corrupted; actual: %v, expected: %v", team.Id, expected.Id)
		}
	})

	t.Run("get notifications", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewUserRepository(db, cfg)

		var request database.Request
		request.From = uint(1)
		request.To = uint(0)
		request.StatusId = uint(cfg.Status.WaitId)
		request.TypeId = uint(cfg.Type.JoinTeamId)

		_, errCreate := repo.CreateRequest(request)
		if errCreate != nil {
			t.Errorf("failed to create request with error: %v", err)
		}


		dbRequests, errGet := repo.GetNotifications(int(request.To))
		if errGet != nil {
			t.Errorf("failed to get request by 'from'notifications with error: %v", err)
		}

		var savedRequest database.Request
		for _, n := range *dbRequests {
				savedRequest = n
		}


		if !reflect.DeepEqual(savedRequest.To, request.To) {
			t.Errorf("request 'To' data is corrupted; actual: %v, expected: %v", savedRequest.To, request.To)
		}
		if !reflect.DeepEqual(savedRequest.From, request.From) {
			t.Errorf("request 'From' data is corrupted; actual: %v, expected: %v", savedRequest.From, request.From)
		}
		if !reflect.DeepEqual(savedRequest.StatusId, request.StatusId) {
			t.Errorf("request 'StatusId' data is corrupted; actual: %v, expected: %v", savedRequest.StatusId, request.StatusId)
		}
		if !reflect.DeepEqual(savedRequest.TypeId, request.TypeId) {
			t.Errorf("request 'TypeId' data is corrupted; actual: %v, expected: %v", savedRequest.TypeId, request.TypeId)
		}
	})
}
