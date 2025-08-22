package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// WebMetaTag represents the WebMetaTag schema from the OpenAPI specification
type WebMetaTag struct {
	Content string `json:"content,omitempty"` // The name of the metadata.
	Name string `json:"name,omitempty"` // The metadata.
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// WebPage represents the WebPage schema from the OpenAPI specification
type WebPage struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Text string `json:"text"` // The query string. Use this string as the query term in a new search request.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL that takes the user to the Bing search results page for the query.Only related search results include this field.
	Displaytext string `json:"displayText,omitempty"` // The display version of the query term. This version of the query term may contain special characters that highlight the search term found in the query string. The string contains the highlighting characters only if the query enabled hit highlighting
	Searchlink string `json:"searchLink,omitempty"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
}

// WebWebAnswer represents the WebWebAnswer schema from the OpenAPI specification
type WebWebAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
}

// QueryContext represents the QueryContext schema from the OpenAPI specification
type QueryContext struct {
	Adultintent bool `json:"adultIntent,omitempty"` // A Boolean value that indicates whether the specified query has adult intent. The value is true if the query has adult intent; otherwise, false.
	Alterationoverridequery string `json:"alterationOverrideQuery,omitempty"` // The query string to use to force Bing to use the original string. For example, if the query string is "saling downwind", the override query string will be "+saling downwind". Remember to encode the query string which results in "%2Bsaling+downwind". This field is included only if the original query string contains a spelling mistake.
	Alteredquery string `json:"alteredQuery,omitempty"` // The query string used by Bing to perform the query. Bing uses the altered query string if the original query string contained spelling mistakes. For example, if the query string is "saling downwind", the altered query string will be "sailing downwind". This field is included only if the original query string contains a spelling mistake.
	Originalquery string `json:"originalQuery"` // The query string as specified in the request.
}
