package gocql

import (
	"github.com/flaxius/portia/pkg/cassandra"
	"github.com/gocql/gocql"
)
//COPIED some funcs from jaeger pkgs

// CQLSession is a wrapper around gocql.Session.
type CQLSession struct {
	session *gocql.Session
}
// WrapCQLSession creates a Session out of *gocql.Session.
func WrapCQLSession(session *gocql.Session) CQLSession {
	return CQLSession{session: session}
}

// Query delegates to gocql.Session#Query and wraps the result as Query.
func (s CQLSession) Query(stmt string, values ...interface{}) cassandra.Query {
	return WrapCQLQuery(s.session.Query(stmt, values...))
}

// Close delegates to gocql.Session#Close.
func (s CQLSession) Close() {
	s.session.Close()
}

// ---

// CQLQuery is a wrapper around gocql.Query.
type CQLQuery struct {
	query *gocql.Query
}

// WrapCQLQuery creates a Query out of *gocql.Query.
func WrapCQLQuery(query *gocql.Query) CQLQuery {
	return CQLQuery{query: query}
}

// Exec delegates to gocql.Query#Exec.
func (q CQLQuery) Exec() error {
	return q.query.Exec()
}

// ScanCAS delegates to gocql.Query#ScanCAS.
func (q CQLQuery) ScanCAS(dest ...interface{}) (bool, error) {
	return q.query.ScanCAS(dest...)
}

// Iter delegates to gocql.Query#Iter and wraps the result as Iterator.
func (q CQLQuery) Iter() cassandra.Iterator {
	return WrapCQLIterator(q.query.Iter())
}

// Bind delegates to gocql.Query#Bind and wraps the result as Query.
func (q CQLQuery) Bind(v ...interface{}) cassandra.Query {
	return WrapCQLQuery(q.query.Bind(v...))
}

// Consistency delegates to gocql.Query#Consistency and wraps the result as Query.
func (q CQLQuery) Consistency(level cassandra.Consistency) cassandra.Query {
	return WrapCQLQuery(q.query.Consistency(gocql.Consistency(level)))
}

// String returns string representation of this query.
func (q CQLQuery) String() string {
	return q.query.String()
}

// PageSize delegates to gocql.Query#PageSize and wraps the result as Query.
func (q CQLQuery) PageSize(n int) cassandra.Query {
	return WrapCQLQuery(q.query.PageSize(n))
}

// ---

// CQLIterator is a wrapper around gocql.Iter.
type CQLIterator struct {
	iter *gocql.Iter
}

// WrapCQLIterator creates an Iterator out of *gocql.Iter.
func WrapCQLIterator(iter *gocql.Iter) CQLIterator {
	return CQLIterator{iter: iter}
}

// Scan delegates to gocql.Iter#Scan.
func (i CQLIterator) Scan(dest ...interface{}) bool {
	return i.iter.Scan(dest...)
}

// Close delegates to gocql.Iter#Close.
func (i CQLIterator) Close() error {
	return i.iter.Close()
}
