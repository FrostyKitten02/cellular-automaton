import {model} from "../wailsjs/go/models";


export default function Cell({cell, row, col, onClick}: {
    cell: model.Cell,
    row: number,
    col: number,
    onClick: (row: number, col: number) => void
}) {
    const size = 20;
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
                     minWidth: size,
                     maxWidth: size,
                     minHeight: size,
                     maxHeight: size,
                     backgroundColor: color,
                     margin: 1
                 }}/>
    )
}