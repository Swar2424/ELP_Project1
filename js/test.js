function create() {
    reponse = "Letter choose : ";
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


function not_lettres(letters, name) {
    var rep = true
    console.log(name)
    for (i = 0 ; i < name.length ; i += 1) {
        rep = rep && (letters.includes(name[i]))
        if (rep){
            var j = 0;
            var r = letters.length;
            while  (letters.length == r){
                if (letters[j] == name[i]){
                    letters.pop(j);
                }
                j += 1;
            }
        }
    };
    return rep
}


function enter_letter(letters) {
    readline.question(`${letters} :`, name => {
        if (not_lettres(letters, name)) {
            let a = create()
            let x = a(name)
            console.log(x)
            fs.writeFile('./test.txt', x, err => {
                if (err) {
                console.error(err);
                } else {
                // file written successfully
                }
            });
            readline.close();
        } else {
            enter_letter(letters)
        }});
}


const fs = require('node:fs');


const lettres = [["a", 14], ["b", 4], ["c", 7] , ["d",5], ["e",19], ["f",2], ["g",4], ["h",2], ["i",11], ["j",1], ["k",1], ["l",6], ["m",5], ["n",9], ["o",8], ["p",4], ["q", 1], ["r", 10], ["s", 7], ["t", 9], ["u", 8], ["v", 2], ["w", 1], ["x", 1], ["y", 1], ["z", 2]]
let pioche = []
for (i=0; i < lettres.length; i += 1 ){
    let l = lettres[i][0], nb = lettres[i][1];
    for (j = 0; j < nb; j += 1){
        pioche.push(l)
    }
}

shuffle(pioche)
var a = pioche.splice(0,6)
enter_letter(a)