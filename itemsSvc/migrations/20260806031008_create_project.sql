-- +goose Up
-- goose Up

create table if not exists slots_type (
                                          id uuid primary key DEFAULT gen_random_uuid(),
    title character varying,
    created_at timestamp default now(),
    updated_at timestamp default now()
    );

create table if not exists slots_rarity (
    id uuid primary key default gen_random_uuid(),
    title character varying not null unique,
    base_weight float not null,
    multiplier float default 1.0,
    color_hex varchar(7) default '#ffffff',
    created_at timestamp default now(),
    updated_at timestamp default now()
    );

create table if not exists slots (
                                     id uuid primary key default gen_random_uuid(),
    title character varying not null,
    slot_type_id  uuid references slots_type(id) not null,
    slot_rarity_id uuid references slots_rarity(id) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
    );

create table if not exists shops (
                                     id uuid primary key default gen_random_uuid(),
    title character varying not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
    );

create table if not exists shops_slots (
                                           slot_id uuid not null references slots(id) on delete cascade,
    shop_id uuid not null references shops(id) on delete cascade,
    price float default 0,
    amount int not null default 0,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    primary key (slot_id, shop_id)
    );

-- +goose StatementBegin
create or replace function update_updated_at_column()
returns trigger as $$
begin
	NEW.updated_at = NOW();
return NEW;
end
$$ language plpgsql;
-- +goose StatementEnd

create trigger trigger_update_slots_updated_at
    before insert or update on "slots"
                         for each row
                         execute function update_updated_at_column();

create trigger trigger_update_slots_type_updated_at
    before insert or update on "slots_type"
                         for each row
                         execute function update_updated_at_column();

create trigger trigger_update_slots_rarity_updated_at
    before insert or update on "slots_rarity"
                         for each row
                         execute function update_updated_at_column();


create trigger trigger_update_shops_updated_at
    before insert or update on "shops"
                         for each row
                         execute function update_updated_at_column();

create trigger trigger_update_shops_slots_updated_at
    before insert or update on "shops_slots"
                         for each row
                         execute function update_updated_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_update_slots_updated_at ON slots;
DROP TRIGGER IF EXISTS trigger_update_slots_type_updated_at ON slots_type;
DROP TRIGGER IF EXISTS trigger_update_slots_rarity_updated_at ON slots_rarity;
DROP TRIGGER IF EXISTS trigger_update_shops_updated_at ON shops;
DROP TRIGGER IF EXISTS trigger_update_shops_slots_updated_at ON shops_slots;


DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;
DROP TABLE IF EXISTS shops_slots;
DROP TABLE IF EXISTS shops;
DROP TABLE IF EXISTS slots;
DROP TABLE IF EXISTS slots_rarity;
DROP TABLE IF EXISTS slots_type;