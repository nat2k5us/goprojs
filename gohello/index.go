package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/araddon/dateparse"
	//"github.com/geoffreybauduin/termtables"
	"encoding/json"
	"time"
)

func main() {
	// loc, err := time.LoadLocation("America/New_York")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// DataTypes()
	// ConditionalStatements(17)
	// CaseStateMents(18)
	// //SimpleHttpServer()
	// StartRestApiServices()
	// LoopTest()
	// table := termtables.CreateTable()
	// for _, dateExample := range examples1 {
	// 	//	ts1 := "2016-06-20T12:41:45.14Z" <-- Final Format
	// 	t, err := time.Parse("11/18/2019 09:16:37 AM", dateExample)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	//	fmt.Println(t)
	// 	//	ret,_ := time.Parse("8/8/1965 12:00:00 AM", dateExample)
	// 	table.AddRow(dateExample, fmt.Sprintf("%v", t))
	// }
	// fmt.Println(table.Render())

	JsonMarshallTimeTest()

	
}
func JsonMarshallTimeTest(){
var u MyUser
	err := json.Unmarshal([]byte(`{"id":1,"name":"Ken","lastSeen":1234567890,"lastSeenRFC3393":"11/18/2019 09:16:37 AM"}`), &u)
	fmt.Println(err, u)
}
type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
	LastSeenRFC3393 string `json:"lastSeenRFC3393"`
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := struct {
		LastSeen int64 `json:"lastSeen"`
		LastSeenRFC3393 string `json:"lastSeenRFC3393"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	u.LastSeenRFC3393 = u.LastSeen.Format(time.RFC3339)
	trfc, err123 := time.Parse("23/11/2019 09:16:37 AM", u.LastSeenRFC3393)
	if err123 != nil {
		return err123
	}
	u.LastSeenRFC3393 = trfc.Format(time.RFC3339)
	return nil
}

var examples1 = []string{
	"11/18/2019 09:16:37 AM",
	"11/18/2019 09:01:46 AM",
	"11/11/2019 15:33:09 AM",
}

var examples = []string{
	"May 8, 2009 5:57:51 PM",
	"oct 7, 1970",
	"oct 7, '70",
	"oct. 7, 1970",
	"oct. 7, 70",
	"Mon Jan  2 15:04:05 2006",
	"Mon Jan  2 15:04:05 MST 2006",
	"Mon Jan 02 15:04:05 -0700 2006",
	"Monday, 02-Jan-06 15:04:05 MST",
	"Mon, 02 Jan 2006 15:04:05 MST",
	"Tue, 11 Jul 2017 16:28:13 +0200 (CEST)",
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"Thu, 4 Jan 2018 17:53:36 +0000",
	"Mon Aug 10 15:44:11 UTC+0100 2015",
	"Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)",
	"September 17, 2012 10:09am",
	"September 17, 2012 at 10:09am PST-08",
	"September 17, 2012, 10:10:09",
	"October 7, 1970",
	"October 7th, 1970",
	"12 Feb 2006, 19:17",
	"12 Feb 2006 19:17",
	"7 oct 70",
	"7 oct 1970",
	"03 February 2013",
	"1 July 2013",
	"2013-Feb-03",
	//   mm/dd/yy
	"3/31/2014",
	"03/31/2014",
	"08/21/71",
	"8/1/71",
	"4/8/2014 22:05",
	"04/08/2014 22:05",
	"4/8/14 22:05",
	"04/2/2014 03:00:51",
	"8/8/1965 12:00:00 AM",
	"11/13/2016 2:00:00 AM",
	"11/13/2016 02:10:10 AM",
	"11/13/2016 22:10:10",
	"8/8/1965 01:00:01 PM",
	"8/8/1965 01:00 PM",
	"8/8/1965 1:00 PM",
	"8/8/1965 12:00 AM",
	"4/02/2014 03:00:51",
	"03/19/2012 10:11:59",
	"03/19/2012 10:11:59.3186369",
	// yyyy/mm/dd
	"2014/3/31",
	"2014/03/31",
	"2014/4/8 22:05",
	"2014/04/08 22:05",
	"2014/04/2 03:00:51",
	"2014/4/02 03:00:51",
	"2012/03/19 10:11:59",
	"2012/03/19 10:11:59.3186369",
	// Chinese
	"2014年04月08日",
	//   yyyy-mm-ddThh
	"2006-01-02T15:04:05+0000",
	"2009-08-12T22:15:09-07:00",
	"2009-08-12T22:15:09",
	"2009-08-12T22:15:09Z",
	//   yyyy-mm-dd hh:mm:ss
	"2014-04-26 17:24:37.3186369",
	"2012-08-03 18:31:59.257000000",
	"2014-04-26 17:24:37.123",
	"2013-04-01 22:43",
	"2013-04-01 22:43:22",
	"2014-12-16 06:20:00 UTC",
	"2014-12-16 06:20:00 GMT",
	"2014-04-26 05:24:37 PM",
	"2014-04-26 13:13:43 +0800",
	"2014-04-26 13:13:43 +0800 +08",
	"2014-04-26 13:13:44 +09:00",
	"2012-08-03 18:31:59.257000000 +0000 UTC",
	"2015-09-30 18:48:56.35272715 +0000 UTC",
	"2015-02-18 00:12:00 +0000 GMT",
	"2015-02-18 00:12:00 +0000 UTC",
	"2015-02-08 03:02:00 +0300 MSK m=+0.000000001",
	"2015-02-08 03:02:00.001 +0300 MSK m=+0.000000001",
	"2017-07-19 03:21:51+00:00",
	"2014-04-26",
	"2014-04",
	"2014",
	"2014-05-11 08:20:13,787",
	// mm.dd.yy
	"3.31.2014",
	"03.31.2014",
	"08.21.71",
	"2014.03",
	"2014.03.30",
	//  yyyymmdd and similar
	"20140601",
	"20140722105203",
	// unix seconds, ms, micro, nano
	"1332151919",
	"1384216367189",
	"1384216367111222",
	"1384216367111222333",
}

type Time struct {
	*time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	const format = "\"2006-01-02T15:04:05+00:00\""
	t_, err := time.Parse(format, string(b))
	if err != nil {
		return err
	}
	*t = Time{&t_}
	return nil
}

type Example struct {
	TimeField *Time `json:"time_field"`
}

func LoopTest() {
	count := 10
	for index := 0; index < count; index++ {
		println(index)
	}
}

func convertToUtcTime(timeStamp string) (ret string) {

	RFC3339local := "2006-01-02T15:04:05Z"

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}

	ts1 := "2016-06-20T12:41:45.14Z"
	t, _ := time.ParseInLocation(RFC3339local, ts1, loc)
	fmt.Println(t)

	ts2 := "2016-01-20T12:41:45.14Z"
	t, _ = time.ParseInLocation(RFC3339local, ts2, loc)
	fmt.Println(t)
	// prints: 2016-01-20 12:41:45.14 -0500 EST
	return t.String()

	// const longForm = "2006-01-02 15:04:05 MST"
	// t, err := time.Parse(longForm, "2016-01-17 20:04:05 IST")
	// fmt.Println(t, err)
	// fmt.Printf("IST to UTC: %v\n\n", t.UTC())

	// s, err1 := time.Parse(longForm, "2016-01-17 23:04:05 PST")
	// fmt.Println(s, err1)
	// fmt.Printf("PST to UTC: %v\n\n", s.UTC())

}

func convertTime(t string) (ret string) {
	// time format
	const (
		RFC3339 = "2006-01-02T15:04:05Z07:00"
	)

	//now time
	//	now := time.Now()
	//	nowtime := now.Unix()

	//	fmt.Println("Nowtime:", nowtime)
	//	fmt.Println("Now:", now)

	//time in db
	//	fmt.Println("Dbtime string:", t)
	udbtime, _ := time.Parse("11/13/2016 2:00:00 AM", t)

	//	fmt.Println("RFC3339: " + RFC3339)
	fmt.Println("dbtime parsed", udbtime)
	//	fmt.Println("dbtime parsed unixtime", udbtime.Unix())

	return udbtime.String()
}

func convertToEasternTime(timeStamp string) (ret string) {
	if len(timeStamp) == 18 {
		t, _ := time.Parse("1/02/06 3:04:05 PM", timeStamp)
		if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
			t, _ := time.Parse("01/2/06 03:04:05 PM", timeStamp)
			if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
				t, _ := time.Parse("1/2/06 03:04:05 PM", timeStamp)
				if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
					t, _ := time.Parse("01/2/06 3:04:05 PM", timeStamp)
					easternTime := t.Add(time.Hour * 3)
					ret = easternTime.Format(time.RFC3339)
					//	fmt.Println(ret)
				} else {
					easternTime := t.Add(time.Hour * 3)
					ret = easternTime.Format(time.RFC3339)
					//	fmt.Println(ret)
				}
			} else {
				easternTime := t.Add(time.Hour * 3)
				ret = easternTime.Format(time.RFC3339)
				//fmt.Println(ret)
			}
		} else {
			easternTime := t.Add(time.Hour * 3)
			ret = easternTime.Format(time.RFC3339)
			//	fmt.Println(ret)
		}
	} else if len(timeStamp) == 19 {
		t, _ := time.Parse("1/02/06 03:04:05 PM", timeStamp)
		if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
			t, _ := time.Parse("01/02/06 3:04:05 PM", timeStamp)
			if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
				t, _ := time.Parse("01/2/06 03:04:05 PM", timeStamp)
				easternTime := t.Add(time.Hour * 3)
				ret = easternTime.Format(time.RFC3339)
				//fmt.Println(ret)
			} else {
				easternTime := t.Add(time.Hour * 3)
				ret = easternTime.Format(time.RFC3339)
				//	fmt.Println(ret)
			}
		} else {
			easternTime := t.Add(time.Hour * 3)
			ret = easternTime.Format(time.RFC3339)
			//	fmt.Println(ret)
		}
	} else if len(timeStamp) == 20 {
		t, _ := time.Parse("01/02/06 03:04:05 PM", timeStamp)
		easternTime := t.Add(time.Hour * 3)
		ret = easternTime.Format(time.RFC3339)
		//	fmt.Println(ret)
	}
	return ret
}

func CaseStateMents(age int) {
	switch age {
	case 10:
		println("Don't run after girls")
		break
	case 18:
		println("Harmones are running high")
		break
	default:
		break
	}
}

func ConditionalStatements(age int) {
	if age < 5 {
		println("you are a baby")
	} else if age < 12 {
		println("you are a toddler")
	} else if age < 18 {
		println("you are a teen")
	} else {
		println("you are the man")
	}
}

func DataTypes() {
	var name string = "Voidy Walker"
	const pi float64 = 3.14124
	win := true
	x := 5

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", name) // prints the DataType as string 'T'
	fmt.Printf("%s \n", name)
	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 34)
	fmt.Printf("%x \n", 15)
	fmt.Printf(" '5 mod 2' = %d \n", 5%2)
	fmt.Printf(" '4 mod  2' = %d \n", 4%2)
}

func SimpleHttpServer() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}

func TestEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "testing server - server is running !!!")
}

func StartRestApiServices() {
	fmt.Println("Starting application")
	log.Print("Starting Go Server at http://localhost:8011")
	router := mux.NewRouter()

	router.HandleFunc("/test", TestEndPoint).Methods("GET")
	//router.HandleFunc("/projects", GetProjectsDataEndPoint).Methods("GET")
	//router.HandleFunc("/project/{projectid}", GetProjectEndPoint).Methods("GET")
	//router.HandleFunc("/project/{projectid}", CreateProjectEndPoint).Methods("POST")
	//router.HandleFunc("/project/{projectid}", DeleteProjectEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8011", router))
}
