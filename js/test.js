const readline = require('node:readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});


function tableau_to_str(tableau) {
    var writeable = "- - - - - - - - - - - - - - - -\n"
    for (j=0 ; (j<tableau.length) ; j+=1) {
        writeable += tableau[j] + "\n"
    }
    writeable += "- - - - - - - - - - - - - - - -\n"
    return writeable
}


//Mélanger la pioche de mots
function shuffle(array) {
    let currentIndex = array.length,  randomIndex;

    while (currentIndex > 0) {

      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex--;

      [array[currentIndex], array[randomIndex]] = [
        array[randomIndex], array[currentIndex]];
    }
  
    return array;
}


function not_lettres(hand, tableau, name, len) {
    var rep = true
    var total_hand = [...hand]
    tab_copy = [...tableau]
    console.log(name)

    //Pour chaque lettre dans le mot à jouer (name)
    for (i = 0 ; i < name.length ; i += 1) {
        rep = rep && (total_hand.includes(name[i]))

        //Si la lettre est disponible
        if (rep){ 
            var j = total_hand.indexOf(name[i]);
            if (tab_copy.includes(name[i])) {
                var k = tab_copy.indexOf(name[i])
                tab_copy.splice(k,1)
                j = len + k
            } else {
                len = len - 1
            }
            total_hand.splice(j,1);
        }
    };
    if (tab_copy.length != 0){
        rep = false;
    }
    if (rep) {
        return [rep, total_hand] 
    } else {
        return [rep, hand] 
    }
}


function end_game(tableaux) {
    console.log("\n-----------------------------------------------------------------\nEND OF THE GAME\n\n")
    console.log(`${tableau_to_str(tableaux[0])}\n${tableau_to_str(tableaux[1])}\n`)
    var point_total = [0,0]
    for (n = 0; n < 2; n +=1){
        var i = 0
        while (i < tableaux[n].length && tableaux[n][i].length != 0){
            point = tableaux[n][i].length
            point_total[n] += point**2
            i += 1
        }
    }
    console.log(`Point Player 1 : ${point_total[0]} | Point Player 2 : ${point_total[1]}`)
    if (point_total[0] > point_total[1]){
        console.log(`Player 1 won !!!`)
    } else {
        console.log(`Player 2 won !!!`)
    }

}


function enter_letter(letters, hands, tableaux, n, jarnac) {
    //prise de l'input
    if (tableaux[n][3].length != 0 || tableaux[(n+1)%2][3].length != 0) {
        end_game(tableaux)
    } else {
        console.log(`${hands[n]}`)
        readline.question(`${tableau_to_str(tableaux[n])}\n Jouer (j) ou Passer (p) : `, choix => {
            if (choix=="j"){
                readline.question(`${tableau_to_str(tableaux[n])}\nChoose row : `, row => {
                    row = parseInt(row) - 1
                    let len = hands[n].length
                    if ((row >= 1 && tableaux[n][row-1].length != 0 || row == 0) && (row < tableaux[n].length)) {
                            var row_split_copy = []
                            if (tableaux[n][row].length != 0){
                                row_split_copy = [...tableaux[n][row]]
                            }
                            //Il faudrait lui demander le numéro du joueur à partir du joueur 2
                            readline.question(`${hands[n]} | ${row_split_copy} : `, name => {
                                if (name.length >= 3){
                                    hands[n] = hands[n].concat(row_split_copy)
                                    rep = not_lettres(hands[n], tableaux[n][row], name, len)
                                    hands[n] = rep[1]

                                    //Si il peut jouer ce mot
                                    if (rep[0]) {
                                        console.log(`Word played by ${i+1} : ${name}`)

                                        if (jarnac) {

                                            //Ajouter le mot dans notre tableau
                                            var replacing = true
                                            j = 0
                                            while (replacing) {
                                                if (tableaux[(n+1)%2][j].length == 0) {
                                                    tableaux[(n+1)%2][j] = name
                                                    replacing = false
                                                }
                                                j += 1
                                            }

                                        //Enlever le mot du tableau adverse
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

                                        var writeable = tableau_to_str(tableaux[n]) + tableau_to_str(tableaux[(n+1)%2])
                                        fs.writeFile('./test.txt', writeable, err => {
                                            if (err) {
                                            console.error(err);
                                            } else {
                                            // file written successfully
                                            }
                                            });

                                    } else {
                                        tableaux[n][row] = name
                                        hands[n].push(letters.splice(0, 1)[0])
                                        var writeable = tableau_to_str(tableaux[n]) + tableau_to_str(tableaux[(n+1)%2])
                                        fs.writeFile('./test.txt', writeable, err => {
                                            if (err) {
                                            console.error(err);
                                            } else {
                                            // file written successfully
                                            }
                                            });
                                    }

                                    enter_letter(letters, hands, tableaux, n, jarnac)
                                    
                                    //Si il ne peut pas jouer ce mot
                                    } else {
                                        console.log("\nInvalide !\n")
                                        hands[n].splice(len, hands[n].length -len);
                                        enter_letter(letters, hands, tableaux, n, false)
                                    };
                                } else {
                                    console.log("\nInvalide !\n")
                                    enter_letter(letters, hands, tableaux, n, false)
                                }  
                        });
                    } else {
                        console.log("\nInvalide !\n")

                        enter_letter(letters, hands, tableaux, n, false)
                    } 
                });                     

                //Si il passe son tour       
                } else if (choix =="p"){
                    if (!jarnac) {
                        console.log(`\n----------------------------------------------------------------\nPlayer ${(n+1)%2+1} is playing\n`)
                        readline.question(`${hands[n]}\n${tableau_to_str(tableaux[n])}\nCall Jarnac ? (Y/N) : `, jarnac_ask => {
                            hands[(n+1)%2].push(letters.splice(0, 1)[0])
                            if (jarnac_ask == "Y") {
                                console.log("\nJARNAC !\n")
                                enter_letter(letters, hands, tableaux, n, true)
                            } else {
                                enter_letter(letters, hands, tableaux, (n+1)%2, false)
                            }
                        });
                    } else {
                        console.log("\nEnd of JARNAC !\n")
                        enter_letter(letters, hands, tableaux, (n+1)%2, false)
                    }
                        
                } else {
                    console.log("\nInvalide !\n")
                    enter_letter(letters, hands, tableaux, n, jarnac)
                };
        });
    }
}



const fs = require('node:fs');

//Création de la pioche mélangée
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
shuffle(pioche)


hands = [[],[]]
tableaux = [[[],[],[],[],[],[],[],[]],[[],[],[],[],[],[],[],[]]]
hands[0] = pioche.splice(0, 6)
hands[1] = pioche.splice(0, 6)

console.log(`\n----------------------------------------------------------------\nPlayer 1 is playing\n`)
enter_letter(pioche, hands, tableaux, 0)