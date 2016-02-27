/*
	Run RTChecks and on faliure update CloudFlare DNS settings

	Author: Rasmus Jönsson (www.rasmusj.se)
	License: MIT
	February 2016
	https://github.com/rasmusj-se/rtcheck
*/

package main

import (
		"fmt"
		"log"
		"github.com/pearkes/cloudflare"
		"github.com/rasmusj-se/rtcheck"
	    "encoding/json"
	    "io/ioutil"
	    "time"
)

type Config struct {
    APIKey      	string
    EmailAddress	string
}

var config Config

type Check struct {
	Domain					string
	Type					string
	PrimaryDestination		string
	BackupDestination		string
	RTcheck 				rtcheck.Check
} 

/* Read config and start checks*/
func main() {

	fmt.Println("Reading configuration...")

	file, err := ioutil.ReadFile("./cloudflare.conf")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Configuration: OK")

    fmt.Println("Reading checks configuration...")

    json.Unmarshal(file, &config)

    //Read all checks and configure

    files, _ := ioutil.ReadDir("./checks/")
    for _, f := range files {

    	file, err := ioutil.ReadFile("./checks/" + f.Name())
	    if err != nil {
	        log.Fatal(err)
	    }

    	var check Check
        err = json.Unmarshal(file, &check)
        if err != nil {
	        log.Fatal(err)
	    }
        check.RTcheck.OnError = func(time time.Time){
        	fmt.Println(check.Domain,"is failing")
			activateError(check)
		}
		check.RTcheck.OnReturn = func(time time.Time){
			fmt.Println(check.Domain,"has returned")
			deactivateError(check)
		}
        rtcheck.AddCheck(check.RTcheck)
    }

    fmt.Println("Checks are running.")

    select{}
}

/* Update CloudFlare DNS to faile state*/
func activateError(check Check){
	client, err := cloudflare.NewClient(config.EmailAddress, config.APIKey)
	if err != nil {
		log.Fatal("CloudFlare Client: ", err)
	}

	opts := cloudflare.UpdateRecord{
		Type:    check.Type,
		Name:    check.Domain,
		Content: check.BackupDestination,
	}

	record, err := client.RetrieveRecordsByName(check.Domain, check.Domain, false)
	if err != nil {
		log.Fatal("CloudFlare RetrieveRecordsByName: ", err)
	}

	err = client.UpdateRecord(check.Domain, record[0].Id, &opts)
	if err != nil {
		log.Fatal("CloudFlare UpdateRecord: ", err)
	}
}

/* Update CloudFlare DNS to ok state*/
func deactivateError(check Check){
	client, err := cloudflare.NewClient(config.EmailAddress, config.APIKey)
	if err != nil {
		log.Fatal("CloudFlare Client: ", err)
	}

	opts := cloudflare.UpdateRecord{
		Type:    check.Type,
		Name:    check.Domain,
		Content: check.PrimaryDestination,
	}

	record, err := client.RetrieveRecordsByName(check.Domain, check.Domain, false)
	if err != nil {
		log.Fatal("CloudFlare RetrieveRecordsByName: ", err)
	}

	err = client.UpdateRecord(check.Domain, record[0].Id, &opts)
	if err != nil {
		log.Fatal("CloudFlare UpdateRecord: ", err)
	}
}