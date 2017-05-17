package epic

import (
	"github.com/pkg/errors"
)

type About struct {
	Name	string
	Version	string
	Website string
}

func ServeHTTPS(user string, password string, server string, database string) error {
	return errors.New("Not yet implimented in Epic.")
}

func ServeHTTP(port int, host string, password string, server string, database string) {
	serveHTTP(port, host, "epic", password, server, database)
}

func ServeAll(user string, password string, server string, database string) error {
	return errors.New("Not yet implimented in Epic.")
}

func ServeRedirect(user string, password string, server string, database string) error {
	return errors.New("Not yet implimented in Epic.")
}

func InstallPostgreSQLDatabase(adminUser string, adminPassword string, server string, database string, epicPassword string) (string, error) {

	/*
	err := dropPostgreSQLSchemaTablesAndConstraints(adminUser, adminPassword, server)
	if err != nil {
		return err
	}
	err = dropPostgreSQLUserAndDatabase(adminUser, adminPassword, server)
	if err != nil {
		return err
	}
	*/
	err := createPostgreSQLUserAndDatabase(adminUser, adminPassword, server, database, epicPassword)
	if err != nil {
		return "", err
	}
	err = createPostgreSQLSchemaTablesAndConstraints(epicPassword, server, database)
	if err != nil {
		return "", err
	}
	appID, err := createPostgreSQLEpicAppAdminUser(epicPassword, server, database)
	if err != nil {
		return "", err
	}
	return appID, nil
}

func UninstallPostgreSQLDatabase(adminUser string, adminPassword string, server string, database string) error {

	err := dropPostgreSQLSchemaTablesAndConstraints(adminUser, adminPassword, server, database)
	if err != nil {
		return err
	}
	err = dropPostgreSQLUserAndDatabase(adminUser, adminPassword, server, database)
	if err != nil {
		return err
	}
	return nil

}