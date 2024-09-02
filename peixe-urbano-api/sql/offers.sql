CREATE DATABASE IF NOT EXISTS peixe_urbano;
USE peixe_urbano;

DROP TABLE IF EXISTS offers;

CREATE TABLE offers(
    id int auto_increment primary key,
    categoria varchar(50) not null,
    titulo varchar(50) not null ,
    descricao varchar(50) not null ,
    anunciante varchar(20) not null ,
    valor decimal(7,2) not null,
    destaque boolean not null
) ENGINE=INNODB;