insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$YS.CNcTZ09vDtnoY/0axzui2zywYgsLnIyvrncJ043kB2sdRn.He."),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$YS.CNcTZ09vDtnoY/0axzui2zywYgsLnIyvrncJ043kB2sdRn.He."),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$YS.CNcTZ09vDtnoY/0axzui2zywYgsLnIyvrncJ043kB2sdRn.He.");


insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacação do Usuário 1", "Conteudo da publicação do usuario 1", 1),
("Publicacação do Usuário 2", "Conteudo da publicação do usuario 2", 2),
("Publicacação do Usuário 3", "Conteudo da publicação do usuario 3", 3);