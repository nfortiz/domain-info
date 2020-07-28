package resources

import (
	"github.com/jackc/pgx"
	"log"
	"time"
	"truora-server/models"
)

type DataBase struct {
	Conn *pgx.Conn
}

const URI = "postgresql://root@DESKTOP-SQKENNG:26257/challenge?sslmode=disable"

func Connect() DataBase {
	config, err := pgx.ParseURI(URI)
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	// Connect to the "bank" database.
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	initDB(conn)
	db := DataBase{Conn:conn}
	return db
}

func initDB(conn *pgx.Conn) {
	// Create Domain Table
	if _, err := conn.Exec(`CREATE TABLE IF NOT EXISTS domains 
		(domain_name VARCHAR(50) PRIMARY KEY, 
		servers_changed BOOL DEFAULT FALSE, 
		ssl_grade VARCHAR(2), 
		previous_ssl_grade VARCHAR(2), 
		logo VARCHAR(200), 
		title VARCHAR(100), 
		is_down BOOL,
		date_checked TIMESTAMP NOT NULL DEFAULT now())
		`); err != nil {
		log.Fatal(err)
	}

	// Create Server Table
	if _, err := conn.Exec(`CREATE TABLE IF NOT EXISTS servers 
		(id SERIAL PRIMARY KEY,
		address VARCHAR(40),
		ssl_grade VARCHAR(2),
		country VARCHAR(10),
		owner VARCHAR(50),
		domain VARCHAR(50),
		FOREIGN KEY (domain) REFERENCES domains(domain_name)
		)
		`); err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) CreateDomain(newDomain models.Domain, domainName string) error {
	if _, err := db.Conn.Exec(`INSERT INTO domains 
			(domain_name, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down)
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
			domainName,
			newDomain.ServersChanged,
			newDomain.SSLGrade,
			newDomain.PreviousSSLGrade,
			newDomain.Logo,
			newDomain.Title,
			newDomain.IsDown); err != nil {
		return err
	}
	return nil
}

func (db *DataBase) CreateServer(newServer models.Server, domainName string) error {
	if _, err := db.Conn.Exec(`INSERT INTO servers 
			(address, ssl_grade, country, owner, domain)
			VALUES($1, $2, $3, $4, $5)`,
		newServer.Address,
		newServer.SSLGrade,
		newServer.Country,
		newServer.Owner,
		domainName); err != nil {
		return  err
	}
	return nil
}

func (db *DataBase) GetDomains () ([]models.Domain, error) {
	var domains []models.Domain
	rows, err := db.Conn.Query(`SELECT
		domain_name,
		servers_changed, 
		ssl_grade, 
		previous_ssl_grade, 
		logo, 
		title, 
		is_down,
		date_checked
		FROM domains
		ORDER BY date_checked DESC`)
	if err != nil {
		return domains, err
	}
	defer rows.Close()
	//ORDER BY date_checked`
	for rows.Next() {
		var domain models.Domain
		if err := rows.Scan(
			&domain.Name,
			&domain.ServersChanged,
			&domain.SSLGrade,
			&domain.PreviousSSLGrade,
			&domain.Logo,
			&domain.Title,
			&domain.IsDown,
			&domain.DateChecked); err != nil {
			return domains, err
		}
		domains = append(domains, domain)
	}
	return domains, nil
}


func (db *DataBase) GetDomainByName (domainName string) (models.Domain, error) {
	var domain models.Domain
	if err := db.Conn.QueryRow(`SELECT 
		servers_changed, 
		ssl_grade, 
		previous_ssl_grade, 
		logo, 
		title, 
		is_down,
		date_checked
		FROM domains
		WHERE domain_name=$1`, domainName).Scan(
		&domain.ServersChanged,
		&domain.SSLGrade,
		&domain.PreviousSSLGrade,
		&domain.Logo,
		&domain.Title,
		&domain.IsDown,
		&domain.DateChecked); err != nil {
		return domain, err
	}
	return domain, nil
}

func (db *DataBase) UpdateDomainChecked(serversChanged bool, domainName string) error {
	if _, err := db.Conn.Exec(
		"UPDATE domains SET servers_changed = $1 WHERE domain_name = $2",
		serversChanged, domainName); err != nil {
		return err
	}

	return nil
}

func (db *DataBase) UpdateDomainServersChanged(domainName string) error {
	if _, err := db.Conn.Exec(
		"UPDATE domains SET date_checked = $1 WHERE domain_name = $2", time.Now(), domainName); err != nil {
		return err
	}

	return nil
}

func (db *DataBase) GetDomainServers(domainName string) ([]models.Server, error) {
	var servers []models.Server
	rows, err := db.Conn.Query("SELECT address, ssl_grade, country, owner FROM servers WHERE domain = $1", domainName)
	if err != nil {
		return servers, err
	}
	defer rows.Close()
	for rows.Next() {
		var server models.Server
		if err := rows.Scan(&server.Address, &server.SSLGrade, &server.Country, &server.Owner); err != nil {
			return servers, err
		} else {
			servers = append(servers, server)
		}
	}

	if rows.Err() != nil {
		return servers, err
	}

	return servers, nil
}

func(db *DataBase) GetServerByAddress(address string) models.Server {
	var server models.Server
	if err := db.Conn.QueryRow("SELECT address, ssl_grade, country, owner FROM servers WHERE address = $1", address).Scan(
		&server.Address,
		&server.SSLGrade,
		&server.Country,
		&server.Owner); err != nil {
		log.Fatal(err)
	}

	return server
}