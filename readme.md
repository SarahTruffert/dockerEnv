Vous développez des fonctionnalités pour une appli qui va utiliser une base de données SQL (moteur au choix : PostgreSQL, MySQL, SQL Server...).

La base de donnée de production sera probablement déployée dans le cloud, mais vous avez besoin d'un moyen rapide et efficace de tester votre code qui utilise une base de données.

Vous allez mettre en place un Dockerfile qui va lancer un conteneur avec un moteur de bases de données, récupérer un fichier .sql de votre choix contenant la structure d'une base de données (les "CREATE TABLE" nécessaires, éventuellement des "INSERT" pour ajouter des données), le lancer sur votre base de données.

Ce qui vous permettra en une seule commande de répliquer la dernière version de votre base de données de production.

Vous devrez aussi fournir un petit programme (dans le langage de programmation que vous utilisez en entreprise) qui va se connecter à votre base de données containérisée, et y effectuer quelques opérations : au minimum "SELECT" et "INSERT".

​

Bonus 1 : utilisez un fichier ".env" pour passez les bonnes variables à votre Dockerfile, et utilisez le même fichier dans votre script (de cette manière votre code et votre base de données utiliseront les mêmes variables sans avoir besoin de les recopier).

​

Bonus 2 : dockerisez votre exemple de code qui intéragit avec la base de données (pour faciliter l'automatisation des tests).

​

Bonus 3 (avancé) : utilisez un ORM (ex : SQLAlchemy) pour générer votre schéma de bases de données. Cela vous demandera un Dockerfile multi-stage (avec un conteneur qui construit votre base de données, et un autre qui récupère les informations dans une nouvelle base de données).




COMMANDS : 

docker build --tag imagesarah .
