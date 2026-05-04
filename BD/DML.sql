/* Inserts de series */
INSERT INTO series (titulo, sinopsis, episodios, pais_origen, genero_principal, portada_url) VALUES
('Breaking Bad', 'Un profesor de química se convierte en fabricante de metanfetaminas.', 62, 'USA', 'Drama', '/uploads/breaking_bad.jpg'),
('Stranger Things', 'Niños enfrentan fenómenos sobrenaturales en un pequeño pueblo.', 34, 'USA', 'Ciencia Ficcion', '/uploads/stranger_things.jpg'),
('Game of Thrones', 'Familias luchan por el control del Trono de Hierro.', 73, 'USA', 'Fantasia', '/uploads/got.jpg'),
('The Office', 'Comedia sobre la vida en una oficina de ventas de papel.', 201, 'USA', 'Comedia', '/uploads/the_office.jpg'),
('Dark', 'Viajes en el tiempo conectan generaciones en Alemania.', 26, 'Alemania', 'Ciencia Ficcion', '/uploads/dark.jpg'),
('Narcos', 'Historia del narcotráfico en Colombia.', 30, 'Colombia', 'Drama', '/uploads/narcos.jpg'),
('Attack on Titan', 'Humanos luchan contra gigantes devoradores.', 87, 'Japon', 'Anime', '/uploads/aot.jpg'),
('The Crown', 'Relato de la vida de la reina Isabel II.', 60, 'UK', 'Drama', '/uploads/the_crown.jpg'),
('Friends', 'Grupo de amigos vive situaciones cómicas en Nueva York.', 236, 'USA', 'Comedia', '/uploads/friends.jpg'),
('The Witcher', 'Un cazador de monstruos navega un mundo oscuro.', 24, 'USA', 'Fantasia', '/uploads/witcher.jpg'),
('Black Mirror', 'Historias sobre el impacto oscuro de la tecnología.', 27, 'UK', 'Ciencia Ficcion', '/uploads/black_mirror.jpg'),
('Better Call Saul', 'Origen del abogado Saul Goodman.', 63, 'USA', 'Drama', '/uploads/bcs.jpg');


/* Inserts de ratings */
INSERT INTO ratings (serie_id, puntaje) VALUES
(1,5),(1,5),(1,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(2,4),(2,5),(2,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(3,5),(3,4),(3,3);
INSERT INTO ratings (serie_id, puntaje) VALUES
(4,5),(4,4),(4,5);
INSERT INTO ratings (serie_id, puntaje) VALUES
(5,5),(5,5),(5,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(6,4),(6,4),(6,5);
INSERT INTO ratings (serie_id, puntaje) VALUES
(7,5),(7,5),(7,5);
INSERT INTO ratings (serie_id, puntaje) VALUES
(8,4),(8,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(9,5),(9,5),(9,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(10,4),(10,3),(10,4);
INSERT INTO ratings (serie_id, puntaje) VALUES
(11,5),(11,4),(11,5);
INSERT INTO ratings (serie_id, puntaje) VALUES
(12,5),(12,5),(12,4);