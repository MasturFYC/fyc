package fyc

func GetIndonesianMonthName(id int, isShort bool) string {
	months := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember"}
	if isShort {
		return months[id][0:3]
	}
	return months[id]
}
