package main

import (
	"fmt"
	"strings"
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
	println("ใส่ string ที่คุณต้องการ :")
	string_What := ""
	Scan(&string_What)

	string_count := strings.Count(string_WhatTheCat, string_What)

	startIndex := 0

	for {
		index := strings.Index(string_WhatTheCat[startIndex:], string_What)
		if index == -1 {
			break
		}

		printf("พบคำว่า 'cat' ที่ตำแหน่งที่: %d\n", startIndex+index)

		startIndex += index + len(string_What)
	}

	if string_count > 0 {
		printf("พบ %s จำนวนทั้งหมด: %d\n", string_What, string_count)
	} else {
		printf("ไม่พบ %s", string_What)
	}

}
