package main

import (
	"fmt"
)

func println(a ...interface{}) {
	fmt.Println(a...)
}

func print(a ...interface{}) {
	fmt.Print(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	/*JennylovedhercatsSheownedsixofthemJennyspentallherdayswiththecatsHerfriendsworriedaboutherYoushouldgooutmoretheysaidIcantstandleavingmycatsalonerepliedJennyOnedayJennymetamanatworkHewasfunnyandhandsomeJennystartedgoingoutwithhimalotShethoughthewasperfectJennyinvitedhimtoherhouseShewasveryexcitedHesgoingtolovemycatsshethoughtThemanwalkedintoherhouse.Hestartedsneezing.HecouldntbreatheHewasallergictocatsIhatecatshescreamed
	จงเขียนcodeการนับคำที่รับจาก input จากบทความนี้ว่ามีกี่คำ
	*/

	string_WhatTheCat := "JennylovedhercatsSheownedsixofthemJennyspentallherdayswiththecatsHerfriendsworriedaboutherYoushouldgooutmoretheysaidIcantstandleavingmycatsalonerepliedJennyOnedayJennymetamanatworkHewasfunnyandhandsomeJennystartedgoingoutwithhimalotShethoughthewasperfectJennyinvitedhimtoherhouseShewasveryexcitedHesgoingtolovemycatsshethoughtThemanwalkedintoherhouse"

	string_number := len(string_WhatTheCat)
	if string_number > 0 {
		printf("พบ ตัวหนังสือทั้งหมด: %d\n", string_number)
	} else {
		print("ไม่พบ")
	}

}
