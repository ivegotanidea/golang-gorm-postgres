package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"gorm.io/gorm"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func SetupPCRouter(profileController *controllers.ProfileController) *gin.Engine {
	r := gin.Default()

	userRouteController := NewRouteProfileController(*profileController)

	api := r.Group("/api")
	userRouteController.ProfileRoute(api)

	return r
}

func SetupPCController() controllers.ProfileController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	profileController := controllers.NewProfileController(initializers.DB)
	profileController.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	if err := profileController.DB.AutoMigrate(
		&models.HairColor{},
		&models.IntimateHairCut{},
		&models.Ethnos{},
		&models.BodyType{},
		&models.ProfileBodyArt{},
		&models.BodyArt{},
		&models.City{},
		&models.User{},
		&models.Profile{},
		&models.Service{},
		&models.Photo{},
		&models.ProfileOption{},
		&models.UserRating{},
		&models.ProfileRating{},
		&models.ProfileTag{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return profileController
}

func PopulateProfileTags(db gorm.DB) []models.ProfileTag {
	var profileTags = []string{
		"Классика",
		"Минет c/без резинки",
		"Глубокий минет с резинкой",
		"Глубокий минет c/без резинки c окончанием",
		"Разрешу куннилингус",
		"Минет c/без резинки c окончанием",
		"Минет с резинкой",
		"Массаж любительский",
		"Массаж профессиональный",
		"Вагинальный фистинг",
		"Расслабляющий массаж",
		"Поцелуи в губы",
		"Массаж простаты",
		"Классический массаж",
		"Поеду отдыхать (в клуб, ресторан и.т.д.). Вечер:",
		"Тайский боди массаж",
		"Глубокий минет c/без резинки",
		"Анилингус, побалую язычком очко",
		"Услуги Госпоже",
		"Услуги семейным парам",
		"Французский поцелуй",
		"Эротический массаж",
		"Секс по телефону",
		"Обслуживаю мальчишники. Вечер:",
		"Групповой секс",
		"Стриптиз любительский",
		"Ветка сакуры",
		"Снимусь на видео",
		"Анальный секс",
		"Ролевые игры, наряды",
		"Фото на память",
		"Сделаю минет",
		"Стриптиз профессиональный",
		"Глубокий минет без резинки c окончанием",
		"Минет без резинки c окончанием",
		"Глубокий минет без резинки",
		"Минет без резинки",
		"Обслуживаю девушек",
		"Обслуживаю парней",
		"Сделаю куннилингус",
		"Обслуживаю девишники/вечеринки. Вечер:",
		"Обслуживаю вечеринки. Вечер:",
	}

	var tags []models.ProfileTag
	for _, profileTag := range profileTags {
		tag := models.ProfileTag{
			Name: profileTag,
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

func PopulateCities(db gorm.DB) []models.City {
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

func PopulateEthnos(db gorm.DB) []models.Ethnos {
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

func PopulateBodyTypes(db gorm.DB) []models.BodyType {
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

func PopulateBodyArts(db gorm.DB) []models.BodyArt {
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

func PopulateIntimateHairCuts(db gorm.DB) []models.IntimateHairCut {
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

func PopulateHairColors(db gorm.DB) []models.HairColor {
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

func TestProfileRoutes(t *testing.T) {

	ac := SetupAuthController()

	//uc := SetupUCController()
	pc := SetupPCController()

	authRouter := SetupACRouter(&ac)
	//userRouter := SetupUCRouter(&uc)
	profileRouter := SetupPCRouter(&pc)

	profileTags := PopulateProfileTags(*pc.DB)

	cities := PopulateCities(*pc.DB)

	// filters

	bodyTypes := PopulateBodyTypes(*pc.DB)

	ethnos := PopulateEthnos(*pc.DB)

	hairColors := PopulateHairColors(*pc.DB)

	intimateHairCuts := PopulateIntimateHairCuts(*pc.DB)

	bodyArts := PopulateBodyArts(*pc.DB)

	//createOwnerUser(profileController.DB)

	random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

	t.Cleanup(func() {
		utils.CleanupTestUsers(pc.DB)
		utils.DropAllTables(pc.DB)
	})

	t.Run("POST /api/profiles/: fail without access token ", func(t *testing.T) {

		w := httptest.NewRecorder()

		payload := &models.CreateProfileRequest{
			Name: "Alice",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.Header.Set("Content-Type", "application/json")
		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("POST /api/profiles/: fail with access_token but bad json (only name) ", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := &models.CreateProfileRequest{
			Name: "Alice",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("POST /api/profiles/: success with access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		assert.Equal(t, profileResponse.Status, "success")
		assert.NotNil(t, profileResponse.Data.ID)
		assert.Equal(t, profileResponse.Data.UserID, user.ID)

		assert.Equal(t, profileResponse.Data.Active, true)
		assert.Equal(t, profileResponse.Data.Verified, false)
		assert.Equal(t, profileResponse.Data.Moderated, false)
		assert.Equal(t, profileResponse.Data.Name, payload.Name)
		assert.Equal(t, profileResponse.Data.Phone, payload.Phone)
		assert.Equal(t, profileResponse.Data.ContactPhone, payload.ContactPhone)
		assert.Equal(t, profileResponse.Data.ContactTG, payload.ContactTG)
		assert.Equal(t, profileResponse.Data.ContactWA, payload.ContactWA)
		assert.NotNil(t, profileResponse.Data.CreatedAt)
		assert.NotNil(t, profileResponse.Data.UpdatedAt)
		assert.Equal(t, profileResponse.Data.UpdatedBy, user.ID)
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

		assert.Equal(t, profileResponse.Data.PrinceSaunaNightRatio, 1.0)
		assert.Equal(t, profileResponse.Data.PriceSaunaContact, payload.PriceSaunaContact)
		assert.Equal(t, profileResponse.Data.PriceSaunaHour, payload.PriceSaunaHour)

		assert.Equal(t, profileResponse.Data.DeletedAt.Valid, false)
		assert.Nil(t, profileResponse.Data.Services)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	//t.Run("PUT /api/profiles/: success profile update", func(t *testing.T) {
	//	user := generateUser(random, authRouter, t)
	//
	//	accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
	//	w := httptest.NewRecorder()
	//
	//	payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)
	//
	//	jsonPayload, err := json.Marshal(payload)
	//	if err != nil {
	//		fmt.Println("Error marshaling payload:", err)
	//		return
	//	}
	//
	//	createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
	//	createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
	//	createProfileReq.Header.Set("Content-Type", "application/json")
	//
	//	profileRouter.ServeHTTP(w, createProfileReq)
	//
	//	var profileResponse CreateProfileResponse
	//	err = json.Unmarshal(w.Body.Bytes(), &profileResponse)
	//
	//	updatePayload := &models.UpdateOwnProfileRequest{}
	//
	//})
}
