CREATE TABLE IF NOT EXISTS ideas (
  id bigserial,
  category varchar(255) NOT NULL,
  description varchar(255) NOT NULL
);


INSERT INTO ideas (category, description) VALUES 
('projects', 'make more docker containers!');
