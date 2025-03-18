import Grid from "./Grid";
// @ts-ignore
import {ResetGrid, Simulate, Step, Increment, InitGrid, Init2} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";
import {useEffect, useState} from "react";


function App() {
    const [grid, setGrid] = useState<main.Grid | null>(null)

    useEffect(() => {
        console.log("Initing")
        InitGrid(20,20).then( res => {
           setGrid(res);
           console.log(res);
        });
    }, [])


    if (grid == null) {
        return (
            <div>
                LOADING...
            </div>
        )
    }


    return (
        <div>
            <button className="btn" onClick={() => {
                Simulate().then((res: main.Grid) => {
                    //TODO rerender grid
                    console.log("Resolved simulate")
                })
            }}>
                Simulate
            </button>
            <button className="btn" onClick={() => {
                Step().then((res: main.Grid) => {
                    //TODO rerender grid
                    setGrid(res);
                    console.log(res)
                    console.log("resolved step")
                })
            }}>
                Step
            </button>
            <button className="btn" onClick={() => {
                ResetGrid().then((res: main.Grid) => {
                    //TODO rerender grid
                    console.log("resolved reset")
                })
            }}>
                Reset
            </button>
            <div style={{display: "flex", justifyContent: "center", marginTop: "20px"}}>
                <Grid grid={grid} />
            </div>
        </div>
    )
}

export default App
