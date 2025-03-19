import {model} from "../wailsjs/go/models";
import Cell from "./Cell";

export default function Grid({ grid, onCellClick}: { grid: model.Grid | null, onCellClick: (row: number, col: number) => void }) {
    if (grid == null) {
        return (
            <></>
        )
    }

    const uiGrid = [];
    const x = grid.Cells[0].length;
    const y = grid.Cells.length;
    for (let row = 0; row < y; row++) {
        const cells = [];
        for (let col = 0; col < x; col++) {
            const cell = grid.Cells[row][col]
            //TODO use UUID for key
            cells.push(
                <Cell cell={cell} row={row} col={col} onClick={onCellClick}/>
            );
        }
        uiGrid.push(
            //TODO also use UUID as key
            <div className={`row-${row}`} key={row} style={{display: "flex"}}>
                {cells}
            </div>
        );
    }

    return <div>{uiGrid}</div>;
}