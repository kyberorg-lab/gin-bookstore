package models

import (
	"fmt"
	"strings"
)

type PostgreParams struct {
	DBHost   string
	DBPort   int
	DBName   string
	User     string
	Password string
	SSL      bool
}

func (pp *PostgreParams) toDSN() string {
	var b strings.Builder
	if pp.DBHost == "" {
		pp.DBHost = "localhost"
	}
	fmt.Fprintf(&b, "host=%s", pp.DBHost)

	if pp.DBPort <= 0 || pp.DBPort > 65535 {
		pp.DBPort = 5432
	}
	fmt.Fprintf(&b, "port=%d", pp.DBPort)

	if pp.DBName == "" {
		pp.DBName = "postgres"
	}
	fmt.Fprintf(&b, "dbname=%s", pp.DBName)

	if pp.User != "" {
		fmt.Fprintf(&b, "user=%s", pp.User)
	}
	if pp.Password != "" {
		fmt.Fprintf(&b, "password=%s", pp.Password)
	}

	if pp.SSL {
		fmt.Fprintf(&b, "sslmode=enable")
	} else {
		fmt.Fprintf(&b, "sslmode=disable")
	}

	fmt.Printf("DSN %s\n", b.String())

	return b.String()
}
