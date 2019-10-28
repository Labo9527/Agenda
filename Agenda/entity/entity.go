package entity

import (
	"fmt"
	"encoding/json"
	"os"
	"strconv"
	"../service"
	"bufio"
)

type User struct {	//首字母大写！！！
    Username string `json:"username"`
	Password string `json:"password"`
	Mail string `json:"mail"`
	Telephone string `json:"telephone"`
}

func Login_service(myname, mypass, mymail, mytele string){
	var myuser User
	myuser.Username=myname
	myuser.Password=mypass
	myuser.Mail=mymail
	myuser.Telephone=mytele
	i:=1

	for true{
		filename := "./entity/User"+strconv.Itoa(i)+".json"
		_, err := os.Stat(filename)
		if err == nil {
			fp, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
			data := make([]byte, 100)
			n, _ := fp.Read(data)
			var user2 User
			err = json.Unmarshal(data[:n], &user2)
			//fmt.Println(user2.Username+" "+user2.Password+" "+myuser.Username+" "+myuser.Password)
			if user2.Username == myuser.Username && user2.Password == myuser.Password{
				_, err := os.Stat("entity/curUser.txt")
				if err == nil{
					service.Record_output("You have logged in.")
					 fmt.Println("You have logged in.")
					 return
				}
				fp2, _ := os.OpenFile("entity/curUser.txt", os.O_RDWR|os.O_CREATE, 0755)
				fp2.WriteString(user2.Username+"\n");
				fp2.WriteString(user2.Password+"\n");
				fp2.WriteString(user2.Mail+"\n");
				fp2.WriteString(user2.Telephone+"\n");
				service.Record_output("Login successfully!")
				fmt.Println("Login successfully!")
				break
			} 
			i=i+1
		}
		if os.IsNotExist(err) {
			service.Record_output("Login failed!")
			fmt.Println("Login failed!")
			break
		}

	}
}

func Exit_service(){
	_, err := os.Stat("entity/curUser.txt")
	if os.IsNotExist(err){
		fmt.Println("You should log in first.")
		return
	}
	err = os.Remove("entity/curUser.txt")
	service.Record_output("exit successfully!")
	fmt.Println("exit successfully!")
}

func Register_service(myname, mypass ,mymail, mytele string){
	var myuser User
	myuser.Username=myname
	myuser.Password=mypass
	myuser.Mail=mymail
	myuser.Telephone=mytele

	i:=1
	for true{
		filename := "entity/User"+strconv.Itoa(i)+".json"

		_, err := os.Stat(filename)
		if err == nil {
			fp, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
			data := make([]byte, 100)
			n, _ := fp.Read(data)
			var user2 User
			err = json.Unmarshal(data[:n], &user2)
			if user2.Username == myuser.Username{
				service.Record_output("Username has existed!")
				fmt.Println("Username has existed!")
				break
			} 
			i=i+1
		}
		if os.IsNotExist(err) {
			fp, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
			data, _ := json.Marshal(myuser)
			_, err = fp.Write(data)
			fp.Close()
			//fmt.Println(filename)
			service.Record_output("Register successfully!")
			fmt.Println("Register successfully!")
			break
		}

	}
}

func Query_service(){
	i:=1
	for true{
		filename := "entity/User"+strconv.Itoa(i)+".json"

		_, err := os.Stat(filename)
		if err == nil {
			fp, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
			data := make([]byte, 100)
			n, _ := fp.Read(data)
			var user2 User
			err = json.Unmarshal(data[:n], &user2)
			service.Record_output("-------------------")
			service.Record_output("Username:"+user2.Username)
			service.Record_output("Mail:"+user2.Mail)
			service.Record_output("Telephone:"+user2.Telephone)
			service.Record_output("-------------------")
			fmt.Println("-------------------")
			fmt.Println("Username:"+user2.Username)
			fmt.Println("Mail:"+user2.Mail)
			fmt.Println("Telephone:"+user2.Telephone)
			fmt.Println("-------------------")
			i=i+1
		}
		if os.IsNotExist(err) {
			break
		}

	}
}

func Delete_service(){
	fpp, err := os.OpenFile("entity/curUser.txt", os.O_RDWR, 0666)
	if err != nil{
		 service.Record_output("You should login first!")
		 fmt.Println("You should log in first.")
		 return
	}
	

	buf := bufio.NewReader(fpp)
	username, err := buf.ReadString('\n')
	password, err := buf.ReadString('\n')
	username = username[:len(username)-1]
	password = password[:len(password)-1]


	i:=1
	for true{
		filename := "entity/User"+strconv.Itoa(i)+".json"

		_, err := os.Stat(filename)
		if err == nil {
			fp, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
			data := make([]byte, 100)
			n, _ := fp.Read(data)
			var user2 User
			err = json.Unmarshal(data[:n], &user2)
			if user2.Username == username && user2.Password == password {
			err = os.Remove(filename)
			err = os.Remove("entity/curUser.txt")

			if err != nil{
				service.Record_output("You should login first!")
				fmt.Println("You should login first!")
			}

			for true{
				filename := "entity/User"+strconv.Itoa(i+1)+".json"
				_, err := os.Stat(filename)
				if os.IsNotExist(err){
					break
				}
				if err == nil{
					err = os.Rename(filename, "entity/User"+strconv.Itoa(i)+".json")
				}
				i = i + 1
			}
				fmt.Println("delete success!")
				service.Record_output("delete success!")
				break
			} 
			i=i+1
		}
		if os.IsNotExist(err) {
			break
		}

	}
}