package elasticstore

import (
	"sync"
)

//Elastic search does not have built in concurrency control
//so we need to use mutexes to make sure we dont have version conflicts when access or modifying data

var newUserLock sync.Mutex
