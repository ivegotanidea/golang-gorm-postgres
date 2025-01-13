package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ivegotanidea/golang-gorm-postgres/initializers"
	"github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/ivegotanidea/golang-gorm-postgres/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var randomPhotos = []string{
	"https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg",
	"https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg",
	"https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg",
	"https://www.caravan.kz/wp-content/uploads/images/635838.jpg",
	"https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg",
}

type UserResponse struct {
	Status string              `json:"status"`
	Data   models.UserResponse `json:"data"`
}

type CreateProfileResponse struct {
	Status string         `json:"status"`
	Data   models.Profile `json:"data"`
}

type ProfilesResponse struct {
	Status string                   `json:"status"`
	Length int                      `json:"results"`
	Data   []models.ProfileResponse `json:"data"`
}

type ServiceResponse struct {
	Status string         `json:"status"`
	Data   models.Service `json:"data"`
}

type ServicesResponse struct {
	Status string           `json:"status"`
	Length int              `json:"results"`
	Data   []models.Service `json:"data"`
}

func ptr(v int) *int {
	return &v
}

func floatPtr(v float32) *float32 {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func getOwnerUser() models.User {
	return models.User{
		ID:             uuid.Max,
		Name:           "He Who Remains",
		Phone:          "77778889900",
		TelegramUserId: 6794234746,
		Password:       "h5sh3d", // Ensure this is hashed
		Avatar:         "https://akm-img-a-in.tosshub.com/indiatoday/images/story/202311/tom-hiddleston-in-a-still-from-loki-2-27480244-16x9_0.jpg",
		Verified:       true,
		HasProfile:     false,
		Tier:           "guru",
		Role:           "owner",
	}
}

func createOwnerUser(db *gorm.DB) models.User {

	owner := getOwnerUser()

	if err := db.Where("role = ?", "owner").FirstOrCreate(&owner).Error; err != nil {
		panic(err)
	}

	return owner
}

func generateUser(random *rand.Rand, authRouter *gin.Engine, t *testing.T, tier string) models.UserResponse {

	if tier == "" {
		tier = "basic"
	}

	name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
	phone := utils.GenerateRandomPhoneNumber(random, 0)
	telegramUserId := fmt.Sprintf("%d", rand.Int64())

	payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/bot/signup", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	authRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var userResponse UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	assert.NoError(t, err)

	user := userResponse.Data

	if tier == "guru" || tier == "expert" {
		tx := initializers.DB.Model(&models.User{}).Where("id = ?", user.ID).Update("tier", tier)
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)
	}

	return user
}

func filterEthnosBySex(ethnos []models.Ethnos, targetSex string) []models.Ethnos {
	var filtered []models.Ethnos
	for _, e := range ethnos {
		if e.Sex == targetSex {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func generateCreateProfileRequest(
	random *rand.Rand,
	cities []models.City,
	ethnos []models.Ethnos,
	profileTags []models.ProfileTag,
	bodyArts []models.BodyArt,
	bodyTypes []models.BodyType,
	hairColors []models.HairColor,
	intimateHairCuts []models.IntimateHairCut) models.CreateProfileRequest {

	bio := "Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking."

	photosPayload := []models.CreatePhotoRequest{
		{URL: randomPhotos[random.IntN(len(randomPhotos))]},
	}

	optionsPayload := []models.CreateProfileOptionRequest{
		{ProfileTagID: profileTags[0].ID, Price: 5000, Comment: "This is my favourite!"},
		{ProfileTagID: profileTags[1].ID, Price: 50000, Comment: "I hate this!"},
	}

	bodyArtsPayload := []models.CreateBodyArtRequest{
		{ID: bodyArts[0].ID},
		{ID: bodyArts[1].ID},
	}

	ethnosSet := &ethnos[random.IntN(len(ethnos))]

	payload := models.CreateProfileRequest{
		Phone:               utils.GenerateRandomPhoneNumber(random, 10),
		Name:                utils.GenerateRandomStringWithPrefix(random, 15, "Alice"),
		Age:                 29,
		Height:              170,
		Weight:              57,
		CityID:              cities[random.IntN(len(cities))].ID,
		Bust:                2.5,
		Sex:                 ethnosSet.Sex,
		BodyTypeID:          &bodyTypes[random.IntN(len(bodyTypes))].ID,
		EthnosID:            &ethnosSet.ID,
		HairColorID:         &hairColors[random.IntN(len(hairColors))].ID,
		IntimateHairCutID:   &intimateHairCuts[random.IntN(len(intimateHairCuts))].ID,
		Bio:                 bio,
		PriceInHouseContact: ptr(10000),
		PriceInHouseHour:    ptr(20000),
		ContactPhone:        "77073778123",
		ContactTG:           "@lovely_mika",
		Photos:              photosPayload,
		Options:             optionsPayload,
		BodyArts:            bodyArtsPayload,
	}

	return payload
}

func assignRole(db *gorm.DB, t *testing.T, authRouter *gin.Engine, userRouter *gin.Engine, id string, role string) models.UserResponse {
	owner := createOwnerUser(db)

	accessTokenCookie, err := loginUserGetAccessToken(t, owner.Password, owner.TelegramUserId, authRouter)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	payload := &models.AssignRole{
		Id:   id,
		Role: role,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return models.UserResponse{}
	}

	url := "/api/users/role"
	assignRoleReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
	assignRoleReq.Header.Set("Content-Type", "application/json")
	assignRoleReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
	userRouter.ServeHTTP(w, assignRoleReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var userResponse UserResponse
	err = json.Unmarshal(w.Body.Bytes(), &userResponse)
	assert.Nil(t, err)
	assert.Equal(t, userResponse.Status, "success")
	assert.NotEmpty(t, userResponse)
	assert.Equal(t, role, userResponse.Data.Role)
	assert.Equal(t, "guru", userResponse.Data.Tier)

	return userResponse.Data
}

func loginUserGetAccessToken(t *testing.T, password string, telegramUserId int64, authRouter *gin.Engine) (*http.Cookie, error) {
	var jsonResponse map[string]interface{}

	w := httptest.NewRecorder()
	payloadLogin := fmt.Sprintf(`{"telegramUserId": "%d", "password": "%s"}`, telegramUserId, password)
	loginReq, _ := http.NewRequest("POST", "/api/auth/bot/login", bytes.NewBuffer([]byte(payloadLogin)))
	loginReq.Header.Set("Content-Type", "application/json")
	authRouter.ServeHTTP(w, loginReq)

	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

	assert.NoError(t, err)
	status := jsonResponse["status"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, status, "success")
	assert.NotEmpty(t, jsonResponse["access_token"])

	// Extract refresh_token from cookies
	cookies := w.Result().Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "access_token" {
			return cookie, err
		}
	}
	return nil, errors.New("cookie not found")
}

func populateProfileTags(db gorm.DB) []models.ProfileTag {
	var profileTags = []models.ProfileTag{
		{Name: "classic", AliasRu: "Классика", AliasEn: "Classic"},
		{Name: "blowjob", AliasRu: "Минет c/без резинки", AliasEn: "Blowjob with/without condom"},
		{Name: "deep-throat-condo", AliasRu: "Глубокий минет с резинкой", AliasEn: "Deep throat with condom"},
		{Name: "deep-throat-cum", AliasRu: "Глубокий минет c/без резинки c окончанием", AliasEn: "Deep throat with/without condom with finish"},
		{Name: "allow-cunnilingus", AliasRu: "Разрешу куннилингус", AliasEn: "Allow cunnilingus"},
		{Name: "blowjob-cum", AliasRu: "Минет c/без резинки c окончанием", AliasEn: "Blowjob with/without condom with finish"},
		{Name: "blowjob-with", AliasRu: "Минет с резинкой", AliasEn: "Blowjob with condom"},
		{Name: "massage-amateur", AliasRu: "Массаж любительский", AliasEn: "Amateur massage"},
		{Name: "massage-pro", AliasRu: "Массаж профессиональный", AliasEn: "Professional massage"},
		{Name: "vaginal-fisting", AliasRu: "Вагинальный фистинг", AliasEn: "Vaginal fisting"},
		{Name: "relaxing-massage", AliasRu: "Расслабляющий массаж", AliasEn: "Relaxing massage"},
		{Name: "kissing", AliasRu: "Поцелуи в губы", AliasEn: "Kissing"},
		{Name: "prostate-massage", AliasRu: "Массаж простаты", AliasEn: "Prostate massage"},
		{Name: "classic-massage", AliasRu: "Классический массаж", AliasEn: "Classic massage"},
		{Name: "evening-out", AliasRu: "Поеду отдыхать (в клуб, ресторан и.т.д.). Вечер:", AliasEn: "Evening out (club, restaurant, etc.). Price:"},
		{Name: "thai-body-massage", AliasRu: "Тайский боди массаж", AliasEn: "Thai body massage"},
		{Name: "deep-throat", AliasRu: "Глубокий минет c/без резинки", AliasEn: "Deep throat with/without condom"},
		{Name: "anilingus", AliasRu: "Анилингус, побалую язычком очко", AliasEn: "Anilingus"},
		{Name: "mistress", AliasRu: "Услуги Госпоже", AliasEn: "Mistress services"},
		{Name: "couples", AliasRu: "Услуги семейным парам", AliasEn: "Services for couples"},
		{Name: "french-kiss", AliasRu: "Французский поцелуй", AliasEn: "French kiss"},
		{Name: "erotic-massage", AliasRu: "Эротический массаж", AliasEn: "Erotic massage"},
		{Name: "phone-sex", AliasRu: "Секс по телефону", AliasEn: "Phone sex"},
		{Name: "stag-men", AliasRu: "Обслуживаю мальчишники. Вечер:", AliasEn: "Bachelor party service. Price:"},
		{Name: "group-sex", AliasRu: "Групповой секс", AliasEn: "Group sex"},
		{Name: "striptease-amateur", AliasRu: "Стриптиз любительский", AliasEn: "Amateur striptease"},
		{Name: "sakura", AliasRu: "Ветка сакуры", AliasEn: "Sakura branch"},
		{Name: "video", AliasRu: "Снимусь на видео", AliasEn: "Video shooting"},
		{Name: "anal", AliasRu: "Анальный секс", AliasEn: "Anal sex"},
		{Name: "role-play", AliasRu: "Ролевые игры, наряды", AliasEn: "Role play, costumes"},
		{Name: "photo", AliasRu: "Фото на память", AliasEn: "Photo memory"},
		{Name: "do-blowjob", AliasRu: "Сделаю минет", AliasEn: "Will do blowjob"},
		{Name: "striptease-pro", AliasRu: "Стриптиз профессиональный", AliasEn: "Professional striptease"},
		{Name: "deep-throat-nocondo-cum", AliasRu: "Глубокий минет без резинки c окончанием", AliasEn: "Deep throat without condom with finish"},
		{Name: "blowjob-nocondo-finish", AliasRu: "Минет без резинки c окончанием", AliasEn: "Blowjob without condom with finish"},
		{Name: "deep-throat-nocondo", AliasRu: "Глубокий минет без резинки", AliasEn: "Deep throat without condom"},
		{Name: "bj-raw", AliasRu: "Минет без резинки", AliasEn: "Blowjob without condom"},
		{Name: "girls", AliasRu: "Обслуживаю девушек", AliasEn: "Service for girls"},
		{Name: "guys", AliasRu: "Обслуживаю парней", AliasEn: "Service for guys"},
		{Name: "cuni", AliasRu: "Сделаю куннилингус", AliasEn: "Will do cunnilingus"},
		{Name: "stag-all", AliasRu: "Обслуживаю девишники/вечеринки. Вечер:", AliasEn: "Bachelorette/party service. Price:"},
		{Name: "party-service", AliasRu: "Обслуживаю вечеринки. Вечер:", AliasEn: "Party service. Price:"},
		{Name: "car-blowjob", AliasRu: "Сделаю минет за рулем", AliasEn: "Car blowjob"},
	}

	var tags []models.ProfileTag
	for _, profileTag := range profileTags {
		tag := models.ProfileTag{
			Name:    profileTag.Name,
			AliasEn: profileTag.AliasEn,
			AliasRu: profileTag.AliasRu,
		}
		tags = append(tags, tag)
	}

	tx := db.Begin()
	// Batch insert options
	if err := tx.Create(&tags).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			tx.Rollback() // Rollback after detecting duplicate key error

			// Now run the fetch query outside the aborted transaction
			var existingTags []models.ProfileTag
			if err := db.Find(&existingTags).Error; err != nil {
				log.Fatalf("failed to fetch existing profile tags: %v", err)
				return nil
			}
			return existingTags
		}

		// Handle other errors
		tx.Rollback()
		panic("failed to create profile tags: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit profile_tags transaction: %v", err)
		return nil
	}

	return tags
}

func populateUserTags(db gorm.DB) []models.UserTag {
	var userTag = []models.UserTag{
		{Name: "hygiene", AliasRu: "Гигиена", AliasEn: "Hygiene"},
		{Name: "neat", AliasRu: "Опрятность", AliasEn: "Neatness"},
		{Name: "generous", AliasRu: "Щедрость", AliasEn: "Generosity"},
		{Name: "punctual", AliasRu: "Пунктуальность", AliasEn: "Punctuality"},
		{Name: "boundaries", AliasRu: "Соблюдение границ", AliasEn: "Respect boundaries"},
		{Name: "sociable", AliasRu: "Общительность", AliasEn: "Sociability"},
	}

	var tags []models.UserTag
	for _, userTag := range userTag {
		tag := models.UserTag{
			Name:    userTag.Name,
			AliasEn: userTag.AliasEn,
			AliasRu: userTag.AliasRu,
		}
		tags = append(tags, tag)
	}

	tx := db.Begin()
	// Batch insert options
	if err := tx.Create(&tags).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			tx.Rollback() // Rollback after detecting duplicate key error

			// Now run the fetch query outside the aborted transaction
			var existingTags []models.UserTag
			if err := db.Find(&existingTags).Error; err != nil {
				log.Fatalf("failed to fetch existing profile tags: %v", err)
				return nil
			}
			return existingTags
		}

		// Handle other errors
		tx.Rollback()
		panic("failed to create profile tags: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit profile_tags transaction: %v", err)
		return nil
	}

	return tags
}

func populateCities(db gorm.DB) []models.City {
	var cities = []models.City{
		{Name: "almaty", AliasRu: "Алматы", AliasEn: "Almaty"},
		{Name: "ust-kamenogorsk", AliasRu: "Усть-Каменогорск", AliasEn: "Ust-Kamenogorsk"},
		{Name: "zhezkazgan", AliasRu: "Жезказган", AliasEn: "Zhezkazgan"},
		{Name: "zhetysai", AliasRu: "Жетысай", AliasEn: "Zhetysai"},
		{Name: "lisakovsk", AliasRu: "Лисаковск", AliasEn: "Lisakovsk"},
		{Name: "astana", AliasRu: "Астана", AliasEn: "Astana"},
		{Name: "kostanay", AliasRu: "Костанай", AliasEn: "Kostanay"},
		{Name: "kapchagay", AliasRu: "Капчагай", AliasEn: "Kapchagay"},
		{Name: "ridder", AliasRu: "Риддер", AliasEn: "Ridder"},
		{Name: "shu", AliasRu: "Шу", AliasEn: "Shu"},
		{Name: "shymkent", AliasRu: "Шымкент", AliasEn: "Shymkent"},
		{Name: "kyzylorda", AliasRu: "Кызылорда", AliasEn: "Kyzylorda"},
		{Name: "balhash", AliasRu: "Балхаш", AliasEn: "Balkhash"},
		{Name: "kaskelen", AliasRu: "Каскелен", AliasEn: "Kaskelen"},
		{Name: "shahtinsk", AliasRu: "Шахтинск", AliasEn: "Shahtinsk"},
		{Name: "karaganda", AliasRu: "Караганда", AliasEn: "Karaganda"},
		{Name: "kokshetau", AliasRu: "Кокшетау", AliasEn: "Kokshetau"},
		{Name: "aksay", AliasRu: "Аксай", AliasEn: "Aksay"},
		{Name: "kulsary", AliasRu: "Кульсары", AliasEn: "Kulsary"},
		{Name: "yesik", AliasRu: "Есик", AliasEn: "Yesik"},
		{Name: "aktau", AliasRu: "Актау", AliasEn: "Aktau"},
		{Name: "taldykorgan", AliasRu: "Талдыкорган", AliasEn: "Taldykorgan"},
		{Name: "shchuchinsk", AliasRu: "Щучинск", AliasEn: "Shchuchinsk"},
		{Name: "stepnogorsk", AliasRu: "Степногорск", AliasEn: "Stepnogorsk"},
		{Name: "zharkent", AliasRu: "Жаркент", AliasEn: "Zharkent"},
		{Name: "aktobe", AliasRu: "Актобе", AliasEn: "Aktobe"},
		{Name: "turkestan", AliasRu: "Туркестан", AliasEn: "Turkestan"},
		{Name: "rudny", AliasRu: "Рудный", AliasEn: "Rudny"},
		{Name: "talgar", AliasRu: "Талгар", AliasEn: "Talgar"},
		{Name: "shardara", AliasRu: "Шардара", AliasEn: "Shardara"},
		{Name: "atyrau", AliasRu: "Атырау", AliasEn: "Atyrau"},
		{Name: "semey", AliasRu: "Семей", AliasEn: "Semey"},
		{Name: "zhanaozen", AliasRu: "Жанаозен", AliasEn: "Zhanaozen"},
		{Name: "saran", AliasRu: "Сарань", AliasEn: "Saran"},
		{Name: "atbasar", AliasRu: "Атбасар", AliasEn: "Atbasar"},
		{Name: "taraz", AliasRu: "Тараз", AliasEn: "Taraz"},
		{Name: "petropavl", AliasRu: "Петропавловск", AliasEn: "Petropavl"},
		{Name: "satpayev", AliasRu: "Сатпаев", AliasEn: "Satpayev"},
		{Name: "aksu", AliasRu: "Аксу", AliasEn: "Aksu"},
		{Name: "tekeli", AliasRu: "Текели", AliasEn: "Tekeli"},
		{Name: "uralsk", AliasRu: "Уральск", AliasEn: "Uralsk"},
		{Name: "temirtau", AliasRu: "Темиртау", AliasEn: "Temirtau"},
		{Name: "kentau", AliasRu: "Кентау", AliasEn: "Kentau"},
		{Name: "zyryanovsk", AliasRu: "Зыряновск", AliasEn: "Zyryanovsk"},
		{Name: "mangistau", AliasRu: "Мангистау", AliasEn: "Mangistau"},
		{Name: "pavlodar", AliasRu: "Павлодар", AliasEn: "Pavlodar"},
		{Name: "ekibastuz", AliasRu: "Экибастуз", AliasEn: "Ekibastuz"},
		{Name: "saryagash", AliasRu: "Сарыагаш", AliasEn: "Saryagash"},
		{Name: "baykonur", AliasRu: "Байконыр", AliasEn: "Baykonur"},
		{Name: "jitiqara", AliasRu: "Житикара", AliasEn: "Jitiqara"},
		{Name: "aral", AliasRu: "Аральск", AliasEn: "Aral"},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&cities).Error; err != nil {
		// Check if it's a duplicate key error
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			// Rollback the failed transaction
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			// Fetch and return existing records
			var existingCities []models.City
			if err := db.Find(&existingCities).Error; err != nil {
				log.Fatalf("failed to fetch existing cities: %v", err)
				return nil
			}
			return existingCities
		}

		// Handle other errors
		tx.Rollback()
		log.Fatalf("failed to create cities: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit cities transaction: %v", err)
		return nil
	}

	return cities
}

func populateEthnos(db gorm.DB) []models.Ethnos {
	var femaleEthnos = []models.Ethnos{
		{Name: "metiska", AliasRu: "Метиска", AliasEn: "Métis", Sex: "female"},
		{Name: "chuvashka", AliasRu: "Чувашка", AliasEn: "Chuvash", Sex: "female"},
		{Name: "kirgizka", AliasRu: "Киргизка", AliasEn: "Kyrgyz", Sex: "female"},
		{Name: "azerbaijanka", AliasRu: "Азербайджанка", AliasEn: "Azerbaijani", Sex: "female"},
		{Name: "iranka", AliasRu: "Иранка", AliasEn: "Iranian", Sex: "female"},
		{Name: "taika", AliasRu: "Тайка", AliasEn: "Thai", Sex: "female"},
		{Name: "ukrainka", AliasRu: "Украинка", AliasEn: "Ukrainian", Sex: "female"},
		{Name: "litovka", AliasRu: "Литовка", AliasEn: "Lithuanian", Sex: "female"},
		{Name: "ingushka", AliasRu: "Ингушка", AliasEn: "Ingush", Sex: "female"},
		{Name: "dagestanka", AliasRu: "Дагестанка", AliasEn: "Dagestani", Sex: "female"},
		{Name: "dunganka", AliasRu: "Дунганка", AliasEn: "Dungan", Sex: "female"},
		{Name: "osetinka", AliasRu: "Осетинка", AliasEn: "Ossetian", Sex: "female"},
		{Name: "turkmenka", AliasRu: "Туркменка", AliasEn: "Turkmen", Sex: "female"},
		{Name: "mulatka", AliasRu: "Мулатка", AliasEn: "Mulatto", Sex: "female"},
		{Name: "evropeyka", AliasRu: "Европейка", AliasEn: "European", Sex: "female"},
		{Name: "koreyanka", AliasRu: "Кореянка", AliasEn: "Korean", Sex: "female"},
		{Name: "beloruska", AliasRu: "Белоруска", AliasEn: "Belarusian", Sex: "female"},
		{Name: "chechenka", AliasRu: "Чеченка", AliasEn: "Chechen", Sex: "female"},
		{Name: "tadzhichka", AliasRu: "Таджичка", AliasEn: "Tajik", Sex: "female"},
		{Name: "kavkazka", AliasRu: "Кавказка", AliasEn: "Caucasian", Sex: "female"},
		{Name: "slavyanka", AliasRu: "Славянка", AliasEn: "Slavic", Sex: "female"},
		{Name: "turchanka", AliasRu: "Турчанка", AliasEn: "Turkish", Sex: "female"},
		{Name: "evreyka", AliasRu: "Еврейка", AliasEn: "Jewish", Sex: "female"},
		{Name: "nemka", AliasRu: "Немка", AliasEn: "German", Sex: "female"},
		{Name: "kazashka", AliasRu: "Казашка", AliasEn: "Kazakh", Sex: "female"},
		{Name: "frantsuzhenka", AliasRu: "Француженка", AliasEn: "French", Sex: "female"},
		{Name: "latyshka", AliasRu: "Латышка", AliasEn: "Latvian", Sex: "female"},
		{Name: "gruzinka", AliasRu: "Грузинка", AliasEn: "Georgian", Sex: "female"},
		{Name: "moldavanka", AliasRu: "Молдаванка", AliasEn: "Moldovan", Sex: "female"},
		{Name: "bolgarka", AliasRu: "Болгарка", AliasEn: "Bulgarian", Sex: "female"},
		{Name: "bashkirka", AliasRu: "Башкирка", AliasEn: "Bashkir", Sex: "female"},
		{Name: "rumynka", AliasRu: "Румынка", AliasEn: "Romanian", Sex: "female"},
		{Name: "grechanka", AliasRu: "Гречанка", AliasEn: "Greek", Sex: "female"},
		{Name: "uzbechka", AliasRu: "Узбечка", AliasEn: "Uzbek", Sex: "female"},
		{Name: "ispanka", AliasRu: "Испанка", AliasEn: "Spanish", Sex: "female"},
		{Name: "tatarka", AliasRu: "Татарка", AliasEn: "Tatar", Sex: "female"},
		{Name: "yakutka", AliasRu: "Якутка", AliasEn: "Yakut", Sex: "female"},
		{Name: "aziatka", AliasRu: "Азиатка", AliasEn: "Asian", Sex: "female"},
		{Name: "mordvinka", AliasRu: "Мордвинка", AliasEn: "Mordvin", Sex: "female"},
		{Name: "kitayanka", AliasRu: "Китаянка", AliasEn: "Chinese", Sex: "female"},
		{Name: "tsyganka", AliasRu: "Цыганка", AliasEn: "Gypsy", Sex: "female"},
		{Name: "armyanka", AliasRu: "Армянка", AliasEn: "Armenian", Sex: "female"},
		{Name: "italyanka", AliasRu: "Итальянка", AliasEn: "Italian", Sex: "female"},
		{Name: "uygurka", AliasRu: "Уйгурка", AliasEn: "Uyghur", Sex: "female"},
		{Name: "polyachka", AliasRu: "Полячка", AliasEn: "Polish", Sex: "female"},
		{Name: "arabka", AliasRu: "Арабка", AliasEn: "Arab", Sex: "female"},
	}
	var maleEthnos = []models.Ethnos{
		{Name: "dagestanets", AliasRu: "Дагестанец", AliasEn: "Dagestani", Sex: "male"},
		{Name: "slavyanin", AliasRu: "Славянин", AliasEn: "Slavic", Sex: "male"},
		{Name: "bolgarin", AliasRu: "Болгарин", AliasEn: "Bulgarian", Sex: "male"},
		{Name: "kavkazets", AliasRu: "Кавказец", AliasEn: "Caucasian", Sex: "male"},
		{Name: "ingush", AliasRu: "Ингуш", AliasEn: "Ingush", Sex: "male"},
		{Name: "osetinets", AliasRu: "Осетинец", AliasEn: "Ossetian", Sex: "male"},
		{Name: "armyanin", AliasRu: "Армянин", AliasEn: "Armenian", Sex: "male"},
		{Name: "kazakh", AliasRu: "Казах", AliasEn: "Kazakh", Sex: "male"},
		{Name: "ukrainets", AliasRu: "Украинец", AliasEn: "Ukrainian", Sex: "male"},
		{Name: "tsyganin", AliasRu: "Цыганин", AliasEn: "Gypsy", Sex: "male"},
		{Name: "gruzin", AliasRu: "Грузин", AliasEn: "Georgian", Sex: "male"},
		{Name: "italyanets", AliasRu: "Итальянец", AliasEn: "Italian", Sex: "male"},
		{Name: "evropeets", AliasRu: "Европеец", AliasEn: "European", Sex: "male"},
		{Name: "litovets", AliasRu: "Литовец", AliasEn: "Lithuanian", Sex: "male"},
		{Name: "tadzhik", AliasRu: "Таджик", AliasEn: "Tajik", Sex: "male"},
		{Name: "frantsuz", AliasRu: "Француз", AliasEn: "French", Sex: "male"},
		{Name: "rumyn", AliasRu: "Румын", AliasEn: "Romanian", Sex: "male"},
		{Name: "ispanets", AliasRu: "Испанец", AliasEn: "Spanish", Sex: "male"},
		{Name: "polyak", AliasRu: "Поляк", AliasEn: "Polish", Sex: "male"},
		{Name: "chuvash", AliasRu: "Чуваш", AliasEn: "Chuvash", Sex: "male"},
		{Name: "turkmen", AliasRu: "Туркмен", AliasEn: "Turkmen", Sex: "male"},
		{Name: "moldavanin", AliasRu: "Молдаванин", AliasEn: "Moldovan", Sex: "male"},
		{Name: "kurd", AliasRu: "Курд", AliasEn: "Kurd", Sex: "male"},
		{Name: "evrey", AliasRu: "Еврей", AliasEn: "Jewish", Sex: "male"},
		{Name: "chechenets", AliasRu: "Чеченец", AliasEn: "Chechen", Sex: "male"},
		{Name: "bashkir", AliasRu: "Башкир", AliasEn: "Bashkir", Sex: "male"},
		{Name: "metis", AliasRu: "Метис", AliasEn: "Métis", Sex: "male"},
		{Name: "nemets", AliasRu: "Немец", AliasEn: "German", Sex: "male"},
		{Name: "mulat", AliasRu: "Мулат", AliasEn: "Mulatto", Sex: "male"},
		{Name: "arab", AliasRu: "Араб", AliasEn: "Arab", Sex: "male"},
		{Name: "latysh", AliasRu: "Латыш", AliasEn: "Latvian", Sex: "male"},
		{Name: "russkiy", AliasRu: "Русский", AliasEn: "Russian", Sex: "male"},
		{Name: "belorus", AliasRu: "Белорус", AliasEn: "Belarusian", Sex: "male"},
		{Name: "dungan", AliasRu: "Дунган", AliasEn: "Dungan", Sex: "male"},
		{Name: "grek", AliasRu: "Грек", AliasEn: "Greek", Sex: "male"},
		{Name: "yakut", AliasRu: "Якут", AliasEn: "Yakut", Sex: "male"},
		{Name: "koreets", AliasRu: "Кореец", AliasEn: "Korean", Sex: "male"},
		{Name: "uygur", AliasRu: "Уйгур", AliasEn: "Uyghur", Sex: "male"},
		{Name: "tatarin", AliasRu: "Татарин", AliasEn: "Tatar", Sex: "male"},
		{Name: "turok", AliasRu: "Турок", AliasEn: "Turkish", Sex: "male"},
		{Name: "kitayets", AliasRu: "Китаец", AliasEn: "Chinese", Sex: "male"},
		{Name: "mordvin", AliasRu: "Мордвин", AliasEn: "Mordvin", Sex: "male"},
		{Name: "iranets", AliasRu: "Иранец", AliasEn: "Iranian", Sex: "male"},
		{Name: "azerbaidzhanets", AliasRu: "Азербайджанец", AliasEn: "Azerbaijani", Sex: "male"},
		{Name: "uzbek", AliasRu: "Узбек", AliasEn: "Uzbek", Sex: "male"},
		{Name: "aziat", AliasRu: "Азиат", AliasEn: "Asian", Sex: "male"},
		{Name: "kirgiz", AliasRu: "Киргиз", AliasEn: "Kyrgyz", Sex: "male"},
	}

	ethnos := append(femaleEthnos, maleEthnos...)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&ethnos).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			var existingEthnos []models.Ethnos
			if err := db.Find(&existingEthnos).Error; err != nil {
				log.Fatalf("failed to fetch existing ethnos: %v", err)
				return nil
			}
			return existingEthnos
		}

		tx.Rollback()
		log.Fatalf("failed to create ethnos: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit ethnos transaction: %v", err)
		return nil
	}

	return ethnos
}

func populateBodyTypes(db gorm.DB) []models.BodyType {
	var bodyTypes = []models.BodyType{
		{Name: "hudaya", AliasRu: "Худая", AliasEn: "Slim"},
		{Name: "stroynaya", AliasRu: "Стройная", AliasEn: "Fit"},
		{Name: "sportivnaya", AliasRu: "Спортивная", AliasEn: "Athletic"},
		{Name: "polnaya", AliasRu: "Полная", AliasEn: "Full-figured"},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&bodyTypes).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			var existingBodyTypes []models.BodyType
			if err := db.Find(&existingBodyTypes).Error; err != nil {
				log.Fatalf("failed to fetch existing body types: %v", err)
				return nil
			}
			return existingBodyTypes
		}

		tx.Rollback()
		log.Fatalf("failed to create body types: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit body types transaction: %v", err)
		return nil
	}

	return bodyTypes
}

func populateBodyArts(db gorm.DB) []models.BodyArt {
	var bodyArts = []models.BodyArt{
		{Name: "tatu", AliasRu: "Татуировки", AliasEn: "Tattoos"},
		{Name: "silikon_v_grudi", AliasRu: "Силикон в груди", AliasEn: "Breast Implants"},
		{Name: "pirsing", AliasRu: "Пирсинг", AliasEn: "Piercing"},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&bodyArts).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			var existingBodyArts []models.BodyArt
			if err := db.Find(&existingBodyArts).Error; err != nil {
				log.Fatalf("failed to fetch existing body arts: %v", err)
				return nil
			}
			return existingBodyArts
		}

		tx.Rollback()
		log.Fatalf("failed to create body arts: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit body arts transaction: %v", err)
		return nil
	}

	return bodyArts
}

func populateIntimateHairCuts(db gorm.DB) []models.IntimateHairCut {
	var intimateHairCuts = []models.IntimateHairCut{
		{Name: "polnaya_depilyatsiya", AliasRu: "Полная депиляция", AliasEn: "Full depilation"},
		{Name: "akkuratnaya_strizhka", AliasRu: "Аккуратная стрижка", AliasEn: "Neat trim"},
		{Name: "naturalnaya", AliasRu: "Натуральная", AliasEn: "Natural"},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&intimateHairCuts).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			var existingIntimateHairCuts []models.IntimateHairCut
			if err := db.Find(&existingIntimateHairCuts).Error; err != nil {
				log.Fatalf("failed to fetch existing intimate haircuts: %v", err)
				return nil
			}
			return existingIntimateHairCuts
		}

		tx.Rollback()
		log.Fatalf("failed to create intimate haircuts: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit intimate haircuts transaction: %v", err)
		return nil
	}

	return intimateHairCuts
}

func populateHairColors(db gorm.DB) []models.HairColor {
	var hairColors = []models.HairColor{
		{Name: "brunetka", AliasRu: "Брюнетка", AliasEn: "Brunette"},
		{Name: "shatenka", AliasRu: "Шатенка", AliasEn: "Brown-haired"},
		{Name: "ryzhaya", AliasRu: "Рыжая", AliasEn: "Red-haired"},
		{Name: "rusaya", AliasRu: "Русая", AliasEn: "Light brown"},
		{Name: "blondinka", AliasRu: "Блондинка", AliasEn: "Blonde"},
		{Name: "lysaya", AliasRu: "Лысая", AliasEn: "Bald"},
		{Name: "tsvetnaya", AliasRu: "Цветная", AliasEn: "Colored"},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&hairColors).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			log.Printf("Duplicate key error: %v", err)
			if rbErr := tx.Rollback().Error; rbErr != nil {
				log.Fatalf("failed to rollback transaction: %v", rbErr)
			}

			var existingHairColors []models.HairColor
			if err := db.Find(&existingHairColors).Error; err != nil {
				log.Fatalf("failed to fetch existing hair colors: %v", err)
				return nil
			}
			return existingHairColors
		}

		tx.Rollback()
		log.Fatalf("failed to create hair colors: %v", err)
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit hair colors transaction: %v", err)
		return nil
	}

	return hairColors
}

func checkProfilesMatch(t *testing.T, userID string, payload models.CreateProfileRequest,
	profileResponse CreateProfileResponse, expectedActive bool, expectedVerified bool, expectedModerated bool) {

	assert.Equal(t, profileResponse.Data.UserID.String(), userID)

	assert.Equal(t, profileResponse.Data.Active, expectedActive)
	assert.Equal(t, profileResponse.Data.Verified, expectedVerified)
	assert.Equal(t, profileResponse.Data.Moderated, expectedModerated)
	assert.Equal(t, profileResponse.Data.Name, payload.Name)
	assert.Equal(t, profileResponse.Data.Phone, payload.Phone)
	assert.Equal(t, profileResponse.Data.ContactPhone, payload.ContactPhone)
	assert.Equal(t, profileResponse.Data.ContactTG, payload.ContactTG)
	assert.Equal(t, profileResponse.Data.ContactWA, payload.ContactWA)
	assert.NotNil(t, profileResponse.Data.CreatedAt)
	assert.NotNil(t, profileResponse.Data.UpdatedAt)
	assert.Equal(t, profileResponse.Data.UpdatedBy.String(), userID)
	assert.Equal(t, len(profileResponse.Data.ProfileOptions), len(payload.Options))

	for i := 0; i < len(payload.Options); i++ {
		assert.Equal(t, profileResponse.Data.ProfileOptions[i].Comment, payload.Options[i].Comment)
		assert.NotNil(t, profileResponse.Data.ProfileOptions[i].ProfileTag.ID)
		assert.NotNil(t, profileResponse.Data.ProfileOptions[i].ProfileTag.Name)
	}

	assert.Equal(t, len(profileResponse.Data.BodyArts), len(payload.BodyArts))

	for i := 0; i < len(payload.BodyArts); i++ {
		assert.NotNil(t, profileResponse.Data.BodyArts[i].ProfileID)
		assert.Equal(t, profileResponse.Data.BodyArts[i].BodyArtID, payload.BodyArts[i].ID)
	}

	assert.Equal(t, len(profileResponse.Data.Photos), len(payload.Photos))

	for i := 0; i < len(payload.Photos); i++ {
		assert.NotNil(t, profileResponse.Data.Photos[i].ID)
		assert.NotNil(t, profileResponse.Data.Photos[i].ProfileID)
		assert.NotNil(t, profileResponse.Data.Photos[i].CreatedAt)

		assert.False(t, profileResponse.Data.Photos[i].Approved)
		assert.False(t, profileResponse.Data.Photos[i].Deleted)
		assert.False(t, profileResponse.Data.Photos[i].Disabled)

		assert.Equal(t, profileResponse.Data.Photos[i].URL, payload.Photos[i].URL)
	}

	assert.Equal(t, profileResponse.Data.BodyTypeID, payload.BodyTypeID)
	assert.Equal(t, profileResponse.Data.EthnosID, payload.EthnosID)
	assert.Equal(t, profileResponse.Data.HairColorID, payload.HairColorID)
	assert.Equal(t, profileResponse.Data.IntimateHairCutID, payload.IntimateHairCutID)
	assert.Equal(t, profileResponse.Data.CityID, payload.CityID)
	assert.Equal(t, profileResponse.Data.Age, payload.Age)
	assert.Equal(t, profileResponse.Data.Height, payload.Height)
	assert.Equal(t, profileResponse.Data.Weight, payload.Weight)
	assert.Equal(t, profileResponse.Data.Bust, payload.Bust)
	assert.Equal(t, profileResponse.Data.Bio, payload.Bio)
	assert.Equal(t, profileResponse.Data.AddressLatitude, payload.AddressLatitude)
	assert.Equal(t, profileResponse.Data.AddressLongitude, payload.AddressLongitude)

	assert.Equal(t, profileResponse.Data.PriceInHouseNightRatio, 1.0)
	assert.Equal(t, profileResponse.Data.PriceInHouseContact, payload.PriceInHouseContact)
	assert.Equal(t, profileResponse.Data.PriceInHouseHour, payload.PriceInHouseHour)

	assert.Equal(t, profileResponse.Data.PriceVisitNightRatio, 1.0)
	assert.Equal(t, profileResponse.Data.PriceVisitContact, payload.PriceVisitContact)
	assert.Equal(t, profileResponse.Data.PriceVisitHour, payload.PriceVisitHour)

	assert.Equal(t, profileResponse.Data.PriceCarNightRatio, 1.0)
	assert.Equal(t, profileResponse.Data.PriceCarContact, payload.PriceCarContact)
	assert.Equal(t, profileResponse.Data.PriceCarHour, payload.PriceCarHour)

	assert.Equal(t, profileResponse.Data.PriceSaunaNightRatio, 1.0)
	assert.Equal(t, profileResponse.Data.PriceSaunaContact, payload.PriceSaunaContact)
	assert.Equal(t, profileResponse.Data.PriceSaunaHour, payload.PriceSaunaHour)

	assert.Equal(t, profileResponse.Data.DeletedAt.Valid, false)
	assert.Empty(t, profileResponse.Data.Services)
}
