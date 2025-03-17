import Grid from "./Grid";
import {ResetGrid, Simulate, Step} from "../wailsjs/go/main/App";


function App() {
    const gridSize = 60;


    return (
        <div>
            <button className="btn" onClick={() => {
                Simulate().then(res => {
                    //TODO rerender grid
                    console.log("Resolved simulate")
                })
            }}>
                Simulate
            </button>
            <button className="btn" onClick={() => {
                Step().then(res => {
                    //TODO rerender grid
                    console.log("resolved step")
                })
            }}>
                Step
            </button>
            <button className="btn" onClick={() => {
                ResetGrid().then((res) => {
                    //TODO rerender grid
                    console.log("resolved reset")
                })
            }}>
                Reset
            </button>
            <div style={{display: "flex", justifyContent: "center", marginTop: "20px"}}>
                <Grid x={gridSize} y={gridSize} />
            </div>
        </div>
    )
}

export default App
