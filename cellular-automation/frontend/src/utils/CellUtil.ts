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
            ["EMPTY", "#0aff00"],
            ["WALL", "#000000"],
            ["SAND", "#ffff00"],
            ["WOOD", "#964B00"],
            ["FIRE", "#ff5100"],
            ["DARK_SMOKE", "#494949"],
            ["WHITE_SMOKE", "#b7b7b7"],
            ["WATER", "#024bf5"]
        ]
    )

    public static getCellColor(cell: model.Cell, gameMode: string): string {
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
            if (cell.cellType == "WATER") {
                return this.colorLuminance(color, -(cell.value - 1)/3)
            }

            return color;
        }

        return this.ERR_COLOR;
    }

    private static colorLuminance(hex: string, lum: number) {

        // validate hex string
        hex = String(hex).replace(/[^0-9a-f]/gi, '');
        if (hex.length < 6) {
            hex = hex[0]+hex[0]+hex[1]+hex[1]+hex[2]+hex[2];
        }
        lum = lum || 0;

        // convert to decimal and change luminosity
        var rgb = "#", c, i;
        for (i = 0; i < 3; i++) {
            c = parseInt(hex.substr(i*2,2), 16);
            c = Math.round(Math.min(Math.max(0, c + (c * lum)), 255)).toString(16);
            rgb += ("00"+c).substr(c.length);
        }

        return rgb;
    }
}
