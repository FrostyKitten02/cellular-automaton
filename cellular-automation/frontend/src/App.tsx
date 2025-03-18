import Grid from "./Grid";
// @ts-ignore
import {ResetGrid, Simulate, Step, Increment, InitGrid, Init2, StopSimulation, EditGrid} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";
import {useEffect, useState} from "react";
import {EventsOff, EventsOn} from "../wailsjs/runtime";


function App() {
    const [grid, setGrid] = useState<main.Grid | null>(null);
    const [canEdit, setCanEdit] = useState<boolean>(true);
    const [gridEdited, setGridEdited] = useState<boolean>(false);

    useEffect(() => {
        InitGrid(20,20).then( res => {
           setGrid(res);
        });

        EventsOn("simulation_stream", (newData: main.Grid) => {
            setGrid(newData)
        });

        return () => {
            EventsOff("simulation_stream")
        }
    }, [])

    //row = y
    //col = x
    const onCellClick = (row: number, col: number) => {
        if (!canEdit) {
            return;
        }
        console.log("Cell click called on cell with col(x)=",col," row(y)=",row)
        setGrid(grid => {
            if (grid?.Cells == null) {
                return grid;
            }
            grid.Cells[row][col].cellType = "ALIVE";
            // return {...grid};
            setGridEdited(true);
            return new main.Grid({...grid})
        })
    }


    if (grid == null) {
        return (
            <div>
                LOADING...
            </div>
        )
    }

    const saveEditAndCallFunc = async <I, O>(i: I, f: (input: I) => Promise<O>) => {
        if (gridEdited) {
            await EditGrid(grid);
            return await f(i);
        }

        return f(i)
    };

    return (
        <div>
            <button className="btn" onClick={() => {
                saveEditAndCallFunc(null, Simulate)
                    .then(() => {
                        setCanEdit(false);
                    })
            }}>
                Simulate
            </button>
            <button className="btn" onClick={() => {
                //probably don't need to save here??
                saveEditAndCallFunc(null, StopSimulation)
                    .then((res) => {
                        setGrid(res);
                        setCanEdit(true);
                    })
            }}>
                Stop
            </button>
            <button className="btn" onClick={() => {
                saveEditAndCallFunc(null, Step).then((res) => {
                    setGrid(res);
                })
            }}>
                Step
            </button>
            <button className="btn" onClick={() => {
                ResetGrid().then((res: main.Grid) => {
                    //TODO implement and render
                    console.log("resolved reset")
                })
            }}>
                Reset
            </button>
            <div style={{display: "flex", justifyContent: "center", marginTop: "20px"}}>
                <Grid grid={grid} onCellClick={onCellClick}/>
            </div>
        </div>
    )
}

export default App
