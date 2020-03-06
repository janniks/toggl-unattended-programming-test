package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"janniks.com/toggl/initial/model"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type createDeckResponse struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type openDeckResponse struct {
	createDeckResponse
	Cards []card `json:"cards"`
}

type drawResponse struct {
	Cards []card `json:"cards"`
}

type card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func Test_CreateDeck(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var jsonResponse map[string]json.RawMessage
	json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Contains(t, jsonResponse, "deck_id")
	assert.Contains(t, jsonResponse, "shuffled")
	assert.Contains(t, jsonResponse, "remaining")
	assert.NotContains(t, jsonResponse, "cards")

	var response createDeckResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	assert.Equal(t, false, response.Shuffled)
	assert.Equal(t, model.CardN, response.Remaining)
}

func Test_CreateDeck_MalformedCardCode(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck?cards=XX", nil)
	router.ServeHTTP(w, req)
	var errorResponse string
	json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid card code provided", errorResponse)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/deck?cards=AC,AD,", nil)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid card parameter", errorResponse)
}

func Test_CreateDeck_MalformedShuffleDefault(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck?shuffle=eurt", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, false, createResponse.Shuffled)
}

func Test_CreateOpenDeck_Standard(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId, nil)
	router.ServeHTTP(w, req)
	var openResponse openDeckResponse
	json.Unmarshal(w.Body.Bytes(), &openResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, model.CardN, len(openResponse.Cards))
}

func Test_CreateOpenDeck_Whitelist(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck?cards=AC,2D,3H,4S", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId, nil)
	router.ServeHTTP(w, req)
	var openResponse openDeckResponse
	json.Unmarshal(w.Body.Bytes(), &openResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 4, openResponse.Remaining)
	assert.Equal(t, 4, len(openResponse.Cards))

	assert.Equal(t, "AC", openResponse.Cards[0].Code)
	assert.Equal(t, "ACE", openResponse.Cards[0].Value)
	assert.Equal(t, "CLUBS", openResponse.Cards[0].Suit)

	assert.Equal(t, "2D", openResponse.Cards[1].Code)
	assert.Equal(t, "2", openResponse.Cards[1].Value)
	assert.Equal(t, "DIAMONDS", openResponse.Cards[1].Suit)

	assert.Equal(t, "3H", openResponse.Cards[2].Code)
	assert.Equal(t, "3", openResponse.Cards[2].Value)
	assert.Equal(t, "HEARTS", openResponse.Cards[2].Suit)

	assert.Equal(t, "4S", openResponse.Cards[3].Code)
	assert.Equal(t, "4", openResponse.Cards[3].Value)
	assert.Equal(t, "SPADES", openResponse.Cards[3].Suit)
}

func Test_CreateOpenDeck_Shuffle(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck?shuffle=true", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, true, createResponse.Shuffled)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId, nil)
	router.ServeHTTP(w, req)
	var openResponse openDeckResponse
	json.Unmarshal(w.Body.Bytes(), &openResponse)

	var ids []int
	for _, card := range openResponse.Cards {
		id, _ := model.CodeToId(card.Code)
		ids = append(ids, int(id))
	}

	var idsCopy []int
	copy(ids, idsCopy)
	sort.Ints(idsCopy)

	assert.NotEqual(t, idsCopy, ids)
}

func Test_OpenDeck_NotFound(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/deck/"+uuid.New().String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func Test_Draw_NotFound(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/deck/"+uuid.New().String()+"/draw", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func Test_CreateDraw_Standard(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw", nil)
	router.ServeHTTP(w, req)
	var drawResponse drawResponse
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(drawResponse.Cards))
}

func Test_CreateDraw_MalformedCount(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw?count=XYZ", nil)
	router.ServeHTTP(w, req)
	var drawResponse string
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid count parameter", drawResponse)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw?count=0", nil)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid count parameter", drawResponse)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw?count=-1", nil)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid count parameter", drawResponse)
}

func Test_CreateDraw_LessAvailable(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	router := SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/deck?cards=AC,AD,AH,AS", nil)
	router.ServeHTTP(w, req)
	var createResponse createDeckResponse
	json.Unmarshal(w.Body.Bytes(), &createResponse)
	assert.Equal(t, 4, createResponse.Remaining)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw?count=2", nil)
	router.ServeHTTP(w, req)
	var drawResponse drawResponse
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(drawResponse.Cards))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw?count=6", nil)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &drawResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(drawResponse.Cards))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw", nil)
	router.ServeHTTP(w, req)
	var drawEmptyResponse string
	json.Unmarshal(w.Body.Bytes(), &drawEmptyResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "deck is empty", drawEmptyResponse)
}
