mkdir -p ~/.streamlit/
echo "[server]
secure = false
domain = "localhost"
base = "http://localhost:1234"
host = "localhost"
port = 8080"  >> ~/.streamlit/config.toml

echo "[admin]
admin_id = 0"  >> ~/.streamlit/config.toml

echo "[status]
waitId = 1
acceptId = 2
rejectId = 3
cancelId = 4" >> ~/.streamlit/config.toml

echo "[type]
registrationId = 1
JoinTeamId = 2
LeaveTeamId = 3" >> ~/.streamlit/config.toml

echo "[role]
adminId = 1
managerId = 2
playerId = 3" >> ~/.streamlit/config.toml

echo "[secrets]
jwt = "secret"" >> ~/.streamlit/config.toml

echo "[db]
host = "localhost"
port = 5432
user = "perceval"
password = "password"
name = "group_organiser"
secure = false" >> ~/.streamlit/config.toml
