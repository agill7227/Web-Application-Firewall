package rules

import (
	"log"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func Define_sql(c *fiber.Ctx) bool {

	SqlInjPatterns := []string{
		`'`,
		`''`,
		"`",
		"``",
		`"`,
		`""`,
		`/`,
		`//`,
		`\\`,
		`\\\\`,
		`;`,
		`' or "`,
		`-- or #`,
		`' OR '1`,
		`' OR 1 -- -`,
		`" OR "" = "`,
		`" OR 1 = 1 -- -`,
		`' OR '' = '`,
		`'='`,
		`'LIKE'`,
		`'=0--+\\\\`,
		` OR 1=1`,
		`' OR 'x'='x`,
		`' AND id IS NULL; --`,
		`'{13}UNION SELECT '2`,
		`%00`,
		`/\\*\\.\\.\\.\\*/`,
		`\\+`,
		// `\\|\\|`,
		`%`,
		`@variable`,
		`@@variable`,
		`# Numeric`,
		`AND 1`,
		`AND 0`,
		`AND true`,
		`AND false`,
		`1-false`,
		`1-true`,
		`1\\*56`,
		`-2`,
		`1' ORDER BY 1--+`,
		`1' ORDER BY 2--+`,
		`1' ORDER BY 3--+`,
		`1' ORDER BY 1,2--+`,
		`1' ORDER BY 1,2,3--+`,
		`1' GROUP BY 1,2,--+`,
		`1' GROUP BY 1,2,3--+`,
		`' GROUP BY columnnames having 1=1 --`,
		`-1' UNION SELECT 1,2,3--+`,
		`' UNION SELECT sum(columnname ) from tablename --`,
		`-1 UNION SELECT 1 INTO @,@`,
		`-1 UNION SELECT 1 INTO @,@,@`,
		`1 AND (SELECT * FROM Users) = 1`,
		`' AND MID\(VERSION\(\),1,1\) = '5';`,
		`' and 1 in \(select min\(name\) from sysobjects where xtype = 'U' and name > '\.'\) --`,
		`,\(select \* from \(select\(sleep\(10\)\)\)a\)`,
		`%2c\(select%20\*%20from%20\(select\(sleep\(10\)\)\)a\)`,
		`';WAITFOR DELAY '0:0:30'--`,
		`#`,
		`/\*`,
		`-- -`,
		`;%00`,
	}

	var inputs []string
	inputs = append(inputs, string(c.Body()))
	inputs = append(inputs, string(c.Path()))

	queryParams := c.OriginalURL()
	inputs = append(inputs, queryParams)

	c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		inputs = append(inputs, string(key))
		inputs = append(inputs, string(value))
	})

	for _, input := range inputs {
		for _, pattern := range SqlInjPatterns {
			match, err := regexp.MatchString(pattern, input)
			if err != nil {
				log.Printf("Error in matching regex '%s': %v", pattern, err)
				continue
			}
			if match {
				log.Printf("Pattern '%s' matched in request body", pattern)
				return true
			}
		}
	}
	return false
}
