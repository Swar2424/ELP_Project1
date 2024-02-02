const { checkPrime } = require("crypto");
const dico = require("./Dico.js")

async function Check_word(hand, tableau, name, len) {
    var rep = true
    var total_hand = [...hand]
    var tab_copy = [...tableau]

    for (i = 0 ; i < name.length ; i += 1) {
        rep = rep && (total_hand.includes(name[i]))

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
    dico.get_word(name).then( word =>{
        if (tab_copy.length != 0 || name.length <= tableau.length || word != name){
            rep = false;
        }
    })
    
    if (rep) {
        return [rep, total_hand] 
    } else {
        return [rep, hand] 
    }
}


module.exports = { Check_word };
/*
hands = ['m', 'n', 's', 'p', 'i', 'e']
noun = "pie"
len = 6
tableaux = [[[],[],[],[],[],[],[],[]],[[],[],[],[],[],[],[],[]]]
tableau = tableaux[0][1]
verif = Check_word(hands, tableau, noun, len)
verif.then(
    result => console.log(result[1]) 
)
*/
