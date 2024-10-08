// Code generated by go-swagger; DO NOT EDIT.

package agent_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchAgentParams creates a new SearchAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAgentParams() *SearchAgentParams {
	return &SearchAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAgentParamsWithTimeout creates a new SearchAgentParams object
// with the ability to set a timeout on a request.
func NewSearchAgentParamsWithTimeout(timeout time.Duration) *SearchAgentParams {
	return &SearchAgentParams{
		timeout: timeout,
	}
}

// NewSearchAgentParamsWithContext creates a new SearchAgentParams object
// with the ability to set a context for a request.
func NewSearchAgentParamsWithContext(ctx context.Context) *SearchAgentParams {
	return &SearchAgentParams{
		Context: ctx,
	}
}

// NewSearchAgentParamsWithHTTPClient creates a new SearchAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAgentParamsWithHTTPClient(client *http.Client) *SearchAgentParams {
	return &SearchAgentParams{
		HTTPClient: client,
	}
}

/*
SearchAgentParams contains all the parameters to send to the API endpoint

	for the search agent operation.

	Typically these are written to a http.Request.
*/
type SearchAgentParams struct {

	// AllowChannels.
	AllowChannels []string

	// AuditorID.
	AuditorID []int64

	// Extension.
	Extension []string

	// Fields.
	Fields []string

	// ID.
	ID []string

	// IsSupervisor.
	//
	// Format: boolean
	IsSupervisor *bool

	// NotSkillID.
	NotSkillID []int64

	// NotSupervisor.
	//
	// Format: boolean
	NotSupervisor *bool

	// NotTeamID.
	NotTeamID []int64

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// QueueID.
	QueueID []int64

	// RegionID.
	RegionID []int64

	// Size.
	//
	// Format: int32
	Size *int32

	// SkillID.
	SkillID []int64

	// Sort.
	Sort *string

	// SupervisorID.
	SupervisorID []int64

	// TeamID.
	TeamID []int64

	// UserID.
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentParams) WithDefaults() *SearchAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search agent params
func (o *SearchAgentParams) WithTimeout(timeout time.Duration) *SearchAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search agent params
func (o *SearchAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search agent params
func (o *SearchAgentParams) WithContext(ctx context.Context) *SearchAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search agent params
func (o *SearchAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search agent params
func (o *SearchAgentParams) WithHTTPClient(client *http.Client) *SearchAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search agent params
func (o *SearchAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowChannels adds the allowChannels to the search agent params
func (o *SearchAgentParams) WithAllowChannels(allowChannels []string) *SearchAgentParams {
	o.SetAllowChannels(allowChannels)
	return o
}

// SetAllowChannels adds the allowChannels to the search agent params
func (o *SearchAgentParams) SetAllowChannels(allowChannels []string) {
	o.AllowChannels = allowChannels
}

// WithAuditorID adds the auditorID to the search agent params
func (o *SearchAgentParams) WithAuditorID(auditorID []int64) *SearchAgentParams {
	o.SetAuditorID(auditorID)
	return o
}

// SetAuditorID adds the auditorId to the search agent params
func (o *SearchAgentParams) SetAuditorID(auditorID []int64) {
	o.AuditorID = auditorID
}

// WithExtension adds the extension to the search agent params
func (o *SearchAgentParams) WithExtension(extension []string) *SearchAgentParams {
	o.SetExtension(extension)
	return o
}

// SetExtension adds the extension to the search agent params
func (o *SearchAgentParams) SetExtension(extension []string) {
	o.Extension = extension
}

// WithFields adds the fields to the search agent params
func (o *SearchAgentParams) WithFields(fields []string) *SearchAgentParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search agent params
func (o *SearchAgentParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search agent params
func (o *SearchAgentParams) WithID(id []string) *SearchAgentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search agent params
func (o *SearchAgentParams) SetID(id []string) {
	o.ID = id
}

// WithIsSupervisor adds the isSupervisor to the search agent params
func (o *SearchAgentParams) WithIsSupervisor(isSupervisor *bool) *SearchAgentParams {
	o.SetIsSupervisor(isSupervisor)
	return o
}

// SetIsSupervisor adds the isSupervisor to the search agent params
func (o *SearchAgentParams) SetIsSupervisor(isSupervisor *bool) {
	o.IsSupervisor = isSupervisor
}

// WithNotSkillID adds the notSkillID to the search agent params
func (o *SearchAgentParams) WithNotSkillID(notSkillID []int64) *SearchAgentParams {
	o.SetNotSkillID(notSkillID)
	return o
}

// SetNotSkillID adds the notSkillId to the search agent params
func (o *SearchAgentParams) SetNotSkillID(notSkillID []int64) {
	o.NotSkillID = notSkillID
}

// WithNotSupervisor adds the notSupervisor to the search agent params
func (o *SearchAgentParams) WithNotSupervisor(notSupervisor *bool) *SearchAgentParams {
	o.SetNotSupervisor(notSupervisor)
	return o
}

// SetNotSupervisor adds the notSupervisor to the search agent params
func (o *SearchAgentParams) SetNotSupervisor(notSupervisor *bool) {
	o.NotSupervisor = notSupervisor
}

// WithNotTeamID adds the notTeamID to the search agent params
func (o *SearchAgentParams) WithNotTeamID(notTeamID []int64) *SearchAgentParams {
	o.SetNotTeamID(notTeamID)
	return o
}

// SetNotTeamID adds the notTeamId to the search agent params
func (o *SearchAgentParams) SetNotTeamID(notTeamID []int64) {
	o.NotTeamID = notTeamID
}

// WithPage adds the page to the search agent params
func (o *SearchAgentParams) WithPage(page *int32) *SearchAgentParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search agent params
func (o *SearchAgentParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search agent params
func (o *SearchAgentParams) WithQ(q *string) *SearchAgentParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search agent params
func (o *SearchAgentParams) SetQ(q *string) {
	o.Q = q
}

// WithQueueID adds the queueID to the search agent params
func (o *SearchAgentParams) WithQueueID(queueID []int64) *SearchAgentParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the search agent params
func (o *SearchAgentParams) SetQueueID(queueID []int64) {
	o.QueueID = queueID
}

// WithRegionID adds the regionID to the search agent params
func (o *SearchAgentParams) WithRegionID(regionID []int64) *SearchAgentParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the search agent params
func (o *SearchAgentParams) SetRegionID(regionID []int64) {
	o.RegionID = regionID
}

// WithSize adds the size to the search agent params
func (o *SearchAgentParams) WithSize(size *int32) *SearchAgentParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search agent params
func (o *SearchAgentParams) SetSize(size *int32) {
	o.Size = size
}

// WithSkillID adds the skillID to the search agent params
func (o *SearchAgentParams) WithSkillID(skillID []int64) *SearchAgentParams {
	o.SetSkillID(skillID)
	return o
}

// SetSkillID adds the skillId to the search agent params
func (o *SearchAgentParams) SetSkillID(skillID []int64) {
	o.SkillID = skillID
}

// WithSort adds the sort to the search agent params
func (o *SearchAgentParams) WithSort(sort *string) *SearchAgentParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search agent params
func (o *SearchAgentParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithSupervisorID adds the supervisorID to the search agent params
func (o *SearchAgentParams) WithSupervisorID(supervisorID []int64) *SearchAgentParams {
	o.SetSupervisorID(supervisorID)
	return o
}

// SetSupervisorID adds the supervisorId to the search agent params
func (o *SearchAgentParams) SetSupervisorID(supervisorID []int64) {
	o.SupervisorID = supervisorID
}

// WithTeamID adds the teamID to the search agent params
func (o *SearchAgentParams) WithTeamID(teamID []int64) *SearchAgentParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the search agent params
func (o *SearchAgentParams) SetTeamID(teamID []int64) {
	o.TeamID = teamID
}

// WithUserID adds the userID to the search agent params
func (o *SearchAgentParams) WithUserID(userID []string) *SearchAgentParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the search agent params
func (o *SearchAgentParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowChannels != nil {

		// binding items for allow_channels
		joinedAllowChannels := o.bindParamAllowChannels(reg)

		// query array param allow_channels
		if err := r.SetQueryParam("allow_channels", joinedAllowChannels...); err != nil {
			return err
		}
	}

	if o.AuditorID != nil {

		// binding items for auditor_id
		joinedAuditorID := o.bindParamAuditorID(reg)

		// query array param auditor_id
		if err := r.SetQueryParam("auditor_id", joinedAuditorID...); err != nil {
			return err
		}
	}

	if o.Extension != nil {

		// binding items for extension
		joinedExtension := o.bindParamExtension(reg)

		// query array param extension
		if err := r.SetQueryParam("extension", joinedExtension...); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.IsSupervisor != nil {

		// query param is_supervisor
		var qrIsSupervisor bool

		if o.IsSupervisor != nil {
			qrIsSupervisor = *o.IsSupervisor
		}
		qIsSupervisor := swag.FormatBool(qrIsSupervisor)
		if qIsSupervisor != "" {

			if err := r.SetQueryParam("is_supervisor", qIsSupervisor); err != nil {
				return err
			}
		}
	}

	if o.NotSkillID != nil {

		// binding items for not_skill_id
		joinedNotSkillID := o.bindParamNotSkillID(reg)

		// query array param not_skill_id
		if err := r.SetQueryParam("not_skill_id", joinedNotSkillID...); err != nil {
			return err
		}
	}

	if o.NotSupervisor != nil {

		// query param not_supervisor
		var qrNotSupervisor bool

		if o.NotSupervisor != nil {
			qrNotSupervisor = *o.NotSupervisor
		}
		qNotSupervisor := swag.FormatBool(qrNotSupervisor)
		if qNotSupervisor != "" {

			if err := r.SetQueryParam("not_supervisor", qNotSupervisor); err != nil {
				return err
			}
		}
	}

	if o.NotTeamID != nil {

		// binding items for not_team_id
		joinedNotTeamID := o.bindParamNotTeamID(reg)

		// query array param not_team_id
		if err := r.SetQueryParam("not_team_id", joinedNotTeamID...); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.QueueID != nil {

		// binding items for queue_id
		joinedQueueID := o.bindParamQueueID(reg)

		// query array param queue_id
		if err := r.SetQueryParam("queue_id", joinedQueueID...); err != nil {
			return err
		}
	}

	if o.RegionID != nil {

		// binding items for region_id
		joinedRegionID := o.bindParamRegionID(reg)

		// query array param region_id
		if err := r.SetQueryParam("region_id", joinedRegionID...); err != nil {
			return err
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.SkillID != nil {

		// binding items for skill_id
		joinedSkillID := o.bindParamSkillID(reg)

		// query array param skill_id
		if err := r.SetQueryParam("skill_id", joinedSkillID...); err != nil {
			return err
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.SupervisorID != nil {

		// binding items for supervisor_id
		joinedSupervisorID := o.bindParamSupervisorID(reg)

		// query array param supervisor_id
		if err := r.SetQueryParam("supervisor_id", joinedSupervisorID...); err != nil {
			return err
		}
	}

	if o.TeamID != nil {

		// binding items for team_id
		joinedTeamID := o.bindParamTeamID(reg)

		// query array param team_id
		if err := r.SetQueryParam("team_id", joinedTeamID...); err != nil {
			return err
		}
	}

	if o.UserID != nil {

		// binding items for user_id
		joinedUserID := o.bindParamUserID(reg)

		// query array param user_id
		if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchAgent binds the parameter allow_channels
func (o *SearchAgentParams) bindParamAllowChannels(formats strfmt.Registry) []string {
	allowChannelsIR := o.AllowChannels

	var allowChannelsIC []string
	for _, allowChannelsIIR := range allowChannelsIR { // explode []string

		allowChannelsIIV := allowChannelsIIR // string as string
		allowChannelsIC = append(allowChannelsIC, allowChannelsIIV)
	}

	// items.CollectionFormat: "multi"
	allowChannelsIS := swag.JoinByFormat(allowChannelsIC, "multi")

	return allowChannelsIS
}

// bindParamSearchAgent binds the parameter auditor_id
func (o *SearchAgentParams) bindParamAuditorID(formats strfmt.Registry) []string {
	auditorIDIR := o.AuditorID

	var auditorIDIC []string
	for _, auditorIDIIR := range auditorIDIR { // explode []int64

		auditorIDIIV := swag.FormatInt64(auditorIDIIR) // int64 as string
		auditorIDIC = append(auditorIDIC, auditorIDIIV)
	}

	// items.CollectionFormat: "multi"
	auditorIDIS := swag.JoinByFormat(auditorIDIC, "multi")

	return auditorIDIS
}

// bindParamSearchAgent binds the parameter extension
func (o *SearchAgentParams) bindParamExtension(formats strfmt.Registry) []string {
	extensionIR := o.Extension

	var extensionIC []string
	for _, extensionIIR := range extensionIR { // explode []string

		extensionIIV := extensionIIR // string as string
		extensionIC = append(extensionIC, extensionIIV)
	}

	// items.CollectionFormat: "multi"
	extensionIS := swag.JoinByFormat(extensionIC, "multi")

	return extensionIS
}

// bindParamSearchAgent binds the parameter fields
func (o *SearchAgentParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}

// bindParamSearchAgent binds the parameter id
func (o *SearchAgentParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamSearchAgent binds the parameter not_skill_id
func (o *SearchAgentParams) bindParamNotSkillID(formats strfmt.Registry) []string {
	notSkillIDIR := o.NotSkillID

	var notSkillIDIC []string
	for _, notSkillIDIIR := range notSkillIDIR { // explode []int64

		notSkillIDIIV := swag.FormatInt64(notSkillIDIIR) // int64 as string
		notSkillIDIC = append(notSkillIDIC, notSkillIDIIV)
	}

	// items.CollectionFormat: "multi"
	notSkillIDIS := swag.JoinByFormat(notSkillIDIC, "multi")

	return notSkillIDIS
}

// bindParamSearchAgent binds the parameter not_team_id
func (o *SearchAgentParams) bindParamNotTeamID(formats strfmt.Registry) []string {
	notTeamIDIR := o.NotTeamID

	var notTeamIDIC []string
	for _, notTeamIDIIR := range notTeamIDIR { // explode []int64

		notTeamIDIIV := swag.FormatInt64(notTeamIDIIR) // int64 as string
		notTeamIDIC = append(notTeamIDIC, notTeamIDIIV)
	}

	// items.CollectionFormat: "multi"
	notTeamIDIS := swag.JoinByFormat(notTeamIDIC, "multi")

	return notTeamIDIS
}

// bindParamSearchAgent binds the parameter queue_id
func (o *SearchAgentParams) bindParamQueueID(formats strfmt.Registry) []string {
	queueIDIR := o.QueueID

	var queueIDIC []string
	for _, queueIDIIR := range queueIDIR { // explode []int64

		queueIDIIV := swag.FormatInt64(queueIDIIR) // int64 as string
		queueIDIC = append(queueIDIC, queueIDIIV)
	}

	// items.CollectionFormat: "multi"
	queueIDIS := swag.JoinByFormat(queueIDIC, "multi")

	return queueIDIS
}

// bindParamSearchAgent binds the parameter region_id
func (o *SearchAgentParams) bindParamRegionID(formats strfmt.Registry) []string {
	regionIDIR := o.RegionID

	var regionIDIC []string
	for _, regionIDIIR := range regionIDIR { // explode []int64

		regionIDIIV := swag.FormatInt64(regionIDIIR) // int64 as string
		regionIDIC = append(regionIDIC, regionIDIIV)
	}

	// items.CollectionFormat: "multi"
	regionIDIS := swag.JoinByFormat(regionIDIC, "multi")

	return regionIDIS
}

// bindParamSearchAgent binds the parameter skill_id
func (o *SearchAgentParams) bindParamSkillID(formats strfmt.Registry) []string {
	skillIDIR := o.SkillID

	var skillIDIC []string
	for _, skillIDIIR := range skillIDIR { // explode []int64

		skillIDIIV := swag.FormatInt64(skillIDIIR) // int64 as string
		skillIDIC = append(skillIDIC, skillIDIIV)
	}

	// items.CollectionFormat: "multi"
	skillIDIS := swag.JoinByFormat(skillIDIC, "multi")

	return skillIDIS
}

// bindParamSearchAgent binds the parameter supervisor_id
func (o *SearchAgentParams) bindParamSupervisorID(formats strfmt.Registry) []string {
	supervisorIDIR := o.SupervisorID

	var supervisorIDIC []string
	for _, supervisorIDIIR := range supervisorIDIR { // explode []int64

		supervisorIDIIV := swag.FormatInt64(supervisorIDIIR) // int64 as string
		supervisorIDIC = append(supervisorIDIC, supervisorIDIIV)
	}

	// items.CollectionFormat: "multi"
	supervisorIDIS := swag.JoinByFormat(supervisorIDIC, "multi")

	return supervisorIDIS
}

// bindParamSearchAgent binds the parameter team_id
func (o *SearchAgentParams) bindParamTeamID(formats strfmt.Registry) []string {
	teamIDIR := o.TeamID

	var teamIDIC []string
	for _, teamIDIIR := range teamIDIR { // explode []int64

		teamIDIIV := swag.FormatInt64(teamIDIIR) // int64 as string
		teamIDIC = append(teamIDIC, teamIDIIV)
	}

	// items.CollectionFormat: "multi"
	teamIDIS := swag.JoinByFormat(teamIDIC, "multi")

	return teamIDIS
}

// bindParamSearchAgent binds the parameter user_id
func (o *SearchAgentParams) bindParamUserID(formats strfmt.Registry) []string {
	userIDIR := o.UserID

	var userIDIC []string
	for _, userIDIIR := range userIDIR { // explode []string

		userIDIIV := userIDIIR // string as string
		userIDIC = append(userIDIC, userIDIIV)
	}

	// items.CollectionFormat: "multi"
	userIDIS := swag.JoinByFormat(userIDIC, "multi")

	return userIDIS
}
