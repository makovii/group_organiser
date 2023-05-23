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

func TestAdminRepository(t *testing.T) {
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

	t.Run("get admins", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		admins, err := repo.GetAdmins()
		if err != nil {
			t.Errorf("error: %v", err)
		}

		expected := database.User{
			Id: 0,
		}

		var admin database.User
		for _, n := range *admins {
			if n.Id == 0 {
				admin = n
			}
		}

		if !reflect.DeepEqual(admin.Id, expected.Id) {
			t.Errorf("wrong data error. actual: %v, expected: %v", admin.Id, expected.Id)
		}
	})

	t.Run("get user by id", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		admin := database.User{
			Id:       uint(0),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: "1234",
			Ban:      false,
			Role:     1,
		}

		user, err := repo.GetUserById(0)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !reflect.DeepEqual(user, &admin) {
			t.Errorf("wrong data error. actual: %v, expected: %v", user, admin)
		}
	})

	t.Run("get team by id", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		team := repo.GetTeamById(0)

		expected := database.Team{
			Id: 0,
		}

		if !reflect.DeepEqual(team.Id, expected.Id) {
			t.Errorf("wrong data error. actual: %v, expected: %v", team.Id, expected.Id)
		}
	})

	t.Run("get teams", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		teams, err := repo.GetTeams()
		if err != nil {
			t.Errorf("failed to find teams with error: %v", err)
		}

		expected := database.Team{
			Id: 0,
		}

		var team database.Team
		for _, n := range *teams {
			if n.Id == 0 {
				team = n
			}
		}

		if !reflect.DeepEqual(team.Id, expected.Id) {
			t.Errorf("manager Id is corrupted; actual: %v, expected: %v", team.Id, expected.Id)
		}
	})

	t.Run("ban by id", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		admin := database.User{
			Id:       uint(0),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: "1234",
			Ban:      true,
			Role:     1,
		}

		user, err := repo.BanById(0)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !reflect.DeepEqual(user, &admin) {
			t.Errorf("wrong data error. actual: %v, expected: %v", user, admin)
		}
	})

	t.Run("accept manager registration", func(t *testing.T) {
		cfg := config.GetConfig()
		db := database.ConnectDatabase()
		repo := NewAdminRepository(db, cfg)

		admin := database.User{
			Id:       uint(0),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: "1234",
			Ban:      false,
			Role:     1,
		}

		user, err := repo.AcceptManagerRegistration(0)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !reflect.DeepEqual(user, &admin) {
			t.Errorf("wrong data error. actual: %v, expected: %v", user, admin)
		}
	})
}
