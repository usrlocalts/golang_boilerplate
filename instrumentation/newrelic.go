package instrumentation

import (
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func StartDataSegmentNowForPostgres(op string, tableName string, txn newrelic.Transaction) newrelic.DatastoreSegment {
	if txn != nil {
		s := newrelic.DatastoreSegment{
			Product:    newrelic.DatastorePostgres,
			Collection: tableName,
			Operation:  op,
		}

		s.StartTime = newrelic.StartSegmentNow(txn)
		return s
	}
	return newrelic.DatastoreSegment{}
}

func GetNewRelicTransaction(w http.ResponseWriter, enabled bool) newrelic.Transaction {
	if enabled {
		newRelicTransaction, _ := w.(newrelic.Transaction)
		return newRelicTransaction
	}
	return nil
}
