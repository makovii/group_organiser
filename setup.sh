mkdir -p ~/.streamlit/
echo "[server]
secure = false
domain = "localhost"
base = "http://localhost:1234"
host = "localhost"
port = 8080"  >> ~/.config/config.toml

echo "[admin]
admin_id = 0"  >> ~/.config/config.toml

echo "[status]
waitId = 1
acceptId = 2
rejectId = 3
cancelId = 4" >> ~/.config/config.toml

echo "[type]
registrationId = 1
JoinTeamId = 2
LeaveTeamId = 3" >> ~/.config/config.toml

echo "[role]
adminId = 1
managerId = 2
playerId = 3" >> ~/.config/config.toml

echo "[secrets]
jwt = "secret"" >> ~/.config/config.toml

echo "[db]
host = "localhost"
port = 5432
user = "perceval"
password = "password"
name = "group_organiser"
secure = false" >> ~/config/config.toml
