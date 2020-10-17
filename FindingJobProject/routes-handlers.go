package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024

var account Account

// Job Information
type (
	Job struct {
		JobID          string `json:"JobID"`
		JobName        string `json:"JobName"`
		JobDescription string `json:"JobDescription"`
		JobSalary      string `json:"JobSalary"`
		JobType        string `json:"JobType"`
		JobImage       string `json:"JobImage"`
	}
	//Account login
	Account struct {
		AccName        string `json:"AccountUsername"`
		AccPassword    string `json:"AccountPassword"`
		AccFullNane    string `json:"FullName"`
		AccDateOfBirth string `json:"DateOfBirth"`
		AccEmail       string `json:"Email"`
		AccPhoneNumber string `json:"PhoneNumber"`
		AccTOE         string `json:"TOE"`
		AccIsEmployee  string `json:"IsEmPloyee"`
		AccImage       string `json:"Image"`
		AccJobApplied  string `json:"JobApplied"`
		AccJobPosted   string `json:"JobPosted"`
	}
)

func checkLoginRequest(respone http.ResponseWriter, request *http.Request) {

	var acclogin Account

	acclogin.AccName = request.URL.Query().Get("username")
	acclogin.AccPassword = request.URL.Query().Get("password")

	loginaccept, err := db.Query("SELECT * FROM `accountinfo` WHERE `AccountUsername` = '" + acclogin.AccName + "' AND `AccountPassword` = '" + acclogin.AccPassword + "'")
	for loginaccept.Next() {
		err2 := loginaccept.Scan(&account.AccName, &account.AccPassword, &account.AccFullNane, &account.AccDateOfBirth, &account.AccEmail, &account.AccPhoneNumber, &account.AccTOE, &account.AccIsEmployee, &account.AccImage, &account.AccJobApplied, &account.AccJobPosted)

		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if err != nil {
		fmt.Println(err)
		returnErrorResponse(respone, request)
	}

	if loginaccept != nil {
		if account.AccName == acclogin.AccName && account.AccPassword == acclogin.AccPassword {
			fmt.Fprintf(respone, "Dang Nhap Thanh Cong")

		} else {
			fmt.Fprintf(respone, "Dang Nhap That Bai")

		}

	}

	fmt.Println(acclogin.AccName)
	fmt.Println(acclogin.AccPassword)

}

func getAccount(respone http.ResponseWriter, request *http.Request) {

	var (
		acclogin  Account
		acclogin2 Account
	)

	acclogin.AccName = request.URL.Query().Get("username")
	acclogin.AccPassword = request.URL.Query().Get("password")

	loginaccept, err := db.Query("SELECT * FROM `accountinfo` WHERE `AccountUsername` = '" + acclogin.AccName + "' AND `AccountPassword` = '" + acclogin.AccPassword + "'")
	for loginaccept.Next() {
		err2 := loginaccept.Scan(&acclogin2.AccName, &acclogin2.AccPassword, &acclogin2.AccFullNane, &acclogin2.AccDateOfBirth, &acclogin2.AccEmail, &acclogin2.AccPhoneNumber, &acclogin2.AccTOE, &acclogin2.AccIsEmployee, &acclogin2.AccImage, &acclogin2.AccJobApplied, &acclogin2.AccJobPosted)

		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if err != nil {
		fmt.Println(err)
		returnErrorResponse(respone, request)
	}

	if loginaccept != nil {
		if acclogin2.AccName == acclogin.AccName && acclogin2.AccPassword == acclogin.AccPassword {

			jsonResponse, jsonError := json.Marshal(acclogin2)
			if jsonError != nil {
				fmt.Println(jsonError)
				returnErrorResponse(respone, request)
			}
			if jsonResponse == nil {
				returnErrorResponse(respone, request)
			} else {
				respone.Header().Set("Content-Type", "application/json")
				respone.Write(jsonResponse)
			}
		} else {

		}
	}

	fmt.Println(acclogin.AccName)
	fmt.Println(acclogin.AccPassword)

}

func getJobs(response http.ResponseWriter, request *http.Request) {
	var (
		job  Job
		jobs []Job
	)
	rows, err := db.Query("SELECT * FROM vieclam.job;")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	for rows.Next() {
		rows.Scan(&job.JobID, &job.JobName, &job.JobDescription, &job.JobSalary, &job.JobType, &job.JobImage)
		jobs = append(jobs, job)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(jobs)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}
	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}

}

func checkAccountCreated(respone http.ResponseWriter, request *http.Request, namecheck string) bool {

	createaccept, err := db.Query("SELECT * from `accountinfo` where `AccountUsername` = '" + namecheck + "'")
	for createaccept.Next() {
		err2 := createaccept.Scan(&account.AccName, &account.AccPassword, &account.AccFullNane, &account.AccDateOfBirth, &account.AccEmail, &account.AccPhoneNumber, &account.AccTOE, &account.AccIsEmployee, &account.AccImage, &account.AccJobApplied, &account.AccJobPosted)

		if err2 != nil {
			fmt.Println(err2)
		}
	}
	if err != nil {
		println(err)
	}
	if createaccept != nil {
		if account.AccName == namecheck {

			return true

		} else {

			return false

		}

	}
	return false

}

func createAccount(response http.ResponseWriter, request *http.Request) {
	var NewAccount Account
	NewAccount.AccName = request.URL.Query().Get("username")
	NewAccount.AccPassword = request.URL.Query().Get("password")
	NewAccount.AccFullNane = request.URL.Query().Get("fullname")
	NewAccount.AccDateOfBirth = request.URL.Query().Get("dateofbirth")
	NewAccount.AccEmail = request.URL.Query().Get("email")
	NewAccount.AccPhoneNumber = request.URL.Query().Get("phonenumber")
	NewAccount.AccTOE = request.URL.Query().Get("toe")
	NewAccount.AccIsEmployee = request.URL.Query().Get("isemployee")
	NewAccount.AccImage = request.URL.Query().Get("image")
	NewAccount.AccJobApplied = request.URL.Query().Get("jobapplied")
	NewAccount.AccJobPosted = request.URL.Query().Get("jobposted")
	db.Query("INSERT INTO `accountinfo`(AccountUsername,AccountPassword,Fullname,DateOfBirth,Email,PhoneNumber,TOE,IsEmployee,Image,JobApplied,JobPosted) VALUES ('" + NewAccount.AccName + "','" + NewAccount.AccPassword + "','" + NewAccount.AccFullNane + "','" + NewAccount.AccDateOfBirth + "','" + NewAccount.AccEmail + "','" + NewAccount.AccPhoneNumber + "','" + NewAccount.AccTOE + "','" + NewAccount.AccIsEmployee + "','" + NewAccount.AccImage + "','" + NewAccount.AccJobApplied + "','" + NewAccount.AccJobPosted + "');")
	fmt.Println(NewAccount)
}

func createAccount2(response http.ResponseWriter, request *http.Request) {
	var NewAccount Account
	NewAccount.AccName = request.URL.Query().Get("username")
	NewAccount.AccPassword = request.URL.Query().Get("password")
	NewAccount.AccFullNane = request.URL.Query().Get("fullname")

	NewAccount.AccIsEmployee = request.URL.Query().Get("isemployee")

	if checkAccountCreated(response, request, NewAccount.AccName) == true {
		fmt.Fprintf(response, "Tai khoan da duoc tao tu truoc")
	} else {
		db.Query("INSERT INTO `vieclam`.`accountinfo` (`AccountUsername`, `AccountPassword`, `FullName`, `IsEmployee`) VALUES ('" + NewAccount.AccName + "', '" + NewAccount.AccPassword + "', '" + NewAccount.AccFullNane + "', '" + NewAccount.AccIsEmployee + "'); ")
		fmt.Fprintf(response, "Tao tai khoan thanh cong")
	}

}

func createJob(response http.ResponseWriter, request *http.Request) {

	var NewJob Job
	NewJob.JobName = request.URL.Query().Get("jobname")
	NewJob.JobDescription = request.URL.Query().Get("jobdescription")
	NewJob.JobSalary = request.URL.Query().Get("jobsalary")
	NewJob.JobType = request.URL.Query().Get("jobtype")
	NewJob.JobImage = request.URL.Query().Get("jobimage")
	_, err := db.Query("INSERT INTO `job`(JobName,JobDescription,JobSalary,JobType,JobImage) VALUES('" + NewJob.JobName + "','" + NewJob.JobDescription + "','" + NewJob.JobSalary + "','" + NewJob.JobType + "','" + NewJob.JobImage + "');")
	if err != nil {
		fmt.Println(err)
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("It's not you it's me.")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join("./tmp", fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func applyJob(w http.ResponseWriter, request *http.Request) {
	var (
		account Account
		job     string
	)

	account.AccName = request.URL.Query().Get("username")
	account.AccPassword = request.URL.Query().Get("password")
	job = request.URL.Query().Get("id")
	db.Query("UPDATE `vieclam`.`accountinfo` SET `JobApplied` = '" + job + "' WHERE (`AccountUsername` = '" + account.AccName + "' AND `AccountPassword` = '" + account.AccPassword + "')")
	if err != nil {
		fmt.Println(err)
	}
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func getPostedjob(response http.ResponseWriter, request *http.Request) {

	var job Job

	job.JobID = request.URL.Query().Get("id")

	loginaccept, err := db.Query("SELECT * FROM `job` WHERE `JobID` = '" + job.JobID + "'")
	for loginaccept.Next() {
		err2 := loginaccept.Scan(&job.JobID, &job.JobName, &job.JobDescription, &job.JobSalary, &job.JobType, &job.JobImage)

		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if err != nil {
		fmt.Println(err)

	}

	if loginaccept != nil {

		jsonResponse, jsonError := json.Marshal(job)
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}

	}

}

func isExist(response http.ResponseWriter, request *http.Request, jobid string) bool {
	var job Job
	jobcreated, err := db.Query("SELECT * from `job` where `JobID` = '" + jobid + "'")
	for jobcreated.Next() {
		err2 := jobcreated.Scan(&job.JobID, &job.JobName, &job.JobDescription, &job.JobSalary, &job.JobType, &job.JobImage)

		if err2 != nil {
			fmt.Println(err2)
		}
	}
	if err != nil {
		println(err)
	}
	if jobcreated != nil {
		if job.JobID == jobid {

			return true

		} else {

			return false

		}

	}
	return false
}

func checkJobExist(response http.ResponseWriter, request *http.Request) {
	var checkIdJob string
	checkIdJob = request.URL.Query().Get("id")
	if isExist(response, request, checkIdJob) {
		fmt.Fprintf(response, "Co the nhan cong viec nay")
	} else {
		fmt.Fprintf(response, "Cong viec nay khong ton tai")
	}
}

func removeJobApplied(response http.ResponseWriter, request *http.Request) {
	var acclogin Account

	acclogin.AccName = request.URL.Query().Get("username")
	acclogin.AccPassword = request.URL.Query().Get("password")

	loginaccept, err := db.Query("SELECT * FROM `accountinfo` WHERE `AccountUsername` = '" + acclogin.AccName + "' AND `AccountPassword` = '" + acclogin.AccPassword + "'")
	for loginaccept.Next() {
		err2 := loginaccept.Scan(&account.AccName, &account.AccPassword, &account.AccFullNane, &account.AccDateOfBirth, &account.AccEmail, &account.AccPhoneNumber, &account.AccTOE, &account.AccIsEmployee, &account.AccImage, &account.AccJobApplied, &account.AccJobPosted)

		if err2 != nil {
			fmt.Println(err2)
		}
	}

	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	_, err3 := db.Query("UPDATE `vieclam`.`accountinfo` SET `JobApplied` = '' WHERE (`AccountUsername` = '" + acclogin.AccName + "' AND `AccountPassword` = '" + acclogin.AccPassword + "')")
	if err3 != nil {
		fmt.Println(err3)
	}
}
