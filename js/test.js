function create() {
    reponse = "BA";
    return function (x) { return x + reponse; }
}

const readline = require('node:readline').createInterface({
    input: process.stdin,
    output: process.stdout,
});



const fs = require('node:fs');





function enter_letter(letters) {
    readline.question(`${letters} :`, name => {
        if (letters.includes(name)) {
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


const lettres = ["a", "b", "c", "d"]

enter_letter(lettres)