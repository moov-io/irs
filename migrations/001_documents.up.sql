create table documents(
  document_id varchar(40) primary key not null,

  pdf  blob not null,
  json blob,

  created_at timestamp not null,
  deleted_at timestamp
);
