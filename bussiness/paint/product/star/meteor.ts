import * as zrender from "zrender";
import config from "@/config";

// 流星
export default (x: number, y: number, rotation: number, toX: number, toY: number, sX: number, sY: number, finishCallBack: () => void) => {
    let meteor = new zrender.Droplet({
        shape: {
            cx: x,
            cy: y,
            width: config.meteorStarWidth,
            height: config.meteorStarHeight
        },
        style: {
            fill: new zrender.LinearGradient(1, 1, 0, 0, [
                {
                    offset: 0.3,
                    color: "rgb(252, 244, 194)"
                },
                {
                    offset: 0.6,
                    color: "rgb(252, 244, 214, 0.8)"
                },
                {
                    offset: 0.8,
                    color: "rgba(255, 255, 255, 0.6)"
                }
            ]),
            shadowBlur: 1,
            shadowColor: "rgba(255, 255, 255, 0.6)",
            shadowOffsetX: sX,
            shadowOffsetY: sY
        },
        rotation: - rotation / 180 * Math.PI,
        originX: x,
        originY: y,
        zlevel: config.dynamicZLevel,
        cursor: "default"
    });
    meteor.animateTo({
        shape: {
            cx: toX,
            cy: toY
        },
    }, {
        duration: config.meteorStarAnimateDuration,
        done: finishCallBack
    });
    return meteor;
}