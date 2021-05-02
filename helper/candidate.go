package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"latihan-web-service-go/connect"
	"os"
)

type Politician struct {
	PoliticianID int     `json:"politician_id"`
	Name         string  `json:"name"`
	Party        string  `json:"party"`
	Location     string  `json:"location"`
	GradeCurrent float32 `json:"grade_current"`
}

type Politicians []Politician

type PoliticianWithTotalVotes struct {
	Politician
	TotalVotes int `json:"total_votes"`
}

func decodePoliticianData(file string) Politicians {
	var politicians Politicians

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &politicians)
	if err != nil {
		fmt.Println(err.Error())
	}

	return politicians
}

func insertPoliticianDataToDb(p Politicians) {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, val := range p {
		_, err = db.Exec("INSERT INTO politicians (name, party, location, grade_current) VALUES (?, ?, ?, ?)", val.Name, val.Party, val.Location, val.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("insert data succeed")
}

// query for end point /politicians_voting
func AllPoliticianWithVoteQuery() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		GROUP BY p.politician_id`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		p = append(p, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	for _, each := range p {
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}

	return p
}

// query for end point /politicians
func AllPoliticianDataQuery() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM politicians`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	var data []Politician

	for rows.Next() {
		var each = Politician{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	return data
}

// query for end point /politicians_il
func ILPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'IL'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_ny
func NYPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'NY'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_tx
func TXPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'TX'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_la
func LAPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'LA'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_wa
func WAPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'WA'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_fl
func FLPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'FL'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

// query for end point /politicians_hi
func HIPoliticians() interface{} {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var data []PoliticianWithTotalVotes

	loc := `'HI'`

	q := fmt.Sprintf(`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total_votes
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = %s
		GROUP BY p.politician_id`, loc)

	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return data
}

func popularPoliticianNYQuery() {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT politician_id, name, party, location, grade_current, MAX(total) total
		FROM (SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
					 COUNT(v.politician_id) as total
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		WHERE p.location = 'NY'
		GROUP BY p.politician_id) cp`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p = append(p, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range p {
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}
}

func threePopularPoliticianQuery() {
	db, err := connect.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT p.politician_id, p.name, p.party, p.location, p.grade_current,
		COUNT(v.politician_id) as total
		FROM politicians AS p
		JOIN votings AS v ON p.politician_id = v.politician_id
		GROUP BY p.politician_id
		ORDER BY total DESC
		LIMIT 3`,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var p []PoliticianWithTotalVotes

	for rows.Next() {
		var each = PoliticianWithTotalVotes{}
		var err = rows.Scan(&each.PoliticianID, &each.Name, &each.Party, &each.Location, &each.GradeCurrent, &each.TotalVotes)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p = append(p, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range p {
		fmt.Println("Politician ID:", each.PoliticianID)
		fmt.Println("Name:", each.Name)
		fmt.Println("Party:", each.Party)
		fmt.Println("Location:", each.Location)
		fmt.Println("Grade Current:", each.GradeCurrent)
		fmt.Println("Total Votes:", each.TotalVotes)
	}
}

func Candidate() {
	// p := decodePoliticianData("json/politicians.json")

	// insertPoliticianDataToDb(p)
	AllPoliticianWithVoteQuery()
	// popularPoliticianNYQuery()
	// threePopularPoliticianQuery()
}
