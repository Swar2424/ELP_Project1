const fs = require('node:fs');
const check = require("./Checkword.js")
const tool = require("./Tools.js")
const end = require("./End.js")
const readline = require('node:readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});


//Fonction implémentant readline avec Promise ;
//Prend en paramètre le string <questions> affiché lors de la demande de l'input ;
//Renvoie la Promise associé à l'événement de l'input
function ask(questions) {
    return new Promise(resolve => {
        readline.question(questions, input => resolve(input));
    });
}


//Fonction principale du programme, récursive tant que le jeu n'est pas fini ;
//Prend en paramètre <letters> la pile de lettres à piocher, <hands> la liste des deux listes qui sont les mains des deux joueurs ;
//<tableaux> qui contient les deux tableaux des joueurs, <n> le numéro du joueur actuel, <jarnac> qui indique si un jarnac est en cours
function player_turn(letters, hands, tableaux, n, jarnac) {

    if (tableaux[0][7].length != 0 || tableaux[1][7].length != 0) {
        end.end_game(tableaux)

    } else {
        writeable = "Player 1 :\n" + hands[0].toString() + "\n" + tool.table_to_str(tableaux[0])
        writeable += "\nPlayer 2 :\n" + hands[1].toString() + "\n" + tool.table_to_str(tableaux[1])
        fs.writeFile('./data.txt', writeable, err => {
            if (err) {
                console.error(err);
            } else {
            }
            });
        console.log(`${hands[n]}`)
        ask(`\nPlay (p) or end turn (e) : `).then(choix => {
            if (choix=="p"){

                ask(`Choose row : `).then(row => {
                    row = parseInt(row) - 1
                    let len = hands[n].length
                    if ((row >= 1 && tableaux[n][row-1].length != 0 || row == 0) && (row < tableaux[n].length)) {
                            var row_split_copy = []
                            if (tableaux[n][row].length != 0){
                                row_split_copy = [...tableaux[n][row]]
                            }

                            ask(`${hands[n]} | ${row_split_copy} : `).then(name => {
                                if (name.length >= 3){
                                    hands[n] = hands[n].concat(row_split_copy)
                                    rep = check.Check_word(hands[n], tableaux[n][row], name, len)
                                    hands[n] = rep[1]

                                    if (rep[0]) {
                                        console.log(`Word played by ${i+1} : ${name}`)

                                        if (jarnac) {

                                            var replacing = true
                                            j = 0
                                            while (replacing) {
                                                if (tableaux[(n+1)%2][j].length == 0) {
                                                    tableaux[(n+1)%2][j] = name
                                                    replacing = false
                                                }
                                                j += 1
                                            }

                                            replacing = true
                                            k = row
                                            while (replacing) {
                                                if (tableaux[n][k+1].length == 0){
                                                    tableaux[n][k] = []
                                                    replacing = false
                                                } else {
                                                    tableaux[n][k] = [...tableaux[n][k+1]]
                                                    k += 1
                                                }
                                            }
                                        } else {
                                            tableaux[n][row] = name
                                            hands[n].push(letters.splice(0, 1)[0])
                                        }
                                        player_turn(letters, hands, tableaux, n, jarnac)
                                    
                                    } else {
                                        console.log("\nInvalide !\n")
                                        hands[n].splice(len, hands[n].length -len);
                                        player_turn(letters, hands, tableaux, n, jarnac)
                                    };
                                } else {
                                    console.log("\nInvalide !\n")
                                    player_turn(letters, hands, tableaux, n, jarnac)
                                }  
                        });
                    } else {
                        console.log("\nInvalide !\n")

                        player_turn(letters, hands, tableaux, n, jarnac)
                    } 
                });                          
                } else if (choix =="e"){
                    if (!jarnac) {
                        console.log(`\n----------------------------------------------------------------\nPlayer ${(n+1)%2+1} is playing\n`)

                        ask(`\nCall Jarnac ? (y/n) : `).then(jarnac_ask => {
                            hands[(n+1)%2].push(letters.splice(0, 1)[0])
                            if (jarnac_ask == "y") {
                                console.log("\nJARNAC !\n")
                                player_turn(letters, hands, tableaux, n, true)
                            } else {
                                player_turn(letters, hands, tableaux, (n+1)%2, false)
                            }
                        });
                    } else {
                        console.log("\nEnd of JARNAC !\n")
                        player_turn(letters, hands, tableaux, (n+1)%2, false)
                    }
                        
                } else {
                    console.log("\nInvalide !\n")
                    player_turn(letters, hands, tableaux, n, jarnac)
                };
        });
    }
}



const lettres = [["a", 14], ["b", 4], ["c", 7] , ["d",5], ["e",19], ["f",2], ["g",4], ["h",2], ["i",11], 
    ["j",1], ["k",1], ["l",6], ["m",5], ["n",9], ["o",8], ["p",4], ["q", 1], ["r", 10], ["s", 7], 
    ["t", 9], ["u", 8], ["v", 2], ["w", 1], ["x", 1], ["y", 1], ["z", 2]]

let pioche = []
for (i=0; i < lettres.length; i += 1 ){
    let l = lettres[i][0], nb = lettres[i][1];
    for (j = 0; j < nb; j += 1){
        pioche.push(l)
    }
}
tool.shuffle(pioche)


hands = [[],[]]
tableaux = [[[],[],[],[],[],[],[],[]],[[],[],[],[],[],[],[],[]]]
hands[0] = pioche.splice(0, 6)
hands[1] = pioche.splice(0, 6)

console.log(`\n----------------------------------------------------------------\nPlayer 1 is playing\n`)
player_turn(pioche, hands, tableaux, 0)