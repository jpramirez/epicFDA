
DROP TABLE epicfda.Product;
DROP TABLE epicfda.Consumer;
DROP TABLE epicfda.FoodEventReactions;
DROP TABLE epicfda.FoodEventOutcomes;
DROP TABLE FoodEvent;


CREATE TABLE epicfda.Product(
    FoodEventID   UUID Primary Key
    ,NameBrand             VARCHAR  
    ,IndustryCode           VARCHAR
    ,Role               VARCHAR
    ,IndustryName       VARCHAR
);

CREATE TABLE epicfda.Consumer(
    FoodEventID   UUID Primary Key
    ,Gender             VARCHAR  
    ,Age           VARCHAR
    ,AgeUnit               VARCHAR
);

CREATE TABLE epicfda.FoodEventReactions (
        FoodEventID  UUID Primary Key ,
        Name    VARCHAR
);

CREATE TABLE epicfda.FoodEventOutcomes (
    FoodEventID UUID Primary Key,
    Name    VARCHAR
);

CREATE TABLE epicfda.FoodEvent (
    FoodEventID UUID 
    ,ReportNumber   VARCHAR
    ,DateCreated    DATE
    ,DateStarted    DATE
    ,PRIMARY KEY (ReportNumber)
);

