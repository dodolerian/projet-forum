
# Projet Forum Open Chat Room 📄 :
Le projet forum est le projet de fin d'année en Bachelor 1 informatique.
Le but du projet est de construire un forum fonctionnel.
Il doit pouvoir contenir:
- Un échange entre plusieurs utilisateurs
- Différentes catégories de postes
- Pouvoir liker et deliker un poste
- Filtrer les postes
Le projet doit égallement contenir une base de donné en SQLITE

## Membres du groupe 👨‍💻 :

- Sulien Payraudeau
- Matias Bellaud
- Maxime Fuzeau 
- Dorian Martin

## Langages 🛠️ :
Le projet à été réalisé en golang pour le back , html, css , js  pour le front et sqLite pour la base de donnée.
Le projet a égallement été dockérisé.

## Comment lancer le projet 🚀 :
__Si le projet n'est pas avec le docker:__ 

- Cloner le git: https://github.com/dodolerian/projet-forum

- Faire ``cd forum``

- Faire ``go mod init forum ``puis `` go mod tidy`` dans le terminal au premier lancement du projet.

- Faire ``go run main/main.go ``dans le terminal.

- Aller sur `` http://localhost:3333/`` 

__Si il est avec docker:__

- Cloner le git: https://github.com/dodolerian/projet-forum

- Ouvir docker.

- Faire ``cd forum``

- Faire ``go mod init forum ``puis `` go mod tidy`` dans le terminal au premier lancement du projet.

- Faire ``cd ..``

- Faire la commande suivante dans le terminale ``docker build --no-cache  -t forum:v3 . ``

- Puis faire ``docker run -p 3333:3333 forum:v3 ``

- Quand dans la console est écrit ``Starting server at port 3333 : http://localhost:3333``, aller sur `` http://localhost:3333/`` 


## Que contient le projet 📖 :
L'utilisateur a la possibilité de créer un compte, de se connecter et se déconnecter.
Quand il est connecté il peut liker et commenter un poste.
Sur sa page de profil il peut égallement ajouter un poste.
Il pourra égallement ajouter un tag au poste.
L'utilisateur peut aussi modifier sa description. 
Le forum est équipé d'un système de rank en fonction du nombre de postes fait par l'utilisateur. La photo de profil dépend du rank de l'utilisateur. 
Il y a égallement un système de censure de certains mots dans les commantaires des postes.
Quand nous cliquons sur la photo de profil, nous pouvons voir la description de l'utilisateur. 
Le forum dispose d'une connection invité qui permet juste de voir les postes et les descriptions.
L'utilisateur du forum peut égallement trié les postes en fonction des tags.