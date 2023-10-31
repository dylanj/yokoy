CREATE TABLE legal_entities (
    id          text primary key,
    code        text,
    language    text,
    name        text,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone
);

CREATE TABLE users (
    id                      text primary key,
    cost_center_id          text,
    email                   text,
    employee_id             integer,
    first_name              text,
    language                text,
    last_name               text,
    legal_entity_id         text,
    line_manager_id         text,
    line_manager_threshold  integer,
    policy_id               text,
    status_active           boolean,
    submission_delegate_id  text
);

CREATE TABLE trips (
    id              text primary key,
    currency        text,
    start_date_time timestamp without time zone,
    end_date_time   timestamp without time zone,
    legal_entity_id text,
    name            text,
    status          text,
    total_claim     integer,
    user_id         text
);

CREATE table company_cards (
    id                  text primary key,
    account_reference   text,
    currency            text,
    description         text,
    name                text,
    number              text,
    owner_id            text,
    status_active       boolean,
    legal_entity_id     text
);

CREATE table categories (
    id                  text primary key,
    account_reference   text,
    charge_to_employee  boolean,
    name                text,
    legal_entity_id     text,
    status_active       boolean
);

CREATE table cost_center (
    id                  text primary key,
    approval_limit      integer,
    approver_id         text,
    auto_approval_limit integer,
    code                text,
    delegate_expiry     timestamp without time zone,
    delegate_id         text,
    description         text,
    name                text,
    parent_id           text,
    legal_entity_id     text,
    status_active       boolean
);
