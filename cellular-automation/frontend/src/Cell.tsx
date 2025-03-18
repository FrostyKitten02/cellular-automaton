import {main} from "../wailsjs/go/models";


export default function Cell({cell, row, col, onClick}: {
    cell: main.Cell,
    row: number,
    col: number,
    onClick: (row: number, col: number) => void
}) {
    let color = "#ff0000" //if red is displayed we have a problem
    if (cell.cellType === "DEAD") {
        color = "#ffffff"
    }

    if (cell.cellType == "ALIVE") {
        color = "#000000"
    }

    return (
        <div className={`r${row}-c${col}`} key={`${row}-${col}`}
             onClick={() => {
                 onClick(row, col)
             }}
             style={
                 {
                     minWidth: 10,
                     maxWidth: 10,
                     minHeight: 10,
                     maxHeight: 10,
                     backgroundColor: color,
                     margin: 1
                 }}/>
    )
}