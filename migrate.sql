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

CREATE table cost_centers (
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
    additional_charges  integer,
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

CREATE table expense_cost_centers (
    expense_id          text,
    cost_center_id      text,
    percent_weight      integer,

    PRIMARY KEY (expense_id, cost_center_id)
);

CREATE table expense_tax_items (
    expense_id          text,
    rate_id             text,
    gross               integer,
    tax                 integer,

    PRIMARY KEY (expense_id, rate_id)
);

CREATE table expense_approver_ids (
    expense_id          text,
    approver_id         text,

    PRIMARY KEY(expense_id, approver_id)
);

CREATE table expense_event_logs (
    id                  serial primary key,
    expense_id          text,
    action_type         text,
    comment             text,
    name                text,
    timestamp           timestamp without time zone,
    user_id             text
);

CREATE table policies (
    id                  text primary key,
    code                text,
    name                text,
    legal_entity_id     text,
    status_active       boolean
);

CREATE table tax_rates (
    id                  text primary key,
    account_reference   text,
    code                text,
    country             text,
    name                text,
    rate                integer,
    legal_entity_id     text,
    status_active       boolean
);

CREATE TABLE tags (
    id                  text primary key,
    dimension_code      text,
    code                text,
    name                text,
    legal_entity_id     text,
    status_active       boolean
);

CREATE TABLE invoices (
    id                  text primary key,
    legal_entity_id     text,
    country             text,
    currency            text,
    dueDate             timestamp without time zone,
    date                timestamp without time zone,
    gross_amount        integer,
    invoice_number      text,
    is_credit_node      boolean,
    net_amount          integer,
    payment_term_id     text,
    posting_date        timestamp without time zone,
    purchase_order_ids  text[],
    service_date        timestamp without time zone,
    status              text,
    submitters          text[],
    supplier_id         text,
    taxable_amount      integer,

    bank_account        text,
    bank_country        text,
    bank_key            text,
    bank_number         text,
    external_id         text,
    iban                text,
    swift_code          text
);

CREATE TABLE invoice_line_items (
    id                      serial primary key,
    invoice_id              text,
    category_id             text,
    cost_object_id          text,
    description             text,
    gross                   integer,
    item_price              integer,
    net                     integer,
    purchase_order_id       text,
    purchase_order_item_id  text,
    quantity                integer,
    tags                    text,
    tax_rate_id             text,
    unit                    text
);

CREATE TABLE invoice_categories (
    id                      text primary key,
    account_reference       text,
    description             text,
    name                    text,
    status_active           boolean,
    legal_entity_id         text
);

CREATE TABLE suppliers (
    id                      text,
    legal_entity_id         text,
    city                    text,
    country_code            text,
    external_id             text,
    name                    text,
    secondary_name          text,
    short_name              text,
    status_active           boolean,
    street                  text,
    tax_number              text,
    url                     text,
    zip_code                text,
    default_approver_id     text,
    default_category_id     text,
    default_cost_center     text,
    default_payment_term_id text,
    supplier_id             text,

    PRIMARY KEY (legal_entity_id, id)
);
