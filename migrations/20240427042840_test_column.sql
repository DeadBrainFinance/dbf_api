-- +goose Up
-- +goose StatementBegin
delete from transaction where transaction.id = 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from transaction where transaction.id == 1;
-- +goose StatementEnd
