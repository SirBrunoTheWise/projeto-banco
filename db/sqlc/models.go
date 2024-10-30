package db

import (
	"database/sql"
	"time"
)

type Card struct {
	CardID          int64       `json:"card_id"`
	CardType        int16       `json:"card_type"`
	CardNumber      int64       `json:"card_number"`
	CardProgression interface{} `json:"card_progression"`
	CardImage       []byte      `json:"card_image"`
}

type CreditCard struct {
	CreditCardName            string        `json:"credit_card_name"`
	CreditCardNumber          string        `json:"credit_card_number"`
	CreditCardCvv             interface{}   `json:"credit_card_cvv"`
	CreditCardExpirationMonth sql.NullInt16 `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  interface{}   `json:"credit_card_expiration_year"`
	UserUserID                int64         `json:"user_user_id"`
}

type Diary struct {
	DateOf   time.Time `json:"date_of"`
	UserID   int64     `json:"user_id"`
	Exercise int64     `json:"exercise"`
	Meal     int64     `json:"meal"`
	Cards    int64     `json:"cards"`
}

type Exercise struct {
	ExerciseID       int64          `json:"exercise_id"`
	TimeStart        time.Time      `json:"time_start"`
	TimeEnd          time.Time      `json:"time_end"`
	CaloriesSpent    int32          `json:"calories_spent"`
	AverageHeartRate int32          `json:"average_heart_rate"`
	ExerciseType     sql.NullString `json:"exercise_type"`
}

type ExerciseBiking struct {
	Coordinates            []interface{}  `json:"coordinates"`
	PaceSeconds            int32          `json:"pace_seconds"`
	DistanceTraveledMeters int32          `json:"distance_traveled_meters"`
	WheelWearLevel         sql.NullInt16  `json:"wheel_wear_level"`
	TerrainType            sql.NullString `json:"terrain_type"`
	ExerciseID             int64          `json:"exercise_id"`
}

type ExerciseHiking struct {
	Coordinates            []interface{}  `json:"coordinates"`
	PaceSeconds            int32          `json:"pace_seconds"`
	DistanceTraveledMeters int32          `json:"distance_traveled_meters"`
	WildernessType         sql.NullString `json:"wilderness_type"`
	ExerciseID             int64          `json:"exercise_id"`
}

type ExerciseJogging struct {
	Coordinates            []interface{} `json:"coordinates"`
	PaceSeconds            int32         `json:"pace_seconds"`
	DistanceTraveledMeters int32         `json:"distance_traveled_meters"`
	Cadence                sql.NullInt32 `json:"cadence"`
	TerrainLevel           sql.NullInt32 `json:"terrain_level"`
	ExerciseExerciseID     int64         `json:"exercise_exercise_id"`
}

type ExerciseWeightlifting struct {
	SheetType              sql.NullString `json:"sheet_type"`
	ExerciseName           sql.NullString `json:"exercise_name"`
	Sets                   sql.NullInt32  `json:"sets"`
	Reps                   sql.NullInt32  `json:"reps"`
	TimeBetweenSetsSeconds sql.NullInt32  `json:"time_between_sets_seconds"`
	TimeBetweenRepsSeconds sql.NullInt32  `json:"time_between_reps_seconds"`
	WeightKgs              sql.NullInt32  `json:"weight_kgs"`
	ExerciseCategory       sql.NullString `json:"exercise_category"`
	ExerciseID             int64          `json:"exercise_id"`
}

type Food struct {
	FoodID       int64         `json:"food_id"`
	FoodCalories sql.NullInt32 `json:"food_calories"`
}

type FoodCategory struct {
	FoodID          int64  `json:"food_id"`
	FoodCategory    string `json:"food_category"`
	FoodDescription string `json:"food_description"`
}

type FoodNutrient struct {
	FoodID            int64  `json:"food_id"`
	FoodNutrientsName string `json:"food_nutrients_name"`
	FoodQuantityUnit  string `json:"food_quantity_unit"`
	FoodQuantity      int32  `json:"food_quantity"`
}

type Liquid struct {
	LiquidID       int64 `json:"liquid_id"`
	LiquidCalories int32 `json:"liquid_calories"`
}

type LiquidCategory struct {
	LiquidID          int64  `json:"liquid_id"`
	LiquidCategory    string `json:"liquid_category"`
	LiquidDescription string `json:"liquid_description"`
}

type LiquidNutrient struct {
	LiquidID            int64  `json:"liquid_id"`
	LiquidNutrientsName string `json:"liquid_nutrients_name"`
	LiquidQuantityUnit  string `json:"liquid_quantity_unit"`
	LiquidQuantity      int32  `json:"liquid_quantity"`
}

type Meal struct {
	MealID   int64        `json:"meal_id"`
	MealTime sql.NullTime `json:"meal_time"`
}

type MealHasFood struct {
	MealMealID int64 `json:"meal_meal_id"`
	FoodFoodID int64 `json:"food_food_id"`
}

type MealHasLiquid struct {
	MealMealID     int64 `json:"meal_meal_id"`
	LiquidLiquidID int64 `json:"liquid_liquid_id"`
}

type User struct {
	UserID      int64        `json:"user_id"`
	Alias       string       `json:"alias"`
	Email       interface{}  `json:"email"`
	Password    string       `json:"password"`
	CreateTime  sql.NullTime `json:"create_time"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Cpf         interface{}  `json:"cpf"`
	PhoneNumber interface{}  `json:"phone_number"`
}
