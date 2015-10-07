# Abstract
The aim of this project is to build a simple REST middleware that would take in input a json representation of a TOSCA Node Type and transfer it to submodules via a 0MQ interface.

*Note* The following description is in french and will be translated soon
# Schéma général de principe
<img src="https://cloud.githubusercontent.com/assets/11583401/10247578/53a444ea-691a-11e5-86de-016857f99738.png"/>

# A propos
1 : Un middleware qui sert des API REST en HTTP. 
Son but est de de faire exécuter des requêtes de création d’éléments d’infrastructure écrites en JSON dans une grammaire qui correspondrait aux node types TOSCA.
Les requètes sont transmises aux différents backend d’exécution (HP CSA, HP OO, VmWare, …) via des drivers (le détail des driver est expliqué dans le paragraphe 5)

# L’authentification
2: L’authentification des utilisateurs est décoréllée des rôles de ces derniers. Ainsi donc l’authentification se fera au niveau d’un SSO via SAMLV2 ou un autre protocole.
Une fois authentifié, le middleware vérifiera les accréditations soit dans une base locale, soit via un serveur d’accréditation (mécanique OauthV2 par exemple).
Un Token (de type JWT par exemple) sera alors généré et il permettra ainsi l’accès à certains périmètres de l’API.
Ainsi donc, il sera possible d’authentifier un utilisateur humain, qui récupérera un token qu’il passera à un automate d'exécution. Ceci permettra une traçabilité des actions via un système de log (par exemple logstash).
De même, à la fin de l’execution, le programme client, ou l’utilisateur pourront annuler le token (logout) rendant ainsi son utilisation impossible. Les tokens “blacklistés” seront gérés par une base NoSQL.

# La liste des offering
*Note*:  Les offerings ne répondent qu’aux méthodes HTTP GET

3 La liste des offering disponible sera intégrée dans une base de fichier YAML. Chaque offering correspondra à une définition de type Node Type TOSCA “allégée” des principe de requirements et de capacités (inutiles dans notre cas). Seules certains paramètres tels que les “attributes” ou les “properties” seront conservées.
La notion de derived_from sera utilisée pour déterminer le backend cible qui permettra d’instancier ce node type;

exemple:
```yaml
tosca.nodes.Compute.BACKEND1.dev:
 derived_from: tosca.nodes.Compute.BACKEND1
 attributes:
   private_address:
     type: string
   public_address:
     type: string
 properties:
   hostname:
     type: string
   VMtype:
     type: string
     valid_values:
       - gold
       - silver
```

Ce type permettra la création d’une VM en utilisant le BACKEND1 et attendra en paramètre le hostname ainsi que le type de la VM (gold ou silver) et retournera une fois terminé les adresses publiques et privées.
Génération d’une requète
4 Lors d’un appel de type POST /request/...  avec un payload JSON décrivant un node template TOSCA qui correspond à une offering, le middleware générera un identifiant unique d’execution qu’il stockera en tant que clé d’une base de donnée NoSQL. La valeur sera 0 pour “initial”.
L’identifiant est ensuite renvoyé en JSON au client.

Le middleware choisira ensuite la backend à utiliser en fonction de la valeur du derived_from du node type correspondant.
Le middleware utilisera ensuite une communication 0MQ vers le driver du backend à utiliser et lui passera le JSON de la requète en l’état. Chaque driver sera donc un programme indépendant, codé dans n’importe quel langage capable d’implémenter un server 0MQ 5.
Les driver auront pour tache de transformer si possible l’offering JSON en ordre du backend et transmettre l’ordre d’éxécution au backend concerné 6.
L’identifiant unique d’execution du backend (généré par le backend lui même ou à défaut par le driver) sera retransmis au middleware qui le stockera en tant que valeur de la clé de la requete en cours dans la base 0MQ 7; l’état de la requete sera également notée dans la base (Voirl e paragraphe 3.3.1 Node state de la norme TOSCA).

# Intérrogation de l’état de la requète
Lors d’une appel GET /request/requestID, le middleware effectue la même communication via le driver pour mettre à jour l’état de la base de donnée.
Si la requête prend plus longtemps qu’un temps donné, l’état retourné est celui de la base de donnée.


# API documentation
The API documentation can be generated via a [swagger](https://github.com/yvasiyarov/swagger) (see [Makefile](Makefile))
