-- +goose Up
-- +goose StatementBegin
alter table category
add description varchar(200);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table category
drop column description;
-- +goose StatementEnd
