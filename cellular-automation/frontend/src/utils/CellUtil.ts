import {model} from "../../wailsjs/go/models";


export default class CellUtil {
    private constructor() {}

    private static readonly ERR_COLOR = "#ff0000" //if red is displayed we have a problem

    private static readonly CONWAY_COLORS = new Map<string, string>(
        [
            ["EMPTY", "#ffffff"],
            ["WALL", "#000000"],
        ]
    )

    private static readonly SANDBOX_COLORS = new Map<string, string>(
        [
            ["EMPTY", "#ffffff"],
            ["WALL", "#000000"],
            ["SAND", "#ffff00"],
            ["WOOD", "#632f02"]
        ]
    )

    public static getCellColor(cell: model.Cell, gameMode: string): string {
        console.log("GETTING COLOR FOR: ", cell)
        if (cell.cellType == undefined || cell.cellType == "") {
            return this.ERR_COLOR;
        }

        if (gameMode == "CONWAY") {
            return this.getConwayCellColor(cell);
        }

        if (gameMode == "SANDBOX") {
            //TODO seperate handler
            return this.getSandboxCellColor(cell);
        }

        return this.ERR_COLOR
    }


    private static getConwayCellColor(cell: model.Cell): string {
        const color = this.CONWAY_COLORS.get(cell.cellType!); //cant be undefined at this point!!

        if (color != undefined) {
            return color;
        }

        return this.ERR_COLOR;
    }

    private static getSandboxCellColor(cell: model.Cell): string {
        const color = this.SANDBOX_COLORS.get(cell.cellType!); //cant be undefined at this point!!

        if (color != undefined) {
            return color;
        }

        return this.ERR_COLOR;
    }
}
