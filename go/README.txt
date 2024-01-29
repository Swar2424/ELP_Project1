Notre projet est l'implémentation de l'algorithme de dijkstra sur un graphe pondéré,
suivi d'un algorithme de comptage qui calcul le nombre de chemins établits par dijkstra
passant par chaque sommet du graphe.

Le dossier contient 3 implémentations différentes de notre algorithme :
 - Le dossier normal_dijkstra pour la version no parallélisée
 - Le dossier goroutine_dijkstra pour une version parallélisée avec une worker pool
 - Les dossiers client_dijkstra et serveur_dijkstra pour la version client/serveur TCP

Tout d'abord, pour générer des matrices d'adjacence de grande taille, il faut exécuter dans le dossier /go la commande :
go run .\Main.go .\Funcs.go

Pour les versions avec et sans goroutines, il faut exécuter dans leurs dossiers respectifs :
go run .\Main.go .\Funcs.go .\Dijkstra.go .\Counting.go .\Load_data.go

Pour la version serveur/client TCP, il faut d'abord exécuter dans server_dijkstra :
go run .\Main.go .\Funcs.go .\Dijkstra.go .\Counting.go .\Load_data.go
Puis dans client_dijkstra, autant de fois que l'on le souhaite, exécuter :
go run .\Main.go .\Funcs.go .\Load_data.go 

Pour voir clairement l'écart de vitesse entre les versions parallélisées et la version standart, il faut faire un graphe
de plus de 1000 sommets. Avec 2000, la différence est très nette, même si l'algorithme prend un certain temps à s'exécuter.