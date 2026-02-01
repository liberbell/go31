module restapi

go 1.25.6

require (
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	go.etcd.io/bbolt v1.4.3 // indirect
	golang.org/x/sys v0.40.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)

replace (
	handlers => ./handlers
)