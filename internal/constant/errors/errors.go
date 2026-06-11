package errors

import (
	"errors"
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	StatusCode int
	Type       *errorx.Type
}

var (
	DBNS          = errorx.NewNamespace("DB")
	DBFetchFailed = errorx.NewType(DBNS, "DB_FETCH_FAILED")
	BankNS        = errorx.NewNamespace("BANK")
	NoActiveBanks = errorx.NewType(BankNS, "NO_ACTIVE_BANKS")

	RequestNS      = errorx.NewNamespace("REQUEST")
	InvalidRequest = errorx.NewType(RequestNS, "INVALID_REQUEST")

	ErrBadJSON           = errors.New("bad json")
	ErrInvalidActionType = errors.New("invalid action type")
)

// list of error namespaceis
var (
	databaseError    = errorx.NewNamespace("database error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	invalidInput     = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	resourceNotFound = errorx.NewNamespace("not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)

	unauthorized = errorx.NewNamespace("unauthorized").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	ineligible   = errorx.NewNamespace("ineligible").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	AccessDenied = errorx.RegisterTrait("You are not authorized to perform the action")
	Ineligible   = errorx.RegisterTrait("You are not eligible to perform the action")
	serverError  = errorx.NewNamespace("INTERNAL_SERVIER_ERROR")

	Unauthenticated = errorx.NewNamespace("user authentication failed")
	ProgramError    = errorx.NewNamespace("program error")

	pgtypeJsonbParseError = errorx.NewNamespace("failed to parse message data")

	billnotfound     = errorx.NewNamespace("BILL_NOT_FOUND ")
	billallreadypaid = errorx.NewNamespace("BILL_ALREADY_PAID ")
	cbsapierror      = errorx.NewNamespace("CBS_API_ERROR ")
)

var (
	ErrInvalidPeriod           = errors.New("invalid period")
	ErrTotalBudgetRequired     = errors.New("total budget is required when categories are set")
	ErrCategorySumExceedsTotal = errors.New("sum of category budgets exceeds total budget")
	ErrInvalidCategory         = errors.New("invalid or inactive category")
	ErrDuplicateCategory       = errors.New("duplicate category in request")
	ErrBudgetNotFound          = errors.New("budget not found for this period")

	ErrBillNotFound             = errorx.NewType(billnotfound, "bill not found")
	ErrBillAllreadyPaid         = errorx.NewType(billallreadypaid, "bill allready paid")
	ErrCBSAPI                   = errorx.NewType(cbsapierror, "cbs api error")
	ErrUnAuthorizedAccess       = errorx.NewType(unauthorized, "unauthorized access")
	ErrFailedToParseMessageData = errorx.NewType(pgtypeJsonbParseError, "failed to parse message data")
	ErrUnableTocreate           = errorx.NewType(databaseError, "unable to create")
	ErrDataAlredyExist          = errorx.NewType(databaseError, "data already exist")
	ErrUnableToGet              = errorx.NewType(databaseError, "unable to get")
	ErrInvalidUserInput         = errorx.NewType(invalidInput, "invalid user input")
	ErrInactiveUserStatus       = errorx.NewType(invalidInput, "Inactive user status")
	ErrTripDeviceChange         = errorx.NewType(invalidInput, "user changed device")
	ErrResourceNotFound         = errorx.NewType(resourceNotFound, "resource not found")
	// ErrCbsLinkedOnepulse        = errorx.NewType(accountcbsmismatch, "cbs,linked and onepulse account for exisiting account mismatch")
	ErrAccessError         = errorx.NewType(unauthorized, "Unauthorized", AccessDenied)
	ErrIneligibleError     = errorx.NewType(ineligible, "Ineligible", Ineligible)
	ErrInternalServerError = errorx.NewType(serverError, "internal server error")
)
var Error = []ErrorType{
	{Type: DBFetchFailed, StatusCode: http.StatusInternalServerError},
	{Type: NoActiveBanks, StatusCode: http.StatusNotFound},
	{Type: InvalidRequest, StatusCode: http.StatusBadRequest},
	{
		StatusCode: http.StatusBadRequest,
		Type:       ErrInvalidUserInput,
	},
	{
		StatusCode: http.StatusForbidden,
		Type:       ErrAccessError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		Type:       ErrInternalServerError,
	},
	// {
	// 	StatusCode: http.StatusBadRequest,
	// 	Type:       ErrAuthClient,
	// },

	// {
	// 	StatusCode: http.StatusUnauthorized,
	// 	Type:       ErrInvalidAccessToken,
	// },

	{
		StatusCode: http.StatusUnauthorized,
		Type:       ErrUnAuthorizedAccess,
	},
	{
		StatusCode: http.StatusInternalServerError,
		Type:       ErrUnableToGet,
	}}
