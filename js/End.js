const fs = require('node:fs');
const tool = require("./Tools.js")

//Fonction appelée à la fin du jeu, chargée d'afficher les scores et de déterminer le vainqueur ;
//Prends en paramètre les tableaux complets des deux joueurs dans la liste <tableaux> ;
//Met fin au programme
function end_game(tableaux) {
    console.log("\n-----------------------------------------------------------------\nEND OF THE GAME\n\n")
    var point_total = [0,0]
    for (n = 0; n < 2; n +=1){
        var i = 0
        while (i < tableaux[n].length && tableaux[n][i].length != 0){
            point = tableaux[n][i].length
            point_total[n] += point**2
            i += 1
        }
    }
    var writeable = ""
    
    if (point_total[0] > point_total[1]){
        writeable += `Player 1 won !!!\n\n`
    } else {
        writeable += `Player 2 won !!!\n\n`
    }
    writeable += `Player 1 Points : ${point_total[0]} | Player 2 Points : ${point_total[1]}\n\n`

    writeable += "Player 1 :\n" + tool.table_to_str(tableaux[0]) + "Player 2 :\n" + tool.table_to_str(tableaux[1])
    
    fs.writeFile('./data.txt', writeable, err => {
        if (err) {
            console.error(err);
        } else {
            process.exit()
        }
        });
}

module.exports = { end_game };