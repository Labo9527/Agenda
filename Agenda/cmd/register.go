/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	//"log"
	"../service"
	"../entity"
	"regexp"
	"fmt"
)

type User struct {	//首字母大写！！！
    Username string `json:"username"`
	Password string `json:"password"`
	Mail string `json:"mail"`
	Telephone string `json:"telephone"`
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Just register",
	Long: `You should user register like agenda register -u username -p password -m mailaddress -t telephonenumber`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("register called")
		myname, _ := cmd.Flags().GetString("username")
		mypass, _ := cmd.Flags().GetString("password")
		mymail, _ := cmd.Flags().GetString("mail")
		mytele, _ := cmd.Flags().GetString("telephone")
		//fmt.Println("username:"+myname)
		//fmt.Println("password:"+mypass)
		//fmt.Println("mail:"+mymail)
		//fmt.Println("telephone:"+mytele)

		//myuser := User{myname,mypass,mymail,mytele}

		service.Record_input("Agenda register -u "+myname+" -p "+mypass+" -m "+mymail+" -t "+mytele)

		if isOk, _ := regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", mymail); isOk {
			
		} else{
			fmt.Println("mail format error.")
			service.Record_output("mail format error.")
			return
		}

		if isOk1, _ := regexp.MatchString("^[0-9]{11}$", mytele); isOk1{

		} else{
			fmt.Println("telephone format error.")
			service.Record_output("telephone format error.")
			return
		}


		entity.Register_service(myname,mypass,mymail,mytele)


		// data, err := json.Marshal(myuser)
		// if err != nil{
		// 	log.Fatal(err)
		// }
		// fmt.Println(myuser)
		// fmt.Println(string(data))

		// fp, err := os.OpenFile("entity/Users.json", os.O_RDWR|os.O_CREATE, 0755)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// _, err = fp.Write(data)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fp.Close()

		// fp, err = os.OpenFile("entity/Users.json", os.O_RDONLY, 0755)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// data = make([]byte, 100)
		// n, err := fp.Read(data)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(string(data[:n]))
		// fmt.Println(data[:n])

		// fp.Close()

		// var user2 User
		// err = json.Unmarshal(data[:n], &user2)
		// fmt.Println(user2)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("username", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("mail", "m", "", "mail address")
	registerCmd.Flags().StringP("telephone", "t", "", "telephone")
}
