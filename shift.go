package main

import "fmt"

// office Logout time Timings
var lateLogin2 = "11:21"

func loginAttendanceForShift22(currentTime string) {
	var startingDayTime = "00:00"
	// var endingDayTime = "23:59"

	// Office login Timeings
	var beforeOfficeLoginTime = "10:30"
	var InTimeofficeLogin = "11:00"
	var endTimeofficeLogin = "11:20"
	var end = "01:00"

	
} 

func lo(){



	// Validation For OTs before Logins
	OTvalidationInLogin := currentTime <= beforeOfficeLoginTime && currentTime >= startingDayTime

	// validation For Exact Office Timings
	validationForexactOfficeTime := currentTime >= InTimeofficeLogin && currentTime <= endTimeofficeLogin

	// In between some Time Duration Consider As Normal Logins
	considerAsNormalLogin := currentTime >= considerbeforeOfficeLoginTime && currentTime <= considerbeforeOfficeLoginTime1

	//validation For Late Logins
	LateValidationLoginTime := currentTime >= lateLogin2 && currentTime <= end


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

}   

	

//case7
func logoutAttendanceForShift2(currentTime string) {
	var officeLogoutTime = "19:30"

	var officeLogoutEndTiming = "20:00"

	var ToLateLogoutTime = "20:30"

	// consider Normal Logouts

	var considerbeforeOfficeLogoutTime = "20:01"
	var considerbeforeOfficeLogoutTime1 = "20:29"

	var EarlyLogout = currentTime <= officeLogoutTime && currentTime >= lateLogin2

	// validation For Exact logout Timings
	var ExactLogoutValidation = currentTime >= officeLogoutTime && currentTime <= officeLogoutEndTiming

	// Validation For Late Logout and applay OTs
	var OTvalidationInLogout = currentTime >= ToLateLogoutTime && currentTime <= endingDayTime

	// In between some Time Duration Consider As Normal Logout
	var considerAsNormalLogout = currentTime >= considerbeforeOfficeLogoutTime && currentTime <= considerbeforeOfficeLogoutTime1
	if EarlyLogout {
		fmt.Println("your are early logout Please Let Me Know Reason")
		fmt.Scan(&reasonORApprovalForLogout)
		updateemployeeLogout(currentTime, "In Progress", "0", reasonORApprovalForLogout)

	} else if ExactLogoutValidation {
		fmt.Println("your are Logout  at Office Timings and time is", currentTime)
		updateemployeeLogout(currentTime, "0", "0", "No")

	} else if OTvalidationInLogout {

		fmt.Println("you are logout To Late and LogOut Time is and you are eligibe for OT", currentTime)
		updateemployeeLogout("0", "17:30 To 18:00", currentTime, "RequiredApprovalForOT")

	} else if considerAsNormalLogout {
		fmt.Println("you are normal logout")
		updateemployeeLogout("0", currentTime, "0", "Normal")
	}

}
