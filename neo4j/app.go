package neo4j

import "github.com/neo4j/neo4j-go-driver/neo4j"

func Main() {

	neo4j.NewDriver("", neo4j.BasicAuth("", "", ""))

}
