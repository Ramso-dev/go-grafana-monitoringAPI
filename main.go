package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Org name = namespace
type Org struct {
	Name string `json:"name,omitempty"`
}

//User struct needed for grafana request
type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

//RequestData  needed for grafana api requests. Email = Login = Username
type RequestData struct {
	OrgId               string `json:"orgId,omitempty"`
	UserId              string `json:"userId,omitempty"`
	Email               string `json:"email,omitempty"`
	TenantID            string `json:"tenantID,omitempty"`
	AuthToken           string `json:"authToken,omitempty"`
	Role                string `json:"role,omitempty"`
	TenantLabel         string `json:"tenantLabel,omitempty"`
	PanelGauges         bool   `json:"panelGauges,omitempty"`
	PanelCpu            bool   `json:"panelCpu,omitempty"`
	PanelMemory         bool   `json:"panelMemory,omitempty"`
	PanelIOpressure     bool   `json:"panelIOpressure,omitempty"`
	PanelResourcequotas bool   `json:"panelResourcequotas,omitempty"`
}

type AddUserToOrg struct {
	LoginOrEmail string `json:"loginOrEmail,omitempty"`
	Role         string `json:"role,omitempty"`
}

//Test environment
var grafanaBaseURL = "https://grafana-cw-portal-plg.playground.itandtel.at" //"http://localhost:8081"
//var prometheusURL = "https://promtest-cw-portal-plg.playground.itandtel.at"
var prometheusURL = "https://promtest2-cw-portal-plg.playground.itandtel.at"

var apiBaseURL = "https://api.playground.itandtel.at"
var OSCPBaseURL = "https://manage.playground.itandtel.at"
var ServiceAccount = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjdy1wb3J0YWwtcGxnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImFwaS1hZG1pbi10b2tlbi1xMXAyaCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhcGktYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI5YmRjNDRjMC0xODY2LTExZTctYmRiNS0wMDUwNTY4YzU3ODEiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6Y3ctcG9ydGFsLXBsZzphcGktYWRtaW4ifQ.Lxz6OfrGe_7DaDWZfmpDwqgmcV0ON6wTCoyRJJJaH5O9TyHDyC2tAm-XB0CYue20U9ymBNFUE8XN9Gam1JUfYwesxNLhwzdwpgN-ML58_f_g6rL7ZmhWrS6GZVe8ajvlGsXdYTCaTCiT5Dct8wnkI7S8Hq7Vlu7IrctQvWwwYuLvQgpBf8B8-He98t5QmtN8SJJedEZGvQ7aJ_YOu8Ho9K0i4W-KCCX13FYHnxW4s0gtlMMKXg4pFCV1Vm6gK4TxdWiXI6Im13BzGza6XVz-A3OGNvI4Lbk54nY2PjhmJNQ9doagDYByuG2eQZY1PzYuKag_IMXsD7MKTOcUrdy8Pw"

//Production
/*
var OSCPBaseURL = "https://manage.cloud.itandtel.at"
var apiBaseURL = "https://api.cloud.itandtel.at"
var ServiceAccount = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjdy1wb3J0YWwiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiYXBpLWFkbWluLXRva2VuLWNpc2YxIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFwaS1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjY2NmIxNTA2LWYxMmQtMTFlNS04MmZmLTAwNTA1NmE2MmZhNiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpjdy1wb3J0YWw6YXBpLWFkbWluIn0.Y947NTQNFKwY_hmojrNOpX3OnpIYPZmPhSOqzWn-QeYgVWpYbSCzjChFuqIiiOvJG-fFg_3TTAakyLGu_gIUvaviPfZcPwND_aaFbzPVnYtBEEWBP-KF109O7xiJpEQL1a-U1YmAcZZUglNTexMu239VpSTFsiK8yks_Vq63lMSotUWpIMJLd5d2ucd-87qotmTbxqB5n4R5pZWTE0wZEvC_WJmZc-IfzwXY8sr5mbis_Bcbifj7yAxmGQUHvsYOShhIKXE_m0afYmdG32fOTHg9IvKpXYxmaf2na_cLWD_i3cogrjFOQ0VN0FD03ZfFkaDvDAP9XlgOipwHku7Zhg"
*/
//Grafana admin
//var grafanaAdminUsername = "admin"
//var grafanaAdminPassword = "admin"
var grafanaAdminUsername = "udo@cloudwerkstatt.com"
var grafanaAdminPassword = "cwsuperadmin1234"

//Grafana standard user
var grafanaUser User
var grafanaUserPassword = "2QkNlf7C8DG2Qtrj"

func InitEnvironmentVariables() {
	if os.Getenv("PROD") == "true" {
		fmt.Println("Using production environment.V2.09.08.2017")
		ServiceAccount = os.Getenv("SERVICE_ACCOUNT")
		grafanaBaseURL = os.Getenv("GRAFANA_URL")
		apiBaseURL = os.Getenv("PORTAL_API")
		prometheusURL = os.Getenv("PROMETHEUS_URL")
		OSCPBaseURL = os.Getenv("OSCP_API")

		if ServiceAccount == "" {
			fmt.Println("SERVICE_ACCOUNT environment var missing or invalid")

		}

		if grafanaBaseURL == "" {
			fmt.Println("GRAFANA_URL environment var missing or invalid")

		}
		if apiBaseURL == "" {
			fmt.Println("PORTAL_API environment var missing or invalid")

		}
		if prometheusURL == "" {
			fmt.Println("PROMETHEUS_URL environment var missing or invalid")

		}
		if OSCPBaseURL == "" {
			fmt.Println("OSCP_API environment var missing or invalid")

		}

		fmt.Println("SERVICE_ACCOUNT: ", ServiceAccount)
		fmt.Println("GRAFANA_URL: ", grafanaBaseURL)
		fmt.Println("PORTAL_API: ", apiBaseURL)
		fmt.Println("PROMETHEUS_URL: ", prometheusURL)
		fmt.Println("OSCP_API: ", OSCPBaseURL)
	} else {
		fmt.Println("Using testing environment.V2.09.08.2017")
	}
}

func createOrg(w http.ResponseWriter, req *http.Request) {

	var org Org
	_ = json.NewDecoder(req.Body).Decode(&org)
	fmt.Println("Org:>", org)

	jsonStr, err := json.Marshal(org)
	if err != nil {
		fmt.Println(err)
		//return
	}

	req, err = http.NewRequest("POST", grafanaBaseURL+"/api/orgs", bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {

		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

	/*json, err := json.Marshal(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func searchOrgByName(w http.ResponseWriter, req *http.Request) {

	var org Org
	_ = json.NewDecoder(req.Body).Decode(&org)
	fmt.Println("Org:>", org)

	req, err := http.NewRequest("GET", grafanaBaseURL+"/api/orgs/name/"+org.Name, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func searchUsersInOrg(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("searchUsersInOrg requestData:>", requestData)

	req, err := http.NewRequest("GET", grafanaBaseURL+"/api/orgs/"+requestData.OrgId+"/users", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func searchUserByEmail(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	req, err := http.NewRequest("GET", grafanaBaseURL+"/api/users/lookup?loginOrEmail="+requestData.Email, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func deleteUser(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData delete:>", requestData)

	req, err := http.NewRequest("DELETE", grafanaBaseURL+"/api/orgs/"+requestData.OrgId+"/users/"+requestData.UserId, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func deleteGlobalUser(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData delete:>", requestData)

	req, err := http.NewRequest("DELETE", grafanaBaseURL+"/api/admin/users/"+requestData.UserId, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func createUser(w http.ResponseWriter, req *http.Request) {

	var user User

	_ = json.NewDecoder(req.Body).Decode(&user)
	fmt.Println("user:>", user)

	user.Password = grafanaUserPassword

	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	grafanaUser = user

	fmt.Println(grafanaUser)

	req, err = http.NewRequest("POST", grafanaBaseURL+"/api/admin/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	/*json, err := json.Marshal(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	//w.Write(json)

	fmt.Fprintf(w, "%s", string(body))

}

func addUserToOrg(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	var addUserToOrg AddUserToOrg
	addUserToOrg.LoginOrEmail = requestData.Email
	addUserToOrg.Role = "Admin"
	fmt.Println("addUserToOrg:>", addUserToOrg)

	jsonStr, err := json.Marshal(addUserToOrg)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err = http.NewRequest("POST", grafanaBaseURL+"/api/orgs/"+requestData.OrgId+"/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func switchUserToOrg(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	req, err := http.NewRequest("POST", grafanaBaseURL+"/api/user/using/"+requestData.OrgId, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(requestData.Email, grafanaUserPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func setViewerRole(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	//roleJSON := `{"role":"Viewer"}`
	roleJSON := `{"role":"` + requestData.Role + `"}`
	rawIn := json.RawMessage(roleJSON)
	roleBytes, err := rawIn.MarshalJSON()

	req, err = http.NewRequest("PATCH", grafanaBaseURL+"/api/orgs/"+requestData.OrgId+"/users/"+requestData.UserId, bytes.NewBuffer(roleBytes))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func removeFromMainOrg(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	req, err := http.NewRequest("DELETE", grafanaBaseURL+"/api/orgs/1/users/"+requestData.UserId, nil)
	req.SetBasicAuth(grafanaAdminUsername, grafanaAdminPassword)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func createDashboard(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	//dashboardJSON := DashboardJSON1A + DashboardTemplating(requestData.TenantLabel) + requestData.TenantID + DashboardJSON1B

	//ONLY FOR TESTING
	//requestData.TenantLabel = "shared"
	//requestData.TenantID = "cw"

	//dashboardJSON := DashboardJSON1A + DashboardTemplating(requestData.TenantLabel, requestData.TenantID) + DashboardJSON1B

	//In case the tenant has no dedicated nodes, the gauges shouldbt be in the dashboard
	if requestData.TenantLabel == "null" {
		requestData.PanelGauges = false
		requestData.PanelIOpressure = false
	}
	dashboardJSON := DashboardPanels(requestData.PanelGauges, requestData.PanelCpu, requestData.PanelMemory, requestData.PanelIOpressure, requestData.PanelResourcequotas) + DashboardTemplating(requestData.TenantLabel, requestData.TenantID) + DashboardJSON1B

	rawIn := json.RawMessage(dashboardJSON)
	dashboardBytes, err := rawIn.MarshalJSON()

	req, err = http.NewRequest("POST", grafanaBaseURL+"/api/dashboards/db", bytes.NewBuffer(dashboardBytes))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(requestData.Email, grafanaUserPassword)
	fmt.Println(grafanaUser.Login)
	fmt.Println(requestData.Email)
	fmt.Println(grafanaUserPassword)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func createSource(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	sourceJSON := `
{
  "name":"prometheus_source",
  "type":"prometheus",
  "url":"` + prometheusURL + `", 
  "access":"proxy",
  "basicAuth":false
}
`

	rawIn := json.RawMessage(sourceJSON)
	sourceBytes, err := rawIn.MarshalJSON()

	req, err = http.NewRequest("POST", grafanaBaseURL+"/api/datasources", bytes.NewBuffer(sourceBytes))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(requestData.Email, grafanaUserPassword)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", string(body))

}

func getUserTenant(w http.ResponseWriter, req *http.Request) {

	var requestData RequestData
	_ = json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("requestData:>", requestData)

	req, err := http.NewRequest("GET", apiBaseURL+"/myTenants", nil)
	req.Header.Set("Authorization", "Bearer "+requestData.AuthToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("response Body:", string(body))

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

type Tenant struct {
	Name        string       `json:"name,omitempty"`
	ID          int          `json:"id,omitempty"`
	NodeLabel   string       `json:"nodeLabel,omitempty"`
	TenantAdmin TenantUser   `json:"tenantAdmin,omitempty"`
	Monitoring  bool         `json:"monitoring,omitempty"`
	Users       []TenantUser `json:"users,omitempty"`
	//my vars
	TenantAdminsArray []TenantUser `json:"tenantAdminsArray,omitempty"`
	AmIAdmin          bool         `json:"amIAdmin,omitempty"`
}

type TenantUser struct {
	Username string `json:"username,omitempty"`
	//my vars
	AmIAdmin bool `json:"amIAdmin,omitempty"`
}

type OSGroup struct {
	//Kind       string `json:"kind,omitempty"`
	//ApiVersion string `json:"apiVersion,omitempty"`
	Items []struct {
		Metadata struct {
			Name              string            `json:"name,omitempty"`
			UID               string            `json:"uid,omitempty"`
			ResourceVersion   string            `json:"resourceVersion,omitempty"`
			CreationTimestamp string            `json:"creationTimestamp,omitempty"`
			Labels            map[string]string `json:"labels,omitempty"`
		} `json:"metadata,omitempty"`
		Users []string `json:"users,omitempty"`
	} `json:"items,omitempty"`
	Metadata struct {
		Name              string            `json:"name,omitempty"`
		UID               string            `json:"uid,omitempty"`
		ResourceVersion   string            `json:"resourceVersion,omitempty"`
		CreationTimestamp string            `json:"creationTimestamp,omitempty"`
		Labels            map[string]string `json:"labels,omitempty"`
	} `json:"metadata,omitempty"`
	Users []string `json:"users,omitempty"`
}

func getTenantGroup(w http.ResponseWriter, req *http.Request) {

	fmt.Println("$$$ Admin Validation $$$")

	KcTokenBearer := req.Header.Get("Authorization")

	KcToken := strings.TrimPrefix(KcTokenBearer, "Bearer ")

	fmt.Println(KcToken)

	token, err := jwt.Parse(KcToken, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//	}

		return []byte("9R9nzUPL4fqWhrRx"), nil
	})

	/*

		if token.Valid {
			fmt.Println("Provider Token validated")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				http.Error(w, "That's not even a token", 403)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				http.Error(w, "Token is expired", 403)
				return
			}

		} //else {
		//fmt.Println("Couldn't handle this token:", err)
		//	return false
		//}
	*/

	//Find out which user am I from the Provider token's claims
	var user = ""
	var issuerApp = ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if claims["email"] != nil {
			user = claims["email"].(string)
			issuerApp = claims["aud"].(string)
			fmt.Println("- User retrieved from provider token's claim: " + user)
			fmt.Println("- From claims, issuer app is : " + issuerApp)
		} else {
			//fmt.Println("Couldnt retrieve user. No user claim found ")
			//http.Error(w, "Forbidden. Couldnt retrieve user. token claims invalid or not found", 401)
			//return
		}
	} else {
		fmt.Println(err)
		return
	}

	//Find out which is the user's TenantID
	req, err = http.NewRequest("GET", apiBaseURL+"/myTenants", nil)
	req.Header.Set("Authorization", "Bearer "+KcToken)

	fmt.Println("- Looking for user's tenant...")

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		http.Error(w, "Could not retrieve tenant", 403)
		return
	}

	fmt.Println("- response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	tenant := make([]Tenant, 0)
	json.Unmarshal([]byte(string(body)), &tenant)
	fmt.Println("- tenant ->", tenant)

	req, err = http.NewRequest("GET", OSCPBaseURL+"/oapi/v1/groups/"+strconv.Itoa(tenant[0].ID)+"-tenantadmin", nil)
	req.Header.Set("Authorization", "Bearer "+ServiceAccount)

	fmt.Println("- Checking if user is tenantadmin...")

	if err != nil {
		fmt.Println(err)
		return
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("- response Status -> ", resp.StatusCode) //fmt.Println("response Headers:", resp.Header)
	body, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	var osgroup OSGroup
	json.Unmarshal([]byte(string(body)), &osgroup)
	//fmt.Println("osgroup:>", osgroup)

	tenant[0].AmIAdmin = false //initialize, cant be null
	//var foundTenantadmin TenantUser
	var aTenantAdmin TenantUser
	for _, aUser := range osgroup.Users {
		if user == aUser {
			fmt.Println("$$$ Validation Success. User found. User is tenantadmin: ", aUser, " $$$")

			tenant[0].AmIAdmin = true

			//foundTenantadmin.Username = aUser
			//tenant[0].TenantAdmin = foundTenantadmin
		}

		//get all the tenantadmins
		aTenantAdmin.Username = aUser
		aTenantAdmin.AmIAdmin = true
		tenant[0].TenantAdminsArray = append(tenant[0].TenantAdminsArray, aTenantAdmin)

	}

	//TODO : remove tenantadmins from userlist??
	fmt.Println("merging users lists")
	for _, aTenantAdmin := range tenant[0].TenantAdminsArray {
		for i, aUser := range tenant[0].Users {
			if aTenantAdmin.Username == aUser.Username {
				tenant[0].Users = append(tenant[0].Users[:i], tenant[0].Users[i+1:]...)
				tenant[0].Users = append(tenant[0].Users, aTenantAdmin)
				continue
			}
		}
	}

	fmt.Println(tenant[0].Users)

	jsonStr, err := json.Marshal(tenant)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(jsonStr))

	//w.WriteHeader(resp.StatusCode)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		w.WriteHeader(resp.StatusCode)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)

}

//MAIN
func handleView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, todoHTML, "Up and Running")

}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleView)

	r.HandleFunc("/grafana/org", createOrg)
	r.HandleFunc("/grafana/org/search", searchOrgByName)
	r.HandleFunc("/grafana/org/users", searchUsersInOrg)
	r.HandleFunc("/grafana/org/join", addUserToOrg)

	r.HandleFunc("/grafana/user", createUser)
	r.HandleFunc("/grafana/user/switch", switchUserToOrg)
	r.HandleFunc("/grafana/user/restrict", removeFromMainOrg)
	r.HandleFunc("/grafana/user/search", searchUserByEmail)
	r.HandleFunc("/grafana/user/delete", deleteGlobalUser)
	r.HandleFunc("/grafana/user/role", setViewerRole)

	r.HandleFunc("/grafana/dashboard", createDashboard)
	r.HandleFunc("/grafana/source", createSource)

	//r.HandleFunc("/deleteuser", deleteUser)

	r.HandleFunc("/cw/user/tenant", getUserTenant)
	r.HandleFunc("/cw/user/tenantgroup", getTenantGroup)

	return r
}

func main() {

	InitEnvironmentVariables()

	router := InitRoutes()
	http.Handle("/", router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":8082", handler)
}

const todoHTML = `<!DOCTYPE html>
<html>
<head></head>
<body>
<div>
<h1>%s</h1>
</div>
</body>
</html>`
