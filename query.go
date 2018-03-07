package bing

import "net/http"

type Query struct {
	//The number of answers that you want the response to include. The answers that Bing returns are based on ranking. For example, if Bing returns webpages, images, videos, and relatedSearches for a request and you set this parameter to two (2), the response includes webpages and images.
	AnswerCount int
	//Country code for answer
	CC string
	//Number of search result
	Count int16
	//Filter seach results by ages values (Day, Week, Month)
	Freshness string
	//Country of the request sender
	Mkt string
	//Number of search results to skip
	Offset  int16
	Promote string
	//The actual research query
	Q string
	//A comma-delimited list of answers to include in the response (Computation, Entities, Images, News, RelatedSearches, SpellSuggestions, TimeZone, Videos, Webpages)
	ResponseFilter string
	//Off, Moderate Strict
	SafeSearch string
	//default languages for the interface
	SetLang string
	//active/deactive text Decoration
	TextDecoration bool
	//The type of markers to use for text decorations (Raw, HTML)
}

//Create a standart Query Object
func NewQuery(query string) *Query {
	return &Query{
		SafeSearch: "Off",
		Q:          query,
	}
}

func (query *Query) buildRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://api.cognitive.microsoft.com/bing/v7.0/search", nil)
	if err != nil {
		return req, err
	}
	query.setDefaultRequestParam(query)
	return req, nil
}

func (query *Query) setDefaultRequestParam(req *http.Request) {
	//Set GET parameters
	k := req.URL.Query()
	k.Add("q", query.Q)
	k.Add("safeSearch", query.SafeSearch)
	req.URL.RawQuery = k.Encode()
}
