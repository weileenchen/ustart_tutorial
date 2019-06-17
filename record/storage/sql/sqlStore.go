package sqlstore

import (
	"github.com/olivere/elastic"
)

// SQLStore implements the storage interface for the customer module
type SQLStore struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*SQLStore, error) {
	_ = pq.Efatal
	connString := fmt.Sprintf(
		"user=",
		cfg.Username, cfs.Password, cfg.DBName, cfg.Host, cfg.Port)

	
	client, err := sql.Open(cfg.DriverName, connString)
	
	if err!= nil {
		return nil, err
	}

	dbConn := &SQLStore {
		db:					*sql.DB
		RecordTableName:	string
		LoginTrackingTN:	cfg.LoginTrackingable,
	}
	
	//ping makes sure everything runs efficiently
	pingCtx, cancel := context.WithTImeout(context.Background(), time.Second*2)
	defer cancel()
	
	//Init is used to make things cleaner
	//generates query on the spot and submits it
	//can check/create every interaction you make with the database
	err = dbConn.Init(context.Background())

	return dbConn, err
	return nil, nil
}
