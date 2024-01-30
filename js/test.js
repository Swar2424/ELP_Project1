function create(i) {
    var reponse = `Letter chosen by ${i}: `;
    return function (x) { return reponse + x ; }
}

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

function write_in(tableau, name) {
    var looping = true
    for (i=0 ; (i<tableau.length && looping) ; i+=1) {
        if (tableau[i].length == 0) {
            tableau[i] = name
            looping = false
        }
    }
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


function not_lettres(hand, tableau, name) {
    var rep = true
    var total_hand = [...hand]
    tab_copy = [...tableau]
    console.log(name)
    for (i = 0 ; i < name.length ; i += 1) {
        rep = rep && (total_hand.includes(name[i]))
        if (rep){
            var j = 0;
            var r = total_hand.length;
            while  (total_hand.length == r){
                if (total_hand[j] == name[i]){
                    if (tab_copy.includes(total_hand[j])){ 
                        var k = 0
                        var l = tab_copy.length
                        while (tab_copy.length == l){
                            if (tab_copy[k]==total_hand[j]){
                                tab_copy.splice(k,1);
                            }
                            k += 1;
                        }
                    }
                    total_hand.splice(j,1);
                }
                j += 1;
            }
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


function enter_letter(letters, hands, tableaux, n) {
    //prise de l'input
    console.log(`${hands[n]}`)
    readline.question(`${tableau_to_str(tableaux[n])}\n${tableau_to_str(tableaux[(n+1)%2])}\n Choose row :`, row => {
        row = parseInt(row) - 1
        console.log(row)
        let len = hands[n].length
        if ((row >= 1 && tableaux[n][row-1].length != 0 || row == 0) && (row < tableaux[n].length)) {
                if (tableaux[n][row].length != 0){
                    var row_split_copy = [...tableaux[n][row]]
                    hands[n] = hands[n].concat(row_split_copy)
                }
                //Il faudrait lui demander le numéro du joueur à partir du joueur 2
                readline.question(`${hands[n]} :`, name => {

                    //Si le joueur joue
                    if (name != "!"){
                        rep = not_lettres(hands[n], tableaux[n][row], name)
                        hands[n] = rep[1]

                        //Si il peut jouer ce mot
                        if (rep[0]) {
                            tableaux[n][row] = name
                            let a = create(n)
                            let x = a(name)
                            console.log(x)
                            hands[n].push(letters.splice(0, 1)[0])
                            var writeable = tableau_to_str(tableaux[n]) + tableau_to_str(tableaux[(n+1)%2])
                            fs.writeFile('./test.txt', writeable, err => {
                                if (err) {
                                console.error(err);
                                } else {
                                // file written successfully
                                }
                            });
                            
                        //Si il ne peut pas jouer ce mot
                        } else {
                            console.log("Invalide")
                            hands[n].splice(len, hands[n].length -len);
                        };

                        enter_letter(letters, hands, tableaux, n)

                    //Si il passe son tour
                    } else {
                    hands[n].push(letters.splice(0, 1)[0])
                    console.log(`Player ${(n+1)%2+1} is playing`)
                    enter_letter(letters, hands, tableaux, (n+1)%2);
                    }

                })
        } else {
            console.log("Invalide !")
            enter_letter(letters, hands, tableaux, n)
        }
    })
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

console.log(`Player 1 is playing`)
enter_letter(pioche, hands, tableaux, 0)