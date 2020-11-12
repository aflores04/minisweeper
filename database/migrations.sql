drop table if exists games;
create table games(
    id      int(11) primary key auto_increment,
    status  varchar(255) not null,
    rows    int(11) not null,
    cols    int(11) not null,
    mines   int(11) not null
);

drop table if exists point;
create table point(
    id      int(11) primary key auto_increment,
    row     int(11) not null,
    col     int(11) not null,
    mine    int(11) not null,
    flag    int(11) not null,
    value  int(11) not null,
    open    int(11) not null,
    game_id int(11) not null,
    foreign key (game_id) references games(id)
);