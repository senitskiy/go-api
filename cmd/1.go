package main

import (
  "encoding/csv"
  "fmt"
  "os"
    // "strings"
)



func main() {

  lines, err := ReadCsv("./ueba.csv")

  if err != nil {
      panic(err)
  }

// stringSearch1 := "5229, 5647"
	stringSearch := []string{"5229", "5647", "5358"}

  // Loop through lines & turn into object
  for _, line := range lines {
    data := CsvLine{
      Column1: line[0],
      Column2: line[1],
      Column3: line[2],
      Column4: line[3],  
      Column5: line[4],  
      Column6: line[5],
      Column7: line[6],
      Column8: line[7],
      Column9: line[8],
      Column10: line[9],
      Column11: line[10],
      Column12: line[11],          
      Column13: line[12],
      Column14: line[13],
      Column15: line[14],
      Column16: line[15],
      Column17: line[16],
      Column18: line[17],
      Column19: line[18],
      Column20: line[19],
      Column21: line[20],
      Column22: line[21],
      Column23: line[22],
      Column24: line[23],
      Column25: line[24],
      Column26: line[25],
      Column27: line[26],
      Column28: line[27],
      Column29: line[28],
      Column30: line[29],
      Column31: line[30],
      Column32: line[31],
      Column33: line[32],
      Column34: line[33],
      Column35: line[34],
      Column36: line[35],
      Column37: line[36],
      Column38: line[37],
      Column39: line[38],
      Column40: line[39],
      Column41: line[40],
      Column42: line[41],
      Column43: line[42],
      Column44: line[43],
      Column45: line[44],
      Column46: line[45],
      Column47: line[46],
      Column48: line[47],			
    }
    // contain := strings.Contains(stringSearch, data.Column2)
    // if contain {
    // // if "872" == data.Column2 {
    //     // fmt.Println(data.Column2 + " " + data.Column5 + " ")
    // }       
    for i := 0; i < len(stringSearch); i++ {      
      if stringSearch[i] == data.Column2 {
        // if "872" == data.Column2 {
        resString :="{ \"id\" : " + "\"" + data.Column1 + "\", " + 
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
        fmt.Println(resString)    
        // return resString
            // fmt.Println(data.Column2 + " " + data.Column5 + " ")
            // fmt.Println("{ \"id\" : " + "\"" + data.Column1 + "\"} ")
      } 
    }         
  }
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