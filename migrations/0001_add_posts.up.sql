CREATE TABLE posts (
  id         BIGSERIAL NOT NULL,
  topic      VARCHAR(100)                NOT NULL,
  body       VARCHAR(255)                NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
