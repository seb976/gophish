-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE imap ADD COLUMN oauth2_enabled boolean DEFAULT 0;
ALTER TABLE imap ADD COLUMN tenant_id varchar(255);
ALTER TABLE imap ADD COLUMN client_id varchar(255);
ALTER TABLE imap ADD COLUMN client_secret varchar(512);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE imap DROP COLUMN oauth2_enabled;
ALTER TABLE imap DROP COLUMN tenant_id;
ALTER TABLE imap DROP COLUMN client_id;
ALTER TABLE imap DROP COLUMN client_secret;
