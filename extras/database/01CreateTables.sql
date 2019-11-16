CREATE TABLE epicfda.FoodEnforcement(
  FoodEnforcementID           UUID primary key,
  ,classification             VARCHAR  
  ,center_classification_date DATE  
  ,report_date                DATE  
  ,postal_code                VARCHAR  
  ,termination_date           DATE  
  ,recall_initiation_date     DATE  
  ,recall_number              VARCHAR 
  ,city                       VARCHAR 
  ,more_code_info             VARCHAR
  ,event_id                   int  
  ,distribution_pattern       VARCHAR 
  ,recalling_firm             VARCHAR 
  ,voluntary_mandated         VARCHAR 
  ,state                      VARCHAR 
  ,reason_for_recall          VARCHAR 
  ,initial_firm_notification  VARCHAR 
  ,status                     VARCHAR 
  ,product_type               VARCHAR 
  ,country                    VARCHAR 
  ,product_description        VARCHAR 
  ,code_info                  VARCHAR 
  ,address_1                  VARCHAR 
  ,address_2                  VARCHAR 
  ,product_quantity           VARCHAR 
);
