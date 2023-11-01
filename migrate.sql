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

CREATE table expenses (
    id                  text primary key,
    category_id         text,
    country             text,
    created             timestamp without time zone,
    currency            text,
    description         text,
    expense_date        timestamp without time zone,
    expense_end_date    timestamp without time zone,
    expense_type        text,
    is_credit_note      boolean,
    last_modified       timestamp without time zone,
    legal_entity_id     text,
    payment_method      text,
    posting_date        timestamp without time zone,
    status              text,
    tax_number          text,
    total_amount        integer,
    total_claim         integer,
    trip_id             text,
    user_id             text
);

CREATE table expense_cost_center (
    expense_id text,
    cost_center_id text,
    percent_weight integer,
    PRIMARY KEY (expense_id, cost_center_id)
);

CREATE table expense_tax_items (
    expense_id text,
    rate_id text,
    gross integer,
    tax integer,
    PRIMARY KEY (expense_id, rate_id)
);
