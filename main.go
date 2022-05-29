package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// shifts document are present in ShiftsForAttendance  ... In these shifts documents employees shifts data are present

var dateOfAttendance string
var leaveApproval string
var superVisorApprovalsOrNotForEmployee string
var updateValuesForEmployeeApprovals string

var documentNameForShifts string

type employeeDetailes struct {
	shift               int
	employeeName        string
	employeeLogInTime   string
	employeeLateLogins  string
	employeeOTInLogin   string
	employeeEarlyLogout string
	employeeLogOut      string
	employeeOTinLogout  string
	resonOrApproval     string
	reqleaveApproval    string
	superVisor          string
	Date                string
}

var doc1, doc2, doc3, doc4, doc5, doc6, doc7, doc8, doc9, doc10, doc11, doc12, doc13, doc16 interface{}

var i string
var start string

var fieldName string
var documentName string

var employeeAttendanceDocuments string

//case1 Varibles
var shifts int
var supervisorNameForNeeds string

// var employeeName string
var employeeLogInTimeToStore string
var employeeLogOutTimeToStore string
var employeeOverTimeWork string
var supervisorName string
var Date string
var userEnterPassword int
var timeForMarkAttendace string
var shiftNumber, employeeName, superVisorName, HR interface{}
var j interface{}

var startingDayTime = "00:00"
var endingDayTime = "23:59"

// Office login Timeings
var beforeOfficeLoginTime string
var InTimeofficeLogin string
var endTimeofficeLogin string
var lateLogin1 string
var end string

// office Logout time Timings

var officeLogoutTime string

var officeLogoutEndTiming string

var ToLateLogoutTime string

// consider Normal Logins
var considerbeforeOfficeLoginTime string
var considerbeforeOfficeLoginTime1 string

// consider Normal Logouts
var considerbeforeOfficeLogoutTime string
var considerbeforeOfficeLogoutTime1 string

// var considerbeforeOfficeLogoutTime = "20:01"
// var considerbeforeOfficeLogoutTime1 = "20:29"

// 	var considerbeforeOfficeLoginTime = "10:31"
// 	var considerbeforeOfficeLoginTime1 = "10:59"

var reasonORApprovalForLogin string
var reasonORApprovalForLogout string
var toCheckPastDetailes string

var loginForHrandManager string
var logoutForHrandManager string
var Designation string
var valueNameForAllocateshifts string
var valueName string
var filedNameForemploNeeds string
var valueNameForemploNeed string

// case 1 for DataBase

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err

}
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

//case 3 Validation

// These function Authenticate the user enterd UID with Company Provided UID
func welcomePortal() {

	fmt.Println("welocme To Attendance Portal")

	var currentTime1 = time.Now().Format("15:04")
	fmt.Println(currentTime1)
	fmt.Println("Welcome To Attendance Portal\n1.Employee\n2.Supervisor\n3.Manager\n4.H.R")

	//Team 1
	var programmerEmployeeShift1 = 123
	var developerEmployeeShift1 = 342

	var programmerEmployeeShift2 = 145
	var developerEmployeeShift2 = 789

	var superVisor = 231
	// var superVisor1 = 143

	var manager = 678
	var HR = 7878

	// fmt.Println("Please ENter Date")
	// fmt.Scan(&dateOfAttendance)

	// // Team 2

	var programmerEmployeeShift1Team2 = 909
	var developerEmployeeShift1Team2 = 709

	var programmerEmployeeShift2Team2 = 129
	var developerEmployeeShift2Team2 = 329

	var superVisorTeam2 = 143

	fmt.Println("Please Unter Your Unique Id")
	fmt.Scan(&userEnterPassword)

	if userEnterPassword == programmerEmployeeShift1 {
		fmt.Println("Hello Shaik Baji Good Morning and Hi Programmer")
		employeeTaks("superVisorName", "shifts", "employees")
	} else if userEnterPassword == developerEmployeeShift1 {

		fmt.Println("Hello SHaik Kalesha Vali Good Morning and Hi Programmer ")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == programmerEmployeeShift2 {
		fmt.Println("Hello Shaik Samad Good Morning and Hi Coder")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == developerEmployeeShift2 {
		fmt.Println("Hello shaik Jani Good Morning and Hi Coder")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == superVisor {
		fmt.Println("Hello Rajesh Good Moring supervisor")
		supervisorNameForNeeds = "Rajesh"
		superVisorTasks("superVisorName", "Rajesh", "shifts")

		// Team 2
	} else if userEnterPassword == programmerEmployeeShift1Team2 {
		fmt.Println("Hello Navazish Good Moring You are a Team 2 Member")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == developerEmployeeShift1Team2 {
		fmt.Println("Hello Eliyaz Good Moring You are a Team 2 Member")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == programmerEmployeeShift2Team2 {
		fmt.Println("Hello Yunus Good Moring You are a Team 2 Member")

		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == developerEmployeeShift2Team2 {
		fmt.Println("Hello Shukur Good Moring You are a Team 2 Member")
		employeeTaks("superVisorName", "shifts", "employees")

	} else if userEnterPassword == superVisorTeam2 {

		fmt.Println("Hello Rafi Good Moring supervisor")
		supervisorNameForNeeds = "Rafi"
		superVisorTasks("superVisorName", "Rafi", "shifts")

	} else if userEnterPassword == manager {
		fmt.Println("Hello Basha Good Morinng ..Have nice day Mananger")
		supervisorNameForNeeds = "Basha"
		managerTasks("managerName", "Basha", "shiftsForSupervisor")

	} else if userEnterPassword == HR {
		fmt.Println("Hello Azmathullah Good Moring and Have A Nice DAY")
		HRTasks()
	}

}

// case 4
// These functions used employee Mark Attendance And stored in DataBase
//1.Employee Tasks Are ENterd UID
//check shifts and Mark Attendance And if need check Past Detailes
// Ask For OTS and Leavs

func employeeTaks(fieldName1 string, documentNameForShifts string, employeeAttendanceDocuments1 string) {

	fmt.Println("Please Enter Date")
	fmt.Scan(&dateOfAttendance)

	// fleid Name In Shift document of shiftsupervisor db//
	fieldName = "superVisorName"
	// document in shiftsupervisor db
	documentNameForShifts = "shifts" // To Retrive Shift
	// To Store data in document
	employeeAttendanceDocuments = "employees" // -->Document Name in employeAttendance db

	Designation = "employee"
	checkingShifts(documentNameForShifts, "shifts")
	fmt.Println("If u want To check Ur Past Detailes Please Enterd One")
	fmt.Scan(&toCheckPastDetailes)
	if toCheckPastDetailes == "1" {

		checkPastDetailes()
	} else {

		fmt.Println("Thank you")

	}

}

//case 5
// superVisor Tasks are
// 1. Login with UID then check shifts and then mark Attendance
// 2. check OTS , Late Logins ,Early Logouts and check Leaves Requests From employees
// 3. Provides Approvals with respected rules
func superVisorTasks(filedNameForShifts string, ValueNameForShifts string, documnetforShifts string) {
	fmt.Println("Please Enter Date")
	fmt.Scan(&dateOfAttendance)

	// fieldName In documents of employeeAttendance db
	filedNameForemploNeeds = "superVisorName"

	// To Store In Db While SuperVisor As A USer
	fieldName = "managerName"
	documentNameForShifts = "shiftsForSupervisor"       // -->Document From ShiftAttendance DB
	employeeAttendanceDocuments = "superVisorEmployees" //-->To store Data In Employee Attendacne -->DocumnetName
	Designation = "superVisor"
	checkingShifts(documentNameForShifts, "shiftsForSupervisor")
	checkingEmployeesNeeds("employees")
	// superVisorTasks("superVisorName", "Rafi", "shifts")
	AllocatesShifts(filedNameForShifts, ValueNameForShifts, documnetforShifts)
	// 1. filedNameForShifts -->Is filedName
	//2.valueNameForShifts -->IS Values
	//3.documentForShifts-->is Document ..
	fmt.Println("Do u want To Check Past Detailes Please press 1")
	fmt.Scan(&toCheckPastDetailes)
	if toCheckPastDetailes == "1" {

		checkPastDetailes()
	} else {

		fmt.Println("Thank you")

	}

}

//case 6

// managerTasks are
// 1.Login with UID and Mark Attendance
// 2.checks OTs, Late Logins ,Early logouts and check Leaves requests From supervisors
// 3.Provides Approvals With respected Rules
func managerTasks(filedNameForShifts string, ValueNameForShifts string, documnetforShifts string) {

	fieldName = "HR"
	employeeAttendanceDocuments = "managerAttendance"
	Designation = "manager"
	filedNameForemploNeeds = "managerName"
	loginTimingsForHRAndManager("manager")
	checkingEmployeesNeeds("superVisorEmployees")
	// managerTasks("managerName", "Basha", "shiftsForSupervisor")
	AllocatesShifts(filedNameForShifts, ValueNameForShifts, documnetforShifts)
	checkPastDetailes()

}
func HRTasks() {

	fieldName = "HrName"
	employeeAttendanceDocuments = "HRAttendance"
	Designation = "HR"
	loginTimingsForHRAndManager("HR")
	HrpaySalary()

}

func main() {

	welcomePortal()

}

// case 7 checking shifts for employees and superVisor

// These function checks the shifts of Employees and superVisor and mark the Attendance
func checkingShifts(documentForToRetriveTheshifts string, ToRetriveShifts string) {
	var markAttendanceOrShiftDetailes int

	var timeToLogin string
	var timeToLogout string
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	fmt.Println("Press 1 To Move To Mark Attendance  and Press 2 To know shift Detailes ")
	fmt.Scan(&markAttendanceOrShiftDetailes)
	if markAttendanceOrShiftDetailes == 1 {

		var filter, option interface{}
		filter = bson.D{{"UID", bson.D{{"$eq", userEnterPassword}}}}

		// Here documentForToRetriveTheshifts are Document Name To Retrive Shifts
		//documentForToRetriveTheshifts == shifts //

		cursor, err := query(client, ctx, "ShiftsForAttendance", documentForToRetriveTheshifts, filter, option)
		if err != nil {

			panic(err)
		}
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}

		for _, doc := range result {

			shiftNumber = doc[1].Value
			employeeName = doc[2].Value
			superVisorName = doc[4].Value
		}
		if shiftNumber == "shift1" {
			fmt.Println("Hii", employeeName, "and Know Your In shift", shiftNumber, "\n...!! The  Shift One Timings Are 09:00 To 17:30", "\n and your superVisorName is ", superVisorName)
			fmt.Println("Shift one Login Timings is 09:00 T0 09:20")
			fmt.Println("Please Enter Time Login Time")
			fmt.Scan(&timeToLogin)
			loginAttendanceForShift1(timeToLogin)
			fmt.Println("SHift one Logout Timings is 17:30 To 18:00")
			fmt.Println("Plese Enter Logout Time")
			fmt.Scan(&timeToLogout)
			logoutAttendanceForShift1(timeToLogout)

		} else if shiftNumber == "shift2" {

			fmt.Println("Hii", employeeName, "and Know Your In shift", shiftNumber, "\n......!!!The Shift Two Timings Is 11:00 To 20:30")
			fmt.Println("Please Enter Time Login Time")
			fmt.Println("Shift Two Logint Timings are 11:00 To 11:20")
			fmt.Scan(&timeToLogin)
			loginAttendanceForShift2(timeToLogin)
			fmt.Println("Shift Two logout Timingsare 19:30 To 20:00")
			fmt.Println("Plese Enter Logout Time")
			fmt.Scan(&timeToLogout)
			logoutAttendanceForShift2(timeToLogout)

		}
	} else if markAttendanceOrShiftDetailes == 2 {

		var filter, option interface{}
		filter = bson.D{{}}

		option = bson.D{{"UID", 0}}
		cursor, err := query(client, ctx, "ShiftsForAttendance", ToRetriveShifts, filter, option)
		if err != nil {

			panic(err)
		}
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}
		for _, doc := range result {

			fmt.Println("The Enter Shifts Detailes Are ", doc)
		}

	}
}

// case 8 --> Login Attendance <--

// The function Follow Login crendentials with related shift1

func loginAttendanceForShift1(currentTime string) {

	considerbeforeOfficeLoginTime = "08:31"
	considerbeforeOfficeLoginTime1 = "08:59"
	beforeOfficeLoginTime = "08:30"
	InTimeofficeLogin = "09:00"
	endTimeofficeLogin = "09:20"
	lateLogin1 = "09:21"
	end = "11:00"
	officeLogins(beforeOfficeLoginTime, InTimeofficeLogin, endTimeofficeLogin, lateLogin1, end, currentTime, considerbeforeOfficeLoginTime,
		considerbeforeOfficeLoginTime1)
}

// The function Follow Login crendentials with related shift2

func loginAttendanceForShift2(currentTime string) {

	considerbeforeOfficeLoginTime = "10:31"
	considerbeforeOfficeLoginTime1 = "10:59"
	beforeOfficeLoginTime = "10:30"
	InTimeofficeLogin = "11:00"
	endTimeofficeLogin = "11:20"
	lateLogin1 = "11:21"
	end = "01:00"

	officeLogins(beforeOfficeLoginTime, InTimeofficeLogin, endTimeofficeLogin, lateLogin1, end, currentTime, considerbeforeOfficeLoginTime, considerbeforeOfficeLoginTime1)
}

// case 9 LOgout Attendace
// The function Follow Logout crendentials with related shift1
func logoutAttendanceForShift1(currentTime string) {

	officeLogoutTime = "17:30"

	officeLogoutEndTiming = "18:00"

	ToLateLogoutTime = "18:30"
	considerbeforeOfficeLogoutTime = "18:01"
	considerbeforeOfficeLogoutTime1 = "18:29"
	officeLogouts(currentTime, officeLogoutTime, officeLogoutEndTiming, ToLateLogoutTime, considerbeforeOfficeLogoutTime, considerbeforeOfficeLogoutTime1)

}

// The function Follow Logout crendentials with related shift2
func logoutAttendanceForShift2(currentTime string) {

	officeLogoutTime = "19:30"

	officeLogoutEndTiming = "20:00"

	ToLateLogoutTime = "20:30"

	// consider Normal Logouts

	considerbeforeOfficeLogoutTime = "20:01"
	considerbeforeOfficeLogoutTime1 = "20:29"

	officeLogouts(currentTime, officeLogoutTime, officeLogoutEndTiming, ToLateLogoutTime, considerbeforeOfficeLogoutTime, considerbeforeOfficeLogoutTime1)

}

// case10 Office Logins

// These function validte the user enterd Login Times
func officeLogins(beforeOfficeLoginTime string, InTimeofficeLogin string, endTimeofficeLogin string, lateLogin1 string, end string, currentTime string, considerbeforeOfficeLoginTime string, considerbeforeOfficeLoginTime1 string) {
	// Validation For OTs before Logins
	var OTvalidationInLogin = currentTime <= beforeOfficeLoginTime && currentTime >= startingDayTime

	// validation For Exact Office Timings
	var validationForexactOfficeTime = currentTime >= InTimeofficeLogin && currentTime <= endTimeofficeLogin

	//validation For Late Logins
	var LateValidationLoginTime = currentTime >= lateLogin1 && currentTime <= end

	// In between some Time Duration Consider As Normal Logins
	var considerAsNormalLogin = currentTime >= considerbeforeOfficeLoginTime && currentTime <= considerbeforeOfficeLoginTime1

	// 1.currentTime -- user Login Time/
	// 2.employee Late Login
	// 3.employee OT Login
	// 4.employee early Logout
	// 5.employee LogoutTime
	// 6.employee Ot In logout
	// 7.Reason
	// 8.SuperName
	// 9.Date
	if OTvalidationInLogin {
		fmt.Println("you are eligine for OT and Login Time is ", currentTime)
		fmt.Println("Congrants You Are eligible for OT  To report To your superVisor ...Please Enter RequiredApproval For OT")
		fmt.Scan(&reasonORApprovalForLogin)
		fmt.Println("Do u want To Request To Leave Approval")
		fmt.Scan(&leaveApproval)

		employeeLoginandLogoutTimeInDB(currentTime, "0", "0", "0", "0", "0", reasonORApprovalForLogin, leaveApproval, supervisorName, dateOfAttendance)

	} else if validationForexactOfficeTime {

		fmt.Println("you are Login in Exact Office Timings and Your Login Time Is ", currentTime)
		fmt.Println("Do u want To Request To Leave Approval")
		fmt.Scan(&leaveApproval)
		employeeLoginandLogoutTimeInDB(currentTime, "0", "0", "0", "0", "0", "No", leaveApproval, supervisorName, dateOfAttendance)

	} else if LateValidationLoginTime {
		fmt.Println("ypur are Late Login and Login Time is", currentTime)
		fmt.Println("Please enter valid Reason")
		fmt.Scan(&reasonORApprovalForLogin)
		fmt.Println("Do u want To Request To Leave Approval")
		fmt.Scan(&leaveApproval)
		employeeLoginandLogoutTimeInDB(currentTime, "0", "0", "0", "0", "0", reasonORApprovalForLogin, leaveApproval, supervisorName, dateOfAttendance)

	} else if considerAsNormalLogin {
		fmt.Println("Sorry you are Normal Login and Not Applicable To OT ", currentTime)
		fmt.Println("Do u want To Request To Leave Approval")
		fmt.Scan(&leaveApproval)
		employeeLoginandLogoutTimeInDB(currentTime, "0", "0", "0", "0", "0", "No", leaveApproval, supervisorName, dateOfAttendance)

	} else {
		fmt.Println("Please Enter A Valid Time")

	}

}

// case12 office Logouts

// These function validte the user enterd Logout Times
func officeLogouts(currentTime string, officeLogoutTime string, officeLogoutEndTiming string, ToLateLogoutTime string, considerbeforeOfficeLogoutTime string, considerbeforeOfficeLogoutTime1 string) {

	// Validation For Early Logout
	var EarlyLogout = currentTime < officeLogoutTime && currentTime >= lateLogin1

	// validation For Exact logout Timings
	var ExactLogoutValidation = currentTime >= officeLogoutTime && currentTime <= officeLogoutEndTiming

	// Validation For Late Logout and applay OTs
	var OTvalidationInLogout = currentTime >= ToLateLogoutTime && currentTime <= endingDayTime

	// In between some Time Duration Consider As Normal Logout
	var considerAsNormalLogout = currentTime >= considerbeforeOfficeLogoutTime && currentTime <= considerbeforeOfficeLogoutTime1
	if EarlyLogout {
		fmt.Println("your are early logout Please Let Me Know Reason and your Logout Time is ", currentTime)
		fmt.Scan(&reasonORApprovalForLogout)
		updateemployeeLogout(currentTime, "0", "0", reasonORApprovalForLogout)

	} else if ExactLogoutValidation {
		fmt.Println("your are Logout  at Office Timings and time is", currentTime)
		updateemployeeLogout(currentTime, "0", "0", "NO")

	} else if OTvalidationInLogout {

		fmt.Println("you are logout To Late and LogOut Time is and you are eligibe for OT", currentTime)
		updateemployeeLogout("0", "17:30 To 18:00", currentTime, "RequiredApprovalForOT")

	} else if considerAsNormalLogout {
		fmt.Println("Sorry To say you are normal logout ..Please Try Again For OT")
		updateemployeeLogout("0", currentTime, "0", "Normal")
	}

}

//  case13 ManagerAndHr Logins And Logouts
func loginTimingsForHRAndManager(documentForToRetriveTheshifts string) {
	fmt.Println("Please Enter date")
	fmt.Scan(&dateOfAttendance)
	fmt.Println("Plese Enter the Login Time")
	fmt.Scan(&loginForHrandManager)
	fmt.Println("Do u want To Leave")
	fmt.Scan(&leaveApproval)

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	fmt.Println("Press 1 To Move To Mark Attendance  and Press 2 To know shift Detailes ")
	fmt.Scan(&shifts)
	if shifts == 1 {

		var filter, option interface{}
		filter = bson.D{{"UID", bson.D{{"$eq", userEnterPassword}}}}

		cursor, err := query(client, ctx, "ShiftsForAttendance", documentForToRetriveTheshifts, filter, option)
		if err != nil {

			panic(err)
		}
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}

		for _, doc := range result {

			employeeName = doc[1].Value
			HR = doc[3].Value

		}
		fmt.Println("The login Employee Name is ", employeeName, "And your Hr is", HR)

	}

	var document interface{}
	document = bson.D{
		{"employeeName", employeeName},
		{"employeeLoginTime", loginForHrandManager},
		{"RequestForLeave", leaveApproval},
		{"LeaveApprovals", "No"},
		{"HRName", HR},
		{"DateOfAttendance", dateOfAttendance},
		{"employeeDesignation", Designation},
	}
	insertOneResult, err := insertOne(client, ctx, "employeeAttendance", employeeAttendanceDocuments, document)
	fmt.Println(insertOneResult)
}

//case6 employee Whole data To Store in DB
// These function is used to stored Employee Login  Detailes
func employeeLoginandLogoutTimeInDB(empLogInTime string, empLatelogin string, empOTlogin string, employeEarlyLogout string, emplLogoutTime string, emplOTinlogout string, reasonOrApproval string, reqLeavApproval string, supervi string, date string) {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	var em = employeeDetailes{
		employeeLogInTime:   empLogInTime,
		employeeLateLogins:  empLatelogin,
		employeeOTInLogin:   empOTlogin,
		employeeEarlyLogout: employeEarlyLogout,
		employeeLogOut:      emplLogoutTime,
		employeeOTinLogout:  emplOTinlogout,
		resonOrApproval:     reasonOrApproval,
		reqleaveApproval:    reqLeavApproval,
		superVisor:          supervi,
		Date:                date,
	}

	var document interface{}
	document = bson.D{
		{"employeeShift", shiftNumber},
		{"employeeName", employeeName},
		{"employeeLogInTime", em.employeeLogInTime},
		{"employeeApplicableForOTs", em.employeeOTInLogin},
		{"employeeLateLogins", em.employeeLateLogins},
		{"employeeLogoutTime", em.employeeOTinLogout},
		{"employeeEarlyLogoutTimes", em.employeeEarlyLogout},
		{"employeeOTinLogoutTime", "0"},
		{"employeeReasonsOrApprovalsForLogins", em.resonOrApproval},
		{"employeeReasonsOrApprovalsForLogouts", "No"},
		{"RequestForLeave", reqLeavApproval},
		{"LeaveApproval", "No"},
		{fieldName, superVisorName},
		{"DateOfAttendance", em.Date},
		{"employeeDesignation", Designation},
	}

	insertOneResult, err := insertOne(client, ctx, "employeeAttendance", employeeAttendanceDocuments, document)
	fmt.Println(insertOneResult)

}

// struct data type for store employee logout detailes in DB
type employeL struct {
	empEarlyLogout string
	empLogoutTime  string
	empOtlogout    string
	reas           string
}

// These function stored the employee Logout detailes
func updateemployeeLogout(empNormalLogout string, empEarlyLogout string, empOTLogout string, reasons string) {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	var employ = employeL{

		empLogoutTime:  empNormalLogout,
		empEarlyLogout: empEarlyLogout,
		empOtlogout:    empOTLogout,
		reas:           reasons,
	}

	filter := bson.D{{"DateOfAttendance", bson.D{{"$eq", dateOfAttendance}}}}
	update := bson.D{{"$set", bson.D{{"employeeLogoutTime", employ.empLogoutTime}, {"employeeEarlyLogoutTimes", employ.empEarlyLogout}, {"employeeOTinLogoutTime", employ.empOtlogout}, {"employeeReasonsOrApprovalsForLogouts", employ.reas}}}}

	fmt.Println("employeeLogoutTime", empNormalLogout)

	result, err := UpdateOne(client, ctx, "employeeAttendance", employeeAttendanceDocuments, filter, update)

	fmt.Println(result)

}

// These function provides employee past detailes
func checkPastDetailes() {

	// retreive data from db upon date or employee Name
	var retrieveDatafromdb int
	fmt.Println("Do u want To Know whole Past Detailes or Particular Day \n 1.if u want know particular day please press 1 and press 2 for whole data")
	fmt.Scan(&retrieveDatafromdb)
	if retrieveDatafromdb == 1 {

		//Retrive Data From Db Upon Date
		var retreiveDatafromdbUponDate string

		fmt.Println("which Date Do U want to know data")
		fmt.Scan(&retreiveDatafromdbUponDate)

		client, ctx, cancel, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}

		defer close(client, ctx, cancel)
		var filter, option interface{}

		filter = bson.D{{"DateOfAttendance", bson.D{{"$eq", retreiveDatafromdbUponDate}}}}
		option = bson.D{{"UID", 0}}
		cursor, err := query(client, ctx, "employeeAttendance", employeeAttendanceDocuments, filter, option)
		if err != nil {

			panic(err)
		}
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}
		for _, doc := range result {

			fmt.Println("The Enter Shifts Detailes Are ,,", doc)
		}

	} else if retrieveDatafromdb == 2 {
		//Retrive data from db upon Name
		var employeeNameforPastDetailes string

		client, ctx, cancel, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}

		defer close(client, ctx, cancel)

		var filter, option interface{}
		fmt.Println("Do u want To Know Your Past Detailes Please Enter Your employeeName")
		fmt.Scan(&employeeNameforPastDetailes)

		filter = bson.D{{"employeeName", bson.D{{"$eq", employeeNameforPastDetailes}}}}
		option = bson.D{{"UID", 0}}
		cursor, err := query(client, ctx, "employeeAttendance", employeeAttendanceDocuments, filter, option)
		if err != nil {

			panic(err)
		}
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}
		for _, doc := range result {

			fmt.Println("The Enter Shifts Detailes Are ,,", doc)

		}

	}

}

// These function Provides employee reports Like Ots, Leaves Requests...

func checkingEmployeesNeeds(checkToemployeedata string) {

	// var superVisorApprovals string
	// var superVisorApprovalsOrNotForEmployee string

	var searchDateForApprovals string
	var updateValuesForEmployeeApprovals string

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	fmt.Println("which date did u want search About employees")
	fmt.Scan(&searchDateForApprovals)
	var filter, option interface{}
	filter = bson.D{{"DateOfAttendance", bson.D{{"$eq", searchDateForApprovals}}}, {filedNameForemploNeeds, bson.D{{"$eq", supervisorNameForNeeds}}}}
	cursor, err := query(client, ctx, "employeeAttendance", checkToemployeedata, filter, option)
	if err != nil {

		panic(err)
	}
	var result []bson.D
	if err := cursor.All(ctx, &result); err != nil {

		panic(err)
	}

	for _, doc := range result {

		fmt.Println(doc[1].Value)
		fmt.Println(doc[2].Value)

		fmt.Println("Your under Reported employees are\n The Employee Name is ", doc[2].Value, "And his Present Shift IS ", doc[1].Value)
		fmt.Println("The Enter Shifts Detailes Are ", doc[2].Value)
		// doc0 = doc[0].Value   // Object ID
		doc1 = doc[1].Value   // Shift Number
		doc2 = doc[2].Value   //Name of employee
		doc3 = doc[3].Value   //employee Login Time
		doc4 = doc[4].Value   // Applicable For OTs
		doc5 = doc[5].Value   // Late Logins
		doc6 = doc[6].Value   // Logout Time
		doc7 = doc[7].Value   // Acceprt Or Rejects in Early Logouts
		doc8 = doc[8].Value   // Ot in logout Time
		doc9 = doc[9].Value   //Requsets or Resons in Logins
		doc10 = doc[10].Value //Requests or Resons in Logouts
		doc11 = doc[11].Value //Requests For Leave
		doc12 = doc[12].Value // Leave Aprovals
		// doc13 = doc[13].Value
		// // doc14 = doc[14].Value //Approval For Leave
		// // doc15 = doc[15].Value //Higher Employee Name Supervisor / Manager
		// doc16 = doc[16].Value // Date of Attendance
		// // doc17 = doc[17].Value // employee Designation

	}

	fmt.Println("The Employee Shift Number is", doc1, "And Name is ", doc2)
	// fmt.Println("Shift Login Timings are", doc3, "\nShift Logout Timings Are ", doc6)
	fmt.Println("The", doc2, "Login Time is ", doc3, "\n The", doc2, "Logout Time is", doc6)
	fmt.Println("The", doc2, "OTs are", doc4, "And Late Logins are", doc5, "\nEarly Logouts are", doc7, "\nOTs In Logouts", doc8)
	fmt.Println("The Reasons for Logins are", doc9, "The Resaons for Logouts are", doc10)
	fmt.Println("The Requests for Leave Approvals are", doc11)
	// fmt.Println("The Date of Attendace is ", doc16)
	// fmt.Println("The ")
	var Approvals string
	var acceptsOrRejects string
	var k int
	var j int
	fmt.Println("do u want start Approves or Rejects Employee Needs Please  press 1")
	fmt.Scan(&j)

	for k = j; k <= 6; k++ {
		fmt.Println("1.To Approves OTS \n2.Accept Or Reject Late Logins \n3.Accept Or Reject Early Logouts \n4. Approves OT In Logot Timings \n5.Approves or Rejects Leavs Requests\n6.Quit the process")
		fmt.Scan(&Approvals)
		switch Approvals {

		case "1":

			fmt.Println(doc2, "Is login Time is ", doc3, "and he is applied For OT", doc9)
			fmt.Println("Do u want To Approves  His OT")
			fmt.Scan(&updateValuesForEmployeeApprovals)
			continue
		case "2":
			fmt.Println(doc2, "Is login Time is ", doc3)
			fmt.Println("Is there Any Accepts or Rejects in Login Time If IS There ENters 1")
			fmt.Scan(&acceptsOrRejects)
			if acceptsOrRejects == "1" {

				fmt.Println("Please Enter Accepts or Rejects in Late logins")
				fmt.Scan(&updateValuesForEmployeeApprovals)

			}

		case "3":
			fmt.Println(doc2, "Is Logout Time is", doc6)
			fmt.Println("Is There Any Accepts or Rejects in Logout Timings If IS There ENters 1")
			fmt.Scan(&acceptsOrRejects)
			if acceptsOrRejects == "1" {

				fmt.Println("Do u Accept Or Reject In Early Logouts")
				fmt.Scan(&updateValuesForEmployeeApprovals)

			}
		case "4":
			fmt.Println(doc2, "is Logout Time is", doc6)
			fmt.Println("Is there Any Approvals For Ots Or Not If IS There ENters 1")
			fmt.Scan(&acceptsOrRejects)
			if acceptsOrRejects == "1" {

				fmt.Println("Do u Accept or Rejects OT In Logout Timings")
				fmt.Scan(&updateValuesForEmployeeApprovals)

			}
		case "5":
			fmt.Println(doc2, "Is Request To Applay For Leave Do u Want To apporves his OT")

			fmt.Scan(&updateValuesForEmployeeApprovals)

		case "6":

			k = 6

		}

	}

	update := bson.D{{"$set", bson.D{{"employeeApplicableForOTs", updateValuesForEmployeeApprovals}, {"employeeLateLogins", updateValuesForEmployeeApprovals}, {"employeeEarlyLogoutTimes", updateValuesForEmployeeApprovals}, {"employeeOTinLogoutTime", updateValuesForEmployeeApprovals}, {"LeaveApproval", updateValuesForEmployeeApprovals}}}}

	UpdateOne(client, ctx, "employeeAttendance", "employees", filter, update)

}

// func for To Approval For Not For employees
func superVisorApprovalsOrNot() {

	if superVisorApprovalsOrNotForEmployee == "1" {
		updateValuesForEmployeeApprovals = "1"

	} else {
		updateValuesForEmployeeApprovals = "0"

	}

}

// The function shows Employees of superVisor
func AllocatesShifts(filedNameForAllocateShifts string, valueForAllocateShifts string, shiftDocumnet string) {
	var changeShiftsOfEmployee string
	var modifiedShifts string
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	var filter, option interface{}
	filter = bson.D{{filedNameForAllocateShifts, bson.D{{"$eq", valueForAllocateShifts}}}}
	cursor, err := query(client, ctx, "ShiftsForAttendance", shiftDocumnet, filter, option)
	if err != nil {

		panic(err)
	}
	var result []bson.D
	if err := cursor.All(ctx, &result); err != nil {

		panic(err)
	}
	var doc1, doc2 interface{}
	for _, doc := range result {

		// doc0 = doc[0].Value
		doc1 = doc[1].Value
		doc2 = doc[2].Value
		// doc3 = doc[3].Value

		fmt.Println("yours Reported Employee Are ", doc2, "And his shifts is", doc1)
		fmt.Println("Do u want To change Shifts Of The Employees  Press 1 To change Shifts And Press 2 To Move Main menu")
		fmt.Scan(&changeShiftsOfEmployee)
	}
	if changeShiftsOfEmployee == "1" {

		fmt.Println("Please Enter Name of employee To change your Required Shifts")
		fmt.Scan(&changeShiftsOfEmployee)
		fmt.Println("Please Enter changed Shifts")
		fmt.Scan(&modifiedShifts)
		if modifiedShifts == "1" {

			filter := bson.D{{filedNameForAllocateShifts, valueForAllocateShifts}}
			update := bson.D{{"$set", bson.D{{"shift", modifiedShifts}}}, {"$set", bson.D{{"shiftOneLoginTimings", "09:00 To 09:20"}}}, {"$set", bson.D{{"shiftOneLogoutTimings", "17:30 To 18:00"}}}}

			UpdateOne(client, ctx, "ShiftsForAttendance", shiftDocumnet, filter, update)
		} else if modifiedShifts == "2" {

			filter := bson.D{{filedNameForAllocateShifts, valueForAllocateShifts}}
			update := bson.D{{"$set", bson.D{{"shift", modifiedShifts}}}, {"$set", bson.D{{"shiftTwoLoginTimings", "11:00 To 11:20"}}}, {"$set", bson.D{{"shiftTwoLogoutTimings", "20:00 To 20:30"}}}}

			UpdateOne(client, ctx, "ShiftsForAttendance", shiftDocumnet, filter, update)
		}

	}
}
func HrpaySalary() {

	var employeeNameToRetrivewData string
	var paySalaries string
	var selectDestination string
	var paySalaryForEmployees string
	var salaryToEmployee string

	fmt.Println("Do u want To retrive Employee data and pay salries")
	fmt.Scan(&paySalaries)

	if paySalaries == "1" {

		fmt.Println("1.employees\n2.managerAttendance\n3.superVisorEmployees")
		fmt.Scan(&selectDestination)
		client, ctx, cancel, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}

		defer close(client, ctx, cancel)

		var filter, option interface{}
		// fmt.Println(" Please Enter Your employeeName")
		// fmt.Scan(&selectDestination)

		// filter = bson.D{{"employeeName", bson.D{{"$eq", }}}}
		filter = bson.D{{}}
		option = bson.D{{"UID", 0}}
		cursor, err := query(client, ctx, "employeeAttendance", selectDestination, filter, option)
		var result []bson.D
		if err := cursor.All(ctx, &result); err != nil {

			panic(err)
		}
		// var doc1 interface{}
		for _, doc := range result {

			fmt.Println("Here List Of EMployees Are shown Below for salaries ", doc[2].Value)

		}
		fmt.Println("Plese enter The Employee Name To Retrview The Employee Attendance")
		fmt.Scan(&employeeNameToRetrivewData)

		filter = bson.D{{"employeeName", bson.D{{"$eq", employeeNameToRetrivewData}}}}
		query(client, ctx, "employeeAttendance", selectDestination, filter, option)

		// if err := cursor.All(ctx, &result); err != nil {

		// 	panic(err1)
		// }
		// //

		for _, doc := range result {

			doc1 = doc[1].Value   // Shift Number
			doc2 = doc[2].Value   //Name of employee
			doc3 = doc[3].Value   //employee Login Time
			doc4 = doc[4].Value   // Applicable For OTs
			doc5 = doc[5].Value   // Late Logins
			doc6 = doc[6].Value   // Logout Time
			doc7 = doc[7].Value   // Acceprt Or Rejects in Early Logouts
			doc8 = doc[8].Value   // Ot in logout Time
			doc9 = doc[9].Value   //Requsets or Resons in Logins
			doc10 = doc[10].Value //Requests or Resons in Logouts
			doc11 = doc[11].Value //Requests For Leave
			doc12 = doc[12].Value // Leave Aprovals

		}
		fmt.Println("The Enter Employee Detailes", "\n The Employee Shift is ", doc1, "The Employee Name is ", doc2)
		fmt.Println("The Employee Login Time is ", doc3, "And Log Out Time Is", doc6)
		fmt.Println("The Employee Late Logins is ", doc5, "The employee OT is ", doc3)
		fmt.Println("The Employee OT In Logout is", doc8)
		fmt.Println("Plese Press 1 To Pay Salries")
		fmt.Scan(&paySalaryForEmployees)
		if paySalaryForEmployees == "1" {

			fmt.Println("Plese  Enter  Employee Name Salary To These Employee")
			fmt.Scan(&salaryToEmployee)

			var document interface{}
			document = bson.D{

				{"employeeName", employeeNameToRetrivewData},
				{"Salary", salaryToEmployee},
			}

			insertOne(client, ctx, "employeeSalaries", selectDestination, document)
		}

	}
	main()

}
