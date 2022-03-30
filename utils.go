package fyc

import (
   "fmt"
   "time"
)

func GetIndonesianMonthName(id int, isShort bool) string {
   months := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember"}
   if isShort {
      return months[id][0:3]
   }
   return months[id]
}

func NestQuerySingle(query string) string {
   return fmt.Sprintf(`(SELECT row_to_json(x) FROM (%s) x)`, query)
}

func NestQuery(query string) string {
   return fmt.Sprintf(`COALESCE((
   SELECT array_to_json(array_agg(row_to_json(x)))
   FROM (%s) x), '[]')`, query)
}

func CreateIndonesianDate(date string, isShort bool) string {

   t, err := time.Parse("2006-01-02", date[0:10])

   if err != nil {
      return date[0:10]
   }
   year, month, day := t.Date()

   return fmt.Sprintf("%02d %s %d", day, GetIndonesianMonthName(int(month-1), isShort), year)
}
