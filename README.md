Execute the following command at the terminal in the folder where the cockroach.exe file is located

./cockroach.exe start-single-node --insecure --store=json-test --listen-addr=localhost:26257 --http-addr=localhost:8080 


Open another terminal and run the following lines

./cockroach.exe sql --insecure --host=localhost:26257
CREATE USER IF NOT EXISTS joe;
CREATE DATABASE androidqueries;
GRANT ALL ON DATABASE androidqueries TO joe;
SET DATABASE = androidqueries
CREATE TABLE IF NOT EXISTS domain (host STRING PRIMARY KEY, sslGrade STRING, sslPreviousGrade STRING, lastSearch TIMESTAMPTZ);
CREATE TABLE IF NOT EXISTS history (host STRING PRIMARY KEY);

Import the go packages to be able to run the go run ./Server.go

Install the apk in the cell phone to have the user interface