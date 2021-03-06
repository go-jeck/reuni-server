package namespace

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/response"
	"github.com/gorilla/mux"
)

var proc processor

func getProcessor() processor {
	if proc == nil {
		proc = &mainProcessor{repo: initRepository(appcontext.GetHelper())}
	}
	return proc
}

func getFromContext(r *http.Request, key string) string {
	data := r.Context().Value(key)
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}

func CreateNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	var namespaceData namespaceView
	var serviceName = mux.Vars(r)["service_name"]
	var organizationName = mux.Vars(r)["organization_name"]
	err := json.NewDecoder(r.Body).Decode(&namespaceData)
	if err != nil {
		response.ResponseError("CreateNamespace", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	namespaceData.CreatedBy = getFromContext(r, "username")
	reg, _ := regexp.Compile(`^[^.|\s]+$`)
	if !reg.MatchString(namespaceData.Namespace) {
		response.ResponseError("CreateNamespace", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Organization name should not contain '.' or any whitespaces"))
		return
	}
	err = getProcessor().createNewNamespaceProcessor(organizationName, serviceName, &namespaceData)
	if err != nil {
		response.ResponseError("CreateNamespace", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func RetrieveAllNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	var serviceName = mux.Vars(r)["service_name"]
	var organizationName = mux.Vars(r)["organization_name"]

	configsjson, err := getProcessor().retrieveAllNamespaceProcessor(organizationName, serviceName)
	if err != nil {
		response.ResponseError("RetrieveAllNamespace", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(configsjson))
}
