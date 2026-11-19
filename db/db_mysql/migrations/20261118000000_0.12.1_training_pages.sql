-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS training_pages (id integer primary key auto_increment,user_id bigint,name varchar(255),html text,modified_date datetime );
ALTER TABLE pages ADD COLUMN training_page_id bigint;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE pages DROP COLUMN training_page_id;
DROP TABLE training_pages;
