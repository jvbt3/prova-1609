create table viagens
(
    id   serial primary key,
    name varchar(255),
    dataChegada varchar(255),
    dataSaida varchar(255),
    valor float
)