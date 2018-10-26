// Package adexchangebuyer2 provides access to the Ad Exchange Buyer API II.
//
// See https://developers.google.com/authorized-buyers/apis/reference/rest/
//
// Usage example:
//
//   import "google.golang.org/api/adexchangebuyer2/v2beta1"
//   ...
//   adexchangebuyer2Service, err := adexchangebuyer2.New(oauthHttpClient)
package adexchangebuyer2 // import "google.golang.org/api/adexchangebuyer2/v2beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "adexchangebuyer2:v2beta1"
const apiName = "adexchangebuyer2"
const apiVersion = "v2beta1"
const basePath = "https://adexchangebuyer.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Manage your Ad Exchange buyer account configuration
	AdexchangeBuyerScope = "https://www.googleapis.com/auth/adexchange.buyer"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)
	s.Bidders = NewBiddersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService

	Bidders *BiddersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.Clients = NewAccountsClientsService(s)
	rs.Creatives = NewAccountsCreativesService(s)
	rs.FinalizedProposals = NewAccountsFinalizedProposalsService(s)
	rs.Products = NewAccountsProductsService(s)
	rs.Proposals = NewAccountsProposalsService(s)
	rs.PublisherProfiles = NewAccountsPublisherProfilesService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Clients *AccountsClientsService

	Creatives *AccountsCreativesService

	FinalizedProposals *AccountsFinalizedProposalsService

	Products *AccountsProductsService

	Proposals *AccountsProposalsService

	PublisherProfiles *AccountsPublisherProfilesService
}

func NewAccountsClientsService(s *Service) *AccountsClientsService {
	rs := &AccountsClientsService{s: s}
	rs.Invitations = NewAccountsClientsInvitationsService(s)
	rs.Users = NewAccountsClientsUsersService(s)
	return rs
}

type AccountsClientsService struct {
	s *Service

	Invitations *AccountsClientsInvitationsService

	Users *AccountsClientsUsersService
}

func NewAccountsClientsInvitationsService(s *Service) *AccountsClientsInvitationsService {
	rs := &AccountsClientsInvitationsService{s: s}
	return rs
}

type AccountsClientsInvitationsService struct {
	s *Service
}

func NewAccountsClientsUsersService(s *Service) *AccountsClientsUsersService {
	rs := &AccountsClientsUsersService{s: s}
	return rs
}

type AccountsClientsUsersService struct {
	s *Service
}

func NewAccountsCreativesService(s *Service) *AccountsCreativesService {
	rs := &AccountsCreativesService{s: s}
	rs.DealAssociations = NewAccountsCreativesDealAssociationsService(s)
	return rs
}

type AccountsCreativesService struct {
	s *Service

	DealAssociations *AccountsCreativesDealAssociationsService
}

func NewAccountsCreativesDealAssociationsService(s *Service) *AccountsCreativesDealAssociationsService {
	rs := &AccountsCreativesDealAssociationsService{s: s}
	return rs
}

type AccountsCreativesDealAssociationsService struct {
	s *Service
}

func NewAccountsFinalizedProposalsService(s *Service) *AccountsFinalizedProposalsService {
	rs := &AccountsFinalizedProposalsService{s: s}
	return rs
}

type AccountsFinalizedProposalsService struct {
	s *Service
}

func NewAccountsProductsService(s *Service) *AccountsProductsService {
	rs := &AccountsProductsService{s: s}
	return rs
}

type AccountsProductsService struct {
	s *Service
}

func NewAccountsProposalsService(s *Service) *AccountsProposalsService {
	rs := &AccountsProposalsService{s: s}
	return rs
}

type AccountsProposalsService struct {
	s *Service
}

func NewAccountsPublisherProfilesService(s *Service) *AccountsPublisherProfilesService {
	rs := &AccountsPublisherProfilesService{s: s}
	return rs
}

type AccountsPublisherProfilesService struct {
	s *Service
}

func NewBiddersService(s *Service) *BiddersService {
	rs := &BiddersService{s: s}
	rs.Accounts = NewBiddersAccountsService(s)
	rs.FilterSets = NewBiddersFilterSetsService(s)
	return rs
}

type BiddersService struct {
	s *Service

	Accounts *BiddersAccountsService

	FilterSets *BiddersFilterSetsService
}

func NewBiddersAccountsService(s *Service) *BiddersAccountsService {
	rs := &BiddersAccountsService{s: s}
	rs.Creatives = NewBiddersAccountsCreativesService(s)
	rs.FilterSets = NewBiddersAccountsFilterSetsService(s)
	return rs
}

type BiddersAccountsService struct {
	s *Service

	Creatives *BiddersAccountsCreativesService

	FilterSets *BiddersAccountsFilterSetsService
}

func NewBiddersAccountsCreativesService(s *Service) *BiddersAccountsCreativesService {
	rs := &BiddersAccountsCreativesService{s: s}
	return rs
}

type BiddersAccountsCreativesService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsService(s *Service) *BiddersAccountsFilterSetsService {
	rs := &BiddersAccountsFilterSetsService{s: s}
	rs.BidMetrics = NewBiddersAccountsFilterSetsBidMetricsService(s)
	rs.BidResponseErrors = NewBiddersAccountsFilterSetsBidResponseErrorsService(s)
	rs.BidResponsesWithoutBids = NewBiddersAccountsFilterSetsBidResponsesWithoutBidsService(s)
	rs.FilteredBidRequests = NewBiddersAccountsFilterSetsFilteredBidRequestsService(s)
	rs.FilteredBids = NewBiddersAccountsFilterSetsFilteredBidsService(s)
	rs.ImpressionMetrics = NewBiddersAccountsFilterSetsImpressionMetricsService(s)
	rs.LosingBids = NewBiddersAccountsFilterSetsLosingBidsService(s)
	rs.NonBillableWinningBids = NewBiddersAccountsFilterSetsNonBillableWinningBidsService(s)
	return rs
}

type BiddersAccountsFilterSetsService struct {
	s *Service

	BidMetrics *BiddersAccountsFilterSetsBidMetricsService

	BidResponseErrors *BiddersAccountsFilterSetsBidResponseErrorsService

	BidResponsesWithoutBids *BiddersAccountsFilterSetsBidResponsesWithoutBidsService

	FilteredBidRequests *BiddersAccountsFilterSetsFilteredBidRequestsService

	FilteredBids *BiddersAccountsFilterSetsFilteredBidsService

	ImpressionMetrics *BiddersAccountsFilterSetsImpressionMetricsService

	LosingBids *BiddersAccountsFilterSetsLosingBidsService

	NonBillableWinningBids *BiddersAccountsFilterSetsNonBillableWinningBidsService
}

func NewBiddersAccountsFilterSetsBidMetricsService(s *Service) *BiddersAccountsFilterSetsBidMetricsService {
	rs := &BiddersAccountsFilterSetsBidMetricsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsBidMetricsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsBidResponseErrorsService(s *Service) *BiddersAccountsFilterSetsBidResponseErrorsService {
	rs := &BiddersAccountsFilterSetsBidResponseErrorsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsBidResponseErrorsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsBidResponsesWithoutBidsService(s *Service) *BiddersAccountsFilterSetsBidResponsesWithoutBidsService {
	rs := &BiddersAccountsFilterSetsBidResponsesWithoutBidsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsBidResponsesWithoutBidsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsFilteredBidRequestsService(s *Service) *BiddersAccountsFilterSetsFilteredBidRequestsService {
	rs := &BiddersAccountsFilterSetsFilteredBidRequestsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsFilteredBidRequestsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsFilteredBidsService(s *Service) *BiddersAccountsFilterSetsFilteredBidsService {
	rs := &BiddersAccountsFilterSetsFilteredBidsService{s: s}
	rs.Creatives = NewBiddersAccountsFilterSetsFilteredBidsCreativesService(s)
	rs.Details = NewBiddersAccountsFilterSetsFilteredBidsDetailsService(s)
	return rs
}

type BiddersAccountsFilterSetsFilteredBidsService struct {
	s *Service

	Creatives *BiddersAccountsFilterSetsFilteredBidsCreativesService

	Details *BiddersAccountsFilterSetsFilteredBidsDetailsService
}

func NewBiddersAccountsFilterSetsFilteredBidsCreativesService(s *Service) *BiddersAccountsFilterSetsFilteredBidsCreativesService {
	rs := &BiddersAccountsFilterSetsFilteredBidsCreativesService{s: s}
	return rs
}

type BiddersAccountsFilterSetsFilteredBidsCreativesService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsFilteredBidsDetailsService(s *Service) *BiddersAccountsFilterSetsFilteredBidsDetailsService {
	rs := &BiddersAccountsFilterSetsFilteredBidsDetailsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsFilteredBidsDetailsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsImpressionMetricsService(s *Service) *BiddersAccountsFilterSetsImpressionMetricsService {
	rs := &BiddersAccountsFilterSetsImpressionMetricsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsImpressionMetricsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsLosingBidsService(s *Service) *BiddersAccountsFilterSetsLosingBidsService {
	rs := &BiddersAccountsFilterSetsLosingBidsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsLosingBidsService struct {
	s *Service
}

func NewBiddersAccountsFilterSetsNonBillableWinningBidsService(s *Service) *BiddersAccountsFilterSetsNonBillableWinningBidsService {
	rs := &BiddersAccountsFilterSetsNonBillableWinningBidsService{s: s}
	return rs
}

type BiddersAccountsFilterSetsNonBillableWinningBidsService struct {
	s *Service
}

func NewBiddersFilterSetsService(s *Service) *BiddersFilterSetsService {
	rs := &BiddersFilterSetsService{s: s}
	rs.BidMetrics = NewBiddersFilterSetsBidMetricsService(s)
	rs.BidResponseErrors = NewBiddersFilterSetsBidResponseErrorsService(s)
	rs.BidResponsesWithoutBids = NewBiddersFilterSetsBidResponsesWithoutBidsService(s)
	rs.FilteredBidRequests = NewBiddersFilterSetsFilteredBidRequestsService(s)
	rs.FilteredBids = NewBiddersFilterSetsFilteredBidsService(s)
	rs.ImpressionMetrics = NewBiddersFilterSetsImpressionMetricsService(s)
	rs.LosingBids = NewBiddersFilterSetsLosingBidsService(s)
	rs.NonBillableWinningBids = NewBiddersFilterSetsNonBillableWinningBidsService(s)
	return rs
}

type BiddersFilterSetsService struct {
	s *Service

	BidMetrics *BiddersFilterSetsBidMetricsService

	BidResponseErrors *BiddersFilterSetsBidResponseErrorsService

	BidResponsesWithoutBids *BiddersFilterSetsBidResponsesWithoutBidsService

	FilteredBidRequests *BiddersFilterSetsFilteredBidRequestsService

	FilteredBids *BiddersFilterSetsFilteredBidsService

	ImpressionMetrics *BiddersFilterSetsImpressionMetricsService

	LosingBids *BiddersFilterSetsLosingBidsService

	NonBillableWinningBids *BiddersFilterSetsNonBillableWinningBidsService
}

func NewBiddersFilterSetsBidMetricsService(s *Service) *BiddersFilterSetsBidMetricsService {
	rs := &BiddersFilterSetsBidMetricsService{s: s}
	return rs
}

type BiddersFilterSetsBidMetricsService struct {
	s *Service
}

func NewBiddersFilterSetsBidResponseErrorsService(s *Service) *BiddersFilterSetsBidResponseErrorsService {
	rs := &BiddersFilterSetsBidResponseErrorsService{s: s}
	return rs
}

type BiddersFilterSetsBidResponseErrorsService struct {
	s *Service
}

func NewBiddersFilterSetsBidResponsesWithoutBidsService(s *Service) *BiddersFilterSetsBidResponsesWithoutBidsService {
	rs := &BiddersFilterSetsBidResponsesWithoutBidsService{s: s}
	return rs
}

type BiddersFilterSetsBidResponsesWithoutBidsService struct {
	s *Service
}

func NewBiddersFilterSetsFilteredBidRequestsService(s *Service) *BiddersFilterSetsFilteredBidRequestsService {
	rs := &BiddersFilterSetsFilteredBidRequestsService{s: s}
	return rs
}

type BiddersFilterSetsFilteredBidRequestsService struct {
	s *Service
}

func NewBiddersFilterSetsFilteredBidsService(s *Service) *BiddersFilterSetsFilteredBidsService {
	rs := &BiddersFilterSetsFilteredBidsService{s: s}
	rs.Creatives = NewBiddersFilterSetsFilteredBidsCreativesService(s)
	rs.Details = NewBiddersFilterSetsFilteredBidsDetailsService(s)
	return rs
}

type BiddersFilterSetsFilteredBidsService struct {
	s *Service

	Creatives *BiddersFilterSetsFilteredBidsCreativesService

	Details *BiddersFilterSetsFilteredBidsDetailsService
}

func NewBiddersFilterSetsFilteredBidsCreativesService(s *Service) *BiddersFilterSetsFilteredBidsCreativesService {
	rs := &BiddersFilterSetsFilteredBidsCreativesService{s: s}
	return rs
}

type BiddersFilterSetsFilteredBidsCreativesService struct {
	s *Service
}

func NewBiddersFilterSetsFilteredBidsDetailsService(s *Service) *BiddersFilterSetsFilteredBidsDetailsService {
	rs := &BiddersFilterSetsFilteredBidsDetailsService{s: s}
	return rs
}

type BiddersFilterSetsFilteredBidsDetailsService struct {
	s *Service
}

func NewBiddersFilterSetsImpressionMetricsService(s *Service) *BiddersFilterSetsImpressionMetricsService {
	rs := &BiddersFilterSetsImpressionMetricsService{s: s}
	return rs
}

type BiddersFilterSetsImpressionMetricsService struct {
	s *Service
}

func NewBiddersFilterSetsLosingBidsService(s *Service) *BiddersFilterSetsLosingBidsService {
	rs := &BiddersFilterSetsLosingBidsService{s: s}
	return rs
}

type BiddersFilterSetsLosingBidsService struct {
	s *Service
}

func NewBiddersFilterSetsNonBillableWinningBidsService(s *Service) *BiddersFilterSetsNonBillableWinningBidsService {
	rs := &BiddersFilterSetsNonBillableWinningBidsService{s: s}
	return rs
}

type BiddersFilterSetsNonBillableWinningBidsService struct {
	s *Service
}

// AbsoluteDateRange: An absolute date range, specified by its start
// date and end date.
// The supported range of dates begins 30 days before today and ends
// today.
// Validity checked upon filter set creation. If a filter set with an
// absolute
// date range is run at a later date more than 30 days after start_date,
// it will
// fail.
type AbsoluteDateRange struct {
	// EndDate: The end date of the range (inclusive).
	// Must be within the 30 days leading up to current date, and must be
	// equal to
	// or after start_date.
	EndDate *Date `json:"endDate,omitempty"`

	// StartDate: The start date of the range (inclusive).
	// Must be within the 30 days leading up to current date, and must be
	// equal to
	// or before end_date.
	StartDate *Date `json:"startDate,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndDate") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndDate") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AbsoluteDateRange) MarshalJSON() ([]byte, error) {
	type NoMethod AbsoluteDateRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AcceptProposalRequest: Request to accept a proposal.
type AcceptProposalRequest struct {
	// ProposalRevision: The last known client revision number of the
	// proposal.
	ProposalRevision int64 `json:"proposalRevision,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ProposalRevision") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ProposalRevision") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AcceptProposalRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AcceptProposalRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AdSize: Represents size of a single ad slot, or a creative.
type AdSize struct {
	// Height: The height of the ad slot in pixels.
	// This field will be present only when size type is `PIXEL`.
	Height int64 `json:"height,omitempty,string"`

	// SizeType: The size type of the ad slot.
	//
	// Possible values:
	//   "SIZE_TYPE_UNSPECIFIED" - A placeholder for an undefined size type.
	//   "PIXEL" - Ad slot with size specified by height and width in
	// pixels.
	//   "INTERSTITIAL" - Special size to describe an interstitial ad slot.
	//   "NATIVE" - Native (mobile) ads rendered by the publisher.
	//   "FLUID" - Fluid size (i.e., responsive size) can be resized
	// automatically with the
	// change of outside environment.
	SizeType string `json:"sizeType,omitempty"`

	// Width: The width of the ad slot in pixels.
	// This field will be present only when size type is `PIXEL`.
	Width int64 `json:"width,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdSize) MarshalJSON() ([]byte, error) {
	type NoMethod AdSize
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddDealAssociationRequest: A request for associating a deal and a
// creative.
type AddDealAssociationRequest struct {
	// Association: The association between a creative and a deal that
	// should be added.
	Association *CreativeDealAssociation `json:"association,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Association") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Association") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddDealAssociationRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AddDealAssociationRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AddNoteRequest: Request message for adding a note to a given
// proposal.
type AddNoteRequest struct {
	// Note: Details of the note to add.
	Note *Note `json:"note,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Note") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Note") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AddNoteRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AddNoteRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AppContext: @OutputOnly The app type the restriction applies to for
// mobile device.
type AppContext struct {
	// AppTypes: The app types this restriction applies to.
	//
	// Possible values:
	//   "NATIVE" - Native app context.
	//   "WEB" - Mobile web app context.
	AppTypes []string `json:"appTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppTypes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppContext) MarshalJSON() ([]byte, error) {
	type NoMethod AppContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuctionContext: @OutputOnly The auction type the restriction applies
// to.
type AuctionContext struct {
	// AuctionTypes: The auction types this restriction applies to.
	//
	// Possible values:
	//   "OPEN_AUCTION" - The restriction applies to open auction.
	//   "DIRECT_DEALS" - The restriction applies to direct deals.
	AuctionTypes []string `json:"auctionTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuctionTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuctionTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AuctionContext) MarshalJSON() ([]byte, error) {
	type NoMethod AuctionContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BidMetricsRow: The set of metrics that are measured in numbers of
// bids, representing how
// many bids with the specified dimension values were considered
// eligible at
// each stage of the bidding funnel;
type BidMetricsRow struct {
	// Bids: The number of bids that Ad Exchange received from the buyer.
	Bids *MetricValue `json:"bids,omitempty"`

	// BidsInAuction: The number of bids that were permitted to compete in
	// the auction.
	BidsInAuction *MetricValue `json:"bidsInAuction,omitempty"`

	// BilledImpressions: The number of bids for which the buyer was billed.
	BilledImpressions *MetricValue `json:"billedImpressions,omitempty"`

	// ImpressionsWon: The number of bids that won an impression.
	ImpressionsWon *MetricValue `json:"impressionsWon,omitempty"`

	// MeasurableImpressions: The number of bids for which the corresponding
	// impression was measurable
	// for viewability (as defined by Active View).
	MeasurableImpressions *MetricValue `json:"measurableImpressions,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// ViewableImpressions: The number of bids for which the corresponding
	// impression was viewable (as
	// defined by Active View).
	ViewableImpressions *MetricValue `json:"viewableImpressions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bids") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bids") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BidMetricsRow) MarshalJSON() ([]byte, error) {
	type NoMethod BidMetricsRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BidResponseWithoutBidsStatusRow: The number of impressions with the
// specified dimension values that were
// considered to have no applicable bids, as described by the specified
// status.
type BidResponseWithoutBidsStatusRow struct {
	// ImpressionCount: The number of impressions for which there was a bid
	// response with the
	// specified status.
	ImpressionCount *MetricValue `json:"impressionCount,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// Status: The status specifying why the bid responses were considered
	// to have no
	// applicable bids.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - A placeholder for an undefined status.
	// This value will never be returned in responses.
	//   "RESPONSES_WITHOUT_BIDS" - The response had no bids.
	//   "RESPONSES_WITHOUT_BIDS_FOR_ACCOUNT" - The response had no bids for
	// the specified account, though it may have
	// included bids on behalf of other accounts.
	//   "RESPONSES_WITHOUT_BIDS_FOR_DEAL" - The response had no bids for
	// the specified deal, though it may have
	// included bids on other deals on behalf of the account to which the
	// deal
	// belongs.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImpressionCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImpressionCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BidResponseWithoutBidsStatusRow) MarshalJSON() ([]byte, error) {
	type NoMethod BidResponseWithoutBidsStatusRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Buyer: Represents a buyer of inventory. Each buyer is identified by a
// unique
// Authorized Buyers account ID.
type Buyer struct {
	// AccountId: Authorized Buyers account ID of the buyer.
	AccountId string `json:"accountId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Buyer) MarshalJSON() ([]byte, error) {
	type NoMethod Buyer
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CalloutStatusRow: The number of impressions with the specified
// dimension values where the
// corresponding bid request or bid response was not successful, as
// described by
// the specified callout status.
type CalloutStatusRow struct {
	// CalloutStatusId: The ID of the callout status.
	// See
	// [callout-status-codes](https://developers.google.com/authorized-buyers
	// /rtb/downloads/callout-status-codes).
	CalloutStatusId int64 `json:"calloutStatusId,omitempty"`

	// ImpressionCount: The number of impressions for which there was a bid
	// request or bid response
	// with the specified callout status.
	ImpressionCount *MetricValue `json:"impressionCount,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CalloutStatusId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CalloutStatusId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CalloutStatusRow) MarshalJSON() ([]byte, error) {
	type NoMethod CalloutStatusRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelNegotiationRequest: Request to cancel an ongoing negotiation.
type CancelNegotiationRequest struct {
}

// Client: A client resource represents a client buyer&mdash;an agency,
// a brand, or an
// advertiser customer of the sponsor buyer. Users associated with the
// client
// buyer have restricted access to the Marketplace and certain other
// sections of
// the Authorized Buyers UI based on the role granted to the client
// buyer. All
// fields are required unless otherwise specified.
type Client struct {
	// ClientAccountId: The globally-unique numerical ID of the client.
	// The value of this field is ignored in create and update operations.
	ClientAccountId int64 `json:"clientAccountId,omitempty,string"`

	// ClientName: Name used to represent this client to publishers.
	// You may have multiple clients that map to the same entity,
	// but for each client the combination of `clientName` and entity
	// must be unique.
	// You can specify this field as empty.
	ClientName string `json:"clientName,omitempty"`

	// EntityId: Numerical identifier of the client entity.
	// The entity can be an advertiser, a brand, or an agency.
	// This identifier is unique among all the entities with the same
	// type.
	//
	// A list of all known advertisers with their identifiers is available
	// in
	// the
	// [advertisers.txt](https://storage.googleapis.com/adx-rtb-dictionar
	// ies/advertisers.txt)
	// file.
	//
	// A list of all known brands with their identifiers is available in
	// the
	// [brands.txt](https://storage.googleapis.com/adx-rtb-dictionaries/b
	// rands.txt)
	// file.
	//
	// A list of all known agencies with their identifiers is available in
	// the
	// [agencies.txt](https://storage.googleapis.com/adx-rtb-dictionaries
	// /agencies.txt)
	// file.
	EntityId int64 `json:"entityId,omitempty,string"`

	// EntityName: The name of the entity. This field is automatically
	// fetched based on
	// the type and ID.
	// The value of this field is ignored in create and update operations.
	EntityName string `json:"entityName,omitempty"`

	// EntityType: The type of the client entity: `ADVERTISER`, `BRAND`, or
	// `AGENCY`.
	//
	// Possible values:
	//   "ENTITY_TYPE_UNSPECIFIED" - A placeholder for an undefined client
	// entity type. Should not be used.
	//   "ADVERTISER" - An advertiser.
	//   "BRAND" - A brand.
	//   "AGENCY" - An advertising agency.
	EntityType string `json:"entityType,omitempty"`

	// PartnerClientId: Optional arbitrary unique identifier of this client
	// buyer from the
	// standpoint of its Ad Exchange sponsor buyer.
	//
	// This field can be used to associate a client buyer with the
	// identifier
	// in the namespace of its sponsor buyer, lookup client buyers by
	// that
	// identifier and verify whether an Ad Exchange counterpart of a given
	// client
	// buyer already exists.
	//
	// If present, must be unique among all the client buyers for its
	// Ad Exchange sponsor buyer.
	PartnerClientId string `json:"partnerClientId,omitempty"`

	// Role: The role which is assigned to the client buyer. Each role
	// implies a set of
	// permissions granted to the client. Must be one of
	// `CLIENT_DEAL_VIEWER`,
	// `CLIENT_DEAL_NEGOTIATOR` or `CLIENT_DEAL_APPROVER`.
	//
	// Possible values:
	//   "CLIENT_ROLE_UNSPECIFIED" - A placeholder for an undefined client
	// role.
	//   "CLIENT_DEAL_VIEWER" - Users associated with this client can see
	// publisher deal offers
	// in the Marketplace.
	// They can neither negotiate proposals nor approve deals.
	// If this client is visible to publishers, they can send deal
	// proposals
	// to this client.
	//   "CLIENT_DEAL_NEGOTIATOR" - Users associated with this client can
	// respond to deal proposals
	// sent to them by publishers. They can also initiate deal proposals
	// of their own.
	//   "CLIENT_DEAL_APPROVER" - Users associated with this client can
	// approve eligible deals
	// on your behalf. Some deals may still explicitly require
	// publisher
	// finalization. If this role is not selected, the sponsor buyer
	// will need to manually approve each of their deals.
	Role string `json:"role,omitempty"`

	// Status: The status of the client buyer.
	//
	// Possible values:
	//   "CLIENT_STATUS_UNSPECIFIED" - A placeholder for an undefined client
	// status.
	//   "DISABLED" - A client that is currently disabled.
	//   "ACTIVE" - A client that is currently active.
	Status string `json:"status,omitempty"`

	// VisibleToSeller: Whether the client buyer will be visible to sellers.
	VisibleToSeller bool `json:"visibleToSeller,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientAccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientAccountId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Client) MarshalJSON() ([]byte, error) {
	type NoMethod Client
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClientUser: A client user is created under a client buyer and has
// restricted access to
// the Marketplace and certain other sections of the Authorized Buyers
// UI based
// on the role granted to the associated client buyer.
//
// The only way a new client user can be created is via accepting
// an
// email invitation
// (see the
// accounts.clients.invitations.create
// method).
//
// All fields are required unless otherwise specified.
type ClientUser struct {
	// ClientAccountId: Numerical account ID of the client buyer
	// with which the user is associated; the
	// buyer must be a client of the current sponsor buyer.
	// The value of this field is ignored in an update operation.
	ClientAccountId int64 `json:"clientAccountId,omitempty,string"`

	// Email: User's email address. The value of this field
	// is ignored in an update operation.
	Email string `json:"email,omitempty"`

	// Status: The status of the client user.
	//
	// Possible values:
	//   "USER_STATUS_UNSPECIFIED" - A placeholder for an undefined user
	// status.
	//   "PENDING" - A user who was already created but hasn't accepted the
	// invitation yet.
	//   "ACTIVE" - A user that is currently active.
	//   "DISABLED" - A user that is currently disabled.
	Status string `json:"status,omitempty"`

	// UserId: The unique numerical ID of the client user
	// that has accepted an invitation.
	// The value of this field is ignored in an update operation.
	UserId int64 `json:"userId,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientAccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientAccountId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClientUser) MarshalJSON() ([]byte, error) {
	type NoMethod ClientUser
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClientUserInvitation: An invitation for a new client user to get
// access to the Authorized Buyers
// UI. All fields are required unless otherwise specified.
type ClientUserInvitation struct {
	// ClientAccountId: Numerical account ID of the client buyer
	// that the invited user is associated with.
	// The value of this field is ignored in create operations.
	ClientAccountId int64 `json:"clientAccountId,omitempty,string"`

	// Email: The email address to which the invitation is sent.
	// Email
	// addresses should be unique among all client users under each
	// sponsor
	// buyer.
	Email string `json:"email,omitempty"`

	// InvitationId: The unique numerical ID of the invitation that is sent
	// to the user.
	// The value of this field is ignored in create operations.
	InvitationId int64 `json:"invitationId,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientAccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientAccountId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClientUserInvitation) MarshalJSON() ([]byte, error) {
	type NoMethod ClientUserInvitation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompleteSetupRequest: Request message for indicating that the
// proposal's setup step is complete.
type CompleteSetupRequest struct {
}

// ContactInformation: Contains information on how a buyer or seller can
// be reached.
type ContactInformation struct {
	// Email: Email address for the contact.
	Email string `json:"email,omitempty"`

	// Name: The name of the contact.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContactInformation) MarshalJSON() ([]byte, error) {
	type NoMethod ContactInformation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Correction: @OutputOnly Shows any corrections that were applied to
// this creative.
type Correction struct {
	// Contexts: The contexts for the correction.
	Contexts []*ServingContext `json:"contexts,omitempty"`

	// Details: Additional details about what was corrected.
	Details []string `json:"details,omitempty"`

	// Type: The type of correction that was applied to the creative.
	//
	// Possible values:
	//   "CORRECTION_TYPE_UNSPECIFIED" - The correction type is unknown.
	// Refer to the details for more information.
	//   "VENDOR_IDS_ADDED" - The ad's declared vendors did not match the
	// vendors that were detected.
	// The detected vendors were added.
	//   "SSL_ATTRIBUTE_REMOVED" - The ad had the SSL attribute declared but
	// was not SSL-compliant.
	// The SSL attribute was removed.
	//   "FLASH_FREE_ATTRIBUTE_REMOVED" - The ad was declared as Flash-free
	// but contained Flash, so the Flash-free
	// attribute was removed.
	//   "FLASH_FREE_ATTRIBUTE_ADDED" - The ad was not declared as
	// Flash-free but it did not reference any flash
	// content, so the Flash-free attribute was added.
	//   "REQUIRED_ATTRIBUTE_ADDED" - The ad did not declare a required
	// creative attribute.
	// The attribute was added.
	//   "REQUIRED_VENDOR_ADDED" - The ad did not declare a required
	// technology vendor.
	// The technology vendor was added.
	//   "SSL_ATTRIBUTE_ADDED" - The ad did not declare the SSL attribute
	// but was SSL-compliant, so the
	// SSL attribute was added.
	//   "IN_BANNER_VIDEO_ATTRIBUTE_ADDED" - Properties consistent with
	// In-banner video were found, so an
	// In-Banner Video attribute was added.
	//   "MRAID_ATTRIBUTE_ADDED" - The ad makes calls to the MRAID API so
	// the MRAID attribute was added.
	//   "FLASH_ATTRIBUTE_REMOVED" - The ad unnecessarily declared the Flash
	// attribute, so the Flash attribute
	// was removed.
	//   "VIDEO_IN_SNIPPET_ATTRIBUTE_ADDED" - The ad contains video content.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Correction) MarshalJSON() ([]byte, error) {
	type NoMethod Correction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Creative: A creative and its classification data.
//
// Next ID: 38
type Creative struct {
	// AccountId: The account that this creative belongs to.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	AccountId string `json:"accountId,omitempty"`

	// AdChoicesDestinationUrl: The link to AdChoices destination page.
	AdChoicesDestinationUrl string `json:"adChoicesDestinationUrl,omitempty"`

	// AdvertiserName: The name of the company being advertised in the
	// creative.
	AdvertiserName string `json:"advertiserName,omitempty"`

	// AgencyId: The agency ID for this creative.
	AgencyId int64 `json:"agencyId,omitempty,string"`

	// ApiUpdateTime: @OutputOnly The last update timestamp of the creative
	// via API.
	ApiUpdateTime string `json:"apiUpdateTime,omitempty"`

	// Attributes: All attributes for the ads that may be shown from this
	// creative.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	//
	// Possible values:
	//   "ATTRIBUTE_UNSPECIFIED" - Do not use. This is a placeholder value
	// only.
	//   "IS_TAGGED" - The creative is tagged.
	//   "IS_COOKIE_TARGETED" - The creative is cookie targeted.
	//   "IS_USER_INTEREST_TARGETED" - The creative is user interest
	// targeted.
	//   "EXPANDING_DIRECTION_NONE" - The creative does not expand.
	//   "EXPANDING_DIRECTION_UP" - The creative expands up.
	//   "EXPANDING_DIRECTION_DOWN" - The creative expands down.
	//   "EXPANDING_DIRECTION_LEFT" - The creative expands left.
	//   "EXPANDING_DIRECTION_RIGHT" - The creative expands right.
	//   "EXPANDING_DIRECTION_UP_LEFT" - The creative expands up and left.
	//   "EXPANDING_DIRECTION_UP_RIGHT" - The creative expands up and right.
	//   "EXPANDING_DIRECTION_DOWN_LEFT" - The creative expands down and
	// left.
	//   "EXPANDING_DIRECTION_DOWN_RIGHT" - The creative expands down and
	// right.
	//   "EXPANDING_DIRECTION_UP_OR_DOWN" - The creative expands up or down.
	//   "EXPANDING_DIRECTION_LEFT_OR_RIGHT" - The creative expands left or
	// right.
	//   "EXPANDING_DIRECTION_ANY_DIAGONAL" - The creative expands on any
	// diagonal.
	//   "EXPANDING_ACTION_ROLLOVER_TO_EXPAND" - The creative expands when
	// rolled over.
	//   "INSTREAM_VAST_VIDEO_TYPE_VPAID_FLASH" - The instream vast video
	// type is vpaid flash.
	//   "RICH_MEDIA_CAPABILITY_TYPE_MRAID" - The creative is MRAID
	//   "RICH_MEDIA_CAPABILITY_TYPE_SSL" - The creative is SSL.
	//   "RICH_MEDIA_CAPABILITY_TYPE_INTERSTITIAL" - The creative is an
	// interstitial.
	//   "NATIVE_ELIGIBILITY_ELIGIBLE" - The creative is eligible for
	// native.
	//   "NATIVE_ELIGIBILITY_NOT_ELIGIBLE" - The creative is not eligible
	// for native.
	//   "RENDERING_SIZELESS_ADX" - The creative can dynamically resize to
	// fill a variety of slot sizes.
	Attributes []string `json:"attributes,omitempty"`

	// ClickThroughUrls: The set of destination URLs for the creative.
	ClickThroughUrls []string `json:"clickThroughUrls,omitempty"`

	// Corrections: @OutputOnly Shows any corrections that were applied to
	// this creative.
	Corrections []*Correction `json:"corrections,omitempty"`

	// CreativeId: The buyer-defined creative ID of this creative.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	CreativeId string `json:"creativeId,omitempty"`

	// DealsStatus: @OutputOnly The top-level deals status of this
	// creative.
	// If disapproved, an entry for 'auctionType=DIRECT_DEALS' (or 'ALL')
	// in
	// serving_restrictions will also exist. Note
	// that this may be nuanced with other contextual restrictions, in which
	// case,
	// it may be preferable to read from serving_restrictions directly.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - The status is unknown.
	//   "NOT_CHECKED" - The creative has not been checked.
	//   "CONDITIONALLY_APPROVED" - The creative has been conditionally
	// approved.
	// See serving_restrictions for details.
	//   "APPROVED" - The creative has been approved.
	//   "DISAPPROVED" - The creative has been disapproved.
	DealsStatus string `json:"dealsStatus,omitempty"`

	// DeclaredClickThroughUrls: The set of declared destination URLs for
	// the creative.
	DeclaredClickThroughUrls []string `json:"declaredClickThroughUrls,omitempty"`

	// DetectedAdvertiserIds: @OutputOnly Detected advertiser IDs, if any.
	DetectedAdvertiserIds googleapi.Int64s `json:"detectedAdvertiserIds,omitempty"`

	// DetectedDomains: @OutputOnly
	// The detected domains for this creative.
	DetectedDomains []string `json:"detectedDomains,omitempty"`

	// DetectedLanguages: @OutputOnly
	// The detected languages for this creative. The order is arbitrary. The
	// codes
	// are 2 or 5 characters and are documented
	// at
	// https://developers.google.com/adwords/api/docs/appendix/languagecod
	// es.
	DetectedLanguages []string `json:"detectedLanguages,omitempty"`

	// DetectedProductCategories: @OutputOnly Detected product categories,
	// if any.
	// See the ad-product-categories.txt file in the technical
	// documentation
	// for a list of IDs.
	DetectedProductCategories []int64 `json:"detectedProductCategories,omitempty"`

	// DetectedSensitiveCategories: @OutputOnly Detected sensitive
	// categories, if any.
	// See the ad-sensitive-categories.txt file in the technical
	// documentation for
	// a list of IDs. You should use these IDs along with
	// the
	// excluded-sensitive-category field in the bid request to filter your
	// bids.
	DetectedSensitiveCategories []int64 `json:"detectedSensitiveCategories,omitempty"`

	// FilteringStats: @OutputOnly The filtering stats for this creative.
	FilteringStats *FilteringStats `json:"filteringStats,omitempty"`

	// Html: An HTML creative.
	Html *HtmlContent `json:"html,omitempty"`

	// ImpressionTrackingUrls: The set of URLs to be called to record an
	// impression.
	ImpressionTrackingUrls []string `json:"impressionTrackingUrls,omitempty"`

	// Native: A native creative.
	Native *NativeContent `json:"native,omitempty"`

	// OpenAuctionStatus: @OutputOnly The top-level open auction status of
	// this creative.
	// If disapproved, an entry for 'auctionType = OPEN_AUCTION' (or 'ALL')
	// in
	// serving_restrictions will also exist. Note
	// that this may be nuanced with other contextual restrictions, in which
	// case,
	// it may be preferable to read from serving_restrictions directly.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - The status is unknown.
	//   "NOT_CHECKED" - The creative has not been checked.
	//   "CONDITIONALLY_APPROVED" - The creative has been conditionally
	// approved.
	// See serving_restrictions for details.
	//   "APPROVED" - The creative has been approved.
	//   "DISAPPROVED" - The creative has been disapproved.
	OpenAuctionStatus string `json:"openAuctionStatus,omitempty"`

	// RestrictedCategories: All restricted categories for the ads that may
	// be shown from this creative.
	//
	// Possible values:
	//   "NO_RESTRICTED_CATEGORIES" - The ad has no restricted categories
	//   "ALCOHOL" - The alcohol restricted category.
	RestrictedCategories []string `json:"restrictedCategories,omitempty"`

	// ServingRestrictions: @OutputOnly The granular status of this ad in
	// specific contexts.
	// A context here relates to where something ultimately serves (for
	// example,
	// a physical location, a platform, an HTTPS vs HTTP request, or the
	// type
	// of auction).
	ServingRestrictions []*ServingRestriction `json:"servingRestrictions,omitempty"`

	// VendorIds: All vendor IDs for the ads that may be shown from this
	// creative.
	// See
	// https://storage.googleapis.com/adx-rtb-dictionaries/vendors.txt
	// for possible values.
	VendorIds []int64 `json:"vendorIds,omitempty"`

	// Version: @OutputOnly The version of this creative.
	Version int64 `json:"version,omitempty"`

	// Video: A video creative.
	Video *VideoContent `json:"video,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Creative) MarshalJSON() ([]byte, error) {
	type NoMethod Creative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreativeDealAssociation: The association between a creative and a
// deal.
type CreativeDealAssociation struct {
	// AccountId: The account the creative belongs to.
	AccountId string `json:"accountId,omitempty"`

	// CreativeId: The ID of the creative associated with the deal.
	CreativeId string `json:"creativeId,omitempty"`

	// DealsId: The externalDealId for the deal associated with the
	// creative.
	DealsId string `json:"dealsId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreativeDealAssociation) MarshalJSON() ([]byte, error) {
	type NoMethod CreativeDealAssociation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreativeRestrictions: Represents creative restrictions associated to
// Programmatic Guaranteed/
// Preferred Deal in Ad Manager.
// This doesn't apply to Private Auction and AdX Preferred Deals.
type CreativeRestrictions struct {
	// CreativeFormat: The format of the environment that the creatives will
	// be displayed in.
	//
	// Possible values:
	//   "CREATIVE_FORMAT_UNSPECIFIED" - A placeholder for an undefined
	// creative format.
	//   "DISPLAY" - A creative that will be displayed in environments such
	// as a browser.
	//   "VIDEO" - A video creative that will be displayed in environments
	// such as a video
	// player.
	CreativeFormat string `json:"creativeFormat,omitempty"`

	CreativeSpecifications []*CreativeSpecification `json:"creativeSpecifications,omitempty"`

	// SkippableAdType: Skippable video ads allow viewers to skip ads after
	// 5 seconds.
	//
	// Possible values:
	//   "SKIPPABLE_AD_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// skippable ad type.
	//   "SKIPPABLE" - This video ad can be skipped after 5 seconds.
	//   "INSTREAM_SELECT" - This video ad can be skipped after 5 seconds,
	// and is counted as
	// engaged view after 30 seconds. The creative is hosted on
	// YouTube only, and viewcount of the YouTube video increments
	// after the engaged view.
	//   "NOT_SKIPPABLE" - This video ad is not skippable.
	SkippableAdType string `json:"skippableAdType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreativeFormat") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeFormat") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CreativeRestrictions) MarshalJSON() ([]byte, error) {
	type NoMethod CreativeRestrictions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreativeSize: Specifies the size of the creative.
type CreativeSize struct {
	// AllowedFormats: What formats are allowed by the publisher.
	// If this repeated field is empty then all formats are allowed.
	// For example, if this field contains AllowedFormatType.AUDIO then
	// the
	// publisher only allows an audio ad (without any video).
	//
	// Possible values:
	//   "UNKNOWN" - A placeholder for an undefined allowed format.
	//   "AUDIO" - An audio-only ad (without any video).
	AllowedFormats []string `json:"allowedFormats,omitempty"`

	// CompanionSizes: For video creatives specifies the sizes of companion
	// ads (if present).
	// Companion sizes may be filled in only when creative_size_type = VIDEO
	CompanionSizes []*Size `json:"companionSizes,omitempty"`

	// CreativeSizeType: The creative size type.
	//
	// Possible values:
	//   "CREATIVE_SIZE_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// creative size type.
	//   "REGULAR" - The creative is a regular desktop creative.
	//   "INTERSTITIAL" - The creative is an interstitial creative.
	//   "VIDEO" - The creative is a video creative.
	//   "NATIVE" - The creative is a native (mobile) creative.
	CreativeSizeType string `json:"creativeSizeType,omitempty"`

	// NativeTemplate: The native template for this creative. It will have a
	// value only if
	// creative_size_type = CreativeSizeType.NATIVE.
	// @OutputOnly
	//
	// Possible values:
	//   "UNKNOWN_NATIVE_TEMPLATE" - A placeholder for an undefined native
	// template.
	//   "NATIVE_CONTENT_AD" - The creative is linked to native content ad.
	//   "NATIVE_APP_INSTALL_AD" - The creative is linked to native app
	// install ad.
	//   "NATIVE_VIDEO_CONTENT_AD" - The creative is linked to native video
	// content ad.
	//   "NATIVE_VIDEO_APP_INSTALL_AD" - The creative is linked to native
	// video app install ad.
	NativeTemplate string `json:"nativeTemplate,omitempty"`

	// Size: For regular or video creative size type, specifies the size
	// of the creative
	Size *Size `json:"size,omitempty"`

	// SkippableAdType: The type of skippable ad for this creative. It will
	// have a value only if
	// creative_size_type = CreativeSizeType.VIDEO.
	//
	// Possible values:
	//   "SKIPPABLE_AD_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// skippable ad type.
	//   "GENERIC" - This video ad can be skipped after 5 seconds.
	//   "INSTREAM_SELECT" - This video ad can be skipped after 5 seconds,
	// and count as
	// engaged view after 30 seconds. The creative is hosted on
	// YouTube only, and viewcount of the YouTube video increments
	// after the engaged view.
	//   "NOT_SKIPPABLE" - This video ad is not skippable.
	SkippableAdType string `json:"skippableAdType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowedFormats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllowedFormats") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CreativeSize) MarshalJSON() ([]byte, error) {
	type NoMethod CreativeSize
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreativeSpecification: Represents information for a creative that is
// associated with a Programmatic
// Guaranteed/Preferred Deal in Ad Manager.
type CreativeSpecification struct {
	// CreativeCompanionSizes: Companion sizes may be filled in only when
	// this is a video creative.
	CreativeCompanionSizes []*AdSize `json:"creativeCompanionSizes,omitempty"`

	// CreativeSize: The size of the creative.
	CreativeSize *AdSize `json:"creativeSize,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CreativeCompanionSizes") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeCompanionSizes")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CreativeSpecification) MarshalJSON() ([]byte, error) {
	type NoMethod CreativeSpecification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreativeStatusRow: The number of bids with the specified dimension
// values that did not win the
// auction (either were filtered pre-auction or lost the auction), as
// described
// by the specified creative status.
type CreativeStatusRow struct {
	// BidCount: The number of bids with the specified status.
	BidCount *MetricValue `json:"bidCount,omitempty"`

	// CreativeStatusId: The ID of the creative status.
	// See
	// [creative-status-codes](https://developers.google.com/authorized-buyer
	// s/rtb/downloads/creative-status-codes).
	CreativeStatusId int64 `json:"creativeStatusId,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BidCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BidCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreativeStatusRow) MarshalJSON() ([]byte, error) {
	type NoMethod CreativeStatusRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CriteriaTargeting: Generic targeting used for targeting dimensions
// that contains a list of
// included and excluded numeric IDs.
type CriteriaTargeting struct {
	// ExcludedCriteriaIds: A list of numeric IDs to be excluded.
	ExcludedCriteriaIds googleapi.Int64s `json:"excludedCriteriaIds,omitempty"`

	// TargetedCriteriaIds: A list of numeric IDs to be included.
	TargetedCriteriaIds googleapi.Int64s `json:"targetedCriteriaIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExcludedCriteriaIds")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludedCriteriaIds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CriteriaTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod CriteriaTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Date: Represents a whole or partial calendar date, e.g. a birthday.
// The time of day
// and time zone are either specified elsewhere or are not significant.
// The date
// is relative to the Proleptic Gregorian Calendar. This can
// represent:
//
// * A full date, with non-zero year, month and day values
// * A month and day value, with a zero year, e.g. an anniversary
// * A year on its own, with zero month and day values
// * A year and month value, with a zero day, e.g. a credit card
// expiration date
//
// Related types are google.type.TimeOfDay and
// `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year by itself or a year and month where the day is
	// not
	// significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12, or 0 if specifying a year
	// without a
	// month and day.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DayPart: Daypart targeting message that specifies if the ad can be
// shown
// only during certain parts of a day/week.
type DayPart struct {
	// DayOfWeek: The day of the week to target. If unspecified, applicable
	// to all days.
	//
	// Possible values:
	//   "DAY_OF_WEEK_UNSPECIFIED" - A placeholder for when the day of the
	// week is not specified.
	//   "MONDAY" - Monday
	//   "TUESDAY" - Tuesday
	//   "WEDNESDAY" - Wednesday
	//   "THURSDAY" - Thursday
	//   "FRIDAY" - Friday
	//   "SATURDAY" - Saturday
	//   "SUNDAY" - Sunday
	DayOfWeek string `json:"dayOfWeek,omitempty"`

	// EndTime: The ending time of the day for the ad to show (minute level
	// granularity).
	// The end time is exclusive.
	// This field is not available for filtering in PQL queries.
	EndTime *TimeOfDay `json:"endTime,omitempty"`

	// StartTime: The starting time of day for the ad to show (minute level
	// granularity).
	// The start time is inclusive.
	// This field is not available for filtering in PQL queries.
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DayOfWeek") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DayOfWeek") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DayPart) MarshalJSON() ([]byte, error) {
	type NoMethod DayPart
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DayPartTargeting: Specifies the day part targeting criteria.
type DayPartTargeting struct {
	// DayParts: A list of day part targeting criterion.
	DayParts []*DayPart `json:"dayParts,omitempty"`

	// TimeZoneType: The timezone to use for interpreting the day part
	// targeting.
	//
	// Possible values:
	//   "TIME_ZONE_SOURCE_UNSPECIFIED" - A placeholder for an undefined
	// time zone source.
	//   "PUBLISHER" - Use publisher's time zone setting.
	//   "USER" - Use the user's time zone setting.
	TimeZoneType string `json:"timeZoneType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DayParts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DayParts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DayPartTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod DayPartTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Deal: A deal represents a segment of inventory for displaying ads
// on.
// A proposal can contain multiple deals. A deal contains the terms and
// targeting information that
// is used for serving.
type Deal struct {
	// AvailableEndTime: Proposed flight end time of the deal.
	// This will generally be stored in a granularity of a second.
	// A value is not required for Private Auction deals or Preferred Deals.
	AvailableEndTime string `json:"availableEndTime,omitempty"`

	// AvailableStartTime: Optional proposed flight start time of the
	// deal.
	// This will generally be stored in the granularity of one second since
	// deal serving
	// starts at seconds boundary. Any time specified with more
	// granularity
	// (e.g., in milliseconds) will be truncated towards the start of time
	// in seconds.
	AvailableStartTime string `json:"availableStartTime,omitempty"`

	// BuyerPrivateData: Buyer private data (hidden from seller).
	BuyerPrivateData *PrivateData `json:"buyerPrivateData,omitempty"`

	// CreateProductId: The product ID from which this deal was
	// created.
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	CreateProductId string `json:"createProductId,omitempty"`

	// CreateProductRevision: Optional revision number of the product that
	// the deal was created from.
	// If present on create, and the server `product_revision` has advanced
	// sinced
	// the passed-in `create_product_revision`, an `ABORTED` error will
	// be
	// returned.
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	CreateProductRevision int64 `json:"createProductRevision,omitempty,string"`

	// CreateTime: The time of the deal creation.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// CreativePreApprovalPolicy: Specifies the creative pre-approval
	// policy.
	// @OutputOnly
	//
	// Possible values:
	//   "CREATIVE_PRE_APPROVAL_POLICY_UNSPECIFIED" - A placeholder for an
	// undefined creative pre-approval policy.
	//   "SELLER_PRE_APPROVAL_REQUIRED" - The seller needs to approve each
	// creative before it can serve.
	//   "SELLER_PRE_APPROVAL_NOT_REQUIRED" - The seller does not need to
	// approve each creative before it can serve.
	CreativePreApprovalPolicy string `json:"creativePreApprovalPolicy,omitempty"`

	// CreativeRestrictions: Restricitions about the creatives associated
	// with the deal (i.e., size)
	// This is available for Programmatic Guaranteed/Preferred Deals in Ad
	// Manager.
	// @OutputOnly
	CreativeRestrictions *CreativeRestrictions `json:"creativeRestrictions,omitempty"`

	// CreativeSafeFrameCompatibility: Specifies whether the creative is
	// safeFrame compatible.
	// @OutputOnly
	//
	// Possible values:
	//   "CREATIVE_SAFE_FRAME_COMPATIBILITY_UNSPECIFIED" - A placeholder for
	// an undefined creative safe-frame compatibility.
	//   "COMPATIBLE" - The creatives need to be compatible with the safe
	// frame option.
	//   "INCOMPATIBLE" - The creatives can be incompatible with the safe
	// frame option.
	CreativeSafeFrameCompatibility string `json:"creativeSafeFrameCompatibility,omitempty"`

	// DealId: A unique deal ID for the deal (server-assigned).
	// @OutputOnly
	DealId string `json:"dealId,omitempty"`

	// DealServingMetadata: Metadata about the serving status of this
	// deal.
	// @OutputOnly
	DealServingMetadata *DealServingMetadata `json:"dealServingMetadata,omitempty"`

	// DealTerms: The negotiable terms of the deal.
	DealTerms *DealTerms `json:"dealTerms,omitempty"`

	// DeliveryControl: The set of fields around delivery control that are
	// interesting for a buyer
	// to see but are non-negotiable. These are set by the publisher.
	DeliveryControl *DeliveryControl `json:"deliveryControl,omitempty"`

	// Description: Description for the deal terms.
	Description string `json:"description,omitempty"`

	// DisplayName: The name of the deal.
	DisplayName string `json:"displayName,omitempty"`

	// ExternalDealId: The external deal ID assigned to this deal once the
	// deal is finalized.
	// This is the deal ID that shows up in serving/reporting
	// etc.
	// @OutputOnly
	ExternalDealId string `json:"externalDealId,omitempty"`

	// IsSetupComplete: True, if the buyside inventory setup is complete for
	// this deal.
	// @OutputOnly
	IsSetupComplete bool `json:"isSetupComplete,omitempty"`

	// ProgrammaticCreativeSource: Specifies the creative source for
	// programmatic deals. PUBLISHER means
	// creative is provided by seller and ADVERTISER means creative
	// is
	// provided by buyer.
	// @OutputOnly
	//
	// Possible values:
	//   "PROGRAMMATIC_CREATIVE_SOURCE_UNSPECIFIED" - A placeholder for an
	// undefined programmatic creative source.
	//   "ADVERTISER" - The advertiser provides the creatives.
	//   "PUBLISHER" - The publisher provides the creatives to be served.
	ProgrammaticCreativeSource string `json:"programmaticCreativeSource,omitempty"`

	// ProposalId: ID of the proposal that this deal is part of.
	// @OutputOnly
	ProposalId string `json:"proposalId,omitempty"`

	// SellerContacts: Seller contact information for the deal.
	// @OutputOnly
	SellerContacts []*ContactInformation `json:"sellerContacts,omitempty"`

	// SyndicationProduct: The syndication product associated with the
	// deal.
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	//
	// Possible values:
	//   "SYNDICATION_PRODUCT_UNSPECIFIED" - A placeholder for an undefined
	// syndication product.
	//   "CONTENT" - This typically represents a web page.
	//   "MOBILE" - This represents a mobile property.
	//   "VIDEO" - This represents video ad formats.
	//   "GAMES" - This represents ads shown within games.
	SyndicationProduct string `json:"syndicationProduct,omitempty"`

	// Targeting: Specifies the subset of inventory targeted by the
	// deal.
	// @OutputOnly
	Targeting *MarketplaceTargeting `json:"targeting,omitempty"`

	// TargetingCriterion: The shared targeting visible to buyers and
	// sellers. Each shared
	// targeting entity is AND'd together.
	TargetingCriterion []*TargetingCriteria `json:"targetingCriterion,omitempty"`

	// UpdateTime: The time when the deal was last updated.
	// @OutputOnly
	UpdateTime string `json:"updateTime,omitempty"`

	// WebPropertyCode: The web property code for the seller copied over
	// from the product.
	WebPropertyCode string `json:"webPropertyCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AvailableEndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailableEndTime") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Deal) MarshalJSON() ([]byte, error) {
	type NoMethod Deal
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DealPauseStatus: Tracks which parties (if any) have paused a
// deal.
// The deal is considered paused if either hasBuyerPaused
// or
// hasSellPaused is true.
type DealPauseStatus struct {
	// BuyerPauseReason: The buyer's reason for pausing, if the buyer paused
	// the deal.
	BuyerPauseReason string `json:"buyerPauseReason,omitempty"`

	// FirstPausedBy: The role of the person who first paused this deal.
	//
	// Possible values:
	//   "BUYER_SELLER_ROLE_UNSPECIFIED" - A placeholder for an undefined
	// buyer/seller role.
	//   "BUYER" - Specifies the role as buyer.
	//   "SELLER" - Specifies the role as seller.
	FirstPausedBy string `json:"firstPausedBy,omitempty"`

	// HasBuyerPaused: True, if the buyer has paused the deal unilaterally.
	HasBuyerPaused bool `json:"hasBuyerPaused,omitempty"`

	// HasSellerPaused: True, if the seller has paused the deal
	// unilaterally.
	HasSellerPaused bool `json:"hasSellerPaused,omitempty"`

	// SellerPauseReason: The seller's reason for pausing, if the seller
	// paused the deal.
	SellerPauseReason string `json:"sellerPauseReason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuyerPauseReason") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuyerPauseReason") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DealPauseStatus) MarshalJSON() ([]byte, error) {
	type NoMethod DealPauseStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DealServingMetadata: Message captures metadata about the serving
// status of a deal.
type DealServingMetadata struct {
	// DealPauseStatus: Tracks which parties (if any) have paused a
	// deal.
	// @OutputOnly
	DealPauseStatus *DealPauseStatus `json:"dealPauseStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DealPauseStatus") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DealPauseStatus") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DealServingMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod DealServingMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DealTerms: The deal terms specify the details of a Product/deal. They
// specify things
// like price per buyer, the type of pricing model (e.g., fixed price,
// auction)
// and expected impressions from the publisher.
type DealTerms struct {
	// BrandingType: Visibility of the URL in bid requests. (default:
	// BRANDED)
	//
	// Possible values:
	//   "BRANDING_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// branding type.
	//   "BRANDED" - Full URL is included in bid requests.
	//   "SEMI_TRANSPARENT" - A TopLevelDomain or masked URL is sent in bid
	// requests
	// rather than the full one.
	BrandingType string `json:"brandingType,omitempty"`

	// Description: Publisher provided description for the terms.
	Description string `json:"description,omitempty"`

	// EstimatedGrossSpend: Non-binding estimate of the estimated gross
	// spend for this deal.
	// Can be set by buyer or seller.
	EstimatedGrossSpend *Price `json:"estimatedGrossSpend,omitempty"`

	// EstimatedImpressionsPerDay: Non-binding estimate of the impressions
	// served per day.
	// Can be set by buyer or seller.
	EstimatedImpressionsPerDay int64 `json:"estimatedImpressionsPerDay,omitempty,string"`

	// GuaranteedFixedPriceTerms: The terms for guaranteed fixed price
	// deals.
	GuaranteedFixedPriceTerms *GuaranteedFixedPriceTerms `json:"guaranteedFixedPriceTerms,omitempty"`

	// NonGuaranteedAuctionTerms: The terms for non-guaranteed auction
	// deals.
	NonGuaranteedAuctionTerms *NonGuaranteedAuctionTerms `json:"nonGuaranteedAuctionTerms,omitempty"`

	// NonGuaranteedFixedPriceTerms: The terms for non-guaranteed fixed
	// price deals.
	NonGuaranteedFixedPriceTerms *NonGuaranteedFixedPriceTerms `json:"nonGuaranteedFixedPriceTerms,omitempty"`

	// SellerTimeZone: The time zone name. For deals with Cost Per Day
	// billing, defines the
	// time zone used to mark the boundaries of a day. It should be an
	// IANA TZ name, such as "America/Los_Angeles". For more
	// information,
	// see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	SellerTimeZone string `json:"sellerTimeZone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BrandingType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BrandingType") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DealTerms) MarshalJSON() ([]byte, error) {
	type NoMethod DealTerms
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeliveryControl: Message contains details about how the deals will be
// paced.
type DeliveryControl struct {
	// CreativeBlockingLevel: Specified the creative blocking levels to be
	// applied.
	// @OutputOnly
	//
	// Possible values:
	//   "CREATIVE_BLOCKING_LEVEL_UNSPECIFIED" - A placeholder for an
	// undefined creative blocking level.
	//   "PUBLISHER_BLOCKING_RULES" - Publisher blocking rules will be
	// applied.
	//   "ADX_POLICY_BLOCKING_ONLY" - The Ad Exchange policy blocking rules
	// will be applied.
	CreativeBlockingLevel string `json:"creativeBlockingLevel,omitempty"`

	// DeliveryRateType: Specifies how the impression delivery will be
	// paced.
	// @OutputOnly
	//
	// Possible values:
	//   "DELIVERY_RATE_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// delivery rate type.
	//   "EVENLY" - Impressions are served uniformly over the life of the
	// deal.
	//   "FRONT_LOADED" - Impressions are served front-loaded.
	//   "AS_FAST_AS_POSSIBLE" - Impressions are served as fast as possible.
	DeliveryRateType string `json:"deliveryRateType,omitempty"`

	// FrequencyCaps: Specifies any frequency caps.
	// @OutputOnly
	FrequencyCaps []*FrequencyCap `json:"frequencyCaps,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CreativeBlockingLevel") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeBlockingLevel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DeliveryControl) MarshalJSON() ([]byte, error) {
	type NoMethod DeliveryControl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Disapproval: @OutputOnly The reason and details for a disapproval.
type Disapproval struct {
	// Details: Additional details about the reason for disapproval.
	Details []string `json:"details,omitempty"`

	// Reason: The categorized reason for disapproval.
	//
	// Possible values:
	//   "LENGTH_OF_IMAGE_ANIMATION" - The length of the image animation is
	// longer than allowed.
	//   "BROKEN_URL" - The click through URL doesn't work properly.
	//   "MEDIA_NOT_FUNCTIONAL" - Something is wrong with the creative
	// itself.
	//   "INVALID_FOURTH_PARTY_CALL" - The ad makes a fourth party call to
	// an unapproved vendor.
	//   "INCORRECT_REMARKETING_DECLARATION" - The ad targets consumers
	// using remarketing lists and/or collects
	// data for subsequent use in retargeting, but does not correctly
	// declare
	// that use.
	//   "LANDING_PAGE_ERROR" - Clicking on the ad leads to an error page.
	//   "AD_SIZE_DOES_NOT_MATCH_AD_SLOT" - The ad size when rendered does
	// not match the declaration.
	//   "NO_BORDER" - Ads with a white background require a border, which
	// was missing.
	//   "FOURTH_PARTY_BROWSER_COOKIES" - The creative attempts to set
	// cookies from a fourth party that is not
	// certified.
	//   "LSO_OBJECTS" - The creative sets an LSO object.
	//   "BLANK_CREATIVE" - The ad serves a blank.
	//   "DESTINATION_URLS_UNDECLARED" - The ad uses rotation, but not all
	// destination URLs were declared.
	//   "PROBLEM_WITH_CLICK_MACRO" - There is a problem with the way the
	// click macro is used.
	//   "INCORRECT_AD_TECHNOLOGY_DECLARATION" - The ad technology
	// declaration is not accurate.
	//   "INCORRECT_DESTINATION_URL_DECLARATION" - The actual destination
	// URL does not match the declared destination URL.
	//   "EXPANDABLE_INCORRECT_DIRECTION" - The declared expanding direction
	// does not match the actual direction.
	//   "EXPANDABLE_DIRECTION_NOT_SUPPORTED" - The ad does not expand in a
	// supported direction.
	//   "EXPANDABLE_INVALID_VENDOR" - The ad uses an expandable vendor that
	// is not supported.
	//   "EXPANDABLE_FUNCTIONALITY" - There was an issue with the expandable
	// ad.
	//   "VIDEO_INVALID_VENDOR" - The ad uses a video vendor that is not
	// supported.
	//   "VIDEO_UNSUPPORTED_LENGTH" - The length of the video ad is not
	// supported.
	//   "VIDEO_UNSUPPORTED_FORMAT" - The format of the video ad is not
	// supported.
	//   "VIDEO_FUNCTIONALITY" - There was an issue with the video ad.
	//   "LANDING_PAGE_DISABLED" - The landing page does not conform to Ad
	// Exchange policy.
	//   "MALWARE_SUSPECTED" - The ad or the landing page may contain
	// malware.
	//   "ADULT_IMAGE_OR_VIDEO" - The ad contains adult images or video
	// content.
	//   "INACCURATE_AD_TEXT" - The ad contains text that is unclear or
	// inaccurate.
	//   "COUNTERFEIT_DESIGNER_GOODS" - The ad promotes counterfeit designer
	// goods.
	//   "POP_UP" - The ad causes a popup window to appear.
	//   "INVALID_RTB_PROTOCOL_USAGE" - The creative does not follow
	// policies set for the RTB protocol.
	//   "RAW_IP_ADDRESS_IN_SNIPPET" - The ad contains a URL that uses a
	// numeric IP address for the domain.
	//   "UNACCEPTABLE_CONTENT_SOFTWARE" - The ad or landing page contains
	// unacceptable content because it initiated
	// a software or executable download.
	//   "UNAUTHORIZED_COOKIE_ON_GOOGLE_DOMAIN" - The ad set an unauthorized
	// cookie on a Google domain.
	//   "UNDECLARED_FLASH_OBJECTS" - Flash content found when no flash was
	// declared.
	//   "INVALID_SSL_DECLARATION" - SSL support declared but not working
	// correctly.
	//   "DIRECT_DOWNLOAD_IN_AD" - Rich Media - Direct Download in Ad (ex.
	// PDF download).
	//   "MAXIMUM_DOWNLOAD_SIZE_EXCEEDED" - Maximum download size exceeded.
	//   "DESTINATION_URL_SITE_NOT_CRAWLABLE" - Bad Destination URL: Site
	// Not Crawlable.
	//   "BAD_URL_LEGAL_DISAPPROVAL" - Bad URL: Legal disapproval.
	//   "PHARMA_GAMBLING_ALCOHOL_NOT_ALLOWED" - Pharmaceuticals, Gambling,
	// Alcohol not allowed and at least one was
	// detected.
	//   "DYNAMIC_DNS_AT_DESTINATION_URL" - Dynamic DNS at Destination URL.
	//   "POOR_IMAGE_OR_VIDEO_QUALITY" - Poor Image / Video Quality.
	//   "UNACCEPTABLE_IMAGE_CONTENT" - For example, Image Trick to Click.
	//   "INCORRECT_IMAGE_LAYOUT" - Incorrect Image Layout.
	//   "IRRELEVANT_IMAGE_OR_VIDEO" - Irrelevant Image / Video.
	//   "DESTINATION_SITE_DOES_NOT_ALLOW_GOING_BACK" - Broken back button.
	//   "MISLEADING_CLAIMS_IN_AD" - Misleading/Inaccurate claims in ads.
	//   "RESTRICTED_PRODUCTS" - Restricted Products.
	//   "UNACCEPTABLE_CONTENT" - Unacceptable content. For example,
	// malware.
	//   "AUTOMATED_AD_CLICKING" - The ad automatically redirects to the
	// destination site without a click,
	// or reports a click when none were made.
	//   "INVALID_URL_PROTOCOL" - The ad uses URL protocols that do not
	// exist or are not allowed on AdX.
	//   "UNDECLARED_RESTRICTED_CONTENT" - Restricted content (for example,
	// alcohol) was found in the ad but not
	// declared.
	//   "INVALID_REMARKETING_LIST_USAGE" - Violation of the remarketing
	// list policy.
	//   "DESTINATION_SITE_NOT_CRAWLABLE_ROBOTS_TXT" - The destination
	// site's robot.txt file prevents it from being crawled.
	//   "CLICK_TO_DOWNLOAD_NOT_AN_APP" - Click to download must link to an
	// app.
	//   "INACCURATE_REVIEW_EXTENSION" - A review extension must be an
	// accurate review.
	//   "SEXUALLY_EXPLICIT_CONTENT" - Sexually explicit content.
	//   "GAINING_AN_UNFAIR_ADVANTAGE" - The ad tries to gain an unfair
	// traffic advantage.
	//   "GAMING_THE_GOOGLE_NETWORK" - The ad tries to circumvent Google's
	// advertising systems.
	//   "DANGEROUS_PRODUCTS_KNIVES" - The ad promotes dangerous knives.
	//   "DANGEROUS_PRODUCTS_EXPLOSIVES" - The ad promotes explosives.
	//   "DANGEROUS_PRODUCTS_GUNS" - The ad promotes guns & parts.
	//   "DANGEROUS_PRODUCTS_DRUGS" - The ad promotes recreational
	// drugs/services & related equipment.
	//   "DANGEROUS_PRODUCTS_TOBACCO" - The ad promotes tobacco
	// products/services & related equipment.
	//   "DANGEROUS_PRODUCTS_WEAPONS" - The ad promotes weapons.
	//   "UNCLEAR_OR_IRRELEVANT_AD" - The ad is unclear or irrelevant to the
	// destination site.
	//   "PROFESSIONAL_STANDARDS" - The ad does not meet professional
	// standards.
	//   "DYSFUNCTIONAL_PROMOTION" - The promotion is unnecessarily
	// difficult to navigate.
	//   "INVALID_INTEREST_BASED_AD" - Violation of Google's policy for
	// interest-based ads.
	//   "MISUSE_OF_PERSONAL_INFORMATION" - Misuse of personal information.
	//   "OMISSION_OF_RELEVANT_INFORMATION" - Omission of relevant
	// information.
	//   "UNAVAILABLE_PROMOTIONS" - Unavailable promotions.
	//   "MISLEADING_PROMOTIONS" - Misleading or unrealistic promotions.
	//   "INAPPROPRIATE_CONTENT" - Offensive or inappropriate content.
	//   "SENSITIVE_EVENTS" - Capitalizing on sensitive events.
	//   "SHOCKING_CONTENT" - Shocking content.
	//   "ENABLING_DISHONEST_BEHAVIOR" - Products & Services that enable
	// dishonest behavior.
	//   "TECHNICAL_REQUIREMENTS" - The ad does not meet technical
	// requirements.
	//   "RESTRICTED_POLITICAL_CONTENT" - Restricted political content.
	//   "UNSUPPORTED_CONTENT" - Unsupported content.
	//   "INVALID_BIDDING_METHOD" - Invalid bidding method.
	//   "VIDEO_TOO_LONG" - Video length exceeds limits.
	//   "VIOLATES_JAPANESE_PHARMACY_LAW" - Unacceptable content: Japanese
	// healthcare.
	//   "UNACCREDITED_PET_PHARMACY" - Online pharmacy ID required.
	//   "ABORTION" - Unacceptable content: Abortion.
	//   "CONTRACEPTIVES" - Unacceptable content: Birth control.
	//   "NEED_CERTIFICATES_TO_ADVERTISE_IN_CHINA" - Restricted in China.
	//   "KCDSP_REGISTRATION" - Unacceptable content: Korean healthcare.
	//   "NOT_FAMILY_SAFE" - Non-family safe or adult content.
	//   "CLINICAL_TRIAL_RECRUITMENT" - Clinical trial recruitment.
	//   "MAXIMUM_NUMBER_OF_HTTP_CALLS_EXCEEDED" - Maximum number of HTTP
	// calls exceeded.
	//   "MAXIMUM_NUMBER_OF_COOKIES_EXCEEDED" - Maximum number of cookies
	// exceeded.
	//   "PERSONAL_LOANS" - Financial service ad does not adhere to
	// specifications.
	//   "UNSUPPORTED_FLASH_CONTENT" - Flash content was found in an
	// unsupported context.
	//   "MISUSE_BY_OMID_SCRIPT" - Misuse by an Open Measurement SDK script.
	//   "NON_WHITELISTED_OMID_VENDOR" - Use of an Open Measurement SDK
	// vendor not on approved whitelist.
	//   "DESTINATION_EXPERIENCE" - Unacceptable landing page.
	//   "UNSUPPORTED_LANGUAGE" - Unsupported language.
	//   "NON_SSL_COMPLIANT" - Non-SSL compliant.
	//   "TEMPORARY_PAUSE" - Temporary pausing of creative.
	//   "BAIL_BONDS" - Promotes services related to bail bonds.
	Reason string `json:"reason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Details") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Details") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Disapproval) MarshalJSON() ([]byte, error) {
	type NoMethod Disapproval
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// FilterSet: A set of filters that is applied to a request for
// data.
// Within a filter set, an AND operation is performed across the
// filters
// represented by each field. An OR operation is performed across the
// filters
// represented by the multiple values of a repeated field,
// e.g.,
// "format=VIDEO AND deal_id=12 AND (seller_network_id=34
// OR
// seller_network_id=56)".
type FilterSet struct {
	// AbsoluteDateRange: An absolute date range, defined by a start date
	// and an end date.
	// Interpreted relative to Pacific time zone.
	AbsoluteDateRange *AbsoluteDateRange `json:"absoluteDateRange,omitempty"`

	// BreakdownDimensions: The set of dimensions along which to break down
	// the response; may be empty.
	// If multiple dimensions are requested, the breakdown is along the
	// Cartesian
	// product of the requested dimensions.
	//
	// Possible values:
	//   "BREAKDOWN_DIMENSION_UNSPECIFIED" - A placeholder for an
	// unspecified dimension; should not be used.
	//   "PUBLISHER_IDENTIFIER" - The response should be broken down by
	// publisher identifier.
	// This option is available only for Open Bidding buyers.
	BreakdownDimensions []string `json:"breakdownDimensions,omitempty"`

	// CreativeId: The ID of the creative on which to filter; optional. This
	// field may be set
	// only for a filter set that accesses account-level troubleshooting
	// data,
	// i.e., one whose name matches the
	// `bidders/*/accounts/*/filterSets/*`
	// pattern.
	CreativeId string `json:"creativeId,omitempty"`

	// DealId: The ID of the deal on which to filter; optional. This field
	// may be set
	// only for a filter set that accesses account-level troubleshooting
	// data,
	// i.e., one whose name matches the
	// `bidders/*/accounts/*/filterSets/*`
	// pattern.
	DealId int64 `json:"dealId,omitempty,string"`

	// Environment: The environment on which to filter; optional.
	//
	// Possible values:
	//   "ENVIRONMENT_UNSPECIFIED" - A placeholder for an undefined
	// environment; indicates that no environment
	// filter will be applied.
	//   "WEB" - The ad impression appears on the web.
	//   "APP" - The ad impression appears in an app.
	Environment string `json:"environment,omitempty"`

	// Formats: The list of formats on which to filter; may be empty. The
	// filters
	// represented by multiple formats are ORed together (i.e., if
	// non-empty,
	// results must match any one of the formats).
	//
	// Possible values:
	//   "FORMAT_UNSPECIFIED" - A placeholder for an undefined format;
	// indicates that no format filter
	// will be applied.
	//   "NATIVE_DISPLAY" - The ad impression is a native ad, and display
	// (i.e., image) format.
	//   "NATIVE_VIDEO" - The ad impression is a native ad, and video
	// format.
	//   "NON_NATIVE_DISPLAY" - The ad impression is not a native ad, and
	// display (i.e., image) format.
	//   "NON_NATIVE_VIDEO" - The ad impression is not a native ad, and
	// video format.
	Formats []string `json:"formats,omitempty"`

	// Name: A user-defined name of the filter set. Filter set names must be
	// unique
	// globally and match one of the patterns:
	//
	// - `bidders/*/filterSets/*` (for accessing bidder-level
	// troubleshooting
	// data)
	// - `bidders/*/accounts/*/filterSets/*` (for accessing
	// account-level
	// troubleshooting data)
	//
	// This field is required in create operations.
	Name string `json:"name,omitempty"`

	// Platforms: The list of platforms on which to filter; may be empty.
	// The filters
	// represented by multiple platforms are ORed together (i.e., if
	// non-empty,
	// results must match any one of the platforms).
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - A placeholder for an undefined platform;
	// indicates that no platform
	// filter will be applied.
	//   "DESKTOP" - The ad impression appears on a desktop.
	//   "TABLET" - The ad impression appears on a tablet.
	//   "MOBILE" - The ad impression appears on a mobile device.
	Platforms []string `json:"platforms,omitempty"`

	// PublisherIdentifiers: For Open Bidding partners only.
	// The list of publisher identifiers on which to filter; may be
	// empty.
	// The filters represented by multiple publisher identifiers are
	// ORed
	// together.
	PublisherIdentifiers []string `json:"publisherIdentifiers,omitempty"`

	// RealtimeTimeRange: An open-ended realtime time range, defined by the
	// aggregation start
	// timestamp.
	RealtimeTimeRange *RealtimeTimeRange `json:"realtimeTimeRange,omitempty"`

	// RelativeDateRange: A relative date range, defined by an offset from
	// today and a duration.
	// Interpreted relative to Pacific time zone.
	RelativeDateRange *RelativeDateRange `json:"relativeDateRange,omitempty"`

	// SellerNetworkIds: For Authorized Buyers only.
	// The list of IDs of the seller (publisher) networks on which to
	// filter;
	// may be empty. The filters represented by multiple seller network IDs
	// are
	// ORed together (i.e., if non-empty, results must match any one of
	// the
	// publisher networks).
	// See
	// [seller-network-ids](https://developers.google.com/authorized-buye
	// rs/rtb/downloads/seller-network-ids)
	// file for the set of existing seller network IDs.
	SellerNetworkIds []int64 `json:"sellerNetworkIds,omitempty"`

	// TimeSeriesGranularity: The granularity of time intervals if a time
	// series breakdown is desired;
	// optional.
	//
	// Possible values:
	//   "TIME_SERIES_GRANULARITY_UNSPECIFIED" - A placeholder for an
	// unspecified interval; no time series is applied.
	// All rows in response will contain data for the entire requested time
	// range.
	//   "HOURLY" - Indicates that data will be broken down by the hour.
	//   "DAILY" - Indicates that data will be broken down by the day.
	TimeSeriesGranularity string `json:"timeSeriesGranularity,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AbsoluteDateRange")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AbsoluteDateRange") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FilterSet) MarshalJSON() ([]byte, error) {
	type NoMethod FilterSet
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilteredBidCreativeRow: The number of filtered bids with the
// specified dimension values that have the
// specified creative.
type FilteredBidCreativeRow struct {
	// BidCount: The number of bids with the specified creative.
	BidCount *MetricValue `json:"bidCount,omitempty"`

	// CreativeId: The ID of the creative.
	CreativeId string `json:"creativeId,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BidCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BidCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilteredBidCreativeRow) MarshalJSON() ([]byte, error) {
	type NoMethod FilteredBidCreativeRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilteredBidDetailRow: The number of filtered bids with the specified
// dimension values, among those
// filtered due to the requested filtering reason (i.e. creative
// status), that
// have the specified detail.
type FilteredBidDetailRow struct {
	// BidCount: The number of bids with the specified detail.
	BidCount *MetricValue `json:"bidCount,omitempty"`

	// DetailId: The ID of the detail. The associated value can be looked up
	// in the
	// dictionary file corresponding to the DetailType in the response
	// message.
	DetailId int64 `json:"detailId,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BidCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BidCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilteredBidDetailRow) MarshalJSON() ([]byte, error) {
	type NoMethod FilteredBidDetailRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FilteringStats: @OutputOnly Filtering reasons for this creative
// during a period of a single
// day (from midnight to midnight Pacific).
type FilteringStats struct {
	// Date: The day during which the data was collected.
	// The data is collected from 00:00:00 to 23:59:59 PT.
	// During switches from PST to PDT and back, the day may
	// contain 23 or 25 hours of data instead of the usual 24.
	Date *Date `json:"date,omitempty"`

	// Reasons: The set of filtering reasons for this date.
	Reasons []*Reason `json:"reasons,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Date") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Date") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FilteringStats) MarshalJSON() ([]byte, error) {
	type NoMethod FilteringStats
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FirstPartyMobileApplicationTargeting: Represents a list of targeted
// and excluded mobile application IDs that
// publishers own.
// Mobile application IDs are from App Store and Google Play
// Store.
// Android App ID, for example, com.google.android.apps.maps, can be
// found in
// Google Play Store URL.
// iOS App ID (which is a number) can be found at the end of iTunes
// store URL.
// First party mobile applications is either included or excluded.
type FirstPartyMobileApplicationTargeting struct {
	// ExcludedAppIds: A list of application IDs to be excluded.
	ExcludedAppIds []string `json:"excludedAppIds,omitempty"`

	// TargetedAppIds: A list of application IDs to be included.
	TargetedAppIds []string `json:"targetedAppIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExcludedAppIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludedAppIds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FirstPartyMobileApplicationTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod FirstPartyMobileApplicationTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FrequencyCap: Frequency cap.
type FrequencyCap struct {
	// MaxImpressions: The maximum number of impressions that can be served
	// to a user within the
	// specified time period.
	MaxImpressions int64 `json:"maxImpressions,omitempty"`

	// NumTimeUnits: The amount of time, in the units specified by
	// time_unit_type. Defines the
	// amount of time over which impressions per user are counted and
	// capped.
	NumTimeUnits int64 `json:"numTimeUnits,omitempty"`

	// TimeUnitType: The time unit. Along with num_time_units defines the
	// amount of time over
	// which impressions per user are counted and capped.
	//
	// Possible values:
	//   "TIME_UNIT_TYPE_UNSPECIFIED" - A placeholder for an undefined time
	// unit type. This just indicates the
	// variable with this value hasn't been initialized.
	//   "MINUTE" - Minute
	//   "HOUR" - Hour
	//   "DAY" - Day
	//   "WEEK" - Week
	//   "MONTH" - Month
	//   "LIFETIME" - Lifetime
	TimeUnitType string `json:"timeUnitType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxImpressions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxImpressions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FrequencyCap) MarshalJSON() ([]byte, error) {
	type NoMethod FrequencyCap
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GuaranteedFixedPriceTerms: Terms for Programmatic Guaranteed Deals.
type GuaranteedFixedPriceTerms struct {
	// FixedPrices: Fixed price for the specified buyer.
	FixedPrices []*PricePerBuyer `json:"fixedPrices,omitempty"`

	// GuaranteedImpressions: Guaranteed impressions as a percentage. This
	// is the percentage
	// of guaranteed looks that the buyer is guaranteeing to buy.
	GuaranteedImpressions int64 `json:"guaranteedImpressions,omitempty,string"`

	// GuaranteedLooks: Count of guaranteed looks. Required for deal,
	// optional for product.
	GuaranteedLooks int64 `json:"guaranteedLooks,omitempty,string"`

	// MinimumDailyLooks: Daily minimum looks for CPD deal types.
	MinimumDailyLooks int64 `json:"minimumDailyLooks,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "FixedPrices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FixedPrices") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GuaranteedFixedPriceTerms) MarshalJSON() ([]byte, error) {
	type NoMethod GuaranteedFixedPriceTerms
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HtmlContent: HTML content for a creative.
type HtmlContent struct {
	// Height: The height of the HTML snippet in pixels.
	Height int64 `json:"height,omitempty"`

	// Snippet: The HTML snippet that displays the ad when inserted in the
	// web page.
	Snippet string `json:"snippet,omitempty"`

	// Width: The width of the HTML snippet in pixels.
	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HtmlContent) MarshalJSON() ([]byte, error) {
	type NoMethod HtmlContent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Image: An image resource. You may provide a larger image than was
// requested,
// so long as the aspect ratio is preserved.
type Image struct {
	// Height: Image height in pixels.
	Height int64 `json:"height,omitempty"`

	// Url: The URL of the image.
	Url string `json:"url,omitempty"`

	// Width: Image width in pixels.
	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Image) MarshalJSON() ([]byte, error) {
	type NoMethod Image
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImpressionMetricsRow: The set of metrics that are measured in numbers
// of impressions, representing
// how many impressions with the specified dimension values were
// considered
// eligible at each stage of the bidding funnel.
type ImpressionMetricsRow struct {
	// AvailableImpressions: The number of impressions available to the
	// buyer on Ad Exchange.
	// In some cases this value may be unavailable.
	AvailableImpressions *MetricValue `json:"availableImpressions,omitempty"`

	// BidRequests: The number of impressions for which Ad Exchange sent the
	// buyer a bid
	// request.
	BidRequests *MetricValue `json:"bidRequests,omitempty"`

	// InventoryMatches: The number of impressions that match the buyer's
	// inventory pretargeting.
	InventoryMatches *MetricValue `json:"inventoryMatches,omitempty"`

	// ResponsesWithBids: The number of impressions for which Ad Exchange
	// received a response from
	// the buyer that contained at least one applicable bid.
	ResponsesWithBids *MetricValue `json:"responsesWithBids,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// SuccessfulResponses: The number of impressions for which the buyer
	// successfully sent a response
	// to Ad Exchange.
	SuccessfulResponses *MetricValue `json:"successfulResponses,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AvailableImpressions") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailableImpressions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ImpressionMetricsRow) MarshalJSON() ([]byte, error) {
	type NoMethod ImpressionMetricsRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InventorySizeTargeting: Represents the size of an ad unit that can be
// targeted on an ad
// request. It only applies to Private Auction, AdX Preferred Deals
// and
// Auction Packages. This targeting does not apply to Programmatic
// Guaranteed
// and Preferred Deals in Ad Manager.
type InventorySizeTargeting struct {
	// ExcludedInventorySizes: A list of inventory sizes to be excluded.
	ExcludedInventorySizes []*AdSize `json:"excludedInventorySizes,omitempty"`

	// TargetedInventorySizes: A list of inventory sizes to be included.
	TargetedInventorySizes []*AdSize `json:"targetedInventorySizes,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExcludedInventorySizes") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludedInventorySizes")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *InventorySizeTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod InventorySizeTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBidMetricsResponse: Response message for listing the metrics that
// are measured in number of bids.
type ListBidMetricsResponse struct {
	// BidMetricsRows: List of rows, each containing a set of bid metrics.
	BidMetricsRows []*BidMetricsRow `json:"bidMetricsRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListBidMetricsRequest.pageToken
	// field in the subsequent call to the bidMetrics.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BidMetricsRows") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BidMetricsRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListBidMetricsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBidMetricsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBidResponseErrorsResponse: Response message for listing all
// reasons that bid responses resulted in an
// error.
type ListBidResponseErrorsResponse struct {
	// CalloutStatusRows: List of rows, with counts of bid responses
	// aggregated by callout status.
	CalloutStatusRows []*CalloutStatusRow `json:"calloutStatusRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListBidResponseErrorsRequest.pageToken
	// field in the subsequent call to the bidResponseErrors.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CalloutStatusRows")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CalloutStatusRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListBidResponseErrorsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBidResponseErrorsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBidResponsesWithoutBidsResponse: Response message for listing all
// reasons that bid responses were considered
// to have no applicable bids.
type ListBidResponsesWithoutBidsResponse struct {
	// BidResponseWithoutBidsStatusRows: List of rows, with counts of bid
	// responses without bids aggregated by
	// status.
	BidResponseWithoutBidsStatusRows []*BidResponseWithoutBidsStatusRow `json:"bidResponseWithoutBidsStatusRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in
	// the
	// ListBidResponsesWithoutBidsRequest.pageToken
	// field in the subsequent call to the
	// bidResponsesWithoutBids.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "BidResponseWithoutBidsStatusRows") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "BidResponseWithoutBidsStatusRows") to include in API requests with
	// the JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListBidResponsesWithoutBidsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBidResponsesWithoutBidsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListClientUserInvitationsResponse struct {
	// Invitations: The returned list of client users.
	Invitations []*ClientUserInvitation `json:"invitations,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in
	// the
	// ListClientUserInvitationsRequest.pageToken
	// field in the subsequent call to the
	// clients.invitations.list
	// method to retrieve the next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Invitations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Invitations") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListClientUserInvitationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListClientUserInvitationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListClientUsersResponse struct {
	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListClientUsersRequest.pageToken
	// field in the subsequent call to the
	// clients.invitations.list
	// method to retrieve the next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Users: The returned list of client users.
	Users []*ClientUser `json:"users,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListClientUsersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListClientUsersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListClientsResponse struct {
	// Clients: The returned list of clients.
	Clients []*Client `json:"clients,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListClientsRequest.pageToken
	// field in the subsequent call to the
	// accounts.clients.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Clients") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Clients") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListClientsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListClientsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCreativeStatusBreakdownByCreativeResponse: Response message for
// listing all creatives associated with a given filtered
// bid reason.
type ListCreativeStatusBreakdownByCreativeResponse struct {
	// FilteredBidCreativeRows: List of rows, with counts of bids with a
	// given creative status aggregated
	// by creative.
	FilteredBidCreativeRows []*FilteredBidCreativeRow `json:"filteredBidCreativeRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in
	// the
	// ListCreativeStatusBreakdownByCreativeRequest.pageToken
	// field in the subsequent call to the
	// filteredBids.creatives.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "FilteredBidCreativeRows") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FilteredBidCreativeRows")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListCreativeStatusBreakdownByCreativeResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCreativeStatusBreakdownByCreativeResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCreativeStatusBreakdownByDetailResponse: Response message for
// listing all details associated with a given filtered bid
// reason.
type ListCreativeStatusBreakdownByDetailResponse struct {
	// DetailType: The type of detail that the detail IDs represent.
	//
	// Possible values:
	//   "DETAIL_TYPE_UNSPECIFIED" - A placeholder for an undefined
	// status.
	// This value will never be returned in responses.
	//   "CREATIVE_ATTRIBUTE" - Indicates that the detail ID refers to a
	// creative attribute;
	// see
	// [publisher-excludable-creative-attributes](https://developers.goog
	// le.com/authorized-buyers/rtb/downloads/publisher-excludable-creative-a
	// ttributes).
	//   "VENDOR" - Indicates that the detail ID refers to a vendor;
	// see
	// [vendors](https://developers.google.com/authorized-buyers/rtb/down
	// loads/vendors).
	//   "SENSITIVE_CATEGORY" - Indicates that the detail ID refers to a
	// sensitive category;
	// see
	// [ad-sensitive-categories](https://developers.google.com/authorized
	// -buyers/rtb/downloads/ad-sensitive-categories).
	//   "PRODUCT_CATEGORY" - Indicates that the detail ID refers to a
	// product category;
	// see
	// [ad-product-categories](https://developers.google.com/authorized-b
	// uyers/rtb/downloads/ad-product-categories).
	//   "DISAPPROVAL_REASON" - Indicates that the detail ID refers to a
	// disapproval reason; see
	// DisapprovalReason enum in
	// [snippet-status-report-proto](https://developers.google.com/authorized
	// -buyers/rtb/downloads/snippet-status-report-proto).
	DetailType string `json:"detailType,omitempty"`

	// FilteredBidDetailRows: List of rows, with counts of bids with a given
	// creative status aggregated
	// by detail.
	FilteredBidDetailRows []*FilteredBidDetailRow `json:"filteredBidDetailRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in
	// the
	// ListCreativeStatusBreakdownByDetailRequest.pageToken
	// field in the subsequent call to the filteredBids.details.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DetailType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DetailType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCreativeStatusBreakdownByDetailResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCreativeStatusBreakdownByDetailResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCreativesResponse: A response for listing creatives.
type ListCreativesResponse struct {
	// Creatives: The list of creatives.
	Creatives []*Creative `json:"creatives,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListCreativesRequest.page_token
	// field in the subsequent call to `ListCreatives` method to retrieve
	// the next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Creatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Creatives") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCreativesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCreativesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListDealAssociationsResponse: A response for listing creative and
// deal associations
type ListDealAssociationsResponse struct {
	// Associations: The list of associations.
	Associations []*CreativeDealAssociation `json:"associations,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListDealAssociationsRequest.page_token
	// field in the subsequent call to 'ListDealAssociation' method to
	// retrieve
	// the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Associations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Associations") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListDealAssociationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListDealAssociationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListFilterSetsResponse: Response message for listing filter sets.
type ListFilterSetsResponse struct {
	// FilterSets: The filter sets belonging to the buyer.
	FilterSets []*FilterSet `json:"filterSets,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListFilterSetsRequest.pageToken
	// field in the subsequent call to the
	// accounts.filterSets.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "FilterSets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FilterSets") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListFilterSetsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListFilterSetsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListFilteredBidRequestsResponse: Response message for listing all
// reasons that bid requests were filtered and
// not sent to the buyer.
type ListFilteredBidRequestsResponse struct {
	// CalloutStatusRows: List of rows, with counts of filtered bid requests
	// aggregated by callout
	// status.
	CalloutStatusRows []*CalloutStatusRow `json:"calloutStatusRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListFilteredBidRequestsRequest.pageToken
	// field in the subsequent call to the filteredBidRequests.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CalloutStatusRows")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CalloutStatusRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListFilteredBidRequestsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListFilteredBidRequestsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListFilteredBidsResponse: Response message for listing all reasons
// that bids were filtered from the
// auction.
type ListFilteredBidsResponse struct {
	// CreativeStatusRows: List of rows, with counts of filtered bids
	// aggregated by filtering reason
	// (i.e. creative status).
	CreativeStatusRows []*CreativeStatusRow `json:"creativeStatusRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListFilteredBidsRequest.pageToken
	// field in the subsequent call to the filteredBids.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreativeStatusRows")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeStatusRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListFilteredBidsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListFilteredBidsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListImpressionMetricsResponse: Response message for listing the
// metrics that are measured in number of
// impressions.
type ListImpressionMetricsResponse struct {
	// ImpressionMetricsRows: List of rows, each containing a set of
	// impression metrics.
	ImpressionMetricsRows []*ImpressionMetricsRow `json:"impressionMetricsRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListImpressionMetricsRequest.pageToken
	// field in the subsequent call to the impressionMetrics.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ImpressionMetricsRows") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImpressionMetricsRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListImpressionMetricsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListImpressionMetricsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListLosingBidsResponse: Response message for listing all reasons that
// bids lost in the auction.
type ListLosingBidsResponse struct {
	// CreativeStatusRows: List of rows, with counts of losing bids
	// aggregated by loss reason (i.e.
	// creative status).
	CreativeStatusRows []*CreativeStatusRow `json:"creativeStatusRows,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in the
	// ListLosingBidsRequest.pageToken
	// field in the subsequent call to the losingBids.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreativeStatusRows")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeStatusRows") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListLosingBidsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListLosingBidsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListNonBillableWinningBidsResponse: Response message for listing all
// reasons for which a buyer was not billed for
// a winning bid.
type ListNonBillableWinningBidsResponse struct {
	// NextPageToken: A token to retrieve the next page of results.
	// Pass this value in
	// the
	// ListNonBillableWinningBidsRequest.pageToken
	// field in the subsequent call to the
	// nonBillableWinningBids.list
	// method to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NonBillableWinningBidStatusRows: List of rows, with counts of bids
	// not billed aggregated by reason.
	NonBillableWinningBidStatusRows []*NonBillableWinningBidStatusRow `json:"nonBillableWinningBidStatusRows,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListNonBillableWinningBidsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListNonBillableWinningBidsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListProductsResponse: Response message for listing products visible
// to the buyer.
type ListProductsResponse struct {
	// NextPageToken: List pagination support.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Products: The list of matching products at their head revision
	// number.
	Products []*Product `json:"products,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListProductsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListProductsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListProposalsResponse: Response message for listing proposals.
type ListProposalsResponse struct {
	// NextPageToken: Continuation token for fetching the next page of
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Proposals: The list of proposals.
	Proposals []*Proposal `json:"proposals,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListProposalsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListProposalsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListPublisherProfilesResponse: Response message for profiles visible
// to the buyer.
type ListPublisherProfilesResponse struct {
	// NextPageToken: List pagination support
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PublisherProfiles: The list of matching publisher profiles.
	PublisherProfiles []*PublisherProfile `json:"publisherProfiles,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListPublisherProfilesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListPublisherProfilesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LocationContext: @OutputOnly The Geo criteria the restriction applies
// to.
type LocationContext struct {
	// GeoCriteriaIds: IDs representing the geo location for this
	// context.
	// Please refer to
	// the
	// [geo-table.csv](https://storage.googleapis.com/adx-rtb-dictionarie
	// s/geo-table.csv)
	// file for different geo criteria IDs.
	GeoCriteriaIds []int64 `json:"geoCriteriaIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GeoCriteriaIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GeoCriteriaIds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LocationContext) MarshalJSON() ([]byte, error) {
	type NoMethod LocationContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MarketplaceTargeting: Targeting represents different criteria that
// can be used by advertisers to
// target ad inventory. For example, they can choose to target ad
// requests only
// if the user is in the US.
// Multiple types of targeting are always applied as a logical AND,
// unless noted
// otherwise.
type MarketplaceTargeting struct {
	// GeoTargeting: Geo criteria IDs to be included/excluded.
	GeoTargeting *CriteriaTargeting `json:"geoTargeting,omitempty"`

	// InventorySizeTargeting: Inventory sizes to be included/excluded.
	InventorySizeTargeting *InventorySizeTargeting `json:"inventorySizeTargeting,omitempty"`

	// PlacementTargeting: Placement targeting information, e.g., URL,
	// mobile applications.
	PlacementTargeting *PlacementTargeting `json:"placementTargeting,omitempty"`

	// TechnologyTargeting: Technology targeting information, e.g.,
	// operating system, device category.
	TechnologyTargeting *TechnologyTargeting `json:"technologyTargeting,omitempty"`

	// VideoTargeting: Video targeting information.
	VideoTargeting *VideoTargeting `json:"videoTargeting,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GeoTargeting") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GeoTargeting") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MarketplaceTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod MarketplaceTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricValue: A metric value, with an expected value and a variance;
// represents a count
// that may be either exact or estimated (i.e. when sampled).
type MetricValue struct {
	// Value: The expected value of the metric.
	Value int64 `json:"value,omitempty,string"`

	// Variance: The variance (i.e. square of the standard deviation) of the
	// metric value.
	// If value is exact, variance is 0.
	// Can be used to calculate margin of error as a percentage of value,
	// using
	// the following formula, where Z is the standard constant that depends
	// on the
	// desired size of the confidence interval (e.g. for 90% confidence
	// interval,
	// use Z = 1.645):
	//
	//   marginOfError = 100 * Z * sqrt(variance) / value
	Variance int64 `json:"variance,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Value") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Value") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricValue) MarshalJSON() ([]byte, error) {
	type NoMethod MetricValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MobileApplicationTargeting: Mobile application targeting settings.
type MobileApplicationTargeting struct {
	// FirstPartyTargeting: Publisher owned apps to be targeted or excluded
	// by the publisher to
	// display the ads in.
	FirstPartyTargeting *FirstPartyMobileApplicationTargeting `json:"firstPartyTargeting,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FirstPartyTargeting")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FirstPartyTargeting") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *MobileApplicationTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod MobileApplicationTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Nanos: Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and
	// `nanos`=-750,000,000.
	Nanos int64 `json:"nanos,omitempty"`

	// Units: The whole units of the amount.
	// For example if `currencyCode` is "USD", then 1 unit is one US
	// dollar.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CurrencyCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CurrencyCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Money) MarshalJSON() ([]byte, error) {
	type NoMethod Money
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NativeContent: Native content for a creative.
type NativeContent struct {
	// AdvertiserName: The name of the advertiser or sponsor, to be
	// displayed in the ad creative.
	AdvertiserName string `json:"advertiserName,omitempty"`

	// AppIcon: The app icon, for app download ads.
	AppIcon *Image `json:"appIcon,omitempty"`

	// Body: A long description of the ad.
	Body string `json:"body,omitempty"`

	// CallToAction: A label for the button that the user is supposed to
	// click.
	CallToAction string `json:"callToAction,omitempty"`

	// ClickLinkUrl: The URL that the browser/SDK will load when the user
	// clicks the ad.
	ClickLinkUrl string `json:"clickLinkUrl,omitempty"`

	// ClickTrackingUrl: The URL to use for click tracking.
	ClickTrackingUrl string `json:"clickTrackingUrl,omitempty"`

	// Headline: A short title for the ad.
	Headline string `json:"headline,omitempty"`

	// Image: A large image.
	Image *Image `json:"image,omitempty"`

	// Logo: A smaller image, for the advertiser's logo.
	Logo *Image `json:"logo,omitempty"`

	// PriceDisplayText: The price of the promoted app including currency
	// info.
	PriceDisplayText string `json:"priceDisplayText,omitempty"`

	// StarRating: The app rating in the app store. Must be in the range
	// [0-5].
	StarRating float64 `json:"starRating,omitempty"`

	// StoreUrl: The URL to the app store to purchase/download the promoted
	// app.
	StoreUrl string `json:"storeUrl,omitempty"`

	// VideoUrl: The URL to fetch a native video ad.
	VideoUrl string `json:"videoUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdvertiserName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdvertiserName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *NativeContent) MarshalJSON() ([]byte, error) {
	type NoMethod NativeContent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *NativeContent) UnmarshalJSON(data []byte) error {
	type NoMethod NativeContent
	var s1 struct {
		StarRating gensupport.JSONFloat64 `json:"starRating"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.StarRating = float64(s1.StarRating)
	return nil
}

// NonBillableWinningBidStatusRow: The number of winning bids with the
// specified dimension values for which the
// buyer was not billed, as described by the specified status.
type NonBillableWinningBidStatusRow struct {
	// BidCount: The number of bids with the specified status.
	BidCount *MetricValue `json:"bidCount,omitempty"`

	// RowDimensions: The values of all dimensions associated with metric
	// values in this row.
	RowDimensions *RowDimensions `json:"rowDimensions,omitempty"`

	// Status: The status specifying why the winning bids were not billed.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - A placeholder for an undefined status.
	// This value will never be returned in responses.
	//   "AD_NOT_RENDERED" - The buyer was not billed because the ad was not
	// rendered by the
	// publisher.
	//   "INVALID_IMPRESSION" - The buyer was not billed because the
	// impression won by the bid was
	// determined to be invalid.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BidCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BidCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonBillableWinningBidStatusRow) MarshalJSON() ([]byte, error) {
	type NoMethod NonBillableWinningBidStatusRow
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonGuaranteedAuctionTerms: Terms for Private Auctions. Note that
// Private Auctions can be created only
// by the seller, but they can be returned in a get or list request.
type NonGuaranteedAuctionTerms struct {
	// AutoOptimizePrivateAuction: True if open auction buyers are allowed
	// to compete with invited buyers
	// in this private auction.
	AutoOptimizePrivateAuction bool `json:"autoOptimizePrivateAuction,omitempty"`

	// ReservePricesPerBuyer: Reserve price for the specified buyer.
	ReservePricesPerBuyer []*PricePerBuyer `json:"reservePricesPerBuyer,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AutoOptimizePrivateAuction") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AutoOptimizePrivateAuction") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonGuaranteedAuctionTerms) MarshalJSON() ([]byte, error) {
	type NoMethod NonGuaranteedAuctionTerms
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonGuaranteedFixedPriceTerms: Terms for Preferred Deals. Note that
// Preferred Deals cannot be created via
// the API at this time, but can be returned in a get or list request.
type NonGuaranteedFixedPriceTerms struct {
	// FixedPrices: Fixed price for the specified buyer.
	FixedPrices []*PricePerBuyer `json:"fixedPrices,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FixedPrices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FixedPrices") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonGuaranteedFixedPriceTerms) MarshalJSON() ([]byte, error) {
	type NoMethod NonGuaranteedFixedPriceTerms
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Note: A proposal may be associated to several notes.
type Note struct {
	// CreateTime: The timestamp for when this note was created.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// CreatorRole: The role of the person (buyer/seller) creating the
	// note.
	// @OutputOnly
	//
	// Possible values:
	//   "BUYER_SELLER_ROLE_UNSPECIFIED" - A placeholder for an undefined
	// buyer/seller role.
	//   "BUYER" - Specifies the role as buyer.
	//   "SELLER" - Specifies the role as seller.
	CreatorRole string `json:"creatorRole,omitempty"`

	// Note: The actual note to attach.
	// (max-length: 1024 unicode code units)
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	Note string `json:"note,omitempty"`

	// NoteId: The unique ID for the note.
	// @OutputOnly
	NoteId string `json:"noteId,omitempty"`

	// ProposalRevision: The revision number of the proposal when the note
	// is created.
	// @OutputOnly
	ProposalRevision int64 `json:"proposalRevision,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Note) MarshalJSON() ([]byte, error) {
	type NoMethod Note
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperatingSystemTargeting: Represents targeting information for
// operating systems.
type OperatingSystemTargeting struct {
	// OperatingSystemCriteria: IDs of operating systems to be
	// included/excluded.
	OperatingSystemCriteria *CriteriaTargeting `json:"operatingSystemCriteria,omitempty"`

	// OperatingSystemVersionCriteria: IDs of operating system versions to
	// be included/excluded.
	OperatingSystemVersionCriteria *CriteriaTargeting `json:"operatingSystemVersionCriteria,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "OperatingSystemCriteria") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OperatingSystemCriteria")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *OperatingSystemTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod OperatingSystemTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PauseProposalRequest: Request message to pause serving for an
// already-finalized proposal.
type PauseProposalRequest struct {
	// Reason: The reason why the proposal is being paused.
	// This human readable message will be displayed in the seller's
	// UI.
	// (Max length: 100 unicode code units.)
	Reason string `json:"reason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Reason") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Reason") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PauseProposalRequest) MarshalJSON() ([]byte, error) {
	type NoMethod PauseProposalRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PlacementTargeting: Represents targeting about where the ads can
// appear, e.g., certain sites or
// mobile applications.
// Different placement targeting types will be logically OR'ed.
type PlacementTargeting struct {
	// MobileApplicationTargeting: Mobile application targeting information
	// in a deal.
	// This doesn't apply to Auction Packages.
	MobileApplicationTargeting *MobileApplicationTargeting `json:"mobileApplicationTargeting,omitempty"`

	// UrlTargeting: URLs to be included/excluded.
	UrlTargeting *UrlTargeting `json:"urlTargeting,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "MobileApplicationTargeting") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "MobileApplicationTargeting") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PlacementTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod PlacementTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PlatformContext: @OutputOnly The type of platform the restriction
// applies to.
type PlatformContext struct {
	// Platforms: The platforms this restriction applies to.
	//
	// Possible values:
	//   "DESKTOP" - Desktop platform.
	//   "ANDROID" - Android platform.
	//   "IOS" - iOS platform.
	Platforms []string `json:"platforms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Platforms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Platforms") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PlatformContext) MarshalJSON() ([]byte, error) {
	type NoMethod PlatformContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Price: Represents a price and a pricing type for a product / deal.
type Price struct {
	// Amount: The actual price with currency specified.
	Amount *Money `json:"amount,omitempty"`

	// PricingType: The pricing type for the deal/product. (default: CPM)
	//
	// Possible values:
	//   "PRICING_TYPE_UNSPECIFIED" - A placeholder for an undefined pricing
	// type. If the pricing type is
	// unpsecified, `COST_PER_MILLE` will be used instead.
	//   "COST_PER_MILLE" - Cost per thousand impressions.
	//   "COST_PER_DAY" - Cost per day
	PricingType string `json:"pricingType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Amount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Amount") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Price) MarshalJSON() ([]byte, error) {
	type NoMethod Price
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PricePerBuyer: Used to specify pricing rules for buyers/advertisers.
// Each PricePerBuyer in
// a product can become 0 or 1 deals. To check if there is a
// PricePerBuyer for
// a particular buyer or buyer/advertiser pair, we look for the most
// specific
// matching rule - we first look for a rule matching the buyer and
// advertiser,
// next a rule with the buyer but an empty advertiser list, and
// otherwise look
// for a matching rule where no buyer is set.
type PricePerBuyer struct {
	// AdvertiserIds: The list of advertisers for this price when associated
	// with this buyer.
	// If empty, all advertisers with this buyer pay this price.
	AdvertiserIds []string `json:"advertiserIds,omitempty"`

	// Buyer: The buyer who will pay this price. If unset, all buyers can
	// pay this price
	// (if the
	// advertisers match, and there's no more specific rule matching the
	// buyer).
	Buyer *Buyer `json:"buyer,omitempty"`

	// Price: The specified price.
	Price *Price `json:"price,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdvertiserIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdvertiserIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PricePerBuyer) MarshalJSON() ([]byte, error) {
	type NoMethod PricePerBuyer
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PrivateData: Buyers are allowed to store certain types of private
// data in a proposal/deal.
type PrivateData struct {
	// ReferenceId: A buyer or seller specified reference ID. This can be
	// queried in the list
	// operations (max-length: 1024 unicode code units).
	ReferenceId string `json:"referenceId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReferenceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ReferenceId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PrivateData) MarshalJSON() ([]byte, error) {
	type NoMethod PrivateData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Product: Note: this resource requires whitelisting for access. Please
// contact your
// account manager for access to Marketplace resources.
//
// A product is a segment of inventory that a seller wishes to sell. It
// is
// associated with certain terms and targeting information which helps
// the buyer
// know more about the inventory.
type Product struct {
	// AvailableEndTime: The proposed end time for the deal. The field will
	// be truncated to the order of
	// seconds during serving.
	AvailableEndTime string `json:"availableEndTime,omitempty"`

	// AvailableStartTime: Inventory availability dates. The start time will
	// be truncated to seconds during serving.
	// Thus, a field specified as 3:23:34.456 (HH:mm:ss.SSS) will be
	// truncated to 3:23:34
	// when serving.
	AvailableStartTime string `json:"availableStartTime,omitempty"`

	// CreateTime: Creation time.
	CreateTime string `json:"createTime,omitempty"`

	// CreatorContacts: Optional contact information for the creator of this
	// product.
	CreatorContacts []*ContactInformation `json:"creatorContacts,omitempty"`

	// DisplayName: The display name for this product as set by the seller.
	DisplayName string `json:"displayName,omitempty"`

	// HasCreatorSignedOff: If the creator has already signed off on the
	// product, then the buyer can
	// finalize the deal by accepting the product as is. When copying to
	// a
	// proposal, if any of the terms are changed, then auto_finalize
	// is
	// automatically set to false.
	HasCreatorSignedOff bool `json:"hasCreatorSignedOff,omitempty"`

	// ProductId: The unique ID for the product.
	ProductId string `json:"productId,omitempty"`

	// ProductRevision: The revision number of the product (auto-assigned by
	// Marketplace).
	ProductRevision int64 `json:"productRevision,omitempty,string"`

	// PublisherProfileId: An ID which can be used by the Publisher Profile
	// API to get more
	// information about the seller that created this product.
	PublisherProfileId string `json:"publisherProfileId,omitempty"`

	// Seller: Information about the seller that created this product.
	Seller *Seller `json:"seller,omitempty"`

	// SyndicationProduct: The syndication product associated with the deal.
	//
	// Possible values:
	//   "SYNDICATION_PRODUCT_UNSPECIFIED" - A placeholder for an undefined
	// syndication product.
	//   "CONTENT" - This typically represents a web page.
	//   "MOBILE" - This represents a mobile property.
	//   "VIDEO" - This represents video ad formats.
	//   "GAMES" - This represents ads shown within games.
	SyndicationProduct string `json:"syndicationProduct,omitempty"`

	// TargetingCriterion: Targeting that is shared between the buyer and
	// the seller. Each targeting
	// criterion has a specified key and for each key there is a list of
	// inclusion
	// value or exclusion values.
	TargetingCriterion []*TargetingCriteria `json:"targetingCriterion,omitempty"`

	// Terms: The negotiable terms of the deal.
	Terms *DealTerms `json:"terms,omitempty"`

	// UpdateTime: Time of last update.
	UpdateTime string `json:"updateTime,omitempty"`

	// WebPropertyCode: The web-property code for the seller. This needs to
	// be copied as is when
	// adding a new deal to a proposal.
	WebPropertyCode string `json:"webPropertyCode,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AvailableEndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailableEndTime") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Product) MarshalJSON() ([]byte, error) {
	type NoMethod Product
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Proposal: Note: this resource requires whitelisting for access.
// Please contact your
// account manager for access to Marketplace resources.
//
// Represents a proposal in the Marketplace. A proposal is the unit
// of
// negotiation between a seller and a buyer and contains deals which
// are served.
//
// Note: you can not update, create, or otherwise modify Private
// Auction or Preferred Deals deals through the API.
//
// Fields are updatable unless noted otherwise.
type Proposal struct {
	// BilledBuyer: Reference to the buyer that will get billed for this
	// proposal.
	// @OutputOnly
	BilledBuyer *Buyer `json:"billedBuyer,omitempty"`

	// Buyer: Reference to the buyer on the proposal.
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	Buyer *Buyer `json:"buyer,omitempty"`

	// BuyerContacts: Contact information for the buyer.
	BuyerContacts []*ContactInformation `json:"buyerContacts,omitempty"`

	// BuyerPrivateData: Private data for buyer. (hidden from seller).
	BuyerPrivateData *PrivateData `json:"buyerPrivateData,omitempty"`

	// Deals: The deals associated with this proposal. For Private Auction
	// proposals (whose deals have
	// NonGuaranteedAuctionTerms), there will only be one deal.
	Deals []*Deal `json:"deals,omitempty"`

	// DisplayName: The name for the proposal.
	DisplayName string `json:"displayName,omitempty"`

	// IsRenegotiating: True if the proposal is being
	// renegotiated.
	// @OutputOnly
	IsRenegotiating bool `json:"isRenegotiating,omitempty"`

	// IsSetupComplete: True, if the buyside inventory setup is complete for
	// this proposal.
	// @OutputOnly
	IsSetupComplete bool `json:"isSetupComplete,omitempty"`

	// LastUpdaterOrCommentorRole: The role of the last user that either
	// updated the proposal or left a
	// comment.
	// @OutputOnly
	//
	// Possible values:
	//   "BUYER_SELLER_ROLE_UNSPECIFIED" - A placeholder for an undefined
	// buyer/seller role.
	//   "BUYER" - Specifies the role as buyer.
	//   "SELLER" - Specifies the role as seller.
	LastUpdaterOrCommentorRole string `json:"lastUpdaterOrCommentorRole,omitempty"`

	// Notes: The notes associated with this proposal.
	// @OutputOnly
	Notes []*Note `json:"notes,omitempty"`

	// OriginatorRole: Indicates whether the buyer/seller created the
	// proposal.
	// @OutputOnly
	//
	// Possible values:
	//   "BUYER_SELLER_ROLE_UNSPECIFIED" - A placeholder for an undefined
	// buyer/seller role.
	//   "BUYER" - Specifies the role as buyer.
	//   "SELLER" - Specifies the role as seller.
	OriginatorRole string `json:"originatorRole,omitempty"`

	// PrivateAuctionId: Private auction ID if this proposal is a private
	// auction proposal.
	// @OutputOnly
	PrivateAuctionId string `json:"privateAuctionId,omitempty"`

	// ProposalId: The unique ID of the proposal.
	// @OutputOnly
	ProposalId string `json:"proposalId,omitempty"`

	// ProposalRevision: The revision number for the proposal.
	// Each update to the proposal or the deal causes the proposal revision
	// number
	// to auto-increment. The buyer keeps track of the last revision number
	// they
	// know of and pass it in when making an update. If the head revision
	// number
	// on the server has since incremented, then an ABORTED error is
	// returned
	// during the update operation to let the buyer know that a subsequent
	// update
	// was made.
	// @OutputOnly
	ProposalRevision int64 `json:"proposalRevision,omitempty,string"`

	// ProposalState: The current state of the proposal.
	// @OutputOnly
	//
	// Possible values:
	//   "PROPOSAL_STATE_UNSPECIFIED" - A placeholder for an undefined
	// proposal state.
	//   "PROPOSED" - The proposal is under negotiation or renegotiation.
	//   "BUYER_ACCEPTED" - The proposal has been accepted by the buyer.
	//   "SELLER_ACCEPTED" - The proposal has been accepted by the seller.
	//   "CANCELED" - The negotiations on the proposal were canceled and the
	// proposal was never
	// finalized.
	//   "FINALIZED" - The proposal is finalized. During renegotiation, the
	// proposal may
	// not be in this state.
	ProposalState string `json:"proposalState,omitempty"`

	// Seller: Reference to the seller on the proposal.
	//
	// Note: This field may be set only when creating the resource.
	// Modifying
	// this field while updating the resource will result in an error.
	Seller *Seller `json:"seller,omitempty"`

	// SellerContacts: Contact information for the seller.
	// @OutputOnly
	SellerContacts []*ContactInformation `json:"sellerContacts,omitempty"`

	// UpdateTime: The time when the proposal was last revised.
	// @OutputOnly
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BilledBuyer") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BilledBuyer") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Proposal) MarshalJSON() ([]byte, error) {
	type NoMethod Proposal
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PublisherProfile: Note: this resource requires whitelisting for
// access. Please contact your
// account manager for access to Marketplace resources.
//
// Represents a publisher profile in Marketplace.
//
// All fields are read only. All string fields are free-form text
// entered by the publisher
// unless noted otherwise.
type PublisherProfile struct {
	// AudienceDescription: Description on the publisher's audience.
	AudienceDescription string `json:"audienceDescription,omitempty"`

	// BuyerPitchStatement: Statement explaining what's unique about
	// publisher's business, and why buyers should
	// partner with the publisher.
	BuyerPitchStatement string `json:"buyerPitchStatement,omitempty"`

	// DirectDealsContact: Contact information for direct reservation deals.
	// This is free text entered by the publisher
	// and may include information like names, phone numbers and email
	// addresses.
	DirectDealsContact string `json:"directDealsContact,omitempty"`

	// DisplayName: Name of the publisher profile.
	DisplayName string `json:"displayName,omitempty"`

	// Domains: The list of domains represented in this publisher profile.
	// Empty if this is a parent profile.
	// These are top private domains, meaning that these will not contain a
	// string like
	// "photos.google.co.uk/123", but will instead contain "google.co.uk".
	Domains []string `json:"domains,omitempty"`

	// GooglePlusUrl: URL to publisher's Google+ page.
	GooglePlusUrl string `json:"googlePlusUrl,omitempty"`

	// LogoUrl: A Google public URL to the logo for this publisher profile.
	// The logo is stored as
	// a PNG, JPG, or GIF image.
	LogoUrl string `json:"logoUrl,omitempty"`

	// MediaKitUrl: URL to additional marketing and sales materials.
	MediaKitUrl string `json:"mediaKitUrl,omitempty"`

	// Overview: Overview of the publisher.
	Overview string `json:"overview,omitempty"`

	// ProgrammaticDealsContact: Contact information for programmatic deals.
	// This is free text entered by the publisher
	// and may include information like names, phone numbers and email
	// addresses.
	ProgrammaticDealsContact string `json:"programmaticDealsContact,omitempty"`

	// PublisherProfileId: Unique ID for publisher profile.
	PublisherProfileId string `json:"publisherProfileId,omitempty"`

	// RateCardInfoUrl: URL to a publisher rate card.
	RateCardInfoUrl string `json:"rateCardInfoUrl,omitempty"`

	// SamplePageUrl: URL to a sample content page.
	SamplePageUrl string `json:"samplePageUrl,omitempty"`

	// Seller: Seller of the publisher profile.
	Seller *Seller `json:"seller,omitempty"`

	// TopHeadlines: Up to three key metrics and rankings. Max 100
	// characters each.
	// For example "#1 Mobile News Site for 20 Straight Months".
	TopHeadlines []string `json:"topHeadlines,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AudienceDescription")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudienceDescription") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PublisherProfile) MarshalJSON() ([]byte, error) {
	type NoMethod PublisherProfile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RealtimeTimeRange: An open-ended realtime time range specified by the
// start timestamp.
// For filter sets that specify a realtime time range RTB metrics
// continue to
// be aggregated throughout the lifetime of the filter set.
type RealtimeTimeRange struct {
	// StartTimestamp: The start timestamp of the real-time RTB metrics
	// aggregation.
	StartTimestamp string `json:"startTimestamp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StartTimestamp") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StartTimestamp") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RealtimeTimeRange) MarshalJSON() ([]byte, error) {
	type NoMethod RealtimeTimeRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Reason: A specific filtering status and how many times it occurred.
type Reason struct {
	// Count: The number of times the creative was filtered for the status.
	// The
	// count is aggregated across all publishers on the exchange.
	Count int64 `json:"count,omitempty,string"`

	// Status: The filtering status code. Please refer to
	// the
	// [creative-status-codes.txt](https://storage.googleapis.com/adx-rtb
	// -dictionaries/creative-status-codes.txt)
	// file for different statuses.
	Status int64 `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Reason) MarshalJSON() ([]byte, error) {
	type NoMethod Reason
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RelativeDateRange: A relative date range, specified by an offset and
// a duration.
// The supported range of dates begins 30 days before today and ends
// today,
// i.e., the limits for these values are:
// offset_days >= 0
// duration_days >= 1
// offset_days + duration_days <= 30
type RelativeDateRange struct {
	// DurationDays: The number of days in the requested date range, e.g.,
	// for a range spanning
	// today: 1. For a range spanning the last 7 days: 7.
	DurationDays int64 `json:"durationDays,omitempty"`

	// OffsetDays: The end date of the filter set, specified as the number
	// of days before
	// today, e.g., for a range where the last date is today: 0.
	OffsetDays int64 `json:"offsetDays,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DurationDays") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DurationDays") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RelativeDateRange) MarshalJSON() ([]byte, error) {
	type NoMethod RelativeDateRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RemoveDealAssociationRequest: A request for removing the association
// between a deal and a creative.
type RemoveDealAssociationRequest struct {
	// Association: The association between a creative and a deal that
	// should be removed.
	Association *CreativeDealAssociation `json:"association,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Association") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Association") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RemoveDealAssociationRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RemoveDealAssociationRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResumeProposalRequest: Request message to resume (unpause) serving
// for an already-finalized
// proposal.
type ResumeProposalRequest struct {
}

// RowDimensions: A response may include multiple rows, breaking down
// along various dimensions.
// Encapsulates the values of all dimensions for a given row.
type RowDimensions struct {
	// PublisherIdentifier: The publisher identifier for this row, if a
	// breakdown by
	// BreakdownDimension.PUBLISHER_IDENTIFIER
	// was requested.
	PublisherIdentifier string `json:"publisherIdentifier,omitempty"`

	// TimeInterval: The time interval that this row represents.
	TimeInterval *TimeInterval `json:"timeInterval,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PublisherIdentifier")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PublisherIdentifier") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RowDimensions) MarshalJSON() ([]byte, error) {
	type NoMethod RowDimensions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SecurityContext: @OutputOnly A security context.
type SecurityContext struct {
	// Securities: The security types in this context.
	//
	// Possible values:
	//   "INSECURE" - Matches impressions that require insecure
	// compatibility.
	//   "SSL" - Matches impressions that require SSL compatibility.
	Securities []string `json:"securities,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Securities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Securities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SecurityContext) MarshalJSON() ([]byte, error) {
	type NoMethod SecurityContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Seller: Represents a seller of inventory. Each seller is identified
// by a unique
// Ad Manager account ID.
type Seller struct {
	// AccountId: The unique ID for the seller. The seller fills in this
	// field.
	// The seller account ID is then available to buyer in the product.
	AccountId string `json:"accountId,omitempty"`

	// SubAccountId: Optional sub-account ID for the seller.
	SubAccountId string `json:"subAccountId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccountId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Seller) MarshalJSON() ([]byte, error) {
	type NoMethod Seller
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServingContext: The serving context for this restriction.
type ServingContext struct {
	// All: Matches all contexts.
	//
	// Possible values:
	//   "SIMPLE_CONTEXT" - A simple context.
	All string `json:"all,omitempty"`

	// AppType: Matches impressions for a particular app type.
	AppType *AppContext `json:"appType,omitempty"`

	// AuctionType: Matches impressions for a particular auction type.
	AuctionType *AuctionContext `json:"auctionType,omitempty"`

	// Location: Matches impressions coming from users *or* publishers in a
	// specific
	// location.
	Location *LocationContext `json:"location,omitempty"`

	// Platform: Matches impressions coming from a particular platform.
	Platform *PlatformContext `json:"platform,omitempty"`

	// SecurityType: Matches impressions for a particular security type.
	SecurityType *SecurityContext `json:"securityType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "All") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "All") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServingContext) MarshalJSON() ([]byte, error) {
	type NoMethod ServingContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServingRestriction: @OutputOnly A representation of the status of an
// ad in a
// specific context. A context here relates to where something
// ultimately serves
// (for example, a user or publisher geo, a platform, an HTTPS vs HTTP
// request,
// or the type of auction).
type ServingRestriction struct {
	// Contexts: The contexts for the restriction.
	Contexts []*ServingContext `json:"contexts,omitempty"`

	// Disapproval: Disapproval bound to this restriction.
	// Only present if status=DISAPPROVED.
	// Can be used to filter the response of the
	// creatives.list
	// method.
	Disapproval *Disapproval `json:"disapproval,omitempty"`

	// DisapprovalReasons: Any disapprovals bound to this restriction.
	// Only present if status=DISAPPROVED.
	// Can be used to filter the response of
	// the
	// creatives.list
	// method.
	// Deprecated; please use
	// disapproval
	// field instead.
	DisapprovalReasons []*Disapproval `json:"disapprovalReasons,omitempty"`

	// Status: The status of the creative in this context (for example, it
	// has been
	// explicitly disapproved or is pending review).
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - The status is not known.
	//   "DISAPPROVAL" - The ad was disapproved in this context.
	//   "PENDING_REVIEW" - The ad is pending review in this context.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServingRestriction) MarshalJSON() ([]byte, error) {
	type NoMethod ServingRestriction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Size: Message depicting the size of the creative. The units of width
// and
// height depend on the type of the targeting.
type Size struct {
	// Height: The height of the creative.
	Height int64 `json:"height,omitempty"`

	// Width: The width of the creative
	Width int64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Height") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Height") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Size) MarshalJSON() ([]byte, error) {
	type NoMethod Size
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StopWatchingCreativeRequest: A request for stopping notifications for
// changes to creative Status.
type StopWatchingCreativeRequest struct {
}

// TargetingCriteria: Advertisers can target different attributes of an
// ad slot. For example,
// they can choose to show ads only if the user is in the U.S.
// Such
// targeting criteria can be specified as part of Shared Targeting.
type TargetingCriteria struct {
	// Exclusions: The list of values to exclude from targeting. Each value
	// is AND'd
	// together.
	Exclusions []*TargetingValue `json:"exclusions,omitempty"`

	// Inclusions: The list of value to include as part of the targeting.
	// Each value is OR'd
	// together.
	Inclusions []*TargetingValue `json:"inclusions,omitempty"`

	// Key: The key representing the shared targeting criterion.
	// Targeting criteria defined by Google ad servers will begin with
	// GOOG_.
	// Third parties may define their own keys.
	// A list of permissible keys along with the acceptable values will
	// be
	// provided as part of the external documentation.
	Key string `json:"key,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Exclusions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Exclusions") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TargetingCriteria) MarshalJSON() ([]byte, error) {
	type NoMethod TargetingCriteria
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TargetingValue: A polymorphic targeting value used as part of Shared
// Targeting.
type TargetingValue struct {
	// CreativeSizeValue: The creative size value to include/exclude.
	// Filled in when key = GOOG_CREATIVE_SIZE
	CreativeSizeValue *CreativeSize `json:"creativeSizeValue,omitempty"`

	// DayPartTargetingValue: The daypart targeting to include /
	// exclude.
	// Filled in when the key is GOOG_DAYPART_TARGETING.
	// The definition of this targeting is derived from the structure
	// used by Ad Manager.
	DayPartTargetingValue *DayPartTargeting `json:"dayPartTargetingValue,omitempty"`

	// LongValue: The long value to include/exclude.
	LongValue int64 `json:"longValue,omitempty,string"`

	// StringValue: The string value to include/exclude.
	StringValue string `json:"stringValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreativeSizeValue")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreativeSizeValue") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TargetingValue) MarshalJSON() ([]byte, error) {
	type NoMethod TargetingValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TechnologyTargeting: Represents targeting about various types of
// technology.
type TechnologyTargeting struct {
	// DeviceCapabilityTargeting: IDs of device capabilities to be
	// included/excluded.
	DeviceCapabilityTargeting *CriteriaTargeting `json:"deviceCapabilityTargeting,omitempty"`

	// DeviceCategoryTargeting: IDs of device categories to be
	// included/excluded.
	DeviceCategoryTargeting *CriteriaTargeting `json:"deviceCategoryTargeting,omitempty"`

	// OperatingSystemTargeting: Operating system related targeting
	// information.
	OperatingSystemTargeting *OperatingSystemTargeting `json:"operatingSystemTargeting,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DeviceCapabilityTargeting") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "DeviceCapabilityTargeting") to include in API requests with the JSON
	// null value. By default, fields with empty values are omitted from API
	// requests. However, any field with an empty value appearing in
	// NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TechnologyTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod TechnologyTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeInterval: An interval of time, with an absolute start and end.
type TimeInterval struct {
	// EndTime: The timestamp marking the end of the range (exclusive) for
	// which data is
	// included.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: The timestamp marking the start of the range (inclusive)
	// for which data is
	// included.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeInterval) MarshalJSON() ([]byte, error) {
	type NoMethod TimeInterval
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeOfDay: Represents a time of day. The date and time zone are
// either not significant
// or are specified elsewhere. An API may choose to allow leap seconds.
// Related
// types are google.type.Date and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours: Hours of day in 24 hour format. Should be from 0 to 23. An API
	// may choose
	// to allow the value "24:00:00" for scenarios like business closing
	// time.
	Hours int64 `json:"hours,omitempty"`

	// Minutes: Minutes of hour of day. Must be from 0 to 59.
	Minutes int64 `json:"minutes,omitempty"`

	// Nanos: Fractions of seconds in nanoseconds. Must be from 0 to
	// 999,999,999.
	Nanos int64 `json:"nanos,omitempty"`

	// Seconds: Seconds of minutes of the time. Must normally be from 0 to
	// 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int64 `json:"seconds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Hours") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Hours") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeOfDay) MarshalJSON() ([]byte, error) {
	type NoMethod TimeOfDay
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UrlTargeting: Represents a list of targeted and excluded URLs (e.g.,
// google.com).
// For Private Auction and AdX Preferred Deals, URLs are either included
// or
// excluded.
// For Programmatic Guaranteed and Preferred Deals, this doesn't
// apply.
type UrlTargeting struct {
	// ExcludedUrls: A list of URLs to be excluded.
	ExcludedUrls []string `json:"excludedUrls,omitempty"`

	// TargetedUrls: A list of URLs to be included.
	TargetedUrls []string `json:"targetedUrls,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExcludedUrls") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludedUrls") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UrlTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod UrlTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VideoContent: Video content for a creative.
type VideoContent struct {
	// VideoUrl: The URL to fetch a video ad.
	VideoUrl string `json:"videoUrl,omitempty"`

	// VideoVastXml: The contents of a VAST document for a video ad.
	// This document should conform to the VAST 2.0 or 3.0 standard.
	VideoVastXml string `json:"videoVastXml,omitempty"`

	// ForceSendFields is a list of field names (e.g. "VideoUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "VideoUrl") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VideoContent) MarshalJSON() ([]byte, error) {
	type NoMethod VideoContent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VideoTargeting: Represents targeting information about video.
type VideoTargeting struct {
	// ExcludedPositionTypes: A list of video positions to be
	// excluded.
	// Position types can either be included or excluded (XOR).
	//
	// Possible values:
	//   "POSITION_TYPE_UNSPECIFIED" - A placeholder for an undefined video
	// position.
	//   "PREROLL" - Ad is played before the video.
	//   "MIDROLL" - Ad is played during the video.
	//   "POSTROLL" - Ad is played after the video.
	ExcludedPositionTypes []string `json:"excludedPositionTypes,omitempty"`

	// TargetedPositionTypes: A list of video positions to be included.
	// When the included list is present, the excluded list must be
	// empty.
	// When the excluded list is present, the included list must be empty.
	//
	// Possible values:
	//   "POSITION_TYPE_UNSPECIFIED" - A placeholder for an undefined video
	// position.
	//   "PREROLL" - Ad is played before the video.
	//   "MIDROLL" - Ad is played during the video.
	//   "POSTROLL" - Ad is played after the video.
	TargetedPositionTypes []string `json:"targetedPositionTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExcludedPositionTypes") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExcludedPositionTypes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VideoTargeting) MarshalJSON() ([]byte, error) {
	type NoMethod VideoTargeting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WatchCreativeRequest: A request for watching changes to creative
// Status.
type WatchCreativeRequest struct {
	// Topic: The Pub/Sub topic to publish notifications to.
	// This topic must already exist and must give permission
	// to
	// ad-exchange-buyside-reports@google.com to write to the topic.
	// This should be the full resource name
	// in
	// "projects/{project_id}/topics/{topic_id}" format.
	Topic string `json:"topic,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Topic") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Topic") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WatchCreativeRequest) MarshalJSON() ([]byte, error) {
	type NoMethod WatchCreativeRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "adexchangebuyer2.accounts.clients.create":

type AccountsClientsCreateCall struct {
	s          *Service
	accountId  int64
	client     *Client
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new client buyer.
func (r *AccountsClientsService) Create(accountId int64, client *Client) *AccountsClientsCreateCall {
	c := &AccountsClientsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.client = client
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsCreateCall) Fields(s ...googleapi.Field) *AccountsClientsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsCreateCall) Context(ctx context.Context) *AccountsClientsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.client)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.create" call.
// Exactly one of *Client or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Client.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsClientsCreateCall) Do(opts ...googleapi.CallOption) (*Client, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Client{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new client buyer.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.clients.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Unique numerical account ID for the buyer of which the client buyer\nis a customer; the sponsor buyer to create a client for. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients",
	//   "request": {
	//     "$ref": "Client"
	//   },
	//   "response": {
	//     "$ref": "Client"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.get":

type AccountsClientsGetCall struct {
	s               *Service
	accountId       int64
	clientAccountId int64
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// Get: Gets a client buyer with a given client account ID.
func (r *AccountsClientsService) Get(accountId int64, clientAccountId int64) *AccountsClientsGetCall {
	c := &AccountsClientsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsGetCall) Fields(s ...googleapi.Field) *AccountsClientsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsGetCall) IfNoneMatch(entityTag string) *AccountsClientsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsGetCall) Context(ctx context.Context) *AccountsClientsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.get" call.
// Exactly one of *Client or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Client.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsClientsGetCall) Do(opts ...googleapi.CallOption) (*Client, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Client{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a client buyer with a given client account ID.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer to retrieve. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}",
	//   "response": {
	//     "$ref": "Client"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.list":

type AccountsClientsListCall struct {
	s            *Service
	accountId    int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all the clients for the current sponsor buyer.
func (r *AccountsClientsService) List(accountId int64) *AccountsClientsListCall {
	c := &AccountsClientsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer clients than requested.
// If unspecified, the server will pick an appropriate default.
func (c *AccountsClientsListCall) PageSize(pageSize int64) *AccountsClientsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListClientsResponse.nextPageToken
// returned from the previous call to the
// accounts.clients.list
// method.
func (c *AccountsClientsListCall) PageToken(pageToken string) *AccountsClientsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// PartnerClientId sets the optional parameter "partnerClientId":
// Optional unique identifier (from the standpoint of an Ad Exchange
// sponsor
// buyer partner) of the client to return.
// If specified, at most one client will be returned in the response.
func (c *AccountsClientsListCall) PartnerClientId(partnerClientId string) *AccountsClientsListCall {
	c.urlParams_.Set("partnerClientId", partnerClientId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsListCall) Fields(s ...googleapi.Field) *AccountsClientsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsListCall) IfNoneMatch(entityTag string) *AccountsClientsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsListCall) Context(ctx context.Context) *AccountsClientsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": strconv.FormatInt(c.accountId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.list" call.
// Exactly one of *ListClientsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListClientsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsClientsListCall) Do(opts ...googleapi.CallOption) (*ListClientsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListClientsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the clients for the current sponsor buyer.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Unique numerical account ID of the sponsor buyer to list the clients for.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer clients than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListClientsResponse.nextPageToken\nreturned from the previous call to the\naccounts.clients.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "partnerClientId": {
	//       "description": "Optional unique identifier (from the standpoint of an Ad Exchange sponsor\nbuyer partner) of the client to return.\nIf specified, at most one client will be returned in the response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients",
	//   "response": {
	//     "$ref": "ListClientsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsClientsListCall) Pages(ctx context.Context, f func(*ListClientsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.clients.update":

type AccountsClientsUpdateCall struct {
	s               *Service
	accountId       int64
	clientAccountId int64
	client          *Client
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Update: Updates an existing client buyer.
func (r *AccountsClientsService) Update(accountId int64, clientAccountId int64, client *Client) *AccountsClientsUpdateCall {
	c := &AccountsClientsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	c.client = client
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsUpdateCall) Fields(s ...googleapi.Field) *AccountsClientsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsUpdateCall) Context(ctx context.Context) *AccountsClientsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.client)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.update" call.
// Exactly one of *Client or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Client.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsClientsUpdateCall) Do(opts ...googleapi.CallOption) (*Client, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Client{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing client buyer.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer2.accounts.clients.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Unique numerical account ID for the buyer of which the client buyer\nis a customer; the sponsor buyer to update a client for. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Unique numerical account ID of the client to update. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}",
	//   "request": {
	//     "$ref": "Client"
	//   },
	//   "response": {
	//     "$ref": "Client"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.invitations.create":

type AccountsClientsInvitationsCreateCall struct {
	s                    *Service
	accountId            int64
	clientAccountId      int64
	clientuserinvitation *ClientUserInvitation
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: Creates and sends out an email invitation to access
// an Ad Exchange client buyer account.
func (r *AccountsClientsInvitationsService) Create(accountId int64, clientAccountId int64, clientuserinvitation *ClientUserInvitation) *AccountsClientsInvitationsCreateCall {
	c := &AccountsClientsInvitationsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	c.clientuserinvitation = clientuserinvitation
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsInvitationsCreateCall) Fields(s ...googleapi.Field) *AccountsClientsInvitationsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsInvitationsCreateCall) Context(ctx context.Context) *AccountsClientsInvitationsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsInvitationsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsInvitationsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.clientuserinvitation)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.invitations.create" call.
// Exactly one of *ClientUserInvitation or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClientUserInvitation.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsClientsInvitationsCreateCall) Do(opts ...googleapi.CallOption) (*ClientUserInvitation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientUserInvitation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates and sends out an email invitation to access\nan Ad Exchange client buyer account.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.clients.invitations.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer that the user\nshould be associated with. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations",
	//   "request": {
	//     "$ref": "ClientUserInvitation"
	//   },
	//   "response": {
	//     "$ref": "ClientUserInvitation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.invitations.get":

type AccountsClientsInvitationsGetCall struct {
	s               *Service
	accountId       int64
	clientAccountId int64
	invitationId    int64
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// Get: Retrieves an existing client user invitation.
func (r *AccountsClientsInvitationsService) Get(accountId int64, clientAccountId int64, invitationId int64) *AccountsClientsInvitationsGetCall {
	c := &AccountsClientsInvitationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	c.invitationId = invitationId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsInvitationsGetCall) Fields(s ...googleapi.Field) *AccountsClientsInvitationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsInvitationsGetCall) IfNoneMatch(entityTag string) *AccountsClientsInvitationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsInvitationsGetCall) Context(ctx context.Context) *AccountsClientsInvitationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsInvitationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsInvitationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations/{invitationId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
		"invitationId":    strconv.FormatInt(c.invitationId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.invitations.get" call.
// Exactly one of *ClientUserInvitation or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClientUserInvitation.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsClientsInvitationsGetCall) Do(opts ...googleapi.CallOption) (*ClientUserInvitation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientUserInvitation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves an existing client user invitation.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations/{invitationId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.invitations.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId",
	//     "invitationId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer that the user invitation\nto be retrieved is associated with. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "invitationId": {
	//       "description": "Numerical identifier of the user invitation to retrieve. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations/{invitationId}",
	//   "response": {
	//     "$ref": "ClientUserInvitation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.invitations.list":

type AccountsClientsInvitationsListCall struct {
	s               *Service
	accountId       int64
	clientAccountId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// List: Lists all the client users invitations for a client
// with a given account ID.
func (r *AccountsClientsInvitationsService) List(accountId int64, clientAccountId string) *AccountsClientsInvitationsListCall {
	c := &AccountsClientsInvitationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Server may return fewer clients than requested.
// If unspecified, server will pick an appropriate default.
func (c *AccountsClientsInvitationsListCall) PageSize(pageSize int64) *AccountsClientsInvitationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListClientUserInvitationsResponse.nextPageToken
// returned from the previous call to
// the
// clients.invitations.list
// method.
func (c *AccountsClientsInvitationsListCall) PageToken(pageToken string) *AccountsClientsInvitationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsInvitationsListCall) Fields(s ...googleapi.Field) *AccountsClientsInvitationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsInvitationsListCall) IfNoneMatch(entityTag string) *AccountsClientsInvitationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsInvitationsListCall) Context(ctx context.Context) *AccountsClientsInvitationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsInvitationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsInvitationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": c.clientAccountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.invitations.list" call.
// Exactly one of *ListClientUserInvitationsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListClientUserInvitationsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AccountsClientsInvitationsListCall) Do(opts ...googleapi.CallOption) (*ListClientUserInvitationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListClientUserInvitationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the client users invitations for a client\nwith a given account ID.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.invitations.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer to list invitations for.\n(required)\nYou must either specify a string representation of a\nnumerical account identifier or the `-` character\nto list all the invitations for all the clients\nof a given sponsor buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Server may return fewer clients than requested.\nIf unspecified, server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListClientUserInvitationsResponse.nextPageToken\nreturned from the previous call to the\nclients.invitations.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/invitations",
	//   "response": {
	//     "$ref": "ListClientUserInvitationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsClientsInvitationsListCall) Pages(ctx context.Context, f func(*ListClientUserInvitationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.clients.users.get":

type AccountsClientsUsersGetCall struct {
	s               *Service
	accountId       int64
	clientAccountId int64
	userId          int64
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// Get: Retrieves an existing client user.
func (r *AccountsClientsUsersService) Get(accountId int64, clientAccountId int64, userId int64) *AccountsClientsUsersGetCall {
	c := &AccountsClientsUsersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsUsersGetCall) Fields(s ...googleapi.Field) *AccountsClientsUsersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsUsersGetCall) IfNoneMatch(entityTag string) *AccountsClientsUsersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsUsersGetCall) Context(ctx context.Context) *AccountsClientsUsersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsUsersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsUsersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
		"userId":          strconv.FormatInt(c.userId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.users.get" call.
// Exactly one of *ClientUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ClientUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsClientsUsersGetCall) Do(opts ...googleapi.CallOption) (*ClientUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves an existing client user.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.users.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer\nthat the user to be retrieved is associated with. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Numerical identifier of the user to retrieve. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}",
	//   "response": {
	//     "$ref": "ClientUser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.clients.users.list":

type AccountsClientsUsersListCall struct {
	s               *Service
	accountId       int64
	clientAccountId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// List: Lists all the known client users for a specified
// sponsor buyer account ID.
func (r *AccountsClientsUsersService) List(accountId int64, clientAccountId string) *AccountsClientsUsersListCall {
	c := &AccountsClientsUsersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer clients than requested.
// If unspecified, the server will pick an appropriate default.
func (c *AccountsClientsUsersListCall) PageSize(pageSize int64) *AccountsClientsUsersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListClientUsersResponse.nextPageToken
// returned from the previous call to the
// accounts.clients.users.list method.
func (c *AccountsClientsUsersListCall) PageToken(pageToken string) *AccountsClientsUsersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsUsersListCall) Fields(s ...googleapi.Field) *AccountsClientsUsersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsClientsUsersListCall) IfNoneMatch(entityTag string) *AccountsClientsUsersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsUsersListCall) Context(ctx context.Context) *AccountsClientsUsersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsUsersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsUsersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": c.clientAccountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.users.list" call.
// Exactly one of *ListClientUsersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListClientUsersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsClientsUsersListCall) Do(opts ...googleapi.CallOption) (*ListClientUsersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListClientUsersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the known client users for a specified\nsponsor buyer account ID.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.clients.users.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the sponsor buyer of the client to list users for.\n(required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "The account ID of the client buyer to list users for. (required)\nYou must specify either a string representation of a\nnumerical account identifier or the `-` character\nto list all the client users for all the clients\nof a given sponsor buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer clients than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListClientUsersResponse.nextPageToken\nreturned from the previous call to the\naccounts.clients.users.list method.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users",
	//   "response": {
	//     "$ref": "ListClientUsersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsClientsUsersListCall) Pages(ctx context.Context, f func(*ListClientUsersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.clients.users.update":

type AccountsClientsUsersUpdateCall struct {
	s               *Service
	accountId       int64
	clientAccountId int64
	userId          int64
	clientuser      *ClientUser
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Update: Updates an existing client user.
// Only the user status can be changed on update.
func (r *AccountsClientsUsersService) Update(accountId int64, clientAccountId int64, userId int64, clientuser *ClientUser) *AccountsClientsUsersUpdateCall {
	c := &AccountsClientsUsersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.clientAccountId = clientAccountId
	c.userId = userId
	c.clientuser = clientuser
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsClientsUsersUpdateCall) Fields(s ...googleapi.Field) *AccountsClientsUsersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsClientsUsersUpdateCall) Context(ctx context.Context) *AccountsClientsUsersUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsClientsUsersUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsClientsUsersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.clientuser)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       strconv.FormatInt(c.accountId, 10),
		"clientAccountId": strconv.FormatInt(c.clientAccountId, 10),
		"userId":          strconv.FormatInt(c.userId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.clients.users.update" call.
// Exactly one of *ClientUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ClientUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsClientsUsersUpdateCall) Do(opts ...googleapi.CallOption) (*ClientUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing client user.\nOnly the user status can be changed on update.",
	//   "flatPath": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer2.accounts.clients.users.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "clientAccountId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Numerical account ID of the client's sponsor buyer. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clientAccountId": {
	//       "description": "Numerical account ID of the client buyer that the user to be retrieved\nis associated with. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Numerical identifier of the user to retrieve. (required)",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/clients/{clientAccountId}/users/{userId}",
	//   "request": {
	//     "$ref": "ClientUser"
	//   },
	//   "response": {
	//     "$ref": "ClientUser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.create":

type AccountsCreativesCreateCall struct {
	s          *Service
	accountId  string
	creative   *Creative
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a creative.
func (r *AccountsCreativesService) Create(accountId string, creative *Creative) *AccountsCreativesCreateCall {
	c := &AccountsCreativesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creative = creative
	return c
}

// DuplicateIdMode sets the optional parameter "duplicateIdMode":
// Indicates if multiple creatives can share an ID or not. Default
// is
// NO_DUPLICATES (one ID per creative).
//
// Possible values:
//   "NO_DUPLICATES"
//   "FORCE_ENABLE_DUPLICATE_IDS"
func (c *AccountsCreativesCreateCall) DuplicateIdMode(duplicateIdMode string) *AccountsCreativesCreateCall {
	c.urlParams_.Set("duplicateIdMode", duplicateIdMode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesCreateCall) Fields(s ...googleapi.Field) *AccountsCreativesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesCreateCall) Context(ctx context.Context) *AccountsCreativesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.creative)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.create" call.
// Exactly one of *Creative or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Creative.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsCreativesCreateCall) Do(opts ...googleapi.CallOption) (*Creative, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Creative{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a creative.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.creatives.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account that this creative belongs to.\nCan be used to filter the response of the\ncreatives.list\nmethod.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "duplicateIdMode": {
	//       "description": "Indicates if multiple creatives can share an ID or not. Default is\nNO_DUPLICATES (one ID per creative).",
	//       "enum": [
	//         "NO_DUPLICATES",
	//         "FORCE_ENABLE_DUPLICATE_IDS"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.get":

type AccountsCreativesGetCall struct {
	s            *Service
	accountId    string
	creativeId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a creative.
func (r *AccountsCreativesService) Get(accountId string, creativeId string) *AccountsCreativesGetCall {
	c := &AccountsCreativesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesGetCall) Fields(s ...googleapi.Field) *AccountsCreativesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCreativesGetCall) IfNoneMatch(entityTag string) *AccountsCreativesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesGetCall) Context(ctx context.Context) *AccountsCreativesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.get" call.
// Exactly one of *Creative or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Creative.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsCreativesGetCall) Do(opts ...googleapi.CallOption) (*Creative, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Creative{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a creative.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.creatives.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account the creative belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The ID of the creative to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}",
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.list":

type AccountsCreativesListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists creatives.
func (r *AccountsCreativesService) List(accountId string) *AccountsCreativesListCall {
	c := &AccountsCreativesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer creatives than requested
// (due to timeout constraint) even if more are available via another
// call.
// If unspecified, server will pick an appropriate default.
// Acceptable values are 1 to 1000, inclusive.
func (c *AccountsCreativesListCall) PageSize(pageSize int64) *AccountsCreativesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListCreativesResponse.next_page_token
// returned from the previous call to 'ListCreatives' method.
func (c *AccountsCreativesListCall) PageToken(pageToken string) *AccountsCreativesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Query sets the optional parameter "query": An optional query string
// to filter creatives. If no filter is specified,
// all active creatives will be returned.
// <p>Supported queries
// are:
// <ul>
// <li>accountId=<i>account_id_string</i>
// <li>creativeId=<i>cre
// ative_id_string</i>
// <li>dealsStatus: {approved, conditionally_approved, disapproved,
//                    not_checked}
// <li>openAuctionStatus: {approved, conditionally_approved,
// disapproved,
//                           not_checked}
// <li>attribute: {a numeric attribute from the list of
// attributes}
// <li>disapprovalReason: {a reason
// from
// DisapprovalReason}
// </ul>
// Example: 'accountId=12345 AND (dealsStatus:disapproved
// AND
// disapprovalReason:unacceptable_content) OR attribute:47'
func (c *AccountsCreativesListCall) Query(query string) *AccountsCreativesListCall {
	c.urlParams_.Set("query", query)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesListCall) Fields(s ...googleapi.Field) *AccountsCreativesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCreativesListCall) IfNoneMatch(entityTag string) *AccountsCreativesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesListCall) Context(ctx context.Context) *AccountsCreativesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.list" call.
// Exactly one of *ListCreativesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCreativesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsCreativesListCall) Do(opts ...googleapi.CallOption) (*ListCreativesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCreativesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists creatives.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.creatives.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account to list the creatives from.\nSpecify \"-\" to list all creatives the current user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer creatives than requested\n(due to timeout constraint) even if more are available via another call.\nIf unspecified, server will pick an appropriate default.\nAcceptable values are 1 to 1000, inclusive.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListCreativesResponse.next_page_token\nreturned from the previous call to 'ListCreatives' method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "An optional query string to filter creatives. If no filter is specified,\nall active creatives will be returned.\n\u003cp\u003eSupported queries are:\n\u003cul\u003e\n\u003cli\u003eaccountId=\u003ci\u003eaccount_id_string\u003c/i\u003e\n\u003cli\u003ecreativeId=\u003ci\u003ecreative_id_string\u003c/i\u003e\n\u003cli\u003edealsStatus: {approved, conditionally_approved, disapproved,\n                   not_checked}\n\u003cli\u003eopenAuctionStatus: {approved, conditionally_approved, disapproved,\n                          not_checked}\n\u003cli\u003eattribute: {a numeric attribute from the list of attributes}\n\u003cli\u003edisapprovalReason: {a reason from\nDisapprovalReason}\n\u003c/ul\u003e\nExample: 'accountId=12345 AND (dealsStatus:disapproved AND\ndisapprovalReason:unacceptable_content) OR attribute:47'",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives",
	//   "response": {
	//     "$ref": "ListCreativesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsCreativesListCall) Pages(ctx context.Context, f func(*ListCreativesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.creatives.stopWatching":

type AccountsCreativesStopWatchingCall struct {
	s                           *Service
	accountId                   string
	creativeId                  string
	stopwatchingcreativerequest *StopWatchingCreativeRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// StopWatching: Stops watching a creative. Will stop push notifications
// being sent to the
// topics when the creative changes status.
func (r *AccountsCreativesService) StopWatching(accountId string, creativeId string, stopwatchingcreativerequest *StopWatchingCreativeRequest) *AccountsCreativesStopWatchingCall {
	c := &AccountsCreativesStopWatchingCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	c.stopwatchingcreativerequest = stopwatchingcreativerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesStopWatchingCall) Fields(s ...googleapi.Field) *AccountsCreativesStopWatchingCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesStopWatchingCall) Context(ctx context.Context) *AccountsCreativesStopWatchingCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesStopWatchingCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesStopWatchingCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stopwatchingcreativerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}:stopWatching")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.stopWatching" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsCreativesStopWatchingCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Stops watching a creative. Will stop push notifications being sent to the\ntopics when the creative changes status.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}:stopWatching",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.creatives.stopWatching",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account of the creative to stop notifications for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The creative ID of the creative to stop notifications for.\nSpecify \"-\" to specify stopping account level notifications.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}:stopWatching",
	//   "request": {
	//     "$ref": "StopWatchingCreativeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.update":

type AccountsCreativesUpdateCall struct {
	s          *Service
	accountId  string
	creativeId string
	creative   *Creative
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates a creative.
func (r *AccountsCreativesService) Update(accountId string, creativeId string, creative *Creative) *AccountsCreativesUpdateCall {
	c := &AccountsCreativesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	c.creative = creative
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesUpdateCall) Fields(s ...googleapi.Field) *AccountsCreativesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesUpdateCall) Context(ctx context.Context) *AccountsCreativesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.creative)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.update" call.
// Exactly one of *Creative or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Creative.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsCreativesUpdateCall) Do(opts ...googleapi.CallOption) (*Creative, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Creative{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a creative.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer2.accounts.creatives.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account that this creative belongs to.\nCan be used to filter the response of the\ncreatives.list\nmethod.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The buyer-defined creative ID of this creative.\nCan be used to filter the response of the\ncreatives.list\nmethod.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.watch":

type AccountsCreativesWatchCall struct {
	s                    *Service
	accountId            string
	creativeId           string
	watchcreativerequest *WatchCreativeRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Watch: Watches a creative. Will result in push notifications being
// sent to the
// topic when the creative changes status.
func (r *AccountsCreativesService) Watch(accountId string, creativeId string, watchcreativerequest *WatchCreativeRequest) *AccountsCreativesWatchCall {
	c := &AccountsCreativesWatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	c.watchcreativerequest = watchcreativerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesWatchCall) Fields(s ...googleapi.Field) *AccountsCreativesWatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesWatchCall) Context(ctx context.Context) *AccountsCreativesWatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesWatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesWatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.watchcreativerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}:watch")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.watch" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsCreativesWatchCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Watches a creative. Will result in push notifications being sent to the\ntopic when the creative changes status.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}:watch",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.creatives.watch",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account of the creative to watch.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The creative ID to watch for status changes.\nSpecify \"-\" to watch all creatives under the above account.\nIf both creative-level and account-level notifications are\nsent, only a single notification will be sent to the\ncreative-level notification topic.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}:watch",
	//   "request": {
	//     "$ref": "WatchCreativeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.dealAssociations.add":

type AccountsCreativesDealAssociationsAddCall struct {
	s                         *Service
	accountId                 string
	creativeId                string
	adddealassociationrequest *AddDealAssociationRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// Add: Associate an existing deal with a creative.
func (r *AccountsCreativesDealAssociationsService) Add(accountId string, creativeId string, adddealassociationrequest *AddDealAssociationRequest) *AccountsCreativesDealAssociationsAddCall {
	c := &AccountsCreativesDealAssociationsAddCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	c.adddealassociationrequest = adddealassociationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesDealAssociationsAddCall) Fields(s ...googleapi.Field) *AccountsCreativesDealAssociationsAddCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesDealAssociationsAddCall) Context(ctx context.Context) *AccountsCreativesDealAssociationsAddCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesDealAssociationsAddCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesDealAssociationsAddCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.adddealassociationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:add")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.dealAssociations.add" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsCreativesDealAssociationsAddCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Associate an existing deal with a creative.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:add",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.creatives.dealAssociations.add",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account the creative belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The ID of the creative associated with the deal.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:add",
	//   "request": {
	//     "$ref": "AddDealAssociationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.creatives.dealAssociations.list":

type AccountsCreativesDealAssociationsListCall struct {
	s            *Service
	accountId    string
	creativeId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all creative-deal associations.
func (r *AccountsCreativesDealAssociationsService) List(accountId string, creativeId string) *AccountsCreativesDealAssociationsListCall {
	c := &AccountsCreativesDealAssociationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Server may return fewer associations than requested.
// If unspecified, server will pick an appropriate default.
func (c *AccountsCreativesDealAssociationsListCall) PageSize(pageSize int64) *AccountsCreativesDealAssociationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListDealAssociationsResponse.next_page_token
// returned from the previous call to 'ListDealAssociations' method.
func (c *AccountsCreativesDealAssociationsListCall) PageToken(pageToken string) *AccountsCreativesDealAssociationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Query sets the optional parameter "query": An optional query string
// to filter deal associations. If no filter is
// specified, all associations will be returned.
// Supported queries
// are:
// <ul>
// <li>accountId=<i>account_id_string</i>
// <li>creativeId=<i>cre
// ative_id_string</i>
// <li>dealsId=<i>deals_id_string</i>
// <li>dealsStatus
// :{approved, conditionally_approved, disapproved,
//                   not_checked}
// <li>openAuctionStatus:{approved, conditionally_approved,
// disapproved,
//                          not_checked}
// </ul>
// Example: 'dealsId=12345 AND dealsStatus:disapproved'
func (c *AccountsCreativesDealAssociationsListCall) Query(query string) *AccountsCreativesDealAssociationsListCall {
	c.urlParams_.Set("query", query)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesDealAssociationsListCall) Fields(s ...googleapi.Field) *AccountsCreativesDealAssociationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCreativesDealAssociationsListCall) IfNoneMatch(entityTag string) *AccountsCreativesDealAssociationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesDealAssociationsListCall) Context(ctx context.Context) *AccountsCreativesDealAssociationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesDealAssociationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesDealAssociationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.dealAssociations.list" call.
// Exactly one of *ListDealAssociationsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListDealAssociationsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsCreativesDealAssociationsListCall) Do(opts ...googleapi.CallOption) (*ListDealAssociationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDealAssociationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all creative-deal associations.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.creatives.dealAssociations.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account to list the associations from.\nSpecify \"-\" to list all creatives the current user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The creative ID to list the associations from.\nSpecify \"-\" to list all creatives under the above account.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Server may return fewer associations than requested.\nIf unspecified, server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListDealAssociationsResponse.next_page_token\nreturned from the previous call to 'ListDealAssociations' method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "An optional query string to filter deal associations. If no filter is\nspecified, all associations will be returned.\nSupported queries are:\n\u003cul\u003e\n\u003cli\u003eaccountId=\u003ci\u003eaccount_id_string\u003c/i\u003e\n\u003cli\u003ecreativeId=\u003ci\u003ecreative_id_string\u003c/i\u003e\n\u003cli\u003edealsId=\u003ci\u003edeals_id_string\u003c/i\u003e\n\u003cli\u003edealsStatus:{approved, conditionally_approved, disapproved,\n                  not_checked}\n\u003cli\u003eopenAuctionStatus:{approved, conditionally_approved, disapproved,\n                         not_checked}\n\u003c/ul\u003e\nExample: 'dealsId=12345 AND dealsStatus:disapproved'",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations",
	//   "response": {
	//     "$ref": "ListDealAssociationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsCreativesDealAssociationsListCall) Pages(ctx context.Context, f func(*ListDealAssociationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.creatives.dealAssociations.remove":

type AccountsCreativesDealAssociationsRemoveCall struct {
	s                            *Service
	accountId                    string
	creativeId                   string
	removedealassociationrequest *RemoveDealAssociationRequest
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// Remove: Remove the association between a deal and a creative.
func (r *AccountsCreativesDealAssociationsService) Remove(accountId string, creativeId string, removedealassociationrequest *RemoveDealAssociationRequest) *AccountsCreativesDealAssociationsRemoveCall {
	c := &AccountsCreativesDealAssociationsRemoveCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.creativeId = creativeId
	c.removedealassociationrequest = removedealassociationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCreativesDealAssociationsRemoveCall) Fields(s ...googleapi.Field) *AccountsCreativesDealAssociationsRemoveCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCreativesDealAssociationsRemoveCall) Context(ctx context.Context) *AccountsCreativesDealAssociationsRemoveCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCreativesDealAssociationsRemoveCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCreativesDealAssociationsRemoveCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.removedealassociationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:remove")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.creatives.dealAssociations.remove" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsCreativesDealAssociationsRemoveCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove the association between a deal and a creative.",
	//   "flatPath": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:remove",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.creatives.dealAssociations.remove",
	//   "parameterOrder": [
	//     "accountId",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The account the creative belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "creativeId": {
	//       "description": "The ID of the creative associated with the deal.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/creatives/{creativeId}/dealAssociations:remove",
	//   "request": {
	//     "$ref": "RemoveDealAssociationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.finalizedProposals.list":

type AccountsFinalizedProposalsListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List finalized proposals, regardless if a proposal is being
// renegotiated.
// A filter expression (PQL query) may be specified to filter the
// results.
// The notes will not be returned.
func (r *AccountsFinalizedProposalsService) List(accountId string) *AccountsFinalizedProposalsListCall {
	c := &AccountsFinalizedProposalsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// Filter sets the optional parameter "filter": An optional PQL filter
// query used to query for proposals.
//
// Nested repeated fields, such as
// proposal.deals.targetingCriterion,
// cannot be filtered.
func (c *AccountsFinalizedProposalsListCall) Filter(filter string) *AccountsFinalizedProposalsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// FilterSyntax sets the optional parameter "filterSyntax": Syntax the
// filter is written in. Current implementation defaults to PQL
// but in the future it will be LIST_FILTER.
//
// Possible values:
//   "FILTER_SYNTAX_UNSPECIFIED"
//   "PQL"
//   "LIST_FILTER"
func (c *AccountsFinalizedProposalsListCall) FilterSyntax(filterSyntax string) *AccountsFinalizedProposalsListCall {
	c.urlParams_.Set("filterSyntax", filterSyntax)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *AccountsFinalizedProposalsListCall) PageSize(pageSize int64) *AccountsFinalizedProposalsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token as
// returned from ListProposalsResponse.
func (c *AccountsFinalizedProposalsListCall) PageToken(pageToken string) *AccountsFinalizedProposalsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsFinalizedProposalsListCall) Fields(s ...googleapi.Field) *AccountsFinalizedProposalsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsFinalizedProposalsListCall) IfNoneMatch(entityTag string) *AccountsFinalizedProposalsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsFinalizedProposalsListCall) Context(ctx context.Context) *AccountsFinalizedProposalsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsFinalizedProposalsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsFinalizedProposalsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/finalizedProposals")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.finalizedProposals.list" call.
// Exactly one of *ListProposalsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListProposalsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsFinalizedProposalsListCall) Do(opts ...googleapi.CallOption) (*ListProposalsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListProposalsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List finalized proposals, regardless if a proposal is being renegotiated.\nA filter expression (PQL query) may be specified to filter the results.\nThe notes will not be returned.",
	//   "flatPath": "v2beta1/accounts/{accountId}/finalizedProposals",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.finalizedProposals.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "An optional PQL filter query used to query for proposals.\n\nNested repeated fields, such as proposal.deals.targetingCriterion,\ncannot be filtered.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filterSyntax": {
	//       "description": "Syntax the filter is written in. Current implementation defaults to PQL\nbut in the future it will be LIST_FILTER.",
	//       "enum": [
	//         "FILTER_SYNTAX_UNSPECIFIED",
	//         "PQL",
	//         "LIST_FILTER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The page token as returned from ListProposalsResponse.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/finalizedProposals",
	//   "response": {
	//     "$ref": "ListProposalsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsFinalizedProposalsListCall) Pages(ctx context.Context, f func(*ListProposalsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.products.get":

type AccountsProductsGetCall struct {
	s            *Service
	accountId    string
	productId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the requested product by ID.
func (r *AccountsProductsService) Get(accountId string, productId string) *AccountsProductsGetCall {
	c := &AccountsProductsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.productId = productId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProductsGetCall) Fields(s ...googleapi.Field) *AccountsProductsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsProductsGetCall) IfNoneMatch(entityTag string) *AccountsProductsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProductsGetCall) Context(ctx context.Context) *AccountsProductsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProductsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProductsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/products/{productId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
		"productId": c.productId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.products.get" call.
// Exactly one of *Product or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Product.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsProductsGetCall) Do(opts ...googleapi.CallOption) (*Product, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Product{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the requested product by ID.",
	//   "flatPath": "v2beta1/accounts/{accountId}/products/{productId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.products.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID for the product to get the head revision for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/products/{productId}",
	//   "response": {
	//     "$ref": "Product"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.products.list":

type AccountsProductsListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all products visible to the buyer (optionally filtered by
// the
// specified PQL query).
func (r *AccountsProductsService) List(accountId string) *AccountsProductsListCall {
	c := &AccountsProductsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// Filter sets the optional parameter "filter": An optional PQL query
// used to query for products.
// See
// https://developers.google.com/ad-manager/docs/pqlreference
// for documentation about PQL and examples.
//
// Nested repeated fields, such as
// product.targetingCriterion.inclusions,
// cannot be filtered.
func (c *AccountsProductsListCall) Filter(filter string) *AccountsProductsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *AccountsProductsListCall) PageSize(pageSize int64) *AccountsProductsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token as
// returned from ListProductsResponse.
func (c *AccountsProductsListCall) PageToken(pageToken string) *AccountsProductsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProductsListCall) Fields(s ...googleapi.Field) *AccountsProductsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsProductsListCall) IfNoneMatch(entityTag string) *AccountsProductsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProductsListCall) Context(ctx context.Context) *AccountsProductsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProductsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProductsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/products")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.products.list" call.
// Exactly one of *ListProductsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListProductsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsProductsListCall) Do(opts ...googleapi.CallOption) (*ListProductsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListProductsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all products visible to the buyer (optionally filtered by the\nspecified PQL query).",
	//   "flatPath": "v2beta1/accounts/{accountId}/products",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.products.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "An optional PQL query used to query for products. See\nhttps://developers.google.com/ad-manager/docs/pqlreference\nfor documentation about PQL and examples.\n\nNested repeated fields, such as product.targetingCriterion.inclusions,\ncannot be filtered.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The page token as returned from ListProductsResponse.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/products",
	//   "response": {
	//     "$ref": "ListProductsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsProductsListCall) Pages(ctx context.Context, f func(*ListProductsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.proposals.accept":

type AccountsProposalsAcceptCall struct {
	s                     *Service
	accountId             string
	proposalId            string
	acceptproposalrequest *AcceptProposalRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Accept: Mark the proposal as accepted at the given revision number.
// If the number
// does not match the server's revision number an `ABORTED` error
// message will
// be returned. This call updates the proposal_state from `PROPOSED`
// to
// `BUYER_ACCEPTED`, or from `SELLER_ACCEPTED` to `FINALIZED`.
func (r *AccountsProposalsService) Accept(accountId string, proposalId string, acceptproposalrequest *AcceptProposalRequest) *AccountsProposalsAcceptCall {
	c := &AccountsProposalsAcceptCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.acceptproposalrequest = acceptproposalrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsAcceptCall) Fields(s ...googleapi.Field) *AccountsProposalsAcceptCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsAcceptCall) Context(ctx context.Context) *AccountsProposalsAcceptCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsAcceptCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsAcceptCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.acceptproposalrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:accept")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.accept" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsAcceptCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Mark the proposal as accepted at the given revision number. If the number\ndoes not match the server's revision number an `ABORTED` error message will\nbe returned. This call updates the proposal_state from `PROPOSED` to\n`BUYER_ACCEPTED`, or from `SELLER_ACCEPTED` to `FINALIZED`.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:accept",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.accept",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to accept.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:accept",
	//   "request": {
	//     "$ref": "AcceptProposalRequest"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.addNote":

type AccountsProposalsAddNoteCall struct {
	s              *Service
	accountId      string
	proposalId     string
	addnoterequest *AddNoteRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// AddNote: Create a new note and attach it to the proposal. The note is
// assigned
// a unique ID by the server.
// The proposal revision number will not increase when associated with
// a
// new note.
func (r *AccountsProposalsService) AddNote(accountId string, proposalId string, addnoterequest *AddNoteRequest) *AccountsProposalsAddNoteCall {
	c := &AccountsProposalsAddNoteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.addnoterequest = addnoterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsAddNoteCall) Fields(s ...googleapi.Field) *AccountsProposalsAddNoteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsAddNoteCall) Context(ctx context.Context) *AccountsProposalsAddNoteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsAddNoteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsAddNoteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addnoterequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:addNote")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.addNote" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AccountsProposalsAddNoteCall) Do(opts ...googleapi.CallOption) (*Note, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Note{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new note and attach it to the proposal. The note is assigned\na unique ID by the server.\nThe proposal revision number will not increase when associated with a\nnew note.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:addNote",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.addNote",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to attach the note to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:addNote",
	//   "request": {
	//     "$ref": "AddNoteRequest"
	//   },
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.cancelNegotiation":

type AccountsProposalsCancelNegotiationCall struct {
	s                        *Service
	accountId                string
	proposalId               string
	cancelnegotiationrequest *CancelNegotiationRequest
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// CancelNegotiation: Cancel an ongoing negotiation on a proposal. This
// does not cancel or end
// serving for the deals if the proposal has been finalized, but only
// cancels
// a negotiation unilaterally.
func (r *AccountsProposalsService) CancelNegotiation(accountId string, proposalId string, cancelnegotiationrequest *CancelNegotiationRequest) *AccountsProposalsCancelNegotiationCall {
	c := &AccountsProposalsCancelNegotiationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.cancelnegotiationrequest = cancelnegotiationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsCancelNegotiationCall) Fields(s ...googleapi.Field) *AccountsProposalsCancelNegotiationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsCancelNegotiationCall) Context(ctx context.Context) *AccountsProposalsCancelNegotiationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsCancelNegotiationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsCancelNegotiationCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cancelnegotiationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:cancelNegotiation")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.cancelNegotiation" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsCancelNegotiationCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel an ongoing negotiation on a proposal. This does not cancel or end\nserving for the deals if the proposal has been finalized, but only cancels\na negotiation unilaterally.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:cancelNegotiation",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.cancelNegotiation",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to cancel negotiation for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:cancelNegotiation",
	//   "request": {
	//     "$ref": "CancelNegotiationRequest"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.completeSetup":

type AccountsProposalsCompleteSetupCall struct {
	s                    *Service
	accountId            string
	proposalId           string
	completesetuprequest *CompleteSetupRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// CompleteSetup: Update the given proposal to indicate that setup has
// been completed.
// This method is called by the buyer when the line items have been
// created
// on their end for a finalized proposal and all the required
// creatives
// have been uploaded using the creatives API. This call updates
// the
// `is_setup_completed` bit on the proposal and also notifies the
// seller.
// The server will advance the revision number of the most recent
// proposal.
func (r *AccountsProposalsService) CompleteSetup(accountId string, proposalId string, completesetuprequest *CompleteSetupRequest) *AccountsProposalsCompleteSetupCall {
	c := &AccountsProposalsCompleteSetupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.completesetuprequest = completesetuprequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsCompleteSetupCall) Fields(s ...googleapi.Field) *AccountsProposalsCompleteSetupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsCompleteSetupCall) Context(ctx context.Context) *AccountsProposalsCompleteSetupCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsCompleteSetupCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsCompleteSetupCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.completesetuprequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:completeSetup")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.completeSetup" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsCompleteSetupCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the given proposal to indicate that setup has been completed.\nThis method is called by the buyer when the line items have been created\non their end for a finalized proposal and all the required creatives\nhave been uploaded using the creatives API. This call updates the\n`is_setup_completed` bit on the proposal and also notifies the seller.\nThe server will advance the revision number of the most recent proposal.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:completeSetup",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.completeSetup",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to mark as setup completed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:completeSetup",
	//   "request": {
	//     "$ref": "CompleteSetupRequest"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.create":

type AccountsProposalsCreateCall struct {
	s          *Service
	accountId  string
	proposal   *Proposal
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Create the given proposal. Each created proposal and any
// deals it contains
// are assigned a unique ID by the server.
func (r *AccountsProposalsService) Create(accountId string, proposal *Proposal) *AccountsProposalsCreateCall {
	c := &AccountsProposalsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposal = proposal
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsCreateCall) Fields(s ...googleapi.Field) *AccountsProposalsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsCreateCall) Context(ctx context.Context) *AccountsProposalsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.proposal)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.create" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsCreateCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create the given proposal. Each created proposal and any deals it contains\nare assigned a unique ID by the server.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals",
	//   "request": {
	//     "$ref": "Proposal"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.get":

type AccountsProposalsGetCall struct {
	s            *Service
	accountId    string
	proposalId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a proposal given its ID. The proposal is returned at its
// head
// revision.
func (r *AccountsProposalsService) Get(accountId string, proposalId string) *AccountsProposalsGetCall {
	c := &AccountsProposalsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsGetCall) Fields(s ...googleapi.Field) *AccountsProposalsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsProposalsGetCall) IfNoneMatch(entityTag string) *AccountsProposalsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsGetCall) Context(ctx context.Context) *AccountsProposalsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.get" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsGetCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a proposal given its ID. The proposal is returned at its head\nrevision.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.proposals.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The unique ID of the proposal",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}",
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.list":

type AccountsProposalsListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List proposals. A filter expression (PQL query) may be
// specified to
// filter the results. To retrieve all finalized proposals, regardless
// if a
// proposal is being renegotiated, see the FinalizedProposals
// resource.
// Note that Bidder/ChildSeat relationships differ from the usual
// behavior.
// A Bidder account can only see its child seats' proposals by
// specifying
// the ChildSeat's accountId in the request path.
func (r *AccountsProposalsService) List(accountId string) *AccountsProposalsListCall {
	c := &AccountsProposalsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// Filter sets the optional parameter "filter": An optional PQL filter
// query used to query for proposals.
//
// Nested repeated fields, such as
// proposal.deals.targetingCriterion,
// cannot be filtered.
func (c *AccountsProposalsListCall) Filter(filter string) *AccountsProposalsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// FilterSyntax sets the optional parameter "filterSyntax": Syntax the
// filter is written in. Current implementation defaults to PQL
// but in the future it will be LIST_FILTER.
//
// Possible values:
//   "FILTER_SYNTAX_UNSPECIFIED"
//   "PQL"
//   "LIST_FILTER"
func (c *AccountsProposalsListCall) FilterSyntax(filterSyntax string) *AccountsProposalsListCall {
	c.urlParams_.Set("filterSyntax", filterSyntax)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *AccountsProposalsListCall) PageSize(pageSize int64) *AccountsProposalsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token as
// returned from ListProposalsResponse.
func (c *AccountsProposalsListCall) PageToken(pageToken string) *AccountsProposalsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsListCall) Fields(s ...googleapi.Field) *AccountsProposalsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsProposalsListCall) IfNoneMatch(entityTag string) *AccountsProposalsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsListCall) Context(ctx context.Context) *AccountsProposalsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.list" call.
// Exactly one of *ListProposalsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListProposalsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsProposalsListCall) Do(opts ...googleapi.CallOption) (*ListProposalsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListProposalsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List proposals. A filter expression (PQL query) may be specified to\nfilter the results. To retrieve all finalized proposals, regardless if a\nproposal is being renegotiated, see the FinalizedProposals resource.\nNote that Bidder/ChildSeat relationships differ from the usual behavior.\nA Bidder account can only see its child seats' proposals by specifying\nthe ChildSeat's accountId in the request path.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.proposals.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "An optional PQL filter query used to query for proposals.\n\nNested repeated fields, such as proposal.deals.targetingCriterion,\ncannot be filtered.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filterSyntax": {
	//       "description": "Syntax the filter is written in. Current implementation defaults to PQL\nbut in the future it will be LIST_FILTER.",
	//       "enum": [
	//         "FILTER_SYNTAX_UNSPECIFIED",
	//         "PQL",
	//         "LIST_FILTER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The page token as returned from ListProposalsResponse.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals",
	//   "response": {
	//     "$ref": "ListProposalsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsProposalsListCall) Pages(ctx context.Context, f func(*ListProposalsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.accounts.proposals.pause":

type AccountsProposalsPauseCall struct {
	s                    *Service
	accountId            string
	proposalId           string
	pauseproposalrequest *PauseProposalRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Pause: Update the given proposal to pause serving.
// This method will set
// the
// `DealServingMetadata.DealPauseStatus.has_buyer_paused` bit to true
// for all
// deals in the proposal.
//
// It is a no-op to pause an already-paused proposal.
// It is an error to call PauseProposal for a proposal that is
// not
// finalized or renegotiating.
func (r *AccountsProposalsService) Pause(accountId string, proposalId string, pauseproposalrequest *PauseProposalRequest) *AccountsProposalsPauseCall {
	c := &AccountsProposalsPauseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.pauseproposalrequest = pauseproposalrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsPauseCall) Fields(s ...googleapi.Field) *AccountsProposalsPauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsPauseCall) Context(ctx context.Context) *AccountsProposalsPauseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsPauseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsPauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pauseproposalrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:pause")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.pause" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsPauseCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the given proposal to pause serving.\nThis method will set the\n`DealServingMetadata.DealPauseStatus.has_buyer_paused` bit to true for all\ndeals in the proposal.\n\nIt is a no-op to pause an already-paused proposal.\nIt is an error to call PauseProposal for a proposal that is not\nfinalized or renegotiating.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:pause",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.pause",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to pause.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:pause",
	//   "request": {
	//     "$ref": "PauseProposalRequest"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.resume":

type AccountsProposalsResumeCall struct {
	s                     *Service
	accountId             string
	proposalId            string
	resumeproposalrequest *ResumeProposalRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Resume: Update the given proposal to resume serving.
// This method will set
// the
// `DealServingMetadata.DealPauseStatus.has_buyer_paused` bit to false
// for all
// deals in the proposal.
//
// Note that if the `has_seller_paused` bit is also set, serving will
// not
// resume until the seller also resumes.
//
// It is a no-op to resume an already-running proposal.
// It is an error to call ResumeProposal for a proposal that is
// not
// finalized or renegotiating.
func (r *AccountsProposalsService) Resume(accountId string, proposalId string, resumeproposalrequest *ResumeProposalRequest) *AccountsProposalsResumeCall {
	c := &AccountsProposalsResumeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.resumeproposalrequest = resumeproposalrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsResumeCall) Fields(s ...googleapi.Field) *AccountsProposalsResumeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsResumeCall) Context(ctx context.Context) *AccountsProposalsResumeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsResumeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsResumeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resumeproposalrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}:resume")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.resume" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsResumeCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the given proposal to resume serving.\nThis method will set the\n`DealServingMetadata.DealPauseStatus.has_buyer_paused` bit to false for all\ndeals in the proposal.\n\nNote that if the `has_seller_paused` bit is also set, serving will not\nresume until the seller also resumes.\n\nIt is a no-op to resume an already-running proposal.\nIt is an error to call ResumeProposal for a proposal that is not\nfinalized or renegotiating.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}:resume",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.accounts.proposals.resume",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The ID of the proposal to resume.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}:resume",
	//   "request": {
	//     "$ref": "ResumeProposalRequest"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.proposals.update":

type AccountsProposalsUpdateCall struct {
	s          *Service
	accountId  string
	proposalId string
	proposal   *Proposal
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Update the given proposal at the client known revision
// number. If the
// server revision has advanced since the
// passed-in
// `proposal.proposal_revision`, an `ABORTED` error message will be
// returned.
// Only the buyer-modifiable fields of the proposal will be
// updated.
//
// Note that the deals in the proposal will be updated to match the
// passed-in
// copy.
// If a passed-in deal does not have a `deal_id`, the server will assign
// a new
// unique ID and create the deal.
// If passed-in deal has a `deal_id`, it will be updated to match
// the
// passed-in copy.
// Any existing deals not present in the passed-in proposal will be
// deleted.
// It is an error to pass in a deal with a `deal_id` not present at
// head.
func (r *AccountsProposalsService) Update(accountId string, proposalId string, proposal *Proposal) *AccountsProposalsUpdateCall {
	c := &AccountsProposalsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.proposalId = proposalId
	c.proposal = proposal
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsProposalsUpdateCall) Fields(s ...googleapi.Field) *AccountsProposalsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsProposalsUpdateCall) Context(ctx context.Context) *AccountsProposalsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsProposalsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsProposalsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.proposal)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/proposals/{proposalId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"proposalId": c.proposalId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.proposals.update" call.
// Exactly one of *Proposal or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Proposal.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsProposalsUpdateCall) Do(opts ...googleapi.CallOption) (*Proposal, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Proposal{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the given proposal at the client known revision number. If the\nserver revision has advanced since the passed-in\n`proposal.proposal_revision`, an `ABORTED` error message will be returned.\nOnly the buyer-modifiable fields of the proposal will be updated.\n\nNote that the deals in the proposal will be updated to match the passed-in\ncopy.\nIf a passed-in deal does not have a `deal_id`, the server will assign a new\nunique ID and create the deal.\nIf passed-in deal has a `deal_id`, it will be updated to match the\npassed-in copy.\nAny existing deals not present in the passed-in proposal will be deleted.\nIt is an error to pass in a deal with a `deal_id` not present at head.",
	//   "flatPath": "v2beta1/accounts/{accountId}/proposals/{proposalId}",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer2.accounts.proposals.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "proposalId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "proposalId": {
	//       "description": "The unique ID of the proposal.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/proposals/{proposalId}",
	//   "request": {
	//     "$ref": "Proposal"
	//   },
	//   "response": {
	//     "$ref": "Proposal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.publisherProfiles.get":

type AccountsPublisherProfilesGetCall struct {
	s                  *Service
	accountId          string
	publisherProfileId string
	urlParams_         gensupport.URLParams
	ifNoneMatch_       string
	ctx_               context.Context
	header_            http.Header
}

// Get: Gets the requested publisher profile by id.
func (r *AccountsPublisherProfilesService) Get(accountId string, publisherProfileId string) *AccountsPublisherProfilesGetCall {
	c := &AccountsPublisherProfilesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.publisherProfileId = publisherProfileId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPublisherProfilesGetCall) Fields(s ...googleapi.Field) *AccountsPublisherProfilesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsPublisherProfilesGetCall) IfNoneMatch(entityTag string) *AccountsPublisherProfilesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsPublisherProfilesGetCall) Context(ctx context.Context) *AccountsPublisherProfilesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsPublisherProfilesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsPublisherProfilesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/publisherProfiles/{publisherProfileId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"publisherProfileId": c.publisherProfileId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.publisherProfiles.get" call.
// Exactly one of *PublisherProfile or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *PublisherProfile.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsPublisherProfilesGetCall) Do(opts ...googleapi.CallOption) (*PublisherProfile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PublisherProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the requested publisher profile by id.",
	//   "flatPath": "v2beta1/accounts/{accountId}/publisherProfiles/{publisherProfileId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.publisherProfiles.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "publisherProfileId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publisherProfileId": {
	//       "description": "The id for the publisher profile to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/publisherProfiles/{publisherProfileId}",
	//   "response": {
	//     "$ref": "PublisherProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.accounts.publisherProfiles.list":

type AccountsPublisherProfilesListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all publisher profiles visible to the buyer
func (r *AccountsPublisherProfilesService) List(accountId string) *AccountsPublisherProfilesListCall {
	c := &AccountsPublisherProfilesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// PageSize sets the optional parameter "pageSize": Specify the number
// of results to include per page.
func (c *AccountsPublisherProfilesListCall) PageSize(pageSize int64) *AccountsPublisherProfilesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The page token as
// return from ListPublisherProfilesResponse.
func (c *AccountsPublisherProfilesListCall) PageToken(pageToken string) *AccountsPublisherProfilesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPublisherProfilesListCall) Fields(s ...googleapi.Field) *AccountsPublisherProfilesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsPublisherProfilesListCall) IfNoneMatch(entityTag string) *AccountsPublisherProfilesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsPublisherProfilesListCall) Context(ctx context.Context) *AccountsPublisherProfilesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsPublisherProfilesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsPublisherProfilesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/accounts/{accountId}/publisherProfiles")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.accounts.publisherProfiles.list" call.
// Exactly one of *ListPublisherProfilesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListPublisherProfilesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsPublisherProfilesListCall) Do(opts ...googleapi.CallOption) (*ListPublisherProfilesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListPublisherProfilesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all publisher profiles visible to the buyer",
	//   "flatPath": "v2beta1/accounts/{accountId}/publisherProfiles",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.accounts.publisherProfiles.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the buyer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Specify the number of results to include per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The page token as return from ListPublisherProfilesResponse.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/accounts/{accountId}/publisherProfiles",
	//   "response": {
	//     "$ref": "ListPublisherProfilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AccountsPublisherProfilesListCall) Pages(ctx context.Context, f func(*ListPublisherProfilesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.creatives.delete":

type BiddersAccountsCreativesDeleteCall struct {
	s          *Service
	ownerName  string
	creativeId string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a single creative.
//
// A creative is deactivated upon deletion and does not count against
// active
// snippet quota. A deleted creative should not be used in bidding (all
// bids
// with that creative will be rejected).
func (r *BiddersAccountsCreativesService) Delete(ownerName string, creativeId string) *BiddersAccountsCreativesDeleteCall {
	c := &BiddersAccountsCreativesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.ownerName = ownerName
	c.creativeId = creativeId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsCreativesDeleteCall) Fields(s ...googleapi.Field) *BiddersAccountsCreativesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsCreativesDeleteCall) Context(ctx context.Context) *BiddersAccountsCreativesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsCreativesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsCreativesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+ownerName}/creatives/{creativeId}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"ownerName":  c.ownerName,
		"creativeId": c.creativeId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.creatives.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BiddersAccountsCreativesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a single creative.\n\nA creative is deactivated upon deletion and does not count against active\nsnippet quota. A deleted creative should not be used in bidding (all bids\nwith that creative will be rejected).",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/creatives/{creativeId}",
	//   "httpMethod": "DELETE",
	//   "id": "adexchangebuyer2.bidders.accounts.creatives.delete",
	//   "parameterOrder": [
	//     "ownerName",
	//     "creativeId"
	//   ],
	//   "parameters": {
	//     "creativeId": {
	//       "description": "The ID of the creative to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ownerName": {
	//       "description": "Name of the owner (bidder or account) of the creative to be deleted.\nFor example:\n\n- For an account-level creative for the buyer account representing bidder\n  123: `bidders/123/accounts/123`\n\n- For an account-level creative for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+ownerName}/creatives/{creativeId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.accounts.filterSets.create":

type BiddersAccountsFilterSetsCreateCall struct {
	s          *Service
	ownerName  string
	filterset  *FilterSet
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates the specified filter set for the account with the
// given account ID.
func (r *BiddersAccountsFilterSetsService) Create(ownerName string, filterset *FilterSet) *BiddersAccountsFilterSetsCreateCall {
	c := &BiddersAccountsFilterSetsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.ownerName = ownerName
	c.filterset = filterset
	return c
}

// IsTransient sets the optional parameter "isTransient": Whether the
// filter set is transient, or should be persisted indefinitely.
// By default, filter sets are not transient.
// If transient, it will be available for at least 1 hour after
// creation.
func (c *BiddersAccountsFilterSetsCreateCall) IsTransient(isTransient bool) *BiddersAccountsFilterSetsCreateCall {
	c.urlParams_.Set("isTransient", fmt.Sprint(isTransient))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsCreateCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsCreateCall) Context(ctx context.Context) *BiddersAccountsFilterSetsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.filterset)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+ownerName}/filterSets")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"ownerName": c.ownerName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.create" call.
// Exactly one of *FilterSet or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *FilterSet.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsCreateCall) Do(opts ...googleapi.CallOption) (*FilterSet, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &FilterSet{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified filter set for the account with the given account ID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.create",
	//   "parameterOrder": [
	//     "ownerName"
	//   ],
	//   "parameters": {
	//     "isTransient": {
	//       "description": "Whether the filter set is transient, or should be persisted indefinitely.\nBy default, filter sets are not transient.\nIf transient, it will be available for at least 1 hour after creation.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ownerName": {
	//       "description": "Name of the owner (bidder or account) of the filter set to be created.\nFor example:\n\n- For a bidder-level filter set for bidder 123: `bidders/123`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+ownerName}/filterSets",
	//   "request": {
	//     "$ref": "FilterSet"
	//   },
	//   "response": {
	//     "$ref": "FilterSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.accounts.filterSets.delete":

type BiddersAccountsFilterSetsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the requested filter set from the account with the
// given account
// ID.
func (r *BiddersAccountsFilterSetsService) Delete(name string) *BiddersAccountsFilterSetsDeleteCall {
	c := &BiddersAccountsFilterSetsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsDeleteCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsDeleteCall) Context(ctx context.Context) *BiddersAccountsFilterSetsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BiddersAccountsFilterSetsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the requested filter set from the account with the given account\nID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}",
	//   "httpMethod": "DELETE",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Full name of the resource to delete.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.accounts.filterSets.get":

type BiddersAccountsFilterSetsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the requested filter set for the account with the
// given account
// ID.
func (r *BiddersAccountsFilterSetsService) Get(name string) *BiddersAccountsFilterSetsGetCall {
	c := &BiddersAccountsFilterSetsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsGetCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsGetCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsGetCall) Context(ctx context.Context) *BiddersAccountsFilterSetsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.get" call.
// Exactly one of *FilterSet or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *FilterSet.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsGetCall) Do(opts ...googleapi.CallOption) (*FilterSet, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &FilterSet{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the requested filter set for the account with the given account\nID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Full name of the resource being requested.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "FilterSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.accounts.filterSets.list":

type BiddersAccountsFilterSetsListCall struct {
	s            *Service
	ownerName    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all filter sets for the account with the given account
// ID.
func (r *BiddersAccountsFilterSetsService) List(ownerName string) *BiddersAccountsFilterSetsListCall {
	c := &BiddersAccountsFilterSetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.ownerName = ownerName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilterSetsResponse.nextPageToken
// returned from the previous call to
// the
// accounts.filterSets.list
// method.
func (c *BiddersAccountsFilterSetsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+ownerName}/filterSets")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"ownerName": c.ownerName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.list" call.
// Exactly one of *ListFilterSetsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListFilterSetsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsListCall) Do(opts ...googleapi.CallOption) (*ListFilterSetsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilterSetsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all filter sets for the account with the given account ID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.list",
	//   "parameterOrder": [
	//     "ownerName"
	//   ],
	//   "parameters": {
	//     "ownerName": {
	//       "description": "Name of the owner (bidder or account) of the filter sets to be listed.\nFor example:\n\n- For a bidder-level filter set for bidder 123: `bidders/123`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilterSetsResponse.nextPageToken\nreturned from the previous call to the\naccounts.filterSets.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+ownerName}/filterSets",
	//   "response": {
	//     "$ref": "ListFilterSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsListCall) Pages(ctx context.Context, f func(*ListFilterSetsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.bidMetrics.list":

type BiddersAccountsFilterSetsBidMetricsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: Lists all metrics that are measured in terms of number of bids.
func (r *BiddersAccountsFilterSetsBidMetricsService) List(filterSetName string) *BiddersAccountsFilterSetsBidMetricsListCall {
	c := &BiddersAccountsFilterSetsBidMetricsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsBidMetricsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidMetricsResponse.nextPageToken
// returned from the previous call to the bidMetrics.list
// method.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsBidMetricsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsBidMetricsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsBidMetricsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsBidMetricsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsBidMetricsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidMetrics")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.bidMetrics.list" call.
// Exactly one of *ListBidMetricsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListBidMetricsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) Do(opts ...googleapi.CallOption) (*ListBidMetricsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidMetricsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all metrics that are measured in terms of number of bids.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/bidMetrics",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.bidMetrics.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidMetricsResponse.nextPageToken\nreturned from the previous call to the bidMetrics.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidMetrics",
	//   "response": {
	//     "$ref": "ListBidMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsBidMetricsListCall) Pages(ctx context.Context, f func(*ListBidMetricsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.bidResponseErrors.list":

type BiddersAccountsFilterSetsBidResponseErrorsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all errors that occurred in bid responses, with the number
// of bid
// responses affected for each reason.
func (r *BiddersAccountsFilterSetsBidResponseErrorsService) List(filterSetName string) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c := &BiddersAccountsFilterSetsBidResponseErrorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidResponseErrorsResponse.nextPageToken
// returned from the previous call to the bidResponseErrors.list
// method.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsBidResponseErrorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidResponseErrors")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.bidResponseErrors.list" call.
// Exactly one of *ListBidResponseErrorsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListBidResponseErrorsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) Do(opts ...googleapi.CallOption) (*ListBidResponseErrorsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidResponseErrorsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all errors that occurred in bid responses, with the number of bid\nresponses affected for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/bidResponseErrors",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.bidResponseErrors.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidResponseErrorsResponse.nextPageToken\nreturned from the previous call to the bidResponseErrors.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidResponseErrors",
	//   "response": {
	//     "$ref": "ListBidResponseErrorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsBidResponseErrorsListCall) Pages(ctx context.Context, f func(*ListBidResponseErrorsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.bidResponsesWithoutBids.list":

type BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bid responses were considered to
// have no
// applicable bids, with the number of bid responses affected for each
// reason.
func (r *BiddersAccountsFilterSetsBidResponsesWithoutBidsService) List(filterSetName string) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c := &BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidResponsesWithoutBidsResponse.nextPageToken
// returned from the previous call to the
// bidResponsesWithoutBids.list
// method.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidResponsesWithoutBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.bidResponsesWithoutBids.list" call.
// Exactly one of *ListBidResponsesWithoutBidsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListBidResponsesWithoutBidsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) Do(opts ...googleapi.CallOption) (*ListBidResponsesWithoutBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidResponsesWithoutBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bid responses were considered to have no\napplicable bids, with the number of bid responses affected for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/bidResponsesWithoutBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.bidResponsesWithoutBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidResponsesWithoutBidsResponse.nextPageToken\nreturned from the previous call to the bidResponsesWithoutBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidResponsesWithoutBids",
	//   "response": {
	//     "$ref": "ListBidResponsesWithoutBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsBidResponsesWithoutBidsListCall) Pages(ctx context.Context, f func(*ListBidResponsesWithoutBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.filteredBidRequests.list":

type BiddersAccountsFilterSetsFilteredBidRequestsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons that caused a bid request not to be sent for
// an
// impression, with the number of bid requests not sent for each reason.
func (r *BiddersAccountsFilterSetsFilteredBidRequestsService) List(filterSetName string) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c := &BiddersAccountsFilterSetsFilteredBidRequestsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilteredBidRequestsResponse.nextPageToken
// returned from the previous call to the
// filteredBidRequests.list
// method.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsFilteredBidRequestsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBidRequests")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.filteredBidRequests.list" call.
// Exactly one of *ListFilteredBidRequestsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListFilteredBidRequestsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) Do(opts ...googleapi.CallOption) (*ListFilteredBidRequestsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilteredBidRequestsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons that caused a bid request not to be sent for an\nimpression, with the number of bid requests not sent for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/filteredBidRequests",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.filteredBidRequests.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilteredBidRequestsResponse.nextPageToken\nreturned from the previous call to the filteredBidRequests.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBidRequests",
	//   "response": {
	//     "$ref": "ListFilteredBidRequestsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsFilteredBidRequestsListCall) Pages(ctx context.Context, f func(*ListFilteredBidRequestsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.list":

type BiddersAccountsFilterSetsFilteredBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bids were filtered, with the number
// of bids
// filtered for each reason.
func (r *BiddersAccountsFilterSetsFilteredBidsService) List(filterSetName string) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c := &BiddersAccountsFilterSetsFilteredBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilteredBidsResponse.nextPageToken
// returned from the previous call to the filteredBids.list
// method.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsFilteredBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsFilteredBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.list" call.
// Exactly one of *ListFilteredBidsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListFilteredBidsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) Do(opts ...googleapi.CallOption) (*ListFilteredBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilteredBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bids were filtered, with the number of bids\nfiltered for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/filteredBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilteredBidsResponse.nextPageToken\nreturned from the previous call to the filteredBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids",
	//   "response": {
	//     "$ref": "ListFilteredBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsFilteredBidsListCall) Pages(ctx context.Context, f func(*ListFilteredBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.creatives.list":

type BiddersAccountsFilterSetsFilteredBidsCreativesListCall struct {
	s                *Service
	filterSetName    string
	creativeStatusId int64
	urlParams_       gensupport.URLParams
	ifNoneMatch_     string
	ctx_             context.Context
	header_          http.Header
}

// List: List all creatives associated with a specific reason for which
// bids were
// filtered, with the number of bids filtered for each creative.
func (r *BiddersAccountsFilterSetsFilteredBidsCreativesService) List(filterSetName string, creativeStatusId int64) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c := &BiddersAccountsFilterSetsFilteredBidsCreativesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	c.creativeStatusId = creativeStatusId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListCreativeStatusBreakdownByCreativeResponse.nextPageToken
// returne
// d from the previous call to the filteredBids.creatives.list
// method.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsFilteredBidsCreativesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/creatives")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName":    c.filterSetName,
		"creativeStatusId": strconv.FormatInt(c.creativeStatusId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.creatives.list" call.
// Exactly one of *ListCreativeStatusBreakdownByCreativeResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListCreativeStatusBreakdownByCreativeResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) Do(opts ...googleapi.CallOption) (*ListCreativeStatusBreakdownByCreativeResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCreativeStatusBreakdownByCreativeResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all creatives associated with a specific reason for which bids were\nfiltered, with the number of bids filtered for each creative.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/filteredBids/{creativeStatusId}/creatives",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.creatives.list",
	//   "parameterOrder": [
	//     "filterSetName",
	//     "creativeStatusId"
	//   ],
	//   "parameters": {
	//     "creativeStatusId": {
	//       "description": "The ID of the creative status for which to retrieve a breakdown by\ncreative.\nSee\n[creative-status-codes](https://developers.google.com/authorized-buyers/rtb/downloads/creative-status-codes).",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListCreativeStatusBreakdownByCreativeResponse.nextPageToken\nreturned from the previous call to the filteredBids.creatives.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/creatives",
	//   "response": {
	//     "$ref": "ListCreativeStatusBreakdownByCreativeResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsFilteredBidsCreativesListCall) Pages(ctx context.Context, f func(*ListCreativeStatusBreakdownByCreativeResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.details.list":

type BiddersAccountsFilterSetsFilteredBidsDetailsListCall struct {
	s                *Service
	filterSetName    string
	creativeStatusId int64
	urlParams_       gensupport.URLParams
	ifNoneMatch_     string
	ctx_             context.Context
	header_          http.Header
}

// List: List all details associated with a specific reason for which
// bids were
// filtered, with the number of bids filtered for each detail.
func (r *BiddersAccountsFilterSetsFilteredBidsDetailsService) List(filterSetName string, creativeStatusId int64) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c := &BiddersAccountsFilterSetsFilteredBidsDetailsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	c.creativeStatusId = creativeStatusId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListCreativeStatusBreakdownByDetailResponse.nextPageToken
// returned from the previous call to the
// filteredBids.details.list
// method.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsFilteredBidsDetailsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/details")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName":    c.filterSetName,
		"creativeStatusId": strconv.FormatInt(c.creativeStatusId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.details.list" call.
// Exactly one of *ListCreativeStatusBreakdownByDetailResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListCreativeStatusBreakdownByDetailResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) Do(opts ...googleapi.CallOption) (*ListCreativeStatusBreakdownByDetailResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCreativeStatusBreakdownByDetailResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all details associated with a specific reason for which bids were\nfiltered, with the number of bids filtered for each detail.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/filteredBids/{creativeStatusId}/details",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.filteredBids.details.list",
	//   "parameterOrder": [
	//     "filterSetName",
	//     "creativeStatusId"
	//   ],
	//   "parameters": {
	//     "creativeStatusId": {
	//       "description": "The ID of the creative status for which to retrieve a breakdown by detail.\nSee\n[creative-status-codes](https://developers.google.com/authorized-buyers/rtb/downloads/creative-status-codes).\nDetails are only available for statuses 10, 14, 15, 17, 18, 19, 86, and 87.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListCreativeStatusBreakdownByDetailResponse.nextPageToken\nreturned from the previous call to the filteredBids.details.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/details",
	//   "response": {
	//     "$ref": "ListCreativeStatusBreakdownByDetailResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsFilteredBidsDetailsListCall) Pages(ctx context.Context, f func(*ListCreativeStatusBreakdownByDetailResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.impressionMetrics.list":

type BiddersAccountsFilterSetsImpressionMetricsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: Lists all metrics that are measured in terms of number of
// impressions.
func (r *BiddersAccountsFilterSetsImpressionMetricsService) List(filterSetName string) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c := &BiddersAccountsFilterSetsImpressionMetricsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListImpressionMetricsResponse.nextPageToken
// returned from the previous call to the impressionMetrics.list
// method.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsImpressionMetricsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/impressionMetrics")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.impressionMetrics.list" call.
// Exactly one of *ListImpressionMetricsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListImpressionMetricsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) Do(opts ...googleapi.CallOption) (*ListImpressionMetricsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListImpressionMetricsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all metrics that are measured in terms of number of impressions.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/impressionMetrics",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.impressionMetrics.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListImpressionMetricsResponse.nextPageToken\nreturned from the previous call to the impressionMetrics.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/impressionMetrics",
	//   "response": {
	//     "$ref": "ListImpressionMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsImpressionMetricsListCall) Pages(ctx context.Context, f func(*ListImpressionMetricsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.losingBids.list":

type BiddersAccountsFilterSetsLosingBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bids lost in the auction, with the
// number of
// bids that lost for each reason.
func (r *BiddersAccountsFilterSetsLosingBidsService) List(filterSetName string) *BiddersAccountsFilterSetsLosingBidsListCall {
	c := &BiddersAccountsFilterSetsLosingBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsLosingBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListLosingBidsResponse.nextPageToken
// returned from the previous call to the losingBids.list
// method.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsLosingBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsLosingBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsLosingBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsLosingBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsLosingBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/losingBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.losingBids.list" call.
// Exactly one of *ListLosingBidsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLosingBidsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) Do(opts ...googleapi.CallOption) (*ListLosingBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListLosingBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bids lost in the auction, with the number of\nbids that lost for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/losingBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.losingBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListLosingBidsResponse.nextPageToken\nreturned from the previous call to the losingBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/losingBids",
	//   "response": {
	//     "$ref": "ListLosingBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsLosingBidsListCall) Pages(ctx context.Context, f func(*ListLosingBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.accounts.filterSets.nonBillableWinningBids.list":

type BiddersAccountsFilterSetsNonBillableWinningBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which winning bids were not billable, with
// the number
// of bids not billed for each reason.
func (r *BiddersAccountsFilterSetsNonBillableWinningBidsService) List(filterSetName string) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c := &BiddersAccountsFilterSetsNonBillableWinningBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) PageSize(pageSize int64) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListNonBillableWinningBidsResponse.nextPageToken
// returned from the previous call to the
// nonBillableWinningBids.list
// method.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) PageToken(pageToken string) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) Fields(s ...googleapi.Field) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) IfNoneMatch(entityTag string) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) Context(ctx context.Context) *BiddersAccountsFilterSetsNonBillableWinningBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/nonBillableWinningBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.accounts.filterSets.nonBillableWinningBids.list" call.
// Exactly one of *ListNonBillableWinningBidsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListNonBillableWinningBidsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) Do(opts ...googleapi.CallOption) (*ListNonBillableWinningBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListNonBillableWinningBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which winning bids were not billable, with the number\nof bids not billed for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/accounts/{accountsId}/filterSets/{filterSetsId}/nonBillableWinningBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.accounts.filterSets.nonBillableWinningBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/accounts/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListNonBillableWinningBidsResponse.nextPageToken\nreturned from the previous call to the nonBillableWinningBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/nonBillableWinningBids",
	//   "response": {
	//     "$ref": "ListNonBillableWinningBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersAccountsFilterSetsNonBillableWinningBidsListCall) Pages(ctx context.Context, f func(*ListNonBillableWinningBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.create":

type BiddersFilterSetsCreateCall struct {
	s          *Service
	ownerName  string
	filterset  *FilterSet
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates the specified filter set for the account with the
// given account ID.
func (r *BiddersFilterSetsService) Create(ownerName string, filterset *FilterSet) *BiddersFilterSetsCreateCall {
	c := &BiddersFilterSetsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.ownerName = ownerName
	c.filterset = filterset
	return c
}

// IsTransient sets the optional parameter "isTransient": Whether the
// filter set is transient, or should be persisted indefinitely.
// By default, filter sets are not transient.
// If transient, it will be available for at least 1 hour after
// creation.
func (c *BiddersFilterSetsCreateCall) IsTransient(isTransient bool) *BiddersFilterSetsCreateCall {
	c.urlParams_.Set("isTransient", fmt.Sprint(isTransient))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsCreateCall) Fields(s ...googleapi.Field) *BiddersFilterSetsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsCreateCall) Context(ctx context.Context) *BiddersFilterSetsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.filterset)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+ownerName}/filterSets")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"ownerName": c.ownerName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.create" call.
// Exactly one of *FilterSet or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *FilterSet.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *BiddersFilterSetsCreateCall) Do(opts ...googleapi.CallOption) (*FilterSet, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &FilterSet{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified filter set for the account with the given account ID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets",
	//   "httpMethod": "POST",
	//   "id": "adexchangebuyer2.bidders.filterSets.create",
	//   "parameterOrder": [
	//     "ownerName"
	//   ],
	//   "parameters": {
	//     "isTransient": {
	//       "description": "Whether the filter set is transient, or should be persisted indefinitely.\nBy default, filter sets are not transient.\nIf transient, it will be available for at least 1 hour after creation.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ownerName": {
	//       "description": "Name of the owner (bidder or account) of the filter set to be created.\nFor example:\n\n- For a bidder-level filter set for bidder 123: `bidders/123`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+ownerName}/filterSets",
	//   "request": {
	//     "$ref": "FilterSet"
	//   },
	//   "response": {
	//     "$ref": "FilterSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.filterSets.delete":

type BiddersFilterSetsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the requested filter set from the account with the
// given account
// ID.
func (r *BiddersFilterSetsService) Delete(name string) *BiddersFilterSetsDeleteCall {
	c := &BiddersFilterSetsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsDeleteCall) Fields(s ...googleapi.Field) *BiddersFilterSetsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsDeleteCall) Context(ctx context.Context) *BiddersFilterSetsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BiddersFilterSetsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the requested filter set from the account with the given account\nID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}",
	//   "httpMethod": "DELETE",
	//   "id": "adexchangebuyer2.bidders.filterSets.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Full name of the resource to delete.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.filterSets.get":

type BiddersFilterSetsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the requested filter set for the account with the
// given account
// ID.
func (r *BiddersFilterSetsService) Get(name string) *BiddersFilterSetsGetCall {
	c := &BiddersFilterSetsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsGetCall) Fields(s ...googleapi.Field) *BiddersFilterSetsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsGetCall) IfNoneMatch(entityTag string) *BiddersFilterSetsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsGetCall) Context(ctx context.Context) *BiddersFilterSetsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.get" call.
// Exactly one of *FilterSet or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *FilterSet.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *BiddersFilterSetsGetCall) Do(opts ...googleapi.CallOption) (*FilterSet, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &FilterSet{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the requested filter set for the account with the given account\nID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Full name of the resource being requested.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "FilterSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// method id "adexchangebuyer2.bidders.filterSets.list":

type BiddersFilterSetsListCall struct {
	s            *Service
	ownerName    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all filter sets for the account with the given account
// ID.
func (r *BiddersFilterSetsService) List(ownerName string) *BiddersFilterSetsListCall {
	c := &BiddersFilterSetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.ownerName = ownerName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsListCall) PageSize(pageSize int64) *BiddersFilterSetsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilterSetsResponse.nextPageToken
// returned from the previous call to
// the
// accounts.filterSets.list
// method.
func (c *BiddersFilterSetsListCall) PageToken(pageToken string) *BiddersFilterSetsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsListCall) Context(ctx context.Context) *BiddersFilterSetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+ownerName}/filterSets")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"ownerName": c.ownerName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.list" call.
// Exactly one of *ListFilterSetsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListFilterSetsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsListCall) Do(opts ...googleapi.CallOption) (*ListFilterSetsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilterSetsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all filter sets for the account with the given account ID.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.list",
	//   "parameterOrder": [
	//     "ownerName"
	//   ],
	//   "parameters": {
	//     "ownerName": {
	//       "description": "Name of the owner (bidder or account) of the filter sets to be listed.\nFor example:\n\n- For a bidder-level filter set for bidder 123: `bidders/123`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilterSetsResponse.nextPageToken\nreturned from the previous call to the\naccounts.filterSets.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+ownerName}/filterSets",
	//   "response": {
	//     "$ref": "ListFilterSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsListCall) Pages(ctx context.Context, f func(*ListFilterSetsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.bidMetrics.list":

type BiddersFilterSetsBidMetricsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: Lists all metrics that are measured in terms of number of bids.
func (r *BiddersFilterSetsBidMetricsService) List(filterSetName string) *BiddersFilterSetsBidMetricsListCall {
	c := &BiddersFilterSetsBidMetricsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsBidMetricsListCall) PageSize(pageSize int64) *BiddersFilterSetsBidMetricsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidMetricsResponse.nextPageToken
// returned from the previous call to the bidMetrics.list
// method.
func (c *BiddersFilterSetsBidMetricsListCall) PageToken(pageToken string) *BiddersFilterSetsBidMetricsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsBidMetricsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsBidMetricsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsBidMetricsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsBidMetricsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsBidMetricsListCall) Context(ctx context.Context) *BiddersFilterSetsBidMetricsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsBidMetricsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsBidMetricsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidMetrics")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.bidMetrics.list" call.
// Exactly one of *ListBidMetricsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListBidMetricsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsBidMetricsListCall) Do(opts ...googleapi.CallOption) (*ListBidMetricsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidMetricsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all metrics that are measured in terms of number of bids.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/bidMetrics",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.bidMetrics.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidMetricsResponse.nextPageToken\nreturned from the previous call to the bidMetrics.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidMetrics",
	//   "response": {
	//     "$ref": "ListBidMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsBidMetricsListCall) Pages(ctx context.Context, f func(*ListBidMetricsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.bidResponseErrors.list":

type BiddersFilterSetsBidResponseErrorsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all errors that occurred in bid responses, with the number
// of bid
// responses affected for each reason.
func (r *BiddersFilterSetsBidResponseErrorsService) List(filterSetName string) *BiddersFilterSetsBidResponseErrorsListCall {
	c := &BiddersFilterSetsBidResponseErrorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsBidResponseErrorsListCall) PageSize(pageSize int64) *BiddersFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidResponseErrorsResponse.nextPageToken
// returned from the previous call to the bidResponseErrors.list
// method.
func (c *BiddersFilterSetsBidResponseErrorsListCall) PageToken(pageToken string) *BiddersFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsBidResponseErrorsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsBidResponseErrorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsBidResponseErrorsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsBidResponseErrorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsBidResponseErrorsListCall) Context(ctx context.Context) *BiddersFilterSetsBidResponseErrorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsBidResponseErrorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsBidResponseErrorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidResponseErrors")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.bidResponseErrors.list" call.
// Exactly one of *ListBidResponseErrorsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListBidResponseErrorsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsBidResponseErrorsListCall) Do(opts ...googleapi.CallOption) (*ListBidResponseErrorsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidResponseErrorsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all errors that occurred in bid responses, with the number of bid\nresponses affected for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/bidResponseErrors",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.bidResponseErrors.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidResponseErrorsResponse.nextPageToken\nreturned from the previous call to the bidResponseErrors.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidResponseErrors",
	//   "response": {
	//     "$ref": "ListBidResponseErrorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsBidResponseErrorsListCall) Pages(ctx context.Context, f func(*ListBidResponseErrorsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.bidResponsesWithoutBids.list":

type BiddersFilterSetsBidResponsesWithoutBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bid responses were considered to
// have no
// applicable bids, with the number of bid responses affected for each
// reason.
func (r *BiddersFilterSetsBidResponsesWithoutBidsService) List(filterSetName string) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c := &BiddersFilterSetsBidResponsesWithoutBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) PageSize(pageSize int64) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListBidResponsesWithoutBidsResponse.nextPageToken
// returned from the previous call to the
// bidResponsesWithoutBids.list
// method.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) PageToken(pageToken string) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) Context(ctx context.Context) *BiddersFilterSetsBidResponsesWithoutBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/bidResponsesWithoutBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.bidResponsesWithoutBids.list" call.
// Exactly one of *ListBidResponsesWithoutBidsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListBidResponsesWithoutBidsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) Do(opts ...googleapi.CallOption) (*ListBidResponsesWithoutBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBidResponsesWithoutBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bid responses were considered to have no\napplicable bids, with the number of bid responses affected for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/bidResponsesWithoutBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.bidResponsesWithoutBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListBidResponsesWithoutBidsResponse.nextPageToken\nreturned from the previous call to the bidResponsesWithoutBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/bidResponsesWithoutBids",
	//   "response": {
	//     "$ref": "ListBidResponsesWithoutBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsBidResponsesWithoutBidsListCall) Pages(ctx context.Context, f func(*ListBidResponsesWithoutBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.filteredBidRequests.list":

type BiddersFilterSetsFilteredBidRequestsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons that caused a bid request not to be sent for
// an
// impression, with the number of bid requests not sent for each reason.
func (r *BiddersFilterSetsFilteredBidRequestsService) List(filterSetName string) *BiddersFilterSetsFilteredBidRequestsListCall {
	c := &BiddersFilterSetsFilteredBidRequestsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) PageSize(pageSize int64) *BiddersFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilteredBidRequestsResponse.nextPageToken
// returned from the previous call to the
// filteredBidRequests.list
// method.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) PageToken(pageToken string) *BiddersFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsFilteredBidRequestsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsFilteredBidRequestsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) Context(ctx context.Context) *BiddersFilterSetsFilteredBidRequestsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsFilteredBidRequestsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBidRequests")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.filteredBidRequests.list" call.
// Exactly one of *ListFilteredBidRequestsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListFilteredBidRequestsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) Do(opts ...googleapi.CallOption) (*ListFilteredBidRequestsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilteredBidRequestsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons that caused a bid request not to be sent for an\nimpression, with the number of bid requests not sent for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/filteredBidRequests",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.filteredBidRequests.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilteredBidRequestsResponse.nextPageToken\nreturned from the previous call to the filteredBidRequests.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBidRequests",
	//   "response": {
	//     "$ref": "ListFilteredBidRequestsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsFilteredBidRequestsListCall) Pages(ctx context.Context, f func(*ListFilteredBidRequestsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.filteredBids.list":

type BiddersFilterSetsFilteredBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bids were filtered, with the number
// of bids
// filtered for each reason.
func (r *BiddersFilterSetsFilteredBidsService) List(filterSetName string) *BiddersFilterSetsFilteredBidsListCall {
	c := &BiddersFilterSetsFilteredBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsFilteredBidsListCall) PageSize(pageSize int64) *BiddersFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListFilteredBidsResponse.nextPageToken
// returned from the previous call to the filteredBids.list
// method.
func (c *BiddersFilterSetsFilteredBidsListCall) PageToken(pageToken string) *BiddersFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsFilteredBidsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsFilteredBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsFilteredBidsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsFilteredBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsFilteredBidsListCall) Context(ctx context.Context) *BiddersFilterSetsFilteredBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsFilteredBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsFilteredBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.filteredBids.list" call.
// Exactly one of *ListFilteredBidsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListFilteredBidsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsFilteredBidsListCall) Do(opts ...googleapi.CallOption) (*ListFilteredBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListFilteredBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bids were filtered, with the number of bids\nfiltered for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/filteredBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.filteredBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListFilteredBidsResponse.nextPageToken\nreturned from the previous call to the filteredBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids",
	//   "response": {
	//     "$ref": "ListFilteredBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsFilteredBidsListCall) Pages(ctx context.Context, f func(*ListFilteredBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.filteredBids.creatives.list":

type BiddersFilterSetsFilteredBidsCreativesListCall struct {
	s                *Service
	filterSetName    string
	creativeStatusId int64
	urlParams_       gensupport.URLParams
	ifNoneMatch_     string
	ctx_             context.Context
	header_          http.Header
}

// List: List all creatives associated with a specific reason for which
// bids were
// filtered, with the number of bids filtered for each creative.
func (r *BiddersFilterSetsFilteredBidsCreativesService) List(filterSetName string, creativeStatusId int64) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c := &BiddersFilterSetsFilteredBidsCreativesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	c.creativeStatusId = creativeStatusId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) PageSize(pageSize int64) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListCreativeStatusBreakdownByCreativeResponse.nextPageToken
// returne
// d from the previous call to the filteredBids.creatives.list
// method.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) PageToken(pageToken string) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) Context(ctx context.Context) *BiddersFilterSetsFilteredBidsCreativesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsFilteredBidsCreativesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/creatives")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName":    c.filterSetName,
		"creativeStatusId": strconv.FormatInt(c.creativeStatusId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.filteredBids.creatives.list" call.
// Exactly one of *ListCreativeStatusBreakdownByCreativeResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListCreativeStatusBreakdownByCreativeResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) Do(opts ...googleapi.CallOption) (*ListCreativeStatusBreakdownByCreativeResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCreativeStatusBreakdownByCreativeResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all creatives associated with a specific reason for which bids were\nfiltered, with the number of bids filtered for each creative.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/filteredBids/{creativeStatusId}/creatives",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.filteredBids.creatives.list",
	//   "parameterOrder": [
	//     "filterSetName",
	//     "creativeStatusId"
	//   ],
	//   "parameters": {
	//     "creativeStatusId": {
	//       "description": "The ID of the creative status for which to retrieve a breakdown by\ncreative.\nSee\n[creative-status-codes](https://developers.google.com/authorized-buyers/rtb/downloads/creative-status-codes).",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListCreativeStatusBreakdownByCreativeResponse.nextPageToken\nreturned from the previous call to the filteredBids.creatives.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/creatives",
	//   "response": {
	//     "$ref": "ListCreativeStatusBreakdownByCreativeResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsFilteredBidsCreativesListCall) Pages(ctx context.Context, f func(*ListCreativeStatusBreakdownByCreativeResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.filteredBids.details.list":

type BiddersFilterSetsFilteredBidsDetailsListCall struct {
	s                *Service
	filterSetName    string
	creativeStatusId int64
	urlParams_       gensupport.URLParams
	ifNoneMatch_     string
	ctx_             context.Context
	header_          http.Header
}

// List: List all details associated with a specific reason for which
// bids were
// filtered, with the number of bids filtered for each detail.
func (r *BiddersFilterSetsFilteredBidsDetailsService) List(filterSetName string, creativeStatusId int64) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c := &BiddersFilterSetsFilteredBidsDetailsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	c.creativeStatusId = creativeStatusId
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) PageSize(pageSize int64) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListCreativeStatusBreakdownByDetailResponse.nextPageToken
// returned from the previous call to the
// filteredBids.details.list
// method.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) PageToken(pageToken string) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) Context(ctx context.Context) *BiddersFilterSetsFilteredBidsDetailsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsFilteredBidsDetailsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/details")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName":    c.filterSetName,
		"creativeStatusId": strconv.FormatInt(c.creativeStatusId, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.filteredBids.details.list" call.
// Exactly one of *ListCreativeStatusBreakdownByDetailResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListCreativeStatusBreakdownByDetailResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) Do(opts ...googleapi.CallOption) (*ListCreativeStatusBreakdownByDetailResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCreativeStatusBreakdownByDetailResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all details associated with a specific reason for which bids were\nfiltered, with the number of bids filtered for each detail.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/filteredBids/{creativeStatusId}/details",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.filteredBids.details.list",
	//   "parameterOrder": [
	//     "filterSetName",
	//     "creativeStatusId"
	//   ],
	//   "parameters": {
	//     "creativeStatusId": {
	//       "description": "The ID of the creative status for which to retrieve a breakdown by detail.\nSee\n[creative-status-codes](https://developers.google.com/authorized-buyers/rtb/downloads/creative-status-codes).\nDetails are only available for statuses 10, 14, 15, 17, 18, 19, 86, and 87.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListCreativeStatusBreakdownByDetailResponse.nextPageToken\nreturned from the previous call to the filteredBids.details.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/filteredBids/{creativeStatusId}/details",
	//   "response": {
	//     "$ref": "ListCreativeStatusBreakdownByDetailResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsFilteredBidsDetailsListCall) Pages(ctx context.Context, f func(*ListCreativeStatusBreakdownByDetailResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.impressionMetrics.list":

type BiddersFilterSetsImpressionMetricsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: Lists all metrics that are measured in terms of number of
// impressions.
func (r *BiddersFilterSetsImpressionMetricsService) List(filterSetName string) *BiddersFilterSetsImpressionMetricsListCall {
	c := &BiddersFilterSetsImpressionMetricsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsImpressionMetricsListCall) PageSize(pageSize int64) *BiddersFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListImpressionMetricsResponse.nextPageToken
// returned from the previous call to the impressionMetrics.list
// method.
func (c *BiddersFilterSetsImpressionMetricsListCall) PageToken(pageToken string) *BiddersFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsImpressionMetricsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsImpressionMetricsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsImpressionMetricsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsImpressionMetricsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsImpressionMetricsListCall) Context(ctx context.Context) *BiddersFilterSetsImpressionMetricsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsImpressionMetricsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsImpressionMetricsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/impressionMetrics")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.impressionMetrics.list" call.
// Exactly one of *ListImpressionMetricsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListImpressionMetricsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsImpressionMetricsListCall) Do(opts ...googleapi.CallOption) (*ListImpressionMetricsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListImpressionMetricsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all metrics that are measured in terms of number of impressions.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/impressionMetrics",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.impressionMetrics.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListImpressionMetricsResponse.nextPageToken\nreturned from the previous call to the impressionMetrics.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/impressionMetrics",
	//   "response": {
	//     "$ref": "ListImpressionMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsImpressionMetricsListCall) Pages(ctx context.Context, f func(*ListImpressionMetricsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.losingBids.list":

type BiddersFilterSetsLosingBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which bids lost in the auction, with the
// number of
// bids that lost for each reason.
func (r *BiddersFilterSetsLosingBidsService) List(filterSetName string) *BiddersFilterSetsLosingBidsListCall {
	c := &BiddersFilterSetsLosingBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsLosingBidsListCall) PageSize(pageSize int64) *BiddersFilterSetsLosingBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListLosingBidsResponse.nextPageToken
// returned from the previous call to the losingBids.list
// method.
func (c *BiddersFilterSetsLosingBidsListCall) PageToken(pageToken string) *BiddersFilterSetsLosingBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsLosingBidsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsLosingBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsLosingBidsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsLosingBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsLosingBidsListCall) Context(ctx context.Context) *BiddersFilterSetsLosingBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsLosingBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsLosingBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/losingBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.losingBids.list" call.
// Exactly one of *ListLosingBidsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLosingBidsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BiddersFilterSetsLosingBidsListCall) Do(opts ...googleapi.CallOption) (*ListLosingBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListLosingBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which bids lost in the auction, with the number of\nbids that lost for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/losingBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.losingBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListLosingBidsResponse.nextPageToken\nreturned from the previous call to the losingBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/losingBids",
	//   "response": {
	//     "$ref": "ListLosingBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsLosingBidsListCall) Pages(ctx context.Context, f func(*ListLosingBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "adexchangebuyer2.bidders.filterSets.nonBillableWinningBids.list":

type BiddersFilterSetsNonBillableWinningBidsListCall struct {
	s             *Service
	filterSetName string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// List: List all reasons for which winning bids were not billable, with
// the number
// of bids not billed for each reason.
func (r *BiddersFilterSetsNonBillableWinningBidsService) List(filterSetName string) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c := &BiddersFilterSetsNonBillableWinningBidsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filterSetName = filterSetName
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The server may return fewer results than requested.
// If unspecified, the server will pick an appropriate default.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) PageSize(pageSize int64) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
// Typically, this is the value
// of
// ListNonBillableWinningBidsResponse.nextPageToken
// returned from the previous call to the
// nonBillableWinningBids.list
// method.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) PageToken(pageToken string) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) Fields(s ...googleapi.Field) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) IfNoneMatch(entityTag string) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) Context(ctx context.Context) *BiddersFilterSetsNonBillableWinningBidsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BiddersFilterSetsNonBillableWinningBidsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+filterSetName}/nonBillableWinningBids")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"filterSetName": c.filterSetName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexchangebuyer2.bidders.filterSets.nonBillableWinningBids.list" call.
// Exactly one of *ListNonBillableWinningBidsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListNonBillableWinningBidsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) Do(opts ...googleapi.CallOption) (*ListNonBillableWinningBidsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListNonBillableWinningBidsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all reasons for which winning bids were not billable, with the number\nof bids not billed for each reason.",
	//   "flatPath": "v2beta1/bidders/{biddersId}/filterSets/{filterSetsId}/nonBillableWinningBids",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer2.bidders.filterSets.nonBillableWinningBids.list",
	//   "parameterOrder": [
	//     "filterSetName"
	//   ],
	//   "parameters": {
	//     "filterSetName": {
	//       "description": "Name of the filter set that should be applied to the requested metrics.\nFor example:\n\n- For a bidder-level filter set for bidder 123:\n  `bidders/123/filterSets/abc`\n\n- For an account-level filter set for the buyer account representing bidder\n  123: `bidders/123/accounts/123/filterSets/abc`\n\n- For an account-level filter set for the child seat buyer account 456\n  whose bidder is 123: `bidders/123/accounts/456/filterSets/abc`",
	//       "location": "path",
	//       "pattern": "^bidders/[^/]+/filterSets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. The server may return fewer results than requested.\nIf unspecified, the server will pick an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.\nTypically, this is the value of\nListNonBillableWinningBidsResponse.nextPageToken\nreturned from the previous call to the nonBillableWinningBids.list\nmethod.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+filterSetName}/nonBillableWinningBids",
	//   "response": {
	//     "$ref": "ListNonBillableWinningBidsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.buyer"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BiddersFilterSetsNonBillableWinningBidsListCall) Pages(ctx context.Context, f func(*ListNonBillableWinningBidsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
