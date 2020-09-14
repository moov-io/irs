create table documents(
  document_id varchar(40) primary key not null,

  pdf   blob,
  ascii blob not null,

  created_at timestamp not null,
  deleted_at timestamp
);
