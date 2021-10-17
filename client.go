package poewatch

import (
	"encoding/json"
	"log"
)

type Client struct {
	// todo
}

var baseURL = "https://api.poe.watch/"

//GetChangeID returns the latest change ID from the top of the river and the time
func GetChangeID() ID {
	st := ID{}
	req := get(baseURL + "status")
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetLeagues list the current leagues. Event leagues appear first.
func GetLeagues() Leagues {
	st := Leagues{}
	req := get(baseURL + "leagues")
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetItemData returns all found items via the stash api
func GetItemData() ItemData {
	st := ItemData{}
	req := get(baseURL + "itemdata")
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetCharacters returns the player character names found through the stash api
func GetCharacters(name string) Characters {
	st := Characters{}
	req := get(baseURL + "account?name=" + name)
	err := json.Unmarshal([]byte(req), &st)
	if err != nil {
		log.Println(err)
	}
	return st
}

//GetAccount returns player account name via a char name
func GetAccount(name string) Accounts {
	st := Accounts{}
	req := get(baseURL + "char?name=" + name)
	err := json.Unmarshal([]byte(req), &st)
	if err != nil {
		log.Println(err)
	}
	return st
}

//GetCategories returns all listed categories
func GetCategories() Categories {
	st := Categories{}
	req := get(baseURL + "categories")
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetItems returns price and item for specified league and category.
//It's advisable to use a singular request to the compact API, rather than multiple requests to this API.
func GetItems(league, category string) Items {
	st := Items{}
	req := get(baseURL + "get?league=" + league + "&category=" + category)
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetCompact returns price data of all items provided in the active league
func GetCompact(league string) Compact {
	st := Compact{}
	req := get(baseURL + "compact?league=" + league)
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetItem returns information about a specific item.
func GetItem(itemID string) Item {
	st := Item{}
	req := get(baseURL + "id?id=" + itemID)
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetItemHistory returns prices from past leagues.
func GetItemHistory(itemID, league string) History {
	st := History{}
	req := get(baseURL + "history?id=" + itemID + "&league=" + league)
	json.Unmarshal([]byte(req), &st)
	return st
}

//GetEnchants returns Enchant data for a given Item in a league.
func GetEnchants(itemID string, league string) []Enchants {
	st := []Enchants{}
	req := get(baseURL + "enchants?id=" + itemID + "&league=" + league)
	json.Unmarshal([]byte(req), &st)
	return st
}
