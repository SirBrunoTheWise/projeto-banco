CREATE DOMAIN email_domain AS VARCHAR(255)
    CHECK (VALUE ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');


CREATE DOMAIN cpf_domain AS VARCHAR(11)
    CHECK (VALUE ~ '^\d{11}$');


CREATE DOMAIN phone_domain AS VARCHAR(15)
    CHECK (VALUE ~ '^\+?[1-9]\d{1,14}$');

CREATE DOMAIN progression_domain AS SMALLINT
    CHECK (VALUE BETWEEN 0 AND 100);


CREATE DOMAIN year_domain AS SMALLINT
    CHECK (VALUE >= EXTRACT(YEAR FROM CURRENT_DATE));


CREATE DOMAIN cvv_domain AS SMALLINT
    CHECK (VALUE BETWEEN 100 AND 999);



CREATE TABLE IF NOT EXISTS users (
    user_ID BIGSERIAL PRIMARY KEY,
    alias VARCHAR(16) NOT NULL,
    email email_domain NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    cpf cpf_domain NOT NULL UNIQUE,
    phone_number phone_domain
);


CREATE TABLE IF NOT EXISTS credit_card (
    credit_card_name VARCHAR(45) NOT NULL,
    credit_card_number VARCHAR(19) NOT NULL,  -- Allowing space for the 16 digits plus potential spaces
    credit_card_CVV cvv_domain NOT NULL,
    credit_card_expiration_month SMALLINT CHECK (credit_card_expiration_month BETWEEN 1 AND 12),
    credit_card_expiration_year year_domain NOT NULL,
    user_user_ID BIGINT NOT NULL,
    PRIMARY KEY (user_user_ID),
    FOREIGN KEY (user_user_ID) REFERENCES users (user_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cards (
    card_ID BIGSERIAL PRIMARY KEY,
    card_type SMALLINT NOT NULL,
    card_number BIGINT NOT NULL,
    card_progression progression_domain NOT NULL,
    card_image BYTEA NOT NULL
);

CREATE TABLE IF NOT EXISTS diary (
    date_of TIMESTAMP NOT NULL,
    user_ID BIGINT NOT NULL,
    exercise BIGINT NOT NULL UNIQUE,
    meal BIGINT NOT NULL UNIQUE,
    cards BIGINT NOT NULL,
    PRIMARY KEY (date_of, exercise, meal),
    FOREIGN KEY (user_ID) REFERENCES users (user_ID) ON DELETE CASCADE,
    FOREIGN KEY (cards) REFERENCES cards (card_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS meal (
    meal_ID BIGSERIAL PRIMARY KEY,
    meal_time TIMESTAMP UNIQUE,
    FOREIGN KEY (meal_ID) REFERENCES diary (meal) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS exercise (
    exercise_ID BIGSERIAL PRIMARY KEY,
    time_start TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_end TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    calories_spent INT NOT NULL,
    average_heart_rate INT NOT NULL,
    exercise_type VARCHAR(45),
    FOREIGN KEY (exercise_ID) REFERENCES diary (exercise) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS liquid (
    liquid_ID BIGSERIAL PRIMARY KEY,
    liquid_calories INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS food (
    food_ID BIGSERIAL PRIMARY KEY,
    food_calories INT
);


CREATE TABLE IF NOT EXISTS liquid_category (
    liquid_ID BIGINT NOT NULL,
    liquid_category VARCHAR(45) NOT NULL,
    liquid_description VARCHAR(45) NOT NULL,
    PRIMARY KEY (liquid_ID),
    FOREIGN KEY (liquid_ID) REFERENCES liquid (liquid_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS liquid_nutrients (
    liquid_ID BIGINT NOT NULL,
    liquid_nutrients_name VARCHAR(45) NOT NULL,
    liquid_quantity_unit VARCHAR(45) NOT NULL,
    liquid_quantity INT NOT NULL,
    PRIMARY KEY (liquid_ID),
    FOREIGN KEY (liquid_ID) REFERENCES liquid (liquid_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS food_category (
    food_ID BIGINT NOT NULL,
    food_category VARCHAR(45) NOT NULL,
    food_description VARCHAR(45) NOT NULL,
    PRIMARY KEY (food_ID),
    FOREIGN KEY (food_ID) REFERENCES food (food_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS food_nutrients (
    food_ID BIGINT NOT NULL,
    food_nutrients_name VARCHAR(45) NOT NULL,
    food_quantity_unit VARCHAR(45) NOT NULL,
    food_quantity INT NOT NULL,
    PRIMARY KEY (food_ID),
    FOREIGN KEY (food_ID) REFERENCES food (food_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS meal_has_liquid (
    meal_meal_ID BIGINT NOT NULL,
    liquid_liquid_ID BIGINT NOT NULL,
    PRIMARY KEY (meal_meal_ID, liquid_liquid_ID),
    FOREIGN KEY (meal_meal_ID) REFERENCES meal (meal_ID) ON DELETE CASCADE,
    FOREIGN KEY (liquid_liquid_ID) REFERENCES liquid (liquid_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS meal_has_food (
    meal_meal_ID BIGINT NOT NULL,
    food_food_ID BIGINT NOT NULL,
    PRIMARY KEY (meal_meal_ID, food_food_ID),
    FOREIGN KEY (meal_meal_ID) REFERENCES meal (meal_ID) ON DELETE CASCADE,
    FOREIGN KEY (food_food_ID) REFERENCES food (food_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS exercise_weightlifting (
    sheet_type VARCHAR(45),
    exercise_name VARCHAR(45),
    sets INT,
    reps INT,
    time_between_sets_seconds INT,
    time_between_reps_seconds INT,
    weight_kgs INT,
    exercise_category VARCHAR(45),
    exercise_ID BIGINT NOT NULL,
    FOREIGN KEY (exercise_ID) REFERENCES exercise (exercise_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS exercise_hiking (
    coordinates POINT[],
    pace_seconds INT NOT NULL,
    distance_traveled_meters INT NOT NULL,
    wilderness_type VARCHAR(45),
    exercise_ID BIGINT NOT NULL,
    FOREIGN KEY (exercise_ID) REFERENCES exercise (exercise_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS exercise_jogging (
    coordinates POINT[],
    pace_seconds INT NOT NULL,
    distance_traveled_meters INT NOT NULL,
    cadence INT,
    terrain_level INT,
    exercise_exercise_ID BIGINT NOT NULL,
    FOREIGN KEY (exercise_exercise_ID) REFERENCES exercise (exercise_ID) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS exercise_biking (
    coordinates POINT[],
    pace_seconds INT NOT NULL,
    distance_traveled_meters INT NOT NULL,
    wheel_wear_level SMALLINT,
    terrain_type VARCHAR(45),
    exercise_ID BIGINT NOT NULL,
    FOREIGN KEY (exercise_ID) REFERENCES exercise (exercise_ID) ON DELETE CASCADE
);
