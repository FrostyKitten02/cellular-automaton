import {main} from "../wailsjs/go/models";

export default function Grid({grid} : {grid: main.Grid | null}) {
    if (grid == null) {
        return (
            <></>
        )
    }

    const uiGrid = [];
    const x = grid.Cells.length;
    const y = grid.Cells[0].length;
    for (let row = 0; row < y; row++) {
        const cells = [];
        for (let col = 0; col < x; col++) {
            //TODO use UUID for key
            const cell = grid.Cells[row][col]

            let color = "#ff0000" //if red is displayed we have a problem
            if (cell.cellType === "DEAD") {
                color = "#ffffff"
            }

            if (cell.cellType == "ALIVE") {
                color = "#000000"
            }
            cells.push(
                <div className={`r${row}-c${col}`} key={`${row}-${col}`}
                     style={
                        {
                            minWidth: 10,
                            maxWidth: 10,
                            minHeight: 10,
                            maxHeight: 10,
                            backgroundColor: color,
                            margin: 1
                        }} />
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