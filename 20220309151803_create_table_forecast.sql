-- +goose Up
-- +goose StatementBegin
CREATE TABLE forecasts (
    id int AUTO_INCREMENT NOT NULL,
    requestDate date,
    forecastDate date,
    min float,
    max float,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE forecasts;
-- +goose StatementEnd
