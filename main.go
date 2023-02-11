//main.go
package main

import (
 "fmt"
 "log"
 "net/http"
 "io"
 "github.com/gorilla/mux"
 "encoding/csv"
 "os"
)

type CsvLine struct {
  Column0 string
  Column1 string
  Column2 string
  Column3 string
  Column4 string
	Column5 string
  Column6 string
  Column7 string
  Column8 string
  Column9 string
  Column10 string
  Column11 string
  Column12 string
  Column13 string
  Column14 string
  Column15 string
  Column16 string
  Column17 string
  Column18 string
  Column19 string
  Column20 string
  Column21 string
  Column22 string
  Column23 string
  Column24 string
  Column25 string
  Column26 string
  Column27 string
  Column28 string
  Column29 string
  Column30 string
  Column31 string
  Column32 string
  Column33 string
  Column34 string
  Column35 string
  Column36 string
  Column37 string
  Column38 string
  Column39 string
  Column40 string
  Column41 string
  Column42 string
  Column43 string
  Column44 string
  Column45 string
  Column46 string
  Column47 string
  Column48 string
}

func main() {
 router := mux.NewRouter()

 router.HandleFunc("/get-items", GetItemsById).Methods("GET")


 fmt.Println("Starting server on the port 3000...")
 log.Fatal(http.ListenAndServe(":3000", router))
}

func GetItemsById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
		fmt.Println(r)
        return
    }
	// fmt.Println(r)
	// r.Context().Value(varsKey)
    // w.Write([]byte("foo"))
	// log.Println(r.Context())
	// fds := r.URL.Query()["id"]
	
	if values := r.URL.Query(); values != nil {
		// log.Println(values)
		for k, v := range values {
			// fmt.Println(k, " => ", v)
			// io.WriteString(w, "{ \"id\" : " + "\"" + r.FormValue("id") + "\"} ")
			// io.WriteString(w, "{ \"id\" : " + "\"" + [0] + "\"} ")
			// fmt.Println(values)
			 //v[0]
			fmt.Println(k)

      if v != nil {        
        for i := range v {
          // v[i] = "New " + v[i]
          fmt.Println(v[i])
          // searchStr(v[i])
          res := searchStr(v[i])
          // io.WriteString(w, "{ \"id\" : " + "\"" + res + "\"} ")
          io.WriteString(w, res)
          fmt.Println(res)
        }
      }
      // fmt.Println(values)
		}

		
		// io.WriteString(r, "id: "+ w("id"))
	}
	// return nil
	
	// log.Println("id: " + r.Form.Get("id"))
	// log.Println(r.URL.Query())
	// log.Println(r.FormValue("id"))
	// log.Println(r.ParseForm())

	// id := r.URL.Query().Get("id")
	// fmt.Println("id =>", id)
	
	// values := r.URL.Query()
	// for k, v := range values {
	// 	fmt.Println(k, " => ", v)
	// }


	// log.Println(r.RequestURI)
}

func searchStr(stringSearch string) string {

	lines, err := ReadCsv("ueba.csv")
	if err != nil {
		panic(err)
	}

  resString := ""
	// stringSearch1 := "5229, 5647"
	// stringSearch := []string{"5229", "5647"}

	// Loop through lines & turn into object
	for _, line := range lines {
    data := CsvLine{
      Column1: line[1],
      Column2: line[2],
      Column3: line[3],
      Column4: line[4],  
      Column5: line[5],  
      Column6: line[6],
      Column7: line[7],
      Column8: line[8],
      Column9: line[9],
      Column10: line[10],
      Column11: line[11],
      Column12: line[12],          
      Column13: line[13],
      Column14: line[14],
      Column15: line[15],
      Column16: line[16],
      Column17: line[17],
      Column18: line[18],
      Column19: line[19],
      Column20: line[20],
      Column21: line[21],
      Column22: line[22],
      Column23: line[23],
      Column24: line[24],
      Column25: line[25],
      Column26: line[26],
      Column27: line[27],
      Column28: line[28],
      Column29: line[29],
      Column30: line[30],
      Column31: line[31],
      Column32: line[32],
      Column33: line[33],
      Column34: line[34],
      Column35: line[35],
      Column36: line[36],
      Column37: line[37],
      Column38: line[38],
      Column39: line[39],
      Column40: line[40],
      Column41: line[41],
      Column42: line[42],
      Column43: line[43],
      Column44: line[44],
      Column45: line[45],
      Column46: line[46],
      Column47: line[47],
      Column48: line[48],			
    }
		// contain := strings.Contains(stringSearch, data.Column2)
		// if contain {
		// // if "872" == data.Column2 {
		//     // fmt.Println(data.Column2 + " " + data.Column5 + " ")
		// }       
		// for i := 0; i < len(stringSearch); i++ {   
      
      // fmt.Println("stringSearch " + stringSearch)
      //   fmt.Println("Column1 " + data.Column2)
    if stringSearch == data.Column1 {
      
      resString ="{ \"id\" : " + "\"" + data.Column1 + "\", " + 
      "\"uid\" : " + "\"" + data.Column2 + "\", " +
      "\"domain\" : " + "\"" + data.Column3 + "\", " +
      "\"cn\" : " + "\"" + data.Column4 + "\", " +
      "\"department\" : " + "\"" + data.Column5 + "\", " +
      "\"title\" : " + "\"" + data.Column6 + "\", " +
      "\"who\" : " + "\"" + data.Column7 + "\", " +
      "\"logon_count\" : " + "\"" + data.Column8 + "\", " +
      "\"num_logons7\" : " + "\"" + data.Column9 + "\", " +
      "\"num_share7\" : " + "\"" + data.Column10 + "\", " +
      "\"num_file7\" : " + "\"" + data.Column11 + "\", " +
      "\"num_ad7\" : " + "\"" + data.Column12 + "\", " +
      "\"num_n7\" : " + "\"" + data.Column13 + "\", " +
      "\"num_logons14\" : " + "\"" + data.Column14 + "\", " +
      "\"num_share14\" : " + "\"" + data.Column15 + "\", " +
      "\"num_file14\" : " + "\"" + data.Column16 + "\", " +
      "\"num_ad14\" : " + "\"" + data.Column17 + "\", " +
      "\"num_n14\" : " + "\"" + data.Column18 + "\", " +
      "\"num_logons30\" : " + "\"" + data.Column19 + "\", " +
      "\"num_share30\" : " + "\"" + data.Column20 + "\", " +
      "\"num_file30\" : " + "\"" + data.Column21 + "\", " +
      "\"num_ad30\" : " + "\"" + data.Column22 + "\", " +
      "\"num_n30\" : " + "\"" + data.Column23 + "\", " +
      "\"num_logons150\" : " + "\"" + data.Column24 + "\", " +
      "\"num_share150\" : " + "\"" + data.Column25 + "\", " +
      "\"num_file150\" : " + "\"" + data.Column26 + "\", " +
      "\"num_ad150\" : " + "\"" + data.Column27 + "\", " +
      "\"num_n150\" : " + "\"" + data.Column28 + "\", " +
      "\"num_logons365\" : " + "\"" + data.Column29 + "\", " +
      "\"num_share365\" : " + "\"" + data.Column30 + "\", " +
      "\"num_file365\" : " + "\"" + data.Column31 + "\", " +
      "\"num_ad365\" : " + "\"" + data.Column32 + "\", " +
      "\"num_n365\" : " + "\"" + data.Column33 + "\", " +
      "\"has_user_principal_name\" : " + "\"" + data.Column34 + "\", " +
      "\"has_mail\" : " + "\"" + data.Column35 + "\", " +
      "\"has_phone\" : " + "\"" + data.Column36 + "\", " +
      "\"flag_disabled\" : " + "\"" + data.Column37 + "\", " +
      "\"flag_lockout\" : " + "\"" + data.Column38 + "\", " +
      "\"flag_password_not_required\" : " + "\"" + data.Column39 + "\", " +
      "\"flag_password_cant_change\" : " + "\"" + data.Column40 + "\", " +
      "\"flag_dont_expire_password\" : " + "\"" + data.Column41 + "\", " +
      "\"owned_files\" : " + "\"" + data.Column41 + "\", " +
      "\"num_mailboxes\" : " + "\"" + data.Column42 + "\", " +
      "\"num_member_of_groups\" : " + "\"" + data.Column43 + "\", " +
      "\"num_member_of_indirect_groups\" : " + "\"" + data.Column44 + "\", " +
      "\"member_of_indirect_groups_ids\" : " + "\"" + data.Column45 + "\", " +
      "\"member_of_groups_ids\" : " + "\"" + data.Column46 + "\", " +
      "\"is_admin\" : " + "\"" + data.Column47 + "\", " +                                
      "\"is_service\" : " + "\"" + data.Column48 + "\"} "
      // fmt.Println(resString)
      // if "872" == data.Column2 {
        // fmt.Println(data.Column2 + " " + data.Column5 + " ")
    } 		
	}
  return resString
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

  // Open CSV file
  f, err := os.Open(filename)
  if err != nil {
    return [][]string{}, err
  }
  defer f.Close()

  // Read File into a Variable
  lines, err := csv.NewReader(f).ReadAll()
  if err != nil {
    return [][]string{}, err
  }

  return lines, nil
}