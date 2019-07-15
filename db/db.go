package db

import "fmt"


type PgConfiguration struct {
	host string
	port int
	user string
	password string
	database string
	sslMode string
}

func (c PgConfiguration) PgConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.user, c.password, c.host, c.port, c.database, c.sslMode)
}

func NewPgConfiguration(host string, port int, user string, password string, database string, sslMode string) PgConfiguration {
	return PgConfiguration{
		host, port, user, password, database, sslMode,
	}
}

func (c PgConfiguration) Host() string {
	return c.host
}

func (c PgConfiguration) Port() int {
	return c.port
}

func (c PgConfiguration) User() string {
	return c.user
}

func (c PgConfiguration) Password() string {
	return c.password
}

func (c PgConfiguration) Database() string {
	return c.database
}

func (c PgConfiguration) SslMode() string {
	return c.sslMode
}

