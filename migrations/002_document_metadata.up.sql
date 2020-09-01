create table document_metadata(
  document_id varchar(40) primary key not null,

  key   varchar(40) not null,
  value varchar(512) not null
);
