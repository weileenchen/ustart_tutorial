package SQLStore

import (
	"context"
	"fmt"
	"strconv"
)

func (dbConn *SQLStore) Insert(ctx context.Context, carID, userID, dateStart string, rate int) error {
	queryString := fmt.Sprintf(
		"INSERT INTO %s (carID, userID, dateStart, rate) VALUES ('%s', '%s', '%s', '"+strconv.Itoa(rate)+"');",
		dbConn.recordTableName, carID, userID, dateStart, rate)

	_, err := dbConn.dc.QueryContext(ctx, queryString)

	return err
}
