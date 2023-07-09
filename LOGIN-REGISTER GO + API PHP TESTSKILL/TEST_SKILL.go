package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func print(a ...interface{}) {
	fmt.Print(a...)
}

func println(a ...interface{}) {
	fmt.Println(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func GETURL(URL string) string {
	url := fmt.Sprintf("http://localhost/api/%s.php", URL)
	return url
}

func Login(username string, password string) {
	url := GETURL("login")
	data := []byte(fmt.Sprintf("username=%s&password=%s", username, password))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Resp", string(body))
}

func Register(username string, password string) {
	url := GETURL("register")
	data := []byte(fmt.Sprintf("username=%s&password=%s", username, password))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Resp", string(body))
}

func main() {
	mode := ""
	println("Login : A | Register : B ") //เวลาเข้างาน
	Scan(&mode)

	userinput := ""
	passwordinput := ""

	if mode == "A" {
		//LOGIN
		println("[LOG] LOGIN")

		print("กรุณาใส่ Username: ")
		Scan(&userinput)

		if len(userinput) > 0 {
			println("[LOG] LOGIN USERNAME")
		} else {
			println("[LOG] กรุณาใส่ username มากกว่า 0 ตัวขึ้นไป")
		}

		print("กรุณาใส่ Password: ")
		Scan(&passwordinput)

		if len(passwordinput) > 0 {
			println("[LOG] LOGIN PASSWORD")
		} else {
			println("[LOG] กรุณาใส่ password มากกว่า 0 ตัวขึ้นไป")
		}

		if userinput != "" && passwordinput != "" {
			Login(userinput, passwordinput)
		}

	} else if mode == "B" {
		//REGISTER
		println("[LOG] REGISTER")

		print("กรุณาใส่ Username: ")
		Scan(&userinput)

		if len(userinput) > 0 {
			println("[LOG] REGISTER USERNAME OK")
		} else {
			println("[LOG] กรุณาใส่ username มากกว่า 0 ตัวขึ้นไป")
		}

		print("กรุณาใส่ Password: ")
		Scan(&passwordinput)

		if len(passwordinput) > 0 {
			println("[LOG] REGISTER PASSWORD OK")
		} else {
			println("[LOG] กรุณาใส่ password มากกว่า 0 ตัวขึ้นไป")
		}

		if userinput != "" && passwordinput != "" {
			Register(userinput, passwordinput)
		}

	} else {
		println("[LOG] กรุณาเลือกโหมดใหม่")
	}

}
