package sqlstore

import "context"

//Init squishes all the strings together
func (dbConn *SQLStore) Init(ctx context.Context) error {

	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS `+dbConn.recordTableName+` (

	carID text NOT NULL,
	userID text NOT NULL,
	datStart text NOT NULL,
	dateReturned text, 
	rate float NOT NULL,
	PRIMARY KEY ()	//combination of different fields
	);
`)
}
