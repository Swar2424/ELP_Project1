function create(i) {
    var reponse = `Letter chosen by ${i}: `;
    return function (x) { return reponse + x ; }
}

const readline = require('node:readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});

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


function not_lettres(hand, name) {
    var rep = true
    var temp_hand = [...hand]
    console.log(name)
    for (i = 0 ; i < name.length ; i += 1) {
        rep = rep && (temp_hand.includes(name[i]))
        if (rep){
            var j = 0;
            var r = temp_hand.length;
            while  (temp_hand.length == r){
                if (temp_hand[j] == name[i]){
                    temp_hand.splice(j,1);
                }
                j += 1;
            }
        }
    };
    if (rep) {
        return [rep, temp_hand] 
    } else {
        return [rep, hand] 
    }
}


function enter_letter(letters, hands, n) {
    //prise de l'input
    readline.question(`${hands[n]} :`, name => {

        //Si le joueur joue
        if (name != "!"){
            rep = not_lettres(hands[n], name)
            hands[n] = rep[1]

            //Si il peut jouer ce mot
            if (rep[0]) {
                let a = create(n)
                let x = a(name)
                console.log(x)
                hands[n].push(letters.splice(0, 1)[0])
                fs.writeFile('./test.txt', x, err => {
                    if (err) {
                    console.error(err);
                    } else {
                    // file written successfully
                    }
                });
                
            //Si il ne peut pas jouer ce mot
            } else {
                console.log("Invalide")
            };

            enter_letter(letters, hands, n)

        //Si il passe son tour
        } else {
        hands[n].push(letters.splice(0, 1)[0])
        enter_letter(letters, hands, (n+1)%2);
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
hands[0] = pioche.splice(0, 6)
hands[1] = pioche.splice(0, 6)

enter_letter(pioche, hands, 0)