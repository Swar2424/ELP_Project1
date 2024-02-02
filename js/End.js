const fs = require('node:fs');
const tool = require("./Tools.js")

function end_game(tableaux) {
    console.log("\n-----------------------------------------------------------------\nEND OF THE GAME\n\n")
    console.log(`${tool.table_to_str(tableaux[0])}\n${tool.table_to_str(tableaux[1])}\n`)
    var point_total = [0,0]
    for (n = 0; n < 2; n +=1){
        var i = 0
        while (i < tableaux[n].length && tableaux[n][i].length != 0){
            point = tableaux[n][i].length
            point_total[n] += point**2
            i += 1
        }
    }
    var writer = ""
    
    console.log(`Point Player 1 : ${point_total[0]} | Point Player 2 : ${point_total[1]}`)
    if (point_total[0] > point_total[1]){
        writer += `Player 1 won !!!`
        console.log(`Player 1 won !!!`)
    } else {
        writer += `Player 2 won !!!\n\n`
        console.log(`Player 2 won !!!`)
    }
    writer += `Point Player 1 : ${point_total[0]} | Point Player 2 : ${point_total[1]}\n\n`

    writer += "Player 1 :\n" + tool.table_to_str(tableaux[0]) + "Player 2 :\n" + tool.table_to_str(tableaux[1])
    writer 
    fs.writeFile('./data.txt', writer, err => {
        if (err) {
            console.error(err);
        } else {
            // file written successfully
        }
        });

    process.exit()
}

module.exports = { end_game };