function create(i) {
    var reponse = `Letter choose by ${i}: `;
    return function (x) { return reponse + x ; }
}

const readline = require('node:readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});

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
    };
    return rep
}


function enter_letter(lettres, n) {
    var letters = lettres.splice(0, 6)
    readline.question(`${letters} :`, name => {
        if (not_lettres(letters, name)) {
            let a = create(n)
            let x = a(name)
            console.log(x)
            fs.writeFile('./test.txt', x, err => {
                if (err) {
                console.error(err);
                } else {
                // file written successfully
                }
            });
            
        } else {
            enter_letter(lettres, n)
        };
        enter_letter(lettres, (n+1)%2);
    })
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
enter_letter(pioche, 0)