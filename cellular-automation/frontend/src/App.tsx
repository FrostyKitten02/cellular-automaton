import Grid from "./Grid";
// @ts-ignore
import {EditGrid, Init, ResetGrid, Simulate, Step, StopSimulation} from "../wailsjs/go/main/App";
import {model} from "../wailsjs/go/models";
import {useEffect, useState} from "react";
import {EventsOn} from "../wailsjs/runtime";


function App() {
    const [grid, setGrid] = useState<model.Grid | null>(null);
    const [canEdit, setCanEdit] = useState<boolean>(true);
    const [gridEdited, setGridEdited] = useState<boolean>(false);
    const [alertMessage, setAlertMessage] = useState<string>("")


    useEffect(() => {
        Init(85,38).then( res => {
           setGrid(res);
        });

        const removeListener = EventsOn("simulation_stream", (newData: model.Grid) => {
            setGrid(newData)
        });

        return () => {
            removeListener()
        }
    }, [])

    //row = y
    //col = x
    const onCellClick = (row: number, col: number) => {
        if (!canEdit) {
            return;
        }
        console.log("Cell click called on cell with col(x)=", col, " row(y)=", row)
        setGrid(grid => {
            if (grid?.Cells == null) {
                return grid;
            }
            grid.Cells[row][col].cellType = "ALIVE";
            // return {...grid};
            setGridEdited(true);
            return new model.Grid({...grid})
        })
    }


    if (grid == null) {
        return (
            <div className="app-main">
                LOADING...
            </div>
        )
    }

    const saveEditAndCallFunc: <I, O>(i: I, f: (input: I) => Promise<O>) => Promise<O | null> = async <I, O>(i: I, f: (input: I) => Promise<O>) => {
        if (gridEdited) {
            try {
                await EditGrid(grid)
            } catch (err) {
                return null;
            }

            return f(i);
        }

        return f(i)
    };

    return (
        <div className="app-wrapper">
            <div className="app-main">
                <button className="btn" onClick={() => {
                    saveEditAndCallFunc(null, Simulate)
                        ?.then(() => {
                            setCanEdit(false);
                        })
                        .catch(err => {
                            setAlertMessage(err);
                        })
                }}>
                    Simulate
                </button>
                <button className="btn" onClick={() => {
                    //probably don't need to save here??
                    saveEditAndCallFunc(null, StopSimulation)
                        ?.then((res) => {
                            setGrid(res);
                            setCanEdit(true);
                        })
                }}>
                    Stop
                </button>
                <button className="btn" onClick={() => {
                    saveEditAndCallFunc(null, Step)
                        ?.then((res) => {
                            setGrid(res);
                        })
                        .catch(err => {
                            setAlertMessage(err);
                        })
                }}>
                    Step
                </button>
                <button className="btn" onClick={() => {
                    ResetGrid().then((res: model.Grid) => {
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
            <div className="alert" style={{display: alertMessage != "" ? "block" : "none"}}>
                <div>
                    <p>{alertMessage}</p>
                    <button className="btn-sm" onClick={() => {setAlertMessage("")}}>
                        Cancel
                    </button>
                </div>
            </div>
        </div>
    )
}

export default App
