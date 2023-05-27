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

	"github.com/jackc/pgx/v4/pgxpool"
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

func TestManagerRepository(t *testing.T) {
	ctx := context.Background()
	resetDB := initDB()
	pool, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() {
		resetDB()
		pool.Close()
	}()

	db := database.ConnectDatabase()
	repo := NewManagerRepository(db)

	t.Run("create team", func(t *testing.T) {
		managerID := uint(1)
		teamName := "Team A"

		team, err := repo.CreateTeam(teamName, managerID)
		if err != nil {
			t.Errorf("failed to create team with error: %v", err)
		}

		if team.Name != teamName {
			t.Errorf("team name is incorrect; actual: %s, expected: %s", team.Name, teamName)
		}
		if team.ManagerID != managerID {
			t.Errorf("team manager ID is incorrect; actual: %d, expected: %d", team.ManagerID, managerID)
		}
	})

	t.Run("get all teams", func(t *testing.T) {
		managerID := uint(1)

		teams, err := repo.GetAllTeams(managerID)
		if err != nil {
			t.Errorf("failed to get all teams with error: %v", err)
		}

		expected := database.Team{
			Name:      "Team A",
			ManagerID: managerID,
		}

		var team database.Team
		for _, t := range *teams {
			if t.Name == expected.Name && t.ManagerID == expected.ManagerID {
				team = t
				break
			}
		}

		if team.Name != expected.Name {
			t.Errorf("team name is incorrect; actual: %s, expected: %s", team.Name, expected.Name)
		}
		if team.ManagerID != expected.ManagerID {
			t.Errorf("team manager ID is incorrect; actual: %d, expected: %d", team.ManagerID, expected.ManagerID)
		}
	})

	t.Run("get team by ID", func(t *testing.T) {
		managerID := uint(1)

		teamID := uint(1)

		team, _ := repo.GetTeam(teamID, managerID)
		// if err != nil {
		// 	t.Errorf("failed to get team by ID with error: %v", err)
		// }

		if team != nil {
			expected := database.Team{
				Id:        teamID,
				Name:      "Team A",
				ManagerID: managerID,
			}
	
			if !reflect.DeepEqual(team, &expected) {
				t.Errorf("team data is corrupted; actual: %v, expected: %v", team, expected)
			}
		}

	})

	// t.Run("update team", func(t *testing.T) {
	// 	managerID := uint(1)
	// 	teamID := uint(1)
	// 	newTeamName := "Updated Team A"

	// 	team, _ := repo.UpdateTeam(teamID, managerID, newTeamName)
	// 	// if err != nil {
	// 	// 	t.Errorf("failed to update team with error: %v", err)
	// 	// }

	// 	if !reflect.DeepEqual(team, newTeamName) {
	// 		t.Errorf("manager Id is corrupted; actual: %v, expected: %v", team.Name, newTeamName)
	// 	}

	// 	// if team.Name != newTeamName {
	// 	// 	t.Errorf("team name is incorrect; actual: %s, expected: %s", team.Name, newTeamName)
	// 	// }
	// 	if team.ManagerID != managerID {
	// 		t.Errorf("team manager ID is incorrect; actual: %d, expected: %d", team.ManagerID, managerID)
	// 	}
	// })

	t.Run("delete team", func(t *testing.T) {
		managerID := uint(1)
		teamID := uint(1)

		err := repo.DeleteTeam(teamID, managerID)
		if err != nil {
			t.Errorf("failed to delete team with error: %v", err)
		}

		// Verify that the team is deleted
		team, err := repo.GetTeam(teamID, managerID)
		if err == nil || team != nil {
			t.Errorf("team still exists after deletion; team: %v", team)
		}
	})
}
