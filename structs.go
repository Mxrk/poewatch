package poewatch

import "time"

//Categories is the reply struct for the getCategories function
type Categories []struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Display string `json:"display"`
	Groups  []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Display string `json:"display"`
	} `json:"groups"`
}

//ID is the reply struct for the GetChangeID function
type ID struct {
	ID string `json:"changeID"`
}

//Leagues is the reply struct for the GetLeagues function
type Leagues []struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Display   string      `json:"display"`
	Hardcore  bool        `json:"hardcore"`
	Upcoming  bool        `json:"upcoming"`
	Active    bool        `json:"active"`
	Event     bool        `json:"event"`
	Challenge bool        `json:"challenge"`
	Start     time.Time   `json:"start"`
	End       interface{} `json:"end"`
}

//ItemData is the reply struct for the GetItemData function
type ItemData []struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	Group          string `json:"group"`
	Frame          int    `json:"frame"`
	GemLevel       int    `json:"gemLevel"`
	GemQuality     int    `json:"gemQuality"`
	GemIsCorrupted bool   `json:"gemIsCorrupted"`
	Icon           string `json:"icon"`
}

//Characters is the reply struct for the GetCharacters function
type Characters struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Found time.Time `json:"found"`
	Seen  time.Time `json:"seen"`
	Chars []struct {
		ID    int       `json:"id"`
		Name  string    `json:"name"`
		Found time.Time `json:"found"`
		Seen  time.Time `json:"seen"`
	} `json:"chars"`
}

//Accounts is the reply struct for the GetAccounts function
type Accounts []struct {
	Account string    `json:"account"`
	Found   time.Time `json:"found"`
	Seen    time.Time `json:"seen"`
}

//Items is the reply struct for the GetItems function
type Items []struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	Group     string    `json:"group"`
	Frame     int       `json:"frame"`
	Variation string    `json:"variation"`
	Icon      string    `json:"icon"`
	Mean      float64   `json:"mean"`
	Median    float64   `json:"median"`
	Mode      float64   `json:"mode"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	Exalted   float64   `json:"exalted"`
	Total     int       `json:"total"`
	Daily     int       `json:"daily"`
	Current   int       `json:"current"`
	Accepted  int       `json:"accepted"`
	Change    float64   `json:"change"`
	History   []float64 `json:"history"`
}

//Compact is the reply struct for the GetCompact function
type Compact []struct {
	ID       int     `json:"id"`
	Mean     float64 `json:"mean"`
	Median   float64 `json:"median"`
	Mode     float64 `json:"mode"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Exalted  float64 `json:"exalted"`
	Total    int     `json:"total"`
	Daily    int     `json:"daily"`
	Current  int     `json:"current"`
	Accepted int     `json:"accepted"`
}

//Item is the reply struct for the GetItem function
type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Group    string `json:"group"`
	Frame    int    `json:"frame"`
	Icon     string `json:"icon"`
	Leagues  []struct {
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Display  string      `json:"display"`
		Active   bool        `json:"active"`
		Start    time.Time   `json:"start"`
		End      interface{} `json:"end"`
		Mean     float64     `json:"mean"`
		Median   float64     `json:"median"`
		Mode     float64     `json:"mode"`
		Min      float64     `json:"min"`
		Max      float64     `json:"max"`
		Exalted  float64     `json:"exalted"`
		Total    int         `json:"total"`
		Daily    int         `json:"daily"`
		Current  int         `json:"current"`
		Accepted int         `json:"accepted"`
	} `json:"leagues"`
}

//History is the reply struct for the GetItemHistory functions
type History []struct {
	Time     time.Time `json:"time"`
	Mean     float64   `json:"mean"`
	Median   float64   `json:"median"`
	Mode     float64   `json:"mode"`
	Daily    int       `json:"daily"`
	Current  int       `json:"current"`
	Accepted int       `json:"accepted"`
}

//Listings is the reply struct for the GetListings function
type Listings []struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Category   string    `json:"category"`
	Group      string    `json:"group"`
	Frame      int       `json:"frame"`
	Icon       string    `json:"icon"`
	Discovered time.Time `json:"discovered"`
	Updated    time.Time `json:"updated"`
	Count      int       `json:"count"`
	Buyout     []struct {
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
		Chaos    float64 `json:"chaos"`
		Count    int     `json:"count"`
	} `json:"buyout"`
}
