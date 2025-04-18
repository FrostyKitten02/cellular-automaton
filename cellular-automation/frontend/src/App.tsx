import Grid from "./Grid";
// @ts-ignore
import {EditGrid, Init, ResetGrid, Simulate, Step, StopSimulation} from "../wailsjs/go/main/App";
import {model} from "../wailsjs/go/models";
import {useEffect, useState} from "react";
import {EventsOn} from "../wailsjs/runtime";
import CellUtil from "./utils/CellUtil";


const WIDTH = 85;
const HEIGHT = 38;

function App() {
    const [grid, setGrid] = useState<model.Grid | null>(null);
    const [canEdit, setCanEdit] = useState<boolean>(true);
    const [gridEdited, setGridEdited] = useState<boolean>(false);
    const [alertMessage, setAlertMessage] = useState<string>("")
    const [paintCell, setPaintCell] = useState<string>("WALL")

    const [gameMode, setGameMode] = useState<string>("CONWAY")
    const [options, setOptions] = useState<Record<string, string>>({})

    useEffect(() => {
        Init(WIDTH, HEIGHT, gameMode, options).then(res => {
            setGrid(res);
        });

        const removeListener = EventsOn("simulation_stream", (newData: model.Grid) => {
            setGrid(newData)
        });

        return () => {
            removeListener()
        }
    }, [])

    const optsChange = () => {
        Init(WIDTH, HEIGHT, gameMode, options)
            .then(res => {
                setGrid(res);
            })
            .catch(err => {
                setAlertMessage(err);
            })
    }


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
            grid.Cells[row][col].cellType = paintCell;
            grid.Cells[row][col].bornGeneration = -1;
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
                await EditGrid(grid);
                setGridEdited(false);
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
                <div
                    style={{
                        display: "flex",
                        justifyContent: "center",
                        flexDirection: "row",
                        columnGap: "10px"
                    }}
                >
                    <div style={{
                        display: "flex",
                        flexDirection: "row",
                        justifyContent: "center",
                        alignItems: "center",
                        columnGap: "10px"
                    }}>
                        <label>
                            Pain
                            <select defaultValue={"WALL"} onChange={(event) => {
                                setPaintCell(event.target.value);
                            }}>
                                <option value="WALL">Wall</option>
                                <option value="EMPTY">Empty</option>
                                <option value="SAND">Sand</option>
                                <option value="WOOD">Wood</option>
                                <option value="FIRE">Fire</option>
                                <option value="DARK_SMOKE">Dark smoke</option>
                                <option value="WHITE_SMOKE">White smoke</option>
                                <option value="WATER">Water</option>
                            </select>
                        </label>
                        <div
                            style={{
                                maxWidth: 10,
                                minWidth: 10,
                                maxHeight: 10,
                                minHeight: 10,
                                backgroundColor: CellUtil.getCellColor(new model.Cell({cellType: paintCell}), gameMode)
                            }}
                        />
                    </div>

                    <label>
                        Gamemode
                        <select defaultValue={"CONWAY"} onChange={(event) => {
                            setGameMode(event.target.value);
                        }}>
                            <option value="CONWAY">Conway</option>
                            <option value="SANDBOX">Sandbox</option>
                            <option value="1D">1D</option>
                        </select>
                    </label>


                    <div>
                        <label>
                            Condition
                            <input name="conditions" value={options.conwayCondition} onChange={(event) => {
                                const conditionString = event.target.value;
                                setOptions(opts => {
                                    return {...opts, conwayCondition: conditionString};
                                });
                            }}/>
                        </label>
                        <label>
                            Rule
                            <input name="conditions" type={"number"} min={0} max={255} value={options.rule} onChange={(event) => {
                                const rule = event.target.value;
                                setOptions(opts => {
                                    return {...opts, rule: rule};
                                });
                            }}/>
                        </label>
                        <label>
                            Wall %
                            <input type={"number"} min={0} max={100} name="alivePercent" value={options.alivePercent}
                                   onChange={(event) => {
                                       const val = event.target.value;
                                       setOptions(opts => {
                                           return {...opts, alivePercent: val}
                                       })
                                   }}/>
                        </label>
                        <label>
                            Pregen cave
                            <input type={"checkbox"} checked={options.pregenCave == "true"}
                                   onChange={(event) => {
                                       const val = event.target.checked
                                       setOptions(opts => {
                                           return {...opts, pregenCave: val.toString()}
                                       })
                                   }}
                            />
                        </label>
                    </div>
                </div>

                <button className="btn" onClick={() => {
                    optsChange()
                }}>
                    Apply
                </button>
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
                    <Grid grid={grid} onCellClick={onCellClick} gameMode={gameMode}/>
                </div>
            </div>
            <div className="alert" style={{display: alertMessage != "" ? "block" : "none"}}>
                <div>
                    <p>{alertMessage}</p>
                    <button className="btn-sm" onClick={() => {
                        setAlertMessage("")
                    }}>
                        Cancel
                    </button>
                </div>
            </div>
        </div>
    )
}

export default App
