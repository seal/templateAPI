package types

type IpLoopup struct {
	IP              string `json:"ip"`
	IPNumber        string `json:"ip_number"`
	IPVersion       int    `json:"ip_version"`
	CountryName     string `json:"country_name"`
	CountryCode2    string `json:"country_code2"`
	Isp             string `json:"isp"`
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
type TbsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Filters []struct {
		Type    string `json:"type"`
		Options []struct {
			Text string `json:"text"`
			Tbs  string `json:"tbs"`
		} `json:"options"`
	} `json:"filters"`

	ShoppingResults []struct { // Changed to [], might break it lol
		Position          int      `json:"position"`
		Title             string   `json:"title"`
		Link              string   `json:"link"`
		ProductLink       string   `json:"product_link"`
		ProductID         string   `json:"product_id"`
		SerpapiProductAPI string   `json:"serpapi_product_api"`
		Source            string   `json:"source"`
		Price             string   `json:"price"`
		ExtractedPrice    float64  `json:"extracted_price"`
		Rating            float64  `json:"rating"`
		Reviews           int      `json:"reviews"`
		Extensions        []string `json:"extensions"`
		Thumbnail         string   `json:"thumbnail"`
		Tag               string   `json:"tag,omitempty"`
		Delivery          string   `json:"delivery,omitempty"`
	}
}
type ShoppingResults []struct {
	Position          int      `json:"position"`
	Title             string   `json:"title"`
	Link              string   `json:"link"`
	ProductLink       string   `json:"product_link"`
	ProductID         string   `json:"product_id"`
	SerpapiProductAPI string   `json:"serpapi_product_api"`
	Source            string   `json:"source"`
	Price             string   `json:"price"`
	ExtractedPrice    float64  `json:"extracted_price"`
	Rating            float64  `json:"rating"`
	Reviews           int      `json:"reviews"`
	Extensions        []string `json:"extensions"`
	Thumbnail         string   `json:"thumbnail"`
	Tag               string   `json:"tag,omitempty"`
	Delivery          string   `json:"delivery,omitempty"`
}

type SearchResponse struct {
	SearchMetadata struct {
		ID             string  `json:"id"`
		Status         string  `json:"status"`
		JSONEndpoint   string  `json:"json_endpoint"`
		CreatedAt      string  `json:"created_at"`
		ProcessedAt    string  `json:"processed_at"`
		GoogleURL      string  `json:"google_url"`
		RawHTMLFile    string  `json:"raw_html_file"`
		TotalTimeTaken float64 `json:"total_time_taken"`
	} `json:"search_metadata"`
	SearchParameters struct {
		Engine       string `json:"engine"`
		Q            string `json:"q"`
		GoogleDomain string `json:"google_domain"`
		Device       string `json:"device"`
		Tbm          string `json:"tbm"`
	} `json:"search_parameters"`
	SearchInformation struct {
		ShoppingResultsState string `json:"shopping_results_state"`
		QueryDisplayed       string `json:"query_displayed"`
		MenuItems            []struct {
			Position    int    `json:"position"`
			Title       string `json:"title"`
			Link        string `json:"link,omitempty"`
			SerpapiLink string `json:"serpapi_link,omitempty"`
		} `json:"menu_items"`
	} `json:"search_information"`
	Filters []struct {
		Type    string `json:"type"`
		Options []struct {
			Text string `json:"text"`
			Tbs  string `json:"tbs"`
		} `json:"options"`
	} `json:"filters"`
	InlineShoppingResults []struct {
		Position       int      `json:"position"`
		BlockPosition  string   `json:"block_position,omitempty"`
		Title          string   `json:"title"`
		Price          string   `json:"price"`
		ExtractedPrice float64  `json:"extracted_price"`
		Link           string   `json:"link"`
		Source         string   `json:"source"`
		Rating         float64  `json:"rating,omitempty"`
		Reviews        int      `json:"reviews,omitempty"`
		Thumbnail      string   `json:"thumbnail"`
		Extensions     []string `json:"extensions,omitempty"`
		Shipping       string   `json:"shipping,omitempty"`
	} `json:"inline_shopping_results"`
	ShoppingResults []struct {
		Position          int      `json:"position"`
		Title             string   `json:"title"`
		Link              string   `json:"link"`
		ProductLink       string   `json:"product_link"`
		ProductID         string   `json:"product_id"`
		SerpapiProductAPI string   `json:"serpapi_product_api"`
		Source            string   `json:"source"`
		Price             string   `json:"price"`
		ExtractedPrice    float64  `json:"extracted_price"`
		Rating            float64  `json:"rating"`
		Reviews           int      `json:"reviews"`
		Extensions        []string `json:"extensions"`
		Thumbnail         string   `json:"thumbnail"`
		Tag               string   `json:"tag,omitempty"`
		Delivery          string   `json:"delivery,omitempty"`
	} `json:"shopping_results"`
	Pagination struct {
		Current    int    `json:"current"`
		Next       string `json:"next"`
		OtherPages struct {
			Num2 string `json:"2"`
			Num3 string `json:"3"`
			Num4 string `json:"4"`
			Num5 string `json:"5"`
			Num6 string `json:"6"`
			Num7 string `json:"7"`
		} `json:"other_pages"`
	} `json:"pagination"`
	SerpapiPagination struct {
		Current    int    `json:"current"`
		NextLink   string `json:"next_link"`
		Next       string `json:"next"`
		OtherPages struct {
			Num2 string `json:"2"`
			Num3 string `json:"3"`
			Num4 string `json:"4"`
			Num5 string `json:"5"`
			Num6 string `json:"6"`
			Num7 string `json:"7"`
		} `json:"other_pages"`
	} `json:"serpapi_pagination"`
}
