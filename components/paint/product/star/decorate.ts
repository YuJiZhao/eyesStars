import * as zrender from "zrender";
import config from "@/config";
import { getDoubleRandom, getIntRandom } from "@/utils/help";

// 背景装饰星星
export default (x: number, y: number) => {
    return new zrender.Star({
        shape: {
            cx: x,
            cy: y,
            n: getIntRandom(config.decorateStarN.min, config.decorateStarN.max),
            r: getDoubleRandom(config.decorateStarR.min, config.decorateStarR.max),
            r0: getDoubleRandom(config.decorateStarR0.min, config.decorateStarR0.max),
        },
        style: {
            fill: `rgb(
                ${getIntRandom(config.decorateStarFill.min, config.decorateStarFill.max)}, 
                ${getIntRandom(config.decorateStarFill.min, config.decorateStarFill.max)}, 
                ${getIntRandom(config.decorateStarFill.min, config.decorateStarFill.max)}
            )`,
            shadowBlur: 2,
            shadowColor: "#FFF",
            opacity: getDoubleRandom(config.decorateStarOpacity.min, config.decorateStarOpacity.max),
        },
        cursor: "default",
        zlevel: config.staticZLevel
    });
}