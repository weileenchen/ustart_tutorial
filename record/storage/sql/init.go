package sqlstore

import "context"

func (dbConn *SQLStore) Init(ctx context.Context) error {

	_, err := dbConn.db.QueryCOntext(ctx, 'CREATE TABLE IF NOT EXISTS -"db.Conn.recordTableName"-

	carID text NOT NULL,
	userID text NOT NULL,
	datStart text NOT NULL,
	dateReturned text, 
	rate float NOT NULL,
	...
	PRIMARY KEY (uuid)
	');

}