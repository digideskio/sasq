package sasq

import (
	"strconv"
	"strings"
)

// ASRecord is the return struct of the QueryWhois with
// all the fields stored in the Shadowserver IP-BGP whois
// database
type ASRecord struct {
	Query  string `json:"query"`
	ASN    int    `json:"asn"`
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
	CC     string `json:"cc"`
	Domain string `json:"domain"`
	ISP    string `json:"isp"`
}

// parse_line is a helper function that takes one line of
// shadowserver return data and parses it into a valid ASRecord
func parse_line(line string) (*ASRecord, error) {
	elems := strings.Split(line, " | ")

	asn, err := strconv.Atoi(elems[1])
	if err != nil {
		return nil, err
	}
	return &ASRecord{
		strings.TrimSpace(elems[0]),
		asn,
		strings.TrimSpace(elems[2]),
		strings.TrimSpace(elems[3]),
		strings.TrimSpace(elems[4]),
		strings.TrimSpace(elems[5]),
		strings.TrimSpace(elems[6]),
	}, nil
}
