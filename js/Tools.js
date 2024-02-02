function table_to_str(tableau) {
    var writeable = "- - - - - - - - - - - - - - - -\n"
    for (j=0 ; (j<tableau.length) ; j+=1) {
        writeable += tableau[j] + "\n"
    }
    writeable += "- - - - - - - - - - - - - - - -\n"
    return writeable
}

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


module.exports = { table_to_str, shuffle };