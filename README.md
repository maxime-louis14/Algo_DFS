```markdown
# Projet : Carte Exploratoire avec Obstacles en Go 🗺️

Ce projet simule un système de déplacement sur une carte 2D dans lequel un personnage cherche un chemin vers une sortie (exit). La carte est représentée sous forme d'une grille 30x30 avec des éléments spécifiques comme des obstacles (murs), un point de départ et une sortie. Le programme utilise une recherche en profondeur (DFS) pour explorer la carte et trouver le chemin.

## Fonctionnalités 🚀

- **Carte générée dynamiquement** : création d'une carte 30x30 initialisée avec des valeurs de base.
- **Positionnement d'obstacles et de bordures** : les murs bordent la carte et des obstacles internes ajoutent de la complexité au parcours.
- **Exploration automatique** : recherche du chemin vers la sortie en marquant les cases visitées.
- **Affichage console** : représentation visuelle de la carte avec le chemin parcouru, les obstacles et la position de l'exit.
- **Paramétrage flexible** : possibilité de changer les dimensions de la carte, les positions de départ et de sortie, ou d'ajouter des obstacles supplémentaires.

## Technologies Utilisées 🛠️

- **Go** : Langage principal pour le développement et l'algorithme de parcours.
- **Console** : Affichage visuel de la carte et du parcours directement dans le terminal.

## Aperçu du Fonctionnement

1. **Initialisation de la Carte** : la carte est générée avec une taille de 30x30 et chaque case est initialisée à `0` (espace libre).
2. **Définition des Points Importants** : un point de départ (`start`), un `nid` (point de repère), et une `exit` (sortie) sont positionnés.
3. **Ajout d'Obstacles** : les murs bordent la carte, et des obstacles supplémentaires sont placés pour créer des challenges.
4. **Algorithme de Parcours** : à partir du point de départ, la fonction de recherche explore les cases voisines (droite, gauche, haut, bas) jusqu'à trouver un chemin vers la sortie.
5. **Affichage du Résultat** : après le parcours, la carte finale est affichée, montrant le chemin pris, les obstacles et l'exit.

## Utilisation 🔧

Clonez ce dépôt et exécutez le projet en lançant la commande suivante depuis le répertoire du projet :

```bash
go run main.go
```

Le résultat s'affichera dans la console, avec le chemin parcouru indiqué par `1`, les murs par `1`, les espaces libres par `0`, et l'exit par `2`.

## Améliorations Futures 🌱

- Ajout d'interfaces graphiques pour une visualisation interactive de la carte.
- Personnalisation avancée des obstacles et des points de repère.
- Optimisation de l'algorithme pour de plus grandes cartes ou des obstacles plus complexes.
```