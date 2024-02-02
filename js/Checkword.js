//Vérifie si un mot peut être joué dans une ligne du tableau ; 
//Prend en paramètre la main du joueur, la ligne du tableau concernée, le mot joué et la taille de la main du joueur ;
//Renvoie <true> et la nouvelle main du joueur si le mot peut être joué ;
//Renvoie <false> et la main initiale du joueur sinon
function Check_word(hand, tableau, name, len) {
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
    if (tab_copy.length != 0 || name.length <= tableau.length){
        rep = false;
    }
    if (rep) {
        return [rep, total_hand] 
    } else {
        return [rep, hand] 
    }
}

module.exports = { Check_word };