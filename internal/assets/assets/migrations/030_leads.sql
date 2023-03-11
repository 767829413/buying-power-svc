-- +migrate Up

create table identity_to_brokers (
    broker text not null,
    identity text not null,
    primary key (identity, broker)
);

-- +migrate Down

drop table identity_to_brokers;
