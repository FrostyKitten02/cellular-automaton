import {model} from "../wailsjs/go/models";
import CellUtil from "./utils/CellUtil";


export default function Cell({cell, row, col, onClick, gameMode}: {
    cell: model.Cell,
    row: number,
    col: number,
    onClick: (row: number, col: number) => void,
    gameMode: string
}) {
    const size = 20;
    const color = CellUtil.getCellColor(cell, gameMode)

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