
export default function Grid({x, y} : {x: number, y: number}) {
    const grid = [];
    for (let row = 0; row < y; row++) {
        let cells = [];
        for (let col = 0; col < x; col++) {
            //TODO use UUID for key
            cells.push(<div className={`r${row}-c${col}`} key={`${row}-${col}`} style={{minWidth: 10, maxWidth: 10, minHeight: 10, maxHeight: 10, backgroundColor: "#000000", margin: 1}}/>);
        }
        grid.push(
            //TODO also use UUID as key
            <div className={`row-${row}`} key={row} style={{display: "flex"}}>
                {cells}
            </div>
        );
    }

    return <div>{grid}</div>;
}