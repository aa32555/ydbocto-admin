package ddl

const (
	Oid = iota
	Rolname = iota
	Rolsuper = iota
	Rolinherit = iota
	Rolcreaterole = iota
	Rolcreatedb = iota
	Rolcanlogin = iota
	Rolreplication = iota
	Rolbypassrls = iota
	Rolconnlimit = iota
	Rolpassword = iota
	Rolvaliduntil = iota
)

type User struct {
	oid int
	rolname string
	rolsuper int
	rolinherit int
	rolcreaterole int
	rolcreatedb int
	rolcanlogin int
	rolreplication int
	rolbypassrls int
	rolconnlimit int
	rolpassword string
	rolvaliduntil string
}
